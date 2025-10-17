// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingSession] class.
var DraggingSessionClass objc.Class

func init() {
	DraggingSessionClass = objc.GetClass("NSDraggingSession")
}

type DraggingSession struct {
	objc.ID
}

func DraggingSessionFrom(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		ID: objc.ID(ptr),
	}
}



