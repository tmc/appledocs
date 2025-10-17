// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cell] class.
var cellClass = _CellClass{objc.GetClass("NSCell")}

type _CellClass struct {
	class objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	objectivec.IObject
	CalcDrawInfo(rect unsafe.Pointer)
	CellAttribute(parameter unsafe.Pointer) int
	CellSizeForBounds(rect unsafe.Pointer) unsafe.Pointer
	Compare(otherCell objc.ID) unsafe.Pointer
	ContinueTrackingAtInView(lastPoint unsafe.Pointer, currentPoint unsafe.Pointer, controlView unsafe.Pointer) bool
	DraggingImageComponentsWithFrameInView(frame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	DrawWithExpansionFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer)
	DrawWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawFocusRingMaskWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawInteriorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawingRectForBounds(rect unsafe.Pointer) unsafe.Pointer
	EditWithFrameInViewEditorDelegateEvent(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	EntryType() int
	ExpansionFrameWithFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer
	FocusRingMaskBoundsForFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	HighlightWithFrameInView(flag bool, cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	HighlightColorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	ImageRectForBounds(rect unsafe.Pointer) unsafe.Pointer
	IsEntryAcceptable(string string) bool
	MenuForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	Mnemonic() unsafe.Pointer
	MnemonicLocation() uint
	PerformClick(sender objc.ID)
	ResetCursorRectInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	SelectWithFrameInViewEditorDelegateStartLength(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int)
	SendActionOn(mask unsafe.Pointer) int
	SetCellAttributeTo(parameter unsafe.Pointer, value int)
	SetEntryType(type_ int)
	SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint)
	SetMnemonicLocation(location uint)
	SetNextState()
	SetTitleWithMnemonic(stringWithAmpersand string)
	SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer
	StartTrackingAtInView(startPoint unsafe.Pointer, controlView unsafe.Pointer) bool
	StopTrackingAtInViewMouseIsUp(lastPoint unsafe.Pointer, stopPoint unsafe.Pointer, controlView unsafe.Pointer, flag bool)
	TakeDoubleValueFrom(sender objc.ID)
	TakeFloatValueFrom(sender objc.ID)
	TakeIntValueFrom(sender objc.ID)
	TakeIntegerValueFrom(sender objc.ID)
	TakeObjectValueFrom(sender objc.ID)
	TakeStringValueFrom(sender objc.ID)
	TitleRectForBounds(rect unsafe.Pointer) unsafe.Pointer
	TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer, flag bool) bool
}

// A mechanism for displaying text or images in a view object without the overhead of a full subclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell

type Cell struct {
	objectivec.Object
}

// CellFrom constructs a [Cell] from an unsafe.Pointer.
//
// A mechanism for displaying text or images in a view object without the overhead of a full subclass.
func CellFrom(ptr unsafe.Pointer) Cell {
	return Cell{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (cc _CellClass) Alloc() Cell {
	rv := objc.Send[Cell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _CellClass) New() Cell {
	rv := objc.Send[Cell](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cell) Init() Cell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cell) Autorelease() Cell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCell creates a new Cell instance.
func NewCell() Cell {
	return cellClass.New()
}
// Returns an object initialized with the specified image and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewCellImageCell(image unsafe.Pointer) Cell {
	instance := cellClass.Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}
// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewCellTextCell(string string) Cell {
	instance := cellClass.Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initTextCell:"), string)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder unsafe.Pointer) Cell {
	instance := cellClass.Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Recalculates the cell geometry. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calcDrawInfo:"), rect)
}
// Returns the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("cellAttribute:"), parameter)
	return rv
}
// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cellSizeForBounds:"), rect)
	return rv
}
// Compares the string values of the receiver another cell, disregarding case. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compare:"), otherCell)
	return rv
}
// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint unsafe.Pointer, currentPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}
// Generates dragging image components with the specified frame in the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}
// Instructs the receiver to draw in an expansion frame. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, view)
}
// Draws the receiver’s border and then draws the interior of the cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}
// Draws the focus ring for the control. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}
// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}
// Returns the rectangle within which the receiver draws itself [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("drawingRectForBounds:"), rect)
	return rv
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}
// Returns the type of data the user can type into the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/entryType
func (c_ Cell) EntryType() int {
	rv := objc.Send[int](c_.ID, objc.Sel("entryType"))
	return rv
}
// Returns the expansion cell frame for the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}
// Returns a custom field editor for editing in the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fieldEditorForView:"), controlView)
	return rv
}
// Returns the bounds of the focus ring mask. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}
// Returns the initial delay and repeat values for continuous sending of action messages to target objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}
// Redraws the receiver with the specified highlight setting. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}
// Returns the color the receiver uses when drawing the selection highlight. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}
// Returns hit testing information for the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}
// Returns the rectangle in which the receiver draws its image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("imageRectForBounds:"), rect)
	return rv
}
// Returns whether a string representing a numeric or date value is formatted in a suitable way for the cell’s entry type. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEntryAcceptable:
func (c_ Cell) IsEntryAcceptable(string string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEntryAcceptable:"), string)
	return rv
}
// Returns the menu associated with the cell and related to the specified event and frame. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}
// Returns the character in the receiver’s title that appears underlined for use as a mnemonic. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mnemonic
func (c_ Cell) Mnemonic() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mnemonic"))
	return rv
}
// Returns the position of the underlined mnemonic character in the receiver’s title. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mnemonicLocation
func (c_ Cell) MnemonicLocation() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("mnemonicLocation"))
	return rv
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}
// Sets the receiver to show the I-beam cursor while it tracks the mouse. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetCursorRect:inView:"), cellFrame, controlView)
}
// Selects the specified text range in the cell’s field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}
// Sets the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter unsafe.Pointer, value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellAttribute:to:"), parameter, value)
}
// Sets how numeric data is formatted in the receiver and places restrictions on acceptable input. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setEntryType:
func (c_ Cell) SetEntryType(type_ int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEntryType:"), type_)
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setFloatingPointFormat:left:right:
func (c_ Cell) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}
// Sets the character of the receiver’s title to be used as a mnemonic character. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setMnemonicLocation:
func (c_ Cell) SetMnemonicLocation(location uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMnemonicLocation:"), location)
}
// Changes cell’s state to the next value in the sequence. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setNextState()
func (c_ Cell) SetNextState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNextState"))
}
// Sets the title of the receiver with one character in the string denoted as an access key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setTitleWithMnemonic:
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleWithMnemonic:"), stringWithAmpersand)
}
// Configures the textual and background attributes of the receiver’s field editor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}
// Begins tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}
// Stops tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint unsafe.Pointer, stopPoint unsafe.Pointer, controlView unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}
// Returns the rectangle in which the receiver draws its title text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("titleRectForBounds:"), rect)
	return rv
}
// Initiates the mouse tracking behavior in a cell. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer, flag bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}

