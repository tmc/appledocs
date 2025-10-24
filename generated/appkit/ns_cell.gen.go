// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSCell */


/* debug [class_header]: Header for NSCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Cell */
// An interface definition for the [Cell] class.
type ICell interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Cell */
	// properties:
	AcceptsFirstResponder() bool
	Action() objc.SEL
	SetAction(value objc.SEL)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	CellSize() Size /* not a class type */
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	ControlView() IView
	SetControlView(value IView)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	Font() IFont
	SetFont(value IFont)
	Formatter() objectivec.IObject
	SetFormatter(value objectivec.IObject)
	HasValidObjectValue() bool
	Image() IImage
	SetImage(value IImage)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	IntValue() int
	SetIntValue(value int)
	IntegerValue() int
	SetIntegerValue(value int)
	InteriorBackgroundStyle() BackgroundStyle
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
	KeyEquivalent() objc.IObject /* cross-framework: NSString */
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
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
	State() ControlStateValue /* typedef */
	SetState(value ControlStateValue /* typedef */)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	Type() CellType
	SetType(value CellType)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Cell */
	// methods:
	CalcDrawInfo(rect Rect /* not a class type */)
	CellAttribute(parameter CellAttribute) int
	CellSizeForBounds(rect Rect /* not a class type */) Size /* not a class type */
	Compare(otherCell objc.IObject) ComparisonResult /* not a class type */
	ContinueTrackingAtInView(lastPoint vision.Point, currentPoint vision.Point, controlView IView) bool
	DraggingImageComponentsWithFrameInView(frame Rect /* not a class type */, view IView) []DraggingImageComponent
	DrawWithExpansionFrameInView(cellFrame Rect /* not a class type */, view IView)
	DrawWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawFocusRingMaskWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawInteriorWithFrameInView(cellFrame Rect /* not a class type */, controlView IView)
	DrawingRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */
	EditWithFrameInViewEditorDelegateEvent(rect Rect /* not a class type */, controlView IView, textObj IText, delegate objc.IObject, event IEvent)
	EndEditing(textObj IText)
	ExpansionFrameWithFrameInView(cellFrame Rect /* not a class type */, view IView) Rect /* not a class type */
	FieldEditorForView(controlView IView) ITextView
	FocusRingMaskBoundsForFrameInView(cellFrame Rect /* not a class type */, controlView IView) Rect /* not a class type */
	GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject)
	HighlightWithFrameInView(flag bool, cellFrame Rect /* not a class type */, controlView IView)
	HighlightColorWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) IColor
	HitTestForEventInRectOfView(event IEvent, cellFrame Rect /* not a class type */, controlView IView) CellHitResult
	ImageRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */
	MenuForEventInRectOfView(event IEvent, cellFrame Rect /* not a class type */, view IView) IMenu
	PerformClick(sender objc.IObject)
	ResetCursorRectInView(cellFrame Rect /* not a class type */, controlView IView)
	SelectWithFrameInViewEditorDelegateStartLength(rect Rect /* not a class type */, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int)
	SendActionOn(mask EventMask) int
	SetCellAttributeTo(parameter CellAttribute, value int)
	SetUpFieldEditorAttributes(textObj IText) IText
	StartTrackingAtInView(startPoint vision.Point, controlView IView) bool
	StopTrackingAtInViewMouseIsUp(lastPoint vision.Point, stopPoint vision.Point, controlView IView, flag bool)
	TakeDoubleValueFrom(sender objc.IObject)
	TakeFloatValueFrom(sender objc.IObject)
	TakeIntValueFrom(sender objc.IObject)
	TakeIntegerValueFrom(sender objc.IObject)
	TakeObjectValueFrom(sender objc.IObject)
	TakeStringValueFrom(sender objc.IObject)
	TitleRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */
	TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame Rect /* not a class type */, controlView IView, flag bool) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Cell */
// Alloc allocates a new instance without initialization.
func (cc _CellClass) Alloc() Cell {
	rv := objc.Send[Cell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Cell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Cell */

// Returns an object initialized with the specified image and set to have the cell’s default menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewCellImageCell(image IImage) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initImageCell:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCellImageCell */


// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewCellTextCell(string_ objc.IObject /* cross-framework: NSString */) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCellTextCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder foundation.Coder) Cell {
	instance := getCellClass().Alloc()
	rv := objc.Send[Cell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Cell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Cell */

// Returns the default type of focus ring for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultFocusRingType
func (cc _CellClass) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](objc.ID(cc.class), objc.Sel("defaultFocusRingType"))
	return rv
}/* debug [class_properties_class/property]: defaultFocusRingType */

// Returns the default menu for instances of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultMenu
func (cc _CellClass) DefaultMenu() IMenu {
	rv := objc.Send[Menu](objc.ID(cc.class), objc.Sel("defaultMenu"))
	return rv
}/* debug [class_properties_class/property]: defaultMenu */

// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/prefersTrackingUntilMouseUp
func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}/* debug [class_properties_class/property]: prefersTrackingUntilMouseUp */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Cell */

// Recalculates the cell geometry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c_ Cell) CalcDrawInfo(rect Rect /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calcDrawInfo:"), rect)
}/* debug [instance_methods/method]: CalcDrawInfo */


// Returns the value for the specified cell attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellAttribute(_:)
func (c_ Cell) CellAttribute(parameter CellAttribute) int {
	rv := objc.Send[int](c_.ID, objc.Sel("cellAttribute:"), parameter)
	return rv
}/* debug [instance_methods/method]: CellAttribute */


// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c_ Cell) CellSizeForBounds(rect Rect /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("cellSizeForBounds:"), rect)
	return rv
}/* debug [instance_methods/method]: CellSizeForBounds */


// Compares the string values of the receiver another cell, disregarding case.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/compare(_:)
func (c_ Cell) Compare(otherCell objc.IObject) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](c_.ID, objc.Sel("compare:"), otherCell)
	return rv
}/* debug [instance_methods/method]: Compare */


// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c_ Cell) ContinueTrackingAtInView(lastPoint vision.Point, currentPoint vision.Point, controlView IView) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}/* debug [instance_methods/method]: ContinueTrackingAtInView */


// Generates dragging image components with the specified frame in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c_ Cell) DraggingImageComponentsWithFrameInView(frame Rect /* not a class type */, view IView) []DraggingImageComponent {
	rv := objc.Send[[]DraggingImageComponent](c_.ID, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}/* debug [instance_methods/method]: DraggingImageComponentsWithFrameInView */


// Instructs the receiver to draw in an expansion frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c_ Cell) DrawWithExpansionFrameInView(cellFrame Rect /* not a class type */, view IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, view)
}/* debug [instance_methods/method]: DrawWithExpansionFrameInView */


// Draws the receiver’s border and then draws the interior of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c_ Cell) DrawWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawWithFrameInView */


// Draws the focus ring for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c_ Cell) DrawFocusRingMaskWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawFocusRingMaskWithFrameInView */


// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c_ Cell) DrawInteriorWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: DrawInteriorWithFrameInView */


// Returns the rectangle within which the receiver draws itself
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c_ Cell) DrawingRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("drawingRectForBounds:"), rect)
	return rv
}/* debug [instance_methods/method]: DrawingRectForBounds */


// Begins editing of the receiver’s text using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c_ Cell) EditWithFrameInViewEditorDelegateEvent(rect Rect /* not a class type */, controlView IView, textObj IText, delegate objc.IObject, event IEvent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}/* debug [instance_methods/method]: EditWithFrameInViewEditorDelegateEvent */


// Ends the editing of text in the receiver using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/endEditing(_:)
func (c_ Cell) EndEditing(textObj IText) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}/* debug [instance_methods/method]: EndEditing */


// Returns the expansion cell frame for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c_ Cell) ExpansionFrameWithFrameInView(cellFrame Rect /* not a class type */, view IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}/* debug [instance_methods/method]: ExpansionFrameWithFrameInView */


// Returns a custom field editor for editing in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/fieldEditor(for:)
func (c_ Cell) FieldEditorForView(controlView IView) ITextView {
	rv := objc.Send[TextView](c_.ID, objc.Sel("fieldEditorForView:"), controlView)
	return rv
}/* debug [instance_methods/method]: FieldEditorForView */


// Returns the bounds of the focus ring mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c_ Cell) FocusRingMaskBoundsForFrameInView(cellFrame Rect /* not a class type */, controlView IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}/* debug [instance_methods/method]: FocusRingMaskBoundsForFrameInView */


// Returns the initial delay and repeat values for continuous sending of action messages to target objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c_ Cell) GetPeriodicDelayInterval(delay objectivec.IObject, interval objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}/* debug [instance_methods/method]: GetPeriodicDelayInterval */


// Redraws the receiver with the specified highlight setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c_ Cell) HighlightWithFrameInView(flag bool, cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}/* debug [instance_methods/method]: HighlightWithFrameInView */


// Returns the color the receiver uses when drawing the selection highlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c_ Cell) HighlightColorWithFrameInView(cellFrame Rect /* not a class type */, controlView IView) IColor {
	rv := objc.Send[Color](c_.ID, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}/* debug [instance_methods/method]: HighlightColorWithFrameInView */


// Returns hit testing information for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c_ Cell) HitTestForEventInRectOfView(event IEvent, cellFrame Rect /* not a class type */, controlView IView) CellHitResult {
	rv := objc.Send[CellHitResult](c_.ID, objc.Sel("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}/* debug [instance_methods/method]: HitTestForEventInRectOfView */


// Returns the rectangle in which the receiver draws its image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c_ Cell) ImageRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("imageRectForBounds:"), rect)
	return rv
}/* debug [instance_methods/method]: ImageRectForBounds */


// Returns the menu associated with the cell and related to the specified event and frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu(for:in:of:)
func (c_ Cell) MenuForEventInRectOfView(event IEvent, cellFrame Rect /* not a class type */, view IView) IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}/* debug [instance_methods/method]: MenuForEventInRectOfView */


// Simulates a single mouse click on the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/performClick(_:)
func (c_ Cell) PerformClick(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}/* debug [instance_methods/method]: PerformClick */


// Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c_ Cell) ResetCursorRectInView(cellFrame Rect /* not a class type */, controlView IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetCursorRect:inView:"), cellFrame, controlView)
}/* debug [instance_methods/method]: ResetCursorRectInView */


// Selects the specified text range in the cell’s field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c_ Cell) SelectWithFrameInViewEditorDelegateStartLength(rect Rect /* not a class type */, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}/* debug [instance_methods/method]: SelectWithFrameInViewEditorDelegateStartLength */


// Sets the conditions on which the receiver sends action messages to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendAction(on:)
func (c_ Cell) SendActionOn(mask EventMask) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}/* debug [instance_methods/method]: SendActionOn */


// Sets the value for the specified cell attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c_ Cell) SetCellAttributeTo(parameter CellAttribute, value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellAttribute:to:"), parameter, value)
}/* debug [instance_methods/method]: SetCellAttributeTo */


// Configures the textual and background attributes of the receiver’s field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) IText {
	rv := objc.Send[Text](c_.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return rv
}/* debug [instance_methods/method]: SetUpFieldEditorAttributes */


// Begins tracking mouse events within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/startTracking(at:in:)
func (c_ Cell) StartTrackingAtInView(startPoint vision.Point, controlView IView) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}/* debug [instance_methods/method]: StartTrackingAtInView */


// Stops tracking mouse events within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c_ Cell) StopTrackingAtInViewMouseIsUp(lastPoint vision.Point, stopPoint vision.Point, controlView IView, flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}/* debug [instance_methods/method]: StopTrackingAtInViewMouseIsUp */


// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c_ Cell) TakeDoubleValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeDoubleValueFrom */


// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c_ Cell) TakeFloatValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeFloatValueFrom */


// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c_ Cell) TakeIntValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeIntValueFrom */


// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c_ Cell) TakeIntegerValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeIntegerValueFrom */


// Sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c_ Cell) TakeObjectValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeObjectValueFrom */


// Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c_ Cell) TakeStringValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeStringValueFrom */


// Returns the rectangle in which the receiver draws its title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c_ Cell) TitleRectForBounds(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("titleRectForBounds:"), rect)
	return rv
}/* debug [instance_methods/method]: TitleRectForBounds */


// Initiates the mouse tracking behavior in a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c_ Cell) TrackMouseInRectOfViewUntilMouseUp(event IEvent, cellFrame Rect /* not a class type */, controlView IView, flag bool) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}/* debug [instance_methods/method]: TrackMouseInRectOfViewUntilMouseUp */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Cell */

// A Boolean value indicating whether the cell accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/acceptsFirstResponder
func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: acceptsFirstResponder */


// The action performed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The action performed by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c_ Cell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The alignment of the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](c_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The alignment of the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c_ Cell) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}/* debug [instance_properties/getter]: allowsEditingTextAttributes */


// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}/* debug [instance_properties/setter]: allowsEditingTextAttributes */


// A Boolean value indicating whether the cell supports three states instead of two.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) AllowsMixedState() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMixedState"))
	return rv
}/* debug [instance_properties/getter]: allowsMixedState */


// A Boolean value indicating whether the cell supports three states instead of two.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMixedState:"), value)
}/* debug [instance_properties/setter]: allowsMixedState */


// A Boolean value indicating whether the cell assumes responsibility for undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) AllowsUndo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsUndo"))
	return rv
}/* debug [instance_properties/getter]: allowsUndo */


// A Boolean value indicating whether the cell assumes responsibility for undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c_ Cell) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsUndo:"), value)
}/* debug [instance_properties/setter]: allowsUndo */


// The cell’s value as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}/* debug [instance_properties/getter]: attributedStringValue */


// The cell’s value as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c_ Cell) SetAttributedStringValue(value foundation.AttributedString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}/* debug [instance_properties/setter]: attributedStringValue */


// The cell’s background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) BackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](c_.ID, objc.Sel("backgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: backgroundStyle */


// The cell’s background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c_ Cell) SetBackgroundStyle(value BackgroundStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundStyle:"), value)
}/* debug [instance_properties/setter]: backgroundStyle */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}/* debug [instance_properties/getter]: baseWritingDirection */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c_ Cell) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}/* debug [instance_properties/setter]: baseWritingDirection */


// The minimum size needed to display the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/cellSize
func (c_ Cell) CellSize() Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("cellSize"))
	return rv
}/* debug [instance_properties/getter]: cellSize */


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](c_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */


// The size of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c_ Cell) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}/* debug [instance_properties/setter]: controlSize */


// The cell’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](c_.ID, objc.Sel("controlTint"))
	return rv
}/* debug [instance_properties/getter]: controlTint */


// The cell’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlTint
func (c_ Cell) SetControlTint(value ControlTint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlTint:"), value)
}/* debug [instance_properties/setter]: controlTint */


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) ControlView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("controlView"))
	return rv
}/* debug [instance_properties/getter]: controlView */


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c_ Cell) SetControlView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlView:"), value)
}/* debug [instance_properties/setter]: controlView */


// Returns the default type of focus ring for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultFocusRingType
func (c_ Cell) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](c_.ID, objc.Sel("defaultFocusRingType"))
	return rv
}/* debug [instance_properties/getter]: defaultFocusRingType */


// Returns the default menu for instances of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/defaultMenu
func (c_ Cell) DefaultMenu() IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("defaultMenu"))
	return rv
}/* debug [instance_properties/getter]: defaultMenu */


// The cell’s value as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The cell’s value as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c_ Cell) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}/* debug [instance_properties/setter]: doubleValue */


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("floatValue"))
	return rv
}/* debug [instance_properties/getter]: floatValue */


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c_ Cell) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}/* debug [instance_properties/setter]: floatValue */


// The type of focus ring to use with the associated view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) FocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](c_.ID, objc.Sel("focusRingType"))
	return rv
}/* debug [instance_properties/getter]: focusRingType */


// The type of focus ring to use with the associated view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c_ Cell) SetFocusRingType(value FocusRingType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusRingType:"), value)
}/* debug [instance_properties/setter]: focusRingType */


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) Font() IFont {
	rv := objc.Send[Font](c_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font that the cell uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c_ Cell) SetFont(value IFont) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The cell’s formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) Formatter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("formatter"))
	return rv
}/* debug [instance_properties/getter]: formatter */


// The cell’s formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c_ Cell) SetFormatter(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}/* debug [instance_properties/setter]: formatter */


// A Boolean value that indicates whether the cell has a valid object value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/hasValidObjectValue
func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidObjectValue"))
	return rv
}/* debug [instance_properties/getter]: hasValidObjectValue */


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) Image() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The image displayed by the cell, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c_ Cell) SetImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// A Boolean value indicating whether the cell supports the importation of images into its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) ImportsGraphics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("importsGraphics"))
	return rv
}/* debug [instance_properties/getter]: importsGraphics */


// A Boolean value indicating whether the cell supports the importation of images into its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c_ Cell) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportsGraphics:"), value)
}/* debug [instance_properties/setter]: importsGraphics */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) IntValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("intValue"))
	return rv
}/* debug [instance_properties/getter]: intValue */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c_ Cell) SetIntValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}/* debug [instance_properties/setter]: intValue */


// The cell’s value as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}/* debug [instance_properties/getter]: integerValue */


// The cell’s value as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c_ Cell) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}/* debug [instance_properties/setter]: integerValue */


// The cell’s interior background style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/interiorBackgroundStyle
func (c_ Cell) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Send[BackgroundStyle](c_.ID, objc.Sel("interiorBackgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: interiorBackgroundStyle */


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) Bezeled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bezeled"))
	return rv
}/* debug [instance_properties/getter]: bezeled */


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c_ Cell) SetBezeled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBezeled:"), value)
}/* debug [instance_properties/setter]: bezeled */


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) Bordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("bordered"))
	return rv
}/* debug [instance_properties/getter]: bordered */


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c_ Cell) SetBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBordered:"), value)
}/* debug [instance_properties/setter]: bordered */


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}/* debug [instance_properties/getter]: continuous */


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c_ Cell) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}/* debug [instance_properties/setter]: continuous */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) Editable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c_ Cell) SetEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c_ Cell) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) Highlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highlighted"))
	return rv
}/* debug [instance_properties/getter]: highlighted */


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c_ Cell) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlighted:"), value)
}/* debug [instance_properties/setter]: highlighted */


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isOpaque
func (c_ Cell) Opaque() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) Scrollable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("scrollable"))
	return rv
}/* debug [instance_properties/getter]: scrollable */


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c_ Cell) SetScrollable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScrollable:"), value)
}/* debug [instance_properties/setter]: scrollable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) Selectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("selectable"))
	return rv
}/* debug [instance_properties/getter]: selectable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c_ Cell) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectable:"), value)
}/* debug [instance_properties/setter]: selectable */


// The key equivalent associated with clicking the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/keyEquivalent
func (c_ Cell) KeyEquivalent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keyEquivalent"))
	return rv
}/* debug [instance_properties/getter]: keyEquivalent */


// The line break mode to use when drawing text in the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}/* debug [instance_properties/getter]: lineBreakMode */


// The line break mode to use when drawing text in the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c_ Cell) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}/* debug [instance_properties/setter]: lineBreakMode */


// The cell’s contextual menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) Menu() IMenu {
	rv := objc.Send[Menu](c_.ID, objc.Sel("menu"))
	return rv
}/* debug [instance_properties/getter]: menu */


// The cell’s contextual menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c_ Cell) SetMenu(value IMenu) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMenu:"), value)
}/* debug [instance_properties/setter]: menu */


// The modifier flags for the last (left) mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/mouseDownFlags
func (c_ Cell) MouseDownFlags() int {
	rv := objc.Send[int](c_.ID, objc.Sel("mouseDownFlags"))
	return rv
}/* debug [instance_properties/getter]: mouseDownFlags */


// The cell’s next state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/nextState
func (c_ Cell) NextState() int {
	rv := objc.Send[int](c_.ID, objc.Sel("nextState"))
	return rv
}/* debug [instance_properties/getter]: nextState */


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValue"))
	return rv
}/* debug [instance_properties/getter]: objectValue */


// The cell’s value as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c_ Cell) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}/* debug [instance_properties/setter]: objectValue */


// Returns a Boolean value that indicates whether tracking stops when the cursor leaves the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/prefersTrackingUntilMouseUp
func (c_ Cell) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}/* debug [instance_properties/getter]: prefersTrackingUntilMouseUp */


// A Boolean value indicating whether the cell refuses the first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: refusesFirstResponder */


// A Boolean value indicating whether the cell refuses the first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}/* debug [instance_properties/setter]: refusesFirstResponder */


// The object represented by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("representedObject"))
	return rv
}/* debug [instance_properties/getter]: representedObject */


// The object represented by the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c_ Cell) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRepresentedObject:"), value)
}/* debug [instance_properties/setter]: representedObject */


// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendsActionOnEndEditing"))
	return rv
}/* debug [instance_properties/getter]: sendsActionOnEndEditing */


// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSendsActionOnEndEditing:"), value)
}/* debug [instance_properties/setter]: sendsActionOnEndEditing */


// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("showsFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: showsFirstResponder */


// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShowsFirstResponder:"), value)
}/* debug [instance_properties/setter]: showsFirstResponder */


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) State() ControlStateValue /* typedef */ {
	rv := objc.Send[int](c_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The cell’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c_ Cell) SetState(value ControlStateValue /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c_ Cell) SetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), value)
}/* debug [instance_properties/setter]: stringValue */


// A tag for identifying the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// A tag for identifying the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c_ Cell) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */


// The object that receives the cell’s action messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The object that receives the cell’s action messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c_ Cell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The cell’s title text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c_ Cell) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("truncatesLastVisibleLine"))
	return rv
}/* debug [instance_properties/getter]: truncatesLastVisibleLine */


// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTruncatesLastVisibleLine:"), value)
}/* debug [instance_properties/setter]: truncatesLastVisibleLine */


// The type of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) Type() CellType {
	rv := objc.Send[CellType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The type of the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c_ Cell) SetType(value CellType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](c_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// The layout direction of the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c_ Cell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */


// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}/* debug [instance_properties/getter]: usesSingleLineMode */


// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}/* debug [instance_properties/setter]: usesSingleLineMode */


// A Boolean value indicating whether the cell’s field editor should post text change notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wantsNotificationForMarkedText
func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wantsNotificationForMarkedText"))
	return rv
}/* debug [instance_properties/getter]: wantsNotificationForMarkedText */


// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) Wraps() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("wraps"))
	return rv
}/* debug [instance_properties/getter]: wraps */


// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c_ Cell) SetWraps(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWraps:"), value)
}/* debug [instance_properties/setter]: wraps */


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbezeled
func (c_ Cell) IsBezeled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBezeled"))
	return rv
}/* debug [instance_properties/getter]: isBezeled */


// A Boolean value indicating whether the cell has a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbezeled
func (c_ Cell) SetIsBezeled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBezeled:"), value)
}/* debug [instance_properties/setter]: isBezeled */


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbordered
func (c_ Cell) IsBordered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBordered"))
	return rv
}/* debug [instance_properties/getter]: isBordered */


// A Boolean value indicating whether the cell draws itself outlined with a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isbordered
func (c_ Cell) SetIsBordered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBordered:"), value)
}/* debug [instance_properties/setter]: isBordered */


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iscontinuous
func (c_ Cell) IsContinuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iscontinuous
func (c_ Cell) SetIsContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (c_ Cell) IsEditable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value indicating whether the cell is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (c_ Cell) SetIsEditable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isenabled
func (c_ Cell) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value indicating whether the cell is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isenabled
func (c_ Cell) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/ishighlighted
func (c_ Cell) IsHighlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighlighted"))
	return rv
}/* debug [instance_properties/getter]: isHighlighted */


// A Boolean value indicating whether the cell has a highlighted appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/ishighlighted
func (c_ Cell) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighlighted:"), value)
}/* debug [instance_properties/setter]: isHighlighted */


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isopaque
func (c_ Cell) IsOpaque() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// A Boolean value indicating whether the cell is completely opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isopaque
func (c_ Cell) SetIsOpaque(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isscrollable
func (c_ Cell) IsScrollable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isScrollable"))
	return rv
}/* debug [instance_properties/getter]: isScrollable */


// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isscrollable
func (c_ Cell) SetIsScrollable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsScrollable:"), value)
}/* debug [instance_properties/setter]: isScrollable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (c_ Cell) IsSelectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelectable"))
	return rv
}/* debug [instance_properties/getter]: isSelectable */


// A Boolean value indicating whether the cell’s text can be selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (c_ Cell) SetIsSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelectable:"), value)
}/* debug [instance_properties/setter]: isSelectable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCell */


