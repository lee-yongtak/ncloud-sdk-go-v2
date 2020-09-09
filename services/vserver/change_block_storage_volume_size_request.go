/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ChangeBlockStorageVolumeSizeRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo"`

	// 블록스토리지사이즈
BlockStorageSize *int32 `json:"blockStorageSize"`
}
