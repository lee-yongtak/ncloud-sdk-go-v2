/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-06-21T02:22:22Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type CreateLaunchConfigurationRequest struct {

	// 론치설정명
LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`

	// 소프트웨어상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode,omitempty"`

	// 회원서버이미지번호
MemberServerImageNo *string `json:"memberServerImageNo,omitempty"`

	// ACG설정번호리스트
AccessControlGroupConfigurationNoList []*string `json:"accessControlGroupConfigurationNoList,omitempty"`

	// 로그인키명
LoginKeyName *string `json:"loginKeyName,omitempty"`

	// 사용자데이터
UserData *string `json:"userData,omitempty"`
}
