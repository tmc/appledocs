// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREventReport] class.
var (
	MTREventReportClass     _MTREventReportClass
	MTREventReportClassOnce sync.Once
)

func getMTREventReportClass() _MTREventReportClass {
	MTREventReportClassOnce.Do(func() {
		MTREventReportClass = _MTREventReportClass{objc.GetClass("MTREventReport")}
	})
	return MTREventReportClass
}

type _MTREventReportClass struct {
	class objc.Class
}

// An interface definition for the [MTREventReport] class.
type IMTREventReport interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventReport
type MTREventReport struct {
	objectivec.Object
}

// MTREventReportFrom constructs a [MTREventReport] from an unsafe.Pointer.
func MTREventReportFrom(ptr unsafe.Pointer) MTREventReport {
	return MTREventReport{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREventReportClass) Alloc() MTREventReport {
	rv := objc.Send[MTREventReport](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREventReportClass) New() MTREventReport {
	rv := objc.Send[MTREventReport](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventReport) Init() MTREventReport {
	rv := objc.Send[MTREventReport](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventReport) Autorelease() MTREventReport {
	rv := objc.Send[MTREventReport](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventReport creates a new MTREventReport instance.
func NewMTREventReport() MTREventReport {
	return getMTREventReportClass().New()
}




