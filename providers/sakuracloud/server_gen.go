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

var serverAllowEmptyValues = []string{""}

var serverAdditionalFields = map[string]interface{}{}

type ServerGenerator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g ServerGenerator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "server", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"sakuracloud_server",
			"sakuracloud",
			obj.attributes(),
			serverAllowEmptyValues,
			serverAdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each server create 1 TerraformResource
// Need server name as ID for terraform resource
func (g *ServerGenerator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.NewServerOp(caller)

	resources, err := findResourcePerZone(ctx, func(ctx context.Context, zone string) ([]interface{}, error) {
		searched, err := op.Find(ctx, zone, &sacloud.FindCondition{

			Count: 10000,
		})
		if err != nil {
			return nil, err
		}
		var res []interface{}
		for _, v := range searched.Servers {
			res = append(res, &sacloudResource{
				resource: v,
				zone:     zone,
			})
		}
		return res, nil
	})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(resources)

	return nil
}
