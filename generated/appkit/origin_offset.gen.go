
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [originOffset] class.
var originOffsetClass _originOffsetClass

func init() {
	originOffsetClass = _originOffsetClass{objc.GetClass("originOffset")}
}

type _originOffsetClass struct {
	objc.Class
}

// An interface definition for the [originOffset] class.
type IoriginOffset interface {
	ID() objc.ID
}

type originOffset struct {
	id objc.ID
}

func originOffsetFrom(ptr unsafe.Pointer) originOffset {
	return originOffset{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ originOffset) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _originOffsetClass) Alloc() originOffset {
	rv := objc.Send[originOffset](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _originOffsetClass) New() originOffset {
	rv := objc.Send[originOffset](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeworiginOffset creates and returns a new initialized instance.
func NeworiginOffset() originOffset {
	return originOffsetClass.New()
}

// Init initializes the instance.
func (o_ originOffset) Init() originOffset {
	rv := objc.Send[originOffset](o_.ID(), selInit)
	return rv
}
