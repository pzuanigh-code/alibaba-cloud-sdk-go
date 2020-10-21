package mts

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

// UpdateTemplate invokes the mts.UpdateTemplate API synchronously
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (response *UpdateTemplateResponse, err error) {
	response = CreateUpdateTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTemplateWithChan invokes the mts.UpdateTemplate API asynchronously
func (client *Client) UpdateTemplateWithChan(request *UpdateTemplateRequest) (<-chan *UpdateTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTemplate(request)
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

// UpdateTemplateWithCallback invokes the mts.UpdateTemplate API asynchronously
func (client *Client) UpdateTemplateWithCallback(request *UpdateTemplateRequest, callback func(response *UpdateTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateTemplate(request)
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

// UpdateTemplateRequest is the request struct for api UpdateTemplate
type UpdateTemplateRequest struct {
	*requests.RpcRequest
	Container            string           `position:"Query" name:"Container"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Video                string           `position:"Query" name:"Video"`
	TransConfig          string           `position:"Query" name:"TransConfig"`
	Audio                string           `position:"Query" name:"Audio"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MuxConfig            string           `position:"Query" name:"MuxConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
	Name                 string           `position:"Query" name:"Name"`
}

// UpdateTemplateResponse is the response struct for api UpdateTemplate
type UpdateTemplateResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Template  Template `json:"Template" xml:"Template"`
}

// CreateUpdateTemplateRequest creates a request to invoke UpdateTemplate API
func CreateUpdateTemplateRequest() (request *UpdateTemplateRequest) {
	request = &UpdateTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateTemplate", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTemplateResponse creates a response to parse from UpdateTemplate response
func CreateUpdateTemplateResponse() (response *UpdateTemplateResponse) {
	response = &UpdateTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
