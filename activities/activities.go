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

package activities

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
	"github.com/coinbase-samples/prime-sdk-go/model"
)

type ActivitiesService interface {
	ListActivities(ctx context.Context, request *ListActivitiesRequest) (*ListActivitiesResponse, error)
	GetActivity(ctx context.Context, request *GetActivityRequest) (*GetActivityResponse, error)
	ListEntityActivities(ctx context.Context, request *ListEntityActivitiesRequest) (*ListEntityActivitiesResponse, error)
	GetEntityActivity(ctx context.Context, request *GetEntityActivityRequest) (*GetEntityActivityResponse, error)
	ServiceConfig() *model.ServiceConfig
}

func NewActivitiesService(c client.RestClient) ActivitiesService {
	return &activitiesServiceImpl{
		client:        c,
		serviceConfig: model.DefaultServiceConfig(),
	}
}

func NewActivitiesServiceWithConfig(c client.RestClient, config *model.ServiceConfig) ActivitiesService {
	return &activitiesServiceImpl{
		client:        c,
		serviceConfig: config,
	}
}

type activitiesServiceImpl struct {
	client        client.RestClient
	serviceConfig *model.ServiceConfig
}

func (s *activitiesServiceImpl) ServiceConfig() *model.ServiceConfig {
	return s.serviceConfig
}
