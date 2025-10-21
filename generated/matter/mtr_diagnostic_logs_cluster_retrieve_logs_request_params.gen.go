// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDiagnosticLogsClusterRetrieveLogsRequestParams] class.
var (
	MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass     _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass
	MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClassOnce sync.Once
)

func getMTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass() _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass {
	MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClassOnce.Do(func() {
		MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass = _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass{objc.GetClass("MTRDiagnosticLogsClusterRetrieveLogsRequestParams")}
	})
	return MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass
}

type _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDiagnosticLogsClusterRetrieveLogsRequestParams] class.
type IMTRDiagnosticLogsClusterRetrieveLogsRequestParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDiagnosticLogsClusterRetrieveLogsRequestParams
type MTRDiagnosticLogsClusterRetrieveLogsRequestParams struct {
	objectivec.Object
}

// MTRDiagnosticLogsClusterRetrieveLogsRequestParamsFrom constructs a [MTRDiagnosticLogsClusterRetrieveLogsRequestParams] from an unsafe.Pointer.
func MTRDiagnosticLogsClusterRetrieveLogsRequestParamsFrom(ptr unsafe.Pointer) MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	return MTRDiagnosticLogsClusterRetrieveLogsRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass) Alloc() MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass) New() MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) Init() MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) Autorelease() MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDiagnosticLogsClusterRetrieveLogsRequestParams creates a new MTRDiagnosticLogsClusterRetrieveLogsRequestParams instance.
func NewMTRDiagnosticLogsClusterRetrieveLogsRequestParams() MTRDiagnosticLogsClusterRetrieveLogsRequestParams {
	return getMTRDiagnosticLogsClusterRetrieveLogsRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/transferfiledesignator
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) TransferFileDesignator() string {
	rv := objc.Send[string](m_.ID, objc.Sel("transferFileDesignator"))
	return rv
}


// SetTransferFileDesignator sets the value of the transferFileDesignator property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/transferfiledesignator
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) SetTransferFileDesignator(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransferFileDesignator:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/timedinvoketimeoutms
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/serversideprocessingtimeout
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/serversideprocessingtimeout
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/requestedprotocol
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) RequestedProtocol() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("requestedProtocol"))
	return rv
}


// SetRequestedProtocol sets the value of the requestedProtocol property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/requestedprotocol
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) SetRequestedProtocol(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestedProtocol:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/intent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) Intent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("intent"))
	return rv
}


// SetIntent sets the value of the intent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdiagnosticlogsclusterretrievelogsrequestparams/intent
func (m_ MTRDiagnosticLogsClusterRetrieveLogsRequestParams) SetIntent(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntent:"), value)
}



