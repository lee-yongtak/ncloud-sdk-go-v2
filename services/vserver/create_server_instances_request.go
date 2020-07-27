/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateServerInstancesRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 회원서버이미지인스턴스번호
MemberServerImageInstanceNo *string `json:"memberServerImageInstanceNo,omitempty"`

	// 서버이름
ServerName *string `json:"serverName,omitempty"`

	// 서버설명
ServerDescription *string `json:"serverDescription,omitempty"`

	// 로그인키이름
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 반납보호여부
IsProtectServerTermination *bool `json:"isProtectServerTermination,omitempty"`

	// 서버생성개수
ServerCreateCount *int32 `json:"serverCreateCount,omitempty"`

	// 서버생성시작번호
ServerCreateStartNo *int32 `json:"serverCreateStartNo,omitempty"`

	// 요금제유형코드
FeeSystemTypeCode *string `json:"feeSystemTypeCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 초기화스크립트번호
InitScriptNo *string `json:"initScriptNo,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`

	// 서브넷번호
SubnetNo *string `json:"subnetNo"`

	// 네트워크인터페이스리스트
NetworkInterfaceList []*NetworkInterfaceParameter `json:"networkInterfaceList"`

	// 물리배치그룹번호
PlacementGroupNo *string `json:"placementGroupNo,omitempty"`

	// 기본블록스토리지볼륨암호화여부
IsEncryptedBaseBlockStorageVolume *bool `json:"isEncryptedBaseBlockStorageVolume,omitempty"`
}