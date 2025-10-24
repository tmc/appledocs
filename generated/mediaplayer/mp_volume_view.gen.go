// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPVolumeView */

/* debug [class_header]: Header for MPVolumeView */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VolumeView */
// An interface definition for the [VolumeView] class.
type IVolumeView interface {
	IView

	/* debug [class_interface_properties]: Properties for VolumeView */
	// properties:
	AreWirelessRoutesAvailable() bool
	SetAreWirelessRoutesAvailable(value bool)
	IsWirelessRouteActive() bool
	SetIsWirelessRouteActive(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VolumeView */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VolumeView */
// Alloc allocates a new instance without initialization.
func (vc _VolumeViewClass) Alloc() VolumeView {
	rv := objc.Send[VolumeView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VolumeView */
// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
//
// Use a volume view to present the user with a slider control for setting the system audio output volume, and a button for choosing the audio output route when the option is available. When first displayed, the slider’s position reflects the current system audio output volume. As the user drags the slider, the changes update the volume view. If the user presses the device volume buttons while sound is playing, the slider moves to reflect the new volume. If there’s an Apple TV or other AirPlay-enabled device in range, the route button allows the user to choose it. If there’s only one audio output route available, the view doesn’t display the route button. The view also doesn’t display a route button when the app runs in visionOS. Use this class by embedding an instance of it in your view hierarchy. The following code snippet assumes you’ve placed an instance of the class on a view using Interface Builder, sizing and positioning it as desired to contain the volume view. Point to the instance with an outlet variable—named, in the case of this example, . You’d typically place code like that shown in the following code in your method. Listing 1. Adding a volume view to your view hierarchy When an audio output route that doesn’t support volume control, such as a car head unit, is active, the system replaces the volume slider with the route name. To instead display a volume slider as an alert, use the functions described in .

// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView
type VolumeView struct {
	View
}

// VolumeViewFrom constructs a [VolumeView] from an unsafe.Pointer.
//
// A slider control for setting the system audio output volume, and a button for choosing the audio output route.
func VolumeViewFrom(ptr unsafe.Pointer) VolumeView {
	return VolumeView{
		View: ViewFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VolumeView */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VolumeView */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VolumeView */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VolumeView */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VolumeView */

// A Boolean value indicating wireless routes are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) AreWirelessRoutesAvailable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("areWirelessRoutesAvailable"))
	return rv
} /* debug [instance_properties/getter]: areWirelessRoutesAvailable */

// A Boolean value indicating wireless routes are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/arewirelessroutesavailable
func (v_ VolumeView) SetAreWirelessRoutesAvailable(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAreWirelessRoutesAvailable:"), value)
} /* debug [instance_properties/setter]: areWirelessRoutesAvailable */

// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) IsWirelessRouteActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isWirelessRouteActive"))
	return rv
} /* debug [instance_properties/getter]: isWirelessRouteActive */

// A Boolean value that indicates whether the wireless route is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/iswirelessrouteactive
func (v_ VolumeView) SetIsWirelessRouteActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsWirelessRouteActive:"), value)
} /* debug [instance_properties/setter]: isWirelessRouteActive */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class MPVolumeView */
