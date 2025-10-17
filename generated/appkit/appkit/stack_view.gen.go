// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StackView] class.
var StackViewClass objc.Class

func init() {
	StackViewClass = objc.GetClass("NSStackView")
}

type StackView struct {
	objc.ID
}

func StackViewFrom(ptr unsafe.Pointer) StackView {
	return StackView{
		ID: objc.ID(ptr),
	}
}


// Creates and returns a stack view with a specified array of views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/init(views:)
func (sc StackView) StackViewWithViews(views unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("stackViewWithViews:")
	ret := objc.ID(StackViewClass).Send(sel, views)
	return unsafe.Pointer(ret)
}
// Adds the specified view to the end of the arranged subviews list. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/addArrangedSubview(_:)
func (s_ StackView) AddArrangedSubview(view unsafe.Pointer) {
	sel := objc.RegisterName("addArrangedSubview:")
	s_.ID.Send(sel, view)
}
// Adds a view to the end of the stack view gravity area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/addView(_:in:)
func (s_ StackView) AddViewInGravity(view unsafe.Pointer, gravity unsafe.Pointer) {
	sel := objc.RegisterName("addView:inGravity:")
	s_.ID.Send(sel, view, gravity)
}
// Returns the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/clippingResistancePriority(for:)
func (s_ StackView) ClippingResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("clippingResistancePriorityForOrientation:")
	ret := s_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Returns the custom spacing, in points, between a specified view in the stack view and the view that follows it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/customSpacing(after:)
func (s_ StackView) CustomSpacingAfterView(view unsafe.Pointer) float64 {
	sel := objc.RegisterName("customSpacingAfterView:")
	ret := s_.ID.Send(sel, view)
	return float64(ret)
}
// Returns the Auto Layout priority for the stack view to minimize its size to fit its contained views as closely as possible, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/huggingPriority(for:)
func (s_ StackView) HuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("huggingPriorityForOrientation:")
	ret := s_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Adds the provided view to the array of arranged subviews at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/insertArrangedSubview(_:at:)
func (s_ StackView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertArrangedSubview:atIndex:")
	s_.ID.Send(sel, view, index)
}
// Adds a view to a stack view gravity area at a specified index position. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/insertView(_:at:in:)
func (s_ StackView) InsertViewAtIndexInGravity(view unsafe.Pointer, index uint, gravity unsafe.Pointer) {
	sel := objc.RegisterName("insertView:atIndex:inGravity:")
	s_.ID.Send(sel, view, index, gravity)
}
// Removes the provided view from the stack’s array of arranged subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/removeArrangedSubview(_:)
func (s_ StackView) RemoveArrangedSubview(view unsafe.Pointer) {
	sel := objc.RegisterName("removeArrangedSubview:")
	s_.ID.Send(sel, view)
}
// Removes a specified view from the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/removeView(_:)
func (s_ StackView) RemoveView(view unsafe.Pointer) {
	sel := objc.RegisterName("removeView:")
	s_.ID.Send(sel, view)
}
// Sets the Auto Layout priority for resisting clipping of views in the stack view when Auto Layout attempts to reduce the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setClippingResistancePriority(_:for:)
func (s_ StackView) SetClippingResistancePriorityForOrientation(clippingResistancePriority unsafe.Pointer, orientation unsafe.Pointer) {
	sel := objc.RegisterName("setClippingResistancePriority:forOrientation:")
	s_.ID.Send(sel, clippingResistancePriority, orientation)
}
// Specifies the custom spacing, in points, between a specified view and the view that follows it in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setCustomSpacing(_:after:)
func (s_ StackView) SetCustomSpacingAfterView(spacing float64, view unsafe.Pointer) {
	sel := objc.RegisterName("setCustomSpacing:afterView:")
	s_.ID.Send(sel, spacing, view)
}
// Sets the Auto Layout priority for the stack view to minimize its size, for a specified user interface axis. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setHuggingPriority(_:for:)
func (s_ StackView) SetHuggingPriorityForOrientation(huggingPriority unsafe.Pointer, orientation unsafe.Pointer) {
	sel := objc.RegisterName("setHuggingPriority:forOrientation:")
	s_.ID.Send(sel, huggingPriority, orientation)
}
// Specifies an array of views for a specified gravity area in the stack view, replacing any previous views in that area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setViews(_:in:)
func (s_ StackView) SetViewsInGravity(views unsafe.Pointer, gravity unsafe.Pointer) {
	sel := objc.RegisterName("setViews:inGravity:")
	s_.ID.Send(sel, views, gravity)
}
// Sets the Auto Layout priority for a view to remain attached to the stack view when Auto Layout reduces the stack view’s size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/setVisibilityPriority(_:for:)
func (s_ StackView) SetVisibilityPriorityForView(priority unsafe.Pointer, view unsafe.Pointer) {
	sel := objc.RegisterName("setVisibilityPriority:forView:")
	s_.ID.Send(sel, priority, view)
}
// Returns the array of views in the specified gravity area in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/views(in:)
func (s_ StackView) ViewsInGravity(gravity unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("viewsInGravity:")
	ret := s_.ID.Send(sel, gravity)
	return unsafe.Pointer(ret)
}
// Returns the visibility priority for a specified view in the stack view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSStackView/visibilityPriority(for:)
func (s_ StackView) VisibilityPriorityForView(view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("visibilityPriorityForView:")
	ret := s_.ID.Send(sel, view)
	return unsafe.Pointer(ret)
}


