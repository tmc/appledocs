// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXDiagnosticPayload] class.
var (
	MXDiagnosticPayloadClass     _MXDiagnosticPayloadClass
	MXDiagnosticPayloadClassOnce sync.Once
)

func getMXDiagnosticPayloadClass() _MXDiagnosticPayloadClass {
	MXDiagnosticPayloadClassOnce.Do(func() {
		MXDiagnosticPayloadClass = _MXDiagnosticPayloadClass{objc.GetClass("MXDiagnosticPayload")}
	})
	return MXDiagnosticPayloadClass
}

type _MXDiagnosticPayloadClass struct {
	class objc.Class
}

// An interface definition for the [MXDiagnosticPayload] class.
type IMXDiagnosticPayload interface {
	objectivec.IObject
	// properties:
	CpuExceptionDiagnostics() []IMXCPUExceptionDiagnostic
	CrashDiagnostics() []IMXCrashDiagnostic
	DiskWriteExceptionDiagnostics() []IMXDiskWriteExceptionDiagnostic
	HangDiagnostics() []IMXHangDiagnostic
	TimeStampBegin() objc.IObject /* cross-framework: NSDate */
	TimeStampEnd() objc.IObject /* cross-framework: NSDate */
	// methods:
	DictionaryRepresentation() objc.IObject /* cross-framework: Dictionary */
	JSONRepresentation() objc.IObject /* cross-framework: Data */
}

// An object that encapsulates a diagnostic report.
//
// The system delivers a diagnostic report as soon as it’s available.


// An object that encapsulates a diagnostic report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload
type MXDiagnosticPayload struct {
	objectivec.Object
}

// MXDiagnosticPayloadFrom constructs a [MXDiagnosticPayload] from an unsafe.Pointer.
//
// An object that encapsulates a diagnostic report.
func MXDiagnosticPayloadFrom(ptr unsafe.Pointer) MXDiagnosticPayload {
	return MXDiagnosticPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXDiagnosticPayloadClass) Alloc() MXDiagnosticPayload {
	rv := objc.Send[MXDiagnosticPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXDiagnosticPayloadClass) New() MXDiagnosticPayload {
	rv := objc.Send[MXDiagnosticPayload](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiagnosticPayload) Init() MXDiagnosticPayload {
	rv := objc.Send[MXDiagnosticPayload](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiagnosticPayload) Autorelease() MXDiagnosticPayload {
	rv := objc.Send[MXDiagnosticPayload](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiagnosticPayload creates a new MXDiagnosticPayload instance.
func NewMXDiagnosticPayload() MXDiagnosticPayload {
	return getMXDiagnosticPayloadClass().New()
}



// Returns the results of the payload as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/dictionaryRepresentation()
func (m_ MXDiagnosticPayload) DictionaryRepresentation() objc.IObject /* cross-framework: Dictionary */ {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}


// Returns the contents of the payload in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/jsonRepresentation()
func (m_ MXDiagnosticPayload) JSONRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}


// The diagnostic reports for fatal and nonfatal CPU exceptions for the app during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/cpuExceptionDiagnostics
func (m_ MXDiagnosticPayload) CpuExceptionDiagnostics() []IMXCPUExceptionDiagnostic {
	rv := objc.Send[[]MXCPUExceptionDiagnostic](m_.ID, objc.Sel("cpuExceptionDiagnostics"))
	return rv
}


// The diagnostic reports for app crashes during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/crashDiagnostics
func (m_ MXDiagnosticPayload) CrashDiagnostics() []IMXCrashDiagnostic {
	rv := objc.Send[[]MXCrashDiagnostic](m_.ID, objc.Sel("crashDiagnostics"))
	return rv
}


// The diagnostic reports for disk write exceptions for the app during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/diskWriteExceptionDiagnostics
func (m_ MXDiagnosticPayload) DiskWriteExceptionDiagnostics() []IMXDiskWriteExceptionDiagnostic {
	rv := objc.Send[[]MXDiskWriteExceptionDiagnostic](m_.ID, objc.Sel("diskWriteExceptionDiagnostics"))
	return rv
}


// The diagnostic reports for times when the app was too busy to handle input responsively during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/hangDiagnostics
func (m_ MXDiagnosticPayload) HangDiagnostics() []IMXHangDiagnostic {
	rv := objc.Send[[]MXHangDiagnostic](m_.ID, objc.Sel("hangDiagnostics"))
	return rv
}


// The starting time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampBegin
func (m_ MXDiagnosticPayload) TimeStampBegin() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampBegin"))
	return rv
}


// The ending time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampEnd
func (m_ MXDiagnosticPayload) TimeStampEnd() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampEnd"))
	return rv
}


