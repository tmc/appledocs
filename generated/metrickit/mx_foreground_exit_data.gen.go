// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXForegroundExitData */


/* debug [class_header]: Header for MXForegroundExitData */
// The class instance for the [MXForegroundExitData] class.
var (
	MXForegroundExitDataClass     _MXForegroundExitDataClass
	MXForegroundExitDataClassOnce sync.Once
)

func getMXForegroundExitDataClass() _MXForegroundExitDataClass {
	MXForegroundExitDataClassOnce.Do(func() {
		MXForegroundExitDataClass = _MXForegroundExitDataClass{objc.GetClass("MXForegroundExitData")}
	})
	return MXForegroundExitDataClass
}

type _MXForegroundExitDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXForegroundExitData */
// An interface definition for the [MXForegroundExitData] class.
type IMXForegroundExitData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXForegroundExitData */
	// properties:
	CumulativeAbnormalExitCount() uint
	CumulativeAppWatchdogExitCount() uint
	CumulativeBadAccessExitCount() uint
	CumulativeIllegalInstructionExitCount() uint
	CumulativeMemoryResourceLimitExitCount() uint
	CumulativeNormalAppExitCount() uint
	ForegroundExitData() IMXForegroundExitData
	SetForegroundExitData(value IMXForegroundExitData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXForegroundExitData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXForegroundExitData */
// Alloc allocates a new instance without initialization.
func (mc _MXForegroundExitDataClass) Alloc() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXForegroundExitDataClass) New() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXForegroundExitData) Init() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXForegroundExitData) Autorelease() MXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXForegroundExitData creates a new MXForegroundExitData instance.
func NewMXForegroundExitData() MXForegroundExitData {
	return getMXForegroundExitDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXForegroundExitData */
// An object representing counts for the different types of foreground app exits.


// An object representing counts for the different types of foreground app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData
type MXForegroundExitData struct {
	objectivec.Object
}

// MXForegroundExitDataFrom constructs a [MXForegroundExitData] from an unsafe.Pointer.
//
// An object representing counts for the different types of foreground app exits.
func MXForegroundExitDataFrom(ptr unsafe.Pointer) MXForegroundExitData {
	return MXForegroundExitData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXForegroundExitData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXForegroundExitData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXForegroundExitData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXForegroundExitData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXForegroundExitData */

// The number of times the app exited abnormally from the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeAbnormalExitCount
func (m_ MXForegroundExitData) CumulativeAbnormalExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAbnormalExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeAbnormalExitCount */


// The number of times the system watchdog terminated the app from the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeAppWatchdogExitCount
func (m_ MXForegroundExitData) CumulativeAppWatchdogExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeAppWatchdogExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeAppWatchdogExitCount */


// The number of times the system terminated the app from the foreground for attempting an invalid memory access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeBadAccessExitCount
func (m_ MXForegroundExitData) CumulativeBadAccessExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeBadAccessExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeBadAccessExitCount */


// The number of times the system terminated the app from the foreground for attempting to execute an illegal or undefined instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeIllegalInstructionExitCount
func (m_ MXForegroundExitData) CumulativeIllegalInstructionExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeIllegalInstructionExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeIllegalInstructionExitCount */


// The number of times the system terminated the app from the foreground for using too much memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeMemoryResourceLimitExitCount
func (m_ MXForegroundExitData) CumulativeMemoryResourceLimitExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeMemoryResourceLimitExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeMemoryResourceLimitExitCount */


// The number of times the app exited normally from the foreground.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXForegroundExitData/cumulativeNormalAppExitCount
func (m_ MXForegroundExitData) CumulativeNormalAppExitCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("cumulativeNormalAppExitCount"))
	return rv
}/* debug [instance_properties/getter]: cumulativeNormalAppExitCount */


// The metrics for the foreground app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/foregroundexitdata
func (m_ MXForegroundExitData) ForegroundExitData() IMXForegroundExitData {
	rv := objc.Send[MXForegroundExitData](m_.ID, objc.Sel("foregroundExitData"))
	return rv
}/* debug [instance_properties/getter]: foregroundExitData */


// The metrics for the foreground app exits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxappexitmetric/foregroundexitdata
func (m_ MXForegroundExitData) SetForegroundExitData(value IMXForegroundExitData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForegroundExitData:"), value)
}/* debug [instance_properties/setter]: foregroundExitData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXForegroundExitData */



