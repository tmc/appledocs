
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [renewGState] class.
var renewGStateClass _renewGStateClass

func init() {
	renewGStateClass = _renewGStateClass{objc.GetClass("renewGState")}
}

type _renewGStateClass struct {
	objc.Class
}

// An interface definition for the [renewGState] class.
type IrenewGState interface {
	ID() objc.ID
}

type renewGState struct {
	id objc.ID
}

func renewGStateFrom(ptr unsafe.Pointer) renewGState {
	return renewGState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ renewGState) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _renewGStateClass) Alloc() renewGState {
	rv := objc.Send[renewGState](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _renewGStateClass) New() renewGState {
	rv := objc.Send[renewGState](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrenewGState creates and returns a new initialized instance.
func NewrenewGState() renewGState {
	return renewGStateClass.New()
}

// Init initializes the instance.
func (r_ renewGState) Init() renewGState {
	rv := objc.Send[renewGState](r_.ID(), selInit)
	return rv
}
