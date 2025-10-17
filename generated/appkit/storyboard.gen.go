
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Storyboard] class.
var StoryboardClass _StoryboardClass

func init() {
	StoryboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}
}

type _StoryboardClass struct {
	objc.Class
}

// An interface definition for the [Storyboard] class.
type IStoryboard interface {
	ID() objc.ID
}

type Storyboard struct {
	id objc.ID
}

func StoryboardFrom(ptr unsafe.Pointer) Storyboard {
	return Storyboard{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Storyboard) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StoryboardClass) Alloc() Storyboard {
	rv := objc.Send[Storyboard](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StoryboardClass) New() Storyboard {
	rv := objc.Send[Storyboard](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStoryboard creates and returns a new initialized instance.
func NewStoryboard() Storyboard {
	return StoryboardClass.New()
}

// Init initializes the instance.
func (s_ Storyboard) Init() Storyboard {
	rv := objc.Send[Storyboard](s_.ID(), selInit)
	return rv
}
