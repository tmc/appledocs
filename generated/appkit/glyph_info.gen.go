
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlyphInfo] class.
var GlyphInfoClass _GlyphInfoClass

func init() {
	GlyphInfoClass = _GlyphInfoClass{objc.GetClass("NSGlyphInfo")}
}

type _GlyphInfoClass struct {
	objc.Class
}

// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	ID() objc.ID
}

type GlyphInfo struct {
	id objc.ID
}

func GlyphInfoFrom(ptr unsafe.Pointer) GlyphInfo {
	return GlyphInfo{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GlyphInfo) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GlyphInfoClass) Alloc() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GlyphInfoClass) New() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGlyphInfo creates and returns a new initialized instance.
func NewGlyphInfo() GlyphInfo {
	return GlyphInfoClass.New()
}

// Init initializes the instance.
func (g_ GlyphInfo) Init() GlyphInfo {
	rv := objc.Send[GlyphInfo](g_.ID(), selInit)
	return rv
}
