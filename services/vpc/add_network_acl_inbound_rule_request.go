/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-09-09T02:06:23Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type AddNetworkAclInboundRuleRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크ACL번호
NetworkAclNo *string `json:"networkAclNo"`

	// 네트워크ACLRule리스트
NetworkAclRuleList []*AddNetworkAclRuleParameter `json:"networkAclRuleList"`
}
