// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMusicPlayerPlayParameters */


/* debug [class_header]: Header for MPMusicPlayerPlayParameters */
// The class instance for the [MusicPlayerPlayParameters] class.
var (
	MusicPlayerPlayParametersClass     _MusicPlayerPlayParametersClass
	MusicPlayerPlayParametersClassOnce sync.Once
)

func getMusicPlayerPlayParametersClass() _MusicPlayerPlayParametersClass {
	MusicPlayerPlayParametersClassOnce.Do(func() {
		MusicPlayerPlayParametersClass = _MusicPlayerPlayParametersClass{objc.GetClass("MPMusicPlayerPlayParameters")}
	})
	return MusicPlayerPlayParametersClass
}

type _MusicPlayerPlayParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MusicPlayerPlayParameters */
// An interface definition for the [MusicPlayerPlayParameters] class.
type IMusicPlayerPlayParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MusicPlayerPlayParameters */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MusicPlayerPlayParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MusicPlayerPlayParameters */
// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerPlayParametersClass) Alloc() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicPlayerPlayParametersClass) New() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerPlayParameters) Init() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerPlayParameters) Autorelease() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerPlayParameters creates a new MusicPlayerPlayParameters instance.
func NewMusicPlayerPlayParameters() MusicPlayerPlayParameters {
	return getMusicPlayerPlayParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MusicPlayerPlayParameters */
// The MusicKit parameters that describe items to play.


// The MusicKit parameters that describe items to play.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParameters
type MusicPlayerPlayParameters struct {
	objectivec.Object
}

// MusicPlayerPlayParametersFrom constructs a [MusicPlayerPlayParameters] from an unsafe.Pointer.
//
// The MusicKit parameters that describe items to play.
func MusicPlayerPlayParametersFrom(ptr unsafe.Pointer) MusicPlayerPlayParameters {
	return MusicPlayerPlayParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MusicPlayerPlayParameters */

// Returns a new play parameters object using information from MusicKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParameters/init(dictionary:)
func NewMusicPlayerPlayParametersWithDictionary(dictionary foundation.IDictionary) MusicPlayerPlayParameters {
	instance := getMusicPlayerPlayParametersClass().Alloc()
	rv := objc.Send[MusicPlayerPlayParameters](instance.ID, objc.Sel("initWithDictionary:"), dictionary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMusicPlayerPlayParametersWithDictionary */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MusicPlayerPlayParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MusicPlayerPlayParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MusicPlayerPlayParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MusicPlayerPlayParameters */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMusicPlayerPlayParameters */


