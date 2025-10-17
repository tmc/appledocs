
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [Control] class.
var ControlClass _ControlClass

func init() {
	ControlClass = _ControlClass{objc.GetClass("NSControl")}
}

type _ControlClass struct {
	objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	ID() objc.ID
	AbortEditing() bool
	CurrentEditor() unsafe.Pointer
	DrawWithExpansionFrameInView(contentFrame foundation.Rect, view unsafe.Pointer)
	EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect
	InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer)
	PerformClick(sender objc.ID)
	SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int)
	SendActionOn(mask unsafe.Pointer) int
	SendActionTo(action objc.SEL, target objc.ID) bool
	SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint)
	SizeThatFits(size foundation.Size) foundation.Size
	SizeToFit()
	TakeDoubleValueFrom(sender objc.ID)
	TakeFloatValueFrom(sender objc.ID)
	TakeIntValueFrom(sender objc.ID)
	TakeIntegerValueFrom(sender objc.ID)
	TakeObjectValueFrom(sender objc.ID)
	TakeStringValueFrom(sender objc.ID)
	ValidateEditing()
}

type Control struct {
	id objc.ID
}

func ControlFrom(ptr unsafe.Pointer) Control {
	return Control{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ Control) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ControlClass) Alloc() Control {
	rv := objc.Send[Control](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ControlClass) New() Control {
	rv := objc.Send[Control](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewControl creates and returns a new initialized instance.
func NewControl() Control {
	return ControlClass.New()
}

// Init initializes the instance.
func (c_ Control) Init() Control {
	rv := objc.Send[Control](c_.ID(), selInit)
	return rv
}
// Terminates the current editing operation and discards any edited text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/abortEditing()
func (c_ Control) AbortEditing() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("abortEditing"))
	return rv
}
// Returns the current field editor for the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/currentEditor()
func (c_ Control) CurrentEditor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("currentEditor"))
	return rv
}
// Performs custom expansion tool tip drawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c_ Control) DrawWithExpansionFrameInView(contentFrame foundation.Rect, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("drawWithExpansionFrame:inView:"), contentFrame, view)
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c_ Control) EditWithFrameEditorDelegateEvent(rect foundation.Rect, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("editWithFrame:editor:delegate:event:"), rect, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/endEditing(_:)
func (c_ Control) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("endEditing:"), textObj)
}
// The frame in which a tool tip can be displayed, if needed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := objc.Send[foundation.Rect](c_.ID(), objc.RegisterName("expansionFrameWithFrame:"), contentFrame)
	return rv
}
// Notifies the control that the intrinsic content size for its cell is no longer valid. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("invalidateIntrinsicContentSizeForCell:"), cell)
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/performClick(_:)
func (c_ Control) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("performClick:"), sender)
}
// Selects the specified text range in the receiver’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect foundation.Rect, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("selectWithFrame:editor:delegate:start:length:"), rect, textObj, delegate, selStart, selLength)
}
// Causes the specified action to be sent to the target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sendAction(_:to:)
func (c_ Control) SendActionTo(action objc.SEL, target objc.ID) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("sendAction:to:"), action, target)
	return rv
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sendAction(on:)
func (c_ Control) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("sendActionOn:"), mask)
	return rv
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/setFloatingPointFormat:left:right:
func (c_ Control) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}
// Asks the control to calculate and return the size that best fits the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := objc.Send[foundation.Size](c_.ID(), objc.RegisterName("sizeThatFits:"), size)
	return rv
}
// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sizeToFit()
func (c_ Control) SizeToFit() {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("sizeToFit"))
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c_ Control) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeDoubleValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c_ Control) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeFloatValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c_ Control) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeIntValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an   value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c_ Control) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeIntegerValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c_ Control) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeObjectValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c_ Control) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeStringValueFrom:"), sender)
}
// Validates changes to any user-typed text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/validateEditing()
func (c_ Control) ValidateEditing() {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("validateEditing"))
}
// The default action-message selector associated with the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/action
func (c_ Control) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID(), objc.RegisterName("action"))
	return rv
}
// SetAction sets the value of the action property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/action
func (c_ Control) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAction:"), value)
}
// The alignment mode of the text in the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/alignment
func (c_ Control) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("alignment"))
	return rv
}
// SetAlignment sets the value of the alignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/alignment
func (c_ Control) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAlignment:"), value)
}
// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsExpansionToolTips"))
	return rv
}
// SetAllowsExpansionToolTips sets the value of the allowsExpansionToolTips property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsExpansionToolTips:"), value)
}
// The value of the receiver’s cell as an attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) AttributedStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("attributedStringValue"))
	return rv
}
// SetAttributedStringValue sets the value of the attributedStringValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) SetAttributedStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAttributedStringValue:"), value)
}
// The initial writing direction used to determine the actual writing direction for text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("baseWritingDirection"))
	return rv
}
// SetBaseWritingDirection sets the value of the baseWritingDirection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBaseWritingDirection:"), value)
}
// The size of the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("controlSize"))
	return rv
}
// SetControlSize sets the value of the controlSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setControlSize:"), value)
}
// The value of the receiver’s cell as a double-precision floating-point number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/doubleValue
func (c_ Control) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID(), objc.RegisterName("doubleValue"))
	return rv
}
// SetDoubleValue sets the value of the doubleValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/doubleValue
func (c_ Control) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDoubleValue:"), value)
}
// The value of the receiver’s cell as a single-precision floating-point number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/floatValue
func (c_ Control) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID(), objc.RegisterName("floatValue"))
	return rv
}
// SetFloatValue sets the value of the floatValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/floatValue
func (c_ Control) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFloatValue:"), value)
}
// The font used to draw text in the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/font
func (c_ Control) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("font"))
	return rv
}
// SetFont sets the value of the font property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/font
func (c_ Control) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFont:"), value)
}
// The receiver’s formatter. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/formatter
func (c_ Control) Formatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("formatter"))
	return rv
}
// SetFormatter sets the value of the formatter property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/formatter
func (c_ Control) SetFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFormatter:"), value)
}
// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("ignoresMultiClick"))
	return rv
}
// SetIgnoresMultiClick sets the value of the ignoresMultiClick property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setIgnoresMultiClick:"), value)
}
// The value of the receiver’s cell as an integer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/intValue
func (c_ Control) IntValue() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("intValue"))
	return rv
}
// SetIntValue sets the value of the intValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/intValue
func (c_ Control) SetIntValue(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setIntValue:"), value)
}
// The value of the receiver’s cell as an   value. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/integerValue
func (c_ Control) IntegerValue() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("integerValue"))
	return rv
}
// SetIntegerValue sets the value of the integerValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/integerValue
func (c_ Control) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setIntegerValue:"), value)
}
// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isContinuous
func (c_ Control) Continuous() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("continuous"))
	return rv
}
// SetContinuous sets the value of the continuous property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isContinuous
func (c_ Control) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setContinuous:"), value)
}
// A Boolean value that indicates whether the receiver reacts to mouse events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isEnabled
func (c_ Control) Enabled() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("enabled"))
	return rv
}
// SetEnabled sets the value of the enabled property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isEnabled
func (c_ Control) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setEnabled:"), value)
}
// A Boolean value that indicates whether the cell is highlighted. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) Highlighted() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("highlighted"))
	return rv
}
// SetHighlighted sets the value of the highlighted property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setHighlighted:"), value)
}
// The line break mode to use for text in the control’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("lineBreakMode"))
	return rv
}
// SetLineBreakMode sets the value of the lineBreakMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setLineBreakMode:"), value)
}
// The value of the receiver’s cell as an Objective-C object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/objectValue
func (c_ Control) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("objectValue"))
	return rv
}
// SetObjectValue sets the value of the objectValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/objectValue
func (c_ Control) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setObjectValue:"), value)
}
// A Boolean value indicating whether the receiver refuses the first responder role. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("refusesFirstResponder"))
	return rv
}
// SetRefusesFirstResponder sets the value of the refusesFirstResponder property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setRefusesFirstResponder:"), value)
}
// The value of the receiver’s cell as an   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/stringValue
func (c_ Control) StringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("stringValue"))
	return rv
}
// SetStringValue sets the value of the stringValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/stringValue
func (c_ Control) SetStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setStringValue:"), value)
}
// The tag identifying the receiver (not the tag of the receiver’s cell). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/tag
func (c_ Control) Tag() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("tag"))
	return rv
}
// SetTag sets the value of the tag property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/tag
func (c_ Control) SetTag(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTag:"), value)
}
// The target object that receives action messages from the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/target
func (c_ Control) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("target"))
	return rv
}
// SetTarget sets the value of the target property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/target
func (c_ Control) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTarget:"), value)
}
// A Boolean value that indicates whether the text in the control’s cell uses single line mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("usesSingleLineMode"))
	return rv
}
// SetUsesSingleLineMode sets the value of the usesSingleLineMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setUsesSingleLineMode:"), value)
}
