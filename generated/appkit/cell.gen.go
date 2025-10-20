// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Cell] class.
var (
	cellClass     _CellClass
	cellClassOnce sync.Once
)

func getCellClass() _CellClass {
	cellClassOnce.Do(func() {
		cellClass = _CellClass{objc.GetClass("NSCell")}
	})
	return cellClass
}

type _CellClass struct {
	class objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	objectivec.IObject
	CalcDrawInfo(rect coregraphics.CGRect)
	CellAttribute(parameter unsafe.Pointer) int
	CellSizeForBounds(rect coregraphics.CGRect) coregraphics.CGSize
	Compare(otherCell objc.ID) unsafe.Pointer
	ContinueTrackingAtInView(lastPoint coregraphics.CGPoint, currentPoint coregraphics.CGPoint, controlView unsafe.Pointer) bool
	DraggingImageComponentsWithFrameInView(frame coregraphics.CGRect, view unsafe.Pointer) []DraggingImageComponent
	DrawWithExpansionFrameInView(cellFrame coregraphics.CGRect, view unsafe.Pointer)
	DrawWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer)
	DrawFocusRingMaskWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer)
	DrawInteriorWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer)
	DrawingRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	EditWithFrameInViewEditorDelegateEvent(rect coregraphics.CGRect, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	EntryType() int
	ExpansionFrameWithFrameInView(cellFrame coregraphics.CGRect, view unsafe.Pointer) coregraphics.CGRect
	FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer
	FocusRingMaskBoundsForFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) coregraphics.CGRect
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	HighlightWithFrameInView(flag bool, cellFrame coregraphics.CGRect, controlView unsafe.Pointer)
	HighlightColorWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) unsafe.Pointer
	HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame coregraphics.CGRect, controlView unsafe.Pointer) unsafe.Pointer
	ImageRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	IsEntryAcceptable(string string) bool
	MenuForEventInRectOfView(event unsafe.Pointer, cellFrame coregraphics.CGRect, view unsafe.Pointer) unsafe.Pointer
	Mnemonic() unsafe.Pointer
	MnemonicLocation() uint
	PerformClick(sender objc.ID)
	ResetCursorRectInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer)
	SelectWithFrameInViewEditorDelegateStartLength(rect coregraphics.CGRect, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int)
	SendActionOn(mask unsafe.Pointer) int
	SetCellAttributeTo(parameter unsafe.Pointer, value int)
	SetEntryType(type_ int)
	SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint)
	SetMnemonicLocation(location uint)
	SetNextState()
	SetTitleWithMnemonic(stringWithAmpersand string)
	SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer
	StartTrackingAtInView(startPoint coregraphics.CGPoint, controlView unsafe.Pointer) bool
	StopTrackingAtInViewMouseIsUp(lastPoint coregraphics.CGPoint, stopPoint coregraphics.CGPoint, controlView unsafe.Pointer, flag bool)
	TakeDoubleValueFrom(sender objc.ID)
	TakeFloatValueFrom(sender objc.ID)
	TakeIntValueFrom(sender objc.ID)
	TakeIntegerValueFrom(sender objc.ID)
	TakeObjectValueFrom(sender objc.ID)
	TakeStringValueFrom(sender objc.ID)
	TitleRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame coregraphics.CGRect, controlView unsafe.Pointer, flag bool) bool
}

// A mechanism for displaying text or images in a view object without the overhead of a full subclass.
//
// Cells are used by most of the classes to implement their internal workings.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getCellClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder unsafe.Pointer) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Returns an object initialized with the specified image and set to have the cell’s default menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewCellImageCell(image unsafe.Pointer) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}

// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewCellTextCell(string string) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initTextCell:"), objc.String(string))
	rv.Autorelease()
	return rv
}


// Recalculates the cell geometry.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calcDrawInfo:"), rect)
}

// Returns the value for the specified cell attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("cellAttribute:"), parameter)
	return rv
}

// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect coregraphics.CGRect) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("cellSizeForBounds:"), rect)
	return rv
}

// Compares the string values of the receiver another cell, disregarding case.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compare:"), otherCell)
	return rv
}

// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint coregraphics.CGPoint, currentPoint coregraphics.CGPoint, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}

// Generates dragging image components with the specified frame in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame coregraphics.CGRect, view unsafe.Pointer) []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](c_.ID, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}

// Instructs the receiver to draw in an expansion frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame coregraphics.CGRect, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, view)
}

// Draws the receiver’s border and then draws the interior of the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}

// Draws the focus ring for the control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}

// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}

// Returns the rectangle within which the receiver draws itself
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("drawingRectForBounds:"), rect)
	return rv
}

// Begins editing of the receiver’s text using the specified field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect coregraphics.CGRect, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}

// Ends the editing of text in the receiver using the specified field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}

// Returns the type of data the user can type into the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/entryType
func (c_ Cell) EntryType() int {
	rv := objc.Send[int](c_.ID, objc.Sel("entryType"))
	return rv
}

// Returns the expansion cell frame for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame coregraphics.CGRect, view unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}

// Returns a custom field editor for editing in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fieldEditorForView:"), controlView)
	return rv
}

// Returns the bounds of the focus ring mask.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}

// Returns the initial delay and repeat values for continuous sending of action messages to target objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}

// Redraws the receiver with the specified highlight setting.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame coregraphics.CGRect, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}

// Returns the color the receiver uses when drawing the selection highlight.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}

// Returns hit testing information for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame coregraphics.CGRect, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}

// Returns the rectangle in which the receiver draws its image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("imageRectForBounds:"), rect)
	return rv
}

// Returns whether a string representing a numeric or date value is formatted in a suitable way for the cell’s entry type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEntryAcceptable:
func (c_ Cell) IsEntryAcceptable(string string) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEntryAcceptable:"), objc.String(string))
	return rv
}

// Returns the menu associated with the cell and related to the specified event and frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event unsafe.Pointer, cellFrame coregraphics.CGRect, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}

// Returns the character in the receiver’s title that appears underlined for use as a mnemonic.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mnemonic
func (c_ Cell) Mnemonic() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("mnemonic"))
	return rv
}

// Returns the position of the underlined mnemonic character in the receiver’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mnemonicLocation
func (c_ Cell) MnemonicLocation() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("mnemonicLocation"))
	return rv
}

// Simulates a single mouse click on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}

// Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame coregraphics.CGRect, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetCursorRect:inView:"), cellFrame, controlView)
}

// Selects the specified text range in the cell’s field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect coregraphics.CGRect, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}

// Sets the conditions on which the receiver sends action messages to its target.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}

// Sets the value for the specified cell attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter unsafe.Pointer, value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellAttribute:to:"), parameter, value)
}

// Sets how numeric data is formatted in the receiver and places restrictions on acceptable input.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setEntryType:
func (c_ Cell) SetEntryType(type_ int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEntryType:"), type_)
}

// Sets the auto-ranging and floating point number format of the receiver’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setFloatingPointFormat:left:right:
func (c_ Cell) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}

// Sets the character of the receiver’s title to be used as a mnemonic character.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setMnemonicLocation:
func (c_ Cell) SetMnemonicLocation(location uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMnemonicLocation:"), location)
}

// Changes cell’s state to the next value in the sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setNextState()
func (c_ Cell) SetNextState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNextState"))
}

// Sets the title of the receiver with one character in the string denoted as an access key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setTitleWithMnemonic:
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitleWithMnemonic:"), objc.String(stringWithAmpersand))
}

// Configures the textual and background attributes of the receiver’s field editor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}

// Begins tracking mouse events within the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint coregraphics.CGPoint, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}

// Stops tracking mouse events within the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint coregraphics.CGPoint, stopPoint coregraphics.CGPoint, controlView unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}

// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}

// Returns the rectangle in which the receiver draws its title text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("titleRectForBounds:"), rect)
	return rv
}

// Initiates the mouse tracking behavior in a cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame coregraphics.CGRect, controlView unsafe.Pointer, flag bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}

// A Boolean value indicating whether the cell accepts first responder status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/acceptsFirstResponder
func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}

// The action performed by the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The action performed by the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}
// The alignment of the cell’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The alignment of the cell’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}
// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}


// SetAllowsEditingTextAttributes sets the value of the allowsEditingTextAttributes property.
// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}
// A Boolean value indicating whether the cell supports three states instead of two.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) AllowsMixedState() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMixedState"))
	return rv
}


// SetAllowsMixedState sets the value of the allowsMixedState property.
// A Boolean value indicating whether the cell supports three states instead of two.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMixedState:"), value)
}
// A Boolean value indicating whether the cell assumes responsibility for undo operations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) AllowsUndo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsUndo"))
	return rv
}


// SetAllowsUndo sets the value of the allowsUndo property.
// A Boolean value indicating whether the cell assumes responsibility for undo operations.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsUndo:"), value)
}
// The cell’s value as an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) AttributedStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}


// SetAttributedStringValue sets the value of the attributedStringValue property.
// The cell’s value as an attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) SetAttributedStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}
// The cell’s background style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) BackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("backgroundStyle"))
	return rv
}


// SetBackgroundStyle sets the value of the backgroundStyle property.
// The cell’s background style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) SetBackgroundStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundStyle:"), value)
}
// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// SetBaseWritingDirection sets the value of the baseWritingDirection property.
// The initial writing direction used to determine the actual writing direction for text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}
// The minimum size needed to display the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize
func (c_ Cell) CellSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("cellSize"))
	return rv
}

// The size of the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlSize"))
	return rv
}


// SetControlSize sets the value of the controlSize property.
// The size of the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}
// The cell’s control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlTint"))
	return rv
}


// SetControlTint sets the value of the controlTint property.
// The cell’s control tint.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlTint:"), value)
}
// The view associated with the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) ControlView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlView"))
	return rv
}


// SetControlView sets the value of the controlView property.
// The view associated with the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) SetControlView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlView:"), value)
}
// The cell’s value as a double-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) DoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The cell’s value as a double-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) SetDoubleValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}
// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) FloatValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("floatValue"))
	return rv
}


// SetFloatValue sets the value of the floatValue property.
// The cell’s value as a single-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) SetFloatValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}
// The type of focus ring to use with the associated view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) FocusRingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusRingType"))
	return rv
}


// SetFocusRingType sets the value of the focusRingType property.
// The type of focus ring to use with the associated view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) SetFocusRingType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusRingType:"), value)
}
// The font that the cell uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font that the cell uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}
// The cell’s formatter object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) Formatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatter"))
	return rv
}


// SetFormatter sets the value of the formatter property.
// The cell’s formatter object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) SetFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}
// A Boolean value that indicates whether the cell has a valid object value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hasValidObjectValue
func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidObjectValue"))
	return rv
}

// The image displayed by the cell, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The image displayed by the cell, if any.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
}
// A Boolean value indicating whether the cell supports the importation of images into its text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) ImportsGraphics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("importsGraphics"))
	return rv
}


// SetImportsGraphics sets the value of the importsGraphics property.
// A Boolean value indicating whether the cell supports the importation of images into its text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportsGraphics:"), value)
}
// The cell’s value as an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) IntValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("intValue"))
	return rv
}


// SetIntValue sets the value of the intValue property.
// The cell’s value as an integer.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) SetIntValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}
// The cell’s value as an type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}


// SetIntegerValue sets the value of the integerValue property.
// The cell’s value as an type.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}
// The cell’s interior background style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/interiorBackgroundStyle
func (c_ Cell) InteriorBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}

// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) Bezeled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bezeled"))
	return rv
}


// SetBezeled sets the value of the bezeled property.
// A Boolean value indicating whether the cell has a bezeled border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) SetBezeled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBezeled:"), value)
}
// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) Bordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bordered"))
	return rv
}


// SetBordered sets the value of the bordered property.
// A Boolean value indicating whether the cell draws itself outlined with a plain border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBordered:"), value)
}
// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}


// SetContinuous sets the value of the continuous property.
// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}
// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) Editable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editable"))
	return rv
}


// SetEditable sets the value of the editable property.
// A Boolean value indicating whether the cell is editable.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) SetEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditable:"), value)
}
// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value indicating whether the cell is currently enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}
// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) Highlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highlighted"))
	return rv
}


// SetHighlighted sets the value of the highlighted property.
// A Boolean value indicating whether the cell has a highlighted appearance.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlighted:"), value)
}
// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isOpaque
func (c_ Cell) Opaque() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("opaque"))
	return rv
}

// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) Scrollable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("scrollable"))
	return rv
}


// SetScrollable sets the value of the scrollable property.
// A Boolean value indicating whether excess text scrolls past the cell’s bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) SetScrollable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollable:"), value)
}
// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) Selectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("selectable"))
	return rv
}


// SetSelectable sets the value of the selectable property.
// A Boolean value indicating whether the cell’s text can be selected.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectable:"), value)
}
// The key equivalent associated with clicking the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/keyEquivalent
func (c_ Cell) KeyEquivalent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("keyEquivalent"))
	return rv
}

// The line break mode to use when drawing text in the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// SetLineBreakMode sets the value of the lineBreakMode property.
// The line break mode to use when drawing text in the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}
// The cell’s contextual menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The cell’s contextual menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) SetMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMenu:"), value)
}
// The modifier flags for the last (left) mouse-down event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mouseDownFlags
func (c_ Cell) MouseDownFlags() int {
	rv := objc.Send[int](c_.ID, objc.Sel("mouseDownFlags"))
	return rv
}

// The cell’s next state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/nextState
func (c_ Cell) NextState() int {
	rv := objc.Send[int](c_.ID, objc.Sel("nextState"))
	return rv
}

// The cell’s value as an Objective-C object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValue"))
	return rv
}


// SetObjectValue sets the value of the objectValue property.
// The cell’s value as an Objective-C object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}
// A Boolean value indicating whether the cell refuses the first responder status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}


// SetRefusesFirstResponder sets the value of the refusesFirstResponder property.
// A Boolean value indicating whether the cell refuses the first responder status.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}
// The object represented by the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("representedObject"))
	return rv
}


// SetRepresentedObject sets the value of the representedObject property.
// The object represented by the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRepresentedObject:"), value)
}
// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendsActionOnEndEditing"))
	return rv
}


// SetSendsActionOnEndEditing sets the value of the sendsActionOnEndEditing property.
// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSendsActionOnEndEditing:"), value)
}
// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsFirstResponder"))
	return rv
}


// SetShowsFirstResponder sets the value of the showsFirstResponder property.
// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsFirstResponder:"), value)
}
// The cell’s current state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The cell’s current state.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}
// The cell’s value as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) StringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
// The cell’s value as a string.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) SetStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), value)
}
// A tag for identifying the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// A tag for identifying the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}
// The object that receives the cell’s action messages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The object that receives the cell’s action messages.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}
// The cell’s title text.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The cell’s title text.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}
// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("truncatesLastVisibleLine"))
	return rv
}


// SetTruncatesLastVisibleLine sets the value of the truncatesLastVisibleLine property.
// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTruncatesLastVisibleLine:"), value)
}
// The type of the cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of the cell.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}
// The layout direction of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) UserInterfaceLayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// The layout direction of the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) SetUserInterfaceLayoutDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}
// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}


// SetUsesSingleLineMode sets the value of the usesSingleLineMode property.
// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}
// A Boolean value indicating whether the cell’s field editor should post text change notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wantsNotificationForMarkedText
func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wantsNotificationForMarkedText"))
	return rv
}

// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) Wraps() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wraps"))
	return rv
}


// SetWraps sets the value of the wraps property.
// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) SetWraps(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWraps:"), value)
}

