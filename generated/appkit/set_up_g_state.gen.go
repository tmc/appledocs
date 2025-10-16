
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [setUpGState] class.
var setUpGStateClass _setUpGStateClass

func init() {
	setUpGStateClass = _setUpGStateClass{objc.GetClass("setUpGState")}
}

type _setUpGStateClass struct {
	objc.Class
}

// An interface definition for the [setUpGState] class.
type IsetUpGState interface {
	ID() objc.ID
}

type setUpGState struct {
	id objc.ID
}

func setUpGStateFrom(ptr unsafe.Pointer) setUpGState {
	return setUpGState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ setUpGState) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _setUpGStateClass) Alloc() setUpGState {
	rv := objc.Send[setUpGState](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _setUpGStateClass) New() setUpGState {
	rv := objc.Send[setUpGState](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsetUpGState creates and returns a new initialized instance.
func NewsetUpGState() setUpGState {
	return setUpGStateClass.New()
}

// Init initializes the instance.
func (s_ setUpGState) Init() setUpGState {
	rv := objc.Send[setUpGState](s_.ID(), selInit)
	return rv
}
