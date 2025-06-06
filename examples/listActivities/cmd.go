/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/activities"
	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/credentials"
)

func main() {

	credentials := &credentials.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("PRIME_CREDENTIALS")), credentials); err != nil {
		log.Fatalf("unable to deserialize prime credentials JSON: %v", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		log.Fatalf("unable to load default http client: %v", err)
	}

	client := client.NewRestClient(credentials, httpClient)

	activitiesSvc := activities.NewActivitiesService(client)
	request := &activities.ListActivitiesRequest{
		PortfolioId: credentials.PortfolioId,
	}

	response, err := activitiesSvc.ListActivities(context.Background(), request)
	if err != nil {
		log.Fatalf("unable to list activities: %v", err)
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("error marshaling response to JSON: %v", err))
	}
	fmt.Println(string(jsonResponse))
}
