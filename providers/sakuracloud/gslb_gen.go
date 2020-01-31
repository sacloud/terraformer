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

var gslbAllowEmptyValues = []string{""}

var gslbAdditionalFields = map[string]interface{}{}

type GSLBGenerator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g GSLBGenerator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "gslb", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"sakuracloud_gslb",
			"sakuracloud",
			obj.attributes(),
			gslbAllowEmptyValues,
			gslbAdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each gslb create 1 TerraformResource
// Need gslb name as ID for terraform resource
func (g *GSLBGenerator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.NewGSLBOp(caller)

	searched, err := op.Find(ctx, &sacloud.FindCondition{})
	if err != nil {
		return err
	}
	var resources []interface{}
	for _, v := range searched.GSLBs {
		resources = append(resources, &sacloudResource{
			resource: v,
			isGlobal: true,
		})
	}
	g.Resources = g.createResources(resources)

	return nil
}
