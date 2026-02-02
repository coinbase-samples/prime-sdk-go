/**
 * Copyright 2026-present Coinbase Global, Inc.
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
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/prime-sdk-go/client"
)

type GetFcmRiskLimitsRequest struct {
	EntityId string `json:"entity_id"`
}

type GetFcmRiskLimitsResponse struct {
	CfmRiskLimit            string                   `json:"cfm_risk_limit"`
	CfmRiskLimitUtilization string                   `json:"cfm_risk_limit_utilization"`
	CfmTotalMargin          string                   `json:"cfm_total_margin"`
	CfmDeltaOte             string                   `json:"cfm_delta_ote"`
	CfmUnsettledRealizedPnl string                   `json:"cfm_unsettled_realized_pnl"`
	Request                 *GetFcmRiskLimitsRequest `json:"-"`
}

func (s *futuresServiceImpl) GetFcmRiskLimits(
	ctx context.Context,
	request *GetFcmRiskLimitsRequest,
) (*GetFcmRiskLimitsResponse, error) {

	path := fmt.Sprintf("/entities/%s/futures/risk_limits", request.EntityId)

	response := &GetFcmRiskLimitsResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
