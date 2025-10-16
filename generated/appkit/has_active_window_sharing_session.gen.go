
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasActiveWindowSharingSession] class.
var hasActiveWindowSharingSessionClass _hasActiveWindowSharingSessionClass

func init() {
	hasActiveWindowSharingSessionClass = _hasActiveWindowSharingSessionClass{objc.GetClass("hasActiveWindowSharingSession")}
}

type _hasActiveWindowSharingSessionClass struct {
	objc.Class
}

// An interface definition for the [hasActiveWindowSharingSession] class.
type IhasActiveWindowSharingSession interface {
	ID() objc.ID
}

type hasActiveWindowSharingSession struct {
	id objc.ID
}

func hasActiveWindowSharingSessionFrom(ptr unsafe.Pointer) hasActiveWindowSharingSession {
	return hasActiveWindowSharingSession{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasActiveWindowSharingSession) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasActiveWindowSharingSessionClass) Alloc() hasActiveWindowSharingSession {
	rv := objc.Send[hasActiveWindowSharingSession](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasActiveWindowSharingSessionClass) New() hasActiveWindowSharingSession {
	rv := objc.Send[hasActiveWindowSharingSession](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasActiveWindowSharingSession creates and returns a new initialized instance.
func NewhasActiveWindowSharingSession() hasActiveWindowSharingSession {
	return hasActiveWindowSharingSessionClass.New()
}

// Init initializes the instance.
func (h_ hasActiveWindowSharingSession) Init() hasActiveWindowSharingSession {
	rv := objc.Send[hasActiveWindowSharingSession](h_.ID(), selInit)
	return rv
}
