
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [gState] class.
var gStateClass _gStateClass

func init() {
	gStateClass = _gStateClass{objc.GetClass("gState")}
}

type _gStateClass struct {
	objc.Class
}

// An interface definition for the [gState] class.
type IgState interface {
	ID() objc.ID
}

type gState struct {
	id objc.ID
}

func gStateFrom(ptr unsafe.Pointer) gState {
	return gState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ gState) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _gStateClass) Alloc() gState {
	rv := objc.Send[gState](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _gStateClass) New() gState {
	rv := objc.Send[gState](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewgState creates and returns a new initialized instance.
func NewgState() gState {
	return gStateClass.New()
}

// Init initializes the instance.
func (g_ gState) Init() gState {
	rv := objc.Send[gState](g_.ID(), selInit)
	return rv
}
