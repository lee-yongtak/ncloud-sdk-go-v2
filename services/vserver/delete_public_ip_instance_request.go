/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DeletePublicIpInstanceRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 공인IP인스턴스번호
PublicIpInstanceNo *string `json:"publicIpInstanceNo"`
}
