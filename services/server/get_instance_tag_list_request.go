/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-21T05:59:54Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetInstanceTagListRequest struct {

	// 인스턴스번호리스트
InstanceNoList []*string `json:"instanceNoList,omitempty"`

	// 태그키리스트
TagKeyList []*string `json:"tagKeyList,omitempty"`

	// 태그값리스트
TagValueList []*string `json:"tagValueList,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}