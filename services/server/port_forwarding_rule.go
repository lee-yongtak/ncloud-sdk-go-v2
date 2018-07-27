/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-05T07:55:47Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type PortForwardingRule struct {

	// 포트포워딩외부포트
PortForwardingExternalPort *int32 `json:"portForwardingExternalPort,omitempty"`

	// 포트포워딩내부포트
PortForwardingInternalPort *int32 `json:"portForwardingInternalPort,omitempty"`

	// 서버인스턴스
ServerInstance *ServerInstance `json:"serverInstance,omitempty"`
}
