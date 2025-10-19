// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AdaptiveImageGlyph] class.
var (
	adaptiveImageGlyphClass     _AdaptiveImageGlyphClass
	adaptiveImageGlyphClassOnce sync.Once
)

func getAdaptiveImageGlyphClass() _AdaptiveImageGlyphClass {
	adaptiveImageGlyphClassOnce.Do(func() {
		adaptiveImageGlyphClass = _AdaptiveImageGlyphClass{objc.GetClass("NSAdaptiveImageGlyph")}
	})
	return adaptiveImageGlyphClass
}

type _AdaptiveImageGlyphClass struct {
	class objc.Class
}

// An interface definition for the [AdaptiveImageGlyph] class.
type IAdaptiveImageGlyph interface {
	objectivec.IObject
}

// A data object for an emoji-like image that can appear in attributed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph

type AdaptiveImageGlyph struct {
	objectivec.Object
}

// AdaptiveImageGlyphFrom constructs a [AdaptiveImageGlyph] from an unsafe.Pointer.
//
// A data object for an emoji-like image that can appear in attributed text.
func AdaptiveImageGlyphFrom(ptr unsafe.Pointer) AdaptiveImageGlyph {
	return AdaptiveImageGlyph{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AdaptiveImageGlyphClass) Alloc() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AdaptiveImageGlyphClass) New() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdaptiveImageGlyph) Init() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdaptiveImageGlyph) Autorelease() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdaptiveImageGlyph creates a new AdaptiveImageGlyph instance.
func NewAdaptiveImageGlyph() AdaptiveImageGlyph {
	return getAdaptiveImageGlyphClass().New()
}




