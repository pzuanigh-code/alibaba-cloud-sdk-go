package cloudcallcenter

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

// DebugIntent invokes the cloudcallcenter.DebugIntent API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugintent.html
func (client *Client) DebugIntent(request *DebugIntentRequest) (response *DebugIntentResponse, err error) {
	response = CreateDebugIntentResponse()
	err = client.DoAction(request, response)
	return
}

// DebugIntentWithChan invokes the cloudcallcenter.DebugIntent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugIntentWithChan(request *DebugIntentRequest) (<-chan *DebugIntentResponse, <-chan error) {
	responseChan := make(chan *DebugIntentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugIntent(request)
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

// DebugIntentWithCallback invokes the cloudcallcenter.DebugIntent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugIntentWithCallback(request *DebugIntentRequest, callback func(response *DebugIntentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugIntentResponse
		var err error
		defer close(result)
		response, err = client.DebugIntent(request)
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

// DebugIntentRequest is the request struct for api DebugIntent
type DebugIntentRequest struct {
	*requests.RpcRequest
	CallId     string           `position:"Query" name:"CallId"`
	SurveyId   string           `position:"Query" name:"SurveyId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	NewSession requests.Boolean `position:"Query" name:"NewSession"`
	ScenarioId string           `position:"Query" name:"ScenarioId"`
	Utterance  string           `position:"Query" name:"Utterance"`
}

// DebugIntentResponse is the response struct for api DebugIntent
type DebugIntentResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Feedback       Feedback `json:"Feedback" xml:"Feedback"`
}

// CreateDebugIntentRequest creates a request to invoke DebugIntent API
func CreateDebugIntentRequest() (request *DebugIntentRequest) {
	request = &DebugIntentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DebugIntent", "", "")
	request.Method = requests.POST
	return
}

// CreateDebugIntentResponse creates a response to parse from DebugIntent response
func CreateDebugIntentResponse() (response *DebugIntentResponse) {
	response = &DebugIntentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}