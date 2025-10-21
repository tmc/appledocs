// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEFilterReport] class.
var (
	NEFilterReportClass     _NEFilterReportClass
	NEFilterReportClassOnce sync.Once
)

func getNEFilterReportClass() _NEFilterReportClass {
	NEFilterReportClassOnce.Do(func() {
		NEFilterReportClass = _NEFilterReportClass{objc.GetClass("NEFilterReport")}
	})
	return NEFilterReportClass
}

type _NEFilterReportClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterReport] class.
type INEFilterReport interface {
	objectivec.IObject
}

// The report of the data provider’s action on a flow.
//
// The system issues a report by calling your control provider’s method with a report instance when the data provider issues a verdict whose property is set to .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport
type NEFilterReport struct {
	objectivec.Object
}

// NEFilterReportFrom constructs a [NEFilterReport] from an unsafe.Pointer.
//
// The report of the data provider’s action on a flow.
func NEFilterReportFrom(ptr unsafe.Pointer) NEFilterReport {
	return NEFilterReport{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterReportClass) Alloc() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterReportClass) New() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterReport) Init() NEFilterReport {
	rv := objc.Send[NEFilterReport](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterReport) Autorelease() NEFilterReport {
	rv := objc.Send[NEFilterReport](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterReport creates a new NEFilterReport instance.
func NewNEFilterReport() NEFilterReport {
	return getNEFilterReportClass().New()
}




