// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCDeviceHaptics */


/* debug [class_header]: Header for GCDeviceHaptics */
// The class instance for the [GCDeviceHaptics] class.
var (
	GCDeviceHapticsClass     _GCDeviceHapticsClass
	GCDeviceHapticsClassOnce sync.Once
)

func getGCDeviceHapticsClass() _GCDeviceHapticsClass {
	GCDeviceHapticsClassOnce.Do(func() {
		GCDeviceHapticsClass = _GCDeviceHapticsClass{objc.GetClass("GCDeviceHaptics")}
	})
	return GCDeviceHapticsClass
}

type _GCDeviceHapticsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDeviceHaptics */
// An interface definition for the [GCDeviceHaptics] class.
type IGCDeviceHaptics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCDeviceHaptics */
	// properties:
	SupportedLocalities() unsafe.Pointer
	SupportsHaptics() bool
	SetSupportsHaptics(value bool)
	GCHapticDurationInfinite() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCDeviceHaptics */
	// methods:
	CreateEngineWithLocality(locality GCHapticsLocality /* typedef */) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDeviceHaptics */
// Alloc allocates a new instance without initialization.
func (gc _GCDeviceHapticsClass) Alloc() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDeviceHapticsClass) New() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceHaptics) Init() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceHaptics) Autorelease() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceHaptics creates a new GCDeviceHaptics instance.
func NewGCDeviceHaptics() GCDeviceHaptics {
	return getGCDeviceHapticsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDeviceHaptics */
// The locations of haptic actuators on a game controller.
//
// Use this class to create a haptic engine with a specified locality. Any patterns you send to that engine play on the specified actuators.


// The locations of haptic actuators on a game controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics
type GCDeviceHaptics struct {
	objectivec.Object
}

// GCDeviceHapticsFrom constructs a [GCDeviceHaptics] from an unsafe.Pointer.
//
// The locations of haptic actuators on a game controller.
func GCDeviceHapticsFrom(ptr unsafe.Pointer) GCDeviceHaptics {
	return GCDeviceHaptics{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDeviceHaptics *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDeviceHaptics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDeviceHaptics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDeviceHaptics */

// Creates a haptics engine with the specified locality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/createEngine(withLocality:)
func (g_ GCDeviceHaptics) CreateEngineWithLocality(locality GCHapticsLocality /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("createEngineWithLocality:"), locality)
	return rv
}/* debug [instance_methods/method]: CreateEngineWithLocality */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDeviceHaptics */

// The locations of haptic actuators on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/supportedLocalities
func (g_ GCDeviceHaptics) SupportedLocalities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("supportedLocalities"))
	return rv
}/* debug [instance_properties/getter]: supportedLocalities */


// A Boolean value that indicates whether the device supports haptic event playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreHaptics/CHHapticDeviceCapability/supportsHaptics
func (g_ GCDeviceHaptics) SupportsHaptics() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsHaptics"))
	return rv
}/* debug [instance_properties/getter]: supportsHaptics */


// A Boolean value that indicates whether the device supports haptic event playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreHaptics/CHHapticDeviceCapability/supportsHaptics
func (g_ GCDeviceHaptics) SetSupportsHaptics(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsHaptics:"), value)
}/* debug [instance_properties/setter]: supportsHaptics */


// An infinite duration for a haptics event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gchapticdurationinfinite
func (g_ GCDeviceHaptics) GCHapticDurationInfinite() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("GCHapticDurationInfinite"))
	return rv
}/* debug [instance_properties/getter]: GCHapticDurationInfinite */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDeviceHaptics */



