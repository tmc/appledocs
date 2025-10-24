// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIProfile */


/* debug [class_header]: Header for MIDICIProfile */
// The class instance for the [MIDICIProfile] class.
var (
	MIDICIProfileClass     _MIDICIProfileClass
	MIDICIProfileClassOnce sync.Once
)

func getMIDICIProfileClass() _MIDICIProfileClass {
	MIDICIProfileClassOnce.Do(func() {
		MIDICIProfileClass = _MIDICIProfileClass{objc.GetClass("MIDICIProfile")}
	})
	return MIDICIProfileClass
}

type _MIDICIProfileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIProfile */
// An interface definition for the [MIDICIProfile] class.
type IMIDICIProfile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIProfile */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	ProfileID() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIProfile */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIProfile */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIProfileClass) Alloc() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDICIProfileClass) New() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIProfile) Init() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIProfile) Autorelease() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIProfile creates a new MIDICIProfile instance.
func NewMIDICIProfile() MIDICIProfile {
	return getMIDICIProfileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIProfile */
// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.


// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile
type MIDICIProfile struct {
	objectivec.Object
}

// MIDICIProfileFrom constructs a [MIDICIProfile] from an unsafe.Pointer.
//
// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.
func MIDICIProfileFrom(ptr unsafe.Pointer) MIDICIProfile {
	return MIDICIProfile{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIProfile */

// Creates a MIDI profile for the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:)
func NewMIDICIProfileWithData(data objc.IObject /* cross-framework: NSData */) MIDICIProfile {
	instance := getMIDICIProfileClass().Alloc()
	rv := objc.Send[MIDICIProfile](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIProfileWithData */


// Creates a named MIDI profile for the specified data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:name:)
func NewMIDICIProfileWithDataName(data objc.IObject /* cross-framework: NSData */, inName objc.IObject /* cross-framework: NSString */) MIDICIProfile {
	instance := getMIDICIProfileClass().Alloc()
	rv := objc.Send[MIDICIProfile](instance.ID, objc.Sel("initWithData:name:"), data, inName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIProfileWithDataName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIProfile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIProfile */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIProfile */

// A string that describes the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/name
func (m_ MIDICIProfile) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The unique five-byte profile identifier that represents the profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/profileID
func (m_ MIDICIProfile) ProfileID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("profileID"))
	return rv
}/* debug [instance_properties/getter]: profileID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIProfile */


