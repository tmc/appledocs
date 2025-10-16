
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resignFirstResponder] class.
var resignFirstResponderClass _resignFirstResponderClass

func init() {
	resignFirstResponderClass = _resignFirstResponderClass{objc.GetClass("resignFirstResponder")}
}

type _resignFirstResponderClass struct {
	objc.Class
}

// An interface definition for the [resignFirstResponder] class.
type IresignFirstResponder interface {
	ID() objc.ID
}

type resignFirstResponder struct {
	id objc.ID
}

func resignFirstResponderFrom(ptr unsafe.Pointer) resignFirstResponder {
	return resignFirstResponder{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resignFirstResponder) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resignFirstResponderClass) Alloc() resignFirstResponder {
	rv := objc.Send[resignFirstResponder](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resignFirstResponderClass) New() resignFirstResponder {
	rv := objc.Send[resignFirstResponder](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresignFirstResponder creates and returns a new initialized instance.
func NewresignFirstResponder() resignFirstResponder {
	return resignFirstResponderClass.New()
}

// Init initializes the instance.
func (r_ resignFirstResponder) Init() resignFirstResponder {
	rv := objc.Send[resignFirstResponder](r_.ID(), selInit)
	return rv
}
