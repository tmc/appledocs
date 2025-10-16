
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [refusesFirstResponder] class.
var refusesFirstResponderClass _refusesFirstResponderClass

func init() {
	refusesFirstResponderClass = _refusesFirstResponderClass{objc.GetClass("refusesFirstResponder")}
}

type _refusesFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [refusesFirstResponder] class.
type IrefusesFirstResponder interface {
	ID() objc.ID
}

type refusesFirstResponder struct {
	id objc.ID
}

func refusesFirstResponderFrom(ptr unsafe.Pointer) refusesFirstResponder {
	return refusesFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ refusesFirstResponder) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _refusesFirstResponderClass) Alloc() refusesFirstResponder {
	rv := objc.Send[refusesFirstResponder](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _refusesFirstResponderClass) New() refusesFirstResponder {
	rv := objc.Send[refusesFirstResponder](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrefusesFirstResponder creates and returns a new initialized instance.
func NewrefusesFirstResponder() refusesFirstResponder {
	return refusesFirstResponderClass.New()
}

// Init initializes the instance.
func (r_ refusesFirstResponder) Init() refusesFirstResponder {
	rv := objc.Send[refusesFirstResponder](r_.ID(), selInit)
	return rv
}
