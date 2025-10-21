// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] class.
var (
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass     _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClassOnce sync.Once
)

func getMTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass() _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass {
	MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClassOnce.Do(func() {
		MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass = _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass{objc.GetClass("MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct")}
	})
	return MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass
}

type _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] class.
type IMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct
type MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct struct {
	objectivec.Object
}

// MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructFrom constructs a [MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct] from an unsafe.Pointer.
func MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructFrom(ptr unsafe.Pointer) MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	return MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass) Alloc() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass) New() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) Init() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct) Autorelease() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	rv := objc.Send[MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct creates a new MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct instance.
func NewMTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct() MTRActivatedCarbonFilterMonitoringClusterReplacementProductStruct {
	return getMTRActivatedCarbonFilterMonitoringClusterReplacementProductStructClass().New()
}




