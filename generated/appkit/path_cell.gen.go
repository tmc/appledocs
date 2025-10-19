// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathCell] class.
var (
	pathCellClass     _PathCellClass
	pathCellClassOnce sync.Once
)

func getPathCellClass() _PathCellClass {
	pathCellClassOnce.Do(func() {
		pathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}
	})
	return pathCellClass
}

type _PathCellClass struct {
	class objc.Class
}

// An interface definition for the [PathCell] class.
type IPathCell interface {
	IActionCell
}

// The user interface of a path control object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell
type PathCell struct {
	ActionCell
}

// PathCellFrom constructs a [PathCell] from an unsafe.Pointer.
//
// The user interface of a path control object.
func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PathCellClass) New() PathCell {
	rv := objc.Send[PathCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathCell) Init() PathCell {
	rv := objc.Send[PathCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathCell) Autorelease() PathCell {
	rv := objc.Send[PathCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathCell creates a new PathCell instance.
func NewPathCell() PathCell {
	return getPathCellClass().New()
}




