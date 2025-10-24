// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScreen */


/* debug [class_header]: Header for NSScreen */
// The class instance for the [Screen] class.
var (
	ScreenClass     _ScreenClass
	ScreenClassOnce sync.Once
)

func getScreenClass() _ScreenClass {
	ScreenClassOnce.Do(func() {
		ScreenClass = _ScreenClass{objc.GetClass("NSScreen")}
	})
	return ScreenClass
}

type _ScreenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Screen */
// An interface definition for the [Screen] class.
type IScreen interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Screen */
	// properties:
	BackingScaleFactor() float64
	ColorSpace() IColorSpace
	Depth() WindowDepth
	DeviceDescription() foundation.IDictionary
	DisplayUpdateGranularity() float64
	Frame() Rect /* not a class type */
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SafeAreaInsets() foundation.EdgeInsets
	SupportedWindowDepths() NSWindowDepth
	VisibleFrame() Rect /* not a class type */
	AuxiliaryTopLeftArea() Rect /* not a class type */
	SetAuxiliaryTopLeftArea(value Rect /* not a class type */)
	AuxiliaryTopRightArea() Rect /* not a class type */
	SetAuxiliaryTopRightArea(value Rect /* not a class type */)
	CgDirectDisplayID() DirectDisplayID /* not a class type */
	SetCgDirectDisplayID(value DirectDisplayID /* not a class type */)
	LastDisplayUpdateTimestamp() float64
	SetLastDisplayUpdateTimestamp(value float64)
	MaximumExtendedDynamicRangeColorComponentValue() float64
	SetMaximumExtendedDynamicRangeColorComponentValue(value float64)
	MaximumFramesPerSecond() int
	SetMaximumFramesPerSecond(value int)
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64
	SetMaximumPotentialExtendedDynamicRangeColorComponentValue(value float64)
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64
	SetMaximumReferenceExtendedDynamicRangeColorComponentValue(value float64)
	MaximumRefreshInterval() float64
	SetMaximumRefreshInterval(value float64)
	MinimumRefreshInterval() float64
	SetMinimumRefreshInterval(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Screen */
	// methods:
	BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Screen */
// Alloc allocates a new instance without initialization.
func (sc _ScreenClass) Alloc() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScreenClass) New() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Screen) Init() Screen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Screen) Autorelease() Screen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScreen creates a new Screen instance.
func NewScreen() Screen {
	return getScreenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Screen */
// An object that describes the attributes of a computer’s monitor or screen.
//
// An app may use an object to retrieve information about a screen and use this information to decide what to display on that screen. For example, an app may use the method to find out which of the available screens can best represent color and then might choose to display all of its windows on that screen. Create the application object before you use the methods in this class, so that the application object can make the necessary connection to the window system. You can make sure the application object exists by invoking the method of . If you created your app with Xcode, the application object is automatically created for you during initialization.


// An object that describes the attributes of a computer’s monitor or screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen
type Screen struct {
	objectivec.Object
}

// ScreenFrom constructs a [Screen] from an unsafe.Pointer.
//
// An object that describes the attributes of a computer’s monitor or screen.
func ScreenFrom(ptr unsafe.Pointer) Screen {
	return Screen{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Screen *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Screen */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Screen */

// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("screensHaveSeparateSpaces"))
	return rv
}/* debug [class_properties_class/property]: screensHaveSeparateSpaces */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Screen */

// Converts a rectangle in global screen coordinates to a pixel aligned rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/backingAlignedRect(_:options:)
func (s_ Screen) BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}/* debug [instance_methods/method]: BackingAlignedRectOptions */


// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}/* debug [instance_methods/method]: CanRepresentDisplayGamut */


// Converts the rectangle from the device pixel aligned coordinates system of a screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectFromBacking(_:)
func (s_ Screen) ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromBacking */


// Converts the rectangle to the device pixel aligned coordinates system of a screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectToBacking(_:)
func (s_ Screen) ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectToBacking */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Screen */

// The backing store pixel scale factor for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/backingScaleFactor
func (s_ Screen) BackingScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("backingScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: backingScaleFactor */


// The color space of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/colorSpace
func (s_ Screen) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](s_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// The current bit depth and colorspace information of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/depth
func (s_ Screen) Depth() WindowDepth {
	rv := objc.Send[WindowDepth](s_.ID, objc.Sel("depth"))
	return rv
}/* debug [instance_properties/getter]: depth */


// The device dictionary for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deviceDescription
func (s_ Screen) DeviceDescription() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("deviceDescription"))
	return rv
}/* debug [instance_properties/getter]: deviceDescription */


// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/displayUpdateGranularity
func (s_ Screen) DisplayUpdateGranularity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("displayUpdateGranularity"))
	return rv
}/* debug [instance_properties/getter]: displayUpdateGranularity */


// The dimensions and location of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/frame
func (s_ Screen) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The localized name of the display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/localizedName
func (s_ Screen) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */


// The distances from the screen’s edges at which content isn’t obscured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/safeAreaInsets
func (s_ Screen) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](s_.ID, objc.Sel("safeAreaInsets"))
	return rv
}/* debug [instance_properties/getter]: safeAreaInsets */


// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (s_ Screen) ScreensHaveSeparateSpaces() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("screensHaveSeparateSpaces"))
	return rv
}/* debug [instance_properties/getter]: screensHaveSeparateSpaces */


// A zero-terminated array of the window depths supported by the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/supportedWindowDepths
func (s_ Screen) SupportedWindowDepths() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](s_.ID, objc.Sel("supportedWindowDepths"))
	return rv
}/* debug [instance_properties/getter]: supportedWindowDepths */


// The current location and dimensions of the visible screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/visibleFrame
func (s_ Screen) VisibleFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("visibleFrame"))
	return rv
}/* debug [instance_properties/getter]: visibleFrame */


// The unobscured portion of the top-left corner of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytopleftarea-uglc
func (s_ Screen) AuxiliaryTopLeftArea() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("auxiliaryTopLeftArea"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryTopLeftArea */


// The unobscured portion of the top-left corner of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytopleftarea-uglc
func (s_ Screen) SetAuxiliaryTopLeftArea(value Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuxiliaryTopLeftArea:"), value)
}/* debug [instance_properties/setter]: auxiliaryTopLeftArea */


// The unobscured portion of the top-right corner of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytoprightarea-gr2n
func (s_ Screen) AuxiliaryTopRightArea() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("auxiliaryTopRightArea"))
	return rv
}/* debug [instance_properties/getter]: auxiliaryTopRightArea */


// The unobscured portion of the top-right corner of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytoprightarea-gr2n
func (s_ Screen) SetAuxiliaryTopRightArea(value Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuxiliaryTopRightArea:"), value)
}/* debug [instance_properties/setter]: auxiliaryTopRightArea */


// The CGDirectDisplayID for this screen. This will return nil if there isn’t one and will never return kCGNullDirectDisplay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/cgdirectdisplayid-8ph5i
func (s_ Screen) CgDirectDisplayID() DirectDisplayID /* not a class type */ {
	rv := objc.Send[DirectDisplayID](s_.ID, objc.Sel("cgDirectDisplayID"))
	return rv
}/* debug [instance_properties/getter]: cgDirectDisplayID */


// The CGDirectDisplayID for this screen. This will return nil if there isn’t one and will never return kCGNullDirectDisplay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/cgdirectdisplayid-8ph5i
func (s_ Screen) SetCgDirectDisplayID(value DirectDisplayID /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCgDirectDisplayID:"), value)
}/* debug [instance_properties/setter]: cgDirectDisplayID */


// The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/lastdisplayupdatetimestamp
func (s_ Screen) LastDisplayUpdateTimestamp() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lastDisplayUpdateTimestamp"))
	return rv
}/* debug [instance_properties/getter]: lastDisplayUpdateTimestamp */


// The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/lastdisplayupdatetimestamp
func (s_ Screen) SetLastDisplayUpdateTimestamp(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLastDisplayUpdateTimestamp:"), value)
}/* debug [instance_properties/setter]: lastDisplayUpdateTimestamp */


// The current maximum color component value for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumextendeddynamicrangecolorcomponentvalue
func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}/* debug [instance_properties/getter]: maximumExtendedDynamicRangeColorComponentValue */


// The current maximum color component value for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumextendeddynamicrangecolorcomponentvalue
func (s_ Screen) SetMaximumExtendedDynamicRangeColorComponentValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumExtendedDynamicRangeColorComponentValue:"), value)
}/* debug [instance_properties/setter]: maximumExtendedDynamicRangeColorComponentValue */


// The maximum number of frames per second that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumframespersecond
func (s_ Screen) MaximumFramesPerSecond() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumFramesPerSecond"))
	return rv
}/* debug [instance_properties/getter]: maximumFramesPerSecond */


// The maximum number of frames per second that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumframespersecond
func (s_ Screen) SetMaximumFramesPerSecond(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumFramesPerSecond:"), value)
}/* debug [instance_properties/setter]: maximumFramesPerSecond */


// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumpotentialextendeddynamicrangecolorcomponentvalue
func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}/* debug [instance_properties/getter]: maximumPotentialExtendedDynamicRangeColorComponentValue */


// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumpotentialextendeddynamicrangecolorcomponentvalue
func (s_ Screen) SetMaximumPotentialExtendedDynamicRangeColorComponentValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumPotentialExtendedDynamicRangeColorComponentValue:"), value)
}/* debug [instance_properties/setter]: maximumPotentialExtendedDynamicRangeColorComponentValue */


// The current maximum color component value for reference rendering to the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumreferenceextendeddynamicrangecolorcomponentvalue
func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}/* debug [instance_properties/getter]: maximumReferenceExtendedDynamicRangeColorComponentValue */


// The current maximum color component value for reference rendering to the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumreferenceextendeddynamicrangecolorcomponentvalue
func (s_ Screen) SetMaximumReferenceExtendedDynamicRangeColorComponentValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumReferenceExtendedDynamicRangeColorComponentValue:"), value)
}/* debug [instance_properties/setter]: maximumReferenceExtendedDynamicRangeColorComponentValue */


// The largest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumrefreshinterval
func (s_ Screen) MaximumRefreshInterval() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumRefreshInterval"))
	return rv
}/* debug [instance_properties/getter]: maximumRefreshInterval */


// The largest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumrefreshinterval
func (s_ Screen) SetMaximumRefreshInterval(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumRefreshInterval:"), value)
}/* debug [instance_properties/setter]: maximumRefreshInterval */


// The shortest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/minimumrefreshinterval
func (s_ Screen) MinimumRefreshInterval() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumRefreshInterval"))
	return rv
}/* debug [instance_properties/getter]: minimumRefreshInterval */


// The shortest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/minimumrefreshinterval
func (s_ Screen) SetMinimumRefreshInterval(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinimumRefreshInterval:"), value)
}/* debug [instance_properties/setter]: minimumRefreshInterval */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScreen */



