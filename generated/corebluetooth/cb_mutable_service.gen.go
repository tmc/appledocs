// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBMutableService */


/* debug [class_header]: Header for CBMutableService */
// The class instance for the [CBMutableService] class.
var (
	CBMutableServiceClass     _CBMutableServiceClass
	CBMutableServiceClassOnce sync.Once
)

func getCBMutableServiceClass() _CBMutableServiceClass {
	CBMutableServiceClassOnce.Do(func() {
		CBMutableServiceClass = _CBMutableServiceClass{objc.GetClass("CBMutableService")}
	})
	return CBMutableServiceClass
}

type _CBMutableServiceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBMutableService */
// An interface definition for the [CBMutableService] class.
type ICBMutableService interface {
	ICBService
	
/* debug [class_interface_properties]: Properties for CBMutableService */
	// properties:
	Characteristics() []CBCharacteristic
	SetCharacteristics(value []CBCharacteristic)
	IncludedServices() []CBService
	SetIncludedServices(value []CBService)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBMutableService */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBMutableService */
// Alloc allocates a new instance without initialization.
func (cc _CBMutableServiceClass) Alloc() CBMutableService {
	rv := objc.Send[CBMutableService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBMutableServiceClass) New() CBMutableService {
	rv := objc.Send[CBMutableService](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBMutableService) Init() CBMutableService {
	rv := objc.Send[CBMutableService](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBMutableService) Autorelease() CBMutableService {
	rv := objc.Send[CBMutableService](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableService creates a new CBMutableService instance.
func NewCBMutableService() CBMutableService {
	return getCBMutableServiceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBMutableService */
// A service with writeable property values.
//
// The class adds write access to all of the properties in the class it inherits from. You use this class to create a service or an included service on a local peripheral device (represented by a object). After creating a service, you can add it to the peripheral’s local database using the method of the class. After you add a service to the peripheral’s local database, Core Bluetooth caches the service and you can no longer make changes to it.


// A service with writeable property values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService
type CBMutableService struct {
	CBService
}

// CBMutableServiceFrom constructs a [CBMutableService] from an unsafe.Pointer.
//
// A service with writeable property values.
func CBMutableServiceFrom(ptr unsafe.Pointer) CBMutableService {
	return CBMutableService{
		CBService: CBServiceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBMutableService */

// Creates a newly initialized mutable service specified by UUID and service type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/init(type:primary:)
func NewCBMutableServiceWithTypePrimary(UUID ICBUUID, isPrimary bool) CBMutableService {
	instance := getCBMutableServiceClass().Alloc()
	rv := objc.Send[CBMutableService](instance.ID, objc.Sel("initWithType:primary:"), UUID, isPrimary)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCBMutableServiceWithTypePrimary */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBMutableService */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBMutableService */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBMutableService */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBMutableService */

// A list of characteristics of a service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/characteristics
func (c_ CBMutableService) Characteristics() []CBCharacteristic {
	rv := objc.Send[[]CBCharacteristic](c_.ID, objc.Sel("characteristics"))
	return rv
}/* debug [instance_properties/getter]: characteristics */


// A list of characteristics of a service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/characteristics
func (c_ CBMutableService) SetCharacteristics(value []CBCharacteristic) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setCharacteristics:"), nsArray)
}/* debug [instance_properties/setter]: characteristics */


// A list of included services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/includedServices
func (c_ CBMutableService) IncludedServices() []CBService {
	rv := objc.Send[[]CBService](c_.ID, objc.Sel("includedServices"))
	return rv
}/* debug [instance_properties/getter]: includedServices */


// A list of included services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableService/includedServices
func (c_ CBMutableService) SetIncludedServices(value []CBService) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setIncludedServices:"), nsArray)
}/* debug [instance_properties/setter]: includedServices */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBMutableService */


