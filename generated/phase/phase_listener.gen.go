// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEListener */


/* debug [class_header]: Header for PHASEListener */
// The class instance for the [PHASEListener] class.
var (
	PHASEListenerClass     _PHASEListenerClass
	PHASEListenerClassOnce sync.Once
)

func getPHASEListenerClass() _PHASEListenerClass {
	PHASEListenerClassOnce.Do(func() {
		PHASEListenerClass = _PHASEListenerClass{objc.GetClass("PHASEListener")}
	})
	return PHASEListenerClass
}

type _PHASEListenerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEListener */
// An interface definition for the [PHASEListener] class.
type IPHASEListener interface {
	IPHASEObject
	
/* debug [class_interface_properties]: Properties for PHASEListener */
	// properties:
	AutomaticHeadTrackingFlags() PHASEAutomaticHeadTrackingFlags
	SetAutomaticHeadTrackingFlags(value PHASEAutomaticHeadTrackingFlags)
	Gain() float64
	SetGain(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEListener */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEListener */
// Alloc allocates a new instance without initialization.
func (pc _PHASEListenerClass) Alloc() PHASEListener {
	rv := objc.Send[PHASEListener](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEListenerClass) New() PHASEListener {
	rv := objc.Send[PHASEListener](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEListener) Init() PHASEListener {
	rv := objc.Send[PHASEListener](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEListener) Autorelease() PHASEListener {
	rv := objc.Send[PHASEListener](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEListener creates a new PHASEListener instance.
func NewPHASEListener() PHASEListener {
	return getPHASEListenerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEListener */
// A central point of reference that defines the location within the scene that’s most audible to the user.
//
// PHASE requires an instance of this class to play ambient or spatial audio. To output sound through an ambient mixer or spatial mixer, the app adds an instance of this class to a sound event by using . For an example that demonstrates listeners, see .


// A central point of reference that defines the location within the scene that’s most audible to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener
type PHASEListener struct {
	PHASEObject
}

// PHASEListenerFrom constructs a [PHASEListener] from an unsafe.Pointer.
//
// A central point of reference that defines the location within the scene that’s most audible to the user.
func PHASEListenerFrom(ptr unsafe.Pointer) PHASEListener {
	return PHASEListener{
		PHASEObject: PHASEObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEListener */

// Creates a listener with the given engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener/init(engine:)
func NewPHASEListenerWithEngine(engine IPHASEEngine) PHASEListener {
	instance := getPHASEListenerClass().Alloc()
	rv := objc.Send[PHASEListener](instance.ID, objc.Sel("initWithEngine:"), engine)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEListenerWithEngine */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEListener */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEListener */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEListener */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEListener */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener/automaticHeadTrackingFlags
func (p_ PHASEListener) AutomaticHeadTrackingFlags() PHASEAutomaticHeadTrackingFlags {
	rv := objc.Send[PHASEAutomaticHeadTrackingFlags](p_.ID, objc.Sel("automaticHeadTrackingFlags"))
	return rv
}/* debug [instance_properties/getter]: automaticHeadTrackingFlags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener/automaticHeadTrackingFlags
func (p_ PHASEListener) SetAutomaticHeadTrackingFlags(value PHASEAutomaticHeadTrackingFlags) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticHeadTrackingFlags:"), value)
}/* debug [instance_properties/setter]: automaticHeadTrackingFlags */


// Modifies the volume of all audio playback for the listener’s mixers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener/gain
func (p_ PHASEListener) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// Modifies the volume of all audio playback for the listener’s mixers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEListener/gain
func (p_ PHASEListener) SetGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}/* debug [instance_properties/setter]: gain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEListener */


