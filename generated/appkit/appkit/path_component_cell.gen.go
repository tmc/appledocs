// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathComponentCell] class.
var (
	pathComponentCellClass     _PathComponentCellClass
	pathComponentCellClassOnce sync.Once
)

func getPathComponentCellClass() _PathComponentCellClass {
	pathComponentCellClassOnce.Do(func() {
		pathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}
	})
	return pathComponentCellClass
}

type _PathComponentCellClass struct {
	class objc.Class
}

// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ITextFieldCell
}

// A component of a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell

type PathComponentCell struct {
	TextFieldCell
}

// PathComponentCellFrom constructs a [PathComponentCell] from an unsafe.Pointer.
//
// A component of a path.
func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.Send[PathComponentCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PathComponentCell) Autorelease() PathComponentCell {
	rv := objc.Send[PathComponentCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPathComponentCell creates a new PathComponentCell instance.
func NewPathComponentCell() PathComponentCell {
	return getPathComponentCellClass().New()
}




