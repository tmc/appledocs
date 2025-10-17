// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var URLSessionDownloadTaskClass = _URLSessionDownloadTaskClass{objc.GetClass("NSURLSessionDownloadTask")}

type _URLSessionDownloadTaskClass struct {
	class objc.Class
}

type URLSessionDownloadTask struct {
	objc.ID
}

func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		ID: objc.ID(ptr),
	}
}




