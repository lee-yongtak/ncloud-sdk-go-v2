/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetPlacementGroupListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 물리배치그룹이름
PlacementGroupName *string `json:"placementGroupName,omitempty"`

	// 물리배치그룹번호리스트
PlacementGroupNoList []*string `json:"placementGroupNoList,omitempty"`
}
