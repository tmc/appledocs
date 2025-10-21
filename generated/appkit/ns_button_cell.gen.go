// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonCell] class.
var (
	ButtonCellClass     _ButtonCellClass
	ButtonCellClassOnce sync.Once
)

func getButtonCellClass() _ButtonCellClass {
	ButtonCellClassOnce.Do(func() {
		ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}
	})
	return ButtonCellClass
}

type _ButtonCellClass struct {
	class objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
	SetKeyEquivalentFontSize(fontName string, fontSize float64)
}

// An object that defines the user interface of a button or other clickable region of a view.
//
// Setting the integer, float, double, or object value of an object results in a call to with the value converted to integer. In the case of , is equivalent to , and a non- object that doesn’t respond to sets the state to . Otherwise, the state is set to the object’s . Similarly, for most button types, querying the integer, float, double, or object value of an returns the current state in the requested representation. In the case of , this is an containing for on, for off, and integer value for the mixed state. For accelerator buttons (type or ) on systems that support pressure sensitivity, querying returns the amount of pressure applied while pressing the button. The configuration of an object controls how the button object appears and behaves, but it’s that sends a message when the control is clicked. For more information on the behavior of , see the and class specifications, and .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell
type ButtonCell struct {
	ActionCell
}

// ButtonCellFrom constructs a [ButtonCell] from an unsafe.Pointer.
//
// An object that defines the user interface of a button or other clickable region of a view.
func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ButtonCell) Autorelease() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButtonCell creates a new ButtonCell instance.
func NewButtonCell() ButtonCell {
	return getButtonCellClass().New()
}


// Sets by name and size of the font used to draw the key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/setKeyEquivalentFont(_:size:)
func (b_ ButtonCell) SetKeyEquivalentFontSize(fontName string, fontSize float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:size:"), objc.String(fontName), fontSize)
}

// The gradient of the button’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) GradientType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("gradientType"))
	return rv
}


// SetGradientType sets the value of the gradientType property.
// The gradient of the button’s border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/gradientType
func (b_ ButtonCell) SetGradientType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setGradientType:"), value)
}

// The button’s key-equivalent character.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) KeyEquivalent() string {
	rv := objc.Send[string](b_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// SetKeyEquivalent sets the value of the keyEquivalent property.
// The button’s key-equivalent character.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalent
func (b_ ButtonCell) SetKeyEquivalent(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The font used to draw the button’s key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) KeyEquivalentFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("keyEquivalentFont"))
	return rv
}


// SetKeyEquivalentFont sets the value of the keyEquivalentFont property.
// The font used to draw the button’s key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentFont
func (b_ ButtonCell) SetKeyEquivalentFont(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentFont:"), value)
}

// The mask that identifies the modifier keys for the button’s key equivalent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](b_.ID, objc.Sel("keyEquivalentModifierMask"))
	return rv
}


// SetKeyEquivalentModifierMask sets the value of the keyEquivalentModifierMask property.
// The mask that identifies the modifier keys for the button’s key equivalent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell/keyEquivalentModifierMask
func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}



