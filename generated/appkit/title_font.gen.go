
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titleFont] class.
var titleFontClass _titleFontClass

func init() {
	titleFontClass = _titleFontClass{objc.GetClass("titleFont")}
}

type _titleFontClass struct {
	objc.Class
}

// An interface definition for the [titleFont] class.
type ItitleFont interface {
	ID() objc.ID
}

type titleFont struct {
	id objc.ID
}

func titleFontFrom(ptr unsafe.Pointer) titleFont {
	return titleFont{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titleFont) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titleFontClass) Alloc() titleFont {
	rv := objc.Send[titleFont](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titleFontClass) New() titleFont {
	rv := objc.Send[titleFont](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitleFont creates and returns a new initialized instance.
func NewtitleFont() titleFont {
	return titleFontClass.New()
}

// Init initializes the instance.
func (t_ titleFont) Init() titleFont {
	rv := objc.Send[titleFont](t_.ID(), selInit)
	return rv
}
