
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StackView] class.
var StackViewClass _StackViewClass

func init() {
	StackViewClass = _StackViewClass{objc.GetClass("NSStackView")}
}

type _StackViewClass struct {
	objc.Class
}

// An interface definition for the [StackView] class.
type IStackView interface {
	ID() objc.ID
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

type StackView struct {
	id objc.ID
}

func StackViewFrom(ptr unsafe.Pointer) StackView {
	return StackView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ StackView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _StackViewClass) Alloc() StackView {
	rv := objc.Send[StackView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _StackViewClass) New() StackView {
	rv := objc.Send[StackView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewStackView creates and returns a new initialized instance.
func NewStackView() StackView {
	return StackViewClass.New()
}

// Init initializes the instance.
func (s_ StackView) Init() StackView {
	rv := objc.Send[StackView](s_.ID(), selInit)
	return rv
}
// Creates and returns a stack view with a specified array of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/init(views:)
func (sc _StackViewClass) StackViewWithViews(views unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("stackViewWithViews:"), views)
	return rv
}

// StackView_StackViewWithViews creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/init(views:)
func StackView_StackViewWithViews(views unsafe.Pointer) unsafe.Pointer {
	return StackViewClass.StackViewWithViews(views)
}
// Adds the specified view to the end of the arranged subviews list. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s_ StackView) AddArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("addArrangedSubview:"), view)
}
// Adds a view to the end of the stack view gravity area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/addView(_:in:)
func (s_ StackView) AddViewInGravity(view unsafe.Pointer, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("addView:inGravity:"), view, gravity)
}
// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}
// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/customSpacing(after:)
func (s_ StackView) CustomSpacingAfterView(view unsafe.Pointer) float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("customSpacingAfterView:"), view)
	return rv
}
// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/huggingPriority(for:)
func (s_ StackView) HuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("huggingPriorityForOrientation:"), orientation)
	return rv
}
// Adds the provided view to the array of arranged subviews at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
func (s_ StackView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("insertArrangedSubview:atIndex:"), view, index)
}
// Adds a view to a stack view gravity area at a specified index position. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/insertView(_:at:in:)
func (s_ StackView) InsertViewAtIndexInGravity(view unsafe.Pointer, index uint, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("insertView:atIndex:inGravity:"), view, index, gravity)
}
// Removes the provided view from the stack’s array of arranged subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s_ StackView) RemoveArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("removeArrangedSubview:"), view)
}
// Removes a specified view from the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/removeView(_:)
func (s_ StackView) RemoveView(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("removeView:"), view)
}
// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}
// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setCustomSpacing:afterView:"), spacing, view)
}
// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}
// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setViews(_:in:)
func (s_ StackView) SetViewsInGravity(views unsafe.Pointer, gravity unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setViews:inGravity:"), views, gravity)
}
// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
func (s_ StackView) SetVisibilityPriorityForView(priority unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVisibilityPriority:forView:"), priority, view)
}
// Returns the array of views in the specified gravity area in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/views(in:)
func (s_ StackView) ViewsInGravity(gravity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("viewsInGravity:"), gravity)
	return rv
}
// Returns the visibility priority for a specified view in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s_ StackView) VisibilityPriorityForView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("visibilityPriorityForView:"), view)
	return rv
}
// The view alignment within the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/alignment
func (s_ StackView) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("alignment"))
	return rv
}
// SetAlignment sets the value of the alignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/alignment
func (s_ StackView) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAlignment:"), value)
}
// The array of views arranged by the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/arrangedSubviews
func (s_ StackView) ArrangedSubviews() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("arrangedSubviews"))
	return rv
}
// The delegate object for the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/delegate
func (s_ StackView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/delegate
func (s_ StackView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDelegate:"), value)
}
// An array that contains the detached views from all the stack view’s gravity areas. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/detachedViews
func (s_ StackView) DetachedViews() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("detachedViews"))
	return rv
}
// A Boolean value that indicates whether the stack view removes hidden views from its view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("detachesHiddenViews"))
	return rv
}
// SetDetachesHiddenViews sets the value of the detachesHiddenViews property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/detachesHiddenViews
func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDetachesHiddenViews:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) Distribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("distribution"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/distribution-swift.property
func (s_ StackView) SetDistribution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDistribution:"), value)
}
// The geometric padding, in points, inside the stack view, surrounding its views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) EdgeInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("edgeInsets"))
	return rv
}
// SetEdgeInsets sets the value of the edgeInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/edgeInsets
func (s_ StackView) SetEdgeInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setEdgeInsets:"), value)
}
// A Boolean value that indicates whether the spacing between adjacent views should be equal to each other. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) HasEqualSpacing() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("hasEqualSpacing"))
	return rv
}
// SetHasEqualSpacing sets the value of the hasEqualSpacing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/hasEqualSpacing
func (s_ StackView) SetHasEqualSpacing(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHasEqualSpacing:"), value)
}
// The horizontal or vertical layout direction of the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/orientation
func (s_ StackView) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("orientation"))
	return rv
}
// SetOrientation sets the value of the orientation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/orientation
func (s_ StackView) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setOrientation:"), value)
}
// The minimum spacing, in points, between adjacent views in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/spacing
func (s_ StackView) Spacing() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("spacing"))
	return rv
}
// SetSpacing sets the value of the spacing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/spacing
func (s_ StackView) SetSpacing(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSpacing:"), value)
}
// The array of views owned by the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/views
func (s_ StackView) Views() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("views"))
	return rv
}
