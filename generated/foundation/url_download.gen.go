// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLDownload] class.
var URLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}

type _URLDownloadClass struct {
	class objc.Class
}

type URLDownload struct {
	objc.ID
}

func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{
		ID: objc.ID(ptr),
	}
}




