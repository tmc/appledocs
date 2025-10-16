
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isAutomaticCustomizeTouchBarMenuItemEnabled] class.
var isAutomaticCustomizeTouchBarMenuItemEnabledClass _isAutomaticCustomizeTouchBarMenuItemEnabledClass

func init() {
	isAutomaticCustomizeTouchBarMenuItemEnabledClass = _isAutomaticCustomizeTouchBarMenuItemEnabledClass{objc.GetClass("isAutomaticCustomizeTouchBarMenuItemEnabled")}
}

type _isAutomaticCustomizeTouchBarMenuItemEnabledClass struct {
	objc.Class
}

// An interface definition for the [isAutomaticCustomizeTouchBarMenuItemEnabled] class.
type IisAutomaticCustomizeTouchBarMenuItemEnabled interface {
	ID() objc.ID
}

type isAutomaticCustomizeTouchBarMenuItemEnabled struct {
	id objc.ID
}

func isAutomaticCustomizeTouchBarMenuItemEnabledFrom(ptr unsafe.Pointer) isAutomaticCustomizeTouchBarMenuItemEnabled {
	return isAutomaticCustomizeTouchBarMenuItemEnabled{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isAutomaticCustomizeTouchBarMenuItemEnabled) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isAutomaticCustomizeTouchBarMenuItemEnabledClass) Alloc() isAutomaticCustomizeTouchBarMenuItemEnabled {
	rv := objc.Send[isAutomaticCustomizeTouchBarMenuItemEnabled](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isAutomaticCustomizeTouchBarMenuItemEnabledClass) New() isAutomaticCustomizeTouchBarMenuItemEnabled {
	rv := objc.Send[isAutomaticCustomizeTouchBarMenuItemEnabled](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisAutomaticCustomizeTouchBarMenuItemEnabled creates and returns a new initialized instance.
func NewisAutomaticCustomizeTouchBarMenuItemEnabled() isAutomaticCustomizeTouchBarMenuItemEnabled {
	return isAutomaticCustomizeTouchBarMenuItemEnabledClass.New()
}

// Init initializes the instance.
func (i_ isAutomaticCustomizeTouchBarMenuItemEnabled) Init() isAutomaticCustomizeTouchBarMenuItemEnabled {
	rv := objc.Send[isAutomaticCustomizeTouchBarMenuItemEnabled](i_.ID(), selInit)
	return rv
}
