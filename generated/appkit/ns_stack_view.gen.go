// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStackView */


/* debug [class_header]: Header for NSStackView */
// The class instance for the [StackView] class.
var (
	StackViewClass     _StackViewClass
	StackViewClassOnce sync.Once
)

func getStackViewClass() _StackViewClass {
	StackViewClassOnce.Do(func() {
		StackViewClass = _StackViewClass{objc.GetClass("NSStackView")}
	})
	return StackViewClass
}

type _StackViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StackView */
// An interface definition for the [StackView] class.
type IStackView interface {
	IView
	
/* debug [class_interface_properties]: Properties for StackView */
	// properties:
	Alignment() LayoutAttribute
	SetAlignment(value LayoutAttribute)
	ArrangedSubviews() []View
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DetachedViews() []View
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
	Distribution() StackViewDistribution
	SetDistribution(value StackViewDistribution)
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(value foundation.EdgeInsets)
	HasEqualSpacing() bool
	SetHasEqualSpacing(value bool)
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(value UserInterfaceLayoutOrientation)
	Spacing() float64
	SetSpacing(value float64)
	Views() []View
	IsHidden() bool
	SetIsHidden(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StackView */
	// methods:
	AddArrangedSubview(view IView)
	AddViewInGravity(view IView, gravity StackViewGravity)
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */
	CustomSpacingAfterView(view IView) float64
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */
	InsertArrangedSubviewAtIndex(view IView, index int)
	InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity)
	RemoveArrangedSubview(view IView)
	RemoveView(view IView)
	SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation)
	SetCustomSpacingAfterView(spacing float64, view IView)
	SetHuggingPriorityForOrientation(huggingPriority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation)
	SetViewsInGravity(views []View, gravity StackViewGravity)
	SetVisibilityPriorityForView(priority StackViewVisibilityPriority /* typedef */, view IView)
	ViewsInGravity(gravity StackViewGravity) []View
	VisibilityPriorityForView(view IView) StackViewVisibilityPriority /* typedef */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StackView */
// Alloc allocates a new instance without initialization.
func (sc _StackViewClass) Alloc() StackView {
	rv := objc.Send[StackView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StackViewClass) New() StackView {
	rv := objc.Send[StackView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StackView) Init() StackView {
	rv := objc.Send[StackView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StackView) Autorelease() StackView {
	rv := objc.Send[StackView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStackView creates a new StackView instance.
func NewStackView() StackView {
	return getStackViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StackView */
// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes.
//
// A stack view employs Auto Layout (the system’s constraint-based layout feature) to arrange and align an array of views according to your specification. To use a stack view effectively, you need to understand the basics of Auto Layout constraints as described in .


// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView
type StackView struct {
	View
}

// StackViewFrom constructs a [StackView] from an unsafe.Pointer.
//
// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes.
func StackViewFrom(ptr unsafe.Pointer) StackView {
	return StackView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StackView */

// Creates and returns a stack view with a specified array of views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func NewStackViewWithViews(views []View) StackView {
	rv := objc.Send[StackView](objc.ID(getStackViewClass().class), objc.Sel("stackViewWithViews:"), views)
	return rv
}/* debug [class_init_methods/constructor]: NewStackViewWithViews */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StackView */

// Creates and returns a stack view with a specified array of views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func (sc _StackViewClass) StackViewWithViews(views []View) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stackViewWithViews:"), views)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StackViewWithViews) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StackView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StackView */

// Adds the specified view to the end of the arranged subviews list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s_ StackView) AddArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}/* debug [instance_methods/method]: AddArrangedSubview */


// Adds a view to the end of the stack view gravity area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addView(_:in:)
func (s_ StackView) AddViewInGravity(view IView, gravity StackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addView:inGravity:"), view, gravity)
}/* debug [instance_methods/method]: AddViewInGravity */


// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ClippingResistancePriorityForOrientation */


// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/customSpacing(after:)
func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("customSpacingAfterView:"), view)
	return rv
}/* debug [instance_methods/method]: CustomSpacingAfterView */


// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/huggingPriority(for:)
func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("huggingPriorityForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: HuggingPriorityForOrientation */


// Adds the provided view to the array of arranged subviews at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
func (s_ StackView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}/* debug [instance_methods/method]: InsertArrangedSubviewAtIndex */


// Adds a view to a stack view gravity area at a specified index position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertView(_:at:in:)
func (s_ StackView) InsertViewAtIndexInGravity(view IView, index uint, gravity StackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertView:atIndex:inGravity:"), view, index, gravity)
}/* debug [instance_methods/method]: InsertViewAtIndexInGravity */


// Removes the provided view from the stack’s array of arranged subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}/* debug [instance_methods/method]: RemoveArrangedSubview */


// Removes a specified view from the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeView(_:)
func (s_ StackView) RemoveView(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeView:"), view)
}/* debug [instance_methods/method]: RemoveView */


// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}/* debug [instance_methods/method]: SetClippingResistancePriorityForOrientation */


// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomSpacing:afterView:"), spacing, view)
}/* debug [instance_methods/method]: SetCustomSpacingAfterView */


// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}/* debug [instance_methods/method]: SetHuggingPriorityForOrientation */


// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setViews(_:in:)
func (s_ StackView) SetViewsInGravity(views []View, gravity StackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViews:inGravity:"), views, gravity)
}/* debug [instance_methods/method]: SetViewsInGravity */


// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
func (s_ StackView) SetVisibilityPriorityForView(priority StackViewVisibilityPriority /* typedef */, view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisibilityPriority:forView:"), priority, view)
}/* debug [instance_methods/method]: SetVisibilityPriorityForView */


// Returns the array of views in the specified gravity area in the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/views(in:)
func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("viewsInGravity:"), gravity)
	return rv
}/* debug [instance_methods/method]: ViewsInGravity */


// Returns the visibility priority for a specified view in the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority /* typedef */ {
	rv := objc.Send[float32](s_.ID, objc.Sel("visibilityPriorityForView:"), view)
	return rv
}/* debug [instance_methods/method]: VisibilityPriorityForView */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StackView */

// The view alignment within the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/alignment
func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.Send[LayoutAttribute](s_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The view alignment within the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/alignment
func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// The array of views arranged by the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/arrangedSubviews
func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("arrangedSubviews"))
	return rv
}/* debug [instance_properties/getter]: arrangedSubviews */


// The delegate object for the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/delegate
func (s_ StackView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object for the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/delegate
func (s_ StackView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// An array that contains the detached views from all the stack view’s gravity areas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachedViews
func (s_ StackView) DetachedViews() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("detachedViews"))
	return rv
}/* debug [instance_properties/getter]: detachedViews */


// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("detachesHiddenViews"))
	return rv
}/* debug [instance_properties/getter]: detachesHiddenViews */


// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDetachesHiddenViews:"), value)
}/* debug [instance_properties/setter]: detachesHiddenViews */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.Send[StackViewDistribution](s_.ID, objc.Sel("distribution"))
	return rv
}/* debug [instance_properties/getter]: distribution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) SetDistribution(value StackViewDistribution) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDistribution:"), value)
}/* debug [instance_properties/setter]: distribution */


// The geometric padding, in points, inside the stack view, surrounding its views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](s_.ID, objc.Sel("edgeInsets"))
	return rv
}/* debug [instance_properties/getter]: edgeInsets */


// The geometric padding, in points, inside the stack view, surrounding its views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEdgeInsets:"), value)
}/* debug [instance_properties/setter]: edgeInsets */


// A Boolean value that indicates whether the spacing between adjacent views should be equal to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) HasEqualSpacing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasEqualSpacing"))
	return rv
}/* debug [instance_properties/getter]: hasEqualSpacing */


// A Boolean value that indicates whether the spacing between adjacent views should be equal to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) SetHasEqualSpacing(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasEqualSpacing:"), value)
}/* debug [instance_properties/setter]: hasEqualSpacing */


// The horizontal or vertical layout direction of the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/orientation
func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.Send[UserInterfaceLayoutOrientation](s_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// The horizontal or vertical layout direction of the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/orientation
func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// The minimum spacing, in points, between adjacent views in the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/spacing
func (s_ StackView) Spacing() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("spacing"))
	return rv
}/* debug [instance_properties/getter]: spacing */


// The minimum spacing, in points, between adjacent views in the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/spacing
func (s_ StackView) SetSpacing(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpacing:"), value)
}/* debug [instance_properties/setter]: spacing */


// The array of views owned by the stack view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/views
func (s_ StackView) Views() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("views"))
	return rv
}/* debug [instance_properties/getter]: views */


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (s_ StackView) IsHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (s_ StackView) SetIsHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/userinterfacelayoutdirection
func (s_ StackView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](s_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/userinterfacelayoutdirection
func (s_ StackView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStackView */


