// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLDownload] class.
var URLDownloadClass objc.Class

func init() {
	URLDownloadClass = objc.GetClass("NSURLDownload")
}

type URLDownload struct {
	objc.ID
}

func URLDownloadFrom(ptr unsafe.Pointer) URLDownload {
	return URLDownload{
		ID: objc.ID(ptr),
	}
}




