/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-08-27T03:07:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DeleteMemberServerImageInstancesRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 회원서버이미지인스턴스번호리스트
MemberServerImageInstanceNoList []*string `json:"memberServerImageInstanceNoList"`
}
