/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-07-09T01:14:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type LoginKey struct {

	// 키이름
KeyName *string `json:"keyName,omitempty"`

	// 핑거프린트
Fingerprint *string `json:"fingerprint,omitempty"`

	// 공개키
PublicKey *string `json:"publicKey,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`
}