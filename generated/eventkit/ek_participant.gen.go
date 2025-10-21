// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EKParticipant] class.
var (
	EKParticipantClass     _EKParticipantClass
	EKParticipantClassOnce sync.Once
)

func getEKParticipantClass() _EKParticipantClass {
	EKParticipantClassOnce.Do(func() {
		EKParticipantClass = _EKParticipantClass{objc.GetClass("EKParticipant")}
	})
	return EKParticipantClass
}

type _EKParticipantClass struct {
	class objc.Class
}

// An interface definition for the [EKParticipant] class.
type IEKParticipant interface {
	IEKObject
	ABPersonInAddressBook(addressBook unsafe.Pointer) unsafe.Pointer
	ABRecordWithAddressBook(addressBook unsafe.Pointer) unsafe.Pointer
}

// A class that represents person, group, or room invited to a calendar event.
//
// Do not create objects directly. Instead, use the property attendees on to return an array of objects. EventKit cannot add participants to an event nor change participant information. Use the properties in this class to get information about a participant. A participant can be a person, group, room, or other resource.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant
type EKParticipant struct {
	EKObject
}

// EKParticipantFrom constructs a [EKParticipant] from an unsafe.Pointer.
//
// A class that represents person, group, or room invited to a calendar event.
func EKParticipantFrom(ptr unsafe.Pointer) EKParticipant {
	return EKParticipant{
		EKObject: EKObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EKParticipantClass) Alloc() EKParticipant {
	rv := objc.Send[EKParticipant](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKParticipantClass) New() EKParticipant {
	rv := objc.Send[EKParticipant](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKParticipant) Init() EKParticipant {
	rv := objc.Send[EKParticipant](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKParticipant) Autorelease() EKParticipant {
	rv := objc.Send[EKParticipant](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKParticipant creates a new EKParticipant instance.
func NewEKParticipant() EKParticipant {
	return getEKParticipantClass().New()
}


// Returns the address book record that represents the participant.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/abPerson(in:)
func (e_ EKParticipant) ABPersonInAddressBook(addressBook unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("ABPersonInAddressBook:"), addressBook)
	return rv
}

// Returns the address book record that represents the participant.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/abRecord(with:)
func (e_ EKParticipant) ABRecordWithAddressBook(addressBook unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("ABRecordWithAddressBook:"), addressBook)
	return rv
}

// A predicate to use with the Contacts framework to retrieve the corresponding contact instance.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/contactPredicate
func (e_ EKParticipant) ContactPredicate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("contactPredicate"))
	return rv
}

// A Boolean value indicating whether this participant represents the owner of this account.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/isCurrentUser
func (e_ EKParticipant) CurrentUser() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("currentUser"))
	return rv
}

// The participant’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/name
func (e_ EKParticipant) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("name"))
	return rv
}

// The participant’s role in the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/participantRole
func (e_ EKParticipant) ParticipantRole() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("participantRole"))
	return rv
}

// The participant’s attendance status.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/participantStatus
func (e_ EKParticipant) ParticipantStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("participantStatus"))
	return rv
}

// The participant’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/participantType
func (e_ EKParticipant) ParticipantType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("participantType"))
	return rv
}

// The URL representing this participant.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKParticipant/url
func (e_ EKParticipant) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("URL"))
	return rv
}



