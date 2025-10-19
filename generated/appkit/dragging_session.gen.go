// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingSession] class.
var draggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}

type _DraggingSessionClass struct {
	class objc.Class
}

// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	objectivec.IObject
}

// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingSession

type DraggingSession struct {
	objectivec.Object
}

// DraggingSessionFrom constructs a [DraggingSession] from an unsafe.Pointer.
//
// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress.
func DraggingSessionFrom(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.Send[DraggingSession](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingSession) Autorelease() DraggingSession {
	rv := objc.Send[DraggingSession](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingSession creates a new DraggingSession instance.
func NewDraggingSession() DraggingSession {
	return draggingSessionClass.New()
}




