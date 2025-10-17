// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Storyboard] class.
var storyboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}

type _StoryboardClass struct {
	class objc.Class
}

// An interface definition for the [Storyboard] class.
type IStoryboard interface {
	objectivec.IObject
}

// An encapsulation of the design-time view controller and window controller graph represented in an Interface Builder storyboard resource file. [Full Topic]
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



