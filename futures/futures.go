/**
 * Copyright 2025-present Coinbase Global, Inc.
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

package futures

import (
	"context"

	"github.com/coinbase-samples/prime-sdk-go/client"
)

type FuturesService interface {
	SetAutoSweep(ctx context.Context, request *SetAutoSweepRequest) (*SetAutoSweepResponse, error)
	GetEntityFcmBalance(ctx context.Context, request *GetEntityFcmBalanceRequest) (*GetEntityFcmBalanceResponse, error)
	GetEntityPositions(ctx context.Context, request *GetEntityPositionsRequest) (*GetEntityPositionsResponse, error)
	ListEntityFuturesSweeps(ctx context.Context, request *ListEntityFuturesSweepsRequest) (*ListEntityFuturesSweepsResponse, error)
	CancelEntityFuturesSweep(ctx context.Context, request *CancelEntityFuturesSweepRequest) (*CancelEntityFuturesSweepResponse, error)
	ScheduleEntityFuturesSweep(ctx context.Context, request *ScheduleEntityFuturesSweepRequest) (*ScheduleEntityFuturesSweepResponse, error)
}

func NewFuturesService(c client.RestClient) FuturesService {
	return &futuresServiceImpl{client: c}
}

type futuresServiceImpl struct {
	client client.RestClient
}
