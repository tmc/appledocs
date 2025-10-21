// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] class.
var (
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass     _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClassOnce sync.Once
)

func getMTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass() _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass {
	MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClassOnce.Do(func() {
		MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass = _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass{objc.GetClass("MTRDiagnosticLogsClusterRetrieveLogsResponseParams")}
	})
	return MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass
}

type _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] class.
type IMTRDiagnosticLogsClusterRetrieveLogsResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDiagnosticLogsClusterRetrieveLogsResponseParams
type MTRDiagnosticLogsClusterRetrieveLogsResponseParams struct {
	objectivec.Object
}

// MTRDiagnosticLogsClusterRetrieveLogsResponseParamsFrom constructs a [MTRDiagnosticLogsClusterRetrieveLogsResponseParams] from an unsafe.Pointer.
func MTRDiagnosticLogsClusterRetrieveLogsResponseParamsFrom(ptr unsafe.Pointer) MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	return MTRDiagnosticLogsClusterRetrieveLogsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass) Alloc() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass) New() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Init() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDiagnosticLogsClusterRetrieveLogsResponseParams) Autorelease() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	rv := objc.Send[MTRDiagnosticLogsClusterRetrieveLogsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDiagnosticLogsClusterRetrieveLogsResponseParams creates a new MTRDiagnosticLogsClusterRetrieveLogsResponseParams instance.
func NewMTRDiagnosticLogsClusterRetrieveLogsResponseParams() MTRDiagnosticLogsClusterRetrieveLogsResponseParams {
	return getMTRDiagnosticLogsClusterRetrieveLogsResponseParamsClass().New()
}




