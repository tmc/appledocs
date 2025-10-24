// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKVirtualConferenceURLDescriptor */


/* debug [class_header]: Header for EKVirtualConferenceURLDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKVirtualConferenceURLDescriptor */
// An interface definition for the [EKVirtualConferenceURLDescriptor] class.
type IEKVirtualConferenceURLDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EKVirtualConferenceURLDescriptor */
	// properties:
	Title() objc.IObject /* cross-framework: NSString */
	URL() objc.IObject /* cross-framework: NSURL */
	ConferenceDetails() objc.IObject /* cross-framework: NSString */
	SetConferenceDetails(value objc.IObject /* cross-framework: NSString */)
	UrlDescriptors() IEKVirtualConferenceURLDescriptor
	SetUrlDescriptors(value IEKVirtualConferenceURLDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKVirtualConferenceURLDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKVirtualConferenceURLDescriptor */
// Alloc allocates a new instance without initialization.
func (ec _EKVirtualConferenceURLDescriptorClass) Alloc() EKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKVirtualConferenceURLDescriptor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKVirtualConferenceURLDescriptor */

// Creates a URL descriptor with the given title and URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/init(title:url:)
func NewEKVirtualConferenceURLDescriptorWithTitleURL(title objc.IObject /* cross-framework: NSString */, URL objc.IObject /* cross-framework: NSURL */) EKVirtualConferenceURLDescriptor {
	instance := getEKVirtualConferenceURLDescriptorClass().Alloc()
	rv := objc.Send[EKVirtualConferenceURLDescriptor](instance.ID, objc.Sel("initWithTitle:URL:"), title, URL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKVirtualConferenceURLDescriptorWithTitleURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKVirtualConferenceURLDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKVirtualConferenceURLDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKVirtualConferenceURLDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKVirtualConferenceURLDescriptor */

// The user-visible name of a room where virtual conferences take place, such as Personal Room or Team Room.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/title
func (e_ EKVirtualConferenceURLDescriptor) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The URL that users open to join a virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceURLDescriptor/url
func (e_ EKVirtualConferenceURLDescriptor) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](e_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// Additional information about the conference that users may find helpful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) ConferenceDetails() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("conferenceDetails"))
	return rv
}/* debug [instance_properties/getter]: conferenceDetails */


// Additional information about the conference that users may find helpful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/conferencedetails
func (e_ EKVirtualConferenceURLDescriptor) SetConferenceDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setConferenceDetails:"), value)
}/* debug [instance_properties/setter]: conferenceDetails */


// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) UrlDescriptors() IEKVirtualConferenceURLDescriptor {
	rv := objc.Send[EKVirtualConferenceURLDescriptor](e_.ID, objc.Sel("urlDescriptors"))
	return rv
}/* debug [instance_properties/getter]: urlDescriptors */


// An array that contains objects with details about where to join the virtual conference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekvirtualconferencedescriptor/urldescriptors
func (e_ EKVirtualConferenceURLDescriptor) SetUrlDescriptors(value IEKVirtualConferenceURLDescriptor) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUrlDescriptors:"), value)
}/* debug [instance_properties/setter]: urlDescriptors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKVirtualConferenceURLDescriptor */


