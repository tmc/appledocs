// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	DictionaryRepresentation() unsafe.Pointer
	JSONRepresentation() unsafe.Pointer
}

// An object that encapsulates a diagnostic report.
//
// The system delivers a diagnostic report as soon as it’s available.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/dictionaryRepresentation()
func (m_ MXDiagnosticPayload) DictionaryRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}

// Returns the contents of the payload in JSON format.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/jsonRepresentation()
func (m_ MXDiagnosticPayload) JSONRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}

// The diagnostic reports for the app launch time.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/appLaunchDiagnostics
func (m_ MXDiagnosticPayload) AppLaunchDiagnostics() []MXAppLaunchDiagnostic {
	rv := objc.Send[[]MXAppLaunchDiagnostic](m_.ID, objc.Sel("appLaunchDiagnostics"))
	return rv
}

// The diagnostic reports for fatal and nonfatal CPU exceptions for the app during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/cpuExceptionDiagnostics
func (m_ MXDiagnosticPayload) CpuExceptionDiagnostics() []MXCPUExceptionDiagnostic {
	rv := objc.Send[[]MXCPUExceptionDiagnostic](m_.ID, objc.Sel("cpuExceptionDiagnostics"))
	return rv
}

// The diagnostic reports for app crashes during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/crashDiagnostics
func (m_ MXDiagnosticPayload) CrashDiagnostics() []MXCrashDiagnostic {
	rv := objc.Send[[]MXCrashDiagnostic](m_.ID, objc.Sel("crashDiagnostics"))
	return rv
}

// The diagnostic reports for disk write exceptions for the app during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/diskWriteExceptionDiagnostics
func (m_ MXDiagnosticPayload) DiskWriteExceptionDiagnostics() []MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[[]MXDiskWriteExceptionDiagnostic](m_.ID, objc.Sel("diskWriteExceptionDiagnostics"))
	return rv
}

// The diagnostic reports for times when the app was too busy to handle input responsively during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/hangDiagnostics
func (m_ MXDiagnosticPayload) HangDiagnostics() []MXHangDiagnostic {
	rv := objc.Send[[]MXHangDiagnostic](m_.ID, objc.Sel("hangDiagnostics"))
	return rv
}

// The starting time of the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampBegin
func (m_ MXDiagnosticPayload) TimeStampBegin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeStampBegin"))
	return rv
}

// The ending time of the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampEnd
func (m_ MXDiagnosticPayload) TimeStampEnd() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeStampEnd"))
	return rv
}



