
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PageLayout] class.
var PageLayoutClass _PageLayoutClass

func init() {
	PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}
}

type _PageLayoutClass struct {
	objc.Class
}

// An interface definition for the [PageLayout] class.
type IPageLayout interface {
	ID() objc.ID
}

type PageLayout struct {
	id objc.ID
}

func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PageLayout) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPageLayout creates and returns a new initialized instance.
func NewPageLayout() PageLayout {
	return PageLayoutClass.New()
}

// Init initializes the instance.
func (p_ PageLayout) Init() PageLayout {
	rv := objc.Send[PageLayout](p_.ID(), selInit)
	return rv
}
