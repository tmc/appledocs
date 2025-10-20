// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingSession] class.
var (
	DraggingSessionClass     _DraggingSessionClass
	DraggingSessionClassOnce sync.Once
)

func getDraggingSessionClass() _DraggingSessionClass {
	DraggingSessionClassOnce.Do(func() {
		DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}
	})
	return DraggingSessionClass
}

type _DraggingSessionClass struct {
	class objc.Class
}

// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	objectivec.IObject
}

// The encapsulation of a drag-and-drop action that supports modification of the drag while in progress.
//
// You start a new dragging session by calling the method method. This method immediately returns and you can further modify the properties of the dragging session. The actual drag begins at the next turn of the run loop.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getDraggingSessionClass().New()
}
