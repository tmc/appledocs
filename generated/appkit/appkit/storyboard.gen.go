// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Storyboard] class.
var StoryboardClass objc.Class

func init() {
	StoryboardClass = objc.GetClass("NSStoryboard")
}

type Storyboard struct {
	objc.ID
}

func StoryboardFrom(ptr unsafe.Pointer) Storyboard {
	return Storyboard{
		ID: objc.ID(ptr),
	}
}



