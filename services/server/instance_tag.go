/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type InstanceTag struct {

	// 인스턴스번호
InstanceNo *string `json:"instanceNo,omitempty"`

	// 인스턴스유형
InstanceType *CommonCode `json:"instanceType,omitempty"`

	// 태그키
TagKey *string `json:"tagKey,omitempty"`

	// 태그값
TagValue *string `json:"tagValue,omitempty"`
}
