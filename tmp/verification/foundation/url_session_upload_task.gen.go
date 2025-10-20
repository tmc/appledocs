// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLSessionUploadTaskClass _URLSessionUploadTaskClass

func init() {
	uRLSessionUploadTaskClass = _URLSessionUploadTaskClass{objc.GetClass("NSURLSessionUploadTask")}
}

type _URLSessionUploadTaskClass struct {
	class objc.Class
}

type URLSessionUploadTask struct {
	objc.ID
}

func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		ID: objc.ID(ptr),
	}
}




