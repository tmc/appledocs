// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFAcousticFeature */


/* debug [class_header]: Header for SFAcousticFeature */
// The class instance for the [SFAcousticFeature] class.
var (
	SFAcousticFeatureClass     _SFAcousticFeatureClass
	SFAcousticFeatureClassOnce sync.Once
)

func getSFAcousticFeatureClass() _SFAcousticFeatureClass {
	SFAcousticFeatureClassOnce.Do(func() {
		SFAcousticFeatureClass = _SFAcousticFeatureClass{objc.GetClass("SFAcousticFeature")}
	})
	return SFAcousticFeatureClass
}

type _SFAcousticFeatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFAcousticFeature */
// An interface definition for the [SFAcousticFeature] class.
type ISFAcousticFeature interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFAcousticFeature */
	// properties:
	AcousticFeatureValuePerFrame() []foundation.Number
	FrameDuration() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFAcousticFeature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFAcousticFeature */
// Alloc allocates a new instance without initialization.
func (sc _SFAcousticFeatureClass) Alloc() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFAcousticFeatureClass) New() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAcousticFeature) Init() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAcousticFeature) Autorelease() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAcousticFeature creates a new SFAcousticFeature instance.
func NewSFAcousticFeature() SFAcousticFeature {
	return getSFAcousticFeatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFAcousticFeature */
// The value of a voice analysis metric.


// The value of a voice analysis metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFAcousticFeature
type SFAcousticFeature struct {
	objectivec.Object
}

// SFAcousticFeatureFrom constructs a [SFAcousticFeature] from an unsafe.Pointer.
//
// The value of a voice analysis metric.
func SFAcousticFeatureFrom(ptr unsafe.Pointer) SFAcousticFeature {
	return SFAcousticFeature{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFAcousticFeature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFAcousticFeature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFAcousticFeature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFAcousticFeature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFAcousticFeature */

// An array of feature values, one value per audio frame, corresponding to a transcript segment of recorded audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFAcousticFeature/acousticFeatureValuePerFrame-gsz5
func (s_ SFAcousticFeature) AcousticFeatureValuePerFrame() []foundation.Number {
	rv := objc.Send[[]foundation.Number](s_.ID, objc.Sel("acousticFeatureValuePerFrame"))
	return rv
}/* debug [instance_properties/getter]: acousticFeatureValuePerFrame */


// The duration of the audio frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFAcousticFeature/frameDuration
func (s_ SFAcousticFeature) FrameDuration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("frameDuration"))
	return rv
}/* debug [instance_properties/getter]: frameDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFAcousticFeature */



