// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LevelIndicatorCell] class.
var (
	levelIndicatorCellClass     _LevelIndicatorCellClass
	levelIndicatorCellClassOnce sync.Once
)

func getLevelIndicatorCellClass() _LevelIndicatorCellClass {
	levelIndicatorCellClassOnce.Do(func() {
		levelIndicatorCellClass = _LevelIndicatorCellClass{objc.GetClass("NSLevelIndicatorCell")}
	})
	return levelIndicatorCellClass
}

type _LevelIndicatorCellClass struct {
	class objc.Class
}

// An interface definition for the [LevelIndicatorCell] class.
type ILevelIndicatorCell interface {
	IActionCell
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
// Alloc allocates a new instance without initialization.
func (lc _LevelIndicatorCellClass) Alloc() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LevelIndicatorCellClass) New() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LevelIndicatorCell) Init() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LevelIndicatorCell) Autorelease() LevelIndicatorCell {
	rv := objc.Send[LevelIndicatorCell](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLevelIndicatorCell creates a new LevelIndicatorCell instance.
func NewLevelIndicatorCell() LevelIndicatorCell {
	return getLevelIndicatorCellClass().New()
}




