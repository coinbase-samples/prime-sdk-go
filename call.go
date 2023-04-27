package prime

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type apiRequest struct {
	url                    string
	httpMethod             string
	body                   []byte
	expectedHttpStatusCode int
	client                 Client
}

type apiResponse struct {
	request        *apiRequest
	body           []byte
	httpStatusCode int
	httpStatusMsg  string
	err            error
	errorMessage   *ErrorMessage
}

func post(
	ctx context.Context,
	client Client,
	url string,
	request,
	response interface{},
) error {
	return call(ctx, client, url, http.MethodPost, http.StatusOK, request, response)
}

func get(
	ctx context.Context,
	client Client,
	url string,
	request,
	response interface{},
) error {
	return call(ctx, client, url, http.MethodGet, http.StatusOK, request, response)
}

func call(
	ctx context.Context,
	client Client,
	url,
	httpMethod string,
	expectedHttpStatusCode int,
	request,
	response interface{},
) error {

	if client.Credentials == nil {
		return errors.New("credentials not set")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp := makeCall(
		ctx,
		&apiRequest{
			url:                    url,
			httpMethod:             httpMethod,
			body:                   body,
			expectedHttpStatusCode: expectedHttpStatusCode,
			client:                 client,
		},
	)

	if resp.err != nil {
		return err
	}

	if err := json.Unmarshal(resp.body, response); err != nil {
		return err
	}

	return nil
}

func makeCall(ctx context.Context, request *apiRequest) *apiResponse {

	response := &apiResponse{
		request: request,
	}

	parsedUrl, err := url.Parse(request.url)
	if err != nil {
		response.err = fmt.Errorf("URL: %s - %w", request.url, err)
		return response
	}

	t := time.Now().Unix()

	req, err := http.NewRequestWithContext(ctx, request.httpMethod, request.url, bytes.NewReader(request.body))
	if err != nil {
		response.err = err
		return response
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CB-ACCESS-KEY", request.client.Credentials.AccessKey)
	req.Header.Add("X-CB-ACCESS-PASSPHRASE", request.client.Credentials.Passphrase)
	req.Header.Add("X-CB-ACCESS-SIGNATURE", sign(parsedUrl.Path, string(request.body), request.httpMethod, request.client.Credentials.SigningKey, t))
	req.Header.Add("X-CB-ACCESS-TIMESTAMP", fmt.Sprintf("%d", t))

	client := http.Client{
		Transport: request.client.Transport,
		Timeout:   request.client.Timeout,
	}

	res, err := client.Do(req)
	if err != nil {
		response.err = err
		return response
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.err = err
		return response
	}

	response.body = resBody
	response.httpStatusCode = res.StatusCode
	response.httpStatusMsg = res.Status

	if request.expectedHttpStatusCode > 0 && res.StatusCode != request.expectedHttpStatusCode {

		var errMsg ErrorMessage
		if strings.Contains(string(response.body), "message") {
			// We do not want to be distracted by an error here. If the message
			// cannot be parsed, the body is available as well.
			_ = json.Unmarshal(response.body, &errMsg)
			response.errorMessage = &errMsg
		}

		responseMsg := string(resBody)

		if response.errorMessage != nil && len(response.errorMessage.Value) > 0 {
			responseMsg = response.errorMessage.Value
		}

		response.err = fmt.Errorf(
			"expected status code: %d - received: %d - status msg: %s - url %s - msg: %s",
			request.expectedHttpStatusCode,
			res.StatusCode,
			res.Status,
			request.url,
			responseMsg,
		)
	}

	return response
}
