/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type AddAccessControlGroupInboundRuleRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// ACGRule리스트
AccessControlGroupRuleList []*AddAccessControlGroupRuleParameter `json:"accessControlGroupRuleList"`
}
