// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXBackgroundExitData */


/* debug [class_header]: Header for MXBackgroundExitData */
// The class instance for the [MXBackgroundExitData] class.
var (
	MXBackgroundExitDataClass     _MXBackgroundExitDataClass
	MXBackgroundExitDataClassOnce sync.Once
)

func getMXBackgroundExitDataClass() _MXBackgroundExitDataClass {
	MXBackgroundExitDataClassOnce.Do(func() {
		MXBackgroundExitDataClass = _MXBackgroundExitDataClass{objc.GetClass("MXBackgroundExitData")}
	})
	return MXBackgroundExitDataClass
}

type _MXBackgroundExitDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXBackgroundExitData */
// An interface definition for the [MXBackgroundExitData] class.
type IMXBackgroundExitData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXBackgroundExitData */
	// properties:
	CumulativeAbnormalExitCount() uint
	CumulativeAppWatchdogExitCount() uint
	CumulativeBackgroundTaskAssertionTimeoutExitCount() uint
	CumulativeBadAccessExitCount() uint
	CumulativeCPUResourceLimitExitCount() uint
	CumulativeIllegalInstructionExitCount() uint
	CumulativeMemoryPressureExitCount() uint
	CumulativeMemoryResourceLimitExitCount() uint
	CumulativeNormalAppExitCount() uint
	CumulativeSuspendedWithLockedFileExitCount() uint
	BackgroundExitData() IMXBackgroundExitData
	SetBackgroundExitData(value IMXBackgroundExitData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXBackgroundExitData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXBackgroundExitData */
// Alloc allocates a new instance without initialization.
func (mc _MXBackgroundExitDataClass) Alloc() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXBackgroundExitDataClass) New() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXBackgroundExitData) Init() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXBackgroundExitData) Autorelease() MXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXBackgroundExitData creates a new MXBackgroundExitData instance.
func NewMXBackgroundExitData() MXBackgroundExitData {
	return getMXBackgroundExitDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXBackgroundExitData */
// An object representing counts for the different types of background app exits.


// An object representing counts for the different types of background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData
type MXBackgroundExitData struct {
	objectivec.Object
}

// MXBackgroundExitDataFrom constructs a [MXBackgroundExitData] from an unsafe.Pointer.
//
// An object representing counts for the different types of background app exits.
func MXBackgroundExitDataFrom(ptr unsafe.Pointer) MXBackgroundExitData {
	return MXBackgroundExitData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXBackgroundExitData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXBackgroundExitData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXBackgroundExitData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXBackgroundExitData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXBackgroundExitData */

// The number of times the app exited abnormally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeAbnormalExitCount
func (m_ MXBackgroundExitData) CumulativeAbnormalExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAbnormalExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeAbnormalExitCount */


// The number of times the system watchdog terminated the app from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeAppWatchdogExitCount
func (m_ MXBackgroundExitData) CumulativeAppWatchdogExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAppWatchdogExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeAppWatchdogExitCount */


// The number of times the system terminated the app from the background for exceeding the allocated time for a background task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeBackgroundTaskAssertionTimeoutExitCount
func (m_ MXBackgroundExitData) CumulativeBackgroundTaskAssertionTimeoutExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBackgroundTaskAssertionTimeoutExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBackgroundTaskAssertionTimeoutExitCount */


// The number of times the system terminated the app from the background for attempting an invalid memory access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeBadAccessExitCount
func (m_ MXBackgroundExitData) CumulativeBadAccessExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBadAccessExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBadAccessExitCount */


// The number of times the system terminated the app from the background for using too much CPU time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeCPUResourceLimitExitCount
func (m_ MXBackgroundExitData) CumulativeCPUResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeCPUResourceLimitExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCPUResourceLimitExitCount */


// The number of times the system terminated the app from the background for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeIllegalInstructionExitCount
func (m_ MXBackgroundExitData) CumulativeIllegalInstructionExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeIllegalInstructionExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeIllegalInstructionExitCount */


// The number of times the system terminated the app from the background to free up memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeMemoryPressureExitCount
func (m_ MXBackgroundExitData) CumulativeMemoryPressureExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryPressureExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeMemoryPressureExitCount */


// The number of times the system terminated the app from the background for using too much memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeMemoryResourceLimitExitCount
func (m_ MXBackgroundExitData) CumulativeMemoryResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryResourceLimitExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeMemoryResourceLimitExitCount */


// The number of times the app exited normally from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeNormalAppExitCount
func (m_ MXBackgroundExitData) CumulativeNormalAppExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeNormalAppExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeNormalAppExitCount */


// The number of times the system terminated the app from the background while being suspended and having file locks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXBackgroundExitData/cumulativeSuspendedWithLockedFileExitCount
func (m_ MXBackgroundExitData) CumulativeSuspendedWithLockedFileExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeSuspendedWithLockedFileExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeSuspendedWithLockedFileExitCount */


// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/backgroundexitdata
func (m_ MXBackgroundExitData) BackgroundExitData() IMXBackgroundExitData {
	rv := objc.Send[MXBackgroundExitData](m_.ID, objc.Sel("backgroundExitData"))
	return rv
}/* debug [instance_properties/getter]: backgroundExitData */


// The metrics for the background app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/backgroundexitdata
func (m_ MXBackgroundExitData) SetBackgroundExitData(value IMXBackgroundExitData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundExitData:"), value)
}/* debug [instance_properties/setter]: backgroundExitData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXBackgroundExitData */



