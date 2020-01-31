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
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
)

type SakuraCloudService struct {
	terraform_utils.Service
}

const (
	DefaultAPIRequestTimeout   = 60
	DefaultAPIRequestRateLimit = 10
)

var sacloudCaller sacloud.APICaller
var initOnce sync.Once

func (s SakuraCloudService) NewClient() sacloud.APICaller {
	initOnce.Do(func() {
		httpClient := &http.Client{
			Timeout:   time.Duration(DefaultAPIRequestTimeout) * time.Second,
			Transport: &sacloud.RateLimitRoundTripper{RateLimitPerSec: DefaultAPIRequestRateLimit},
		}
		caller := &sacloud.Client{
			AccessToken:       s.Args["token"].(string),
			AccessTokenSecret: s.Args["secret"].(string),
			UserAgent:         "terraformer/v0.7.8 (sacloud version)",
			AcceptLanguage:    sacloud.APIDefaultAcceptLanguage,
			RetryMax:          sacloud.APIDefaultRetryMax,
			RetryWaitMax:      sacloud.APIDefaultRetryWaitMax,
			RetryWaitMin:      sacloud.APIDefaultRetryWaitMin,
			HTTPClient:        httpClient,
		}

		if traceMode := os.Getenv("SAKURACLOUD_TRACE"); traceMode != "" {
			enableAPITrace := true
			enableHTTPTrace := true

			mode := strings.ToLower(traceMode)
			switch mode {
			case "api":
				enableHTTPTrace = false
			case "http":
				enableAPITrace = false
			}

			if enableAPITrace {
				trace.AddClientFactoryHooks()
			}
			if enableHTTPTrace {
				caller.HTTPClient.Transport = &sacloud.TracingRoundTripper{
					Transport: caller.HTTPClient.Transport,
				}
			}
		}
		sacloudCaller = caller
	})
	return sacloudCaller
}
