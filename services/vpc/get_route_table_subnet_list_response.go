/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-09-09T02:06:23Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetRouteTableSubnetListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

SubnetList []*Subnet `json:"subnetList,omitempty"`
}
