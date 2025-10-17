
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontAssetRequest] class.
var FontAssetRequestClass _FontAssetRequestClass

func init() {
	FontAssetRequestClass = _FontAssetRequestClass{objc.GetClass("NSFontAssetRequest")}
}

type _FontAssetRequestClass struct {
	objc.Class
}

// An interface definition for the [FontAssetRequest] class.
type IFontAssetRequest interface {
	ID() objc.ID
}

type FontAssetRequest struct {
	id objc.ID
}

func FontAssetRequestFrom(ptr unsafe.Pointer) FontAssetRequest {
	return FontAssetRequest{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FontAssetRequest) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontAssetRequestClass) Alloc() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontAssetRequestClass) New() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFontAssetRequest creates and returns a new initialized instance.
func NewFontAssetRequest() FontAssetRequest {
	return FontAssetRequestClass.New()
}

// Init initializes the instance.
func (f_ FontAssetRequest) Init() FontAssetRequest {
	rv := objc.Send[FontAssetRequest](f_.ID(), selInit)
	return rv
}
