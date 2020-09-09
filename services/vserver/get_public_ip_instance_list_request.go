/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.beta-apigw.ntruss.com/vserver/v2
 *
 * API version: 2020-09-04T05:39:38Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type GetPublicIpInstanceListRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 공인IP인스턴스번호리스트
PublicIpInstanceNoList []*string `json:"publicIpInstanceNoList,omitempty"`

	// 공인IP주소
PublicIp *string `json:"publicIp,omitempty"`

	// 비공인IP주소
PrivateIp *string `json:"privateIp,omitempty"`

	// 할당여부
IsAssociated *bool `json:"isAssociated,omitempty"`

	// 서버이름
ServerName *string `json:"serverName,omitempty"`

	// 공인IP인스턴스상태코드
PublicIpInstanceStatusCode *string `json:"publicIpInstanceStatusCode,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`
}
