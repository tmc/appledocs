// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCommissionableBrowserResult] class.
var (
	MTRCommissionableBrowserResultClass     _MTRCommissionableBrowserResultClass
	MTRCommissionableBrowserResultClassOnce sync.Once
)

func getMTRCommissionableBrowserResultClass() _MTRCommissionableBrowserResultClass {
	MTRCommissionableBrowserResultClassOnce.Do(func() {
		MTRCommissionableBrowserResultClass = _MTRCommissionableBrowserResultClass{objc.GetClass("MTRCommissionableBrowserResult")}
	})
	return MTRCommissionableBrowserResultClass
}

type _MTRCommissionableBrowserResultClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissionableBrowserResult] class.
type IMTRCommissionableBrowserResult interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionableBrowserResult
type MTRCommissionableBrowserResult struct {
	objectivec.Object
}

// MTRCommissionableBrowserResultFrom constructs a [MTRCommissionableBrowserResult] from an unsafe.Pointer.
func MTRCommissionableBrowserResultFrom(ptr unsafe.Pointer) MTRCommissionableBrowserResult {
	return MTRCommissionableBrowserResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionableBrowserResultClass) Alloc() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissionableBrowserResultClass) New() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionableBrowserResult) Init() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionableBrowserResult) Autorelease() MTRCommissionableBrowserResult {
	rv := objc.Send[MTRCommissionableBrowserResult](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionableBrowserResult creates a new MTRCommissionableBrowserResult instance.
func NewMTRCommissionableBrowserResult() MTRCommissionableBrowserResult {
	return getMTRCommissionableBrowserResultClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/commissioningmode
func (m_ MTRCommissionableBrowserResult) CommissioningMode() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("commissioningMode"))
	return rv
}


// SetCommissioningMode sets the value of the commissioningMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/commissioningmode
func (m_ MTRCommissionableBrowserResult) SetCommissioningMode(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/discriminator
func (m_ MTRCommissionableBrowserResult) Discriminator() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("discriminator"))
	return rv
}


// SetDiscriminator sets the value of the discriminator property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/discriminator
func (m_ MTRCommissionableBrowserResult) SetDiscriminator(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/instancename
func (m_ MTRCommissionableBrowserResult) InstanceName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("instanceName"))
	return rv
}


// SetInstanceName sets the value of the instanceName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/instancename
func (m_ MTRCommissionableBrowserResult) SetInstanceName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstanceName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/productid
func (m_ MTRCommissionableBrowserResult) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/productid
func (m_ MTRCommissionableBrowserResult) SetProductID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/vendorid
func (m_ MTRCommissionableBrowserResult) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionablebrowserresult/vendorid
func (m_ MTRCommissionableBrowserResult) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



