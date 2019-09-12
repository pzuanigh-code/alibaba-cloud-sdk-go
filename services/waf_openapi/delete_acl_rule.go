package waf_openapi

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

// DeleteAclRule invokes the waf_openapi.DeleteAclRule API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/deleteaclrule.html
func (client *Client) DeleteAclRule(request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
	response = CreateDeleteAclRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAclRuleWithChan invokes the waf_openapi.DeleteAclRule API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/deleteaclrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAclRuleWithChan(request *DeleteAclRuleRequest) (<-chan *DeleteAclRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteAclRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAclRule(request)
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

// DeleteAclRuleWithCallback invokes the waf_openapi.DeleteAclRule API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/deleteaclrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAclRuleWithCallback(request *DeleteAclRuleRequest, callback func(response *DeleteAclRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAclRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteAclRule(request)
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

// DeleteAclRuleRequest is the request struct for api DeleteAclRule
type DeleteAclRuleRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	Domain     string           `position:"Query" name:"Domain"`
	Lang       string           `position:"Query" name:"Lang"`
	RuleId     requests.Integer `position:"Query" name:"RuleId"`
	Region     string           `position:"Query" name:"Region"`
}

// DeleteAclRuleResponse is the response struct for api DeleteAclRule
type DeleteAclRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeleteAclRuleRequest creates a request to invoke DeleteAclRule API
func CreateDeleteAclRuleRequest() (request *DeleteAclRuleRequest) {
	request = &DeleteAclRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "DeleteAclRule", "waf", "openAPI")
	return
}

// CreateDeleteAclRuleResponse creates a response to parse from DeleteAclRule response
func CreateDeleteAclRuleResponse() (response *DeleteAclRuleResponse) {
	response = &DeleteAclRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}