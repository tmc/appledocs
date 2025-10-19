// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StackView] class.
var (
	stackViewClass     _StackViewClass
	stackViewClassOnce sync.Once
)

func getStackViewClass() _StackViewClass {
	stackViewClassOnce.Do(func() {
		stackViewClass = _StackViewClass{objc.GetClass("NSStackView")}
	})
	return stackViewClass
}

type _StackViewClass struct {
	class objc.Class
}

// An interface definition for the [StackView] class.
type IStackView interface {
	IView
	AddArrangedSubview(view unsafe.Pointer)
	AddViewInGravity(view unsafe.Pointer, gravity unsafe.Pointer)
	ClippingResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	CustomSpacingAfterView(view unsafe.Pointer) float64
	HuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int)
	InsertViewAtIndexInGravity(view unsafe.Pointer, index uint, gravity unsafe.Pointer)
	RemoveArrangedSubview(view unsafe.Pointer)
	RemoveView(view unsafe.Pointer)
	SetClippingResistancePriorityForOrientation(clippingResistancePriority unsafe.Pointer, orientation unsafe.Pointer)
	SetCustomSpacingAfterView(spacing float64, view unsafe.Pointer)
	SetHuggingPriorityForOrientation(huggingPriority unsafe.Pointer, orientation unsafe.Pointer)
	SetViewsInGravity(views unsafe.Pointer, gravity unsafe.Pointer)
	SetVisibilityPriorityForView(priority unsafe.Pointer, view unsafe.Pointer)
	ViewsInGravity(gravity unsafe.Pointer) unsafe.Pointer
	VisibilityPriorityForView(view unsafe.Pointer) unsafe.Pointer
}

// A view that arranges an array of views horizontally or vertically and updates their placement and sizing when the window size changes. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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


// Creates and returns a stack view with a specified array of views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func NewStackViewWithViews(views unsafe.Pointer) StackView {
	rv := objc.Send[StackView](objc.ID(getStackViewClass().class), objc.Sel("stackViewWithViews:"), views)
	rv.Autorelease()
	return rv
}


// Creates and returns a stack view with a specified array of views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/init(views:)
func (sc _StackViewClass) StackViewWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stackViewWithViews:"), views)
	return rv
}
// Adds the specified view to the end of the arranged subviews list. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s_ StackView) AddArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}
// Adds a view to the end of the stack view gravity area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/addView(_:in:)
func (s_ StackView) AddViewInGravity(view unsafe.Pointer, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addView:inGravity:"), view, gravity)
}
// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}
// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/customSpacing(after:)
func (s_ StackView) CustomSpacingAfterView(view unsafe.Pointer) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("customSpacingAfterView:"), view)
	return rv
}
// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/huggingPriority(for:)
func (s_ StackView) HuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("huggingPriorityForOrientation:"), orientation)
	return rv
}
// Adds the provided view to the array of arranged subviews at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
func (s_ StackView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}
// Adds a view to a stack view gravity area at a specified index position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/insertView(_:at:in:)
func (s_ StackView) InsertViewAtIndexInGravity(view unsafe.Pointer, index uint, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertView:atIndex:inGravity:"), view, index, gravity)
}
// Removes the provided view from the stack’s array of arranged subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s_ StackView) RemoveArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}
// Removes a specified view from the stack view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/removeView(_:)
func (s_ StackView) RemoveView(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeView:"), view)
}
// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}
// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomSpacing:afterView:"), spacing, view)
}
// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}
// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setViews(_:in:)
func (s_ StackView) SetViewsInGravity(views unsafe.Pointer, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setViews:inGravity:"), views, gravity)
}
// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
func (s_ StackView) SetVisibilityPriorityForView(priority unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisibilityPriority:forView:"), priority, view)
}
// Returns the array of views in the specified gravity area in the stack view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/views(in:)
func (s_ StackView) ViewsInGravity(gravity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("viewsInGravity:"), gravity)
	return rv
}
// Returns the visibility priority for a specified view in the stack view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s_ StackView) VisibilityPriorityForView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("visibilityPriorityForView:"), view)
	return rv
}

