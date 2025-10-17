// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitView] class.
var SplitViewClass objc.Class

func init() {
	SplitViewClass = objc.GetClass("NSSplitView")
}

type SplitView struct {
	objc.ID
}

func SplitViewFrom(ptr unsafe.Pointer) SplitView {
	return SplitView{
		ID: objc.ID(ptr),
	}
}


// Adds a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s_ SplitView) AddArrangedSubview(view unsafe.Pointer) {
	sel := objc.RegisterName("addArrangedSubview:")
	s_.ID.Send(sel, view)
}
// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/adjustSubviews()
func (s_ SplitView) AdjustSubviews() {
	sel := objc.RegisterName("adjustSubviews")
	s_.ID.Send(sel)
}
// Draws a divider between two of the split view’s subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s_ SplitView) DrawDividerInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("drawDividerInRect:")
	s_.ID.Send(sel, rect)
}
// Returns the priority of the subview’s width or height when resizing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) unsafe.Pointer {
	sel := objc.RegisterName("holdingPriorityForSubviewAtIndex:")
	ret := s_.ID.Send(sel, subviewIndex)
	return unsafe.Pointer(ret)
}
// Adds a view as an arranged split pane at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s_ SplitView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertArrangedSubview:atIndex:")
	s_.ID.Send(sel, view, index)
}
// The type of pane splitter. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isPaneSplitter
func (s_ SplitView) IsPaneSplitter() bool {
	sel := objc.RegisterName("isPaneSplitter")
	ret := s_.ID.Send(sel)
	return ret != 0
}
// Returns whether the specified view is in a collapsed state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s_ SplitView) IsSubviewCollapsed(subview unsafe.Pointer) bool {
	sel := objc.RegisterName("isSubviewCollapsed:")
	ret := s_.ID.Send(sel, subview)
	return ret != 0
}
// Returns the maximum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	sel := objc.RegisterName("maxPossiblePositionOfDividerAtIndex:")
	ret := s_.ID.Send(sel, dividerIndex)
	return float64(ret)
}
// Returns the minimum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	sel := objc.RegisterName("minPossiblePositionOfDividerAtIndex:")
	ret := s_.ID.Send(sel, dividerIndex)
	return float64(ret)
}
// Removes a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s_ SplitView) RemoveArrangedSubview(view unsafe.Pointer) {
	sel := objc.RegisterName("removeArrangedSubview:")
	s_.ID.Send(sel, view)
}
// Sets the priority for split view subviews to maintain their width or height. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority unsafe.Pointer, subviewIndex int) {
	sel := objc.RegisterName("setHoldingPriority:forSubviewAtIndex:")
	s_.ID.Send(sel, priority, subviewIndex)
}
// Sets the type of splitter. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setIsPaneSplitter:
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	sel := objc.RegisterName("setIsPaneSplitter:")
	s_.ID.Send(sel, flag)
}
// Updates the location of a divider you specify by index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	sel := objc.RegisterName("setPosition:ofDividerAtIndex:")
	s_.ID.Send(sel, position, dividerIndex)
}

