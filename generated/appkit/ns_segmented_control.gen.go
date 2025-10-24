// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSegmentedControl */


/* debug [class_header]: Header for NSSegmentedControl */
// The class instance for the [SegmentedControl] class.
var (
	SegmentedControlClass     _SegmentedControlClass
	SegmentedControlClassOnce sync.Once
)

func getSegmentedControlClass() _SegmentedControlClass {
	SegmentedControlClassOnce.Do(func() {
		SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}
	})
	return SegmentedControlClass
}

type _SegmentedControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SegmentedControl */
// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	IControl
	
/* debug [class_interface_properties]: Properties for SegmentedControl */
	// properties:
	ActiveCompressionOptions() IUserInterfaceCompressionOptions
	BorderShape() ControlBorderShape
	SetBorderShape(value ControlBorderShape)
	DoubleValueForSelectedSegment() float64
	IndexOfSelectedItem() int
	SpringLoaded() bool
	SetSpringLoaded(value bool)
	SegmentCount() int
	SetSegmentCount(value int)
	SegmentDistribution() SegmentDistribution
	SetSegmentDistribution(value SegmentDistribution)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SelectedSegment() int
	SetSelectedSegment(value int)
	SelectedSegmentBezelColor() IColor
	SetSelectedSegmentBezelColor(value IColor)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	IsSpringLoaded() bool
	SetIsSpringLoaded(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SegmentedControl */
	// methods:
	AlignmentForSegment(segment int) TextAlignment
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	ImageForSegment(segment int) IImage
	ImageScalingForSegment(segment int) ImageScaling
	IsEnabledForSegment(segment int) bool
	IsSelectedForSegment(segment int) bool
	LabelForSegment(segment int) foundation.String
	MenuForSegment(segment int) IMenu
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) Size /* not a class type */
	SelectSegmentWithTag(tag int) bool
	SetAlignmentForSegment(alignment TextAlignment, segment int)
	SetEnabledForSegment(enabled bool, segment int)
	SetImageForSegment(image IImage, segment int)
	SetImageScalingForSegment(scaling ImageScaling, segment int)
	SetLabelForSegment(label objc.IObject /* cross-framework: NSString */, segment int)
	SetMenuForSegment(menu IMenu, segment int)
	SetSelectedForSegment(selected bool, segment int)
	SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int)
	SetTagForSegment(tag int, segment int)
	SetToolTipForSegment(toolTip objc.IObject /* cross-framework: NSString */, segment int)
	SetWidthForSegment(width float64, segment int)
	ShowsMenuIndicatorForSegment(segment int) bool
	TagForSegment(segment int) int
	ToolTipForSegment(segment int) foundation.String
	WidthForSegment(segment int) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SegmentedControl */
// Alloc allocates a new instance without initialization.
func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.Send[SegmentedControl](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SegmentedControl) Autorelease() SegmentedControl {
	rv := objc.Send[SegmentedControl](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSegmentedControl creates a new SegmentedControl instance.
func NewSegmentedControl() SegmentedControl {
	return getSegmentedControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SegmentedControl */
// Display one or more buttons in a single horizontal group.
//
// The class uses an class to implement much of the control’s functionality. Most methods in are simply cover methods that call the corresponding method in . The methods of that do not have covers relate to accessing and setting values for tags and tooltips, programatically setting the key segment, and establishing the mode of the control. The features of a segmented control include the following: A segment can have an image, text (label), menu, tooltip, and tag. A segmented control can contain images or text, but not both. Either the control or individual segments can be enabled or disabled. Segmented controls have four tracking modes, described in . You use these modes with the property. Each segment can be either a fixed width or autosized to fit the contents. If a segment has text and is marked as autosizing, then the text may be truncated so that the control completely fits. If an image is too large to fit in a segment, it is clipped. If Full Keyboard Access is enabled in System Preferences > Keyboard, the keyboard may be used to move between and select segments.


// Display one or more buttons in a single horizontal group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl
type SegmentedControl struct {
	Control
}

// SegmentedControlFrom constructs a [SegmentedControl] from an unsafe.Pointer.
//
// Display one or more buttons in a single horizontal group.
func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: ControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SegmentedControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(images:trackingMode:target:action:)
func NewSegmentedControlWithImagesTrackingModeTargetAction(images []Image, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.SEL) SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(getSegmentedControlClass().class), objc.Sel("segmentedControlWithImages:trackingMode:target:action:"), images, trackingMode, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewSegmentedControlWithImagesTrackingModeTargetAction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(labels:trackingMode:target:action:)
func NewSegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.SEL) SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(getSegmentedControlClass().class), objc.Sel("segmentedControlWithLabels:trackingMode:target:action:"), labels, trackingMode, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewSegmentedControlWithLabelsTrackingModeTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SegmentedControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(images:trackingMode:target:action:)
func (sc _SegmentedControlClass) SegmentedControlWithImagesTrackingModeTargetAction(images []Image, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("segmentedControlWithImages:trackingMode:target:action:"), images, trackingMode, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SegmentedControlWithImagesTrackingModeTargetAction) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/init(labels:trackingMode:target:action:)
func (sc _SegmentedControlClass) SegmentedControlWithLabelsTrackingModeTargetAction(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("segmentedControlWithLabels:trackingMode:target:action:"), labels, trackingMode, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SegmentedControlWithLabelsTrackingModeTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SegmentedControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SegmentedControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/alignment(forSegment:)
func (s_ SegmentedControl) AlignmentForSegment(segment int) TextAlignment {
	rv := objc.Send[TextAlignment](s_.ID, objc.Sel("alignmentForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: AlignmentForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/compress(withPrioritizedCompressionOptions:)
func (s_ SegmentedControl) CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) {
	objc.Send[objc.ID](s_.ID, objc.Sel("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}/* debug [instance_methods/method]: CompressWithPrioritizedCompressionOptions */


// Returns the image associated with the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/image(forSegment:)
func (s_ SegmentedControl) ImageForSegment(segment int) IImage {
	rv := objc.Send[Image](s_.ID, objc.Sel("imageForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: ImageForSegment */


// Returns the scaling mode used to display the specified segment’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/imageScaling(forSegment:)
func (s_ SegmentedControl) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.Send[ImageScaling](s_.ID, objc.Sel("imageScalingForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: ImageScalingForSegment */


// Returns a Boolean value indicating whether the specified segment is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isEnabled(forSegment:)
func (s_ SegmentedControl) IsEnabledForSegment(segment int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabledForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: IsEnabledForSegment */


// Returns a Boolean value indicating whether the specified segment is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isSelected(forSegment:)
func (s_ SegmentedControl) IsSelectedForSegment(segment int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSelectedForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: IsSelectedForSegment */


// Returns the label of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/label(forSegment:)
func (s_ SegmentedControl) LabelForSegment(segment int) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("labelForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: LabelForSegment */


// Returns the menu for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/menu(forSegment:)
func (s_ SegmentedControl) MenuForSegment(segment int) IMenu {
	rv := objc.Send[Menu](s_.ID, objc.Sel("menuForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: MenuForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/minimumSize(withPrioritizedCompressionOptions:)
func (s_ SegmentedControl) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) Size /* not a class type */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}/* debug [instance_methods/method]: MinimumSizeWithPrioritizedCompressionOptions */


// Selects the segment with the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectSegment(withTag:)
func (s_ SegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("selectSegmentWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: SelectSegmentWithTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setAlignment(_:forSegment:)
func (s_ SegmentedControl) SetAlignmentForSegment(alignment TextAlignment, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlignment:forSegment:"), alignment, segment)
}/* debug [instance_methods/method]: SetAlignmentForSegment */


// Sets the enabled state of the specified segment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setEnabled(_:forSegment:)
func (s_ SegmentedControl) SetEnabledForSegment(enabled bool, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:forSegment:"), enabled, segment)
}/* debug [instance_methods/method]: SetEnabledForSegment */


// Sets the image for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setImage(_:forSegment:)
func (s_ SegmentedControl) SetImageForSegment(image IImage, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImage:forSegment:"), image, segment)
}/* debug [instance_methods/method]: SetImageForSegment */


// Sets the scaling mode used to display the specified segment’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setImageScaling(_:forSegment:)
func (s_ SegmentedControl) SetImageScalingForSegment(scaling ImageScaling, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageScaling:forSegment:"), scaling, segment)
}/* debug [instance_methods/method]: SetImageScalingForSegment */


// Sets the label for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setLabel(_:forSegment:)
func (s_ SegmentedControl) SetLabelForSegment(label objc.IObject /* cross-framework: NSString */, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:forSegment:"), label, segment)
}/* debug [instance_methods/method]: SetLabelForSegment */


// Sets the menu for the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setMenu(_:forSegment:)
func (s_ SegmentedControl) SetMenuForSegment(menu IMenu, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMenu:forSegment:"), menu, segment)
}/* debug [instance_methods/method]: SetMenuForSegment */


// Sets the selection state of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setSelected(_:forSegment:)
func (s_ SegmentedControl) SetSelectedForSegment(selected bool, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelected:forSegment:"), selected, segment)
}/* debug [instance_methods/method]: SetSelectedForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setShowsMenuIndicator(_:forSegment:)
func (s_ SegmentedControl) SetShowsMenuIndicatorForSegment(showsMenuIndicator bool, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsMenuIndicator:forSegment:"), showsMenuIndicator, segment)
}/* debug [instance_methods/method]: SetShowsMenuIndicatorForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setTag(_:forSegment:)
func (s_ SegmentedControl) SetTagForSegment(tag int, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTag:forSegment:"), tag, segment)
}/* debug [instance_methods/method]: SetTagForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setToolTip(_:forSegment:)
func (s_ SegmentedControl) SetToolTipForSegment(toolTip objc.IObject /* cross-framework: NSString */, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToolTip:forSegment:"), toolTip, segment)
}/* debug [instance_methods/method]: SetToolTipForSegment */


// Sets the width of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/setWidth(_:forSegment:)
func (s_ SegmentedControl) SetWidthForSegment(width float64, segment int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWidth:forSegment:"), width, segment)
}/* debug [instance_methods/method]: SetWidthForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/showsMenuIndicator(forSegment:)
func (s_ SegmentedControl) ShowsMenuIndicatorForSegment(segment int) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsMenuIndicatorForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: ShowsMenuIndicatorForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/tag(forSegment:)
func (s_ SegmentedControl) TagForSegment(segment int) int {
	rv := objc.Send[int](s_.ID, objc.Sel("tagForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: TagForSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/toolTip(forSegment:)
func (s_ SegmentedControl) ToolTipForSegment(segment int) foundation.String {
	rv := objc.Send[foundation.String](s_.ID, objc.Sel("toolTipForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: ToolTipForSegment */


// Returns the width of the specified segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/width(forSegment:)
func (s_ SegmentedControl) WidthForSegment(segment int) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("widthForSegment:"), segment)
	return rv
}/* debug [instance_methods/method]: WidthForSegment */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SegmentedControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/activeCompressionOptions
func (s_ SegmentedControl) ActiveCompressionOptions() IUserInterfaceCompressionOptions {
	rv := objc.Send[UserInterfaceCompressionOptions](s_.ID, objc.Sel("activeCompressionOptions"))
	return rv
}/* debug [instance_properties/getter]: activeCompressionOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/borderShape
func (s_ SegmentedControl) BorderShape() ControlBorderShape {
	rv := objc.Send[ControlBorderShape](s_.ID, objc.Sel("borderShape"))
	return rv
}/* debug [instance_properties/getter]: borderShape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/borderShape
func (s_ SegmentedControl) SetBorderShape(value ControlBorderShape) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderShape:"), value)
}/* debug [instance_properties/setter]: borderShape */


// When the tracking mode for the control is set to use a momentary accelerator, returns a value for the selected segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/doubleValueForSelectedSegment
func (s_ SegmentedControl) DoubleValueForSelectedSegment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("doubleValueForSelectedSegment"))
	return rv
}/* debug [instance_properties/getter]: doubleValueForSelectedSegment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/indexOfSelectedItem
func (s_ SegmentedControl) IndexOfSelectedItem() int {
	rv := objc.Send[int](s_.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}/* debug [instance_properties/getter]: indexOfSelectedItem */


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isSpringLoaded
func (s_ SegmentedControl) SpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("springLoaded"))
	return rv
}/* debug [instance_properties/getter]: springLoaded */


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/isSpringLoaded
func (s_ SegmentedControl) SetSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSpringLoaded:"), value)
}/* debug [instance_properties/setter]: springLoaded */


// The number of segments in the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentCount
func (s_ SegmentedControl) SegmentCount() int {
	rv := objc.Send[int](s_.ID, objc.Sel("segmentCount"))
	return rv
}/* debug [instance_properties/getter]: segmentCount */


// The number of segments in the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentCount
func (s_ SegmentedControl) SetSegmentCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentCount:"), value)
}/* debug [instance_properties/setter]: segmentCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentDistribution
func (s_ SegmentedControl) SegmentDistribution() SegmentDistribution {
	rv := objc.Send[SegmentDistribution](s_.ID, objc.Sel("segmentDistribution"))
	return rv
}/* debug [instance_properties/getter]: segmentDistribution */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentDistribution
func (s_ SegmentedControl) SetSegmentDistribution(value SegmentDistribution) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentDistribution:"), value)
}/* debug [instance_properties/setter]: segmentDistribution */


// The visual style used to display the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentStyle
func (s_ SegmentedControl) SegmentStyle() SegmentStyle {
	rv := objc.Send[SegmentStyle](s_.ID, objc.Sel("segmentStyle"))
	return rv
}/* debug [instance_properties/getter]: segmentStyle */


// The visual style used to display the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/segmentStyle
func (s_ SegmentedControl) SetSegmentStyle(value SegmentStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSegmentStyle:"), value)
}/* debug [instance_properties/setter]: segmentStyle */


// The index of the selected segment of the control, or if no segment is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegment
func (s_ SegmentedControl) SelectedSegment() int {
	rv := objc.Send[int](s_.ID, objc.Sel("selectedSegment"))
	return rv
}/* debug [instance_properties/getter]: selectedSegment */


// The index of the selected segment of the control, or if no segment is selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegment
func (s_ SegmentedControl) SetSelectedSegment(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegment:"), value)
}/* debug [instance_properties/setter]: selectedSegment */


// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SelectedSegmentBezelColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("selectedSegmentBezelColor"))
	return rv
}/* debug [instance_properties/getter]: selectedSegmentBezelColor */


// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}/* debug [instance_properties/setter]: selectedSegmentBezelColor */


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/trackingMode
func (s_ SegmentedControl) TrackingMode() SegmentSwitchTracking {
	rv := objc.Send[SegmentSwitchTracking](s_.ID, objc.Sel("trackingMode"))
	return rv
}/* debug [instance_properties/getter]: trackingMode */


// The type of tracking behavior the control exhibits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/trackingMode
func (s_ SegmentedControl) SetTrackingMode(value SegmentSwitchTracking) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTrackingMode:"), value)
}/* debug [instance_properties/setter]: trackingMode */


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) IsSpringLoaded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSpringLoaded"))
	return rv
}/* debug [instance_properties/getter]: isSpringLoaded */


// A Boolean value that indicates whether spring loading is enabled for the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/isspringloaded
func (s_ SegmentedControl) SetIsSpringLoaded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSpringLoaded:"), value)
}/* debug [instance_properties/setter]: isSpringLoaded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSegmentedControl */


