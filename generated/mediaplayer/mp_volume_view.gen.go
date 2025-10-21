// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [VolumeView] class.
var (
	VolumeViewClass     _VolumeViewClass
	VolumeViewClassOnce sync.Once
)

func getVolumeViewClass() _VolumeViewClass {
	VolumeViewClassOnce.Do(func() {
		VolumeViewClass = _VolumeViewClass{objc.GetClass("MPVolumeView")}
	})
	return VolumeViewClass
}

type _VolumeViewClass struct {
	class objc.Class
}

// An interface definition for the [VolumeView] class.
type IVolumeView interface {
	appkit.IView
	MaximumVolumeSliderImageForState(state unsafe.Pointer) unsafe.Pointer
	MinimumVolumeSliderImageForState(state unsafe.Pointer) unsafe.Pointer
	RouteButtonImageForState(state unsafe.Pointer) unsafe.Pointer
	RouteButtonRectForBounds(bounds coregraphics.CGRect) coregraphics.CGRect
	SetMaximumVolumeSliderImageForState(image unsafe.Pointer, state unsafe.Pointer)
	SetMinimumVolumeSliderImageForState(image unsafe.Pointer, state unsafe.Pointer)
	SetRouteButtonImageForState(image unsafe.Pointer, state unsafe.Pointer)
	SetVolumeThumbImageForState(image unsafe.Pointer, state unsafe.Pointer)
	VolumeSliderRectForBounds(bounds coregraphics.CGRect) coregraphics.CGRect
	VolumeThumbImageForState(state unsafe.Pointer) unsafe.Pointer
	VolumeThumbRectForBoundsVolumeSliderRectValue(bounds coregraphics.CGRect, rect coregraphics.CGRect, value unsafe.Pointer) coregraphics.CGRect
}

// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
//
// Use a volume view to present the user with a slider control for setting the system audio output volume, and a button for choosing the audio output route when the option is available. When first displayed, the slider’s position reflects the current system audio output volume. As the user drags the slider, the changes update the volume view. If the user presses the device volume buttons while sound is playing, the slider moves to reflect the new volume. If there’s an Apple TV or other AirPlay-enabled device in range, the route button allows the user to choose it. If there’s only one audio output route available, the view doesn’t display the route button. The view also doesn’t display a route button when the app runs in visionOS. Use this class by embedding an instance of it in your view hierarchy. The following code snippet assumes you’ve placed an instance of the class on a view using Interface Builder, sizing and positioning it as desired to contain the volume view. Point to the instance with an outlet variable—named, in the case of this example, . You’d typically place code like that shown in the following code in your method. Listing 1. Adding a volume view to your view hierarchy When an audio output route that doesn’t support volume control, such as a car head unit, is active, the system replaces the volume slider with the route name. To instead display a volume slider as an alert, use the functions described in .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView
type VolumeView struct {
	appkit.View
}

// VolumeViewFrom constructs a [VolumeView] from an unsafe.Pointer.
//
// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
func VolumeViewFrom(ptr unsafe.Pointer) VolumeView {
	return VolumeView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VolumeViewClass) Alloc() VolumeView {
	rv := objc.Send[VolumeView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VolumeViewClass) New() VolumeView {
	rv := objc.Send[VolumeView](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VolumeView) Init() VolumeView {
	rv := objc.Send[VolumeView](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VolumeView) Autorelease() VolumeView {
	rv := objc.Send[VolumeView](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVolumeView creates a new VolumeView instance.
func NewVolumeView() VolumeView {
	return getVolumeViewClass().New()
}


// Returns the maximum volume image associated with the specified control state.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/maximumVolumeSliderImage(for:)
func (v_ VolumeView) MaximumVolumeSliderImageForState(state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("maximumVolumeSliderImageForState:"), state)
	return rv
}

// Returns the minimum volume image associated with the specified control state.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/minimumVolumeSliderImage(for:)
func (v_ VolumeView) MinimumVolumeSliderImageForState(state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("minimumVolumeSliderImageForState:"), state)
	return rv
}

// Returns the button image associated with the specified control state.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/routeButtonImage(for:)
func (v_ VolumeView) RouteButtonImageForState(state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("routeButtonImageForState:"), state)
	return rv
}

// Returns the drawing rectangle for the route button.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/routeButtonRect(forBounds:)
func (v_ VolumeView) RouteButtonRectForBounds(bounds coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("routeButtonRectForBounds:"), bounds)
	return rv
}

// Assigns a maximum volume slider image to the specified control states.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setMaximumVolumeSliderImage(_:for:)
func (v_ VolumeView) SetMaximumVolumeSliderImageForState(image unsafe.Pointer, state unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumVolumeSliderImage:forState:"), image, state)
}

// Assigns a minimum volume slider image to the specified control states.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setMinimumVolumeSliderImage(_:for:)
func (v_ VolumeView) SetMinimumVolumeSliderImageForState(image unsafe.Pointer, state unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMinimumVolumeSliderImage:forState:"), image, state)
}

// Assigns a button image to the specified control states.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setRouteButtonImage(_:for:)
func (v_ VolumeView) SetRouteButtonImageForState(image unsafe.Pointer, state unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRouteButtonImage:forState:"), image, state)
}

// Assigns a thumb image to the specified control states.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setVolumeThumbImage(_:for:)
func (v_ VolumeView) SetVolumeThumbImageForState(image unsafe.Pointer, state unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVolumeThumbImage:forState:"), image, state)
}

// Returns the drawing rectangle for the slider’s track.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeSliderRect(forBounds:)
func (v_ VolumeView) VolumeSliderRectForBounds(bounds coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("volumeSliderRectForBounds:"), bounds)
	return rv
}

// Returns the thumb image associated with the specified control state.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeThumbImage(for:)
func (v_ VolumeView) VolumeThumbImageForState(state unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("volumeThumbImageForState:"), state)
	return rv
}

// Returns the drawing rectangle for the volume slider’s thumb image.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeThumbRect(forBounds:volumeSliderRect:value:)
func (v_ VolumeView) VolumeThumbRectForBoundsVolumeSliderRectValue(bounds coregraphics.CGRect, rect coregraphics.CGRect, value unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("volumeThumbRectForBounds:volumeSliderRect:value:"), bounds, rect, value)
	return rv
}

// A Boolean value indicating wireless routes are available.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/areWirelessRoutesAvailable
func (v_ VolumeView) WirelessRoutesAvailable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wirelessRoutesAvailable"))
	return rv
}

// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/isWirelessRouteActive
func (v_ VolumeView) WirelessRouteActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wirelessRouteActive"))
	return rv
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsRouteButton
func (v_ VolumeView) ShowsRouteButton() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsRouteButton"))
	return rv
}


// SetShowsRouteButton sets the value of the showsRouteButton property.
// A Boolean value that indicates whether the route button is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsRouteButton
func (v_ VolumeView) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsRouteButton:"), value)
}

// A Boolean value that indicates the volume slider is visible in the volume view.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsVolumeSlider
func (v_ VolumeView) ShowsVolumeSlider() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsVolumeSlider"))
	return rv
}


// SetShowsVolumeSlider sets the value of the showsVolumeSlider property.
// A Boolean value that indicates the volume slider is visible in the volume view.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsVolumeSlider
func (v_ VolumeView) SetShowsVolumeSlider(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsVolumeSlider:"), value)
}

// The image used to designate the European Union volume limit.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeWarningSliderImage
func (v_ VolumeView) VolumeWarningSliderImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("volumeWarningSliderImage"))
	return rv
}


// SetVolumeWarningSliderImage sets the value of the volumeWarningSliderImage property.
// The image used to designate the European Union volume limit.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeWarningSliderImage
func (v_ VolumeView) SetVolumeWarningSliderImage(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVolumeWarningSliderImage:"), value)
}

// A Boolean value indicating wireless routes are available.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) AreWirelessRoutesAvailable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("areWirelessRoutesAvailable"))
	return rv
}


// SetAreWirelessRoutesAvailable sets the value of the areWirelessRoutesAvailable property.
// A Boolean value indicating wireless routes are available.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) SetAreWirelessRoutesAvailable(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAreWirelessRoutesAvailable:"), value)
}

// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) IsWirelessRouteActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isWirelessRouteActive"))
	return rv
}


// SetIsWirelessRouteActive sets the value of the isWirelessRouteActive property.
// A Boolean value that indicates whether the wireless route is active.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) SetIsWirelessRouteActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsWirelessRouteActive:"), value)
}




