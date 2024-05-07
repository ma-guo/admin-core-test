// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 列举出指定 UploadId 所属任务所有已经上传成功的分片
package resumable_upload_v2_list_parts

import (
	"encoding/json"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
	uptoken "github.com/qiniu/go-sdk/v7/storagev2/uptoken"
)

// 调用 API 所用的请求
type Request struct {
	BucketName       string           // 存储空间名称
	ObjectName       *string          // 对象名称
	UploadId         string           // 在服务端申请的 Multipart Upload 任务 id
	MaxParts         int64            // 响应中的最大分片数目。默认值：1000，最大值：1000
	PartNumberMarker int64            // 指定列举的起始位置，只有 partNumber 值大于该参数的分片会被列出
	UpToken          uptoken.Provider // 上传凭证，如果为空，则使用 HTTPClientOptions 中的 UpToken
}

// 获取 API 所用的响应
type Response struct {
	UploadId         string      // 在服务端申请的 Multipart Upload 任务 id
	ExpiredAt        int64       // UploadId 的过期时间 UNIX 时间戳，过期之后 UploadId 不可用
	PartNumberMarker int64       // 下次继续列举的起始位置，0 表示列举结束，没有更多分片
	Parts            ListedParts // 返回所有已经上传成功的分片信息
}

// 单个已经上传的分片信息
type ListedPartInfo struct {
	Size       int64  // 分片大小
	Etag       string // 分片内容的 etag
	PartNumber int64  // 每一个上传的分片都有一个标识它的号码
	PutTime    int64  // 分片上传时间 UNIX 时间戳
}
type jsonListedPartInfo struct {
	Size       int64  `json:"size,omitempty"` // 分片大小
	Etag       string `json:"etag"`           // 分片内容的 etag
	PartNumber int64  `json:"partNumber"`     // 每一个上传的分片都有一个标识它的号码
	PutTime    int64  `json:"putTime"`        // 分片上传时间 UNIX 时间戳
}

func (j *ListedPartInfo) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonListedPartInfo{Size: j.Size, Etag: j.Etag, PartNumber: j.PartNumber, PutTime: j.PutTime})
}
func (j *ListedPartInfo) UnmarshalJSON(data []byte) error {
	var nj jsonListedPartInfo
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Size = nj.Size
	j.Etag = nj.Etag
	j.PartNumber = nj.PartNumber
	j.PutTime = nj.PutTime
	return nil
}
func (j *ListedPartInfo) validate() error {
	if j.Etag == "" {
		return errors.MissingRequiredFieldError{Name: "Etag"}
	}
	if j.PartNumber == 0 {
		return errors.MissingRequiredFieldError{Name: "PartNumber"}
	}
	if j.PutTime == 0 {
		return errors.MissingRequiredFieldError{Name: "PutTime"}
	}
	return nil
}

// 所有已经上传的分片信息
type ListedParts = []ListedPartInfo

// 返回所有已经上传成功的分片信息
type ListedPartsResponse = Response
type jsonResponse struct {
	UploadId         string      `json:"uploadId"`         // 在服务端申请的 Multipart Upload 任务 id
	ExpiredAt        int64       `json:"expireAt"`         // UploadId 的过期时间 UNIX 时间戳，过期之后 UploadId 不可用
	PartNumberMarker int64       `json:"partNumberMarker"` // 下次继续列举的起始位置，0 表示列举结束，没有更多分片
	Parts            ListedParts `json:"parts"`            // 返回所有已经上传成功的分片信息
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{UploadId: j.UploadId, ExpiredAt: j.ExpiredAt, PartNumberMarker: j.PartNumberMarker, Parts: j.Parts})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.UploadId = nj.UploadId
	j.ExpiredAt = nj.ExpiredAt
	j.PartNumberMarker = nj.PartNumberMarker
	j.Parts = nj.Parts
	return nil
}
func (j *Response) validate() error {
	if j.UploadId == "" {
		return errors.MissingRequiredFieldError{Name: "UploadId"}
	}
	if j.ExpiredAt == 0 {
		return errors.MissingRequiredFieldError{Name: "ExpiredAt"}
	}
	if j.PartNumberMarker == 0 {
		return errors.MissingRequiredFieldError{Name: "PartNumberMarker"}
	}
	if len(j.Parts) == 0 {
		return errors.MissingRequiredFieldError{Name: "Parts"}
	}
	for _, value := range j.Parts {
		if err := value.validate(); err != nil {
			return err
		}
	}
	return nil
}
