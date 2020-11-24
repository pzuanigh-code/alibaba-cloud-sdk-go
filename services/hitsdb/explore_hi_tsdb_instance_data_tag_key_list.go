package hitsdb

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

// ExploreHiTSDBInstanceDataTagKeyList invokes the hitsdb.ExploreHiTSDBInstanceDataTagKeyList API synchronously
func (client *Client) ExploreHiTSDBInstanceDataTagKeyList(request *ExploreHiTSDBInstanceDataTagKeyListRequest) (response *ExploreHiTSDBInstanceDataTagKeyListResponse, err error) {
	response = CreateExploreHiTSDBInstanceDataTagKeyListResponse()
	err = client.DoAction(request, response)
	return
}

// ExploreHiTSDBInstanceDataTagKeyListWithChan invokes the hitsdb.ExploreHiTSDBInstanceDataTagKeyList API asynchronously
func (client *Client) ExploreHiTSDBInstanceDataTagKeyListWithChan(request *ExploreHiTSDBInstanceDataTagKeyListRequest) (<-chan *ExploreHiTSDBInstanceDataTagKeyListResponse, <-chan error) {
	responseChan := make(chan *ExploreHiTSDBInstanceDataTagKeyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExploreHiTSDBInstanceDataTagKeyList(request)
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

// ExploreHiTSDBInstanceDataTagKeyListWithCallback invokes the hitsdb.ExploreHiTSDBInstanceDataTagKeyList API asynchronously
func (client *Client) ExploreHiTSDBInstanceDataTagKeyListWithCallback(request *ExploreHiTSDBInstanceDataTagKeyListRequest, callback func(response *ExploreHiTSDBInstanceDataTagKeyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExploreHiTSDBInstanceDataTagKeyListResponse
		var err error
		defer close(result)
		response, err = client.ExploreHiTSDBInstanceDataTagKeyList(request)
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

// ExploreHiTSDBInstanceDataTagKeyListRequest is the request struct for api ExploreHiTSDBInstanceDataTagKeyList
type ExploreHiTSDBInstanceDataTagKeyListRequest struct {
	*requests.RpcRequest
	ReverseVpcIp         string           `position:"Query" name:"ReverseVpcIp"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Prefix               string           `position:"Query" name:"Prefix"`
	ReverseVpcPort       requests.Integer `position:"Query" name:"ReverseVpcPort"`
	PassWord             string           `position:"Query" name:"PassWord"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Max                  requests.Integer `position:"Query" name:"Max"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Metric               string           `position:"Query" name:"Metric"`
	UserName             string           `position:"Query" name:"UserName"`
}

// ExploreHiTSDBInstanceDataTagKeyListResponse is the response struct for api ExploreHiTSDBInstanceDataTagKeyList
type ExploreHiTSDBInstanceDataTagKeyListResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	TagKeyList []Data `json:"TagKeyList" xml:"TagKeyList"`
}

// CreateExploreHiTSDBInstanceDataTagKeyListRequest creates a request to invoke ExploreHiTSDBInstanceDataTagKeyList API
func CreateExploreHiTSDBInstanceDataTagKeyListRequest() (request *ExploreHiTSDBInstanceDataTagKeyListRequest) {
	request = &ExploreHiTSDBInstanceDataTagKeyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2017-06-01", "ExploreHiTSDBInstanceDataTagKeyList", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExploreHiTSDBInstanceDataTagKeyListResponse creates a response to parse from ExploreHiTSDBInstanceDataTagKeyList response
func CreateExploreHiTSDBInstanceDataTagKeyListResponse() (response *ExploreHiTSDBInstanceDataTagKeyListResponse) {
	response = &ExploreHiTSDBInstanceDataTagKeyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}