// Copyright 2019 Kazumichi Yamamoto.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sakuracloud

import (
	"context"
)

var zones = []string{"is1a", "is1b", "tk1a", "tk1v"}

type findFunc func(ctx context.Context, zone string) ([]interface{}, error)

func findResourcePerZone(ctx context.Context, findFunc findFunc) ([]interface{}, error) {
	var res []interface{}
	for _, zone := range zones {
		searched, err := findFunc(ctx, zone)
		if err != nil {
			return nil, err
		}
		for _, v := range searched {
			res = append(res, v)
		}
	}
	return res, nil
}
