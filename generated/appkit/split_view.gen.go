// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SplitView] class.
var splitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}

type _SplitViewClass struct {
	class objc.Class
}

// A view that arranges two or more views in a linear stack running horizontally or vertically. [Full Topic]
//
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

// Adds a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s_ SplitView) AddArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}
// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/adjustSubviews()
func (s_ SplitView) AdjustSubviews() {
	objc.Send[objc.ID](s_.ID, objc.Sel("adjustSubviews"))
}
// Draws a divider between two of the split view’s subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s_ SplitView) DrawDividerInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawDividerInRect:"), rect)
}
// Returns the priority of the subview’s width or height when resizing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}
// Adds a view as an arranged split pane at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s_ SplitView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}
// The type of pane splitter. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isPaneSplitter
func (s_ SplitView) IsPaneSplitter() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPaneSplitter"))
	return rv
}
// Returns whether the specified view is in a collapsed state. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s_ SplitView) IsSubviewCollapsed(subview unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSubviewCollapsed:"), subview)
	return rv
}
// Returns the maximum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Returns the minimum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Removes a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s_ SplitView) RemoveArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}
// Sets the priority for split view subviews to maintain their width or height. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority unsafe.Pointer, subviewIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}
// Sets the type of splitter. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setIsPaneSplitter:
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPaneSplitter:"), flag)
}
// Updates the location of a divider you specify by index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}


