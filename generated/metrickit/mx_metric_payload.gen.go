// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXMetricPayload] class.
var (
	MXMetricPayloadClass     _MXMetricPayloadClass
	MXMetricPayloadClassOnce sync.Once
)

func getMXMetricPayloadClass() _MXMetricPayloadClass {
	MXMetricPayloadClassOnce.Do(func() {
		MXMetricPayloadClass = _MXMetricPayloadClass{objc.GetClass("MXMetricPayload")}
	})
	return MXMetricPayloadClass
}

type _MXMetricPayloadClass struct {
	class objc.Class
}

// An interface definition for the [MXMetricPayload] class.
type IMXMetricPayload interface {
	objectivec.IObject
	DictionaryRepresentation() unsafe.Pointer
	JSONRepresentation() unsafe.Pointer
}

// An object that encapsulates a daily metrics report.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload
type MXMetricPayload struct {
	objectivec.Object
}

// MXMetricPayloadFrom constructs a [MXMetricPayload] from an unsafe.Pointer.
//
// An object that encapsulates a daily metrics report.
func MXMetricPayloadFrom(ptr unsafe.Pointer) MXMetricPayload {
	return MXMetricPayload{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXMetricPayloadClass) Alloc() MXMetricPayload {
	rv := objc.Send[MXMetricPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXMetricPayloadClass) New() MXMetricPayload {
	rv := objc.Send[MXMetricPayload](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXMetricPayload) Init() MXMetricPayload {
	rv := objc.Send[MXMetricPayload](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXMetricPayload) Autorelease() MXMetricPayload {
	rv := objc.Send[MXMetricPayload](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXMetricPayload creates a new MXMetricPayload instance.
func NewMXMetricPayload() MXMetricPayload {
	return getMXMetricPayloadClass().New()
}


// Returns the results of the payload as a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/DictionaryRepresentation-1nrgx
func (m_ MXMetricPayload) DictionaryRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("DictionaryRepresentation"))
	return rv
}

// Returns the contents of the payload in JSON format.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/jsonRepresentation()
func (m_ MXMetricPayload) JSONRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}

// The metrics for the responsiveness of app animations for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/animationMetrics
func (m_ MXMetricPayload) AnimationMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("animationMetrics"))
	return rv
}

// The app foreground and background exit metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationExitMetrics
func (m_ MXMetricPayload) ApplicationExitMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("applicationExitMetrics"))
	return rv
}

// The app launch and resume metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationLaunchMetrics
func (m_ MXMetricPayload) ApplicationLaunchMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("applicationLaunchMetrics"))
	return rv
}

// The metrics indicating an app’s responsiveness to user interaction for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationResponsivenessMetrics
func (m_ MXMetricPayload) ApplicationResponsivenessMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("applicationResponsivenessMetrics"))
	return rv
}

// The app foreground and background time metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationTimeMetrics
func (m_ MXMetricPayload) ApplicationTimeMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("applicationTimeMetrics"))
	return rv
}

// The cellular condition measurements for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/cellularConditionMetrics
func (m_ MXMetricPayload) CellularConditionMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cellularConditionMetrics"))
	return rv
}

// The CPU metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/cpuMetrics
func (m_ MXMetricPayload) CpuMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cpuMetrics"))
	return rv
}

// The storage metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/diskIOMetrics
func (m_ MXMetricPayload) DiskIOMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("diskIOMetrics"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/diskSpaceUsageMetrics
func (m_ MXMetricPayload) DiskSpaceUsageMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("diskSpaceUsageMetrics"))
	return rv
}

// The display metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/displayMetrics
func (m_ MXMetricPayload) DisplayMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("displayMetrics"))
	return rv
}

// The GPU metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/gpuMetrics
func (m_ MXMetricPayload) GpuMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("gpuMetrics"))
	return rv
}

// A Boolean indicating if the version of the app changed at least once during the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/includesMultipleApplicationVersions
func (m_ MXMetricPayload) IncludesMultipleApplicationVersions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("includesMultipleApplicationVersions"))
	return rv
}

// The version of the app on the device at the end of the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/latestApplicationVersion
func (m_ MXMetricPayload) LatestApplicationVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("latestApplicationVersion"))
	return rv
}

// The location-tracking activity for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/locationActivityMetrics
func (m_ MXMetricPayload) LocationActivityMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("locationActivityMetrics"))
	return rv
}

// The memory metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/memoryMetrics
func (m_ MXMetricPayload) MemoryMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("memoryMetrics"))
	return rv
}

// A set of system-level information for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/metaData
func (m_ MXMetricPayload) MetaData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("metaData"))
	return rv
}

// The network-transfer activity for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/networkTransferMetrics
func (m_ MXMetricPayload) NetworkTransferMetrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("networkTransferMetrics"))
	return rv
}

// An array of the custom metrics for the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/signpostMetrics
func (m_ MXMetricPayload) SignpostMetrics() []MXSignpostMetric {
	rv := objc.Send[[]MXSignpostMetric](m_.ID, objc.Sel("signpostMetrics"))
	return rv
}

// The starting time of the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/timeStampBegin
func (m_ MXMetricPayload) TimeStampBegin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeStampBegin"))
	return rv
}

// The ending time of the reporting period.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/timeStampEnd
func (m_ MXMetricPayload) TimeStampEnd() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeStampEnd"))
	return rv
}



