// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionDataTask] class.
var URLSessionDataTaskClass objc.Class

func init() {
	URLSessionDataTaskClass = objc.GetClass("NSURLSessionDataTask")
}

type URLSessionDataTask struct {
	objc.ID
}

func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{
		ID: objc.ID(ptr),
	}
}




