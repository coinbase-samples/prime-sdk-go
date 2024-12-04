# Prime Go SDK README

[![GoDoc](https://godoc.org/github.com/coinbase-samples/prime-sdk-go?status.svg)](https://godoc.org/github.com/coinbase-samples/prime-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/coinbase-samples/prime-sdk-go)](https://goreportcard.com/report/coinbase-samples/prime-sdk-go)

## Overview

The *Prime Go SDK* is a sample libary that demonstrates the structure of a [Coinbase Prime](https://prime.coinbase.com/) driver for
the [REST APIs](https://docs.cloud.coinbase.com/prime/reference).

## License

The *Prime Go SDK* sample library is free and open source and released under the [Apache License, Version 2.0](LICENSE).

The application and code are only available for demonstration purposes.

## Usage

To use the *Prime Go SDK*, initialize the [Credentials](credentials.go) struct and create a new client. The Credentials struct is JSON
enabled. Ensure that Prime API credentials are stored in a secure manner.

```
primeCredentials, err := credentials.ReadEnvCredentials("PRIME_CREDENTIALS")
if err != nil {
    log.Fatalf("unable to load prime credentials: %v", err)
}

httpClient, err := client.DefaultHttpClient()
if err != nil {
    log.Fatalf("unable to load default http client: %v", err)
}

client := prime.NewRestClient(primeCredentials, httpClient)
```

The credentials.ReadEvnCredentials is a convenience function to read the credentials from an environment variable and deserialize the JSON structure. Use credentials.UnmarshalCredentials,
if pulled from a different source. The JSON format expected by both is:

```
{
  "accessKey": "",
  "passphrase": "",
  "signingKey": "",
  "portfolioId": "",
  "svcAccountId": "",
  "entityId": ""
}
```

Coinbase Prime API credentials can be created in the Prime web console under Settings -> APIs. Entity ID can be retrieved by calling [Get Portfolio](https://docs.cdp.coinbase.com/prime/reference/primerestapi_getportfolio).

Once the client is initialized, instantiate a service to make the desired call. For example, to list portfolios, create the service, pass in the request object, check for an error, and if nil, process the response.

```
service := portfolios.NewPortfoliosService(client)

response, err := service.ListPortfolios(ctx, &prime.ListPortfoliosRequest{})
```

## Build

To build the sample library, ensure that [Go](https://go.dev/) 1.19+ is installed and then run:

```bash
go build ./...
```
