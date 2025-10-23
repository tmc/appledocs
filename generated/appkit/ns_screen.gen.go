// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Screen] class.
type IScreen interface {
	objectivec.IObject
	// properties:
	CGDirectDisplayID() DirectDisplayID /* not a class type */
	AuxiliaryTopLeftArea() coregraphics.CGRect
	AuxiliaryTopRightArea() coregraphics.CGRect
	BackingScaleFactor() float64 /* primitive/slice/pointer. */
	ColorSpace() IColorSpace
	Depth() WindowDepth
	DeviceDescription() foundation.IDictionary /* already interface */
	DisplayUpdateGranularity() float64 /* primitive/slice/pointer. */
	Frame() coregraphics.CGRect
	LastDisplayUpdateTimestamp() float64 /* primitive/slice/pointer. */
	LocalizedName() string /* primitive/slice/pointer. */
	MaximumExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */
	MaximumFramesPerSecond() int /* primitive/slice/pointer. */
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */
	MaximumRefreshInterval() float64 /* primitive/slice/pointer. */
	MinimumRefreshInterval() float64 /* primitive/slice/pointer. */
	SafeAreaInsets() EdgeInsets /* not a class type */
	SupportedWindowDepths() NSWindowDepth
	VisibleFrame() coregraphics.CGRect
	// methods:
	BackingAlignedRectOptions(rect coregraphics.CGRect, options AlignmentOptions /* not a class type */) coregraphics.CGRect
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool /* primitive/slice/pointer. */
	ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objc.IObject /* cross-framework: DisplayLink */
}

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

// Alloc allocates a new instance without initialization.
func (sc _ScreenClass) Alloc() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a screen object representing the screen that can best represent color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deepest
func (sc _ScreenClass) DeepestScreen() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("deepestScreen"))
	return rv
}

// Returns the screen object containing the window with the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/main
func (sc _ScreenClass) MainScreen() Screen {
	rv := objc.Send[Screen](objc.ID(sc.class), objc.Sel("mainScreen"))
	return rv
}

// Returns an array of screen objects representing all of the screens available on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screens
func (sc _ScreenClass) Screens() []Screen /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Screen](objc.ID(sc.class), objc.Sel("screens"))
	return rv
}

// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("screensHaveSeparateSpaces"))
	return rv
}

// Converts a rectangle in global screen coordinates to a pixel aligned rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/backingAlignedRect(_:options:)
func (s_ Screen) BackingAlignedRectOptions(rect coregraphics.CGRect, options AlignmentOptions /* not a class type */) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}


// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}


// Converts the rectangle from the device pixel aligned coordinates system of a screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectFromBacking(_:)
func (s_ Screen) ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}


// Converts the rectangle to the device pixel aligned coordinates system of a screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectToBacking(_:)
func (s_ Screen) ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/displayLink(target:selector:)
func (s_ Screen) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objc.IObject /* cross-framework: DisplayLink */ {
	rv := objc.Send[DisplayLink](s_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}


// The CGDirectDisplayID for this screen. This will return kCGNullDirectDisplay if there isn’t one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/CGDirectDisplayID-7uvhw
func (s_ Screen) CGDirectDisplayID() DirectDisplayID /* not a class type */ {
	rv := objc.Send[DirectDisplayID](s_.ID, objc.Sel("CGDirectDisplayID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/auxiliaryTopLeftArea-4ow3p
func (s_ Screen) AuxiliaryTopLeftArea() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("auxiliaryTopLeftArea"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/auxiliaryTopRightArea-6gb2v
func (s_ Screen) AuxiliaryTopRightArea() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("auxiliaryTopRightArea"))
	return rv
}


// The backing store pixel scale factor for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/backingScaleFactor
func (s_ Screen) BackingScaleFactor() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("backingScaleFactor"))
	return rv
}


// The color space of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/colorSpace
func (s_ Screen) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](s_.ID, objc.Sel("colorSpace"))
	return rv
}


// Returns a screen object representing the screen that can best represent color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deepest
func (s_ Screen) DeepestScreen() IScreen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("deepestScreen"))
	return rv
}


// The current bit depth and colorspace information of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/depth
func (s_ Screen) Depth() WindowDepth {
	rv := objc.Send[WindowDepth](s_.ID, objc.Sel("depth"))
	return rv
}


// The device dictionary for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deviceDescription
func (s_ Screen) DeviceDescription() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("deviceDescription"))
	return rv
}


// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/displayUpdateGranularity
func (s_ Screen) DisplayUpdateGranularity() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("displayUpdateGranularity"))
	return rv
}


// The dimensions and location of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/frame
func (s_ Screen) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("frame"))
	return rv
}


// The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/lastDisplayUpdateTimestamp
func (s_ Screen) LastDisplayUpdateTimestamp() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("lastDisplayUpdateTimestamp"))
	return rv
}


// The localized name of the display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/localizedName
func (s_ Screen) LocalizedName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("localizedName"))
	return rv
}


// Returns the screen object containing the window with the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/main
func (s_ Screen) MainScreen() IScreen {
	rv := objc.Send[Screen](s_.ID, objc.Sel("mainScreen"))
	return rv
}


// The current maximum color component value for the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}


// The maximum number of frames per second that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumFramesPerSecond
func (s_ Screen) MaximumFramesPerSecond() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumFramesPerSecond"))
	return rv
}


// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumPotentialExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}


// The current maximum color component value for reference rendering to the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumReferenceExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}


// The largest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumRefreshInterval
func (s_ Screen) MaximumRefreshInterval() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumRefreshInterval"))
	return rv
}


// The shortest refresh interval that the screen supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/minimumRefreshInterval
func (s_ Screen) MinimumRefreshInterval() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumRefreshInterval"))
	return rv
}


// The distances from the screen’s edges at which content isn’t obscured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/safeAreaInsets
func (s_ Screen) SafeAreaInsets() EdgeInsets /* not a class type */ {
	rv := objc.Send[EdgeInsets](s_.ID, objc.Sel("safeAreaInsets"))
	return rv
}


// Returns an array of screen objects representing all of the screens available on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screens
func (s_ Screen) Screens() []Screen /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Screen](s_.ID, objc.Sel("screens"))
	return rv
}


// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (s_ Screen) ScreensHaveSeparateSpaces() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("screensHaveSeparateSpaces"))
	return rv
}


// A zero-terminated array of the window depths supported by the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/supportedWindowDepths
func (s_ Screen) SupportedWindowDepths() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](s_.ID, objc.Sel("supportedWindowDepths"))
	return rv
}


// The current location and dimensions of the visible screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/visibleFrame
func (s_ Screen) VisibleFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("visibleFrame"))
	return rv
}



