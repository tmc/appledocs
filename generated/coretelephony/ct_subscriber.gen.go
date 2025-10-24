// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTSubscriber */


/* debug [class_header]: Header for CTSubscriber */
// The class instance for the [Subscriber] class.
var (
	SubscriberClass     _SubscriberClass
	SubscriberClassOnce sync.Once
)

func getSubscriberClass() _SubscriberClass {
	SubscriberClassOnce.Do(func() {
		SubscriberClass = _SubscriberClass{objc.GetClass("CTSubscriber")}
	})
	return SubscriberClass
}

type _SubscriberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Subscriber */
// An interface definition for the [Subscriber] class.
type ISubscriber interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Subscriber */
	// properties:
	IsSIMInserted() bool
	SetIsSIMInserted(value bool)
	CTSubscriberTokenRefreshed() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Subscriber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Subscriber */
// Alloc allocates a new instance without initialization.
func (sc _SubscriberClass) Alloc() Subscriber {
	rv := objc.Send[Subscriber](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubscriberClass) New() Subscriber {
	rv := objc.Send[Subscriber](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Subscriber) Init() Subscriber {
	rv := objc.Send[Subscriber](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Subscriber) Autorelease() Subscriber {
	rv := objc.Send[Subscriber](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubscriber creates a new Subscriber instance.
func NewSubscriber() Subscriber {
	return getSubscriberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Subscriber */
// A cellular network subscriber.


// A cellular network subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber
type Subscriber struct {
	objectivec.Object
}

// SubscriberFrom constructs a [Subscriber] from an unsafe.Pointer.
//
// A cellular network subscriber.
func SubscriberFrom(ptr unsafe.Pointer) Subscriber {
	return Subscriber{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Subscriber *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Subscriber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Subscriber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Subscriber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Subscriber */

// A Boolean property that indicates whether a SIM is present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/issiminserted
func (s_ Subscriber) IsSIMInserted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSIMInserted"))
	return rv
}/* debug [instance_properties/getter]: isSIMInserted */


// A Boolean property that indicates whether a SIM is present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/issiminserted
func (s_ Subscriber) SetIsSIMInserted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSIMInserted:"), value)
}/* debug [instance_properties/setter]: isSIMInserted */


// The name of the notification indicating that the carrier token is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscribertokenrefreshed
func (s_ Subscriber) CTSubscriberTokenRefreshed() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("CTSubscriberTokenRefreshed"))
	return rv
}/* debug [instance_properties/getter]: CTSubscriberTokenRefreshed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTSubscriber */


