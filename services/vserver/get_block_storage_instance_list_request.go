/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetBlockStorageInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 블록스토리지유형코드리스트
BlockStorageTypeCodeList []*string `json:"blockStorageTypeCodeList,omitempty"`

	// 블록스토리지인스턴스상태코드
BlockStorageInstanceStatusCode *string `json:"blockStorageInstanceStatusCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *int32 `json:"blockStorageSize,omitempty"`

	// 블록스토리지인스턴스번호리스트
BlockStorageInstanceNoList []*string `json:"blockStorageInstanceNoList,omitempty"`

	// 블록스토리지이름
BlockStorageName *string `json:"blockStorageName,omitempty"`

	// 서버이름
ServerName *string `json:"serverName,omitempty"`

	// 연결정보
ConnectionInfo *string `json:"connectionInfo,omitempty"`

	// 블록스토리지디스크유형코드
BlockStorageDiskTypeCode *string `json:"blockStorageDiskTypeCode,omitempty"`

	// 블록스토리지디스크상세유형코드
BlockStorageDiskDetailTypeCode *string `json:"blockStorageDiskDetailTypeCode,omitempty"`
}
