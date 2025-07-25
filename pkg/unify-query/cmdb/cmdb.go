// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package cmdb

import (
	"context"
)

type CMDB interface {
	// QueryResourceMatcher 获取目标的关键维度和值（instant 查询）
	QueryResourceMatcher(ctx context.Context, lookBackDelta, spaceUid string, ts int64, target, source Resource, indexesMatcher, expandMatcher Matcher, expandShow bool, pathResource []Resource) (Resource, Matcher, []string, Resource, Matchers, error)

	// QueryResourceMatcherRange 获取目标的关键维度和值（query_range 查询）
	QueryResourceMatcherRange(ctx context.Context, lookBackDelta, spaceUid string, step string, startTs, endTs int64, target, source Resource, indexesMatcher, expandMatcher Matcher, expandShow bool, pathResource []Resource) (Resource, Matcher, []string, Resource, []MatchersWithTimestamp, error)
}
