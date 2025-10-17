// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BundleResourceRequest] class.
var BundleResourceRequestClass objc.Class

func init() {
	BundleResourceRequestClass = objc.GetClass("NSBundleResourceRequest")
}

type BundleResourceRequest struct {
	objc.ID
}

func BundleResourceRequestFrom(ptr unsafe.Pointer) BundleResourceRequest {
	return BundleResourceRequest{
		ID: objc.ID(ptr),
	}
}




