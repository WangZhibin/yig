/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"git.letv.cn/yig/yig/api/datatype"
	"git.letv.cn/yig/yig/iam"
	"git.letv.cn/yig/yig/meta"
	"io"
)

// ObjectLayer implements primitives for object API layer.
type ObjectLayer interface {
	// Bucket operations.
	MakeBucket(bucket string, acl datatype.Acl, credential iam.Credential) error
	SetBucketAcl(bucket string, acl datatype.Acl, credential iam.Credential) error
	SetBucketCors(bucket string, cors datatype.Cors, credential iam.Credential) error
	SetBucketVersioning(bucket string, versioning datatype.Versioning, credential iam.Credential) error
	DeleteBucketCors(bucket string, credential iam.Credential) error
	GetBucketVersioning(bucket string, credential iam.Credential) (datatype.Versioning, error)
	GetBucketCors(bucket string, credential iam.Credential) (datatype.Cors, error)
	GetBucket(bucketName string) (bucket meta.Bucket, err error) // For INTERNAL USE ONLY
	GetBucketInfo(bucket string, credential iam.Credential) (bucketInfo meta.Bucket, err error)
	ListBuckets(credential iam.Credential) (buckets []meta.Bucket, err error)
	DeleteBucket(bucket string, credential iam.Credential) error
	ListObjects(credential iam.Credential, bucket string,
		request datatype.ListObjectsRequest) (result meta.ListObjectsInfo, err error)
	ListVersionedObjects(credential iam.Credential, bucket string,
		request datatype.ListObjectsRequest) (result meta.VersionedListObjectsInfo, err error)

	// Object operations.
	GetObject(object meta.Object, startOffset int64, length int64, writer io.Writer) (err error)
	GetObjectInfo(bucket, object, version string) (objInfo meta.Object, err error)
	PutObject(bucket, object string, size int64, data io.Reader,
		metadata map[string]string, acl datatype.Acl,
		sse datatype.SseRequest) (result datatype.PutObjectResult, err error)
	CopyObject(object meta.Object, source io.Reader,
		credential iam.Credential) (result datatype.PutObjectResult, err error)
	SetObjectAcl(bucket string, object string, version string, acl datatype.Acl,
		credential iam.Credential) error
	DeleteObject(bucket, object, version string, credential iam.Credential) (datatype.DeleteObjectResult,
		error)

	// Multipart operations.
	ListMultipartUploads(credential iam.Credential, bucket string,
		request datatype.ListUploadsRequest) (result datatype.ListMultipartUploadsResponse, err error)
	NewMultipartUpload(credential iam.Credential, bucket, object string,
		metadata map[string]string, acl datatype.Acl) (uploadID string, err error)
	PutObjectPart(bucket, object, uploadID string, partID int, size int64,
		data io.Reader, md5Hex string) (md5 string, err error)
	CopyObjectPart(bucketName, objectName, uploadId string, partId int, size int64, data io.Reader,
		credential iam.Credential) (result datatype.PutObjectResult, err error)
	ListObjectParts(credential iam.Credential, bucket, object string,
		request datatype.ListPartsRequest) (result datatype.ListPartsResponse, err error)
	AbortMultipartUpload(credential iam.Credential, bucket, object, uploadID string) error
	CompleteMultipartUpload(credential iam.Credential, bucket, object, uploadID string,
		uploadedParts []meta.CompletePart) (result datatype.CompleteMultipartResult, err error)
}
