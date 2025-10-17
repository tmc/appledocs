// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StoryboardSegue] class.
var StoryboardSegueClass objc.Class

func init() {
	StoryboardSegueClass = objc.GetClass("NSStoryboardSegue")
}

type StoryboardSegue struct {
	objc.ID
}

func StoryboardSegueFrom(ptr unsafe.Pointer) StoryboardSegue {
	return StoryboardSegue{
		ID: objc.ID(ptr),
	}
}



