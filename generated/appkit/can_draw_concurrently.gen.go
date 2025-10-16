
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canDrawConcurrently] class.
var canDrawConcurrentlyClass _canDrawConcurrentlyClass

func init() {
	canDrawConcurrentlyClass = _canDrawConcurrentlyClass{objc.GetClass("canDrawConcurrently")}
}

type _canDrawConcurrentlyClass struct {
	objc.Class
}

// An interface definition for the [canDrawConcurrently] class.
type IcanDrawConcurrently interface {
	ID() objc.ID
}

type canDrawConcurrently struct {
	id objc.ID
}

func canDrawConcurrentlyFrom(ptr unsafe.Pointer) canDrawConcurrently {
	return canDrawConcurrently{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canDrawConcurrently) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canDrawConcurrentlyClass) Alloc() canDrawConcurrently {
	rv := objc.Send[canDrawConcurrently](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canDrawConcurrentlyClass) New() canDrawConcurrently {
	rv := objc.Send[canDrawConcurrently](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanDrawConcurrently creates and returns a new initialized instance.
func NewcanDrawConcurrently() canDrawConcurrently {
	return canDrawConcurrentlyClass.New()
}

// Init initializes the instance.
func (c_ canDrawConcurrently) Init() canDrawConcurrently {
	rv := objc.Send[canDrawConcurrently](c_.ID(), selInit)
	return rv
}
