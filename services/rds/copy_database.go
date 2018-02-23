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

func (client *Client) CopyDatabase(request *CopyDatabaseRequest) (response *CopyDatabaseResponse, err error) {
	response = CreateCopyDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CopyDatabaseWithChan(request *CopyDatabaseRequest) (<-chan *CopyDatabaseResponse, <-chan error) {
	responseChan := make(chan *CopyDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyDatabase(request)
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

func (client *Client) CopyDatabaseWithCallback(request *CopyDatabaseRequest, callback func(response *CopyDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyDatabaseResponse
		var err error
		defer close(result)
		response, err = client.CopyDatabase(request)
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

type CopyDatabaseRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
}

type CopyDatabaseResponse struct {
	*responses.BaseResponse
	DBName   string `json:"DBName" xml:"DBName"`
	DBStatus string `json:"DBStatus" xml:"DBStatus"`
	TaskId   string `json:"TaskId" xml:"TaskId"`
}

func CreateCopyDatabaseRequest() (request *CopyDatabaseRequest) {
	request = &CopyDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CopyDatabase", "rds", "openAPI")
	return
}

func CreateCopyDatabaseResponse() (response *CopyDatabaseResponse) {
	response = &CopyDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
