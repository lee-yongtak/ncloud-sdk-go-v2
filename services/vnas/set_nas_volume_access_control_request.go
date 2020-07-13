/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * API version: 2020-07-06T00:54:55Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type SetNasVolumeAccessControlRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// NAS볼륨인스턴스번호
NasVolumeInstanceNo *string `json:"nasVolumeInstanceNo"`

	// 서버인스턴스번호리스트
ServerInstanceNoList []*string `json:"serverInstanceNoList,omitempty"`
}
