// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsClass     _MTRSoftwareDiagnosticsClusterThreadMetricsClass
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsClass() _MTRSoftwareDiagnosticsClusterThreadMetricsClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsClass = _MTRSoftwareDiagnosticsClusterThreadMetricsClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetrics")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetrics] class.
type IMTRSoftwareDiagnosticsClusterThreadMetrics interface {
	IMTRSoftwareDiagnosticsClusterThreadMetricsStruct
	// properties:
	Id() objc.IObject /* cross-framework: NSNumber */
	SetId(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */
	SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */)
	StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */
	SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */)
	StackSize() objc.IObject /* cross-framework: NSNumber */
	SetStackSize(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetrics
type MTRSoftwareDiagnosticsClusterThreadMetrics struct {
	MTRSoftwareDiagnosticsClusterThreadMetricsStruct
}

// MTRSoftwareDiagnosticsClusterThreadMetricsFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetrics] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetrics {
	return MTRSoftwareDiagnosticsClusterThreadMetrics{
		MTRSoftwareDiagnosticsClusterThreadMetricsStruct: MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsClass) New() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Init() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetrics {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetrics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetrics creates a new MTRSoftwareDiagnosticsClusterThreadMetrics instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetrics() MTRSoftwareDiagnosticsClusterThreadMetrics {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeCurrent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeCurrent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeMinimum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeMinimum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) StackSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetrics/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetrics) SetStackSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackSize:"), value)
}



