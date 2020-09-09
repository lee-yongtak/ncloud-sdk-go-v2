/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RemoveAccessControlGroupOutboundRuleRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// ACGRule리스트
AccessControlGroupRuleList []*RemoveAccessControlGroupRuleParameter `json:"accessControlGroupRuleList"`
}
