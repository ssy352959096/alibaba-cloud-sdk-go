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

// SubmitMediaCensorJob invokes the mts.SubmitMediaCensorJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitmediacensorjob.html
func (client *Client) SubmitMediaCensorJob(request *SubmitMediaCensorJobRequest) (response *SubmitMediaCensorJobResponse, err error) {
	response = CreateSubmitMediaCensorJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitMediaCensorJobWithChan invokes the mts.SubmitMediaCensorJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmediacensorjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMediaCensorJobWithChan(request *SubmitMediaCensorJobRequest) (<-chan *SubmitMediaCensorJobResponse, <-chan error) {
	responseChan := make(chan *SubmitMediaCensorJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitMediaCensorJob(request)
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

// SubmitMediaCensorJobWithCallback invokes the mts.SubmitMediaCensorJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmediacensorjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMediaCensorJobWithCallback(request *SubmitMediaCensorJobRequest, callback func(response *SubmitMediaCensorJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitMediaCensorJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitMediaCensorJob(request)
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

// SubmitMediaCensorJobRequest is the request struct for api SubmitMediaCensorJob
type SubmitMediaCensorJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CoverImages          string           `position:"Query" name:"CoverImages"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Title                string           `position:"Query" name:"Title"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	VideoCensorConfig    string           `position:"Query" name:"VideoCensorConfig"`
	Input                string           `position:"Query" name:"Input"`
	UserData             string           `position:"Query" name:"UserData"`
	Barrages             string           `position:"Query" name:"Barrages"`
}

// SubmitMediaCensorJobResponse is the response struct for api SubmitMediaCensorJob
type SubmitMediaCensorJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitMediaCensorJobRequest creates a request to invoke SubmitMediaCensorJob API
func CreateSubmitMediaCensorJobRequest() (request *SubmitMediaCensorJobRequest) {
	request = &SubmitMediaCensorJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitMediaCensorJob", "mts", "openAPI")
	return
}

// CreateSubmitMediaCensorJobResponse creates a response to parse from SubmitMediaCensorJob response
func CreateSubmitMediaCensorJobResponse() (response *SubmitMediaCensorJobResponse) {
	response = &SubmitMediaCensorJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
