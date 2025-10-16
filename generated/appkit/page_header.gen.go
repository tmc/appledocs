
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pageHeader] class.
var pageHeaderClass _pageHeaderClass

func init() {
	pageHeaderClass = _pageHeaderClass{objc.GetClass("pageHeader")}
}

type _pageHeaderClass struct {
	objc.Class
}

// An interface definition for the [pageHeader] class.
type IpageHeader interface {
	ID() objc.ID
}

type pageHeader struct {
	id objc.ID
}

func pageHeaderFrom(ptr unsafe.Pointer) pageHeader {
	return pageHeader{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pageHeader) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pageHeaderClass) Alloc() pageHeader {
	rv := objc.Send[pageHeader](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pageHeaderClass) New() pageHeader {
	rv := objc.Send[pageHeader](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpageHeader creates and returns a new initialized instance.
func NewpageHeader() pageHeader {
	return pageHeaderClass.New()
}

// Init initializes the instance.
func (p_ pageHeader) Init() pageHeader {
	rv := objc.Send[pageHeader](p_.ID(), selInit)
	return rv
}
