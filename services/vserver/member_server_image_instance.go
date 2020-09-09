/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type MemberServerImageInstance struct {

	// 회원서버이미지인스턴스번호
MemberServerImageInstanceNo *string `json:"memberServerImageInstanceNo,omitempty"`

	// 회원서버이미지이름
MemberServerImageName *string `json:"memberServerImageName,omitempty"`

	// 회원서버이미지설명
MemberServerImageDescription *string `json:"memberServerImageDescription,omitempty"`

	// 원본서버인스턴스번호
OriginalServerInstanceNo *string `json:"originalServerInstanceNo,omitempty"`

	// 원본서버이미지상품코드
OriginalServerImageProductCode *string `json:"originalServerImageProductCode,omitempty"`

	// 회원서버이미지인스턴스상태
MemberServerImageInstanceStatus *CommonCode `json:"memberServerImageInstanceStatus,omitempty"`

	// 회원서버이미지인스턴스OP
MemberServerImageInstanceOperation *CommonCode `json:"memberServerImageInstanceOperation,omitempty"`

	// 회원서버이미지인스턴스상태이름
MemberServerImageInstanceStatusName *string `json:"memberServerImageInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 회원서버이미지블록스토리지총개수
MemberServerImageBlockStorageTotalRows *int32 `json:"memberServerImageBlockStorageTotalRows,omitempty"`

	// 회원서버이미지블록스토리지총사이즈
MemberServerImageBlockStorageTotalSize *int64 `json:"memberServerImageBlockStorageTotalSize,omitempty"`
}
