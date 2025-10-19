// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Control] class.
var controlClass = _ControlClass{objc.GetClass("NSControl")}

type _ControlClass struct {
	class objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	IView
	AbortEditing() bool
	CurrentEditor() unsafe.Pointer
	DrawWithExpansionFrameInView(contentFrame unsafe.Pointer, view unsafe.Pointer)
	EditWithFrameEditorDelegateEvent(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	ExpansionFrameWithFrame(contentFrame unsafe.Pointer) unsafe.Pointer
	InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer)
	PerformClick(sender objc.ID)
	SelectWithFrameEditorDelegateStartLength(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int)
	SendActionTo(action objc.SEL, target objc.ID) bool
	SendActionOn(mask unsafe.Pointer) int
	SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint)
	SizeThatFits(size unsafe.Pointer) unsafe.Pointer
	SizeToFit()
	TakeDoubleValueFrom(sender objc.ID)
	TakeFloatValueFrom(sender objc.ID)
	TakeIntValueFrom(sender objc.ID)
	TakeIntegerValueFrom(sender objc.ID)
	TakeObjectValueFrom(sender objc.ID)
	TakeStringValueFrom(sender objc.ID)
	ValidateEditing()
}

// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return controlClass.New()
}


// Initializes a control with data in an unarchiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewControlWithCoder(coder unsafe.Pointer) Control {
	instance := controlClass.Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
// Initializes a control with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewControlWithFrame(frameRect unsafe.Pointer) Control {
	instance := controlClass.Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Terminates the current editing operation and discards any edited text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/abortEditing()
func (c_ Control) AbortEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("abortEditing"))
	return rv
}
// Returns the current field editor for the control. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/currentEditor()
func (c_ Control) CurrentEditor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentEditor"))
	return rv
}
// Performs custom expansion tool tip drawing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c_ Control) DrawWithExpansionFrameInView(contentFrame unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), contentFrame, view)
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c_ Control) EditWithFrameEditorDelegateEvent(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:editor:delegate:event:"), rect, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/endEditing(_:)
func (c_ Control) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}
// The frame in which a tool tip can be displayed, if needed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c_ Control) ExpansionFrameWithFrame(contentFrame unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("expansionFrameWithFrame:"), contentFrame)
	return rv
}
// Notifies the control that the intrinsic content size for its cell is no longer valid. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateIntrinsicContentSizeForCell:"), cell)
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/performClick(_:)
func (c_ Control) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}
// Selects the specified text range in the receiver’s field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:editor:delegate:start:length:"), rect, textObj, delegate, selStart, selLength)
}
// Causes the specified action to be sent to the target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(_:to:)
func (c_ Control) SendActionTo(action objc.SEL, target objc.ID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendAction:to:"), action, target)
	return rv
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(on:)
func (c_ Control) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/setFloatingPointFormat:left:right:
func (c_ Control) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}
// Asks the control to calculate and return the size that best fits the specified size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c_ Control) SizeThatFits(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sizeThatFits:"), size)
	return rv
}
// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeToFit()
func (c_ Control) SizeToFit() {
	objc.Send[objc.ID](c_.ID, objc.Sel("sizeToFit"))
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c_ Control) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c_ Control) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c_ Control) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c_ Control) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c_ Control) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c_ Control) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}
// Validates changes to any user-typed text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/validateEditing()
func (c_ Control) ValidateEditing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("validateEditing"))
}

