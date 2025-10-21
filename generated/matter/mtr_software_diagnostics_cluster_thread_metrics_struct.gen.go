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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Id() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("id"))
	return rv
}


// SetId sets the value of the id property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/id
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/name
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeCurrent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stackFreeCurrent"))
	return rv
}


// SetStackFreeCurrent sets the value of the stackFreeCurrent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreecurrent
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeCurrent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackFreeMinimum() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stackFreeMinimum"))
	return rv
}


// SetStackFreeMinimum sets the value of the stackFreeMinimum property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stackfreeminimum
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackFreeMinimum(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackFreeMinimum:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) StackSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("stackSize"))
	return rv
}


// SetStackSize sets the value of the stackSize property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclusterthreadmetricsstruct/stacksize
func (m_ MTRSoftwareDiagnosticsClusterThreadMetricsStruct) SetStackSize(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStackSize:"), value)
}



