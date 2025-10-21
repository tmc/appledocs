// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Storyboard] class.
var (
	StoryboardClass     _StoryboardClass
	StoryboardClassOnce sync.Once
)

func getStoryboardClass() _StoryboardClass {
	StoryboardClassOnce.Do(func() {
		StoryboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}
	})
	return StoryboardClass
}

type _StoryboardClass struct {
	class objc.Class
}

// An interface definition for the [Storyboard] class.
type IStoryboard interface {
	objectivec.IObject
	InstantiateControllerWithIdentifier(identifier IStoryboardSceneIdentifier) objc.ID
	InstantiateControllerWithIdentifierCreator(identifier IStoryboardSceneIdentifier, block unsafe.Pointer) objc.ID
	InstantiateInitialController() objc.ID
	InstantiateInitialControllerWithCreator(block unsafe.Pointer) objc.ID
}

// An encapsulation of the design-time view controller and window controller graph represented in an Interface Builder storyboard resource file.
//
// You can use storyboard files to define the view and window controllers for all or part of an app’s user interface. Typically, AppKit creates these objects automatically in response to actions defined within a storyboard file itself, such as the clicking of a button or the choosing of a menu item. However, you can use a storyboard object to directly instantiate the initial view controller from a storyboard file or to instantiate other view or window controllers that you want to present programmatically. In the context of a storyboard file, each contained controller is called a . A transition from one scene to another in a storyboard is called a . This same term, and the same Cocoa APIs, express a containment relationship between two scenes. In macOS, containment (rather than transition) is the more common notion for storyboards. For descriptions of the related APIs, refer to and .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard
type Storyboard struct {
	objectivec.Object
}

// StoryboardFrom constructs a [Storyboard] from an unsafe.Pointer.
//
// An encapsulation of the design-time view controller and window controller graph represented in an Interface Builder storyboard resource file.
func StoryboardFrom(ptr unsafe.Pointer) Storyboard {
	return Storyboard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StoryboardClass) Alloc() Storyboard {
	rv := objc.Send[Storyboard](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StoryboardClass) New() Storyboard {
	rv := objc.Send[Storyboard](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Storyboard) Init() Storyboard {
	rv := objc.Send[Storyboard](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Storyboard) Autorelease() Storyboard {
	rv := objc.Send[Storyboard](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStoryboard creates a new Storyboard instance.
func NewStoryboard() Storyboard {
	return getStoryboardClass().New()
}




// Creates a storyboard based on the named storyboard file in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/init(name:bundle:)
func NewStoryboardWithNameBundle(name IStoryboardName, storyboardBundleOrNil foundation.IBundle) Storyboard {
	rv := objc.Send[Storyboard](objc.ID(getStoryboardClass().class), objc.Sel("storyboardWithName:bundle:"), name, storyboardBundleOrNil)
	return rv
}


// Creates a storyboard based on the named storyboard file in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/init(name:bundle:)
func (sc _StoryboardClass) StoryboardWithNameBundle(name IStoryboardName, storyboardBundleOrNil foundation.IBundle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("storyboardWithName:bundle:"), name, storyboardBundleOrNil)
	return rv
}

// The app’s main storyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/main
func (sc _StoryboardClass) MainStoryboard() Storyboard {
	rv := objc.Send[NSStoryboard](objc.ID(sc.class), objc.Sel("mainStoryboard"))
	return rv
}
// Instantiates a specified view controller or window controller from a storyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateController(withIdentifier:)
func (s_ Storyboard) InstantiateControllerWithIdentifier(identifier IStoryboardSceneIdentifier) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("instantiateControllerWithIdentifier:"), identifier)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateControllerWithIdentifier:creator:
func (s_ Storyboard) InstantiateControllerWithIdentifierCreator(identifier IStoryboardSceneIdentifier, block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("instantiateControllerWithIdentifier:creator:"), identifier, block)
	return rv
}

// Creates the initial view controller or window controller from a storyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateInitialController()
func (s_ Storyboard) InstantiateInitialController() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("instantiateInitialController"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateInitialControllerWithCreator:
func (s_ Storyboard) InstantiateInitialControllerWithCreator(block unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("instantiateInitialControllerWithCreator:"), block)
	return rv
}

// The app’s main storyboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboard/main
func (s_ Storyboard) MainStoryboard() NSStoryboard {
	rv := objc.Send[NSStoryboard](s_.ID, objc.Sel("mainStoryboard"))
	return rv
}


