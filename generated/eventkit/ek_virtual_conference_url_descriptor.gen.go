// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EKVirtualConferenceURLDescriptor] class.
var (
	EKVirtualConferenceURLDescriptorClass     _EKVirtualConferenceURLDescriptorClass
	EKVirtualConferenceURLDescriptorClassOnce sync.Once
)

func getEKVirtualConferenceURLDescriptorClass() _EKVirtualConferenceURLDescriptorClass {
	EKVirtualConferenceURLDescriptorClassOnce.Do(func() {
		EKVirtualConferenceURLDescriptorClass = _EKVirtualConferenceURLDescriptorClass{objc.GetClass("EKVirtualConferenceURLDescriptor")}
	})
	return EKVirtualConferenceURLDescriptorClass
}

type _EKVirtualConferenceURLDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [EKVirtualConferenceURLDescriptor] class.
type IEKVirtualConferenceURLDescriptor interface {
	objectivec.IObject
}

// Details about how users join a virtual conference, including a title and URL.
//
// To let users join a virtual conference, you provide one or more URL descriptor objects in the for the conference. Calendar uses the first URL descriptor as the preferred way for users to join a virtual conference and displays any additional links you provide in the virtual conference details.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor
type EKVirtualConferenceURLDescriptor struct {
	objectivec.Object
}

// EKVirtualConferenceURLDescriptorFrom constructs a [EKVirtualConferenceURLDescriptor] from an unsafe.Pointer.
//
// Details about how users join a virtual conference, including a title and URL.
func EKVirtualConferenceURLDescriptorFrom(ptr unsafe.Pointer) EKVirtualConferenceURLDescriptor {
	return EKVirtualConferenceURLDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EKVirtualConferenceURLDescriptorClass) Alloc() EKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKVirtualConferenceURLDescriptorClass) New() EKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKVirtualConferenceURLDescriptor) Init() EKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKVirtualConferenceURLDescriptor) Autorelease() EKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKVirtualConferenceURLDescriptor creates a new EKVirtualConferenceURLDescriptor instance.
func NewEKVirtualConferenceURLDescriptor() EKVirtualConferenceURLDescriptor {
	return getEKVirtualConferenceURLDescriptorClass().New()
}




// Creates a URL descriptor with the given title and URL.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/init(title:url:)
func NewEKVirtualConferenceURLDescriptorWithTitleURL(title string, URL foundation.URL) EKVirtualConferenceURLDescriptor {
	instance := getEKVirtualConferenceURLDescriptorClass().Alloc()
	rv := objc.Send[EKVirtualConferenceURLDescriptor](instance.ID, objc.Sel("initWithTitle:URL:"), objc.String(title), URL)
	rv.Autorelease()
	return rv
}


// The user-visible name of a room where virtual conferences take place, such as Personal Room or Team Room.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/title
func (e_ EKVirtualConferenceURLDescriptor) Title() string {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}

// The URL that users open to join a virtual conference.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/url
func (e_ EKVirtualConferenceURLDescriptor) URL() foundation.URL {
	rv := objc.Send[foundation.URL](e_.ID, objc.Sel("URL"))
	return rv
}

// Additional information about the conference that users may find helpful.
//
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) ConferenceDetails() string {
	rv := objc.Send[string](e_.ID, objc.Sel("conferenceDetails"))
	return rv
}


// SetConferenceDetails sets the value of the conferenceDetails property.
// Additional information about the conference that users may find helpful.

//
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) SetConferenceDetails(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConferenceDetails:"), objc.String(value))
}

// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) UrlDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("urlDescriptors"))
	return rv
}


// SetUrlDescriptors sets the value of the urlDescriptors property.
// An array that contains objects with details about where to join the virtual conference.

//
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) SetUrlDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrlDescriptors:"), value)
}


