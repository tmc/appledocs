// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Error() foundation.Error
	SetError(value foundation.IError)
	EventNumber() foundation.Number
	SetEventNumber(value foundation.INumber)
	EventTimeType() MTREventTimeType
	SetEventTimeType(value MTREventTimeType)
	Path() MTREventPath
	SetPath(value IMTREventPath)
	Priority() foundation.Number
	SetPriority(value foundation.INumber)
	SystemUpTime() unsafe.Pointer
	SetSystemUpTime(value unsafe.Pointer)
	Timestamp() foundation.Number
	SetTimestamp(value foundation.INumber)
	TimestampDate() foundation.Date
	SetTimestampDate(value foundation.IDate)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/error
func (m_ MTREventReport) Error() foundation.Error {
	rv := objc.Send[foundation.Error](m_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/error
func (m_ MTREventReport) SetError(value foundation.IError) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventnumber
func (m_ MTREventReport) EventNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("eventNumber"))
	return rv
}


// SetEventNumber sets the value of the eventNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventnumber
func (m_ MTREventReport) SetEventNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEventNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventtimetype
func (m_ MTREventReport) EventTimeType() MTREventTimeType {
	rv := objc.Send[MTREventTimeType](m_.ID, objc.Sel("eventTimeType"))
	return rv
}


// SetEventTimeType sets the value of the eventTimeType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/eventtimetype
func (m_ MTREventReport) SetEventTimeType(value MTREventTimeType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEventTimeType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/path
func (m_ MTREventReport) Path() MTREventPath {
	rv := objc.Send[MTREventPath](m_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/path
func (m_ MTREventReport) SetPath(value IMTREventPath) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/priority
func (m_ MTREventReport) Priority() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("priority"))
	return rv
}


// SetPriority sets the value of the priority property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/priority
func (m_ MTREventReport) SetPriority(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/systemuptime
func (m_ MTREventReport) SystemUpTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("systemUpTime"))
	return rv
}


// SetSystemUpTime sets the value of the systemUpTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/systemuptime
func (m_ MTREventReport) SetSystemUpTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemUpTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestamp
func (m_ MTREventReport) Timestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestamp
func (m_ MTREventReport) SetTimestamp(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestampdate
func (m_ MTREventReport) TimestampDate() foundation.Date {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("timestampDate"))
	return rv
}


// SetTimestampDate sets the value of the timestampDate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/timestampdate
func (m_ MTREventReport) SetTimestampDate(value foundation.IDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimestampDate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/value
func (m_ MTREventReport) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventreport/value
func (m_ MTREventReport) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



