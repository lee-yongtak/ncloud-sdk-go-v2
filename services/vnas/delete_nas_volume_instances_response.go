/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vnas/v2
 *
 * API version: 2020-08-27T01:44:24Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type DeleteNasVolumeInstancesResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

NasVolumeInstanceList []*NasVolumeInstance `json:"nasVolumeInstanceList,omitempty"`
}
