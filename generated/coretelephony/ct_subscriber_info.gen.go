// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTSubscriberInfo */


/* debug [class_header]: Header for CTSubscriberInfo */
// The class instance for the [SubscriberInfo] class.
var (
	SubscriberInfoClass     _SubscriberInfoClass
	SubscriberInfoClassOnce sync.Once
)

func getSubscriberInfoClass() _SubscriberInfoClass {
	SubscriberInfoClassOnce.Do(func() {
		SubscriberInfoClass = _SubscriberInfoClass{objc.GetClass("CTSubscriberInfo")}
	})
	return SubscriberInfoClass
}

type _SubscriberInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SubscriberInfo */
// An interface definition for the [SubscriberInfo] class.
type ISubscriberInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SubscriberInfo */
	// properties:
	CarrierToken() foundation.Data
	SetCarrierToken(value foundation.Data)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SubscriberInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SubscriberInfo */
// Alloc allocates a new instance without initialization.
func (sc _SubscriberInfoClass) Alloc() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubscriberInfoClass) New() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SubscriberInfo) Init() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SubscriberInfo) Autorelease() SubscriberInfo {
	rv := objc.Send[SubscriberInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubscriberInfo creates a new SubscriberInfo instance.
func NewSubscriberInfo() SubscriberInfo {
	return getSubscriberInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SubscriberInfo */
// An object that provides an array of cellular network subscribers.
//
// Use the instances provided by this class to identify individual subscribers by their or properties.


// An object that provides an array of cellular network subscribers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriberInfo
type SubscriberInfo struct {
	objectivec.Object
}

// SubscriberInfoFrom constructs a [SubscriberInfo] from an unsafe.Pointer.
//
// An object that provides an array of cellular network subscribers.
func SubscriberInfoFrom(ptr unsafe.Pointer) SubscriberInfo {
	return SubscriberInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SubscriberInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SubscriberInfo */

// Returns the cellular network subscribers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriberInfo/subscriber()
func (sc _SubscriberInfoClass) Subscriber() ISubscriber {
	rv := objc.Send[Subscriber](objc.ID(sc.class), objc.Sel("subscriber"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Subscriber) */


// Returns the cellular network subscribers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriberInfo/subscribers()
func (sc _SubscriberInfoClass) Subscribers() []Subscriber {
	rv := objc.Send[[]Subscriber](objc.ID(sc.class), objc.Sel("subscribers"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Subscribers) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SubscriberInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SubscriberInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SubscriberInfo */

// A data object containing authorization information about the subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/carriertoken
func (s_ SubscriberInfo) CarrierToken() foundation.Data {
	rv := objc.Send[foundation.Data](s_.ID, objc.Sel("carrierToken"))
	return rv
}/* debug [instance_properties/getter]: carrierToken */


// A data object containing authorization information about the subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/carriertoken
func (s_ SubscriberInfo) SetCarrierToken(value foundation.Data) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCarrierToken:"), value)
}/* debug [instance_properties/setter]: carrierToken */


// An implementation-defined identifier used to correlate this subscriber with information vended by other APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/identifier
func (s_ SubscriberInfo) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// An implementation-defined identifier used to correlate this subscriber with information vended by other APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/identifier
func (s_ SubscriberInfo) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTSubscriberInfo */


