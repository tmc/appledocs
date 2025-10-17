// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var levelIndicatorCellClass = _LevelIndicatorCellClass{objc.GetClass("NSLevelIndicatorCell")}

type _LevelIndicatorCellClass struct {
	class objc.Class
}

// is a subclass of that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicatorCell

type LevelIndicatorCell struct {
	ActionCell
}

// LevelIndicatorCellFrom constructs a [LevelIndicatorCell] from an unsafe.Pointer.
//
// is a subclass of that provides several level indicator display styles including: capacity, ranking and relevancy. The capacity style provides both continuous and discrete modes.
func LevelIndicatorCellFrom(ptr unsafe.Pointer) LevelIndicatorCell {
	return LevelIndicatorCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



