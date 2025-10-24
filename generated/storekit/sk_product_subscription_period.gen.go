// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ProductSubscriptionPeriod] class.
type IProductSubscriptionPeriod interface {
	objectivec.IObject
	// properties:
	SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetSubscriptionGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	SubscriptionPeriod() ISKProductSubscriptionPeriod
	SetSubscriptionPeriod(value ISKProductSubscriptionPeriod)
	NumberOfUnits() int /* primitive/slice/pointer. */
	SetNumberOfUnits(value int /* primitive/slice/pointer. */)
	Unit() unsafe.Pointer
	SetUnit(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _ProductSubscriptionPeriodClass) Alloc() ProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptiongroupidentifier
func (p_ ProductSubscriptionPeriod) SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subscriptionGroupIdentifier"))
	return rv
}


// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptiongroupidentifier
func (p_ ProductSubscriptionPeriod) SetSubscriptionGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSubscriptionGroupIdentifier:"), value)
}


// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ ProductSubscriptionPeriod) SubscriptionPeriod() ISKProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}


// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ ProductSubscriptionPeriod) SetSubscriptionPeriod(value ISKProductSubscriptionPeriod) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSubscriptionPeriod:"), value)
}


// The number of units per subscription period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproductsubscriptionperiod/numberofunits
func (p_ ProductSubscriptionPeriod) NumberOfUnits() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfUnits"))
	return rv
}


// The number of units per subscription period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproductsubscriptionperiod/numberofunits
func (p_ ProductSubscriptionPeriod) SetNumberOfUnits(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfUnits:"), value)
}


// The increment of time that a subscription period is specified in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproductsubscriptionperiod/unit
func (p_ ProductSubscriptionPeriod) Unit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("unit"))
	return rv
}


// The increment of time that a subscription period is specified in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproductsubscriptionperiod/unit
func (p_ ProductSubscriptionPeriod) SetUnit(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnit:"), value)
}



