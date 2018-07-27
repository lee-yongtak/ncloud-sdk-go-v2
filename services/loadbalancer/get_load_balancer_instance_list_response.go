/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * API version: 2018-06-21T02:19:18Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type GetLoadBalancerInstanceListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

LoadBalancerInstanceList []*LoadBalancerInstance `json:"loadBalancerInstanceList,omitempty"`
}
