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
}

// A class that represents an alarm.
//
// An object represents an alarm in Event Kit. Use the and class methods to create an alarm and use the properties to set information about an alarm. In macOS Mountain Lion, you can specify an action to trigger when the alarm fires via the , , or property.
//
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
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func NewEKAlarmWithAbsoluteDate(date unsafe.Pointer) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}



// Creates and returns an alarm with a relative offset.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func NewEKAlarmWithRelativeOffset(offset foundation.TimeInterval) EKAlarm {
	rv := objc.Send[EKAlarm](objc.ID(getEKAlarmClass().class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}


// Creates and returns an alarm with an absolute date.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(absoluteDate:)
func (ec _EKAlarmClass) AlarmWithAbsoluteDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("alarmWithAbsoluteDate:"), date)
	return rv
}

// Creates and returns an alarm with a relative offset.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/init(relativeOffset:)
func (ec _EKAlarmClass) AlarmWithRelativeOffset(offset foundation.TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("alarmWithRelativeOffset:"), offset)
	return rv
}

// The absolute date for the alarm.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) AbsoluteDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("absoluteDate"))
	return rv
}


// SetAbsoluteDate sets the value of the absoluteDate property.
// The absolute date for the alarm.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/absoluteDate
func (e_ EKAlarm) SetAbsoluteDate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAbsoluteDate:"), value)
}

// The recipient of an email to send when the alarm triggers.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) EmailAddress() string {
	rv := objc.Send[string](e_.ID, objc.Sel("emailAddress"))
	return rv
}


// SetEmailAddress sets the value of the emailAddress property.
// The recipient of an email to send when the alarm triggers.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/emailAddress
func (e_ EKAlarm) SetEmailAddress(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEmailAddress:"), objc.String(value))
}

// A value indicating how a location-based alarm is triggered.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) Proximity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("proximity"))
	return rv
}


// SetProximity sets the value of the proximity property.
// A value indicating how a location-based alarm is triggered.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/proximity
func (e_ EKAlarm) SetProximity(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setProximity:"), value)
}

// The offset from the start of an event, at which the alarm fires.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) RelativeOffset() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](e_.ID, objc.Sel("relativeOffset"))
	return rv
}


// SetRelativeOffset sets the value of the relativeOffset property.
// The offset from the start of an event, at which the alarm fires.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/relativeOffset
func (e_ EKAlarm) SetRelativeOffset(value foundation.TimeInterval) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setRelativeOffset:"), value)
}

// The name of the sound to play when the alarm triggers.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SoundName() string {
	rv := objc.Send[string](e_.ID, objc.Sel("soundName"))
	return rv
}


// SetSoundName sets the value of the soundName property.
// The name of the sound to play when the alarm triggers.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/soundName
func (e_ EKAlarm) SetSoundName(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSoundName:"), objc.String(value))
}

// The location to trigger an alarm.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) StructuredLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("structuredLocation"))
	return rv
}


// SetStructuredLocation sets the value of the structuredLocation property.
// The location to trigger an alarm.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/structuredLocation
func (e_ EKAlarm) SetStructuredLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}

// The type of action to trigger when the alarm fires.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/type
func (e_ EKAlarm) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("type"))
	return rv
}

// The URL to open when the alarm triggers.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) Url() foundation.URL {
	rv := objc.Send[foundation.URL](e_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL to open when the alarm triggers.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKAlarm/url
func (e_ EKAlarm) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrl:"), value)
}


