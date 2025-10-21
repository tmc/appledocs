// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CBService] class.
type ICBService interface {
	ICBAttribute
}

// A collection of data and associated behaviors that accomplish a function or feature of a device.
//
// objects represent services of a remote peripheral. Services are either primary or secondary and may contain multiple characteristics or included services (references to other services).
//
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

// Alloc allocates a new instance without initialization.
func (cc _CBServiceClass) Alloc() CBService {
	rv := objc.Send[CBService](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A list of characteristics discovered in this service.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/characteristics
func (c_ CBService) Characteristics() []CBCharacteristic {
	rv := objc.Send[[]CBCharacteristic](c_.ID, objc.Sel("characteristics"))
	return rv
}

// A list of included services discovered in this service.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/includedServices
func (c_ CBService) IncludedServices() []CBService {
	rv := objc.Send[[]CBService](c_.ID, objc.Sel("includedServices"))
	return rv
}

// A Boolean value that indicates whether the type of service is primary or secondary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/isPrimary
func (c_ CBService) IsPrimary() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPrimary"))
	return rv
}

// The peripheral to which this service belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBService/peripheral
func (c_ CBService) Peripheral() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("peripheral"))
	return rv
}



