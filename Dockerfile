# Copyright 2019 Kazumichi Yamamoto.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.13 as builder
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>

RUN  apt-get update && apt-get -y install \
        bash \
        git  \
        make \
        zip  \
        bzr  \
      && apt-get clean \
      && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

ADD . /go/src/github.com/sacloud/terraformer
WORKDIR /go/src/github.com/sacloud/terraformer
RUN make build

# ---

FROM hashicorp/terraform:0.12.9
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>

RUN mkdir -p ~/.terraform.d/plugins/linux_amd64
ADD https://github.com/sacloud/terraform-provider-sakuracloud/releases/download/v1.16.4/terraform-provider-sakuracloud_1.16.4_linux-amd64.zip ./
RUN unzip terraform-provider-sakuracloud_1.16.4_linux-amd64.zip -d ~/.terraform.d/plugins/linux_amd64
RUN rm -f terraform-provider-sakuracloud_1.16.4_linux-amd64.zip

COPY --from=builder /go/src/github.com/sacloud/terraformer/bin/terraformer /bin/

ENTRYPOINT ["/bin/terraformer"]
WORKDIR /work
CMD ["--help"]
