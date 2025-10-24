// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUAudioUnit */


/* debug [class_header]: Header for AUAudioUnit */
// The class instance for the [AudioUnit] class.
var (
	AudioUnitClass     _AudioUnitClass
	AudioUnitClassOnce sync.Once
)

func getAudioUnitClass() _AudioUnitClass {
	AudioUnitClassOnce.Do(func() {
		AudioUnitClass = _AudioUnitClass{objc.GetClass("AUAudioUnit")}
	})
	return AudioUnitClass
}

type _AudioUnitClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnit */
// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnit */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnit */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnit */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitClass) New() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnit) Init() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnit) Autorelease() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnit creates a new AudioUnit instance.
func NewAudioUnit() AudioUnit {
	return getAudioUnitClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnit */
// A parent class referenced by other AVFAudio classes.


// A parent class referenced by other AVFAudio classes. [Full Topic]
type AudioUnit struct {
	objectivec.Object
}

// AudioUnitFrom constructs a [AudioUnit] from an unsafe.Pointer.
//
// A parent class referenced by other AVFAudio classes.
func AudioUnitFrom(ptr unsafe.Pointer) AudioUnit {
	return AudioUnit{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnit *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnit */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnit */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnit */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUAudioUnit */



