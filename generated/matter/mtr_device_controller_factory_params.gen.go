// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceControllerFactoryParams] class.
var (
	MTRDeviceControllerFactoryParamsClass     _MTRDeviceControllerFactoryParamsClass
	MTRDeviceControllerFactoryParamsClassOnce sync.Once
)

func getMTRDeviceControllerFactoryParamsClass() _MTRDeviceControllerFactoryParamsClass {
	MTRDeviceControllerFactoryParamsClassOnce.Do(func() {
		MTRDeviceControllerFactoryParamsClass = _MTRDeviceControllerFactoryParamsClass{objc.GetClass("MTRDeviceControllerFactoryParams")}
	})
	return MTRDeviceControllerFactoryParamsClass
}

type _MTRDeviceControllerFactoryParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerFactoryParams] class.
type IMTRDeviceControllerFactoryParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams
type MTRDeviceControllerFactoryParams struct {
	objectivec.Object
}

// MTRDeviceControllerFactoryParamsFrom constructs a [MTRDeviceControllerFactoryParams] from an unsafe.Pointer.
func MTRDeviceControllerFactoryParamsFrom(ptr unsafe.Pointer) MTRDeviceControllerFactoryParams {
	return MTRDeviceControllerFactoryParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerFactoryParamsClass) Alloc() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerFactoryParamsClass) New() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerFactoryParams) Init() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerFactoryParams) Autorelease() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerFactoryParams creates a new MTRDeviceControllerFactoryParams instance.
func NewMTRDeviceControllerFactoryParams() MTRDeviceControllerFactoryParams {
	return getMTRDeviceControllerFactoryParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/productAttestationAuthorityCertificates
func (m_ MTRDeviceControllerFactoryParams) ProductAttestationAuthorityCertificates() []addressbook.NSData {
	rv := objc.Send[[]addressbook.NSData](m_.ID, objc.Sel("productAttestationAuthorityCertificates"))
	return rv
}


// SetProductAttestationAuthorityCertificates sets the value of the productAttestationAuthorityCertificates property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/productAttestationAuthorityCertificates
func (m_ MTRDeviceControllerFactoryParams) SetProductAttestationAuthorityCertificates(value []addressbook.NSData) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductAttestationAuthorityCertificates:"), nsArray)
}


