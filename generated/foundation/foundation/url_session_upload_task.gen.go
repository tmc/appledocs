// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionUploadTask] class.
var URLSessionUploadTaskClass objc.Class

func init() {
	URLSessionUploadTaskClass = objc.GetClass("NSURLSessionUploadTask")
}

type URLSessionUploadTask struct {
	objc.ID
}

func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		ID: objc.ID(ptr),
	}
}




