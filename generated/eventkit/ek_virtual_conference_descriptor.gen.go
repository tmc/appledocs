// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EKVirtualConferenceDescriptor] class.
var (
	EKVirtualConferenceDescriptorClass     _EKVirtualConferenceDescriptorClass
	EKVirtualConferenceDescriptorClassOnce sync.Once
)

func getEKVirtualConferenceDescriptorClass() _EKVirtualConferenceDescriptorClass {
	EKVirtualConferenceDescriptorClassOnce.Do(func() {
		EKVirtualConferenceDescriptorClass = _EKVirtualConferenceDescriptorClass{objc.GetClass("EKVirtualConferenceDescriptor")}
	})
	return EKVirtualConferenceDescriptorClass
}

type _EKVirtualConferenceDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [EKVirtualConferenceDescriptor] class.
type IEKVirtualConferenceDescriptor interface {
	objectivec.IObject
	ConferenceDetails() string
	Title() string
	URLDescriptors() []EKVirtualConferenceURLDescriptor
}

// Details about a virtual conference that uses a custom room type.
//
// When users add events to their calendars and use one of the room types that your provider defines, EventKit requests a virtual conference descriptor from your provider. Each virtual conference descriptor contains: A user-visible name for the virtual conference One or more URLs that the users open to join the virtual conference Optional details about the conference that may be helpful to users Calendar uses the first URL that you provide as the preferred way for users to join a virtual conference and displays additional URLs as links in the virtual conference details.


// Details about a virtual conference that uses a custom room type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceDescriptor

type EKVirtualConferenceDescriptor struct {
	objectivec.Object
}

// EKVirtualConferenceDescriptorFrom constructs a [EKVirtualConferenceDescriptor] from an unsafe.Pointer.
//
// Details about a virtual conference that uses a custom room type.
func EKVirtualConferenceDescriptorFrom(ptr unsafe.Pointer) EKVirtualConferenceDescriptor {
	return EKVirtualConferenceDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EKVirtualConferenceDescriptorClass) Alloc() EKVirtualConferenceDescriptor {
	rv := objc.Send[EKVirtualConferenceDescriptor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKVirtualConferenceDescriptorClass) New() EKVirtualConferenceDescriptor {
	rv := objc.Send[EKVirtualConferenceDescriptor](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKVirtualConferenceDescriptor) Init() EKVirtualConferenceDescriptor {
	rv := objc.Send[EKVirtualConferenceDescriptor](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKVirtualConferenceDescriptor) Autorelease() EKVirtualConferenceDescriptor {
	rv := objc.Send[EKVirtualConferenceDescriptor](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKVirtualConferenceDescriptor creates a new EKVirtualConferenceDescriptor instance.
func NewEKVirtualConferenceDescriptor() EKVirtualConferenceDescriptor {
	return getEKVirtualConferenceDescriptorClass().New()
}




// Creates an object that describes a virtual conference, including a name and URL to join the conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceDescriptor/init(title:urlDescriptors:conferenceDetails:)

func NewEKVirtualConferenceDescriptorWithTitleURLDescriptorsConferenceDetails(title string, URLDescriptors []EKVirtualConferenceURLDescriptor, conferenceDetails string) EKVirtualConferenceDescriptor {
	instance := getEKVirtualConferenceDescriptorClass().Alloc()
	rv := objc.Send[EKVirtualConferenceDescriptor](instance.ID, objc.Sel("initWithTitle:URLDescriptors:conferenceDetails:"), objc.String(title), URLDescriptors, objc.String(conferenceDetails))
	rv.Autorelease()
	return rv
}



// Additional information about the conference that users may find helpful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceDescriptor/conferenceDetails

func (e_ EKVirtualConferenceDescriptor) ConferenceDetails() string {
	rv := objc.Send[string](e_.ID, objc.Sel("conferenceDetails"))
	return rv
}


// The user-visible name of the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceDescriptor/title

func (e_ EKVirtualConferenceDescriptor) Title() string {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}


// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceDescriptor/urlDescriptors

func (e_ EKVirtualConferenceDescriptor) URLDescriptors() []EKVirtualConferenceURLDescriptor {
	rv := objc.Send[[]EKVirtualConferenceURLDescriptor](e_.ID, objc.Sel("URLDescriptors"))
	return rv
}


