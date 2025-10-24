// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDIUMPCIProfile */


/* debug [class_header]: Header for MIDIUMPCIProfile */
// The class instance for the [MIDIUMPCIProfile] class.
var (
	MIDIUMPCIProfileClass     _MIDIUMPCIProfileClass
	MIDIUMPCIProfileClassOnce sync.Once
)

func getMIDIUMPCIProfileClass() _MIDIUMPCIProfileClass {
	MIDIUMPCIProfileClassOnce.Do(func() {
		MIDIUMPCIProfileClass = _MIDIUMPCIProfileClass{objc.GetClass("MIDIUMPCIProfile")}
	})
	return MIDIUMPCIProfileClass
}

type _MIDIUMPCIProfileClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDIUMPCIProfile */
// An interface definition for the [MIDIUMPCIProfile] class.
type IMIDIUMPCIProfile interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDIUMPCIProfile */
	// properties:
	EnabledChannelCount() MIDIUInteger14 /* typedef */
	FirstChannel() MIDIChannelNumber /* typedef */
	GroupOffset() MIDIUMPGroupNumber /* typedef */
	IsEnabled() bool
	Name() objc.IObject /* cross-framework: NSString */
	ProfileID() objectivec.IObject
	ProfileType() MIDICIProfileType
	TotalChannelCount() MIDIUInteger14 /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDIUMPCIProfile */
	// methods:
	SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount MIDIUInteger14 /* typedef */, error_ objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDIUMPCIProfile */
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPCIProfileClass) Alloc() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDIUMPCIProfileClass) New() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPCIProfile) Init() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPCIProfile) Autorelease() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPCIProfile creates a new MIDIUMPCIProfile instance.
func NewMIDIUMPCIProfile() MIDIUMPCIProfile {
	return getMIDIUMPCIProfileClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDIUMPCIProfile */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile
type MIDIUMPCIProfile struct {
	objectivec.Object
}

// MIDIUMPCIProfileFrom constructs a [MIDIUMPCIProfile] from an unsafe.Pointer.
func MIDIUMPCIProfileFrom(ptr unsafe.Pointer) MIDIUMPCIProfile {
	return MIDIUMPCIProfile{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDIUMPCIProfile *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDIUMPCIProfile */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDIUMPCIProfile */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDIUMPCIProfile */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/setProfileState(_:enabledChannelCount:)
func (m_ MIDIUMPCIProfile) SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount MIDIUInteger14 /* typedef */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setProfileState:enabledChannelCount:error:"), isEnabled, enabledChannelCount, error_)
	return rv
}/* debug [instance_methods/method]: SetProfileStateEnabledChannelCountError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDIUMPCIProfile */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/enabledChannelCount
func (m_ MIDIUMPCIProfile) EnabledChannelCount() MIDIUInteger14 /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("enabledChannelCount"))
	return rv
}/* debug [instance_properties/getter]: enabledChannelCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/firstChannel
func (m_ MIDIUMPCIProfile) FirstChannel() MIDIChannelNumber /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("firstChannel"))
	return rv
}/* debug [instance_properties/getter]: firstChannel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/groupOffset
func (m_ MIDIUMPCIProfile) GroupOffset() MIDIUMPGroupNumber /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("groupOffset"))
	return rv
}/* debug [instance_properties/getter]: groupOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/isEnabled
func (m_ MIDIUMPCIProfile) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/name
func (m_ MIDIUMPCIProfile) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileID
func (m_ MIDIUMPCIProfile) ProfileID() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("profileID"))
	return rv
}/* debug [instance_properties/getter]: profileID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileType
func (m_ MIDIUMPCIProfile) ProfileType() MIDICIProfileType {
	rv := objc.Send[MIDICIProfileType](m_.ID, objc.Sel("profileType"))
	return rv
}/* debug [instance_properties/getter]: profileType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/totalChannelCount
func (m_ MIDIUMPCIProfile) TotalChannelCount() MIDIUInteger14 /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("totalChannelCount"))
	return rv
}/* debug [instance_properties/getter]: totalChannelCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDIUMPCIProfile */



