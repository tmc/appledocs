// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
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
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.DisplayLink
	UserSpaceScaleFactor() float64
}

// An object that describes the attributes of a computer’s monitor or screen.
//
// An app may use an object to retrieve information about a screen and use this information to decide what to display on that screen. For example, an app may use the method to find out which of the available screens can best represent color and then might choose to display all of its windows on that screen. Create the application object before you use the methods in this class, so that the application object can make the necessary connection to the window system. You can make sure the application object exists by invoking the method of . If you created your app with Xcode, the application object is automatically created for you during initialization.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deepest
func (sc _ScreenClass) DeepestScreen() Screen {
	rv := objc.Send[NSScreen](objc.ID(sc.class), objc.Sel("deepestScreen"))
	return rv
}
// Returns the screen object containing the window with the keyboard focus.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/main
func (sc _ScreenClass) MainScreen() Screen {
	rv := objc.Send[NSScreen](objc.ID(sc.class), objc.Sel("mainScreen"))
	return rv
}
// Returns an array of screen objects representing all of the screens available on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screens
func (sc _ScreenClass) Screens() []Screen {
	rv := objc.Send[[]Screen](objc.ID(sc.class), objc.Sel("screens"))
	return rv
}
// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("screensHaveSeparateSpaces"))
	return rv
}
// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/canRepresent(_:)
func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/displayLink(target:selector:)
func (s_ Screen) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.DisplayLink {
	rv := objc.Send[quartzcore.DisplayLink](s_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}

// Returns the scaling factor from user space to device space on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/userSpaceScaleFactor
func (s_ Screen) UserSpaceScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("userSpaceScaleFactor"))
	return rv
}

// The CGDirectDisplayID for this screen. This will return kCGNullDirectDisplay if there isn’t one.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/CGDirectDisplayID-7uvhw
func (s_ Screen) CGDirectDisplayID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("CGDirectDisplayID"))
	return rv
}

// The color space of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/colorSpace
func (s_ Screen) ColorSpace() NSColorSpace {
	rv := objc.Send[NSColorSpace](s_.ID, objc.Sel("colorSpace"))
	return rv
}

// Returns a screen object representing the screen that can best represent color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deepest
func (s_ Screen) DeepestScreen() NSScreen {
	rv := objc.Send[NSScreen](s_.ID, objc.Sel("deepestScreen"))
	return rv
}

// The current bit depth and colorspace information of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/depth
func (s_ Screen) Depth() WindowDepth {
	rv := objc.Send[WindowDepth](s_.ID, objc.Sel("depth"))
	return rv
}

// The device dictionary for the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/deviceDescription
func (s_ Screen) DeviceDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("deviceDescription"))
	return rv
}

// The dimensions and location of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/frame
func (s_ Screen) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("frame"))
	return rv
}

// The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/lastDisplayUpdateTimestamp
func (s_ Screen) LastDisplayUpdateTimestamp() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lastDisplayUpdateTimestamp"))
	return rv
}

// The localized name of the display.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/localizedName
func (s_ Screen) LocalizedName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("localizedName"))
	return rv
}

// Returns the screen object containing the window with the keyboard focus.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/main
func (s_ Screen) MainScreen() NSScreen {
	rv := objc.Send[NSScreen](s_.ID, objc.Sel("mainScreen"))
	return rv
}

// The current maximum color component value for the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}

// The maximum number of frames per second that the screen supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumFramesPerSecond
func (s_ Screen) MaximumFramesPerSecond() int {
	rv := objc.Send[int](s_.ID, objc.Sel("maximumFramesPerSecond"))
	return rv
}

// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumPotentialExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}

// The current maximum color component value for reference rendering to the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumReferenceExtendedDynamicRangeColorComponentValue
func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}

// The shortest refresh interval that the screen supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/minimumRefreshInterval
func (s_ Screen) MinimumRefreshInterval() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minimumRefreshInterval"))
	return rv
}

// The distances from the screen’s edges at which content isn’t obscured.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/safeAreaInsets
func (s_ Screen) SafeAreaInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("safeAreaInsets"))
	return rv
}

// Returns an array of screen objects representing all of the screens available on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screens
func (s_ Screen) Screens() []Screen {
	rv := objc.Send[[]Screen](s_.ID, objc.Sel("screens"))
	return rv
}

// Returns a Boolean value indicating whether each screen can have its own set of spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (s_ Screen) ScreensHaveSeparateSpaces() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("screensHaveSeparateSpaces"))
	return rv
}

// A zero-terminated array of the window depths supported by the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScreen/supportedWindowDepths
func (s_ Screen) SupportedWindowDepths() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](s_.ID, objc.Sel("supportedWindowDepths"))
	return rv
}

// The unobscured portion of the top-left corner of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytopleftarea-uglc
func (s_ Screen) AuxiliaryTopLeftArea() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("auxiliaryTopLeftArea"))
	return rv
}


// SetAuxiliaryTopLeftArea sets the value of the auxiliaryTopLeftArea property.
// The unobscured portion of the top-left corner of the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytopleftarea-uglc
func (s_ Screen) SetAuxiliaryTopLeftArea(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuxiliaryTopLeftArea:"), value)
}

// The unobscured portion of the top-right corner of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytoprightarea-gr2n
func (s_ Screen) AuxiliaryTopRightArea() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("auxiliaryTopRightArea"))
	return rv
}


// SetAuxiliaryTopRightArea sets the value of the auxiliaryTopRightArea property.
// The unobscured portion of the top-right corner of the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytoprightarea-gr2n
func (s_ Screen) SetAuxiliaryTopRightArea(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAuxiliaryTopRightArea:"), value)
}

// The backing store pixel scale factor for the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/backingscalefactor
func (s_ Screen) BackingScaleFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("backingScaleFactor"))
	return rv
}


// SetBackingScaleFactor sets the value of the backingScaleFactor property.
// The backing store pixel scale factor for the screen.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/backingscalefactor
func (s_ Screen) SetBackingScaleFactor(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackingScaleFactor:"), value)
}

// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/displayupdategranularity
func (s_ Screen) DisplayUpdateGranularity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("displayUpdateGranularity"))
	return rv
}


// SetDisplayUpdateGranularity sets the value of the displayUpdateGranularity property.
// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/displayupdategranularity
func (s_ Screen) SetDisplayUpdateGranularity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayUpdateGranularity:"), value)
}

// The largest refresh interval that the screen supports.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumrefreshinterval
func (s_ Screen) MaximumRefreshInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("maximumRefreshInterval"))
	return rv
}


// SetMaximumRefreshInterval sets the value of the maximumRefreshInterval property.
// The largest refresh interval that the screen supports.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/maximumrefreshinterval
func (s_ Screen) SetMaximumRefreshInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaximumRefreshInterval:"), value)
}

// The current location and dimensions of the visible screen.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/visibleframe
func (s_ Screen) VisibleFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("visibleFrame"))
	return rv
}


// SetVisibleFrame sets the value of the visibleFrame property.
// The current location and dimensions of the visible screen.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscreen/visibleframe
func (s_ Screen) SetVisibleFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVisibleFrame:"), value)
}



