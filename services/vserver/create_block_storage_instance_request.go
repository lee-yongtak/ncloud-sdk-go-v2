/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateBlockStorageInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 블록스토리지이름
BlockStorageName *string `json:"blockStorageName,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *int32 `json:"blockStorageSize"`

	// 블록스토리지디스크상세유형코드
BlockStorageDiskDetailTypeCode *string `json:"blockStorageDiskDetailTypeCode,omitempty"`

	// 블록스토리지설명
BlockStorageDescription *string `json:"blockStorageDescription,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 블록스토리지스냅샷인스턴스번호
BlockStorageSnapshotInstanceNo *string `json:"blockStorageSnapshotInstanceNo,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`
}
