// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceStorageBehaviorConfiguration] class.
var (
	MTRDeviceStorageBehaviorConfigurationClass     _MTRDeviceStorageBehaviorConfigurationClass
	MTRDeviceStorageBehaviorConfigurationClassOnce sync.Once
)

func getMTRDeviceStorageBehaviorConfigurationClass() _MTRDeviceStorageBehaviorConfigurationClass {
	MTRDeviceStorageBehaviorConfigurationClassOnce.Do(func() {
		MTRDeviceStorageBehaviorConfigurationClass = _MTRDeviceStorageBehaviorConfigurationClass{objc.GetClass("MTRDeviceStorageBehaviorConfiguration")}
	})
	return MTRDeviceStorageBehaviorConfigurationClass
}

type _MTRDeviceStorageBehaviorConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceStorageBehaviorConfiguration] class.
type IMTRDeviceStorageBehaviorConfiguration interface {
	objectivec.IObject
}

// Class that configures how MTRDevice objects persist their attributes to storage, so as to not overwhelm the underlying storage system.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceStorageBehaviorConfiguration
type MTRDeviceStorageBehaviorConfiguration struct {
	objectivec.Object
}

// MTRDeviceStorageBehaviorConfigurationFrom constructs a [MTRDeviceStorageBehaviorConfiguration] from an unsafe.Pointer.
//
// Class that configures how MTRDevice objects persist their attributes to storage, so as to not overwhelm the underlying storage system.
func MTRDeviceStorageBehaviorConfigurationFrom(ptr unsafe.Pointer) MTRDeviceStorageBehaviorConfiguration {
	return MTRDeviceStorageBehaviorConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceStorageBehaviorConfigurationClass) Alloc() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceStorageBehaviorConfigurationClass) New() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceStorageBehaviorConfiguration) Init() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceStorageBehaviorConfiguration) Autorelease() MTRDeviceStorageBehaviorConfiguration {
	rv := objc.Send[MTRDeviceStorageBehaviorConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceStorageBehaviorConfiguration creates a new MTRDeviceStorageBehaviorConfiguration instance.
func NewMTRDeviceStorageBehaviorConfiguration() MTRDeviceStorageBehaviorConfiguration {
	return getMTRDeviceStorageBehaviorConfigurationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/devicereportingexcessivelyintervalthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) DeviceReportingExcessivelyIntervalThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceReportingExcessivelyIntervalThreshold"))
	return rv
}


// SetDeviceReportingExcessivelyIntervalThreshold sets the value of the deviceReportingExcessivelyIntervalThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/devicereportingexcessivelyintervalthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) SetDeviceReportingExcessivelyIntervalThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceReportingExcessivelyIntervalThreshold:"), value)
}

// If disableStorageBehaviorOptimization is set to YES, then all the waiting mechanism as described above
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/disablestoragebehavioroptimization
func (m_ MTRDeviceStorageBehaviorConfiguration) DisableStorageBehaviorOptimization() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("disableStorageBehaviorOptimization"))
	return rv
}


// SetDisableStorageBehaviorOptimization sets the value of the disableStorageBehaviorOptimization property.
// If disableStorageBehaviorOptimization is set to YES, then all the waiting mechanism as described above

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/disablestoragebehavioroptimization
func (m_ MTRDeviceStorageBehaviorConfiguration) SetDisableStorageBehaviorOptimization(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisableStorageBehaviorOptimization:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/recentreporttimesmaxcount
func (m_ MTRDeviceStorageBehaviorConfiguration) RecentReportTimesMaxCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("recentReportTimesMaxCount"))
	return rv
}


// SetRecentReportTimesMaxCount sets the value of the recentReportTimesMaxCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/recentreporttimesmaxcount
func (m_ MTRDeviceStorageBehaviorConfiguration) SetRecentReportTimesMaxCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRecentReportTimesMaxCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaymaxmultiplier
func (m_ MTRDeviceStorageBehaviorConfiguration) ReportToPersistenceDelayMaxMultiplier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("reportToPersistenceDelayMaxMultiplier"))
	return rv
}


// SetReportToPersistenceDelayMaxMultiplier sets the value of the reportToPersistenceDelayMaxMultiplier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaymaxmultiplier
func (m_ MTRDeviceStorageBehaviorConfiguration) SetReportToPersistenceDelayMaxMultiplier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReportToPersistenceDelayMaxMultiplier:"), value)
}

// If any of these properties are set to be out of the documented limits, these default values will
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaytime
func (m_ MTRDeviceStorageBehaviorConfiguration) ReportToPersistenceDelayTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("reportToPersistenceDelayTime"))
	return rv
}


// SetReportToPersistenceDelayTime sets the value of the reportToPersistenceDelayTime property.
// If any of these properties are set to be out of the documented limits, these default values will

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaytime
func (m_ MTRDeviceStorageBehaviorConfiguration) SetReportToPersistenceDelayTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReportToPersistenceDelayTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaytimemax
func (m_ MTRDeviceStorageBehaviorConfiguration) ReportToPersistenceDelayTimeMax() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("reportToPersistenceDelayTimeMax"))
	return rv
}


// SetReportToPersistenceDelayTimeMax sets the value of the reportToPersistenceDelayTimeMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/reporttopersistencedelaytimemax
func (m_ MTRDeviceStorageBehaviorConfiguration) SetReportToPersistenceDelayTimeMax(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReportToPersistenceDelayTimeMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/timebetweenreportstooshortminthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) TimeBetweenReportsTooShortMinThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeBetweenReportsTooShortMinThreshold"))
	return rv
}


// SetTimeBetweenReportsTooShortMinThreshold sets the value of the timeBetweenReportsTooShortMinThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/timebetweenreportstooshortminthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) SetTimeBetweenReportsTooShortMinThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeBetweenReportsTooShortMinThreshold:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/timebetweenreportstooshortthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) TimeBetweenReportsTooShortThreshold() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeBetweenReportsTooShortThreshold"))
	return rv
}


// SetTimeBetweenReportsTooShortThreshold sets the value of the timeBetweenReportsTooShortThreshold property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicestoragebehaviorconfiguration/timebetweenreportstooshortthreshold
func (m_ MTRDeviceStorageBehaviorConfiguration) SetTimeBetweenReportsTooShortThreshold(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeBetweenReportsTooShortThreshold:"), value)
}



