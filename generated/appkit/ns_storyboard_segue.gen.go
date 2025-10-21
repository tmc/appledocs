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
	StoryboardSegueClass     _StoryboardSegueClass
	StoryboardSegueClassOnce sync.Once
)

func getStoryboardSegueClass() _StoryboardSegueClass {
	StoryboardSegueClassOnce.Do(func() {
		StoryboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}
	})
	return StoryboardSegueClass
}

type _StoryboardSegueClass struct {
	class objc.Class
}

// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	objectivec.IObject
	Perform()
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

// The designated initializer for a storyboard segue.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:)
func NewStoryboardSegueWithIdentifierSourceDestination(identifier unsafe.Pointer, sourceController objc.ID, destinationController objc.ID) StoryboardSegue {
	instance := getStoryboardSegueClass().Alloc()
	rv := objc.Send[StoryboardSegue](instance.ID, objc.Sel("initWithIdentifier:source:destination:"), identifier, sourceController, destinationController)
	rv.Autorelease()
	return rv
}

// Performs a visual transition from one controller to another.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/perform()
func (s_ StoryboardSegue) Perform() {
	objc.Send[objc.ID](s_.ID, objc.Sel("perform"))
}

// An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/identifier-swift.property
func (s_ StoryboardSegue) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("identifier"))
	return rv
}

// The starting/containing view controller or window controller for the storyboard segue.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/sourceController
func (s_ StoryboardSegue) SourceController() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("sourceController"))
	return rv
}
