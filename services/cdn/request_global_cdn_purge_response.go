/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-04T02:02:10Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type RequestGlobalCdnPurgeResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

GlobalCdnPurgeHistoryList []GlobalCdnPurgeHistory `json:"globalCdnPurgeHistoryList,omitempty"`
}