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
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
)

type sacloudResource struct {
	resource interface{}
	isGlobal bool
	zone     string
}

func (r *sacloudResource) id() string {
	return r.resource.(accessor.ID).GetID().String()
}

type nameAccessor interface {
	GetName() string
}

func (r *sacloudResource) name() string {
	if v, ok := r.resource.(nameAccessor); ok {
		return v.GetName()
	}
	return ""
}

func (r *sacloudResource) attributes() map[string]string {
	attrs := make(map[string]string)
	if v, ok := r.resource.(nameAccessor); ok {
		attrs["name"] = v.GetName()
	}
	if !r.isGlobal {
		attrs["zone"] = r.zone
	}
	return attrs
}

func (r *sacloudResource) isFailed() bool {
	if v, ok := r.resource.(accessor.Availability); ok {
		return v.GetAvailability().IsFailed()
	}

	return false
}
