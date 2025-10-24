// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CBService */


/* debug [class_header]: Header for CBService */
// The class instance for the [CBService] class.
var (
	CBServiceClass     _CBServiceClass
	CBServiceClassOnce sync.Once
)

func getCBServiceClass() _CBServiceClass {
	CBServiceClassOnce.Do(func() {
		CBServiceClass = _CBServiceClass{objc.GetClass("CBService")}
	})
	return CBServiceClass
}

type _CBServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBService */
// An interface definition for the [CBService] class.
type ICBService interface {
	ICBAttribute
	
/* debug [class_interface_properties]: Properties for CBService */
	// properties:
	Characteristics() []CBCharacteristic
	IncludedServices() []CBService
	IsPrimary() bool
	Peripheral() ICBPeripheral
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBService */
// Alloc allocates a new instance without initialization.
func (cc _CBServiceClass) Alloc() CBService {
	rv := objc.Send[CBService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBServiceClass) New() CBService {
	rv := objc.Send[CBService](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBService) Init() CBService {
	rv := objc.Send[CBService](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBService) Autorelease() CBService {
	rv := objc.Send[CBService](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBService creates a new CBService instance.
func NewCBService() CBService {
	return getCBServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBService */
// A collection of data and associated behaviors that accomplish a function or feature of a device.
//
// objects represent services of a remote peripheral. Services are either primary or secondary and may contain multiple characteristics or included services (references to other services).


// A collection of data and associated behaviors that accomplish a function or feature of a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService
type CBService struct {
	CBAttribute
}

// CBServiceFrom constructs a [CBService] from an unsafe.Pointer.
//
// A collection of data and associated behaviors that accomplish a function or feature of a device.
func CBServiceFrom(ptr unsafe.Pointer) CBService {
	return CBService{
		CBAttribute: CBAttributeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBService *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBService */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBService */

// A list of characteristics discovered in this service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/characteristics
func (c_ CBService) Characteristics() []CBCharacteristic {
	rv := objc.Send[[]CBCharacteristic](c_.ID, objc.Sel("characteristics"))
	return rv
}/* debug [instance_properties/getter]: characteristics */


// A list of included services discovered in this service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/includedServices
func (c_ CBService) IncludedServices() []CBService {
	rv := objc.Send[[]CBService](c_.ID, objc.Sel("includedServices"))
	return rv
}/* debug [instance_properties/getter]: includedServices */


// A Boolean value that indicates whether the type of service is primary or secondary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/isPrimary
func (c_ CBService) IsPrimary() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPrimary"))
	return rv
}/* debug [instance_properties/getter]: isPrimary */


// The peripheral to which this service belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/peripheral
func (c_ CBService) Peripheral() ICBPeripheral {
	rv := objc.Send[CBPeripheral](c_.ID, objc.Sel("peripheral"))
	return rv
}/* debug [instance_properties/getter]: peripheral */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBService */



