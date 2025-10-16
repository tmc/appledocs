
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [frameAutosaveName] class.
var frameAutosaveNameClass _frameAutosaveNameClass

func init() {
	frameAutosaveNameClass = _frameAutosaveNameClass{objc.GetClass("frameAutosaveName")}
}

type _frameAutosaveNameClass struct {
	objc.Class
}

// An interface definition for the [frameAutosaveName] class.
type IframeAutosaveName interface {
	ID() objc.ID
}

type frameAutosaveName struct {
	id objc.ID
}

func frameAutosaveNameFrom(ptr unsafe.Pointer) frameAutosaveName {
	return frameAutosaveName{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ frameAutosaveName) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _frameAutosaveNameClass) Alloc() frameAutosaveName {
	rv := objc.Send[frameAutosaveName](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _frameAutosaveNameClass) New() frameAutosaveName {
	rv := objc.Send[frameAutosaveName](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewframeAutosaveName creates and returns a new initialized instance.
func NewframeAutosaveName() frameAutosaveName {
	return frameAutosaveNameClass.New()
}

// Init initializes the instance.
func (f_ frameAutosaveName) Init() frameAutosaveName {
	rv := objc.Send[frameAutosaveName](f_.ID(), selInit)
	return rv
}
