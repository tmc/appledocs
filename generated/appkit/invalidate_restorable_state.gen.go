
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [invalidateRestorableState] class.
var invalidateRestorableStateClass _invalidateRestorableStateClass

func init() {
	invalidateRestorableStateClass = _invalidateRestorableStateClass{objc.GetClass("invalidateRestorableState")}
}

type _invalidateRestorableStateClass struct {
	objc.Class
}

// An interface definition for the [invalidateRestorableState] class.
type IinvalidateRestorableState interface {
	ID() objc.ID
}

type invalidateRestorableState struct {
	id objc.ID
}

func invalidateRestorableStateFrom(ptr unsafe.Pointer) invalidateRestorableState {
	return invalidateRestorableState{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ invalidateRestorableState) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateRestorableStateClass) Alloc() invalidateRestorableState {
	rv := objc.Send[invalidateRestorableState](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _invalidateRestorableStateClass) New() invalidateRestorableState {
	rv := objc.Send[invalidateRestorableState](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinvalidateRestorableState creates and returns a new initialized instance.
func NewinvalidateRestorableState() invalidateRestorableState {
	return invalidateRestorableStateClass.New()
}

// Init initializes the instance.
func (i_ invalidateRestorableState) Init() invalidateRestorableState {
	rv := objc.Send[invalidateRestorableState](i_.ID(), selInit)
	return rv
}
