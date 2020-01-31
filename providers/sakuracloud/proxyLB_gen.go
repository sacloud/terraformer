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

// AUTO-GENERATED CODE. DO NOT EDIT.

package sakuracloud

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

var proxyLBAllowEmptyValues = []string{""}

var proxyLBAdditionalFields = map[string]interface{}{}

type ProxyLBGenerator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g ProxyLBGenerator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "proxyLB", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"sakuracloud_proxylb",
			"sakuracloud",
			obj.attributes(),
			proxyLBAllowEmptyValues,
			proxyLBAdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each proxyLB create 1 TerraformResource
// Need proxyLB name as ID for terraform resource
func (g *ProxyLBGenerator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.NewProxyLBOp(caller)

	searched, err := op.Find(ctx, &sacloud.FindCondition{})
	if err != nil {
		return err
	}
	var resources []interface{}
	for _, v := range searched.ProxyLBs {
		resources = append(resources, &sacloudResource{
			resource: v,
			isGlobal: true,
		})
	}
	g.Resources = g.createResources(resources)

	return nil
}
