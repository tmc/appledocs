// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StoryboardSegue] class.
var (
	storyboardSegueClass     _StoryboardSegueClass
	storyboardSegueClassOnce sync.Once
)

func getStoryboardSegueClass() _StoryboardSegueClass {
	storyboardSegueClassOnce.Do(func() {
		storyboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}
	})
	return storyboardSegueClass
}

type _StoryboardSegueClass struct {
	class objc.Class
}

// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	objectivec.IObject
}

// A transition or containment relationship between two scenes in a storyboard.
//
// In this context, a is a view controller or a window controller and a is an instance of the class. A storyboard segue has a procedural notion of being invoked, known in the API as being . You can take advantage of hooks into the segue performance process by way of the protocol. You do not create storyboard segue objects directly. Instead, the system creates them as needed as segues are invoked. To run code during initialization and performance of a segue, override the and methods. You can initiate a segue programmatically with the method of the protocol. For example, you might do this to transition from a scene in one storyboard file to a scene in another.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getStoryboardSegueClass().New()
}




