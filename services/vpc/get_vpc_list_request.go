/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-07-06T00:53:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetVpcListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`

	// VPC상태코드
VpcStatusCode *string `json:"vpcStatusCode,omitempty"`

	// VPC번호리스트
VpcNoList []*string `json:"vpcNoList,omitempty"`
}