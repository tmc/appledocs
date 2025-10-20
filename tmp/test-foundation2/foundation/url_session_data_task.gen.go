// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLSessionDataTaskClass _URLSessionDataTaskClass

func init() {
	uRLSessionDataTaskClass = _URLSessionDataTaskClass{objc.GetClass("NSURLSessionDataTask")}
}

type _URLSessionDataTaskClass struct {
	class objc.Class
}

type URLSessionDataTask struct {
	objc.ID
}

func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{
		ID: objc.ID(ptr),
	}
}




