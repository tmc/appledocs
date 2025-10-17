// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var URLSessionDownloadTaskClass objc.Class

func init() {
	URLSessionDownloadTaskClass = objc.GetClass("NSURLSessionDownloadTask")
}

type URLSessionDownloadTask struct {
	objc.ID
}

func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		ID: objc.ID(ptr),
	}
}



