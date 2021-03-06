/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-11-01T03:57:11Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type DmsFile struct {

	// 파일길이
FileLength *int64 `json:"fileLength,omitempty"`

	// 최종작성시각
LastWriteTime *string `json:"lastWriteTime,omitempty"`

	// 파일이름
FileName *string `json:"fileName,omitempty"`
}
