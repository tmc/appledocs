// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Event] class.
var EventClass objc.Class

func init() {
	EventClass = objc.GetClass("NSEvent")
}

type Event struct {
	objc.ID
}

func EventFrom(ptr unsafe.Pointer) Event {
	return Event{
		ID: objc.ID(ptr),
	}
}




