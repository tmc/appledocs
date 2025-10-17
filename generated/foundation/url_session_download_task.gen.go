// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionDownloadTask] class.
var uRLSessionDownloadTaskClass = _URLSessionDownloadTaskClass{objc.GetClass("NSURLSessionDownloadTask")}

type _URLSessionDownloadTaskClass struct {
	class objc.Class
}

// A URL session task that stores downloaded data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionDownloadTask

type URLSessionDownloadTask struct {
	URLSessionTask
}

// URLSessionDownloadTaskFrom constructs a [URLSessionDownloadTask] from an unsafe.Pointer.
//
// A URL session task that stores downloaded data to a file.
func URLSessionDownloadTaskFrom(ptr unsafe.Pointer) URLSessionDownloadTask {
	return URLSessionDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}



