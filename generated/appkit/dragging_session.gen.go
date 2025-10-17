
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingSession] class.
var DraggingSessionClass _DraggingSessionClass

func init() {
	DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}
}

type _DraggingSessionClass struct {
	objc.Class
}

// An interface definition for the [DraggingSession] class.
type IDraggingSession interface {
	ID() objc.ID
}

type DraggingSession struct {
	id objc.ID
}

func DraggingSessionFrom(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DraggingSession) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.Send[DraggingSession](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDraggingSession creates and returns a new initialized instance.
func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

// Init initializes the instance.
func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.Send[DraggingSession](d_.ID(), selInit)
	return rv
}
