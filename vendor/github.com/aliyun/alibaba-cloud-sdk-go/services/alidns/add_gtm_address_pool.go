package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddGtmAddressPool invokes the alidns.AddGtmAddressPool API synchronously
// api document: https://help.aliyun.com/api/alidns/addgtmaddresspool.html
func (client *Client) AddGtmAddressPool(request *AddGtmAddressPoolRequest) (response *AddGtmAddressPoolResponse, err error) {
	response = CreateAddGtmAddressPoolResponse()
	err = client.DoAction(request, response)
	return
}

// AddGtmAddressPoolWithChan invokes the alidns.AddGtmAddressPool API asynchronously
// api document: https://help.aliyun.com/api/alidns/addgtmaddresspool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGtmAddressPoolWithChan(request *AddGtmAddressPoolRequest) (<-chan *AddGtmAddressPoolResponse, <-chan error) {
	responseChan := make(chan *AddGtmAddressPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddGtmAddressPool(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddGtmAddressPoolWithCallback invokes the alidns.AddGtmAddressPool API asynchronously
// api document: https://help.aliyun.com/api/alidns/addgtmaddresspool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGtmAddressPoolWithCallback(request *AddGtmAddressPoolRequest, callback func(response *AddGtmAddressPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddGtmAddressPoolResponse
		var err error
		defer close(result)
		response, err = client.AddGtmAddressPool(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddGtmAddressPoolRequest is the request struct for api AddGtmAddressPool
type AddGtmAddressPoolRequest struct {
	*requests.RpcRequest
	InstanceId          string                   `position:"Query" name:"InstanceId"`
	UserClientIp        string                   `position:"Query" name:"UserClientIp"`
	Name                string                   `position:"Query" name:"Name"`
	Lang                string                   `position:"Query" name:"Lang"`
	Type                string                   `position:"Query" name:"Type"`
	Addr                *[]AddGtmAddressPoolAddr `position:"Query" name:"Addr"  type:"Repeated"`
	MinAvailableAddrNum requests.Integer         `position:"Query" name:"MinAvailableAddrNum"`
}

// AddGtmAddressPoolAddr is a repeated param struct in AddGtmAddressPoolRequest
type AddGtmAddressPoolAddr struct {
	Mode      string `name:"Mode"`
	LbaWeight string `name:"LbaWeight"`
	Value     string `name:"Value"`
}

// AddGtmAddressPoolResponse is the response struct for api AddGtmAddressPool
type AddGtmAddressPoolResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	AddrPoolId string `json:"AddrPoolId" xml:"AddrPoolId"`
}

// CreateAddGtmAddressPoolRequest creates a request to invoke AddGtmAddressPool API
func CreateAddGtmAddressPoolRequest() (request *AddGtmAddressPoolRequest) {
	request = &AddGtmAddressPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddGtmAddressPool", "Alidns", "openAPI")
	return
}

// CreateAddGtmAddressPoolResponse creates a response to parse from AddGtmAddressPool response
func CreateAddGtmAddressPoolResponse() (response *AddGtmAddressPoolResponse) {
	response = &AddGtmAddressPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
