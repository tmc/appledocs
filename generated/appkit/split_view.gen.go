
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [SplitView] class.
var SplitViewClass _SplitViewClass

func init() {
	SplitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}
}

type _SplitViewClass struct {
	objc.Class
}

// An interface definition for the [SplitView] class.
type ISplitView interface {
	ID() objc.ID
	AddArrangedSubview(view unsafe.Pointer)
	AdjustSubviews()
	DrawDividerInRect(rect foundation.Rect)
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

type SplitView struct {
	id objc.ID
}

func SplitViewFrom(ptr unsafe.Pointer) SplitView {
	return SplitView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SplitView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.Send[SplitView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SplitViewClass) New() SplitView {
	rv := objc.Send[SplitView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSplitView creates and returns a new initialized instance.
func NewSplitView() SplitView {
	return SplitViewClass.New()
}

// Init initializes the instance.
func (s_ SplitView) Init() SplitView {
	rv := objc.Send[SplitView](s_.ID(), selInit)
	return rv
}
// Adds a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/addArrangedSubview(_:)
func (s_ SplitView) AddArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("addArrangedSubview:"), view)
}
// Adjusts the sizes of the split view’s subviews so they (plus the dividers) fill the split view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/adjustSubviews()
func (s_ SplitView) AdjustSubviews() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("adjustSubviews"))
}
// Draws a divider between two of the split view’s subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/drawDivider(in:)
func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("drawDividerInRect:"), rect)
}
// Returns the priority of the subview’s width or height when resizing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/holdingPriorityForSubview(at:)
func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}
// Adds a view as an arranged split pane at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/insertArrangedSubview(_:at:)
func (s_ SplitView) InsertArrangedSubviewAtIndex(view unsafe.Pointer, index int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("insertArrangedSubview:atIndex:"), view, index)
}
// The type of pane splitter. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isPaneSplitter
func (s_ SplitView) IsPaneSplitter() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("isPaneSplitter"))
	return rv
}
// Returns whether the specified view is in a collapsed state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isSubviewCollapsed(_:)
func (s_ SplitView) IsSubviewCollapsed(subview unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("isSubviewCollapsed:"), subview)
	return rv
}
// Returns the maximum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/maxPossiblePositionOfDivider(at:)
func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Returns the minimum possible position of the divider at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/minPossiblePositionOfDivider(at:)
func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}
// Removes a view as an arranged split pane. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/removeArrangedSubview(_:)
func (s_ SplitView) RemoveArrangedSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("removeArrangedSubview:"), view)
}
// Sets the priority for split view subviews to maintain their width or height. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setHoldingPriority(_:forSubviewAt:)
func (s_ SplitView) SetHoldingPriorityForSubviewAtIndex(priority unsafe.Pointer, subviewIndex int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}
// Sets the type of splitter. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setIsPaneSplitter:
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setIsPaneSplitter:"), flag)
}
// Updates the location of a divider you specify by index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/setPosition(_:ofDividerAt:)
func (s_ SplitView) SetPositionOfDividerAtIndex(position float64, dividerIndex int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}
// The array of views that the split view arranges as its split panes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/arrangedSubviews
func (s_ SplitView) ArrangedSubviews() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("arrangedSubviews"))
	return rv
}
// A Boolean value that determines whether the split view arranges all of its subviews as split panes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("arrangesAllSubviews"))
	return rv
}
// SetArrangesAllSubviews sets the value of the arrangesAllSubviews property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/arrangesAllSubviews
func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setArrangesAllSubviews:"), value)
}
// The name to use when the system automatically saves the split view’s divider configuration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) AutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("autosaveName"))
	return rv
}
// SetAutosaveName sets the value of the autosaveName property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/autosaveName-swift.property
func (s_ SplitView) SetAutosaveName(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAutosaveName:"), value)
}
// The split view’s delegate. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/delegate
func (s_ SplitView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDelegate:"), value)
}
// The color of the dividers that the split view draws between subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/dividerColor
func (s_ SplitView) DividerColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("dividerColor"))
	return rv
}
// The style of divider between views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) DividerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("dividerStyle"))
	return rv
}
// SetDividerStyle sets the value of the dividerStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/dividerStyle-swift.property
func (s_ SplitView) SetDividerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDividerStyle:"), value)
}
// The thickness of the dividers for the split view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/dividerThickness
func (s_ SplitView) DividerThickness() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("dividerThickness"))
	return rv
}
// A Boolean value that determines the geometric orientation of the split view’s dividers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) Vertical() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("vertical"))
	return rv
}
// SetVertical sets the value of the vertical property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSplitView/isVertical
func (s_ SplitView) SetVertical(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVertical:"), value)
}
