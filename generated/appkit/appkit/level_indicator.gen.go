// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LevelIndicator] class.
var LevelIndicatorClass objc.Class

func init() {
	LevelIndicatorClass = objc.GetClass("NSLevelIndicator")
}

type LevelIndicator struct {
	objc.ID
}

func LevelIndicatorFrom(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		ID: objc.ID(ptr),
	}
}



