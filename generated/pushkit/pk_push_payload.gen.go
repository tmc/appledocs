// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKPushPayload */


/* debug [class_header]: Header for PKPushPayload */
// The class instance for the [PushPayload] class.
var (
	PushPayloadClass     _PushPayloadClass
	PushPayloadClassOnce sync.Once
)

func getPushPayloadClass() _PushPayloadClass {
	PushPayloadClassOnce.Do(func() {
		PushPayloadClass = _PushPayloadClass{objc.GetClass("PKPushPayload")}
	})
	return PushPayloadClass
}

type _PushPayloadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PushPayload */
// An interface definition for the [PushPayload] class.
type IPushPayload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PushPayload */
	// properties:
	DictionaryPayload() objc.IObject /* cross-framework: NSDictionary */
	Type() PushType /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PushPayload */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PushPayload */
// Alloc allocates a new instance without initialization.
func (pc _PushPayloadClass) Alloc() PushPayload {
	rv := objc.Send[PushPayload](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PushPayloadClass) New() PushPayload {
	rv := objc.Send[PushPayload](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PushPayload) Init() PushPayload {
	rv := objc.Send[PushPayload](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PushPayload) Autorelease() PushPayload {
	rv := objc.Send[PushPayload](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPushPayload creates a new PushPayload instance.
func NewPushPayload() PushPayload {
	return getPushPayloadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PushPayload */
// An object that contains information about a received PushKit notification.


// An object that contains information about a received PushKit notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushPayload
type PushPayload struct {
	objectivec.Object
}

// PushPayloadFrom constructs a [PushPayload] from an unsafe.Pointer.
//
// An object that contains information about a received PushKit notification.
func PushPayloadFrom(ptr unsafe.Pointer) PushPayload {
	return PushPayload{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PushPayload *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PushPayload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PushPayload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PushPayload */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PushPayload */

// The contents of the received payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushPayload/dictionaryPayload
func (p_ PushPayload) DictionaryPayload() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("dictionaryPayload"))
	return rv
}/* debug [instance_properties/getter]: dictionaryPayload */


// The type value indicating how to interpret the payload.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushPayload/type
func (p_ PushPayload) Type() PushType /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKPushPayload */



