// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Storyboard] class.
var (
	storyboardClass     _StoryboardClass
	storyboardClassOnce sync.Once
)

func getStoryboardClass() _StoryboardClass {
	storyboardClassOnce.Do(func() {
		storyboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}
	})
	return storyboardClass
}

type _StoryboardClass struct {
	class objc.Class
}

// An interface definition for the [Storyboard] class.
type IStoryboard interface {
	objectivec.IObject
}

// An encapsulation of the design-time view controller and window controller graph represented in an Interface Builder storyboard resource file.
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




