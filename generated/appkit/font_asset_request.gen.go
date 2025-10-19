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

// An interface definition for the [FontAssetRequest] class.
type IFontAssetRequest interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (fc _FontAssetRequestClass) Alloc() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (fc _FontAssetRequestClass) New() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontAssetRequest) Init() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontAssetRequest) Autorelease() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontAssetRequest creates a new FontAssetRequest instance.
func NewFontAssetRequest() FontAssetRequest {
	return fontAssetRequestClass.New()
}




