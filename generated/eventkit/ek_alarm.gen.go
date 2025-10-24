// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKAlarm */


/* debug [class_header]: Header for EKAlarm */
// The class instance for the [EKAlarm] class.
var (
	EKAlarmClass     _EKAlarmClass
	EKAlarmClassOnce sync.Once
)

func getEKAlarmClass() _EKAlarmClass {
	EKAlarmClassOnce.Do(func() {
		EKAlarmClass = _EKAlarmClass{objc.GetClass("EKAlarm")}
	})
	return EKAlarmClass
}

type _EKAlarmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKAlarm */
// An interface definition for the [EKAlarm] class.
type IEKAlarm interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKAlarm */
	// properties:
	AbsoluteDate() objc.IObject /* cross-framework: NSDate */
	SetAbsoluteDate(value objc.IObject /* cross-framework: NSDate */)
	EmailAddress() objc.IObject /* cross-framework: NSString */
	SetEmailAddress(value objc.IObject /* cross-framework: NSString */)
	Proximity() EKAlarmProximity
	SetProximity(value EKAlarmProximity)
	RelativeOffset() float64
	SetRelativeOffset(value float64)
	SoundName() objc.IObject /* cross-framework: NSString */
	SetSoundName(value objc.IObject /* cross-framework: NSString */)
	StructuredLocation() IEKStructuredLocation
	SetStructuredLocation(value IEKStructuredLocation)
	Type() EKAlarmType
	Url() objc.IObject /* cross-framework: NSURL */
	SetUrl(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKAlarm */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKAlarm */
// Alloc allocates a new instance without initialization.
func (ec _EKAlarmClass) Alloc() EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKAlarmClass) New() EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKAlarm) Init() EKAlarm {
	rv := objc.Send[EKAlarm](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKAlarm) Autorelease() EKAlarm {
	rv := objc.Send[EKAlarm](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKAlarm creates a new EKAlarm instance.
func NewEKAlarm() EKAlarm {
	return getEKAlarmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKAlarm */
// A class that represents an alarm.
//
// An object represents an alarm in Event Kit. Use the and class methods to create an alarm and use the properties to set information about an alarm. In macOS Mountain Lion, you can specify an action to trigger when the alarm fires via the , , or property.


// A class that represents an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm
type EKAlarm struct {
	EKObject
}

// EKAlarmFrom constructs a [EKAlarm] from an unsafe.Pointer.
//
// A class that represents an alarm.
func EKAlarmFrom(ptr unsafe.Pointer) EKAlarm {
	return EKAlarm{
		EKObject: EKObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKAlarm */

// Creates and returns an alarm with an absolute date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func NewEKAlarmWithAbsoluteDate(date objc.IObject /* cross-framework: NSDate */) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}/* debug [class_init_methods/constructor]: NewEKAlarmWithAbsoluteDate */


// Creates and returns an alarm with a relative offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func NewEKAlarmWithRelativeOffset(offset float64) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}/* debug [class_init_methods/constructor]: NewEKAlarmWithRelativeOffset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKAlarm */

// Creates and returns an alarm with an absolute date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func (ec _EKAlarmClass) AlarmWithAbsoluteDate(date objc.IObject /* cross-framework: NSDate */) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AlarmWithAbsoluteDate) */


// Creates and returns an alarm with a relative offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func (ec _EKAlarmClass) AlarmWithRelativeOffset(offset float64) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AlarmWithRelativeOffset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKAlarm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKAlarm */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKAlarm */

// The absolute date for the alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) AbsoluteDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("absoluteDate"))
	return rv
}/* debug [instance_properties/getter]: absoluteDate */


// The absolute date for the alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) SetAbsoluteDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteDate:"), value)
}/* debug [instance_properties/setter]: absoluteDate */


// The recipient of an email to send when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) EmailAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("emailAddress"))
	return rv
}/* debug [instance_properties/getter]: emailAddress */


// The recipient of an email to send when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) SetEmailAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmailAddress:"), value)
}/* debug [instance_properties/setter]: emailAddress */


// A value indicating how a location-based alarm is triggered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) Proximity() EKAlarmProximity {
	rv := objc.Send[EKAlarmProximity](e_.ID, objc.Sel("proximity"))
	return rv
}/* debug [instance_properties/getter]: proximity */


// A value indicating how a location-based alarm is triggered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) SetProximity(value EKAlarmProximity) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProximity:"), value)
}/* debug [instance_properties/setter]: proximity */


// The offset from the start of an event, at which the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) RelativeOffset() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("relativeOffset"))
	return rv
}/* debug [instance_properties/getter]: relativeOffset */


// The offset from the start of an event, at which the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) SetRelativeOffset(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelativeOffset:"), value)
}/* debug [instance_properties/setter]: relativeOffset */


// The name of the sound to play when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SoundName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("soundName"))
	return rv
}/* debug [instance_properties/getter]: soundName */


// The name of the sound to play when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SetSoundName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSoundName:"), value)
}/* debug [instance_properties/setter]: soundName */


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) StructuredLocation() IEKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("structuredLocation"))
	return rv
}/* debug [instance_properties/getter]: structuredLocation */


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) SetStructuredLocation(value IEKStructuredLocation) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}/* debug [instance_properties/setter]: structuredLocation */


// The type of action to trigger when the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/type
func (e_ EKAlarm) Type() EKAlarmType {
	rv := objc.Send[EKAlarmType](e_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The URL to open when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](e_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL to open when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKAlarm */


