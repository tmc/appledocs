// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var bundleResourceRequestClass _BundleResourceRequestClass

func init() {
	bundleResourceRequestClass = _BundleResourceRequestClass{objc.GetClass("NSBundleResourceRequest")}
}

type _BundleResourceRequestClass struct {
	class objc.Class
}

type BundleResourceRequest struct {
	objc.ID
}

func BundleResourceRequestFrom(ptr unsafe.Pointer) BundleResourceRequest {
	return BundleResourceRequest{
		ID: objc.ID(ptr),
	}
}




