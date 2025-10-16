
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lineBreakMode] class.
var lineBreakModeClass _lineBreakModeClass

func init() {
	lineBreakModeClass = _lineBreakModeClass{objc.GetClass("lineBreakMode")}
}

type _lineBreakModeClass struct {
	objc.Class
}

// An interface definition for the [lineBreakMode] class.
type IlineBreakMode interface {
	ID() objc.ID
}

type lineBreakMode struct {
	id objc.ID
}

func lineBreakModeFrom(ptr unsafe.Pointer) lineBreakMode {
	return lineBreakMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lineBreakMode) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lineBreakModeClass) Alloc() lineBreakMode {
	rv := objc.Send[lineBreakMode](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lineBreakModeClass) New() lineBreakMode {
	rv := objc.Send[lineBreakMode](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlineBreakMode creates and returns a new initialized instance.
func NewlineBreakMode() lineBreakMode {
	return lineBreakModeClass.New()
}

// Init initializes the instance.
func (l_ lineBreakMode) Init() lineBreakMode {
	rv := objc.Send[lineBreakMode](l_.ID(), selInit)
	return rv
}
