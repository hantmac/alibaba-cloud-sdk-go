package rds

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

func (client *Client) ImportDatabaseBetweenInstances(request *ImportDatabaseBetweenInstancesRequest) (response *ImportDatabaseBetweenInstancesResponse, err error) {
	response = CreateImportDatabaseBetweenInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ImportDatabaseBetweenInstancesWithChan(request *ImportDatabaseBetweenInstancesRequest) (<-chan *ImportDatabaseBetweenInstancesResponse, <-chan error) {
	responseChan := make(chan *ImportDatabaseBetweenInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportDatabaseBetweenInstances(request)
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

func (client *Client) ImportDatabaseBetweenInstancesWithCallback(request *ImportDatabaseBetweenInstancesRequest, callback func(response *ImportDatabaseBetweenInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportDatabaseBetweenInstancesResponse
		var err error
		defer close(result)
		response, err = client.ImportDatabaseBetweenInstances(request)
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

type ImportDatabaseBetweenInstancesRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SourceDBInstanceId   string           `position:"Query" name:"SourceDBInstanceId"`
	DBInfo               string           `position:"Query" name:"DBInfo"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ImportDatabaseBetweenInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImportId  string `json:"ImportId" xml:"ImportId"`
}

func CreateImportDatabaseBetweenInstancesRequest() (request *ImportDatabaseBetweenInstancesRequest) {
	request = &ImportDatabaseBetweenInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ImportDatabaseBetweenInstances", "rds", "openAPI")
	return
}

func CreateImportDatabaseBetweenInstancesResponse() (response *ImportDatabaseBetweenInstancesResponse) {
	response = &ImportDatabaseBetweenInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
