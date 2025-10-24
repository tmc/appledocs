// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EKVirtualConferenceRoomTypeDescriptor */


/* debug [class_header]: Header for EKVirtualConferenceRoomTypeDescriptor */
// The class instance for the [EKVirtualConferenceRoomTypeDescriptor] class.
var (
	EKVirtualConferenceRoomTypeDescriptorClass     _EKVirtualConferenceRoomTypeDescriptorClass
	EKVirtualConferenceRoomTypeDescriptorClassOnce sync.Once
)

func getEKVirtualConferenceRoomTypeDescriptorClass() _EKVirtualConferenceRoomTypeDescriptorClass {
	EKVirtualConferenceRoomTypeDescriptorClassOnce.Do(func() {
		EKVirtualConferenceRoomTypeDescriptorClass = _EKVirtualConferenceRoomTypeDescriptorClass{objc.GetClass("EKVirtualConferenceRoomTypeDescriptor")}
	})
	return EKVirtualConferenceRoomTypeDescriptorClass
}

type _EKVirtualConferenceRoomTypeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKVirtualConferenceRoomTypeDescriptor */
// An interface definition for the [EKVirtualConferenceRoomTypeDescriptor] class.
type IEKVirtualConferenceRoomTypeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EKVirtualConferenceRoomTypeDescriptor */
	// properties:
	Identifier() EKVirtualConferenceRoomTypeIdentifier /* typedef */
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKVirtualConferenceRoomTypeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKVirtualConferenceRoomTypeDescriptor */
// Alloc allocates a new instance without initialization.
func (ec _EKVirtualConferenceRoomTypeDescriptorClass) Alloc() EKVirtualConferenceRoomTypeDescriptor {
	rv := objc.Send[EKVirtualConferenceRoomTypeDescriptor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKVirtualConferenceRoomTypeDescriptorClass) New() EKVirtualConferenceRoomTypeDescriptor {
	rv := objc.Send[EKVirtualConferenceRoomTypeDescriptor](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKVirtualConferenceRoomTypeDescriptor) Init() EKVirtualConferenceRoomTypeDescriptor {
	rv := objc.Send[EKVirtualConferenceRoomTypeDescriptor](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKVirtualConferenceRoomTypeDescriptor) Autorelease() EKVirtualConferenceRoomTypeDescriptor {
	rv := objc.Send[EKVirtualConferenceRoomTypeDescriptor](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKVirtualConferenceRoomTypeDescriptor creates a new EKVirtualConferenceRoomTypeDescriptor instance.
func NewEKVirtualConferenceRoomTypeDescriptor() EKVirtualConferenceRoomTypeDescriptor {
	return getEKVirtualConferenceRoomTypeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKVirtualConferenceRoomTypeDescriptor */
// Details about a room where virtual conferences take place.
//
// To present a list of rooms where a virtual conference takes place, your virtual conference provider creates one or more room type descriptors. Each descriptor contains a user-visible title and an identifier of your choosing. When users create events using one of the rooms you provide, EventKit calls and passes the room’s identifier.


// Details about a room where virtual conferences take place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceRoomTypeDescriptor
type EKVirtualConferenceRoomTypeDescriptor struct {
	objectivec.Object
}

// EKVirtualConferenceRoomTypeDescriptorFrom constructs a [EKVirtualConferenceRoomTypeDescriptor] from an unsafe.Pointer.
//
// Details about a room where virtual conferences take place.
func EKVirtualConferenceRoomTypeDescriptorFrom(ptr unsafe.Pointer) EKVirtualConferenceRoomTypeDescriptor {
	return EKVirtualConferenceRoomTypeDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKVirtualConferenceRoomTypeDescriptor */

// Creates an object that describes a location where a virtual conference takes place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceRoomTypeDescriptor/init(title:identifier:)
func NewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier(title objc.IObject /* cross-framework: NSString */, identifier EKVirtualConferenceRoomTypeIdentifier /* typedef */) EKVirtualConferenceRoomTypeDescriptor {
	instance := getEKVirtualConferenceRoomTypeDescriptorClass().Alloc()
	rv := objc.Send[EKVirtualConferenceRoomTypeDescriptor](instance.ID, objc.Sel("initWithTitle:identifier:"), title, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEKVirtualConferenceRoomTypeDescriptorWithTitleIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKVirtualConferenceRoomTypeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKVirtualConferenceRoomTypeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKVirtualConferenceRoomTypeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKVirtualConferenceRoomTypeDescriptor */

// A unique string you choose that identifies the room.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceRoomTypeDescriptor/identifier
func (e_ EKVirtualConferenceRoomTypeDescriptor) Identifier() EKVirtualConferenceRoomTypeIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The user-visible name of a room where virtual conferences take place, such as Personal Room or Team Room.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKVirtualConferenceRoomTypeDescriptor/title
func (e_ EKVirtualConferenceRoomTypeDescriptor) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKVirtualConferenceRoomTypeDescriptor */


