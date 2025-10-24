// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXMetaData] class.
var (
	MXMetaDataClass     _MXMetaDataClass
	MXMetaDataClassOnce sync.Once
)

func getMXMetaDataClass() _MXMetaDataClass {
	MXMetaDataClassOnce.Do(func() {
		MXMetaDataClass = _MXMetaDataClass{objc.GetClass("MXMetaData")}
	})
	return MXMetaDataClass
}

type _MXMetaDataClass struct {
	class objc.Class
}

// An interface definition for the [MXMetaData] class.
type IMXMetaData interface {
	objectivec.IObject
	// properties:
	ApplicationBuildVersion() objc.IObject /* cross-framework: NSString */
	BundleIdentifier() objc.IObject /* cross-framework: NSString */
	DeviceType() objc.IObject /* cross-framework: NSString */
	IsTestFlightApp() bool
	LowPowerModeEnabled() bool
	OsVersion() objc.IObject /* cross-framework: NSString */
	Pid() unsafe.Pointer
	PlatformArchitecture() objc.IObject /* cross-framework: NSString */
	RegionFormat() objc.IObject /* cross-framework: NSString */
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
	JSONRepresentation() objc.IObject /* cross-framework: Data */
}

// An object containing system-level information about the device.


// An object containing system-level information about the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData
type MXMetaData struct {
	objectivec.Object
}

// MXMetaDataFrom constructs a [MXMetaData] from an unsafe.Pointer.
//
// An object containing system-level information about the device.
func MXMetaDataFrom(ptr unsafe.Pointer) MXMetaData {
	return MXMetaData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXMetaDataClass) Alloc() MXMetaData {
	rv := objc.Send[MXMetaData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXMetaDataClass) New() MXMetaData {
	rv := objc.Send[MXMetaData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXMetaData) Init() MXMetaData {
	rv := objc.Send[MXMetaData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXMetaData) Autorelease() MXMetaData {
	rv := objc.Send[MXMetaData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXMetaData creates a new MXMetaData instance.
func NewMXMetaData() MXMetaData {
	return getMXMetaDataClass().New()
}



// Returns the contents of the metadata in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/jsonRepresentation()
func (m_ MXMetaData) JSONRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}


// The value of the bundle version key in the app’s property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/applicationBuildVersion
func (m_ MXMetaData) ApplicationBuildVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationBuildVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/bundleIdentifier
func (m_ MXMetaData) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// The hardware identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/deviceType
func (m_ MXMetaData) DeviceType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("deviceType"))
	return rv
}


// Indicates whether the app is registered with TestFlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/isTestFlightApp
func (m_ MXMetaData) IsTestFlightApp() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isTestFlightApp"))
	return rv
}


// Indicates whether low power mode is enabled on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/lowPowerModeEnabled
func (m_ MXMetaData) LowPowerModeEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("lowPowerModeEnabled"))
	return rv
}


// The version of the OS on the device including the type of OS, version number, and build number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/osVersion
func (m_ MXMetaData) OsVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("osVersion"))
	return rv
}


// The process ID (PID) of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/pid
func (m_ MXMetaData) Pid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pid"))
	return rv
}


// The name of the processor architecture for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/platformArchitecture
func (m_ MXMetaData) PlatformArchitecture() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("platformArchitecture"))
	return rv
}


// The short country code for the region format setting of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetaData/regionFormat
func (m_ MXMetaData) RegionFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("regionFormat"))
	return rv
}


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXMetaData) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}


