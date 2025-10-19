// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GlyphInfo] class.
var glyphInfoClass = _GlyphInfoClass{objc.GetClass("NSGlyphInfo")}

type _GlyphInfoClass struct {
	class objc.Class
}

// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	objectivec.IObject
}

// A glyph attribute in an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo

type GlyphInfo struct {
	objectivec.Object
}

// GlyphInfoFrom constructs a [GlyphInfo] from an unsafe.Pointer.
//
// A glyph attribute in an attributed string.
func GlyphInfoFrom(ptr unsafe.Pointer) GlyphInfo {
	return GlyphInfo{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (gc _GlyphInfoClass) Alloc() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GlyphInfoClass) New() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlyphInfo) Init() GlyphInfo {
	rv := objc.Send[GlyphInfo](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlyphInfo) Autorelease() GlyphInfo {
	rv := objc.Send[GlyphInfo](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlyphInfo creates a new GlyphInfo instance.
func NewGlyphInfo() GlyphInfo {
	return glyphInfoClass.New()
}




