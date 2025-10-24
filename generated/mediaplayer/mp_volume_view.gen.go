// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	// properties:
	AreWirelessRoutesAvailable() bool
	SetAreWirelessRoutesAvailable(value bool)
	IsWirelessRouteActive() bool
	SetIsWirelessRouteActive(value bool)
	// methods:
}

// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
//
// Use a volume view to present the user with a slider control for setting the system audio output volume, and a button for choosing the audio output route when the option is available. When first displayed, the slider’s position reflects the current system audio output volume. As the user drags the slider, the changes update the volume view. If the user presses the device volume buttons while sound is playing, the slider moves to reflect the new volume. If there’s an Apple TV or other AirPlay-enabled device in range, the route button allows the user to choose it. If there’s only one audio output route available, the view doesn’t display the route button. The view also doesn’t display a route button when the app runs in visionOS. Use this class by embedding an instance of it in your view hierarchy. The following code snippet assumes you’ve placed an instance of the class on a view using Interface Builder, sizing and positioning it as desired to contain the volume view. Point to the instance with an outlet variable—named, in the case of this example, . You’d typically place code like that shown in the following code in your method. Listing 1. Adding a volume view to your view hierarchy When an audio output route that doesn’t support volume control, such as a car head unit, is active, the system replaces the volume slider with the route name. To instead display a volume slider as an alert, use the functions described in .


// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
//
// [Full Topic]
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



// A Boolean value indicating wireless routes are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) AreWirelessRoutesAvailable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("areWirelessRoutesAvailable"))
	return rv
}


// A Boolean value indicating wireless routes are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) SetAreWirelessRoutesAvailable(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAreWirelessRoutesAvailable:"), value)
}


// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) IsWirelessRouteActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isWirelessRouteActive"))
	return rv
}


// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) SetIsWirelessRouteActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsWirelessRouteActive:"), value)
}


