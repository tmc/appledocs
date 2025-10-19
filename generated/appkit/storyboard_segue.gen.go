// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StoryboardSegue] class.
var storyboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}

type _StoryboardSegueClass struct {
	class objc.Class
}

// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	objectivec.IObject
}

// A transition or containment relationship between two scenes in a storyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue

type StoryboardSegue struct {
	objectivec.Object
}

// StoryboardSegueFrom constructs a [StoryboardSegue] from an unsafe.Pointer.
//
// A transition or containment relationship between two scenes in a storyboard.
func StoryboardSegueFrom(ptr unsafe.Pointer) StoryboardSegue {
	return StoryboardSegue{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _StoryboardSegueClass) Alloc() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _StoryboardSegueClass) New() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StoryboardSegue) Init() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StoryboardSegue) Autorelease() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStoryboardSegue creates a new StoryboardSegue instance.
func NewStoryboardSegue() StoryboardSegue {
	return storyboardSegueClass.New()
}




