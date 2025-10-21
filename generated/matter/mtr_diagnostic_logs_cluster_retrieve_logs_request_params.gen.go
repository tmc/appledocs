// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




