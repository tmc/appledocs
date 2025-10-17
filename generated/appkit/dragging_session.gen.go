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



