/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-07-06T00:53:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type CreateNetworkAclRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크ACL설명
NetworkAclDescription *string `json:"networkAclDescription,omitempty"`

	// 네트워크ACL이름
NetworkAclName *string `json:"networkAclName,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`
}
