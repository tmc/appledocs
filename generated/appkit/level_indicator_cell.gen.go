// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var LevelIndicatorCellClass objc.Class

func init() {
	LevelIndicatorCellClass = objc.GetClass("NSLevelIndicatorCell")
}

type LevelIndicatorCell struct {
	objc.ID
}

func LevelIndicatorCellFrom(ptr unsafe.Pointer) LevelIndicatorCell {
	return LevelIndicatorCell{
		ID: objc.ID(ptr),
	}
}



