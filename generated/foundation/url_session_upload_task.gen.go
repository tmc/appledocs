// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionUploadTask] class.
var uRLSessionUploadTaskClass = _URLSessionUploadTaskClass{objc.GetClass("NSURLSessionUploadTask")}

type _URLSessionUploadTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionUploadTask] class.
type IURLSessionUploadTask interface {
	IURLSessionDataTask
}

// A URL session task that uploads data to the network in a request body. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionUploadTask

type URLSessionUploadTask struct {
	URLSessionDataTask
}

// URLSessionUploadTaskFrom constructs a [URLSessionUploadTask] from an unsafe.Pointer.
//
// A URL session task that uploads data to the network in a request body.
func URLSessionUploadTaskFrom(ptr unsafe.Pointer) URLSessionUploadTask {
	return URLSessionUploadTask{
		URLSessionDataTask: URLSessionDataTaskFrom(ptr),
	}
}



