// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Control] class.
var (
	ControlClass     _ControlClass
	ControlClassOnce sync.Once
)

func getControlClass() _ControlClass {
	ControlClassOnce.Do(func() {
		ControlClass = _ControlClass{objc.GetClass("NSControl")}
	})
	return ControlClass
}

type _ControlClass struct {
	class objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	IView
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	Alignment() unsafe.Pointer
	SetAlignment(value unsafe.Pointer)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	BaseWritingDirection() unsafe.Pointer
	SetBaseWritingDirection(value unsafe.Pointer)
	ControlSize() unsafe.Pointer
	SetControlSize(value unsafe.Pointer)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	Font() IFont
	SetFont(value IFont)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.Formatter)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
	IntValue() unsafe.Pointer
	SetIntValue(value unsafe.Pointer)
	IntegerValue() int
	SetIntegerValue(value int)
	IsContinuous() bool
	SetIsContinuous(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	LineBreakMode() unsafe.Pointer
	SetLineBreakMode(value unsafe.Pointer)
	ObjectValue() unsafe.Pointer
	SetObjectValue(value unsafe.Pointer)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	StringValue() string
	SetStringValue(value string)
	Tag() int
	SetTag(value int)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
}

// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
//
// The class is abstract and must be subclassed to be used. Although you can subclass it yourself, more often you use one of the subclasses already defined by AppKit. A control draws content on the screen, automatically handles user interactions with that content, and calls the action method of its target object for any significant user interactions.


// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl
type Control struct {
	View
}

// ControlFrom constructs a [Control] from an unsafe.Pointer.
//
// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
func ControlFrom(ptr unsafe.Pointer) Control {
	return Control{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ControlClass) Alloc() Control {
	rv := objc.Send[Control](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ControlClass) New() Control {
	rv := objc.Send[Control](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Control) Init() Control {
	rv := objc.Send[Control](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Control) Autorelease() Control {
	rv := objc.Send[Control](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewControl creates a new Control instance.
func NewControl() Control {
	return getControlClass().New()
}



// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (c_ Control) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("action"))
	return rv
}


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (c_ Control) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}


// The alignment mode of the text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/alignment
func (c_ Control) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alignment"))
	return rv
}


// The alignment mode of the text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/alignment
func (c_ Control) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}


// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/allowsexpansiontooltips
func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsExpansionToolTips"))
	return rv
}


// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/allowsexpansiontooltips
func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsExpansionToolTips:"), value)
}


// The value of the receiver’s cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/attributedstringvalue
func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}


// The value of the receiver’s cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/attributedstringvalue
func (c_ Control) SetAttributedStringValue(value foundation.AttributedString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/basewritingdirection
func (c_ Control) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/basewritingdirection
func (c_ Control) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// The size of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/controlsize-swift.property
func (c_ Control) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlSize"))
	return rv
}


// The size of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/controlsize-swift.property
func (c_ Control) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (c_ Control) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("doubleValue"))
	return rv
}


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/doublevalue
func (c_ Control) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}


// The value of the receiver’s cell as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/floatvalue
func (c_ Control) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("floatValue"))
	return rv
}


// The value of the receiver’s cell as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/floatvalue
func (c_ Control) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}


// The font used to draw text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/font
func (c_ Control) Font() IFont {
	rv := objc.Send[Font](c_.ID, objc.Sel("font"))
	return rv
}


// The font used to draw text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/font
func (c_ Control) SetFont(value IFont) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}


// The receiver’s formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/formatter
func (c_ Control) Formatter() foundation.Formatter {
	rv := objc.Send[foundation.Formatter](c_.ID, objc.Sel("formatter"))
	return rv
}


// The receiver’s formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/formatter
func (c_ Control) SetFormatter(value foundation.Formatter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}


// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ignoresmulticlick
func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresMultiClick"))
	return rv
}


// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ignoresmulticlick
func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresMultiClick:"), value)
}


// The value of the receiver’s cell as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/intvalue
func (c_ Control) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("intValue"))
	return rv
}


// The value of the receiver’s cell as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/intvalue
func (c_ Control) SetIntValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}


// The value of the receiver’s cell as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/integervalue
func (c_ Control) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}


// The value of the receiver’s cell as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/integervalue
func (c_ Control) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (c_ Control) IsContinuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (c_ Control) SetIsContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuous:"), value)
}


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (c_ Control) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (c_ Control) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ishighlighted
func (c_ Control) IsHighlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ishighlighted
func (c_ Control) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighlighted:"), value)
}


// The line break mode to use for text in the control’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/linebreakmode
func (c_ Control) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The line break mode to use for text in the control’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/linebreakmode
func (c_ Control) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/objectvalue
func (c_ Control) ObjectValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectValue"))
	return rv
}


// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/objectvalue
func (c_ Control) SetObjectValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}


// A Boolean value indicating whether the receiver refuses the first responder role.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/refusesfirstresponder
func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}


// A Boolean value indicating whether the receiver refuses the first responder role.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/refusesfirstresponder
func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}


// The value of the receiver’s cell as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/stringvalue
func (c_ Control) StringValue() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stringValue"))
	return rv
}


// The value of the receiver’s cell as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/stringvalue
func (c_ Control) SetStringValue(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), objc.String(value))
}


// The tag identifying the receiver (not the tag of the receiver’s cell).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/tag
func (c_ Control) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}


// The tag identifying the receiver (not the tag of the receiver’s cell).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/tag
func (c_ Control) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/target
func (c_ Control) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("target"))
	return rv
}


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/target
func (c_ Control) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}


// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/usessinglelinemode
func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}


// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/usessinglelinemode
func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}



