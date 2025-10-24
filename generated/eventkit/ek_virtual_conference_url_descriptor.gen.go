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
	// properties:
	ConferenceDetails() objc.IObject /* cross-framework: NSString */
	SetConferenceDetails(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UrlDescriptors() IEKVirtualConferenceURLDescriptor
	SetUrlDescriptors(value IEKVirtualConferenceURLDescriptor)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// Details about how users join a virtual conference, including a title and URL.
//
// To let users join a virtual conference, you provide one or more URL descriptor objects in the for the conference. Calendar uses the first URL descriptor as the preferred way for users to join a virtual conference and displays any additional links you provide in the virtual conference details.


// Details about how users join a virtual conference, including a title and URL.
//
// [Full Topic]
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



// Additional information about the conference that users may find helpful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) ConferenceDetails() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("conferenceDetails"))
	return rv
}


// Additional information about the conference that users may find helpful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) SetConferenceDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConferenceDetails:"), value)
}


// The user-visible name of the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/title
func (e_ EKVirtualConferenceURLDescriptor) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}


// The user-visible name of the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/title
func (e_ EKVirtualConferenceURLDescriptor) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), value)
}


// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) UrlDescriptors() IEKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](e_.ID, objc.Sel("urlDescriptors"))
	return rv
}


// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) SetUrlDescriptors(value IEKVirtualConferenceURLDescriptor) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrlDescriptors:"), value)
}


// The URL that users open to join a virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferenceurldescriptor/url
func (e_ EKVirtualConferenceURLDescriptor) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](e_.ID, objc.Sel("url"))
	return rv
}


// The URL that users open to join a virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferenceurldescriptor/url
func (e_ EKVirtualConferenceURLDescriptor) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrl:"), value)
}




