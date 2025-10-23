// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cell] class.
var (
	CellClass     _CellClass
	CellClassOnce sync.Once
)

func getCellClass() _CellClass {
	CellClassOnce.Do(func() {
		CellClass = _CellClass{objc.GetClass("NSCell")}
	})
	return CellClass
}

type _CellClass struct {
	class objc.Class
}

// An interface definition for the [Cell] class.
type ICell interface {
	objectivec.IObject
	AcceptsFirstResponder() bool
	Action() objc.SEL
	SetAction(value objc.SEL)
	Alignment() unsafe.Pointer
	SetAlignment(value unsafe.Pointer)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	BackgroundStyle() NSBackgroundStyle
	SetBackgroundStyle(value NSBackgroundStyle)
	BaseWritingDirection() unsafe.Pointer
	SetBaseWritingDirection(value unsafe.Pointer)
	CellSize() coregraphics.CGSize
	ControlSize() unsafe.Pointer
	SetControlSize(value unsafe.Pointer)
	ControlTint() unsafe.Pointer
	SetControlTint(value unsafe.Pointer)
	ControlView() IView
	SetControlView(value IView)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	FocusRingType() NSFocusRingType
	SetFocusRingType(value NSFocusRingType)
	Font() IFont
	SetFont(value IFont)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.Formatter)
	HasValidObjectValue() bool
	Image() IImage
	SetImage(value IImage)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	IntValue() int
	SetIntValue(value int)
	IntegerValue() int
	SetIntegerValue(value int)
	InteriorBackgroundStyle() NSBackgroundStyle
	Bezeled() bool
	SetBezeled(value bool)
	Bordered() bool
	SetBordered(value bool)
	Continuous() bool
	SetContinuous(value bool)
	Editable() bool
	SetEditable(value bool)
	Enabled() bool
	SetEnabled(value bool)
	Highlighted() bool
	SetHighlighted(value bool)
	Opaque() bool
	Scrollable() bool
	SetScrollable(value bool)
	Selectable() bool
	SetSelectable(value bool)
	KeyEquivalent() string
	LineBreakMode() unsafe.Pointer
	SetLineBreakMode(value unsafe.Pointer)
	Menu() IMenu
	SetMenu(value IMenu)
	MouseDownFlags() int
	NextState() int
	ObjectValue() objc.ID
	SetObjectValue(value objc.ID)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	RepresentedObject() objc.ID
	SetRepresentedObject(value objc.ID)
	SendsActionOnEndEditing() bool
	SetSendsActionOnEndEditing(value bool)
	ShowsFirstResponder() bool
	SetShowsFirstResponder(value bool)
	State() unsafe.Pointer
	SetState(value unsafe.Pointer)
	StringValue() string
	SetStringValue(value string)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() string
	SetTitle(value string)
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	Type() NSCellType
	SetType(value NSCellType)
	UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	WantsNotificationForMarkedText() bool
	Wraps() bool
	SetWraps(value bool)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsBordered() bool
	SetIsBordered(value bool)
	IsContinuous() bool
	SetIsContinuous(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsOpaque() bool
	SetIsOpaque(value bool)
	IsScrollable() bool
	SetIsScrollable(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	CalcDrawInfo(rect coregraphics.CGRect)
	CellAttribute(parameter NSCellAttribute) int
	CellSizeForBounds(rect coregraphics.CGRect) coregraphics.CGSize
	Compare(otherCell objectivec.IObject) unsafe.Pointer
	ContinueTrackingAtInView(lastPoint coregraphics.CGPoint, currentPoint coregraphics.CGPoint, controlView IView) bool
	DraggingImageComponentsWithFrameInView(frame coregraphics.CGRect, view IView) []DraggingImageComponent
	DrawWithExpansionFrameInView(cellFrame coregraphics.CGRect, view IView)
	DrawWithFrameInView(cellFrame coregraphics.CGRect, controlView IView)
	DrawFocusRingMaskWithFrameInView(cellFrame coregraphics.CGRect, controlView IView)
	DrawInteriorWithFrameInView(cellFrame coregraphics.CGRect, controlView IView)
	DrawingRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	EditWithFrameInViewEditorDelegateEvent(rect coregraphics.CGRect, controlView IView, textObj IText, delegate objectivec.IObject, event IEvent)
	EndEditing(textObj IText)
	ExpansionFrameWithFrameInView(cellFrame coregraphics.CGRect, view IView) coregraphics.CGRect
	FieldEditorForView(controlView IView) ITextView
	FocusRingMaskBoundsForFrameInView(cellFrame coregraphics.CGRect, controlView IView) coregraphics.CGRect
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)
	HighlightWithFrameInView(flag bool, cellFrame coregraphics.CGRect, controlView IView)
	HighlightColorWithFrameInView(cellFrame coregraphics.CGRect, controlView IView) IColor
	HitTestForEventInRectOfView(event IEvent, cellFrame coregraphics.CGRect, controlView IView) NSCellHitResult
	ImageRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	MenuForEventInRectOfView(event IEvent, cellFrame coregraphics.CGRect, view IView) IMenu
	PerformClick(sender objectivec.IObject)
	ResetCursorRectInView(cellFrame coregraphics.CGRect, controlView IView)
	SelectWithFrameInViewEditorDelegateStartLength(rect coregraphics.CGRect, controlView IView, textObj IText, delegate objectivec.IObject, selStart int, selLength int)
	SendActionOn(mask NSEventMask) int
	SetCellAttributeTo(parameter NSCellAttribute, value int)
	SetUpFieldEditorAttributes(textObj IText) IText
	StartTrackingAtInView(startPoint coregraphics.CGPoint, controlView IView) bool
	StopTrackingAtInViewMouseIsUp(lastPoint coregraphics.CGPoint, stopPoint coregraphics.CGPoint, controlView IView, flag bool)
	TakeDoubleValueFrom(sender objectivec.IObject)
	TakeFloatValueFrom(sender objectivec.IObject)
	TakeIntValueFrom(sender objectivec.IObject)
	TakeIntegerValueFrom(sender objectivec.IObject)
	TakeObjectValueFrom(sender objectivec.IObject)
	TakeStringValueFrom(sender objectivec.IObject)
	TitleRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect
	TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame coregraphics.CGRect, controlView IView, flag bool) bool
}

// A mechanism for displaying text or images in a view object without the overhead of a full subclass.
//
// Cells are used by most of the classes to implement their internal workings.


// A mechanism for displaying text or images in a view object without the overhead of a full subclass.
//
// [Full Topic]
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



// Returns an object initialized with the specified image and set to have the cell’s default menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewCellImageCell(image IImage) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}


// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewCellTextCell(string_ string) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder foundation.Coder) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns the default type of focus ring for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultFocusRingType
func (cc _CellClass) DefaultFocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](objc.ID(cc.class), objc.Sel("defaultFocusRingType"))
	return rv
}

// Returns the default menu for instances of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultMenu
func (cc _CellClass) DefaultMenu() IMenu {
	rv := objc.Send[Menu](objc.ID(cc.class), objc.Sel("defaultMenu"))
	return rv
}

// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/prefersTrackingUntilMouseUp
func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}

// Recalculates the cell geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calcDrawInfo:"), rect)
}


// Returns the value for the specified cell attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter NSCellAttribute) int {
	rv := objc.Send[int](c_.ID, objc.Sel("cellAttribute:"), parameter)
	return rv
}


// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect coregraphics.CGRect) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("cellSizeForBounds:"), rect)
	return rv
}


// Compares the string values of the receiver another cell, disregarding case.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compare:"), otherCell)
	return rv
}


// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint coregraphics.CGPoint, currentPoint coregraphics.CGPoint, controlView IView) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}


// Generates dragging image components with the specified frame in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame coregraphics.CGRect, view IView) []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](c_.ID, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}


// Instructs the receiver to draw in an expansion frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame coregraphics.CGRect, view IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, view)
}


// Draws the receiver’s border and then draws the interior of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame coregraphics.CGRect, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}


// Draws the focus ring for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame coregraphics.CGRect, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}


// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame coregraphics.CGRect, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}


// Returns the rectangle within which the receiver draws itself
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("drawingRectForBounds:"), rect)
	return rv
}


// Begins editing of the receiver’s text using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect coregraphics.CGRect, controlView IView, textObj IText, delegate objectivec.IObject, event IEvent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}


// Ends the editing of text in the receiver using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj IText) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}


// Returns the expansion cell frame for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame coregraphics.CGRect, view IView) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}


// Returns a custom field editor for editing in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView IView) ITextView {
	rv := objc.Send[TextView](c_.ID, objc.Sel("fieldEditorForView:"), controlView)
	return rv
}


// Returns the bounds of the focus ring mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame coregraphics.CGRect, controlView IView) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}


// Returns the initial delay and repeat values for continuous sending of action messages to target objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}


// Redraws the receiver with the specified highlight setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame coregraphics.CGRect, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}


// Returns the color the receiver uses when drawing the selection highlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame coregraphics.CGRect, controlView IView) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}


// Returns hit testing information for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event IEvent, cellFrame coregraphics.CGRect, controlView IView) NSCellHitResult {
	rv := objc.Send[NSCellHitResult](c_.ID, objc.Sel("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}


// Returns the rectangle in which the receiver draws its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("imageRectForBounds:"), rect)
	return rv
}


// Returns the menu associated with the cell and related to the specified event and frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event IEvent, cellFrame coregraphics.CGRect, view IView) IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}


// Simulates a single mouse click on the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}


// Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame coregraphics.CGRect, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetCursorRect:inView:"), cellFrame, controlView)
}


// Selects the specified text range in the cell’s field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect coregraphics.CGRect, controlView IView, textObj IText, delegate objectivec.IObject, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}


// Sets the conditions on which the receiver sends action messages to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask NSEventMask) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}


// Sets the value for the specified cell attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter NSCellAttribute, value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellAttribute:to:"), parameter, value)
}


// Configures the textual and background attributes of the receiver’s field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) IText {
	rv := objc.Send[Text](c_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}


// Begins tracking mouse events within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint coregraphics.CGPoint, controlView IView) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}


// Stops tracking mouse events within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint coregraphics.CGPoint, stopPoint coregraphics.CGPoint, controlView IView, flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}


// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}


// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}


// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}


// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}


// Sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}


// Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}


// Returns the rectangle in which the receiver draws its title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("titleRectForBounds:"), rect)
	return rv
}


// Initiates the mouse tracking behavior in a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame coregraphics.CGRect, controlView IView, flag bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}


// A Boolean value indicating whether the cell accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/acceptsFirstResponder
func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}


// The action performed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}


// The action performed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}


// The alignment of the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alignment"))
	return rv
}


// The alignment of the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}


// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}


// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}


// A Boolean value indicating whether the cell supports three states instead of two.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) AllowsMixedState() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMixedState"))
	return rv
}


// A Boolean value indicating whether the cell supports three states instead of two.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMixedState:"), value)
}


// A Boolean value indicating whether the cell assumes responsibility for undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) AllowsUndo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsUndo"))
	return rv
}


// A Boolean value indicating whether the cell assumes responsibility for undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsUndo:"), value)
}


// The cell’s value as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}


// The cell’s value as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) SetAttributedStringValue(value foundation.AttributedString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}


// The cell’s background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) BackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](c_.ID, objc.Sel("backgroundStyle"))
	return rv
}


// The cell’s background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) SetBackgroundStyle(value NSBackgroundStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundStyle:"), value)
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) BaseWritingDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) SetBaseWritingDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}


// The minimum size needed to display the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize
func (c_ Cell) CellSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("cellSize"))
	return rv
}


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) ControlSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlSize"))
	return rv
}


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) SetControlSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}


// The cell’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) ControlTint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlTint"))
	return rv
}


// The cell’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) SetControlTint(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlTint:"), value)
}


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) ControlView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("controlView"))
	return rv
}


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) SetControlView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlView:"), value)
}


// Returns the default type of focus ring for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultFocusRingType
func (c_ Cell) DefaultFocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](c_.ID, objc.Sel("defaultFocusRingType"))
	return rv
}


// Returns the default menu for instances of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultMenu
func (c_ Cell) DefaultMenu() IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("defaultMenu"))
	return rv
}


// The cell’s value as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("doubleValue"))
	return rv
}


// The cell’s value as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("floatValue"))
	return rv
}


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}


// The type of focus ring to use with the associated view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) FocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](c_.ID, objc.Sel("focusRingType"))
	return rv
}


// The type of focus ring to use with the associated view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) SetFocusRingType(value NSFocusRingType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusRingType:"), value)
}


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) Font() IFont {
	rv := objc.Send[Font](c_.ID, objc.Sel("font"))
	return rv
}


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) SetFont(value IFont) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}


// The cell’s formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) Formatter() foundation.Formatter {
	rv := objc.Send[foundation.Formatter](c_.ID, objc.Sel("formatter"))
	return rv
}


// The cell’s formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) SetFormatter(value foundation.Formatter) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}


// A Boolean value that indicates whether the cell has a valid object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hasValidObjectValue
func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidObjectValue"))
	return rv
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) Image() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) SetImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
}


// A Boolean value indicating whether the cell supports the importation of images into its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) ImportsGraphics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value indicating whether the cell supports the importation of images into its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportsGraphics:"), value)
}


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) IntValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("intValue"))
	return rv
}


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) SetIntValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}


// The cell’s value as an type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}


// The cell’s value as an type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}


// The cell’s interior background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/interiorBackgroundStyle
func (c_ Cell) InteriorBackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](c_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) Bezeled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bezeled"))
	return rv
}


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) SetBezeled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBezeled:"), value)
}


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) Bordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bordered"))
	return rv
}


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBordered:"), value)
}


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) Editable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) SetEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditable:"), value)
}


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) Highlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highlighted"))
	return rv
}


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlighted:"), value)
}


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isOpaque
func (c_ Cell) Opaque() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("opaque"))
	return rv
}


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) Scrollable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("scrollable"))
	return rv
}


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) SetScrollable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollable:"), value)
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) Selectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("selectable"))
	return rv
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectable:"), value)
}


// The key equivalent associated with clicking the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/keyEquivalent
func (c_ Cell) KeyEquivalent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("keyEquivalent"))
	return rv
}


// The line break mode to use when drawing text in the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) LineBreakMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}


// The line break mode to use when drawing text in the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) SetLineBreakMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}


// The cell’s contextual menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) Menu() IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("menu"))
	return rv
}


// The cell’s contextual menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) SetMenu(value IMenu) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMenu:"), value)
}


// The modifier flags for the last (left) mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mouseDownFlags
func (c_ Cell) MouseDownFlags() int {
	rv := objc.Send[int](c_.ID, objc.Sel("mouseDownFlags"))
	return rv
}


// The cell’s next state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/nextState
func (c_ Cell) NextState() int {
	rv := objc.Send[int](c_.ID, objc.Sel("nextState"))
	return rv
}


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValue"))
	return rv
}


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}


// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/prefersTrackingUntilMouseUp
func (c_ Cell) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}


// A Boolean value indicating whether the cell refuses the first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}


// A Boolean value indicating whether the cell refuses the first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}


// The object represented by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("representedObject"))
	return rv
}


// The object represented by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRepresentedObject:"), value)
}


// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendsActionOnEndEditing"))
	return rv
}


// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSendsActionOnEndEditing:"), value)
}


// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsFirstResponder"))
	return rv
}


// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsFirstResponder:"), value)
}


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("state"))
	return rv
}


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) StringValue() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stringValue"))
	return rv
}


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) SetStringValue(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), objc.String(value))
}


// A tag for identifying the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}


// A tag for identifying the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}


// The object that receives the cell’s action messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}


// The object that receives the cell’s action messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("truncatesLastVisibleLine"))
	return rv
}


// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTruncatesLastVisibleLine:"), value)
}


// The type of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) Type() NSCellType {
	rv := objc.Send[NSCellType](c_.ID, objc.Sel("type"))
	return rv
}


// The type of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) SetType(value NSCellType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](c_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}


// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}


// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}


// A Boolean value indicating whether the cell’s field editor should post text change notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wantsNotificationForMarkedText
func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wantsNotificationForMarkedText"))
	return rv
}


// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) Wraps() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wraps"))
	return rv
}


// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) SetWraps(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWraps:"), value)
}


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbezeled
func (c_ Cell) IsBezeled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBezeled"))
	return rv
}


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbezeled
func (c_ Cell) SetIsBezeled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBezeled:"), value)
}


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbordered
func (c_ Cell) IsBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBordered"))
	return rv
}


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbordered
func (c_ Cell) SetIsBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBordered:"), value)
}


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iscontinuous
func (c_ Cell) IsContinuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iscontinuous
func (c_ Cell) SetIsContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuous:"), value)
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (c_ Cell) IsEditable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (c_ Cell) SetIsEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isenabled
func (c_ Cell) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isenabled
func (c_ Cell) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/ishighlighted
func (c_ Cell) IsHighlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/ishighlighted
func (c_ Cell) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighlighted:"), value)
}


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isopaque
func (c_ Cell) IsOpaque() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOpaque"))
	return rv
}


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isopaque
func (c_ Cell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOpaque:"), value)
}


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isscrollable
func (c_ Cell) IsScrollable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isScrollable"))
	return rv
}


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isscrollable
func (c_ Cell) SetIsScrollable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsScrollable:"), value)
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (c_ Cell) IsSelectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (c_ Cell) SetIsSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelectable:"), value)
}


