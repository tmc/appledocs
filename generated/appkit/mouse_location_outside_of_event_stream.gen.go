
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [mouseLocationOutsideOfEventStream] class.
var mouseLocationOutsideOfEventStreamClass _mouseLocationOutsideOfEventStreamClass

func init() {
	mouseLocationOutsideOfEventStreamClass = _mouseLocationOutsideOfEventStreamClass{objc.GetClass("mouseLocationOutsideOfEventStream")}
}

type _mouseLocationOutsideOfEventStreamClass struct {
	objc.Class
}

// An interface definition for the [mouseLocationOutsideOfEventStream] class.
type ImouseLocationOutsideOfEventStream interface {
	ID() objc.ID
}

type mouseLocationOutsideOfEventStream struct {
	id objc.ID
}

func mouseLocationOutsideOfEventStreamFrom(ptr unsafe.Pointer) mouseLocationOutsideOfEventStream {
	return mouseLocationOutsideOfEventStream{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ mouseLocationOutsideOfEventStream) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _mouseLocationOutsideOfEventStreamClass) Alloc() mouseLocationOutsideOfEventStream {
	rv := objc.Send[mouseLocationOutsideOfEventStream](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _mouseLocationOutsideOfEventStreamClass) New() mouseLocationOutsideOfEventStream {
	rv := objc.Send[mouseLocationOutsideOfEventStream](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmouseLocationOutsideOfEventStream creates and returns a new initialized instance.
func NewmouseLocationOutsideOfEventStream() mouseLocationOutsideOfEventStream {
	return mouseLocationOutsideOfEventStreamClass.New()
}

// Init initializes the instance.
func (m_ mouseLocationOutsideOfEventStream) Init() mouseLocationOutsideOfEventStream {
	rv := objc.Send[mouseLocationOutsideOfEventStream](m_.ID(), selInit)
	return rv
}
