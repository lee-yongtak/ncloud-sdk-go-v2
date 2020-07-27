/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * API version: 2020-07-06T00:54:55Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type GetNasVolumeInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 볼륨할당프로토콜유형코드
VolumeAllotmentProtocolTypeCode *string `json:"volumeAllotmentProtocolTypeCode,omitempty"`

	// 이벤트설정여부
IsEventConfiguration *bool `json:"isEventConfiguration,omitempty"`

	// 스냅샷설정여부
IsSnapshotConfiguration *bool `json:"isSnapshotConfiguration,omitempty"`

	// NAS볼륨인스턴스번호리스트
NasVolumeInstanceNoList []*string `json:"nasVolumeInstanceNoList,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 볼륨이름
VolumeName *string `json:"volumeName,omitempty"`

	// 정렬대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 정렬순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}