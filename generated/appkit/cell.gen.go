
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Cell] class.
var CellClass _CellClass

func init() {
	CellClass = _CellClass{objc.GetClass("NSCell")}
}

type _CellClass struct {
	objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	ID() objc.ID
	CalcDrawInfo(rect unsafe.Pointer)
	CellAttribute(parameter unsafe.Pointer) int
	CellSizeForBounds(rect unsafe.Pointer) unsafe.Pointer
	Compare(otherCell objc.ID) unsafe.Pointer
	ContinueTrackingAtInView(lastPoint unsafe.Pointer, currentPoint unsafe.Pointer, controlView unsafe.Pointer) bool
	DraggingImageComponentsWithFrameInView(frame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	DrawFocusRingMaskWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawInteriorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawWithExpansionFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer)
	DrawWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	DrawingRectForBounds(rect unsafe.Pointer) unsafe.Pointer
	EditWithFrameInViewEditorDelegateEvent(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer)
	EndEditing(textObj unsafe.Pointer)
	EntryType() int
	ExpansionFrameWithFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer
	FocusRingMaskBoundsForFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	GetPeriodicDelayInterval(delay float32, interval float32)
	HighlightColorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	HighlightWithFrameInView(flag bool, cellFrame unsafe.Pointer, controlView unsafe.Pointer)
	HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer
	ImageRectForBounds(rect unsafe.Pointer) unsafe.Pointer
	InitImageCell(image unsafe.Pointer) unsafe.Pointer
	InitTextCell(string unsafe.Pointer) unsafe.Pointer
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	IsEntryAcceptable(string unsafe.Pointer) bool
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
	SetTitleWithMnemonic(stringWithAmpersand unsafe.Pointer)
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

type Cell struct {
	id objc.ID
}

func CellFrom(ptr unsafe.Pointer) Cell {
	return Cell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ Cell) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CellClass) Alloc() Cell {
	rv := objc.Send[Cell](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CellClass) New() Cell {
	rv := objc.Send[Cell](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCell creates and returns a new initialized instance.
func NewCell() Cell {
	return CellClass.New()
}

// Init initializes the instance.
func (c_ Cell) Init() Cell {
	rv := objc.Send[Cell](c_.ID(), selInit)
	return rv
}
// Recalculates the cell geometry. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("calcDrawInfo:"), rect)
}
// Returns the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("cellAttribute:"), parameter)
	return rv
}
// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("cellSizeForBounds:"), rect)
	return rv
}
// Compares the string values of the receiver another cell, disregarding case. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("compare:"), otherCell)
	return rv
}
// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint unsafe.Pointer, currentPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}
// Generates dragging image components with the specified frame in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}
// Instructs the receiver to draw in an expansion frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("drawWithExpansionFrame:inView:"), cellFrame, view)
}
// Draws the receiver’s border and then draws the interior of the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("drawWithFrame:inView:"), cellFrame, controlView)
}
// Draws the focus ring for the control. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}
// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}
// Returns the rectangle within which the receiver draws itself [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("drawingRectForBounds:"), rect)
	return rv
}
// Begins editing of the receiver’s text using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, event unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}
// Ends the editing of text in the receiver using the specified field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("endEditing:"), textObj)
}
// Returns the type of data the user can type into the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/entryType
func (c_ Cell) EntryType() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("entryType"))
	return rv
}
// Returns the expansion cell frame for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}
// Returns a custom field editor for editing in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("fieldEditorForView:"), controlView)
	return rv
}
// Returns the bounds of the focus ring mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}
// Returns the initial delay and repeat values for continuous sending of action messages to target objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay float32, interval float32) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("getPeriodicDelay:interval:"), delay, interval)
}
// Redraws the receiver with the specified highlight setting. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}
// Returns the color the receiver uses when drawing the selection highlight. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}
// Returns hit testing information for the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}
// Returns the rectangle in which the receiver draws its image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("imageRectForBounds:"), rect)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(coder:)
func (c_ Cell) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
// Returns an   object initialized with the specified image and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(imageCell:)
func (c_ Cell) InitImageCell(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("initImageCell:"), image)
	return rv
}
// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/init(textCell:)
func (c_ Cell) InitTextCell(string unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("initTextCell:"), string)
	return rv
}
// Returns whether a string representing a numeric or date value is formatted in a suitable way for the cell’s entry type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEntryAcceptable:
func (c_ Cell) IsEntryAcceptable(string unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("isEntryAcceptable:"), string)
	return rv
}
// Returns the menu associated with the cell and related to the specified event and frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event unsafe.Pointer, cellFrame unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}
// Returns the character in the receiver’s title that appears underlined for use as a mnemonic. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/mnemonic
func (c_ Cell) Mnemonic() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("mnemonic"))
	return rv
}
// Returns the position of the underlined mnemonic character in the receiver’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/mnemonicLocation
func (c_ Cell) MnemonicLocation() uint {
	rv := objc.Send[uint](c_.ID(), objc.RegisterName("mnemonicLocation"))
	return rv
}
// Simulates a single mouse click on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("performClick:"), sender)
}
// Sets the receiver to show the I-beam cursor while it tracks the mouse. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame unsafe.Pointer, controlView unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("resetCursorRect:inView:"), cellFrame, controlView)
}
// Selects the specified text range in the cell’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect unsafe.Pointer, controlView unsafe.Pointer, textObj unsafe.Pointer, delegate objc.ID, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}
// Sets the conditions on which the receiver sends action messages to its target. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask unsafe.Pointer) int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("sendActionOn:"), mask)
	return rv
}
// Sets the value for the specified cell attribute. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter unsafe.Pointer, value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setCellAttribute:to:"), parameter, value)
}
// Sets how numeric data is formatted in the receiver and places restrictions on acceptable input. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setEntryType:
func (c_ Cell) SetEntryType(type_ int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setEntryType:"), type_)
}
// Sets the auto-ranging and floating point number format of the receiver’s cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setFloatingPointFormat:left:right:
func (c_ Cell) SetFloatingPointFormatLeftRight(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}
// Sets the character of the receiver’s title to be used as a mnemonic character. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setMnemonicLocation:
func (c_ Cell) SetMnemonicLocation(location uint) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMnemonicLocation:"), location)
}
// Changes cell’s state to the next value in the sequence. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setNextState()
func (c_ Cell) SetNextState() {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setNextState"))
}
// Sets the title of the receiver with one character in the string denoted as an access key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setTitleWithMnemonic:
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTitleWithMnemonic:"), stringWithAmpersand)
}
// Configures the textual and background attributes of the receiver’s field editor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("setUpFieldEditorAttributes:"), textObj)
	return rv
}
// Begins tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint unsafe.Pointer, controlView unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}
// Stops tracking mouse events within the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint unsafe.Pointer, stopPoint unsafe.Pointer, controlView unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}
// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeDoubleValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeFloatValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeIntValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to an integer value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeIntegerValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the object value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeObjectValueFrom:"), sender)
}
// Sets the value of the receiver’s cell to the string value obtained from the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("takeStringValueFrom:"), sender)
}
// Returns the rectangle in which the receiver draws its title text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("titleRectForBounds:"), rect)
	return rv
}
// Initiates the mouse tracking behavior in a cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event unsafe.Pointer, cellFrame unsafe.Pointer, controlView unsafe.Pointer, flag bool) bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}
// A Boolean value indicating whether the cell accepts first responder status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/acceptsFirstResponder
func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("acceptsFirstResponder"))
	return rv
}
// The action performed by the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/action
func (c_ Cell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID(), objc.RegisterName("action"))
	return rv
}
// SetAction sets the value of the action property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/action
func (c_ Cell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAction:"), value)
}
// The alignment of the cell’s text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/alignment
func (c_ Cell) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("alignment"))
	return rv
}
// SetAlignment sets the value of the alignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/alignment
func (c_ Cell) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAlignment:"), value)
}
// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsEditingTextAttributes"))
	return rv
}
// SetAllowsEditingTextAttributes sets the value of the allowsEditingTextAttributes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsEditingTextAttributes:"), value)
}
// A Boolean value indicating whether the cell supports three states instead of two. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) AllowsMixedState() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsMixedState"))
	return rv
}
// SetAllowsMixedState sets the value of the allowsMixedState property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsMixedState:"), value)
}
// A Boolean value indicating whether the cell assumes responsibility for undo operations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) AllowsUndo() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsUndo"))
	return rv
}
// SetAllowsUndo sets the value of the allowsUndo property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsUndo:"), value)
}
// The cell’s value as an attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) AttributedStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("attributedStringValue"))
	return rv
}
// SetAttributedStringValue sets the value of the attributedStringValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) SetAttributedStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAttributedStringValue:"), value)
}
// The cell’s background style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) BackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("backgroundStyle"))
	return rv
}
// SetBackgroundStyle sets the value of the backgroundStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) SetBackgroundStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBackgroundStyle:"), value)
}
// The initial writing direction used to determine the actual writing direction for text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("baseWritingDirection"))
	return rv
}
// SetBaseWritingDirection sets the value of the baseWritingDirection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBaseWritingDirection:"), value)
}
// The minimum size needed to display the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/cellSize
func (c_ Cell) CellSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("cellSize"))
	return rv
}
// The size of the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlSize
func (c_ Cell) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("controlSize"))
	return rv
}
// SetControlSize sets the value of the controlSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlSize
func (c_ Cell) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setControlSize:"), value)
}
// The cell’s control tint. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlTint
func (c_ Cell) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("controlTint"))
	return rv
}
// SetControlTint sets the value of the controlTint property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlTint
func (c_ Cell) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setControlTint:"), value)
}
// The view associated with the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlView
func (c_ Cell) ControlView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("controlView"))
	return rv
}
// SetControlView sets the value of the controlView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/controlView
func (c_ Cell) SetControlView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setControlView:"), value)
}
// The cell’s value as a double-precision floating-point number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID(), objc.RegisterName("doubleValue"))
	return rv
}
// SetDoubleValue sets the value of the doubleValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDoubleValue:"), value)
}
// The cell’s value as a single-precision floating-point number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/floatValue
func (c_ Cell) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID(), objc.RegisterName("floatValue"))
	return rv
}
// SetFloatValue sets the value of the floatValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/floatValue
func (c_ Cell) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFloatValue:"), value)
}
// The type of focus ring to use with the associated view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) FocusRingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("focusRingType"))
	return rv
}
// SetFocusRingType sets the value of the focusRingType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) SetFocusRingType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFocusRingType:"), value)
}
// The font that the cell uses to display text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/font
func (c_ Cell) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("font"))
	return rv
}
// SetFont sets the value of the font property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/font
func (c_ Cell) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFont:"), value)
}
// The cell’s formatter object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/formatter
func (c_ Cell) Formatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("formatter"))
	return rv
}
// SetFormatter sets the value of the formatter property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/formatter
func (c_ Cell) SetFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setFormatter:"), value)
}
// A Boolean value that indicates whether the cell has a valid object value. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/hasValidObjectValue
func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("hasValidObjectValue"))
	return rv
}
// The image displayed by the cell, if any. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/image
func (c_ Cell) Image() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("image"))
	return rv
}
// SetImage sets the value of the image property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/image
func (c_ Cell) SetImage(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setImage:"), value)
}
// A Boolean value indicating whether the cell supports the importation of images into its text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) ImportsGraphics() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("importsGraphics"))
	return rv
}
// SetImportsGraphics sets the value of the importsGraphics property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setImportsGraphics:"), value)
}
// The cell’s value as an integer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/intValue
func (c_ Cell) IntValue() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("intValue"))
	return rv
}
// SetIntValue sets the value of the intValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/intValue
func (c_ Cell) SetIntValue(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setIntValue:"), value)
}
// The cell’s value as an   type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/integerValue
func (c_ Cell) IntegerValue() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("integerValue"))
	return rv
}
// SetIntegerValue sets the value of the integerValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/integerValue
func (c_ Cell) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setIntegerValue:"), value)
}
// The cell’s interior background style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/interiorBackgroundStyle
func (c_ Cell) InteriorBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("interiorBackgroundStyle"))
	return rv
}
// A Boolean value indicating whether the cell has a bezeled border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) Bezeled() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("bezeled"))
	return rv
}
// SetBezeled sets the value of the bezeled property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) SetBezeled(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBezeled:"), value)
}
// A Boolean value indicating whether the cell draws itself outlined with a plain border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isBordered
func (c_ Cell) Bordered() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("bordered"))
	return rv
}
// SetBordered sets the value of the bordered property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isBordered
func (c_ Cell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBordered:"), value)
}
// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) Continuous() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("continuous"))
	return rv
}
// SetContinuous sets the value of the continuous property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setContinuous:"), value)
}
// A Boolean value indicating whether the cell is editable. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEditable
func (c_ Cell) Editable() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("editable"))
	return rv
}
// SetEditable sets the value of the editable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEditable
func (c_ Cell) SetEditable(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setEditable:"), value)
}
// A Boolean value indicating whether the cell is currently enabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) Enabled() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("enabled"))
	return rv
}
// SetEnabled sets the value of the enabled property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setEnabled:"), value)
}
// A Boolean value indicating whether the cell has a highlighted appearance. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) Highlighted() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("highlighted"))
	return rv
}
// SetHighlighted sets the value of the highlighted property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setHighlighted:"), value)
}
// A Boolean value indicating whether the cell is completely opaque. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isOpaque
func (c_ Cell) Opaque() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("opaque"))
	return rv
}
// A Boolean value indicating whether excess text scrolls past the cell’s bounds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) Scrollable() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("scrollable"))
	return rv
}
// SetScrollable sets the value of the scrollable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) SetScrollable(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setScrollable:"), value)
}
// A Boolean value indicating whether the cell’s text can be selected. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) Selectable() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("selectable"))
	return rv
}
// SetSelectable sets the value of the selectable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setSelectable:"), value)
}
// The key equivalent associated with clicking the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/keyEquivalent
func (c_ Cell) KeyEquivalent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("keyEquivalent"))
	return rv
}
// The line break mode to use when drawing text in the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("lineBreakMode"))
	return rv
}
// SetLineBreakMode sets the value of the lineBreakMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setLineBreakMode:"), value)
}
// The cell’s contextual menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/menu
func (c_ Cell) Menu() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("menu"))
	return rv
}
// SetMenu sets the value of the menu property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/menu
func (c_ Cell) SetMenu(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMenu:"), value)
}
// The modifier flags for the last (left) mouse-down event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/mouseDownFlags
func (c_ Cell) MouseDownFlags() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("mouseDownFlags"))
	return rv
}
// The cell’s next state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/nextState
func (c_ Cell) NextState() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("nextState"))
	return rv
}
// The cell’s value as an Objective-C object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/objectValue
func (c_ Cell) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("objectValue"))
	return rv
}
// SetObjectValue sets the value of the objectValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/objectValue
func (c_ Cell) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setObjectValue:"), value)
}
// A Boolean value indicating whether the cell refuses the first responder status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("refusesFirstResponder"))
	return rv
}
// SetRefusesFirstResponder sets the value of the refusesFirstResponder property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setRefusesFirstResponder:"), value)
}
// The object represented by the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/representedObject
func (c_ Cell) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("representedObject"))
	return rv
}
// SetRepresentedObject sets the value of the representedObject property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/representedObject
func (c_ Cell) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setRepresentedObject:"), value)
}
// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("sendsActionOnEndEditing"))
	return rv
}
// SetSendsActionOnEndEditing sets the value of the sendsActionOnEndEditing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setSendsActionOnEndEditing:"), value)
}
// A Boolean value indicating whether the cell provides a visual indication that it is the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("showsFirstResponder"))
	return rv
}
// SetShowsFirstResponder sets the value of the showsFirstResponder property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setShowsFirstResponder:"), value)
}
// The cell’s current state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/state
func (c_ Cell) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("state"))
	return rv
}
// SetState sets the value of the state property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/state
func (c_ Cell) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setState:"), value)
}
// The cell’s value as a string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/stringValue
func (c_ Cell) StringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("stringValue"))
	return rv
}
// SetStringValue sets the value of the stringValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/stringValue
func (c_ Cell) SetStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setStringValue:"), value)
}
// A tag for identifying the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/tag
func (c_ Cell) Tag() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("tag"))
	return rv
}
// SetTag sets the value of the tag property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/tag
func (c_ Cell) SetTag(value int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTag:"), value)
}
// The object that receives the cell’s action messages. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/target
func (c_ Cell) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID(), objc.RegisterName("target"))
	return rv
}
// SetTarget sets the value of the target property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/target
func (c_ Cell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTarget:"), value)
}
// The cell’s title text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/title
func (c_ Cell) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("title"))
	return rv
}
// SetTitle sets the value of the title property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/title
func (c_ Cell) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTitle:"), value)
}
// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("truncatesLastVisibleLine"))
	return rv
}
// SetTruncatesLastVisibleLine sets the value of the truncatesLastVisibleLine property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setTruncatesLastVisibleLine:"), value)
}
// The type of the cell. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/type
func (c_ Cell) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("type"))
	return rv
}
// SetType sets the value of the type property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/type
func (c_ Cell) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setType:"), value)
}
// The layout direction of the user interface. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) UserInterfaceLayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("userInterfaceLayoutDirection"))
	return rv
}
// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) SetUserInterfaceLayoutDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setUserInterfaceLayoutDirection:"), value)
}
// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("usesSingleLineMode"))
	return rv
}
// SetUsesSingleLineMode sets the value of the usesSingleLineMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setUsesSingleLineMode:"), value)
}
// A Boolean value indicating whether the cell’s field editor should post text change notifications. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/wantsNotificationForMarkedText
func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("wantsNotificationForMarkedText"))
	return rv
}
// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/wraps
func (c_ Cell) Wraps() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("wraps"))
	return rv
}
// SetWraps sets the value of the wraps property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCell/wraps
func (c_ Cell) SetWraps(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setWraps:"), value)
}
