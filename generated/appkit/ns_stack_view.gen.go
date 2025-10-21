// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [StackView] class.
type IStackView interface {
	IView
	AddArrangedSubview(view IView)
	AddViewInGravity(view IView, gravity IStackViewGravity)
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	CustomSpacingAfterView(view IView) float64
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	InsertArrangedSubviewAtIndex(view IView, index int)
	InsertViewAtIndexInGravity(view IView, index uint, gravity IStackViewGravity)
	RemoveArrangedSubview(view IView)
	RemoveView(view IView)
	SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetCustomSpacingAfterView(spacing float64, view IView)
	SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetViewsInGravity(views []View, gravity IStackViewGravity)
	SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView)
	ViewsInGravity(gravity IStackViewGravity) []View
	VisibilityPriorityForView(view IView) StackViewVisibilityPriority
}

// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes.
//
// A stack view employs Auto Layout (the system’s constraint-based layout feature) to arrange and align an array of views according to your specification. To use a stack view effectively, you need to understand the basics of Auto Layout constraints as described in .
//
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

// Alloc allocates a new instance without initialization.
func (sc _StackViewClass) Alloc() StackView {
	rv := objc.Send[StackView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates and returns a stack view with a specified array of views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func NewStackViewWithViews(views []View) StackView {
	rv := objc.Send[StackView](objc.ID(getStackViewClass().class), objc.Sel("stackViewWithViews:"), views)
	return rv
}


// Creates and returns a stack view with a specified array of views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func (sc _StackViewClass) StackViewWithViews(views []View) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stackViewWithViews:"), views)
	return rv
}

// Adds the specified view to the end of the arranged subviews list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s_ StackView) AddArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}

// Adds a view to the end of the stack view gravity area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addView(_:in:)
func (s_ StackView) AddViewInGravity(view IView, gravity IStackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addView:inGravity:"), view, gravity)
}

// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Send[LayoutPriority](s_.ID, objc.Sel("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}

// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/customSpacing(after:)
func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("customSpacingAfterView:"), view)
	return rv
}

// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/huggingPriority(for:)
func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Send[LayoutPriority](s_.ID, objc.Sel("huggingPriorityForOrientation:"), orientation)
	return rv
}

// Adds the provided view to the array of arranged subviews at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
func (s_ StackView) InsertArrangedSubviewAtIndex(view IView, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}

// Adds a view to a stack view gravity area at a specified index position.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertView(_:at:in:)
func (s_ StackView) InsertViewAtIndexInGravity(view IView, index uint, gravity IStackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertView:atIndex:inGravity:"), view, index, gravity)
}

// Removes the provided view from the stack’s array of arranged subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}

// Removes a specified view from the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeView(_:)
func (s_ StackView) RemoveView(view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeView:"), view)
}

// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomSpacing:afterView:"), spacing, view)
}

// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setViews(_:in:)
func (s_ StackView) SetViewsInGravity(views []View, gravity IStackViewGravity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViews:inGravity:"), views, gravity)
}

// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
func (s_ StackView) SetVisibilityPriorityForView(priority StackViewVisibilityPriority, view IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisibilityPriority:forView:"), priority, view)
}

// Returns the array of views in the specified gravity area in the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/views(in:)
func (s_ StackView) ViewsInGravity(gravity IStackViewGravity) []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("viewsInGravity:"), gravity)
	return rv
}

// Returns the visibility priority for a specified view in the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := objc.Send[StackViewVisibilityPriority](s_.ID, objc.Sel("visibilityPriorityForView:"), view)
	return rv
}

// The view alignment within the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/alignment
func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.Send[LayoutAttribute](s_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The view alignment within the stack view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/alignment
func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlignment:"), value)
}

// The array of views arranged by the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/arrangedSubviews
func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("arrangedSubviews"))
	return rv
}

// The delegate object for the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/delegate
func (s_ StackView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object for the stack view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/delegate
func (s_ StackView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// An array that contains the detached views from all the stack view’s gravity areas.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachedViews
func (s_ StackView) DetachedViews() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("detachedViews"))
	return rv
}

// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("detachesHiddenViews"))
	return rv
}


// SetDetachesHiddenViews sets the value of the detachesHiddenViews property.
// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDetachesHiddenViews:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.Send[StackViewDistribution](s_.ID, objc.Sel("distribution"))
	return rv
}


// SetDistribution sets the value of the distribution property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) SetDistribution(value IStackViewDistribution) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDistribution:"), value)
}

// The geometric padding, in points, inside the stack view, surrounding its views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) EdgeInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("edgeInsets"))
	return rv
}


// SetEdgeInsets sets the value of the edgeInsets property.
// The geometric padding, in points, inside the stack view, surrounding its views.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) SetEdgeInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEdgeInsets:"), value)
}

// A Boolean value that indicates whether the spacing between adjacent views should be equal to each other.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) HasEqualSpacing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasEqualSpacing"))
	return rv
}


// SetHasEqualSpacing sets the value of the hasEqualSpacing property.
// A Boolean value that indicates whether the spacing between adjacent views should be equal to each other.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) SetHasEqualSpacing(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasEqualSpacing:"), value)
}

// The horizontal or vertical layout direction of the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/orientation
func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.Send[UserInterfaceLayoutOrientation](s_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// The horizontal or vertical layout direction of the stack view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/orientation
func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOrientation:"), value)
}

// The minimum spacing, in points, between adjacent views in the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/spacing
func (s_ StackView) Spacing() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("spacing"))
	return rv
}


// SetSpacing sets the value of the spacing property.
// The minimum spacing, in points, between adjacent views in the stack view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/spacing
func (s_ StackView) SetSpacing(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpacing:"), value)
}

// The array of views owned by the stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/views
func (s_ StackView) Views() []View {
	rv := objc.Send[[]View](s_.ID, objc.Sel("views"))
	return rv
}

// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (s_ StackView) IsHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value indicating whether the view is hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (s_ StackView) SetIsHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsHidden:"), value)
}

// The layout direction for content in the view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/userinterfacelayoutdirection
func (s_ StackView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](s_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// The layout direction for content in the view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/userinterfacelayoutdirection
func (s_ StackView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}


