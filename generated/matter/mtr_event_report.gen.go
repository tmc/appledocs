// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	SetError(value objc.IObject /* cross-framework: Error */)
	EventNumber() objc.IObject /* cross-framework: NSNumber */
	SetEventNumber(value objc.IObject /* cross-framework: NSNumber */)
	EventTimeType() MTREventTimeType
	SetEventTimeType(value MTREventTimeType)
	Path() IMTREventPath
	SetPath(value IMTREventPath)
	Priority() objc.IObject /* cross-framework: NSNumber */
	SetPriority(value objc.IObject /* cross-framework: NSNumber */)
	SystemUpTime() float64
	SetSystemUpTime(value float64)
	Timestamp() objc.IObject /* cross-framework: NSNumber */
	SetTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	TimestampDate() objc.IObject /* cross-framework: Date */
	SetTimestampDate(value objc.IObject /* cross-framework: Date */)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/error
func (m_ MTREventReport) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](m_.ID, objc.Sel("error"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/error
func (m_ MTREventReport) SetError(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventnumber
func (m_ MTREventReport) EventNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("eventNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventnumber
func (m_ MTREventReport) SetEventNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEventNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventtimetype
func (m_ MTREventReport) EventTimeType() MTREventTimeType {
	rv := objc.Send[MTREventTimeType](m_.ID, objc.Sel("eventTimeType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventtimetype
func (m_ MTREventReport) SetEventTimeType(value MTREventTimeType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEventTimeType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/path
func (m_ MTREventReport) Path() IMTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("path"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/path
func (m_ MTREventReport) SetPath(value IMTREventPath) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/priority
func (m_ MTREventReport) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/priority
func (m_ MTREventReport) SetPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/systemuptime
func (m_ MTREventReport) SystemUpTime() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("systemUpTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/systemuptime
func (m_ MTREventReport) SetSystemUpTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemUpTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestamp
func (m_ MTREventReport) Timestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestamp
func (m_ MTREventReport) SetTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimestamp:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestampdate
func (m_ MTREventReport) TimestampDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("timestampDate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestampdate
func (m_ MTREventReport) SetTimestampDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimestampDate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/value
func (m_ MTREventReport) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/value
func (m_ MTREventReport) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



