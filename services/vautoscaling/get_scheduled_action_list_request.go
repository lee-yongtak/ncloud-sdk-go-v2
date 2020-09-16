/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type GetScheduledActionListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹이름
AutoScalingGroupName *string `json:"autoScalingGroupName"`

	// 스케쥴액션이름리스트
ScheduledActionNameList []*string `json:"scheduledActionNameList,omitempty"`

	// 스케쥴시작일시
StartTime *string `json:"startTime,omitempty"`

	// 스케쥴종료일시
EndTime *string `json:"endTime,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// 정렬대상
SortedBy *string `json:"sortedBy,omitempty"`

	// 정렬순서
SortingOrder *string `json:"sortingOrder,omitempty"`
}
