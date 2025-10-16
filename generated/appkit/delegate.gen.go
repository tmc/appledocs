
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [delegate] class.
var delegateClass _delegateClass

func init() {
	delegateClass = _delegateClass{objc.GetClass("delegate")}
}

type _delegateClass struct {
	objc.Class
}

// An interface definition for the [delegate] class.
type Idelegate interface {
	ID() objc.ID
}

type delegate struct {
	id objc.ID
}

func delegateFrom(ptr unsafe.Pointer) delegate {
	return delegate{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ delegate) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _delegateClass) Alloc() delegate {
	rv := objc.Send[delegate](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _delegateClass) New() delegate {
	rv := objc.Send[delegate](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newdelegate creates and returns a new initialized instance.
func Newdelegate() delegate {
	return delegateClass.New()
}

// Init initializes the instance.
func (d_ delegate) Init() delegate {
	rv := objc.Send[delegate](d_.ID(), selInit)
	return rv
}
