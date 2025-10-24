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
	// properties:
	CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	CredentialData() objc.IObject /* cross-framework: Data */
	SetCredentialData(value objc.IObject /* cross-framework: Data */)
	CredentialExists() objc.IObject /* cross-framework: NSNumber */
	SetCredentialExists(value objc.IObject /* cross-framework: NSNumber */)
	LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */
	SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UserIndex() objc.IObject /* cross-framework: NSNumber */
	SetUserIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CreatorFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("creatorFabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/creatorfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCreatorFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCreatorFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialdata
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("credentialData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialdata
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialexists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) CredentialExists() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("credentialExists"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/credentialexists
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetCredentialExists(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCredentialExists:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) LastModifiedFabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastModifiedFabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/lastmodifiedfabricindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetLastModifiedFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastModifiedFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) NextCredentialIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextCredentialIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/nextcredentialindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetNextCredentialIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextCredentialIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/timedinvoketimeoutms
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/userindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) UserIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclustergetcredentialstatusresponseparams/userindex
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) SetUserIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserIndex:"), value)
}



