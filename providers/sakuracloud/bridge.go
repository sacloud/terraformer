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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

var bridgeAllowEmptyValues = []string{""}

var bridgeAdditionalFields = map[string]string{}

type BridgeGenerator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g BridgeGenerator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "bridge", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"sakuracloud_bridge",
			"sakuracloud",
			obj.attributes(),
			bridgeAllowEmptyValues,
			bridgeAdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each bridge create 1 TerraformResource
// Need bridge name as ID for terraform resource
func (g *BridgeGenerator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.NewBridgeOp(caller)

	// bridgeのみどのゾーンでリクエストしても同一リソースが返る
	searched, err := op.Find(ctx, sacloud.APIDefaultZone, &sacloud.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return err
	}
	var resources []interface{}
	for _, v := range searched.Bridges {
		resources = append(resources, &sacloudResource{
			resource: v,
			isGlobal: true,
		})
	}
	g.Resources = g.createResources(resources)

	g.PopulateIgnoreKeys()
	return nil
}
