// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontAssetRequest] class.
var fontAssetRequestClass = _FontAssetRequestClass{objc.GetClass("NSFontAssetRequest")}

type _FontAssetRequestClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest

type FontAssetRequest struct {
	objectivec.Object
}

// FontAssetRequestFrom constructs a [FontAssetRequest] from an unsafe.Pointer.
func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{objectivec.Object{objc.ID(ptr)}}
}



