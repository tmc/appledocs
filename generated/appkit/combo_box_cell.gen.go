// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboBoxCell] class.
var (
	comboBoxCellClass     _ComboBoxCellClass
	comboBoxCellClassOnce sync.Once
)

func getComboBoxCellClass() _ComboBoxCellClass {
	comboBoxCellClassOnce.Do(func() {
		comboBoxCellClass = _ComboBoxCellClass{objc.GetClass("NSComboBoxCell")}
	})
	return comboBoxCellClass
}

type _ComboBoxCellClass struct {
	class objc.Class
}

// An interface definition for the [ComboBoxCell] class.
type IComboBoxCell interface {
	ITextFieldCell
}

// The user interface of a combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboBoxCell

type ComboBoxCell struct {
	TextFieldCell
}

// ComboBoxCellFrom constructs a [ComboBoxCell] from an unsafe.Pointer.
//
// The user interface of a combo box.
func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (cc _ComboBoxCellClass) Alloc() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComboBoxCellClass) New() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComboBoxCell) Init() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComboBoxCell) Autorelease() ComboBoxCell {
	rv := objc.Send[ComboBoxCell](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComboBoxCell creates a new ComboBoxCell instance.
func NewComboBoxCell() ComboBoxCell {
	return getComboBoxCellClass().New()
}




