// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SplitView] class.
var (
	splitViewClass     _SplitViewClass
	splitViewClassOnce sync.Once
)

func getSplitViewClass() _SplitViewClass {
	splitViewClassOnce.Do(func() {
		splitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}
	})
	return splitViewClass
}

type _SplitViewClass struct {
	class objc.Class
}

// An interface definition for the [SplitView] class.
type ISplitView interface {
	IView
	AddArrangedSubview(view unsafe.Pointer)
	AdjustSubviews()
	DrawDividerInRect(rect coregraphics.CGRect)
	HoldingPriorityForSubviewAtIndex(subviewIndex int) unsafe.Pointer
	InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int)
	IsPaneSplitter() bool
	IsSubviewCollapsed(subview unsafe.Pointer) bool
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	RemoveArrangedSubview(view unsafe.Pointer)
	SetHoldingPriorityForSubviewAtIndex(priority unsafe.Pointer, subviewIndex int)
	SetIsPaneSplitter(flag bool)
	SetPositionOfDividerAtIndex(position float64, dividerIndex int)
}

// A view that arranges two or more views in a linear stack running horizontally or vertically.
//
// A split view manages the dividers and orientation for a split view controller ( ). By default, dividers have a horizontal orientation so that the split view arranges its panes vertically from top to bottom. Divider indices are zero-based. If the property is , which is the default value, the top divider has an index of . If is , the leading divider has an index of .
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

// Alloc allocates a new instance without initialization.
func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.Send[SplitView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds a view as an arranged split pane.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s_ SplitView) AddArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addArrangedSubview:"), view)
}

// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/adjustSubviews()
func (s_ SplitView) AdjustSubviews() {
	objc.Send[objc.ID](s_.ID, objc.Sel("adjustSubviews"))
}

// Draws a divider between two of the split view’s subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s_ SplitView) DrawDividerInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawDividerInRect:"), rect)
}

// Returns the priority of the subview’s width or height when resizing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}

// Adds a view as an arranged split pane at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s_ SplitView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertArrangedSubview:atIndex:"), view, index)
}

// The type of pane splitter.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isPaneSplitter
func (s_ SplitView) IsPaneSplitter() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPaneSplitter"))
	return rv
}

// Returns whether the specified view is in a collapsed state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s_ SplitView) IsSubviewCollapsed(subview unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSubviewCollapsed:"), subview)
	return rv
}

// Returns the maximum possible position of the divider at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

// Returns the minimum possible position of the divider at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

// Removes a view as an arranged split pane.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s_ SplitView) RemoveArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeArrangedSubview:"), view)
}

// Sets the priority for split view subviews to maintain their width or height.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority unsafe.Pointer, subviewIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}

// Sets the type of splitter.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setIsPaneSplitter:
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPaneSplitter:"), flag)
}

// Updates the location of a divider you specify by index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}

// The array of views that the split view arranges as its split panes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangedSubviews
func (s_ SplitView) ArrangedSubviews() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("arrangedSubviews"))
	return rv
}
// A Boolean value that determines whether the split view arranges all of its subviews as split panes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("arrangesAllSubviews"))
	return rv
}

// SetArrangesAllSubviews sets the value of the arrangesAllSubviews property.
// A Boolean value that determines whether the split view arranges all of its subviews as split panes.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArrangesAllSubviews:"), value)
}
// The name to use when the system automatically saves the split view’s divider configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) AutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("autosaveName"))
	return rv
}

// SetAutosaveName sets the value of the autosaveName property.
// The name to use when the system automatically saves the split view’s divider configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) SetAutosaveName(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutosaveName:"), value)
}
// The split view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The split view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}
// The color of the dividers that the split view draws between subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerColor
func (s_ SplitView) DividerColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dividerColor"))
	return rv
}
// The style of divider between views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) DividerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dividerStyle"))
	return rv
}

// SetDividerStyle sets the value of the dividerStyle property.
// The style of divider between views.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) SetDividerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDividerStyle:"), value)
}
// The thickness of the dividers for the split view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/dividerThickness
func (s_ SplitView) DividerThickness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("dividerThickness"))
	return rv
}
// A Boolean value that determines the geometric orientation of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) Vertical() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("vertical"))
	return rv
}

// SetVertical sets the value of the vertical property.
// A Boolean value that determines the geometric orientation of the split view’s dividers.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVertical:"), value)
}


