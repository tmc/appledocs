
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [occlusionState] class.
var occlusionStateClass _occlusionStateClass

func init() {
	occlusionStateClass = _occlusionStateClass{objc.GetClass("occlusionState")}
}

type _occlusionStateClass struct {
	objc.Class
}

// An interface definition for the [occlusionState] class.
type IocclusionState interface {
	ID() objc.ID
}

type occlusionState struct {
	id objc.ID
}

func occlusionStateFrom(ptr unsafe.Pointer) occlusionState {
	return occlusionState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ occlusionState) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _occlusionStateClass) Alloc() occlusionState {
	rv := objc.Send[occlusionState](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _occlusionStateClass) New() occlusionState {
	rv := objc.Send[occlusionState](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewocclusionState creates and returns a new initialized instance.
func NewocclusionState() occlusionState {
	return occlusionStateClass.New()
}

// Init initializes the instance.
func (o_ occlusionState) Init() occlusionState {
	rv := objc.Send[occlusionState](o_.ID(), selInit)
	return rv
}
