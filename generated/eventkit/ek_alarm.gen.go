// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [EKAlarm] class.
type IEKAlarm interface {
	IEKObject
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
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ec _EKAlarmClass) Alloc() EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates and returns an alarm with an absolute date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func NewEKAlarmWithAbsoluteDate(date objc.IObject /* cross-framework: NSDate */) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}


// Creates and returns an alarm with a relative offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func NewEKAlarmWithRelativeOffset(offset float64) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}



// Creates and returns an alarm with an absolute date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func (ec _EKAlarmClass) AlarmWithAbsoluteDate(date objc.IObject /* cross-framework: NSDate */) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}


// Creates and returns an alarm with a relative offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func (ec _EKAlarmClass) AlarmWithRelativeOffset(offset float64) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(ec.class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}


// The absolute date for the alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) AbsoluteDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("absoluteDate"))
	return rv
}


// The absolute date for the alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) SetAbsoluteDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteDate:"), value)
}


// The recipient of an email to send when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) EmailAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("emailAddress"))
	return rv
}


// The recipient of an email to send when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) SetEmailAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmailAddress:"), value)
}


// A value indicating how a location-based alarm is triggered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) Proximity() EKAlarmProximity {
	rv := objc.Send[EKAlarmProximity](e_.ID, objc.Sel("proximity"))
	return rv
}


// A value indicating how a location-based alarm is triggered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) SetProximity(value EKAlarmProximity) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProximity:"), value)
}


// The offset from the start of an event, at which the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) RelativeOffset() float64 {
	rv := objc.Send[TimeInterval](e_.ID, objc.Sel("relativeOffset"))
	return rv
}


// The offset from the start of an event, at which the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) SetRelativeOffset(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelativeOffset:"), value)
}


// The name of the sound to play when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SoundName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("soundName"))
	return rv
}


// The name of the sound to play when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SetSoundName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSoundName:"), value)
}


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) StructuredLocation() IEKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("structuredLocation"))
	return rv
}


// The location to trigger an alarm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) SetStructuredLocation(value IEKStructuredLocation) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}


// The type of action to trigger when the alarm fires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/type
func (e_ EKAlarm) Type() EKAlarmType {
	rv := objc.Send[EKAlarmType](e_.ID, objc.Sel("type"))
	return rv
}


// The URL to open when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](e_.ID, objc.Sel("url"))
	return rv
}


// The URL to open when the alarm triggers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrl:"), value)
}


