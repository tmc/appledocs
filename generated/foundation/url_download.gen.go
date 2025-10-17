// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLDownload] class.
var uRLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}

type _URLDownloadClass struct {
	class objc.Class
}

// An interface definition for the [URLDownload] class.
type IURLDownload interface {
	objectivec.IObject
}

// An object that downloads a resource asynchronously and saves the data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLDownload

type URLDownload struct {
	objectivec.Object
}

// URLDownloadFrom constructs a [URLDownload] from an unsafe.Pointer.
//
// An object that downloads a resource asynchronously and saves the data to a file.
func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{objectivec.Object{objc.ID(ptr)}}
}



