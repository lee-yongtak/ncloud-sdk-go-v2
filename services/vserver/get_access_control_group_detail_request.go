/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetAccessControlGroupDetailRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ACG번호
AccessControlGroupNo *string `json:"accessControlGroupNo"`
}
