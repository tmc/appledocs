// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLDownloadClass _URLDownloadClass

func init() {
	uRLDownloadClass = _URLDownloadClass{objc.GetClass("NSURLDownload")}
}

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




