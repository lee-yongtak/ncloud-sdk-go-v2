/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-08-27T03:07:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetPlacementGroupDetailRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 물리배치그룹번호리스트
PlacementGroupNo *string `json:"placementGroupNo"`
}
