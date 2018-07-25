/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-04T02:02:10Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GetGlobalCdnPurgeHistoryListRequest struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo"`

	// 퍼지ID리스트
PurgeIdList []string `json:"purgeIdList,omitempty"`
}