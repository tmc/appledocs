// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Font] class.
var (
	fontClass     _FontClass
	fontClassOnce sync.Once
)

func getFontClass() _FontClass {
	fontClassOnce.Do(func() {
		fontClass = _FontClass{objc.GetClass("NSFont")}
	})
	return fontClass
}

type _FontClass struct {
	class objc.Class
}

// An interface definition for the [Font] class.
type IFont interface {
	objectivec.IObject
}

// The representation of a font in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont

type Font struct {
	objectivec.Object
}

// FontFrom constructs a [Font] from an unsafe.Pointer.
//
// The representation of a font in an app.
func FontFrom(ptr unsafe.Pointer) Font {
	return Font{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (fc _FontClass) Alloc() Font {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (fc _FontClass) New() Font {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Font) Init() Font {
	rv := objc.Send[Font](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Font) Autorelease() Font {
	rv := objc.Send[Font](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFont creates a new Font instance.
func NewFont() Font {
	return getFontClass().New()
}




