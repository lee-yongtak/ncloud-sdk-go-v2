/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreatePlacementGroupRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 물리배치그룹이름
PlacementGroupName *string `json:"placementGroupName,omitempty"`

	// 물리배치그룹유형코드
PlacementGroupTypeCode *string `json:"placementGroupTypeCode,omitempty"`
}
