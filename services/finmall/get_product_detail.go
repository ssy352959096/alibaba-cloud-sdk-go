package finmall

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

// GetProductDetail invokes the finmall.GetProductDetail API synchronously
// api document: https://help.aliyun.com/api/finmall/getproductdetail.html
func (client *Client) GetProductDetail(request *GetProductDetailRequest) (response *GetProductDetailResponse, err error) {
	response = CreateGetProductDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetProductDetailWithChan invokes the finmall.GetProductDetail API asynchronously
// api document: https://help.aliyun.com/api/finmall/getproductdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductDetailWithChan(request *GetProductDetailRequest) (<-chan *GetProductDetailResponse, <-chan error) {
	responseChan := make(chan *GetProductDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProductDetail(request)
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

// GetProductDetailWithCallback invokes the finmall.GetProductDetail API asynchronously
// api document: https://help.aliyun.com/api/finmall/getproductdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProductDetailWithCallback(request *GetProductDetailRequest, callback func(response *GetProductDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProductDetailResponse
		var err error
		defer close(result)
		response, err = client.GetProductDetail(request)
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

// GetProductDetailRequest is the request struct for api GetProductDetail
type GetProductDetailRequest struct {
	*requests.RpcRequest
	ProductId   string `position:"Query" name:"ProductId"`
	FundPartyId string `position:"Query" name:"FundPartyId"`
	UserId      string `position:"Query" name:"UserId"`
}

// GetProductDetailResponse is the response struct for api GetProductDetail
type GetProductDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetProductDetailRequest creates a request to invoke GetProductDetail API
func CreateGetProductDetailRequest() (request *GetProductDetailRequest) {
	request = &GetProductDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetProductDetail", "finmall", "openAPI")
	return
}

// CreateGetProductDetailResponse creates a response to parse from GetProductDetail response
func CreateGetProductDetailResponse() (response *GetProductDetailResponse) {
	response = &GetProductDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
