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

package main

type sakuraCloudResourceRenderable interface {
	getTerraformName() string
	getPluralResourceName() string
	getAdditionalFields() map[string]string
	getAllowEmptyValues() []string
	isUseGlobalZone() bool
	isUseScopeFilter() bool
}

type sakuraCloudResource struct {
	terraformName      string
	pluralResourceName string
	allowEmptyValues   []string
	additionalFields   map[string]string
	globalZone         bool
	useScopeFilter     bool
}

func (r sakuraCloudResource) getTerraformName() string {
	return r.terraformName
}

func (r sakuraCloudResource) getPluralResourceName() string {
	return r.pluralResourceName
}

func (r sakuraCloudResource) getAdditionalFields() map[string]string {
	return r.additionalFields
}

func (r sakuraCloudResource) getAllowEmptyValues() []string {
	return r.allowEmptyValues
}

func (r sakuraCloudResource) isUseGlobalZone() bool {
	return r.globalZone
}

func (r sakuraCloudResource) isUseScopeFilter() bool {
	return r.useScopeFilter
}
