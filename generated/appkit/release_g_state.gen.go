
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [releaseGState] class.
var releaseGStateClass _releaseGStateClass

func init() {
	releaseGStateClass = _releaseGStateClass{objc.GetClass("releaseGState")}
}

type _releaseGStateClass struct {
	objc.Class
}

// An interface definition for the [releaseGState] class.
type IreleaseGState interface {
	ID() objc.ID
}

type releaseGState struct {
	id objc.ID
}

func releaseGStateFrom(ptr unsafe.Pointer) releaseGState {
	return releaseGState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ releaseGState) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _releaseGStateClass) Alloc() releaseGState {
	rv := objc.Send[releaseGState](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _releaseGStateClass) New() releaseGState {
	rv := objc.Send[releaseGState](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewreleaseGState creates and returns a new initialized instance.
func NewreleaseGState() releaseGState {
	return releaseGStateClass.New()
}

// Init initializes the instance.
func (r_ releaseGState) Init() releaseGState {
	rv := objc.Send[releaseGState](r_.ID(), selInit)
	return rv
}
