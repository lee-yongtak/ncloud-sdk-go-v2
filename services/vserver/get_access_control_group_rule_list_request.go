/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetAccessControlGroupRuleListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo"`

	// ACGRule유형코드
AccessControlGroupRuleTypeCode *string `json:"accessControlGroupRuleTypeCode,omitempty"`
}
