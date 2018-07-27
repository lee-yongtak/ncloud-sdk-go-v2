/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * API version: 2018-06-21T02:19:18Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

import (
	"bytes"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/hmac"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the loadbalancer API v2018-06-21T02:19:18Z
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	V2Api *V2ApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.V2Api = (*V2ApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	path string,
	method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	// APIKey Authentication
	if auth := c.cfg.APIKey; auth != nil {
		timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		signer := hmac.NewSigner(auth.SecretKey, crypto.SHA256)
		signature, _ := signer.Sign(method, path, auth.AccessKey, timestamp)

		localVarRequest.Header.Add("x-ncp-apigw-timestamp", timestamp)
		localVarRequest.Header.Add("x-ncp-iam-access-key", auth.AccessKey)
		localVarRequest.Header.Add("x-ncp-apigw-signature-v1", signature)
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func toLowerFirstChar(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	result := "responseFormatType=json"
	s := reflect.ValueOf(body).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if !f.IsNil() {
			key := toLowerFirstChar(typeOfT.Field(i).Name)
			if f.Kind() == reflect.Ptr {
				switch f.Type().String() {
				case "*string":
					if key == "userData" {
						result += fmt.Sprintf("&%s=%s", key, base64.StdEncoding.EncodeToString([]byte(ncloud.StringValue(f.Interface().(*string)))))
					} else {
						result += fmt.Sprintf("&%s=%s", key, url.QueryEscape(ncloud.StringValue(f.Interface().(*string))))
					}
				case "*bool":
					result += fmt.Sprintf("&%s=%t", key, ncloud.BoolValue(f.Interface().(*bool)))
				case "*int":
					result += fmt.Sprintf("&%s=%d", key, ncloud.IntValue(f.Interface().(*int)))
				case "*int32":
					result += fmt.Sprintf("&%s=%d", key, ncloud.Int32Value(f.Interface().(*int32)))
				case "*int64":
					result += fmt.Sprintf("&%s=%d", key, ncloud.Int64Value(f.Interface().(*int64)))
				case "*float32":
					result += fmt.Sprintf("&%s=%f", key, ncloud.Float32Value(f.Interface().(*float32)))
				}
			} else if f.Kind() == reflect.Slice {
				for i := 0; i < f.Len(); i++ {
					item := f.Index(i)

					if item.Elem().Kind() == reflect.Struct {
						item := item.Elem()
						typeOfSubItem := item.Type()

						for j := 0; j < item.NumField(); j++ {
							subItem := item.Field(j)
							subKey := toLowerFirstChar(typeOfSubItem.Field(j).Name)

							switch subItem.Type().String() {
							case "*string":
								if key == "userData" {
									result += fmt.Sprintf("&%s.%d.%s=%s", key, i+1, subKey, base64.StdEncoding.EncodeToString([]byte(ncloud.StringValue(subItem.Interface().(*string)))))
								} else {
									result += fmt.Sprintf("&%s.%d.%s=%s", key, i+1, subKey, url.QueryEscape(ncloud.StringValue(subItem.Interface().(*string))))
								}
							case "*bool":
								result += fmt.Sprintf("&%s.%d.%s=%t", key, i+1, subKey, ncloud.BoolValue(subItem.Interface().(*bool)))
							case "*int":
								result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, subKey, ncloud.IntValue(subItem.Interface().(*int)))
							case "*int32":
								result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, subKey, ncloud.Int32Value(subItem.Interface().(*int32)))
							case "*int64":
								result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, subKey, ncloud.Int64Value(subItem.Interface().(*int64)))
							case "*float32":
								result += fmt.Sprintf("&%s.%d.%s=%f", key, i+1, subKey, ncloud.Float32Value(subItem.Interface().(*float32)))
							}
						}
					} else {
						switch item.Type().String() {
						case "*string":
							if key == "userData" {
								result += fmt.Sprintf("&%s.%d=%s", key, i+1, base64.StdEncoding.EncodeToString([]byte(*item.Interface().(*string))))
							} else {
								result += fmt.Sprintf("&%s.%d=%s", key, i+1, url.QueryEscape(*item.Interface().(*string)))
							}
						case "*bool":
							result += fmt.Sprintf("&%s.%d.%s=%t", key, i+1, ncloud.BoolValue(item.Interface().(*bool)))
						case "*int":
							result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, item.Interface().(*int))
						case "*int32":
							result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, item.Interface().(*int32))
						case "*int64":
							result += fmt.Sprintf("&%s.%d.%s=%d", key, i+1, item.Interface().(*int64))
						case "*float32":
							result += fmt.Sprintf("&%s.%d.%s=%f", key, i+1, item.Interface().(*float32))
						}
					}
				}
			}
		}
	}

	if err != nil {
		return nil, err
	}

	bodyBuf.WriteString(result)

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}
