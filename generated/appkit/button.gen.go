
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Button] class.
var ButtonClass _ButtonClass

func init() {
	ButtonClass = _ButtonClass{objc.GetClass("NSButton")}
}

type _ButtonClass struct {
	objc.Class
}

// An interface definition for the [Button] class.
type IButton interface {
	ID() objc.ID
	CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer)
	SetPeriodicDelayInterval(delay float32, interval float32)
}

type Button struct {
	id objc.ID
}

func ButtonFrom(ptr unsafe.Pointer) Button {
	return Button{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ Button) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonClass) Alloc() Button {
	rv := objc.Send[Button](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _ButtonClass) New() Button {
	rv := objc.Send[Button](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewButton creates and returns a new initialized instance.
func NewButton() Button {
	return ButtonClass.New()
}

// Init initializes the instance.
func (b_ Button) Init() Button {
	rv := objc.Send[Button](b_.ID(), selInit)
	return rv
}
// Creates a standard checkbox with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func (bc _ButtonClass) CheckboxWithTitleTargetAction(title unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.Class), objc.RegisterName("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}
// Creates a standard push button with a title and image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:image:target:action:)
func (bc _ButtonClass) ButtonWithTitleImageTargetAction(title unsafe.Pointer, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.Class), objc.RegisterName("buttonWithTitle:image:target:action:"), title, image, target, action)
	return rv
}

// Button_ButtonWithTitleImageTargetAction creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:image:target:action:)
func Button_ButtonWithTitleImageTargetAction(title unsafe.Pointer, image unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	return ButtonClass.ButtonWithTitleImageTargetAction(title, image, target, action)
}
// Creates a standard push button with the title you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:target:action:)
func (bc _ButtonClass) ButtonWithTitleTargetAction(title unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.Class), objc.RegisterName("buttonWithTitle:target:action:"), title, target, action)
	return rv
}

// Button_ButtonWithTitleTargetAction creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/init(title:target:action:)
func Button_ButtonWithTitleTargetAction(title unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	return ButtonClass.ButtonWithTitleTargetAction(title, target, action)
}
// Sets the priority compression options for this button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/compress(withPrioritizedCompressionOptions:)
func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}
// Sets the message delay and interval periods for a continuous button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/setPeriodicDelay(_:interval:)
func (b_ Button) SetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setPeriodicDelay:interval:"), delay, interval)
}
// The compression options active for this button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/activeCompressionOptions
func (b_ Button) ActiveCompressionOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("activeCompressionOptions"))
	return rv
}
// The title that the button displays in an off state, as an attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) AttributedTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("attributedTitle"))
	return rv
}
// SetAttributedTitle sets the value of the attributedTitle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/attributedTitle
func (b_ Button) SetAttributedTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setAttributedTitle:"), value)
}
// The color of the button’s bezel, in appearances that support it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/bezelColor
func (b_ Button) BezelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("bezelColor"))
	return rv
}
// SetBezelColor sets the value of the bezelColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/bezelColor
func (b_ Button) SetBezelColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBezelColor:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/borderShape
func (b_ Button) BorderShape() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("borderShape"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/borderShape
func (b_ Button) SetBorderShape(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBorderShape:"), value)
}
// A tint color to use for the template image and text content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) ContentTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("contentTintColor"))
	return rv
}
// SetContentTintColor sets the value of the contentTintColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/contentTintColor
func (b_ Button) SetContentTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setContentTintColor:"), value)
}
// A Boolean value that determines whether the button has a border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/isBordered
func (b_ Button) Bordered() bool {
	rv := objc.Send[bool](b_.ID(), objc.RegisterName("bordered"))
	return rv
}
// SetBordered sets the value of the bordered property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/isBordered
func (b_ Button) SetBordered(value bool) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBordered:"), value)
}
// A Boolean value that indicates whether the button is transparent. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/isTransparent
func (b_ Button) Transparent() bool {
	rv := objc.Send[bool](b_.ID(), objc.RegisterName("transparent"))
	return rv
}
// SetTransparent sets the value of the transparent property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/isTransparent
func (b_ Button) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTransparent:"), value)
}
// The sound that plays when the user clicks the button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/sound
func (b_ Button) Sound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("sound"))
	return rv
}
// SetSound sets the value of the sound property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSButton/sound
func (b_ Button) SetSound(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setSound:"), value)
}
