// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionDataTask] class.
var uRLSessionDataTaskClass = _URLSessionDataTaskClass{objc.GetClass("NSURLSessionDataTask")}

type _URLSessionDataTaskClass struct {
	class objc.Class
}

// A URL session task that returns downloaded data directly to the app in memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDataTask

type URLSessionDataTask struct {
	URLSessionTask
}

// URLSessionDataTaskFrom constructs a [URLSessionDataTask] from an unsafe.Pointer.
//
// A URL session task that returns downloaded data directly to the app in memory.
func URLSessionDataTaskFrom(ptr unsafe.Pointer) URLSessionDataTask {
	return URLSessionDataTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}



