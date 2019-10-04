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

	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

var noteAllowEmptyValues = []string{""}

var noteAdditionalFields = map[string]interface{}{}

type NoteGenerator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g NoteGenerator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "note", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"sakuracloud_note",
			"sakuracloud",
			obj.attributes(),
			noteAllowEmptyValues,
			noteAdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each note create 1 TerraformResource
// Need note name as ID for terraform resource
func (g *NoteGenerator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.NewNoteOp(caller)

	searched, err := op.Find(ctx, &sacloud.FindCondition{

		Filter: search.Filter{
			search.Key(keys.Scope): search.AndEqual(string(types.Scopes.User)),
		},
		Count: 10000,
	})
	if err != nil {
		return err
	}
	var resources []interface{}
	for _, v := range searched.Notes {
		resources = append(resources, &sacloudResource{
			resource: v,
			isGlobal: true,
		})
	}
	g.Resources = g.createResources(resources)

	return nil
}
