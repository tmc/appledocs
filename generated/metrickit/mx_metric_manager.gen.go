// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXMetricManager */


/* debug [class_header]: Header for MXMetricManager */
// The class instance for the [MXMetricManager] class.
var (
	MXMetricManagerClass     _MXMetricManagerClass
	MXMetricManagerClassOnce sync.Once
)

func getMXMetricManagerClass() _MXMetricManagerClass {
	MXMetricManagerClassOnce.Do(func() {
		MXMetricManagerClass = _MXMetricManagerClass{objc.GetClass("MXMetricManager")}
	})
	return MXMetricManagerClass
}

type _MXMetricManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXMetricManager */
// An interface definition for the [MXMetricManager] class.
type IMXMetricManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXMetricManager */
	// properties:
	PastDiagnosticPayloads() []MXDiagnosticPayload
	PastPayloads() []MXMetricPayload
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXMetricManager */
	// methods:
	AddSubscriber(subscriber unsafe.Pointer)
	RemoveSubscriber(subscriber unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXMetricManager */
// Alloc allocates a new instance without initialization.
func (mc _MXMetricManagerClass) Alloc() MXMetricManager {
	rv := objc.Send[MXMetricManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXMetricManagerClass) New() MXMetricManager {
	rv := objc.Send[MXMetricManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXMetricManager) Init() MXMetricManager {
	rv := objc.Send[MXMetricManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXMetricManager) Autorelease() MXMetricManager {
	rv := objc.Send[MXMetricManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXMetricManager creates a new MXMetricManager instance.
func NewMXMetricManager() MXMetricManager {
	return getMXMetricManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXMetricManager */
// The shared object that registers you to receive metrics, creates logs for custom metrics, and gives access to past reports.
//
// The shared object manages your subscription for receiving on-device daily metrics. MetricKit starts accumulating reports for your app after calling for the first time. To receive the reports, call with an object that adopts the protocol. The system then delivers metric reports at most once per day, and diagnostic reports immediately in iOS 15 and later and macOS 12 and later. The reports contain the metrics from the past 24 hours and any previously undelivered daily reports. To pause receiving reports, call . The calls to add a subscriber and for receiving reports are safe to use in performance-sensitive code, such as app launch. The snippet below shows a simple class for using MetricKit.


// The shared object that registers you to receive metrics, creates logs for custom metrics, and gives access to past reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager
type MXMetricManager struct {
	objectivec.Object
}

// MXMetricManagerFrom constructs a [MXMetricManager] from an unsafe.Pointer.
//
// The shared object that registers you to receive metrics, creates logs for custom metrics, and gives access to past reports.
func MXMetricManagerFrom(ptr unsafe.Pointer) MXMetricManager {
	return MXMetricManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXMetricManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXMetricManager */

// Starts to measure an extended launch task with the given task identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/extendLaunchMeasurement(forTaskID:)
func (mc _MXMetricManagerClass) ExtendLaunchMeasurementForTaskIDError(taskID MXLaunchTaskID /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("extendLaunchMeasurementForTaskID:error:"), taskID, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExtendLaunchMeasurementForTaskIDError) */


// Signals the end of an extended launch task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/finishExtendedLaunchMeasurement(forTaskID:)
func (mc _MXMetricManagerClass) FinishExtendedLaunchMeasurementForTaskIDError(taskID MXLaunchTaskID /* typedef */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("finishExtendedLaunchMeasurementForTaskID:error:"), taskID, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FinishExtendedLaunchMeasurementForTaskIDError) */


// Returns a log handle used for writing custom metric events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/makeLogHandle(category:)
func (mc _MXMetricManagerClass) MakeLogHandleWithCategory(category objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("makeLogHandleWithCategory:"), category)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MakeLogHandleWithCategory) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXMetricManager */

// An object that returns the shared metrics manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/shared
func (mc _MXMetricManagerClass) SharedManager() MXMetricManager {
	rv := objc.Send[MXMetricManager](objc.ID(mc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_properties_class/property]: sharedManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXMetricManager */

// Registers to receive a daily report of app metrics from the metrics manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/add(_:)
func (m_ MXMetricManager) AddSubscriber(subscriber unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addSubscriber:"), subscriber)
}/* debug [instance_methods/method]: AddSubscriber */


// Unsubscribes from daily reports of app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/remove(_:)
func (m_ MXMetricManager) RemoveSubscriber(subscriber unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeSubscriber:"), subscriber)
}/* debug [instance_methods/method]: RemoveSubscriber */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXMetricManager */

// Returns an array of the diagnostic reports generated since the last allocation of the shared manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/pastDiagnosticPayloads
func (m_ MXMetricManager) PastDiagnosticPayloads() []MXDiagnosticPayload {
	rv := objc.Send[[]MXDiagnosticPayload](m_.ID, objc.Sel("pastDiagnosticPayloads"))
	return rv
}/* debug [instance_properties/getter]: pastDiagnosticPayloads */


// Returns an array of the daily metrics reports generated since the last allocation of the shared manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/pastPayloads
func (m_ MXMetricManager) PastPayloads() []MXMetricPayload {
	rv := objc.Send[[]MXMetricPayload](m_.ID, objc.Sel("pastPayloads"))
	return rv
}/* debug [instance_properties/getter]: pastPayloads */


// An object that returns the shared metrics manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/shared
func (m_ MXMetricManager) SharedManager() IMXMetricManager {
	rv := objc.Send[MXMetricManager](m_.ID, objc.Sel("sharedManager"))
	return rv
}/* debug [instance_properties/getter]: sharedManager */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXMetricManager */



