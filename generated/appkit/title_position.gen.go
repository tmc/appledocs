
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titlePosition] class.
var titlePositionClass _titlePositionClass

func init() {
	titlePositionClass = _titlePositionClass{objc.GetClass("titlePosition")}
}

type _titlePositionClass struct {
	objc.Class
}

// An interface definition for the [titlePosition] class.
type ItitlePosition interface {
	ID() objc.ID
}

type titlePosition struct {
	id objc.ID
}

func titlePositionFrom(ptr unsafe.Pointer) titlePosition {
	return titlePosition{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titlePosition) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titlePositionClass) Alloc() titlePosition {
	rv := objc.Send[titlePosition](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titlePositionClass) New() titlePosition {
	rv := objc.Send[titlePosition](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitlePosition creates and returns a new initialized instance.
func NewtitlePosition() titlePosition {
	return titlePositionClass.New()
}

// Init initializes the instance.
func (t_ titlePosition) Init() titlePosition {
	rv := objc.Send[titlePosition](t_.ID(), selInit)
	return rv
}
