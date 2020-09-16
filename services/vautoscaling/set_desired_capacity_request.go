/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type SetDesiredCapacityRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹이름
AutoScalingGroupName *string `json:"autoScalingGroupName"`

	// 기대용량
DesiredCapacity *int32 `json:"desiredCapacity"`
}
