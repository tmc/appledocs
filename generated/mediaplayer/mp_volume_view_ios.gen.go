//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// iOS-only methods for VolumeView


// Returns the maximum volume image associated with the specified control state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/maximumVolumeSliderImage(for:)
func (v_ VolumeView) MaximumVolumeSliderImageForState(state ControlState /* not a class type */) objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](v_.ID, objc.Sel("maximumVolumeSliderImageForState:"), state)
	return rv
}

// Returns the minimum volume image associated with the specified control state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/minimumVolumeSliderImage(for:)
func (v_ VolumeView) MinimumVolumeSliderImageForState(state ControlState /* not a class type */) objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](v_.ID, objc.Sel("minimumVolumeSliderImageForState:"), state)
	return rv
}

// Assigns a maximum volume slider image to the specified control states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setMaximumVolumeSliderImage(_:for:)
func (v_ VolumeView) SetMaximumVolumeSliderImageForState(image objc.IObject /* cross-framework: Image */, state ControlState /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMaximumVolumeSliderImage:forState:"), image, state)
}

// Assigns a minimum volume slider image to the specified control states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setMinimumVolumeSliderImage(_:for:)
func (v_ VolumeView) SetMinimumVolumeSliderImageForState(image objc.IObject /* cross-framework: Image */, state ControlState /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMinimumVolumeSliderImage:forState:"), image, state)
}

// Assigns a thumb image to the specified control states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/setVolumeThumbImage(_:for:)
func (v_ VolumeView) SetVolumeThumbImageForState(image objc.IObject /* cross-framework: Image */, state ControlState /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVolumeThumbImage:forState:"), image, state)
}

// Returns the drawing rectangle for the slider’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeSliderRect(forBounds:)
func (v_ VolumeView) VolumeSliderRectForBounds(bounds objc.IObject /* cross-framework: Rect */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](v_.ID, objc.Sel("volumeSliderRectForBounds:"), bounds)
	return rv
}

// Returns the thumb image associated with the specified control state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeThumbImage(for:)
func (v_ VolumeView) VolumeThumbImageForState(state ControlState /* not a class type */) objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](v_.ID, objc.Sel("volumeThumbImageForState:"), state)
	return rv
}

// Returns the drawing rectangle for the volume slider’s thumb image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeThumbRect(forBounds:volumeSliderRect:value:)
func (v_ VolumeView) VolumeThumbRectForBoundsVolumeSliderRectValue(bounds objc.IObject /* cross-framework: Rect */, rect objc.IObject /* cross-framework: Rect */, value float32) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](v_.ID, objc.Sel("volumeThumbRectForBounds:volumeSliderRect:value:"), bounds, rect, value)
	return rv
}

// iOS-only properties

// A Boolean value indicating wireless routes are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/areWirelessRoutesAvailable
func (v_ VolumeView) WirelessRoutesAvailable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wirelessRoutesAvailable"))
	return rv
}

// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/isWirelessRouteActive
func (v_ VolumeView) WirelessRouteActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wirelessRouteActive"))
	return rv
}

// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsRouteButton
func (v_ VolumeView) ShowsRouteButton() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsRouteButton"))
	return rv
}
func (v_ VolumeView) SetShowsRouteButton(value bool) {
	v_.ID.Send(objc.RegisterName("setShowsRouteButton:"), value)
}

// A Boolean value that indicates the volume slider is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/showsVolumeSlider
func (v_ VolumeView) ShowsVolumeSlider() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsVolumeSlider"))
	return rv
}
func (v_ VolumeView) SetShowsVolumeSlider(value bool) {
	v_.ID.Send(objc.RegisterName("setShowsVolumeSlider:"), value)
}

// The image used to designate the European Union volume limit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView/volumeWarningSliderImage
func (v_ VolumeView) VolumeWarningSliderImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](v_.ID, objc.Sel("volumeWarningSliderImage"))
	return rv
}
func (v_ VolumeView) SetVolumeWarningSliderImage(value objc.IObject /* cross-framework: Image */) {
	v_.ID.Send(objc.RegisterName("setVolumeWarningSliderImage:"), value)
}






