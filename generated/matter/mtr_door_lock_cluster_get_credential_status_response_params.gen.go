// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
var (
	MTRDoorLockClusterGetCredentialStatusResponseParamsClass     _MTRDoorLockClusterGetCredentialStatusResponseParamsClass
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetCredentialStatusResponseParamsClass() _MTRDoorLockClusterGetCredentialStatusResponseParamsClass {
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetCredentialStatusResponseParamsClass = _MTRDoorLockClusterGetCredentialStatusResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetCredentialStatusResponseParams")}
	})
	return MTRDoorLockClusterGetCredentialStatusResponseParamsClass
}

type _MTRDoorLockClusterGetCredentialStatusResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
type IMTRDoorLockClusterGetCredentialStatusResponseParams interface {
	objectivec.IObject
	CreatorFabricIndex() foundation.Number
	SetCreatorFabricIndex(value foundation.INumber)
	CredentialData() foundation.Data
	SetCredentialData(value foundation.IData)
	CredentialExists() foundation.Number
	SetCredentialExists(value foundation.INumber)
	LastModifiedFabricIndex() foundation.Number
	SetLastModifiedFabricIndex(value foundation.INumber)
	NextCredentialIndex() foundation.Number
	SetNextCredentialIndex(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UserIndex() foundation.Number
	SetUserIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams
type MTRDoorLockClusterGetCredentialStatusResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetCredentialStatusResponseParamsFrom constructs a [MTRDoorLockClusterGetCredentialStatusResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetCredentialStatusResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusResponseParams {
	return MTRDoorLockClusterGetCredentialStatusResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) Alloc() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) New() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Init() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Autorelease() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetCredentialStatusResponseParams creates a new MTRDoorLockClusterGetCredentialStatusResponseParams instance.
func NewMTRDoorLockClusterGetCredentialStatusResponseParams() MTRDoorLockClusterGetCredentialStatusResponseParams {
	return getMTRDoorLockClusterGetCredentialStatusResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CreatorFabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}


// SetCreatorFabricIndex sets the value of the creatorFabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCreatorFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialdata
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialData() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("credentialData"))
	return rv
}


// SetCredentialData sets the value of the credentialData property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialdata
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialexists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialExists() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("credentialExists"))
	return rv
}


// SetCredentialExists sets the value of the credentialExists property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialexists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialExists(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialExists:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) LastModifiedFabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}


// SetLastModifiedFabricIndex sets the value of the lastModifiedFabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetLastModifiedFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) NextCredentialIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nextCredentialIndex"))
	return rv
}


// SetNextCredentialIndex sets the value of the nextCredentialIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetNextCredentialIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextCredentialIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/userindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) UserIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userIndex"))
	return rv
}


// SetUserIndex sets the value of the userIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/userindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetUserIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



