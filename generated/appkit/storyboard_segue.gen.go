
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StoryboardSegue] class.
var StoryboardSegueClass _StoryboardSegueClass

func init() {
	StoryboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}
}

type _StoryboardSegueClass struct {
	objc.Class
}

// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	ID() objc.ID
}

type StoryboardSegue struct {
	id objc.ID
}

func StoryboardSegueFrom(ptr unsafe.Pointer) StoryboardSegue {
	return StoryboardSegue{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StoryboardSegue) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StoryboardSegueClass) Alloc() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StoryboardSegueClass) New() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStoryboardSegue creates and returns a new initialized instance.
func NewStoryboardSegue() StoryboardSegue {
	return StoryboardSegueClass.New()
}

// Init initializes the instance.
func (s_ StoryboardSegue) Init() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](s_.ID(), selInit)
	return rv
}
