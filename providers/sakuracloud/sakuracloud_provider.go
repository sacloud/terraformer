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
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/zclconf/go-cty/cty"
)

type SakuraCloudProvider struct {
	terraform_utils.Provider
	token  string
	secret string
}

const sakuraCloudProviderVersion = "~>v1.16.0"

func (p SakuraCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"archive": {
			"icon": []string{"icon_id", "id"},
		},
		"autoBackup": {
			"disk": []string{"disk_id", "id"},
			"icon": []string{"icon_id", "id"},
		},
		"cdrom": {
			"icon": []string{"icon_id", "id"},
		},
		"database": {
			"switch": []string{"switch_id", "id"},
			"icon":   []string{"icon_id", "id"},
		},
		"disk": {
			"archive": []string{"source_archive_id", "id"},
			"disk":    []string{"source_disk_id", "id"},
			"icon":    []string{"icon_id", "id"},
		},
		"dns": {
			"icon": []string{"icon_id", "id"},
		},
		"gslb": {
			"icon": []string{"icon_id", "id"},
		},
		"internet": {
			"icon": []string{"icon_id", "id"},
		},
		"loadBalancer": {
			"switch": []string{"switch_id", "id"},
			"icon":   []string{"icon_id", "id"},
		},
		"mobileGateway": {
			"switch": []string{"switch_id", "id"},
			"icon":   []string{"icon_id", "id"},
		},
		"nfs": {
			"switch": []string{"switch_id", "id"},
			"icon":   []string{"icon_id", "id"},
		},
		"note": {
			"icon": []string{"icon_id", "id"},
		},
		"privateHost": {
			"icon": []string{"icon_id", "id"},
		},
		"proxyLB": {
			"icon": []string{"icon_id", "id"},
		},
		"server": {
			"disk":         []string{"disks", "id"},
			"cdrom":        []string{"cdrom_id", "id"},
			"privateHost":  []string{"private_host_id", "id"},
			"switch":       []string{"additional_nics", "id"},
			"packetFilter": []string{"packet_filter_ids", "id"},
			"icon":         []string{"icon_id", "id"},
		},
		"sim": {
			"mobileGateway": []string{"mobile_gateway_id", "id"},
			"icon":          []string{"icon_id", "id"},
		},
		"simpleMonitor": {
			"icon": []string{"icon_id", "id"},
		},
		"switch": {
			"icon": []string{"icon_id", "id"},
		},
		"vpcRouter": {
			"icon":   []string{"icon_id", "id"},
			"switch": []string{"switch_id", "id"},
		},
	}
}

func (p SakuraCloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"sakuracloud": map[string]interface{}{},
		},
	}
}

func (p *SakuraCloudProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"token":  cty.StringVal(p.token),
		"secret": cty.StringVal(p.secret),
	})
}

// Init SakuraCloudProvider with API token
func (p *SakuraCloudProvider) Init(args []string) error {
	if args[0] != "" {
		p.token = args[0]
	} else {
		if token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN"); token != "" {
			p.token = token
		} else {
			return errors.New("token requirement")
		}
	}

	if args[1] != "" {
		p.secret = args[1]
	} else {
		if secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"); secret != "" {
			p.secret = secret
		} else {
			return errors.New("secret requirement")
		}
	}

	return nil
}

func (p *SakuraCloudProvider) GetName() string {
	return "sakuracloud"
}

func (p *SakuraCloudProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token":  p.token,
		"secret": p.secret,
	})
	return nil
}

// GetSupportedService return map of support service for SakuraCloud
func (p *SakuraCloudProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"archive":       &ArchiveGenerator{},
		"autoBackup":    &AutoBackupGenerator{},
		"bridge":        &BridgeGenerator{},
		"cdrom":         &CDROMGenerator{},
		"database":      &DatabaseGenerator{},
		"disk":          &DiskGenerator{},
		"dns":           &DNSGenerator{},
		"gslb":          &GSLBGenerator{},
		"icon":          &IconGenerator{},
		"internet":      &InternetGenerator{},
		"loadBalancer":  &LoadBalancerGenerator{},
		"mobileGateway": &MobileGatewayGenerator{},
		"nfs":           &NFSGenerator{},
		"note":          &NoteGenerator{},
		"packetFilter":  &PacketFilterGenerator{},
		"privateHost":   &PrivateHostGenerator{},
		"proxyLB":       &ProxyLBGenerator{},
		"server":        &ServerGenerator{},
		"sim":           &SIMGenerator{},
		"simpleMonitor": &SimpleMonitorGenerator{},
		"sshKey":        &SSHKeyGenerator{},
		"switch":        &SwitchGenerator{},
		"vpcRouter":     &VPCRouterGenerator{},
	}
}
