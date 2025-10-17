// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontAssetRequest] class.
var FontAssetRequestClass objc.Class

func init() {
	FontAssetRequestClass = objc.GetClass("NSFontAssetRequest")
}

type FontAssetRequest struct {
	objc.ID
}

func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{
		ID: objc.ID(ptr),
	}
}



