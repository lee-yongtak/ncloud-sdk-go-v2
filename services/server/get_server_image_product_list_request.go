/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetServerImageProductListRequest struct {

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 플랫폼유형코드리스트
PlatformTypeCodeList []*string `json:"platformTypeCodeList,omitempty"`

	// 블록스토리지사이즈
BlockStorageSize *int32 `json:"blockStorageSize,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// 인프라자원상세구분코드
InfraResourceDetailTypeCode *string `json:"infraResourceDetailTypeCode,omitempty"`
}
