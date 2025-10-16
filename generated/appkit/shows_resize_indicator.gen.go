
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [showsResizeIndicator] class.
var showsResizeIndicatorClass _showsResizeIndicatorClass

func init() {
	showsResizeIndicatorClass = _showsResizeIndicatorClass{objc.GetClass("showsResizeIndicator")}
}

type _showsResizeIndicatorClass struct {
	objc.Class
}

// An interface definition for the [showsResizeIndicator] class.
type IshowsResizeIndicator interface {
	ID() objc.ID
}

type showsResizeIndicator struct {
	id objc.ID
}

func showsResizeIndicatorFrom(ptr unsafe.Pointer) showsResizeIndicator {
	return showsResizeIndicator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ showsResizeIndicator) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _showsResizeIndicatorClass) Alloc() showsResizeIndicator {
	rv := objc.Send[showsResizeIndicator](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _showsResizeIndicatorClass) New() showsResizeIndicator {
	rv := objc.Send[showsResizeIndicator](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshowsResizeIndicator creates and returns a new initialized instance.
func NewshowsResizeIndicator() showsResizeIndicator {
	return showsResizeIndicatorClass.New()
}

// Init initializes the instance.
func (s_ showsResizeIndicator) Init() showsResizeIndicator {
	rv := objc.Send[showsResizeIndicator](s_.ID(), selInit)
	return rv
}
