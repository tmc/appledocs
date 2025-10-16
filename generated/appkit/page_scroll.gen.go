
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pageScroll] class.
var pageScrollClass _pageScrollClass

func init() {
	pageScrollClass = _pageScrollClass{objc.GetClass("pageScroll")}
}

type _pageScrollClass struct {
	objc.Class
}

// An interface definition for the [pageScroll] class.
type IpageScroll interface {
	ID() objc.ID
}

type pageScroll struct {
	id objc.ID
}

func pageScrollFrom(ptr unsafe.Pointer) pageScroll {
	return pageScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pageScroll) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pageScrollClass) Alloc() pageScroll {
	rv := objc.Send[pageScroll](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pageScrollClass) New() pageScroll {
	rv := objc.Send[pageScroll](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpageScroll creates and returns a new initialized instance.
func NewpageScroll() pageScroll {
	return pageScrollClass.New()
}

// Init initializes the instance.
func (p_ pageScroll) Init() pageScroll {
	rv := objc.Send[pageScroll](p_.ID(), selInit)
	return rv
}
