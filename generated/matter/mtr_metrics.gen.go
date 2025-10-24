// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMetrics] class.
var (
	MTRMetricsClass     _MTRMetricsClass
	MTRMetricsClassOnce sync.Once
)

func getMTRMetricsClass() _MTRMetricsClass {
	MTRMetricsClassOnce.Do(func() {
		MTRMetricsClass = _MTRMetricsClass{objc.GetClass("MTRMetrics")}
	})
	return MTRMetricsClass
}

type _MTRMetricsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMetrics] class.
type IMTRMetrics interface {
	objectivec.IObject
	// properties:
	AllKeys() objc.IObject /* cross-framework: NSString */
	SetAllKeys(value objc.IObject /* cross-framework: NSString */)
	UniqueIdentifier() objc.IObject /* cross-framework: UUID */
	SetUniqueIdentifier(value objc.IObject /* cross-framework: UUID */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMetrics
type MTRMetrics struct {
	objectivec.Object
}

// MTRMetricsFrom constructs a [MTRMetrics] from an unsafe.Pointer.
func MTRMetricsFrom(ptr unsafe.Pointer) MTRMetrics {
	return MTRMetrics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMetricsClass) Alloc() MTRMetrics {
	rv := objc.Send[MTRMetrics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMetricsClass) New() MTRMetrics {
	rv := objc.Send[MTRMetrics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMetrics) Init() MTRMetrics {
	rv := objc.Send[MTRMetrics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMetrics) Autorelease() MTRMetrics {
	rv := objc.Send[MTRMetrics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMetrics creates a new MTRMetrics instance.
func NewMTRMetrics() MTRMetrics {
	return getMTRMetricsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetrics/allkeys
func (m_ MTRMetrics) AllKeys() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("allKeys"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetrics/allkeys
func (m_ MTRMetrics) SetAllKeys(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllKeys:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetrics/uniqueidentifier
func (m_ MTRMetrics) UniqueIdentifier() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetrics/uniqueidentifier
func (m_ MTRMetrics) SetUniqueIdentifier(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUniqueIdentifier:"), value)
}



