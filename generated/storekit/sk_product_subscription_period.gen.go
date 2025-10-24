// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProductSubscriptionPeriod */


/* debug [class_header]: Header for SKProductSubscriptionPeriod */
// The class instance for the [ProductSubscriptionPeriod] class.
var (
	ProductSubscriptionPeriodClass     _ProductSubscriptionPeriodClass
	ProductSubscriptionPeriodClassOnce sync.Once
)

func getProductSubscriptionPeriodClass() _ProductSubscriptionPeriodClass {
	ProductSubscriptionPeriodClassOnce.Do(func() {
		ProductSubscriptionPeriodClass = _ProductSubscriptionPeriodClass{objc.GetClass("SKProductSubscriptionPeriod")}
	})
	return ProductSubscriptionPeriodClass
}

type _ProductSubscriptionPeriodClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProductSubscriptionPeriod */
// An interface definition for the [ProductSubscriptionPeriod] class.
type IProductSubscriptionPeriod interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ProductSubscriptionPeriod */
	// properties:
	NumberOfUnits() uint
	Unit() ProductPeriodUnit
	SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetSubscriptionGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	SubscriptionPeriod() ISKProductSubscriptionPeriod
	SetSubscriptionPeriod(value ISKProductSubscriptionPeriod)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProductSubscriptionPeriod */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProductSubscriptionPeriod */
// Alloc allocates a new instance without initialization.
func (pc _ProductSubscriptionPeriodClass) Alloc() ProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ProductSubscriptionPeriodClass) New() ProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProductSubscriptionPeriod) Init() ProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProductSubscriptionPeriod) Autorelease() ProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProductSubscriptionPeriod creates a new ProductSubscriptionPeriod instance.
func NewProductSubscriptionPeriod() ProductSubscriptionPeriod {
	return getProductSubscriptionPeriodClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProductSubscriptionPeriod */
// An object containing the subscription period duration information.
//
// A subscription period is a duration of time defined as some number of units, where a unit can be a , , , or . For example, a subscription period of two weeks has a of a , and a equal to .


// An object containing the subscription period duration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductSubscriptionPeriod
type ProductSubscriptionPeriod struct {
	objectivec.Object
}

// ProductSubscriptionPeriodFrom constructs a [ProductSubscriptionPeriod] from an unsafe.Pointer.
//
// An object containing the subscription period duration information.
func ProductSubscriptionPeriodFrom(ptr unsafe.Pointer) ProductSubscriptionPeriod {
	return ProductSubscriptionPeriod{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProductSubscriptionPeriod *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProductSubscriptionPeriod */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProductSubscriptionPeriod */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProductSubscriptionPeriod */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProductSubscriptionPeriod */

// The number of units per subscription period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductSubscriptionPeriod/numberOfUnits
func (p_ ProductSubscriptionPeriod) NumberOfUnits() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfUnits"))
	return rv
}/* debug [instance_properties/getter]: numberOfUnits */


// The increment of time that a subscription period is specified in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductSubscriptionPeriod/unit
func (p_ ProductSubscriptionPeriod) Unit() ProductPeriodUnit {
	rv := objc.Send[ProductPeriodUnit](p_.ID, objc.Sel("unit"))
	return rv
}/* debug [instance_properties/getter]: unit */


// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptiongroupidentifier
func (p_ ProductSubscriptionPeriod) SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subscriptionGroupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: subscriptionGroupIdentifier */


// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptiongroupidentifier
func (p_ ProductSubscriptionPeriod) SetSubscriptionGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSubscriptionGroupIdentifier:"), value)
}/* debug [instance_properties/setter]: subscriptionGroupIdentifier */


// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ ProductSubscriptionPeriod) SubscriptionPeriod() ISKProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}/* debug [instance_properties/getter]: subscriptionPeriod */


// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ ProductSubscriptionPeriod) SetSubscriptionPeriod(value ISKProductSubscriptionPeriod) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSubscriptionPeriod:"), value)
}/* debug [instance_properties/setter]: subscriptionPeriod */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKProductSubscriptionPeriod */



