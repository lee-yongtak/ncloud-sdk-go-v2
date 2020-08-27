/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-08-27T03:07:34Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RootPasswordServerInstanceParameter struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 개인키
PrivateKey *string `json:"privateKey,omitempty"`
}
