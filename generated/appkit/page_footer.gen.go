
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pageFooter] class.
var pageFooterClass _pageFooterClass

func init() {
	pageFooterClass = _pageFooterClass{objc.GetClass("pageFooter")}
}

type _pageFooterClass struct {
	objc.Class
}

// An interface definition for the [pageFooter] class.
type IpageFooter interface {
	ID() objc.ID
}

type pageFooter struct {
	id objc.ID
}

func pageFooterFrom(ptr unsafe.Pointer) pageFooter {
	return pageFooter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pageFooter) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pageFooterClass) Alloc() pageFooter {
	rv := objc.Send[pageFooter](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pageFooterClass) New() pageFooter {
	rv := objc.Send[pageFooter](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpageFooter creates and returns a new initialized instance.
func NewpageFooter() pageFooter {
	return pageFooterClass.New()
}

// Init initializes the instance.
func (p_ pageFooter) Init() pageFooter {
	rv := objc.Send[pageFooter](p_.ID(), selInit)
	return rv
}
