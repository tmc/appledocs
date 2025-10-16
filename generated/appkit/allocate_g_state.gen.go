
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allocateGState] class.
var allocateGStateClass _allocateGStateClass

func init() {
	allocateGStateClass = _allocateGStateClass{objc.GetClass("allocateGState")}
}

type _allocateGStateClass struct {
	objc.Class
}

// An interface definition for the [allocateGState] class.
type IallocateGState interface {
	ID() objc.ID
}

type allocateGState struct {
	id objc.ID
}

func allocateGStateFrom(ptr unsafe.Pointer) allocateGState {
	return allocateGState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allocateGState) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allocateGStateClass) Alloc() allocateGState {
	rv := objc.Send[allocateGState](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allocateGStateClass) New() allocateGState {
	rv := objc.Send[allocateGState](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallocateGState creates and returns a new initialized instance.
func NewallocateGState() allocateGState {
	return allocateGStateClass.New()
}

// Init initializes the instance.
func (a_ allocateGState) Init() allocateGState {
	rv := objc.Send[allocateGState](a_.ID(), selInit)
	return rv
}
