/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-07-06T00:53:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type GetNetworkAclRuleListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크ACL번호
NetworkAclNo *string `json:"networkAclNo"`

	// 네트워크ACLRule유형코드
NetworkAclRuleTypeCode *string `json:"networkAclRuleTypeCode,omitempty"`
}
