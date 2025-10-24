// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSplitView */


/* debug [class_header]: Header for NSSplitView */
// The class instance for the [SplitView] class.
var (
	SplitViewClass     _SplitViewClass
	SplitViewClassOnce sync.Once
)

func getSplitViewClass() _SplitViewClass {
	SplitViewClassOnce.Do(func() {
		SplitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}
	})
	return SplitViewClass
}

type _SplitViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SplitView */
// An interface definition for the [SplitView] class.
type ISplitView interface {
	IView
	
/* debug [class_interface_properties]: Properties for SplitView */
	// properties:
	ArrangedSubviews() []View
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
	AutosaveName() SplitViewAutosaveName /* typedef */
	SetAutosaveName(value SplitViewAutosaveName /* typedef */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DividerColor() IColor
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(value SplitViewDividerStyle)
	DividerThickness() float64
	Vertical() bool
	SetVertical(value bool)
	IsVertical() bool
	SetIsVertical(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SplitView */
	// methods:
	AddArrangedSubview(view IView)
	AdjustSubviews()
	DrawDividerInRect(rect Rect /* not a class type */)
	HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority /* typedef */
	InsertArrangedSubviewAtIndex(view IView, index int)
	IsSubviewCollapsed(subview IView) bool
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	RemoveArrangedSubview(view IView)
	SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority /* typedef */, subviewIndex int)
	SetPositionOfDividerAtIndex(position float64, dividerIndex int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SplitView */
// Alloc allocates a new instance without initialization.
func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.Send[SplitView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SplitViewClass) New() SplitView {
	rv := objc.Send[SplitView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SplitView) Init() SplitView {
	rv := objc.Send[SplitView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SplitView) Autorelease() SplitView {
	rv := objc.Send[SplitView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSplitView creates a new SplitView instance.
func NewSplitView() SplitView {
	return getSplitViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SplitView */
// A view that arranges two or more views in a linear stack running horizontally or vertically.
//
// A split view manages the dividers and orientation for a split view controller ( ). By default, dividers have a horizontal orientation so that the split view arranges its panes vertically from top to bottom. Divider indices are zero-based. If the property is , which is the default value, the top divider has an index of . If is , the leading divider has an index of .


// A view that arranges two or more views in a linear stack running horizontally or vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView
type SplitView struct {
	View
}

// SplitViewFrom constructs a [SplitView] from an unsafe.Pointer.
//
// A view that arranges two or more views in a linear stack running horizontally or vertically.
func SplitViewFrom(ptr unsafe.Pointer) SplitView {
	return SplitView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SplitView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SplitView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SplitView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SplitView */

// Adds a view as an arranged split pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s_ SplitView) AddArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}/* debug [instance_methods/method]: AddArrangedSubview */


// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/adjustSubviews()
func (s_ SplitView) AdjustSubviews() {
	objc.Send[objc.ID](s_.ID, objc.Sel("adjustSubviews"))
}/* debug [instance_methods/method]: AdjustSubviews */


// Draws a divider between two of the split view’s subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s_ SplitView) DrawDividerInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawDividerInRect:"), rect)
}/* debug [instance_methods/method]: DrawDividerInRect */


// Returns the priority of the subview’s width or height when resizing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority /* typedef */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}/* debug [instance_methods/method]: HoldingPriorityForSubviewAtIndex */


// Adds a view as an arranged split pane at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s_ SplitView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}/* debug [instance_methods/method]: InsertArrangedSubviewAtIndex */


// Returns whether the specified view is in a collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSubviewCollapsed:"), subview)
	return rv
}/* debug [instance_methods/method]: IsSubviewCollapsed */


// Returns the maximum possible position of the divider at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}/* debug [instance_methods/method]: MaxPossiblePositionOfDividerAtIndex */


// Returns the minimum possible position of the divider at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}/* debug [instance_methods/method]: MinPossiblePositionOfDividerAtIndex */


// Removes a view as an arranged split pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s_ SplitView) RemoveArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}/* debug [instance_methods/method]: RemoveArrangedSubview */


// Sets the priority for split view subviews to maintain their width or height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority LayoutPriority /* typedef */, subviewIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}/* debug [instance_methods/method]: SetHoldingPriorityForSubviewAtIndex */


// Updates the location of a divider you specify by index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}/* debug [instance_methods/method]: SetPositionOfDividerAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SplitView */

// The array of views that the split view arranges as its split panes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangedSubviews
func (s_ SplitView) ArrangedSubviews() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("arrangedSubviews"))
	return rv
}/* debug [instance_properties/getter]: arrangedSubviews */


// A Boolean value that determines whether the split view arranges all of its subviews as split panes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("arrangesAllSubviews"))
	return rv
}/* debug [instance_properties/getter]: arrangesAllSubviews */


// A Boolean value that determines whether the split view arranges all of its subviews as split panes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArrangesAllSubviews:"), value)
}/* debug [instance_properties/setter]: arrangesAllSubviews */


// The name to use when the system automatically saves the split view’s divider configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) AutosaveName() SplitViewAutosaveName /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("autosaveName"))
	return rv
}/* debug [instance_properties/getter]: autosaveName */


// The name to use when the system automatically saves the split view’s divider configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutosaveName:"), value)
}/* debug [instance_properties/setter]: autosaveName */


// The split view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The split view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The color of the dividers that the split view draws between subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerColor
func (s_ SplitView) DividerColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("dividerColor"))
	return rv
}/* debug [instance_properties/getter]: dividerColor */


// The style of divider between views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := objc.Send[SplitViewDividerStyle](s_.ID, objc.Sel("dividerStyle"))
	return rv
}/* debug [instance_properties/getter]: dividerStyle */


// The style of divider between views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDividerStyle:"), value)
}/* debug [instance_properties/setter]: dividerStyle */


// The thickness of the dividers for the split view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerThickness
func (s_ SplitView) DividerThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("dividerThickness"))
	return rv
}/* debug [instance_properties/getter]: dividerThickness */


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}/* debug [instance_properties/getter]: vertical */


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}/* debug [instance_properties/setter]: vertical */


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitView) IsVertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isVertical"))
	return rv
}/* debug [instance_properties/getter]: isVertical */


// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitview/isvertical
func (s_ SplitView) SetIsVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsVertical:"), value)
}/* debug [instance_properties/setter]: isVertical */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSplitView */



