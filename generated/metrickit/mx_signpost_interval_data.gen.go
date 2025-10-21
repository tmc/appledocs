// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXSignpostIntervalData] class.
var (
	MXSignpostIntervalDataClass     _MXSignpostIntervalDataClass
	MXSignpostIntervalDataClassOnce sync.Once
)

func getMXSignpostIntervalDataClass() _MXSignpostIntervalDataClass {
	MXSignpostIntervalDataClassOnce.Do(func() {
		MXSignpostIntervalDataClass = _MXSignpostIntervalDataClass{objc.GetClass("MXSignpostIntervalData")}
	})
	return MXSignpostIntervalDataClass
}

type _MXSignpostIntervalDataClass struct {
	class objc.Class
}

// An interface definition for the [MXSignpostIntervalData] class.
type IMXSignpostIntervalData interface {
	objectivec.IObject
}

// A data object representing the captured data for a custom metric.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData
type MXSignpostIntervalData struct {
	objectivec.Object
}

// MXSignpostIntervalDataFrom constructs a [MXSignpostIntervalData] from an unsafe.Pointer.
//
// A data object representing the captured data for a custom metric.
func MXSignpostIntervalDataFrom(ptr unsafe.Pointer) MXSignpostIntervalData {
	return MXSignpostIntervalData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXSignpostIntervalDataClass) Alloc() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXSignpostIntervalDataClass) New() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXSignpostIntervalData) Init() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXSignpostIntervalData) Autorelease() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXSignpostIntervalData creates a new MXSignpostIntervalData instance.
func NewMXSignpostIntervalData() MXSignpostIntervalData {
	return getMXSignpostIntervalDataClass().New()
}




