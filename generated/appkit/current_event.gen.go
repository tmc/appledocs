
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [currentEvent] class.
var currentEventClass _currentEventClass

func init() {
	currentEventClass = _currentEventClass{objc.GetClass("currentEvent")}
}

type _currentEventClass struct {
	objc.Class
}

// An interface definition for the [currentEvent] class.
type IcurrentEvent interface {
	ID() objc.ID
}

type currentEvent struct {
	id objc.ID
}

func currentEventFrom(ptr unsafe.Pointer) currentEvent {
	return currentEvent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ currentEvent) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _currentEventClass) Alloc() currentEvent {
	rv := objc.Send[currentEvent](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _currentEventClass) New() currentEvent {
	rv := objc.Send[currentEvent](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcurrentEvent creates and returns a new initialized instance.
func NewcurrentEvent() currentEvent {
	return currentEventClass.New()
}

// Init initializes the instance.
func (c_ currentEvent) Init() currentEvent {
	rv := objc.Send[currentEvent](c_.ID(), selInit)
	return rv
}
