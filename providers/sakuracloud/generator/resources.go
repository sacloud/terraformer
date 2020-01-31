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

var terraformResources = map[string]sakuraCloudResourceRenderable{
	"archive": sakuraCloudResource{
		terraformName:      "sakuracloud_archive",
		pluralResourceName: "Archives",
		useScopeFilter:     true,
	},
	//
	// NOTE: autoBackupとbridgeは複数ゾーンでのリソース検索に癖があるためマニュアル実装とする
	//
	//"autoBackup": sakuraCloudResource{
	//	terraformName:      "sakuracloud_auto_backup",
	//	pluralResourceName: "AutoBackups",
	//},
	//"bridge": sakuraCloudResource{
	//	terraformName:      "sakuracloud_bridge",
	//	pluralResourceName: "Bridges",
	//},

	"cdrom": sakuraCloudResource{
		terraformName:      "sakuracloud_cdrom",
		pluralResourceName: "CDROMs",
		useScopeFilter:     true,
	},
	"containerRegistry": sakuraCloudResource{
		terraformName:      "sakuracloud_container_registry",
		pluralResourceName: "ContainerRegistries",
		globalZone:         true,
	},
	"database": sakuraCloudResource{
		terraformName:      "sakuracloud_database",
		pluralResourceName: "Databases",
	},
	"disk": sakuraCloudResource{
		terraformName:      "sakuracloud_disk",
		pluralResourceName: "Disks",
	},
	"dns": sakuraCloudResource{
		terraformName:      "sakuracloud_dns",
		pluralResourceName: "DNS",
		globalZone:         true,
	},
	"gslb": sakuraCloudResource{
		terraformName:      "sakuracloud_gslb",
		pluralResourceName: "GSLBs",
		globalZone:         true,
	},
	"icon": sakuraCloudResource{
		terraformName:      "sakuracloud_icon",
		pluralResourceName: "Icons",
		globalZone:         true,
		useScopeFilter:     true,
	},
	"internet": sakuraCloudResource{
		terraformName:      "sakuracloud_internet",
		pluralResourceName: "Internet",
	},
	"loadBalancer": sakuraCloudResource{
		terraformName:      "sakuracloud_load_balancer",
		pluralResourceName: "LoadBalancers",
	},
	"localRouter": sakuraCloudResource{
		terraformName:      "sakuracloud_local_router",
		pluralResourceName: "LocalRouters",
		globalZone:         true,
	},
	"mobileGateway": sakuraCloudResource{
		terraformName:      "sakuracloud_mobile_gateway",
		pluralResourceName: "MobileGateways",
	},
	"nfs": sakuraCloudResource{
		terraformName:      "sakuracloud_nfs",
		pluralResourceName: "NFS",
	},
	"note": sakuraCloudResource{
		terraformName:      "sakuracloud_note",
		pluralResourceName: "Notes",
		globalZone:         true,
		useScopeFilter:     true,
	},
	"packetFilter": sakuraCloudResource{
		terraformName:      "sakuracloud_packet_filter",
		pluralResourceName: "PacketFilters",
	},
	"privateHost": sakuraCloudResource{
		terraformName:      "sakuracloud_private_host",
		pluralResourceName: "PrivateHosts",
	},
	"proxyLB": sakuraCloudResource{
		terraformName:      "sakuracloud_proxylb",
		pluralResourceName: "ProxyLBs",
		globalZone:         true,
	},
	"server": sakuraCloudResource{
		terraformName:      "sakuracloud_server",
		pluralResourceName: "Servers",
	},
	"sim": sakuraCloudResource{
		terraformName:      "sakuracloud_sim",
		pluralResourceName: "SIMs",
		globalZone:         true,
	},
	"simpleMonitor": sakuraCloudResource{
		terraformName:      "sakuracloud_simple_monitor",
		pluralResourceName: "SimpleMonitors",
		globalZone:         true,
	},
	"sshKey": sakuraCloudResource{
		terraformName:      "sakuracloud_ssh_key",
		pluralResourceName: "SSHKeys",
		globalZone:         true,
	},
	"switch": sakuraCloudResource{
		terraformName:      "sakuracloud_switch",
		pluralResourceName: "Switches",
	},
	"vpcRouter": sakuraCloudResource{
		terraformName:      "sakuracloud_vpc_router",
		pluralResourceName: "VPCRouters",
	},
}
