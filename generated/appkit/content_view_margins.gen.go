
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentViewMargins] class.
var contentViewMarginsClass _contentViewMarginsClass

func init() {
	contentViewMarginsClass = _contentViewMarginsClass{objc.GetClass("contentViewMargins")}
}

type _contentViewMarginsClass struct {
	objc.Class
}

// An interface definition for the [contentViewMargins] class.
type IcontentViewMargins interface {
	ID() objc.ID
}

type contentViewMargins struct {
	id objc.ID
}

func contentViewMarginsFrom(ptr unsafe.Pointer) contentViewMargins {
	return contentViewMargins{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentViewMargins) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentViewMarginsClass) Alloc() contentViewMargins {
	rv := objc.Send[contentViewMargins](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentViewMarginsClass) New() contentViewMargins {
	rv := objc.Send[contentViewMargins](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentViewMargins creates and returns a new initialized instance.
func NewcontentViewMargins() contentViewMargins {
	return contentViewMarginsClass.New()
}

// Init initializes the instance.
func (c_ contentViewMargins) Init() contentViewMargins {
	rv := objc.Send[contentViewMargins](c_.ID(), selInit)
	return rv
}
