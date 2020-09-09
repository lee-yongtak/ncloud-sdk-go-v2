/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vpc/v2
 *
 * API version: 2020-09-09T02:06:23Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type RemoveNetworkAclRuleParameter struct {

	// IP블록
IpBlock *string `json:"ipBlock"`

	// Rule액션코드
RuleActionCode *string `json:"ruleActionCode"`

	// 포트범위
PortRange *string `json:"portRange,omitempty"`

	// 우선순위
Priority *int32 `json:"priority"`

	// 프로토콜유형코드
ProtocolTypeCode *string `json:"protocolTypeCode"`
}
