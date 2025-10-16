
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titlebarAppearsTransparent] class.
var titlebarAppearsTransparentClass _titlebarAppearsTransparentClass

func init() {
	titlebarAppearsTransparentClass = _titlebarAppearsTransparentClass{objc.GetClass("titlebarAppearsTransparent")}
}

type _titlebarAppearsTransparentClass struct {
	objc.Class
}

// An interface definition for the [titlebarAppearsTransparent] class.
type ItitlebarAppearsTransparent interface {
	ID() objc.ID
}

type titlebarAppearsTransparent struct {
	id objc.ID
}

func titlebarAppearsTransparentFrom(ptr unsafe.Pointer) titlebarAppearsTransparent {
	return titlebarAppearsTransparent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titlebarAppearsTransparent) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titlebarAppearsTransparentClass) Alloc() titlebarAppearsTransparent {
	rv := objc.Send[titlebarAppearsTransparent](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titlebarAppearsTransparentClass) New() titlebarAppearsTransparent {
	rv := objc.Send[titlebarAppearsTransparent](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitlebarAppearsTransparent creates and returns a new initialized instance.
func NewtitlebarAppearsTransparent() titlebarAppearsTransparent {
	return titlebarAppearsTransparentClass.New()
}

// Init initializes the instance.
func (t_ titlebarAppearsTransparent) Init() titlebarAppearsTransparent {
	rv := objc.Send[titlebarAppearsTransparent](t_.ID(), selInit)
	return rv
}
