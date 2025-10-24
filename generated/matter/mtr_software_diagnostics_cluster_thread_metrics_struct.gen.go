// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
var (
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClass     _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass() _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass {
	MTRSoftwareDiagnosticsClusterThreadMetricsStructClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterThreadMetricsStructClass = _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass{objc.GetClass("MTRSoftwareDiagnosticsClusterThreadMetricsStruct")}
	})
	return MTRSoftwareDiagnosticsClusterThreadMetricsStructClass
}

type _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] class.
type IMTRSoftwareDiagnosticsClusterThreadMetricsStruct interface {
	objectivec.IObject
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterThreadMetricsStruct
type MTRSoftwareDiagnosticsClusterThreadMetricsStruct struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom constructs a [MTRSoftwareDiagnosticsClusterThreadMetricsStruct] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterThreadMetricsStructFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return MTRSoftwareDiagnosticsClusterThreadMetricsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) Alloc() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterThreadMetricsStructClass) New() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Init() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Autorelease() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterThreadMetricsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct creates a new MTRSoftwareDiagnosticsClusterThreadMetricsStruct instance.
func NewMTRSoftwareDiagnosticsClusterThreadMetricsStruct() MTRSoftwareDiagnosticsClusterThreadMetricsStruct {
	return getMTRSoftwareDiagnosticsClusterThreadMetricsStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeCurrent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeCurrent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeMinimum() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackFreeMinimum"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeMinimum(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeMinimum:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stackSize"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackSize:"), value)
}



