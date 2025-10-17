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



