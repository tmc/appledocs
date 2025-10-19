// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontManager] class.
var (
	fontManagerClass     _FontManagerClass
	fontManagerClassOnce sync.Once
)

func getFontManagerClass() _FontManagerClass {
	fontManagerClassOnce.Do(func() {
		fontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}
	})
	return fontManagerClass
}

type _FontManagerClass struct {
	class objc.Class
}

// An interface definition for the [FontManager] class.
type IFontManager interface {
	objectivec.IObject
}

// The center of activity for the font-conversion system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager
type FontManager struct {
	objectivec.Object
}

// FontManagerFrom constructs a [FontManager] from an unsafe.Pointer.
//
// The center of activity for the font-conversion system.
func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontManagerClass) New() FontManager {
	rv := objc.Send[FontManager](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontManager) Init() FontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontManager) Autorelease() FontManager {
	rv := objc.Send[FontManager](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontManager creates a new FontManager instance.
func NewFontManager() FontManager {
	return getFontManagerClass().New()
}




