// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BundleResourceRequest] class.
var bundleResourceRequestClass = _BundleResourceRequestClass{objc.GetClass("NSBundleResourceRequest")}

type _BundleResourceRequestClass struct {
	class objc.Class
}

// A resource manager you use to download content hosted on the App Store at the time your app needs it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest

type BundleResourceRequest struct {
	objectivec.Object
}

// BundleResourceRequestFrom constructs a [BundleResourceRequest] from an unsafe.Pointer.
//
// A resource manager you use to download content hosted on the App Store at the time your app needs it.
func BundleResourceRequestFrom(ptr unsafe.Pointer) BundleResourceRequest {
	return BundleResourceRequest{objectivec.Object{objc.ID(ptr)}}
}



