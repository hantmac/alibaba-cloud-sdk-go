package slb

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

func (client *Client) CreateMasterSlaveVServerGroup(request *CreateMasterSlaveVServerGroupRequest) (response *CreateMasterSlaveVServerGroupResponse, err error) {
	response = CreateCreateMasterSlaveVServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateMasterSlaveVServerGroupWithChan(request *CreateMasterSlaveVServerGroupRequest) (<-chan *CreateMasterSlaveVServerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateMasterSlaveVServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMasterSlaveVServerGroup(request)
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

func (client *Client) CreateMasterSlaveVServerGroupWithCallback(request *CreateMasterSlaveVServerGroupRequest, callback func(response *CreateMasterSlaveVServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMasterSlaveVServerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateMasterSlaveVServerGroup(request)
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

type CreateMasterSlaveVServerGroupRequest struct {
	*requests.RpcRequest
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId                 string           `position:"Query" name:"access_key_id"`
	Tags                        string           `position:"Query" name:"Tags"`
	LoadBalancerId              string           `position:"Query" name:"LoadBalancerId"`
	MasterSlaveVServerGroupName string           `position:"Query" name:"MasterSlaveVServerGroupName"`
	MasterSlaveBackendServers   string           `position:"Query" name:"MasterSlaveBackendServers"`
}

type CreateMasterSlaveVServerGroupResponse struct {
	*responses.BaseResponse
	RequestId                 string                                                   `json:"RequestId" xml:"RequestId"`
	MasterSlaveVServerGroupId string                                                   `json:"MasterSlaveVServerGroupId" xml:"MasterSlaveVServerGroupId"`
	MasterSlaveBackendServers MasterSlaveBackendServersInCreateMasterSlaveVServerGroup `json:"MasterSlaveBackendServers" xml:"MasterSlaveBackendServers"`
}

func CreateCreateMasterSlaveVServerGroupRequest() (request *CreateMasterSlaveVServerGroupRequest) {
	request = &CreateMasterSlaveVServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveVServerGroup", "slb", "openAPI")
	return
}

func CreateCreateMasterSlaveVServerGroupResponse() (response *CreateMasterSlaveVServerGroupResponse) {
	response = &CreateMasterSlaveVServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
