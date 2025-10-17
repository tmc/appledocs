// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Control] class.
var ControlClass objc.Class

func init() {
	ControlClass = objc.GetClass("NSControl")
}

type Control struct {
	objc.ID
}

func ControlFrom(ptr unsafe.Pointer) Control {
	return Control{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc Control) Alloc() Control {
	ret := objc.ID(ControlClass).Send(objc.RegisterName("alloc"))
	return Control{ret}
}

// Init initializes the instance.
func (c_ Control) Init() Control {
	ret := c_.ID.Send(objc.RegisterName("init"))
	return Control{ret}
}
// Initializes a control with data in an unarchiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/init(coder:)
func NewControlWithCoder(coder unsafe.Pointer) Control {
	instance := Control{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Control{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes a control with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/init(frame:)
func NewControlWithFrame(frameRect unsafe.Pointer) Control {
	instance := Control{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = Control{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Terminates the current editing operation and discards any edited text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/abortEditing()
func (c_ Control) AbortEditing() bool {
	sel := objc.RegisterName("abortEditing")
	ret := c_.ID.Send(sel)
	return ret != 0
}
// Returns the current field editor for the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/currentEditor()
func (c_ Control) CurrentEditor() unsafe.Pointer {
	sel := objc.RegisterName("currentEditor")
	ret := c_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Performs custom expansion tool tip drawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c_ Control) DrawWithExpansionFrameInView(contentFrame unsafe.Pointer, view unsafe.Pointer) {
	sel := objc.RegisterName("drawWithExpansionFrame:inView:")
	c_.ID.Send(sel, contentFrame, view)
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c_ Control) EditWithFrameEditorDelegateEvent(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	sel := objc.RegisterName("editWithFrame:editor:delegate:event:")
	c_.ID.Send(sel, rect, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/endEditing(_:)
func (c_ Control) EndEditing(textObj unsafe.Pointer) {
	sel := objc.RegisterName("endEditing:")
	c_.ID.Send(sel, textObj)
}
// The frame in which a tool tip can be displayed, if needed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c_ Control) ExpansionFrameWithFrame(contentFrame unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("expansionFrameWithFrame:")
	ret := c_.ID.Send(sel, contentFrame)
	return unsafe.Pointer(ret)
}
// Notifies the control that the intrinsic content size for its cell is no longer valid. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell unsafe.Pointer) {
	sel := objc.RegisterName("invalidateIntrinsicContentSizeForCell:")
	c_.ID.Send(sel, cell)
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/performClick(_:)
func (c_ Control) PerformClick(sender objc.ID) {
	sel := objc.RegisterName("performClick:")
	c_.ID.Send(sel, sender)
}
// Selects the specified text range in the receiver’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	sel := objc.RegisterName("selectWithFrame:editor:delegate:start:length:")
	c_.ID.Send(sel, rect, textObj, delegate, selStart, selLength)
}
// Causes the specified action to be sent to the target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sendAction(_:to:)
func (c_ Control) SendActionTo(action objc.SEL, target objc.ID) bool {
	sel := objc.RegisterName("sendAction:to:")
	ret := c_.ID.Send(sel, action, target)
	return ret != 0
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sendAction(on:)
func (c_ Control) SendActionOn(mask unsafe.Pointer) int {
	sel := objc.RegisterName("sendActionOn:")
	ret := c_.ID.Send(sel, mask)
	return int(ret)
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/setFloatingPointFormat:left:right:
func (c_ Control) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	sel := objc.RegisterName("setFloatingPointFormat:left:right:")
	c_.ID.Send(sel, autoRange, leftDigits, rightDigits)
}
// Asks the control to calculate and return the size that best fits the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c_ Control) SizeThatFits(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sizeThatFits:")
	ret := c_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/sizeToFit()
func (c_ Control) SizeToFit() {
	sel := objc.RegisterName("sizeToFit")
	c_.ID.Send(sel)
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c_ Control) TakeDoubleValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeDoubleValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c_ Control) TakeFloatValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeFloatValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c_ Control) TakeIntValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeIntValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to an   value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c_ Control) TakeIntegerValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeIntegerValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c_ Control) TakeObjectValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeObjectValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c_ Control) TakeStringValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeStringValueFrom:")
	c_.ID.Send(sel, sender)
}
// Validates changes to any user-typed text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSControl/validateEditing()
func (c_ Control) ValidateEditing() {
	sel := objc.RegisterName("validateEditing")
	c_.ID.Send(sel)
}

