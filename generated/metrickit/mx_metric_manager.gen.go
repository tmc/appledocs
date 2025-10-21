// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MXMetricManager] class.
type IMXMetricManager interface {
	objectivec.IObject
	AddSubscriber(subscriber objc.ID)
	RemoveSubscriber(subscriber objc.ID)
}

// The shared object that registers you to receive metrics, creates logs for custom metrics, and gives access to past reports.
//
// The shared object manages your subscription for receiving on-device daily metrics. MetricKit starts accumulating reports for your app after calling for the first time. To receive the reports, call with an object that adopts the protocol. The system then delivers metric reports at most once per day, and diagnostic reports immediately in iOS 15 and later and macOS 12 and later. The reports contain the metrics from the past 24 hours and any previously undelivered daily reports. To pause receiving reports, call . The calls to add a subscriber and for receiving reports are safe to use in performance-sensitive code, such as app launch. The snippet below shows a simple class for using MetricKit.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXMetricManagerClass) Alloc() MXMetricManager {
	rv := objc.Send[MXMetricManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Starts to measure an extended launch task with the given task identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/extendLaunchMeasurement(forTaskID:)
func (mc _MXMetricManagerClass) ExtendLaunchMeasurementForTaskIDError(taskID unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("extendLaunchMeasurementForTaskID:error:"), taskID, error_)
	return rv
}

// Signals the end of an extended launch task.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/finishExtendedLaunchMeasurement(forTaskID:)
func (mc _MXMetricManagerClass) FinishExtendedLaunchMeasurementForTaskIDError(taskID unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("finishExtendedLaunchMeasurementForTaskID:error:"), taskID, error_)
	return rv
}

// Returns a log handle used for writing custom metric events.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/makeLogHandle(category:)
func (mc _MXMetricManagerClass) MakeLogHandleWithCategory(category string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("makeLogHandleWithCategory:"), objc.String(category))
	return rv
}

// An object that returns the shared metrics manager instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/shared
func (mc _MXMetricManagerClass) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("sharedManager"))
	return rv
}
// Registers to receive a daily report of app metrics from the metrics manager.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/add(_:)
func (m_ MXMetricManager) AddSubscriber(subscriber objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addSubscriber:"), subscriber)
}

// Unsubscribes from daily reports of app metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/remove(_:)
func (m_ MXMetricManager) RemoveSubscriber(subscriber objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeSubscriber:"), subscriber)
}

// Returns an array of the diagnostic reports generated since the last allocation of the shared manager instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/pastDiagnosticPayloads
func (m_ MXMetricManager) PastDiagnosticPayloads() []MXDiagnosticPayload {
	rv := objc.Send[[]MXDiagnosticPayload](m_.ID, objc.Sel("pastDiagnosticPayloads"))
	return rv
}

// Returns an array of the daily metrics reports generated since the last allocation of the shared manager instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/pastPayloads
func (m_ MXMetricManager) PastPayloads() []MXMetricPayload {
	rv := objc.Send[[]MXMetricPayload](m_.ID, objc.Sel("pastPayloads"))
	return rv
}

// An object that returns the shared metrics manager instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricManager/shared
func (m_ MXMetricManager) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sharedManager"))
	return rv
}



