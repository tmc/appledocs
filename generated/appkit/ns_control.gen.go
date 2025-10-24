// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSControl */


/* debug [class_header]: Header for NSControl */
// The class instance for the [Control] class.
var (
	ControlClass     _ControlClass
	ControlClassOnce sync.Once
)

func getControlClass() _ControlClass {
	ControlClassOnce.Do(func() {
		ControlClass = _ControlClass{objc.GetClass("NSControl")}
	})
	return ControlClass
}

type _ControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Control */
// An interface definition for the [Control] class.
type IControl interface {
	IView
	
/* debug [class_interface_properties]: Properties for Control */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.AttributedString)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	Cell() ICell
	SetCell(value ICell)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	Font() IFont
	SetFont(value IFont)
	Formatter() objectivec.IObject
	SetFormatter(value objectivec.IObject)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
	IntValue() int
	SetIntValue(value int)
	IntegerValue() int
	SetIntegerValue(value int)
	Continuous() bool
	SetContinuous(value bool)
	Enabled() bool
	SetEnabled(value bool)
	Highlighted() bool
	SetHighlighted(value bool)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	ObjectValue() objc.ID
	SetObjectValue(value objc.ID)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	IsContinuous() bool
	SetIsContinuous(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Control */
	// methods:
	AbortEditing() bool
	CurrentEditor() IText
	DrawWithExpansionFrameInView(contentFrame Rect /* not a class type */, view IView)
	DrawCell(cell ICell)
	DrawCellInside(cell ICell)
	EditWithFrameEditorDelegateEvent(rect Rect /* not a class type */, textObj IText, delegate objc.IObject, event IEvent)
	EndEditing(textObj IText)
	ExpansionFrameWithFrame(contentFrame Rect /* not a class type */) Rect /* not a class type */
	InvalidateIntrinsicContentSizeForCell(cell ICell)
	PerformClick(sender objc.IObject)
	SelectWithFrameEditorDelegateStartLength(rect Rect /* not a class type */, textObj IText, delegate objc.IObject, selStart int, selLength int)
	SelectCell(cell ICell)
	SelectedCell() ICell
	SelectedTag() int
	SendActionTo(action objc.SEL, target objc.IObject) bool
	SendActionOn(mask EventMask) int
	SizeThatFits(size Size /* not a class type */) Size /* not a class type */
	SizeToFit()
	TakeDoubleValueFrom(sender objc.IObject)
	TakeFloatValueFrom(sender objc.IObject)
	TakeIntValueFrom(sender objc.IObject)
	TakeIntegerValueFrom(sender objc.IObject)
	TakeObjectValueFrom(sender objc.IObject)
	TakeStringValueFrom(sender objc.IObject)
	UpdateCell(cell ICell)
	UpdateCellInside(cell ICell)
	ValidateEditing()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Control */
// Alloc allocates a new instance without initialization.
func (cc _ControlClass) Alloc() Control {
	rv := objc.Send[Control](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Control */
// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
//
// The class is abstract and must be subclassed to be used. Although you can subclass it yourself, more often you use one of the subclasses already defined by AppKit. A control draws content on the screen, automatically handles user interactions with that content, and calls the action method of its target object for any significant user interactions.


// A specialized view, such as a button or text field, that notifies your app of relevant events using the target-action design pattern.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Control */

// Initializes a control with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewControlWithCoder(coder foundation.Coder) Control {
	instance := getControlClass().Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewControlWithCoder */


// Initializes a control with the specified frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewControlWithFrame(frameRect Rect /* not a class type */) Control {
	instance := getControlClass().Alloc()
	rv := objc.Send[Control](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewControlWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Control */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Control */

// Returns the type of cell used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/cellClass
func (cc _ControlClass) CellClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(cc.class), objc.Sel("cellClass"))
	return rv
}/* debug [class_properties_class/property]: cellClass */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Control */

// Terminates the current editing operation and discards any edited text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/abortEditing()
func (c_ Control) AbortEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("abortEditing"))
	return rv
}/* debug [instance_methods/method]: AbortEditing */


// Returns the current field editor for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/currentEditor()
func (c_ Control) CurrentEditor() IText {
	rv := objc.Send[Text](c_.ID, objc.Sel("currentEditor"))
	return rv
}/* debug [instance_methods/method]: CurrentEditor */


// Performs custom expansion tool tip drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c_ Control) DrawWithExpansionFrameInView(contentFrame Rect /* not a class type */, view IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawWithExpansionFrame:inView:"), contentFrame, view)
}/* debug [instance_methods/method]: DrawWithExpansionFrameInView */


// Draws the specified cell, as long as it belongs to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/drawCell(_:)
func (c_ Control) DrawCell(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawCell:"), cell)
}/* debug [instance_methods/method]: DrawCell */


// Draws the inside of the receiver’s cell (the area within the bezel or border)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/drawCellInside(_:)
func (c_ Control) DrawCellInside(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawCellInside:"), cell)
}/* debug [instance_methods/method]: DrawCellInside */


// Begins editing of the receiver’s text using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c_ Control) EditWithFrameEditorDelegateEvent(rect Rect /* not a class type */, textObj IText, delegate objc.IObject, event IEvent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("editWithFrame:editor:delegate:event:"), rect, textObj, delegate, event)
}/* debug [instance_methods/method]: EditWithFrameEditorDelegateEvent */


// Ends the editing of text in the receiver using the specified field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/endEditing(_:)
func (c_ Control) EndEditing(textObj IText) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endEditing:"), textObj)
}/* debug [instance_methods/method]: EndEditing */


// The frame in which a tool tip can be displayed, if needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c_ Control) ExpansionFrameWithFrame(contentFrame Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](c_.ID, objc.Sel("expansionFrameWithFrame:"), contentFrame)
	return rv
}/* debug [instance_methods/method]: ExpansionFrameWithFrame */


// Notifies the control that the intrinsic content size for its cell is no longer valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidateIntrinsicContentSizeForCell:"), cell)
}/* debug [instance_methods/method]: InvalidateIntrinsicContentSizeForCell */


// Simulates a single mouse click on the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/performClick(_:)
func (c_ Control) PerformClick(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performClick:"), sender)
}/* debug [instance_methods/method]: PerformClick */


// Selects the specified text range in the receiver’s field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c_ Control) SelectWithFrameEditorDelegateStartLength(rect Rect /* not a class type */, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectWithFrame:editor:delegate:start:length:"), rect, textObj, delegate, selStart, selLength)
}/* debug [instance_methods/method]: SelectWithFrameEditorDelegateStartLength */


// Selects the specified cell and redraws the control as needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/selectCell(_:)
func (c_ Control) SelectCell(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectCell:"), cell)
}/* debug [instance_methods/method]: SelectCell */


// Returns the receiver’s selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/selectedCell()
func (c_ Control) SelectedCell() ICell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("selectedCell"))
	return rv
}/* debug [instance_methods/method]: SelectedCell */


// Returns the tag of the receiver’s selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/selectedTag()
func (c_ Control) SelectedTag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("selectedTag"))
	return rv
}/* debug [instance_methods/method]: SelectedTag */


// Causes the specified action to be sent to the target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(_:to:)
func (c_ Control) SendActionTo(action objc.SEL, target objc.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("sendAction:to:"), action, target)
	return rv
}/* debug [instance_methods/method]: SendActionTo */


// Sets the conditions on which the receiver sends action messages to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(on:)
func (c_ Control) SendActionOn(mask EventMask) int {
	rv := objc.Send[int](c_.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}/* debug [instance_methods/method]: SendActionOn */


// Asks the control to calculate and return the size that best fits the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c_ Control) SizeThatFits(size Size /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](c_.ID, objc.Sel("sizeThatFits:"), size)
	return rv
}/* debug [instance_methods/method]: SizeThatFits */


// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/sizeToFit()
func (c_ Control) SizeToFit() {
	objc.Send[objc.ID](c_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */


// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeDoubleValueFrom */


// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeFloatValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeFloatValueFrom */


// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeIntValueFrom */


// Sets the value of the receiver’s cell to an value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeIntegerValueFrom */


// Sets the value of the receiver’s cell to the object value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeObjectValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeObjectValueFrom */


// Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("takeStringValueFrom:"), sender)
}/* debug [instance_methods/method]: TakeStringValueFrom */


// Marks the specified cell as in need of redrawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/updateCell(_:)
func (c_ Control) UpdateCell(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateCell:"), cell)
}/* debug [instance_methods/method]: UpdateCell */


// Marks the inside of the specified cell as in need of redrawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/updateCellInside(_:)
func (c_ Control) UpdateCellInside(cell ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateCellInside:"), cell)
}/* debug [instance_methods/method]: UpdateCellInside */


// Validates changes to any user-typed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/validateEditing()
func (c_ Control) ValidateEditing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("validateEditing"))
}/* debug [instance_methods/method]: ValidateEditing */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Control */

// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/action
func (c_ Control) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/action
func (c_ Control) SetAction(value objc.SEL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The alignment mode of the text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/alignment
func (c_ Control) Alignment() TextAlignment {
	rv := objc.Send[TextAlignment](c_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The alignment mode of the text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/alignment
func (c_ Control) SetAlignment(value TextAlignment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsExpansionToolTips"))
	return rv
}/* debug [instance_properties/getter]: allowsExpansionToolTips */


// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsExpansionToolTips:"), value)
}/* debug [instance_properties/setter]: allowsExpansionToolTips */


// The value of the receiver’s cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](c_.ID, objc.Sel("attributedStringValue"))
	return rv
}/* debug [instance_properties/getter]: attributedStringValue */


// The value of the receiver’s cell as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/attributedStringValue
func (c_ Control) SetAttributedStringValue(value foundation.AttributedString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttributedStringValue:"), value)
}/* debug [instance_properties/setter]: attributedStringValue */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := objc.Send[WritingDirection](c_.ID, objc.Sel("baseWritingDirection"))
	return rv
}/* debug [instance_properties/getter]: baseWritingDirection */


// The initial writing direction used to determine the actual writing direction for text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/baseWritingDirection
func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBaseWritingDirection:"), value)
}/* debug [instance_properties/setter]: baseWritingDirection */


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/cell
func (c_ Control) Cell() ICell {
	rv := objc.Send[Cell](c_.ID, objc.Sel("cell"))
	return rv
}/* debug [instance_properties/getter]: cell */


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/cell
func (c_ Control) SetCell(value ICell) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCell:"), value)
}/* debug [instance_properties/setter]: cell */


// Returns the type of cell used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/cellClass
func (c_ Control) CellClass() objc.Class {
	rv := objc.Send[objc.Class](c_.ID, objc.Sel("cellClass"))
	return rv
}/* debug [instance_properties/getter]: cellClass */


// Returns the type of cell used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/cellClass
func (c_ Control) SetCellClass(value objc.Class) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCellClass:"), value)
}/* debug [instance_properties/setter]: cellClass */


// The size of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](c_.ID, objc.Sel("controlSize"))
	return rv
}/* debug [instance_properties/getter]: controlSize */


// The size of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/controlSize-swift.property
func (c_ Control) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlSize:"), value)
}/* debug [instance_properties/setter]: controlSize */


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/doubleValue
func (c_ Control) DoubleValue() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The value of the receiver’s cell as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/doubleValue
func (c_ Control) SetDoubleValue(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDoubleValue:"), value)
}/* debug [instance_properties/setter]: doubleValue */


// The value of the receiver’s cell as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/floatValue
func (c_ Control) FloatValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("floatValue"))
	return rv
}/* debug [instance_properties/getter]: floatValue */


// The value of the receiver’s cell as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/floatValue
func (c_ Control) SetFloatValue(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFloatValue:"), value)
}/* debug [instance_properties/setter]: floatValue */


// The font used to draw text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/font
func (c_ Control) Font() IFont {
	rv := objc.Send[Font](c_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font used to draw text in the receiver’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/font
func (c_ Control) SetFont(value IFont) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The receiver’s formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/formatter
func (c_ Control) Formatter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("formatter"))
	return rv
}/* debug [instance_properties/getter]: formatter */


// The receiver’s formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/formatter
func (c_ Control) SetFormatter(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatter:"), value)
}/* debug [instance_properties/setter]: formatter */


// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("ignoresMultiClick"))
	return rv
}/* debug [instance_properties/getter]: ignoresMultiClick */


// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ignoresMultiClick
func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIgnoresMultiClick:"), value)
}/* debug [instance_properties/setter]: ignoresMultiClick */


// The value of the receiver’s cell as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/intValue
func (c_ Control) IntValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("intValue"))
	return rv
}/* debug [instance_properties/getter]: intValue */


// The value of the receiver’s cell as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/intValue
func (c_ Control) SetIntValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntValue:"), value)
}/* debug [instance_properties/setter]: intValue */


// The value of the receiver’s cell as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/integerValue
func (c_ Control) IntegerValue() int {
	rv := objc.Send[int](c_.ID, objc.Sel("integerValue"))
	return rv
}/* debug [instance_properties/getter]: integerValue */


// The value of the receiver’s cell as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/integerValue
func (c_ Control) SetIntegerValue(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntegerValue:"), value)
}/* debug [instance_properties/setter]: integerValue */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isContinuous
func (c_ Control) Continuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuous"))
	return rv
}/* debug [instance_properties/getter]: continuous */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isContinuous
func (c_ Control) SetContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContinuous:"), value)
}/* debug [instance_properties/setter]: continuous */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isEnabled
func (c_ Control) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isEnabled
func (c_ Control) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) Highlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highlighted"))
	return rv
}/* debug [instance_properties/getter]: highlighted */


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/isHighlighted
func (c_ Control) SetHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighlighted:"), value)
}/* debug [instance_properties/setter]: highlighted */


// The line break mode to use for text in the control’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) LineBreakMode() LineBreakMode {
	rv := objc.Send[LineBreakMode](c_.ID, objc.Sel("lineBreakMode"))
	return rv
}/* debug [instance_properties/getter]: lineBreakMode */


// The line break mode to use for text in the control’s cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/lineBreakMode
func (c_ Control) SetLineBreakMode(value LineBreakMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLineBreakMode:"), value)
}/* debug [instance_properties/setter]: lineBreakMode */


// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/objectValue
func (c_ Control) ObjectValue() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("objectValue"))
	return rv
}/* debug [instance_properties/getter]: objectValue */


// The value of the receiver’s cell as an Objective-C object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/objectValue
func (c_ Control) SetObjectValue(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObjectValue:"), value)
}/* debug [instance_properties/setter]: objectValue */


// A Boolean value indicating whether the receiver refuses the first responder role.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("refusesFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: refusesFirstResponder */


// A Boolean value indicating whether the receiver refuses the first responder role.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/refusesFirstResponder
func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRefusesFirstResponder:"), value)
}/* debug [instance_properties/setter]: refusesFirstResponder */


// The value of the receiver’s cell as an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/stringValue
func (c_ Control) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The value of the receiver’s cell as an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/stringValue
func (c_ Control) SetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStringValue:"), value)
}/* debug [instance_properties/setter]: stringValue */


// The tag identifying the receiver (not the tag of the receiver’s cell).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/tag
func (c_ Control) Tag() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// The tag identifying the receiver (not the tag of the receiver’s cell).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/tag
func (c_ Control) SetTag(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/target
func (c_ Control) Target() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/target
func (c_ Control) SetTarget(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesSingleLineMode"))
	return rv
}/* debug [instance_properties/getter]: usesSingleLineMode */


// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/usesSingleLineMode
func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesSingleLineMode:"), value)
}/* debug [instance_properties/setter]: usesSingleLineMode */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (c_ Control) IsContinuous() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (c_ Control) SetIsContinuous(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (c_ Control) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/isenabled
func (c_ Control) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ishighlighted
func (c_ Control) IsHighlighted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighlighted"))
	return rv
}/* debug [instance_properties/getter]: isHighlighted */


// A Boolean value that indicates whether the cell is highlighted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/ishighlighted
func (c_ Control) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighlighted:"), value)
}/* debug [instance_properties/setter]: isHighlighted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSControl */


