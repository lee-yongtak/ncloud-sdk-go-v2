/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetMemberServerImageInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 회원서버이미지이름
MemberServerImageName *string `json:"memberServerImageName,omitempty"`

	// 회원서버이미지인스턴스상태코드
MemberServerImageInstanceStatusCode *string `json:"memberServerImageInstanceStatusCode,omitempty"`

	// 회원서버이미지인스턴스번호리스트
MemberServerImageInstanceNoList []*string `json:"memberServerImageInstanceNoList,omitempty"`

	// 플랫폼유형코드리스트
PlatformTypeCodeList []*string `json:"platformTypeCodeList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 정렬순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}
