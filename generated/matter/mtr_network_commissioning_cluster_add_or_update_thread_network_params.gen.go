// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] class.
var (
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass     _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass() _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass {
	MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClassOnce.Do(func() {
		MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass = _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass{objc.GetClass("MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams")}
	})
	return MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass
}

type _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] class.
type IMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams
type MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsFrom constructs a [MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	return MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass) Alloc() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass) New() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) Init() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) Autorelease() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	rv := objc.Send[MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams creates a new MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams instance.
func NewMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams() MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams {
	return getMTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) Breadcrumb() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// SetBreadcrumb sets the value of the breadcrumb property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/breadcrumb
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) SetBreadcrumb(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/operationaldataset
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) OperationalDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalDataset"))
	return rv
}


// SetOperationalDataset sets the value of the operationalDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/operationaldataset
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) SetOperationalDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalDataset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/serversideprocessingtimeout
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrnetworkcommissioningclusteraddorupdatethreadnetworkparams/timedinvoketimeoutms
func (m_ MTRNetworkCommissioningClusterAddOrUpdateThreadNetworkParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



