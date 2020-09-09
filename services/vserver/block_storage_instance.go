/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type BlockStorageInstance struct {

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 블록스토리지이름
BlockStorageName *string `json:"blockStorageName,omitempty"`

	// 블록스토리지유형
BlockStorageType *CommonCode `json:"blockStorageType,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *int64 `json:"blockStorageSize,omitempty"`

	// 디바이스이름
DeviceName *string `json:"deviceName,omitempty"`

	// 블록스토리지상품코드
BlockStorageProductCode *string `json:"blockStorageProductCode,omitempty"`

	// 블록스토리지인스턴스상태
BlockStorageInstanceStatus *CommonCode `json:"blockStorageInstanceStatus,omitempty"`

	// 블록스토리지인스턴스OP
BlockStorageInstanceOperation *CommonCode `json:"blockStorageInstanceOperation,omitempty"`

	// 블록스토리지인스턴스상태이름
BlockStorageInstanceStatusName *string `json:"blockStorageInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 블록스토리지설명
BlockStorageDescription *string `json:"blockStorageDescription,omitempty"`

	// 블록스토리지디스크유형
BlockStorageDiskType *CommonCode `json:"blockStorageDiskType,omitempty"`

	// 블록스토리지디스크상세유형
BlockStorageDiskDetailType *CommonCode `json:"blockStorageDiskDetailType,omitempty"`

	// 최대IOPS
MaxIopsThroughput *int32 `json:"maxIopsThroughput,omitempty"`

	// 볼륨암호화여부
IsEncryptedVolume *string `json:"isEncryptedVolume,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`
}
