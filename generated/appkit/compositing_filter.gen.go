
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [compositingFilter] class.
var compositingFilterClass _compositingFilterClass

func init() {
	compositingFilterClass = _compositingFilterClass{objc.GetClass("compositingFilter")}
}

type _compositingFilterClass struct {
	objc.Class
}

// An interface definition for the [compositingFilter] class.
type IcompositingFilter interface {
	ID() objc.ID
}

type compositingFilter struct {
	id objc.ID
}

func compositingFilterFrom(ptr unsafe.Pointer) compositingFilter {
	return compositingFilter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ compositingFilter) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _compositingFilterClass) Alloc() compositingFilter {
	rv := objc.Send[compositingFilter](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _compositingFilterClass) New() compositingFilter {
	rv := objc.Send[compositingFilter](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcompositingFilter creates and returns a new initialized instance.
func NewcompositingFilter() compositingFilter {
	return compositingFilterClass.New()
}

// Init initializes the instance.
func (c_ compositingFilter) Init() compositingFilter {
	rv := objc.Send[compositingFilter](c_.ID(), selInit)
	return rv
}
