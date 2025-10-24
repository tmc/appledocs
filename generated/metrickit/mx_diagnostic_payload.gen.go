// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXDiagnosticPayload */


/* debug [class_header]: Header for MXDiagnosticPayload */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDiagnosticPayload */
// An interface definition for the [MXDiagnosticPayload] class.
type IMXDiagnosticPayload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXDiagnosticPayload */
	// properties:
	CpuExceptionDiagnostics() []MXCPUExceptionDiagnostic
	CrashDiagnostics() []MXCrashDiagnostic
	DiskWriteExceptionDiagnostics() []MXDiskWriteExceptionDiagnostic
	HangDiagnostics() []MXHangDiagnostic
	TimeStampBegin() objc.IObject /* cross-framework: NSDate */
	TimeStampEnd() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDiagnosticPayload */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDiagnosticPayload */
// Alloc allocates a new instance without initialization.
func (mc _MXDiagnosticPayloadClass) Alloc() MXDiagnosticPayload {
	rv := objc.Send[MXDiagnosticPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDiagnosticPayload */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDiagnosticPayload *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDiagnosticPayload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDiagnosticPayload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDiagnosticPayload */

// Returns the results of the payload as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/dictionaryRepresentation()
func (m_ MXDiagnosticPayload) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns the contents of the payload in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/jsonRepresentation()
func (m_ MXDiagnosticPayload) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDiagnosticPayload */

// The diagnostic reports for fatal and nonfatal CPU exceptions for the app during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/cpuExceptionDiagnostics
func (m_ MXDiagnosticPayload) CpuExceptionDiagnostics() []MXCPUExceptionDiagnostic {
	rv := objc.Send[[]MXCPUExceptionDiagnostic](m_.ID, objc.Sel("cpuExceptionDiagnostics"))
	return rv
}/* debug [instance_properties/getter]: cpuExceptionDiagnostics */


// The diagnostic reports for app crashes during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/crashDiagnostics
func (m_ MXDiagnosticPayload) CrashDiagnostics() []MXCrashDiagnostic {
	rv := objc.Send[[]MXCrashDiagnostic](m_.ID, objc.Sel("crashDiagnostics"))
	return rv
}/* debug [instance_properties/getter]: crashDiagnostics */


// The diagnostic reports for disk write exceptions for the app during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/diskWriteExceptionDiagnostics
func (m_ MXDiagnosticPayload) DiskWriteExceptionDiagnostics() []MXDiskWriteExceptionDiagnostic {
	rv := objc.Send[[]MXDiskWriteExceptionDiagnostic](m_.ID, objc.Sel("diskWriteExceptionDiagnostics"))
	return rv
}/* debug [instance_properties/getter]: diskWriteExceptionDiagnostics */


// The diagnostic reports for times when the app was too busy to handle input responsively during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/hangDiagnostics
func (m_ MXDiagnosticPayload) HangDiagnostics() []MXHangDiagnostic {
	rv := objc.Send[[]MXHangDiagnostic](m_.ID, objc.Sel("hangDiagnostics"))
	return rv
}/* debug [instance_properties/getter]: hangDiagnostics */


// The starting time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampBegin
func (m_ MXDiagnosticPayload) TimeStampBegin() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampBegin"))
	return rv
}/* debug [instance_properties/getter]: timeStampBegin */


// The ending time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiagnosticPayload/timeStampEnd
func (m_ MXDiagnosticPayload) TimeStampEnd() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampEnd"))
	return rv
}/* debug [instance_properties/getter]: timeStampEnd */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDiagnosticPayload */


