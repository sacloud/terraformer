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

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

const pathForGenerateFiles = "/providers/sakuracloud/"
const serviceTemplate = `
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
	{{ if .isUseScopeFilter }}
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/types"{{end}}
)

var {{.resource}}AllowEmptyValues = []string{"{{join .allowEmptyValues "\",\"" }}"}

var {{.resource}}AdditionalFields = map[string]interface{}{
	{{ range $key,$value := .additionalFields}}
	"{{$key}}":			"{{$value}}",{{end}}
}

type {{.titleResourceName}}Generator struct {
	SakuraCloudService
}

// Create for each TerraformResource
func (g {{.titleResourceName}}Generator) createResources(searched []interface{}) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for i, resource := range searched {
		obj := resource.(*sacloudResource)
		if obj.isFailed() {
			continue
		}
		resourceID := obj.id()
		resourceName := fmt.Sprintf("%s-%03d-%s", "{{.resource}}", i, obj.name())
		resources = append(resources, terraform_utils.NewResource(
			resourceID,
			resourceName,
			"{{.terraformName}}",
			"sakuracloud",
			obj.attributes(),
			{{.resource}}AllowEmptyValues,
			{{.resource}}AdditionalFields,
		))
	}

	return resources
}

// Generate TerraformResources from SakuraCloud API,
// from each {{.resource}} create 1 TerraformResource
// Need {{.resource}} name as ID for terraform resource
func (g *{{.titleResourceName}}Generator) InitResources() error {
	caller := g.NewClient()
	ctx := context.Background()
	op := sacloud.New{{.titleResourceName}}Op(caller)

	{{ if .isUseGlobalZone }}
	searched, err := op.Find(ctx, &sacloud.FindCondition{
		{{ if .isUseScopeFilter }}
		Filter: search.Filter{
			search.Key(keys.Scope): search.AndEqual(string(types.Scopes.User)),
		},{{ end }}
	})
	if err != nil {
		return err
	}
	var resources []interface{}
	for _ , v := range searched.{{.pluralResourceName}} {
		resources = append(resources, &sacloudResource{
			resource: v,
			isGlobal: true,
		})
	}
	g.Resources = g.createResources(resources)
	{{ else }}
	resources, err := findResourcePerZone(ctx, func(ctx context.Context, zone string) ([]interface{}, error) {
		searched, err := op.Find(ctx, zone, &sacloud.FindCondition{
			{{ if .isUseScopeFilter }}
			Filter: search.Filter{
				search.Key(keys.Scope): search.AndEqual(string(types.Scopes.User)),
			},{{ end }}
		})
		if err != nil {
			return nil, err
		}
		var res []interface{}
		for _, v := range searched.{{.pluralResourceName}} {
			res = append(res, &sacloudResource{
				resource: v,
				zone: zone,
			})
		}
		return res, nil
	})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(resources)
	{{ end }}

	return nil
}
`

func main() {
	funcMap := template.FuncMap{
		"title":   strings.Title,
		"toLower": strings.ToLower,
		"join":    strings.Join,
	}

	for resourceName, resourceDef := range terraformResources {

		var titleResourceName string
		switch resourceName {
		case "cdrom":
			titleResourceName = "CDROM"
		case "dns":
			titleResourceName = "DNS"
		case "gslb":
			titleResourceName = "GSLB"
		case "nfs":
			titleResourceName = "NFS"
		case "sim":
			titleResourceName = "SIM"
		case "sshKey":
			titleResourceName = "SSHKey"
		case "vpcRouter":
			titleResourceName = "VPCRouter"
		default:
			titleResourceName = strings.Title(resourceName)
		}

		var tpl bytes.Buffer
		t := template.Must(template.New("resourceName.go").Funcs(funcMap).Parse(serviceTemplate))
		err := t.Execute(&tpl, map[string]interface{}{
			"resource":           resourceName,
			"titleResourceName":  titleResourceName,
			"pluralResourceName": resourceDef.getPluralResourceName(),
			"terraformName":      resourceDef.getTerraformName(),
			"additionalFields":   resourceDef.getAdditionalFields(),
			"allowEmptyValues":   resourceDef.getAllowEmptyValues(),
			"isUseGlobalZone":    resourceDef.isUseGlobalZone(),
			"isUseScopeFilter":   resourceDef.isUseScopeFilter(),
		})
		if err != nil {
			log.Printf("rendering template with %s is failed: %s", resourceName, err)
			continue
		}
		rootPath, _ := os.Getwd()
		currentPath := rootPath + pathForGenerateFiles
		err = os.MkdirAll(currentPath, os.ModePerm)
		if err != nil {
			log.Print(resourceName, err)
			continue
		}
		err = ioutil.WriteFile(currentPath+"/"+resourceName+"_gen.go", codeFormat(tpl.Bytes()), os.ModePerm)
		if err != nil {
			log.Printf("writing file %s_gen.go is failed: %s", resourceName, err)
			continue
		}
	}
}

func codeFormat(src []byte) []byte {
	code, err := format.Source(src)
	if err != nil {
		log.Println(err)
	}
	return code
}
