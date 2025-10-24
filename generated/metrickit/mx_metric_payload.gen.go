// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXMetricPayload */


/* debug [class_header]: Header for MXMetricPayload */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXMetricPayload */
// An interface definition for the [MXMetricPayload] class.
type IMXMetricPayload interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXMetricPayload */
	// properties:
	AnimationMetrics() IMXAnimationMetric
	ApplicationExitMetrics() IMXAppExitMetric
	ApplicationLaunchMetrics() IMXAppLaunchMetric
	ApplicationResponsivenessMetrics() IMXAppResponsivenessMetric
	ApplicationTimeMetrics() IMXAppRunTimeMetric
	CellularConditionMetrics() IMXCellularConditionMetric
	CpuMetrics() IMXCPUMetric
	DiskIOMetrics() IMXDiskIOMetric
	DiskSpaceUsageMetrics() IMXDiskSpaceUsageMetric
	DisplayMetrics() IMXDisplayMetric
	GpuMetrics() IMXGPUMetric
	IncludesMultipleApplicationVersions() bool
	LatestApplicationVersion() objc.IObject /* cross-framework: NSString */
	LocationActivityMetrics() IMXLocationActivityMetric
	MemoryMetrics() IMXMemoryMetric
	MetaData() IMXMetaData
	NetworkTransferMetrics() IMXNetworkTransferMetric
	SignpostMetrics() []MXSignpostMetric
	TimeStampBegin() objc.IObject /* cross-framework: NSDate */
	TimeStampEnd() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXMetricPayload */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXMetricPayload */
// Alloc allocates a new instance without initialization.
func (mc _MXMetricPayloadClass) Alloc() MXMetricPayload {
	rv := objc.Send[MXMetricPayload](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXMetricPayload */
// An object that encapsulates a daily metrics report.


// An object that encapsulates a daily metrics report.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXMetricPayload *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXMetricPayload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXMetricPayload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXMetricPayload */

// Returns the results of the payload as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/dictionaryRepresentation()
func (m_ MXMetricPayload) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns the contents of the payload in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/jsonRepresentation()
func (m_ MXMetricPayload) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXMetricPayload */

// The metrics for the responsiveness of app animations for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/animationMetrics
func (m_ MXMetricPayload) AnimationMetrics() IMXAnimationMetric {
	rv := objc.Send[MXAnimationMetric](m_.ID, objc.Sel("animationMetrics"))
	return rv
}/* debug [instance_properties/getter]: animationMetrics */


// The app foreground and background exit metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationExitMetrics
func (m_ MXMetricPayload) ApplicationExitMetrics() IMXAppExitMetric {
	rv := objc.Send[MXAppExitMetric](m_.ID, objc.Sel("applicationExitMetrics"))
	return rv
}/* debug [instance_properties/getter]: applicationExitMetrics */


// The app launch and resume metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationLaunchMetrics
func (m_ MXMetricPayload) ApplicationLaunchMetrics() IMXAppLaunchMetric {
	rv := objc.Send[MXAppLaunchMetric](m_.ID, objc.Sel("applicationLaunchMetrics"))
	return rv
}/* debug [instance_properties/getter]: applicationLaunchMetrics */


// The metrics indicating an app’s responsiveness to user interaction for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationResponsivenessMetrics
func (m_ MXMetricPayload) ApplicationResponsivenessMetrics() IMXAppResponsivenessMetric {
	rv := objc.Send[MXAppResponsivenessMetric](m_.ID, objc.Sel("applicationResponsivenessMetrics"))
	return rv
}/* debug [instance_properties/getter]: applicationResponsivenessMetrics */


// The app foreground and background time metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/applicationTimeMetrics
func (m_ MXMetricPayload) ApplicationTimeMetrics() IMXAppRunTimeMetric {
	rv := objc.Send[MXAppRunTimeMetric](m_.ID, objc.Sel("applicationTimeMetrics"))
	return rv
}/* debug [instance_properties/getter]: applicationTimeMetrics */


// The cellular condition measurements for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/cellularConditionMetrics
func (m_ MXMetricPayload) CellularConditionMetrics() IMXCellularConditionMetric {
	rv := objc.Send[MXCellularConditionMetric](m_.ID, objc.Sel("cellularConditionMetrics"))
	return rv
}/* debug [instance_properties/getter]: cellularConditionMetrics */


// The CPU metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/cpuMetrics
func (m_ MXMetricPayload) CpuMetrics() IMXCPUMetric {
	rv := objc.Send[MXCPUMetric](m_.ID, objc.Sel("cpuMetrics"))
	return rv
}/* debug [instance_properties/getter]: cpuMetrics */


// The storage metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/diskIOMetrics
func (m_ MXMetricPayload) DiskIOMetrics() IMXDiskIOMetric {
	rv := objc.Send[MXDiskIOMetric](m_.ID, objc.Sel("diskIOMetrics"))
	return rv
}/* debug [instance_properties/getter]: diskIOMetrics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/diskSpaceUsageMetrics
func (m_ MXMetricPayload) DiskSpaceUsageMetrics() IMXDiskSpaceUsageMetric {
	rv := objc.Send[MXDiskSpaceUsageMetric](m_.ID, objc.Sel("diskSpaceUsageMetrics"))
	return rv
}/* debug [instance_properties/getter]: diskSpaceUsageMetrics */


// The display metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/displayMetrics
func (m_ MXMetricPayload) DisplayMetrics() IMXDisplayMetric {
	rv := objc.Send[MXDisplayMetric](m_.ID, objc.Sel("displayMetrics"))
	return rv
}/* debug [instance_properties/getter]: displayMetrics */


// The GPU metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/gpuMetrics
func (m_ MXMetricPayload) GpuMetrics() IMXGPUMetric {
	rv := objc.Send[MXGPUMetric](m_.ID, objc.Sel("gpuMetrics"))
	return rv
}/* debug [instance_properties/getter]: gpuMetrics */


// A Boolean indicating if the version of the app changed at least once during the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/includesMultipleApplicationVersions
func (m_ MXMetricPayload) IncludesMultipleApplicationVersions() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("includesMultipleApplicationVersions"))
	return rv
}/* debug [instance_properties/getter]: includesMultipleApplicationVersions */


// The version of the app on the device at the end of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/latestApplicationVersion
func (m_ MXMetricPayload) LatestApplicationVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("latestApplicationVersion"))
	return rv
}/* debug [instance_properties/getter]: latestApplicationVersion */


// The location-tracking activity for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/locationActivityMetrics
func (m_ MXMetricPayload) LocationActivityMetrics() IMXLocationActivityMetric {
	rv := objc.Send[MXLocationActivityMetric](m_.ID, objc.Sel("locationActivityMetrics"))
	return rv
}/* debug [instance_properties/getter]: locationActivityMetrics */


// The memory metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/memoryMetrics
func (m_ MXMetricPayload) MemoryMetrics() IMXMemoryMetric {
	rv := objc.Send[MXMemoryMetric](m_.ID, objc.Sel("memoryMetrics"))
	return rv
}/* debug [instance_properties/getter]: memoryMetrics */


// A set of system-level information for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/metaData
func (m_ MXMetricPayload) MetaData() IMXMetaData {
	rv := objc.Send[MXMetaData](m_.ID, objc.Sel("metaData"))
	return rv
}/* debug [instance_properties/getter]: metaData */


// The network-transfer activity for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/networkTransferMetrics
func (m_ MXMetricPayload) NetworkTransferMetrics() IMXNetworkTransferMetric {
	rv := objc.Send[MXNetworkTransferMetric](m_.ID, objc.Sel("networkTransferMetrics"))
	return rv
}/* debug [instance_properties/getter]: networkTransferMetrics */


// An array of the custom metrics for the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/signpostMetrics
func (m_ MXMetricPayload) SignpostMetrics() []MXSignpostMetric {
	rv := objc.Send[[]MXSignpostMetric](m_.ID, objc.Sel("signpostMetrics"))
	return rv
}/* debug [instance_properties/getter]: signpostMetrics */


// The starting time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/timeStampBegin
func (m_ MXMetricPayload) TimeStampBegin() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampBegin"))
	return rv
}/* debug [instance_properties/getter]: timeStampBegin */


// The ending time of the reporting period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetricPayload/timeStampEnd
func (m_ MXMetricPayload) TimeStampEnd() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("timeStampEnd"))
	return rv
}/* debug [instance_properties/getter]: timeStampEnd */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXMetricPayload */


