// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] class.
var (
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass     _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClassOnce sync.Once
)

func getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass() _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass {
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClassOnce.Do(func() {
		MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass = _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass{objc.GetClass("MTRCommissionerControlClusterReverseOpenCommissioningWindowParams")}
	})
	return MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass
}

type _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] class.
type IMTRCommissionerControlClusterReverseOpenCommissioningWindowParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams
type MTRCommissionerControlClusterReverseOpenCommissioningWindowParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsFrom constructs a [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	return MTRCommissionerControlClusterReverseOpenCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass) Alloc() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass) New() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Init() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Autorelease() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParams creates a new MTRCommissionerControlClusterReverseOpenCommissioningWindowParams instance.
func NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParams() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	return getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass().New()
}




// Initialize an MTRCommissionerControlClusterReverseOpenCommissioningWindowParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/init(responseValue:)
func NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	instance := getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass().Alloc()
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/commissioningTimeout
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) CommissioningTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("commissioningTimeout"))
	return rv
}


// SetCommissioningTimeout sets the value of the commissioningTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/commissioningTimeout
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetCommissioningTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/discriminator
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Discriminator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("discriminator"))
	return rv
}


// SetDiscriminator sets the value of the discriminator property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/discriminator
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetDiscriminator(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/iterations
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Iterations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iterations"))
	return rv
}


// SetIterations sets the value of the iterations property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/iterations
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetIterations(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIterations:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/pakePasscodeVerifier
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) PakePasscodeVerifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pakePasscodeVerifier"))
	return rv
}


// SetPakePasscodeVerifier sets the value of the pakePasscodeVerifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/pakePasscodeVerifier
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetPakePasscodeVerifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakePasscodeVerifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/salt
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Salt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("salt"))
	return rv
}


// SetSalt sets the value of the salt property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/salt
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetSalt(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSalt:"), value)
}


