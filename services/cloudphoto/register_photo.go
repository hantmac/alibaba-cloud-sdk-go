package cloudphoto

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

func (client *Client) RegisterPhoto(request *RegisterPhotoRequest) (response *RegisterPhotoResponse, err error) {
	response = CreateRegisterPhotoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RegisterPhotoWithChan(request *RegisterPhotoRequest) (<-chan *RegisterPhotoResponse, <-chan error) {
	responseChan := make(chan *RegisterPhotoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterPhoto(request)
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

func (client *Client) RegisterPhotoWithCallback(request *RegisterPhotoRequest, callback func(response *RegisterPhotoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterPhotoResponse
		var err error
		defer close(result)
		response, err = client.RegisterPhoto(request)
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

type RegisterPhotoRequest struct {
	*requests.RpcRequest
	TakenAt    requests.Integer `position:"Query" name:"TakenAt"`
	Location   string           `position:"Query" name:"Location"`
	StoreName  string           `position:"Query" name:"StoreName"`
	LibraryId  string           `position:"Query" name:"LibraryId"`
	Latitude   requests.Float   `position:"Query" name:"Latitude"`
	Longitude  requests.Float   `position:"Query" name:"Longitude"`
	Width      requests.Integer `position:"Query" name:"Width"`
	Height     requests.Integer `position:"Query" name:"Height"`
	IsVideo    string           `position:"Query" name:"IsVideo"`
	Md5        string           `position:"Query" name:"Md5"`
	Size       requests.Integer `position:"Query" name:"Size"`
	PhotoTitle string           `position:"Query" name:"PhotoTitle"`
	Remark     string           `position:"Query" name:"Remark"`
}

type RegisterPhotoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Photo     Photo  `json:"Photo" xml:"Photo"`
}

func CreateRegisterPhotoRequest() (request *RegisterPhotoRequest) {
	request = &RegisterPhotoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "RegisterPhoto", "cloudphoto", "openAPI")
	return
}

func CreateRegisterPhotoResponse() (response *RegisterPhotoResponse) {
	response = &RegisterPhotoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
