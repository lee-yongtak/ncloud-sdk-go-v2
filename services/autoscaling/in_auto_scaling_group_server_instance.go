/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-06-21T02:22:22Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type InAutoScalingGroupServerInstance struct {

	// 헬스상태
HealthStatus *CommonCode `json:"healthStatus,omitempty"`

	// 라이프사이클상태
LifecycleState *CommonCode `json:"lifecycleState,omitempty"`

	// 론치설정
LaunchConfiguration *LaunchConfiguration `json:"launchConfiguration,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 서버인스턴스명
ServerInstanceName *string `json:"serverInstanceName,omitempty"`
}
