// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [SegmentedCell] class.
var (
	SegmentedCellClass     _SegmentedCellClass
	SegmentedCellClassOnce sync.Once
)

func getSegmentedCellClass() _SegmentedCellClass {
	SegmentedCellClassOnce.Do(func() {
		SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}
	})
	return SegmentedCellClass
}

type _SegmentedCellClass struct {
	class objc.Class
}





// An interface definition for the [SegmentedCell] class.
type ISegmentedCell interface {
	IActionCell
	

	// properties:
	SegmentCount() int
	SetSegmentCount(value int)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SelectedSegment() int
	SetSelectedSegment(value int)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)


	

	// methods:
	DrawSegmentInFrameWithView(segment int, frame corefoundation.CGRect, controlView IView)
	ImageForSegment(segment int) IImage
	ImageScalingForSegment(segment int) ImageScaling
	InteriorBackgroundStyleForSegment(segment int) BackgroundStyle
	IsEnabledForSegment(segment int) bool
	IsSelectedForSegment(segment int) bool
	LabelForSegment(segment int) foundation.String
	MakeNextSegmentKey()
	MakePreviousSegmentKey()
	MenuForSegment(segment int) IMenu
	SelectSegmentWithTag(tag int) bool
	SetEnabledForSegment(enabled bool, segment int)
	SetImageForSegment(image IImage, segment int)
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	SetLabelForSegment(label foundation.foundation.INSString, segment int)
	SetMenuForSegment(menu IMenu, segment int)
	SetSelectedForSegment(selected bool, segment int)
	SetTagForSegment(tag int, segment int)
	SetToolTipForSegment(toolTip foundation.foundation.INSString, segment int)
	SetWidthForSegment(width float64, segment int)
	TagForSegment(segment int) int
	ToolTipForSegment(segment int) foundation.String
	WidthForSegment(segment int) float64


}





// Alloc allocates a new instance without initialization.
func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.Send[SegmentedCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.Send[SegmentedCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SegmentedCell) Autorelease() SegmentedCell {
	rv := objc.Send[SegmentedCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSegmentedCell creates a new SegmentedCell instance.
func NewSegmentedCell() SegmentedCell {
	return getSegmentedCellClass().New()
}





// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control.
//
// Use the methods of to customize the attributes of a segmented control. To customize the appearance of individual segments, you can also subclass and override the method.


// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell
type SegmentedCell struct {
	ActionCell
}

// SegmentedCellFrom constructs a [SegmentedCell] from an unsafe.Pointer.
//
// An object implements the appearance and behavior of a horizontal button divided into multiple segments. This class is used in conjunction with the class to implement a segmented control.
func SegmentedCellFrom(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ActionCell: ActionCellFrom(ptr),
	}
}




















// Draws the image and label of the segment in the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/drawSegment(_:inFrame:with:)
func (s_ SegmentedCell) DrawSegmentInFrameWithView(segment int, frame corefoundation.CGRect, controlView IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawSegment:inFrame:withView:"), segment, frame, controlView)
}


// Returns the image associated with the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/image(forSegment:)
func (s_ SegmentedCell) ImageForSegment(segment int) IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("imageForSegment:"), segment)
	return rv
}


// Returns the image scaling mode associated with the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/imageScaling(forSegment:)
func (s_ SegmentedCell) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.Send[ImageScaling](s_.ID, objc.Sel("imageScalingForSegment:"), segment)
	return rv
}


// Returns the interior background style for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/interiorBackgroundStyle(forSegment:)
func (s_ SegmentedCell) InteriorBackgroundStyleForSegment(segment int) BackgroundStyle {
	rv := objc.Send[BackgroundStyle](s_.ID, objc.Sel("interiorBackgroundStyleForSegment:"), segment)
	return rv
}


// Returns a Boolean value indicating whether the specified segment is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/isEnabled(forSegment:)
func (s_ SegmentedCell) IsEnabledForSegment(segment int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}


// Returns a Boolean value indicating whether the specified segment is selected,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/isSelected(forSegment:)
func (s_ SegmentedCell) IsSelectedForSegment(segment int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}


// Returns the label of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/label(forSegment:)
func (s_ SegmentedCell) LabelForSegment(segment int) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("labelForSegment:"), segment)
	return rv
}


// Selects the next segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/makeNextSegmentKey()
func (s_ SegmentedCell) MakeNextSegmentKey() {
	objc.Send[objc.ID](s_.ID, objc.Sel("makeNextSegmentKey"))
}


// Selects the previous segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/makePreviousSegmentKey()
func (s_ SegmentedCell) MakePreviousSegmentKey() {
	objc.Send[objc.ID](s_.ID, objc.Sel("makePreviousSegmentKey"))
}


// Returns the menu for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/menu(forSegment:)
func (s_ SegmentedCell) MenuForSegment(segment int) IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("menuForSegment:"), segment)
	return rv
}


// Selects the segment with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/selectSegment(withTag:)
func (s_ SegmentedCell) SelectSegmentWithTag(tag int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}


// Sets the enabled state of the specified segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setEnabled(_:forSegment:)
func (s_ SegmentedCell) SetEnabledForSegment(enabled bool, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}


// Sets the image for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setImage(_:forSegment:)
func (s_ SegmentedCell) SetImageForSegment(image IImage, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:forSegment:"), image, segment)
}


// Sets the image scaling mode for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setImageScaling(_:forSegment:)
func (s_ SegmentedCell) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}


// Sets the label for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setLabel(_:forSegment:)
func (s_ SegmentedCell) SetLabelForSegment(label foundation.foundation.INSString, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:forSegment:"), label, segment)
}


// Sets the menu for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setMenu(_:forSegment:)
func (s_ SegmentedCell) SetMenuForSegment(menu IMenu, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenu:forSegment:"), menu, segment)
}


// Sets the selection state of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setSelected(_:forSegment:)
func (s_ SegmentedCell) SetSelectedForSegment(selected bool, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelected:forSegment:"), selected, segment)
}


// Sets the tag for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setTag(_:forSegment:)
func (s_ SegmentedCell) SetTagForSegment(tag int, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTag:forSegment:"), tag, segment)
}


// Sets the tooltip for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setToolTip(_:forSegment:)
func (s_ SegmentedCell) SetToolTipForSegment(toolTip foundation.foundation.INSString, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToolTip:forSegment:"), toolTip, segment)
}


// Sets the width of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/setWidth(_:forSegment:)
func (s_ SegmentedCell) SetWidthForSegment(width float64, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:forSegment:"), width, segment)
}


// Returns the tag of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/tag(forSegment:)
func (s_ SegmentedCell) TagForSegment(segment int) int {
	rv := objc.Send[int](s_.ID, objc.Sel("tagForSegment:"), segment)
	return rv
}


// Returns the tooltip of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/toolTip(forSegment:)
func (s_ SegmentedCell) ToolTipForSegment(segment int) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("toolTipForSegment:"), segment)
	return rv
}


// Returns the width of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/width(forSegment:)
func (s_ SegmentedCell) WidthForSegment(segment int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("widthForSegment:"), segment)
	return rv
}







// The number of segments in the segmented control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentCount
func (s_ SegmentedCell) SegmentCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("segmentCount"))
	return rv
}


// The number of segments in the segmented control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentCount
func (s_ SegmentedCell) SetSegmentCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentCount:"), value)
}


// The visual style used to display the segmented control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentStyle
func (s_ SegmentedCell) SegmentStyle() SegmentStyle {
	rv := objc.Send[SegmentStyle](s_.ID, objc.Sel("segmentStyle"))
	return rv
}


// The visual style used to display the segmented control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/segmentStyle
func (s_ SegmentedCell) SetSegmentStyle(value SegmentStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentStyle:"), value)
}


// The index of the selected segment of the control, or if no segment is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/selectedSegment
func (s_ SegmentedCell) SelectedSegment() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedSegment"))
	return rv
}


// The index of the selected segment of the control, or if no segment is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/selectedSegment
func (s_ SegmentedCell) SetSelectedSegment(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegment:"), value)
}


// The tracking mode used for the segments of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/trackingMode
func (s_ SegmentedCell) TrackingMode() SegmentSwitchTracking {
	rv := objc.Send[SegmentSwitchTracking](s_.ID, objc.Sel("trackingMode"))
	return rv
}


// The tracking mode used for the segments of the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedCell/trackingMode
func (s_ SegmentedCell) SetTrackingMode(value SegmentSwitchTracking) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackingMode:"), value)
}








