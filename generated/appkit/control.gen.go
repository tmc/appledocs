// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Control] class.
var (
	controlClass     _ControlClass
	controlClassOnce sync.Once
)

func getControlClass() _ControlClass {
	controlClassOnce.Do(func() {
		controlClass = _ControlClass{objc.GetClass("NSControl")}
	})
	return controlClass
}

type _ControlClass struct {
	class objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	IView
	AbortEditing() bool
	CurrentEditor() unsafe.Pointer
	DrawWithExpansionFrameInView(contentFrame coregraphics.CGRect, view unsafe.Pointer)
	EditWithFrameEditorDelegateEvent(rect coregraphics.CGRect, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	ExpansionFrameWithFrame(contentFrame coregraphics.CGRect) coregraphics.CGRect
	InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer)
	PerformClick(sender objc.ID)
	SelectWithFrameEditorDelegateStartLength(rect coregraphics.CGRect, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int)
	SendActionTo(action objc.SEL, target objc.ID) bool
	SendActionOn(mask unsafe.Pointer) int
	SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint)
	SizeThatFits(size coregraphics.CGSize) coregraphics.CGSize
	SizeToFit()
	TakeDoubleValueFrom(sender objc.ID)
	TakeFloatValueFrom(sender objc.ID)
	TakeIntValueFrom(sender objc.ID)
	TakeIntegerValueFrom(sender objc.ID)
	TakeObjectValueFrom(sender objc.ID)
	TakeStringValueFrom(sender objc.ID)
	ValidateEditing()
}

// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
//
// The class is abstract and must be subclassed to be used. Although you can subclass it yourself, more often you use one of the subclasses already defined by AppKit. A control draws content on the screen, automatically handles user interactions with that content, and calls the action method of its target object for any significant user interactions.
//
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


// Initializes a control with data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewControlWithCoder(coder unsafe.Pointer) Control {
	instance := getControlClass().Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Initializes a control with the specified frame rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewControlWithFrame(frameRect coregraphics.CGRect) Control {
	instance := getControlClass().Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Terminates the current editing operation and discards any edited text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/abortEditing()
func (c_ Control) AbortEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("abortEditing"))
	return rv
}

// Returns the current field editor for the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/currentEditor()
func (c_ Control) CurrentEditor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentEditor"))
	return rv
}

// Performs custom expansion tool tip drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c_ Control) DrawWithExpansionFrameInView(contentFrame coregraphics.CGRect, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), contentFrame, view)
}

// Begins editing of the receiver’s text using the specified field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c_ Control) EditWithFrameEditorDelegateEvent(rect coregraphics.CGRect, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:editor:delegate:event:"), rect, textObj, delegate, event)
}

// Ends the editing of text in the receiver using the specified field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/endEditing(_:)
func (c_ Control) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}

// The frame in which a tool tip can be displayed, if needed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c_ Control) ExpansionFrameWithFrame(contentFrame coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("expansionFrameWithFrame:"), contentFrame)
	return rv
}

// Notifies the control that the intrinsic content size for its cell is no longer valid.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateIntrinsicContentSizeForCell:"), cell)
}

// Simulates a single mouse click on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/performClick(_:)
func (c_ Control) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}

// Selects the specified text range in the receiver’s field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect coregraphics.CGRect, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:editor:delegate:start:length:"), rect, textObj, delegate, selStart, selLength)
}

// Causes the specified action to be sent to the target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(_:to:)
func (c_ Control) SendActionTo(action objc.SEL, target objc.ID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendAction:to:"), action, target)
	return rv
}

// Sets the conditions on which the receiver sends action messages to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(on:)
func (c_ Control) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}

// Sets the auto-ranging and floating point number format of the receiver’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/setFloatingPointFormat:left:right:
func (c_ Control) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}

// Asks the control to calculate and return the size that best fits the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c_ Control) SizeThatFits(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("sizeThatFits:"), size)
	return rv
}

// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeToFit()
func (c_ Control) SizeToFit() {
	objc.Send[objc.ID](c_.ID, objc.Sel("sizeToFit"))
}

// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c_ Control) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c_ Control) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c_ Control) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c_ Control) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c_ Control) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c_ Control) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}

// Validates changes to any user-typed text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/validateEditing()
func (c_ Control) ValidateEditing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("validateEditing"))
}

// The default action-message selector associated with the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/action
func (c_ Control) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The default action-message selector associated with the control.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/action
func (c_ Control) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}
// The alignment mode of the text in the receiver’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/alignment
func (c_ Control) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The alignment mode of the text in the receiver’s cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/alignment
func (c_ Control) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}
// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsExpansionToolTips"))
	return rv
}


// SetAllowsExpansionToolTips sets the value of the allowsExpansionToolTips property.
// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsExpansionToolTips:"), value)
}
// The value of the receiver’s cell as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) AttributedStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}


// SetAttributedStringValue sets the value of the attributedStringValue property.
// The value of the receiver’s cell as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) SetAttributedStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}
// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// SetBaseWritingDirection sets the value of the baseWritingDirection property.
// The initial writing direction used to determine the actual writing direction for text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}
// The size of the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlSize"))
	return rv
}


// SetControlSize sets the value of the controlSize property.
// The size of the control.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}
// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/doubleValue
func (c_ Control) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The value of the receiver’s cell as a double-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/doubleValue
func (c_ Control) SetDoubleValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}
// The value of the receiver’s cell as a single-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/floatValue
func (c_ Control) FloatValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("floatValue"))
	return rv
}


// SetFloatValue sets the value of the floatValue property.
// The value of the receiver’s cell as a single-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/floatValue
func (c_ Control) SetFloatValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}
// The font used to draw text in the receiver’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/font
func (c_ Control) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font used to draw text in the receiver’s cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/font
func (c_ Control) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}
// The receiver’s formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/formatter
func (c_ Control) Formatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatter"))
	return rv
}


// SetFormatter sets the value of the formatter property.
// The receiver’s formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/formatter
func (c_ Control) SetFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}
// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresMultiClick"))
	return rv
}


// SetIgnoresMultiClick sets the value of the ignoresMultiClick property.
// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresMultiClick:"), value)
}
// The value of the receiver’s cell as an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/intValue
func (c_ Control) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("intValue"))
	return rv
}


// SetIntValue sets the value of the intValue property.
// The value of the receiver’s cell as an integer.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/intValue
func (c_ Control) SetIntValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}
// The value of the receiver’s cell as an value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/integerValue
func (c_ Control) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}


// SetIntegerValue sets the value of the integerValue property.
// The value of the receiver’s cell as an value.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/integerValue
func (c_ Control) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}
// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isContinuous
func (c_ Control) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}


// SetContinuous sets the value of the continuous property.
// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isContinuous
func (c_ Control) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}
// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isEnabled
func (c_ Control) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether the receiver reacts to mouse events.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isEnabled
func (c_ Control) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}
// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) Highlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highlighted"))
	return rv
}


// SetHighlighted sets the value of the highlighted property.
// A Boolean value that indicates whether the cell is highlighted.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlighted:"), value)
}
// The line break mode to use for text in the control’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// SetLineBreakMode sets the value of the lineBreakMode property.
// The line break mode to use for text in the control’s cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}
// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/objectValue
func (c_ Control) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValue"))
	return rv
}


// SetObjectValue sets the value of the objectValue property.
// The value of the receiver’s cell as an Objective-C object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/objectValue
func (c_ Control) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}
// A Boolean value indicating whether the receiver refuses the first responder role.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}


// SetRefusesFirstResponder sets the value of the refusesFirstResponder property.
// A Boolean value indicating whether the receiver refuses the first responder role.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}
// The value of the receiver’s cell as an object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/stringValue
func (c_ Control) StringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
// The value of the receiver’s cell as an object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/stringValue
func (c_ Control) SetStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), value)
}
// The tag identifying the receiver (not the tag of the receiver’s cell).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/tag
func (c_ Control) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// The tag identifying the receiver (not the tag of the receiver’s cell).

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/tag
func (c_ Control) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}
// The target object that receives action messages from the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/target
func (c_ Control) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The target object that receives action messages from the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/target
func (c_ Control) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}
// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}


// SetUsesSingleLineMode sets the value of the usesSingleLineMode property.
// A Boolean value that indicates whether the text in the control’s cell uses single line mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}

