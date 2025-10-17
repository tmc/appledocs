// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Cell] class.
var CellClass objc.Class

func init() {
	CellClass = objc.GetClass("NSCell")
}

type Cell struct {
	objc.ID
}

func CellFrom(ptr unsafe.Pointer) Cell {
	return Cell{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc Cell) Alloc() Cell {
	ret := objc.ID(CellClass).Send(objc.RegisterName("alloc"))
	return Cell{ret}
}

// Init initializes the instance.
func (c_ Cell) Init() Cell {
	ret := c_.ID.Send(objc.RegisterName("init"))
	return Cell{ret}
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init()
func NewCell() Cell {
	instance := Cell{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder unsafe.Pointer) Cell {
	instance := Cell{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Cell{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an   object initialized with the specified image and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(imageCell:)
func NewCellImageCell(image unsafe.Pointer) Cell {
	instance := Cell{}.Alloc()
	sel := objc.RegisterName("initImageCell:")
	ret := instance.ID.Send(sel, image)
	instance = Cell{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(textCell:)
func NewCellTextCell(string string) Cell {
	instance := Cell{}.Alloc()
	sel := objc.RegisterName("initTextCell:")
	ret := instance.ID.Send(sel, string)
	instance = Cell{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Recalculates the cell geometry. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect unsafe.Pointer) {
	sel := objc.RegisterName("calcDrawInfo:")
	c_.ID.Send(sel, rect)
}
// Returns the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter unsafe.Pointer) int {
	sel := objc.RegisterName("cellAttribute:")
	ret := c_.ID.Send(sel, parameter)
	return int(ret)
}
// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("cellSizeForBounds:")
	ret := c_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Compares the string values of the receiver another cell, disregarding case. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("compare:")
	ret := c_.ID.Send(sel, otherCell)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint unsafe.Pointer, currentPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	sel := objc.RegisterName("continueTracking:at:inView:")
	ret := c_.ID.Send(sel, lastPoint, currentPoint, controlView)
	return ret != 0
}
// Generates dragging image components with the specified frame in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("draggingImageComponentsWithFrame:inView:")
	ret := c_.ID.Send(sel, frame, view)
	return unsafe.Pointer(ret)
}
// Instructs the receiver to draw in an expansion frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) {
	sel := objc.RegisterName("drawWithExpansionFrame:inView:")
	c_.ID.Send(sel, cellFrame, view)
}
// Draws the receiver’s border and then draws the interior of the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	sel := objc.RegisterName("drawWithFrame:inView:")
	c_.ID.Send(sel, cellFrame, controlView)
}
// Draws the focus ring for the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	sel := objc.RegisterName("drawFocusRingMaskWithFrame:inView:")
	c_.ID.Send(sel, cellFrame, controlView)
}
// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	sel := objc.RegisterName("drawInteriorWithFrame:inView:")
	c_.ID.Send(sel, cellFrame, controlView)
}
// Returns the rectangle within which the receiver draws itself [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("drawingRectForBounds:")
	ret := c_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	sel := objc.RegisterName("editWithFrame:inView:editor:delegate:event:")
	c_.ID.Send(sel, rect, controlView, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj unsafe.Pointer) {
	sel := objc.RegisterName("endEditing:")
	c_.ID.Send(sel, textObj)
}
// Returns the type of data the user can type into the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/entryType
func (c_ Cell) EntryType() int {
	sel := objc.RegisterName("entryType")
	ret := c_.ID.Send(sel)
	return int(ret)
}
// Returns the expansion cell frame for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("expansionFrameWithFrame:inView:")
	ret := c_.ID.Send(sel, cellFrame, view)
	return unsafe.Pointer(ret)
}
// Returns a custom field editor for editing in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("fieldEditorForView:")
	ret := c_.ID.Send(sel, controlView)
	return unsafe.Pointer(ret)
}
// Returns the bounds of the focus ring mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("focusRingMaskBoundsForFrame:inView:")
	ret := c_.ID.Send(sel, cellFrame, controlView)
	return unsafe.Pointer(ret)
}
// Returns the initial delay and repeat values for continuous sending of action messages to target objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	sel := objc.RegisterName("getPeriodicDelay:interval:")
	c_.ID.Send(sel, delay, interval)
}
// Redraws the receiver with the specified highlight setting. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	sel := objc.RegisterName("highlight:withFrame:inView:")
	c_.ID.Send(sel, flag, cellFrame, controlView)
}
// Returns the color the receiver uses when drawing the selection highlight. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("highlightColorWithFrame:inView:")
	ret := c_.ID.Send(sel, cellFrame, controlView)
	return unsafe.Pointer(ret)
}
// Returns hit testing information for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("hitTestForEvent:inRect:ofView:")
	ret := c_.ID.Send(sel, event, cellFrame, controlView)
	return unsafe.Pointer(ret)
}
// Returns the rectangle in which the receiver draws its image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("imageRectForBounds:")
	ret := c_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Returns whether a string representing a numeric or date value is formatted in a suitable way for the cell’s entry type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEntryAcceptable:
func (c_ Cell) IsEntryAcceptable(string string) bool {
	sel := objc.RegisterName("isEntryAcceptable:")
	ret := c_.ID.Send(sel, string)
	return ret != 0
}
// Returns the menu associated with the cell and related to the specified event and frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("menuForEvent:inRect:ofView:")
	ret := c_.ID.Send(sel, event, cellFrame, view)
	return unsafe.Pointer(ret)
}
// Returns the character in the receiver’s title that appears underlined for use as a mnemonic. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/mnemonic
func (c_ Cell) Mnemonic() unsafe.Pointer {
	sel := objc.RegisterName("mnemonic")
	ret := c_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the position of the underlined mnemonic character in the receiver’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/mnemonicLocation
func (c_ Cell) MnemonicLocation() uint {
	sel := objc.RegisterName("mnemonicLocation")
	ret := c_.ID.Send(sel)
	return uint(ret)
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objc.ID) {
	sel := objc.RegisterName("performClick:")
	c_.ID.Send(sel, sender)
}
// Sets the receiver to show the I-beam cursor while it tracks the mouse. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	sel := objc.RegisterName("resetCursorRect:inView:")
	c_.ID.Send(sel, cellFrame, controlView)
}
// Selects the specified text range in the cell’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	sel := objc.RegisterName("selectWithFrame:inView:editor:delegate:start:length:")
	c_.ID.Send(sel, rect, controlView, textObj, delegate, selStart, selLength)
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask unsafe.Pointer) int {
	sel := objc.RegisterName("sendActionOn:")
	ret := c_.ID.Send(sel, mask)
	return int(ret)
}
// Sets the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter unsafe.Pointer, value int) {
	sel := objc.RegisterName("setCellAttribute:to:")
	c_.ID.Send(sel, parameter, value)
}
// Sets how numeric data is formatted in the receiver and places restrictions on acceptable input. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setEntryType:
func (c_ Cell) SetEntryType(type_ int) {
	sel := objc.RegisterName("setEntryType:")
	c_.ID.Send(sel, type_)
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setFloatingPointFormat:left:right:
func (c_ Cell) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	sel := objc.RegisterName("setFloatingPointFormat:left:right:")
	c_.ID.Send(sel, autoRange, leftDigits, rightDigits)
}
// Sets the character of the receiver’s title to be used as a mnemonic character. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setMnemonicLocation:
func (c_ Cell) SetMnemonicLocation(location uint) {
	sel := objc.RegisterName("setMnemonicLocation:")
	c_.ID.Send(sel, location)
}
// Changes cell’s state to the next value in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setNextState()
func (c_ Cell) SetNextState() {
	sel := objc.RegisterName("setNextState")
	c_.ID.Send(sel)
}
// Sets the title of the receiver with one character in the string denoted as an access key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setTitleWithMnemonic:
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	sel := objc.RegisterName("setTitleWithMnemonic:")
	c_.ID.Send(sel, stringWithAmpersand)
}
// Configures the textual and background attributes of the receiver’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("setUpFieldEditorAttributes:")
	ret := c_.ID.Send(sel, textObj)
	return unsafe.Pointer(ret)
}
// Begins tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	sel := objc.RegisterName("startTrackingAt:inView:")
	ret := c_.ID.Send(sel, startPoint, controlView)
	return ret != 0
}
// Stops tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint unsafe.Pointer, stopPoint unsafe.Pointer, controlView unsafe.Pointer, flag bool) {
	sel := objc.RegisterName("stopTracking:at:inView:mouseIsUp:")
	c_.ID.Send(sel, lastPoint, stopPoint, controlView, flag)
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeDoubleValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeFloatValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeIntValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeIntegerValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeObjectValueFrom:")
	c_.ID.Send(sel, sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objc.ID) {
	sel := objc.RegisterName("takeStringValueFrom:")
	c_.ID.Send(sel, sender)
}
// Returns the rectangle in which the receiver draws its title text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("titleRectForBounds:")
	ret := c_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Initiates the mouse tracking behavior in a cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer, flag bool) bool {
	sel := objc.RegisterName("trackMouse:inRect:ofView:untilMouseUp:")
	ret := c_.ID.Send(sel, event, cellFrame, controlView, flag)
	return ret != 0
}

