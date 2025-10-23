// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProcessInfo] class.
var (
	ProcessInfoClass     _ProcessInfoClass
	ProcessInfoClassOnce sync.Once
)

func getProcessInfoClass() _ProcessInfoClass {
	ProcessInfoClassOnce.Do(func() {
		ProcessInfoClass = _ProcessInfoClass{objc.GetClass("NSProcessInfo")}
	})
	return ProcessInfoClass
}

type _ProcessInfoClass struct {
	class objc.Class
}

// An interface definition for the [ProcessInfo] class.
type IProcessInfo interface {
	objectivec.IObject
	// properties:
	ActiveProcessorCount() uint /* primitive/slice/pointer. */
	Arguments() []string /* primitive/slice/pointer. */
	AutomaticTerminationSupportEnabled() bool /* primitive/slice/pointer. */
	SetAutomaticTerminationSupportEnabled(value bool /* primitive/slice/pointer. */)
	Environment() IDictionary /* already interface */
	FullUserName() IString
	GloballyUniqueString() IString
	HostName() IString
	LowPowerModeEnabled() bool /* primitive/slice/pointer. */
	MacCatalystApp() bool /* primitive/slice/pointer. */
	IOSAppOnMac() bool /* primitive/slice/pointer. */
	IOSAppOnVision() bool /* primitive/slice/pointer. */
	OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */
	OperatingSystemVersionString() IString
	PhysicalMemory() uint64 /* primitive/slice/pointer. */
	ProcessIdentifier() int /* primitive/slice/pointer. */
	ProcessName() IString
	SetProcessName(value IString)
	ProcessorCount() uint /* primitive/slice/pointer. */
	SystemUptime() objc.IObject /* cross-framework: TimeInterval */
	ThermalState() ProcessInfoThermalState
	UserName() IString
	IsLowPowerModeEnabled() bool /* primitive/slice/pointer. */
	SetIsLowPowerModeEnabled(value bool /* primitive/slice/pointer. */)
	IsMacCatalystApp() bool /* primitive/slice/pointer. */
	SetIsMacCatalystApp(value bool /* primitive/slice/pointer. */)
	IsiOSAppOnMac() bool /* primitive/slice/pointer. */
	SetIsiOSAppOnMac(value bool /* primitive/slice/pointer. */)
	IsiOSAppOnVision() bool /* primitive/slice/pointer. */
	SetIsiOSAppOnVision(value bool /* primitive/slice/pointer. */)
	// methods:
	BeginActivityWithOptionsReason(options ActivityOptions, reason IString) objc.ID
	DisableAutomaticTermination(reason IString)
	DisableSuddenTermination()
	EnableAutomaticTermination(reason IString)
	EnableSuddenTermination()
	EndActivity(activity objectivec.IObject)
	HasPerformanceProfile(performanceProfile ProcessPerformanceProfile /* not a class type */) bool /* primitive/slice/pointer. */
	IsDeviceCertifiedFor(performanceTier DeviceCertification /* not a class type */) bool /* primitive/slice/pointer. */
	IsOperatingSystemAtLeastVersion(version objc.IObject /* cross-framework OperatingSystemVersion */) bool /* primitive/slice/pointer. */
	PerformActivityWithOptionsReasonUsingBlock(options ActivityOptions, reason IString, block unsafe.Pointer)
	PerformExpiringActivityWithReasonUsingBlock(reason IString, block unsafe.Pointer)
}

// A collection of information about the current process.
//
// Each process has a single, shared object known as a that can return information such as arguments, environment variables, host name, and process name. The class method returns the shared agent for the current process. For example, the following line returns the object, which then provides the name of the current process: The class also includes the property, which returns an structure identifying the operating system version on which the process is executing. objects attempt to interpret environment variables and command-line arguments in the user’s default C string encoding if they can’t convert to Unicode as UTF-8 strings. If neither the Unicode nor C string conversion works, the object ignores these values.


// A collection of information about the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo
type ProcessInfo struct {
	objectivec.Object
}

// ProcessInfoFrom constructs a [ProcessInfo] from an unsafe.Pointer.
//
// A collection of information about the current process.
func ProcessInfoFrom(ptr unsafe.Pointer) ProcessInfo {
	return ProcessInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProcessInfoClass) Alloc() ProcessInfo {
	rv := objc.Send[ProcessInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProcessInfoClass) New() ProcessInfo {
	rv := objc.Send[ProcessInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProcessInfo) Init() ProcessInfo {
	rv := objc.Send[ProcessInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProcessInfo) Autorelease() ProcessInfo {
	rv := objc.Send[ProcessInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProcessInfo creates a new ProcessInfo instance.
func NewProcessInfo() ProcessInfo {
	return getProcessInfoClass().New()
}



// Returns the process information agent for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processInfo
func (pc _ProcessInfoClass) ProcessInfo() ProcessInfo {
	rv := objc.Send[ProcessInfo](objc.ID(pc.class), objc.Sel("processInfo"))
	return rv
}

// Begin an activity using the given options and reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/beginActivity(options:reason:)
func (p_ ProcessInfo) BeginActivityWithOptionsReason(options ActivityOptions, reason IString) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("beginActivityWithOptions:reason:"), options, reason)
	return rv
}


// Disables automatic termination for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/disableAutomaticTermination(_:)
func (p_ ProcessInfo) DisableAutomaticTermination(reason IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("disableAutomaticTermination:"), reason)
}


// Disables the application for quickly killing using sudden termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/disableSuddenTermination()
func (p_ ProcessInfo) DisableSuddenTermination() {
	objc.Send[objc.ID](p_.ID, objc.Sel("disableSuddenTermination"))
}


// Enables automatic termination for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableAutomaticTermination(_:)
func (p_ ProcessInfo) EnableAutomaticTermination(reason IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("enableAutomaticTermination:"), reason)
}


// Enables the application for quick killing using sudden termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableSuddenTermination()
func (p_ ProcessInfo) EnableSuddenTermination() {
	objc.Send[objc.ID](p_.ID, objc.Sel("enableSuddenTermination"))
}


// Ends the given activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/endActivity(_:)
func (p_ ProcessInfo) EndActivity(activity objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("endActivity:"), activity)
}


// Indicates whether an app is running under a known performance profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/hasPerformanceProfile(_:)
func (p_ ProcessInfo) HasPerformanceProfile(performanceProfile ProcessPerformanceProfile /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasPerformanceProfile:"), performanceProfile)
	return rv
}


// Indicates whether the device supports the requested performance tier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isDeviceCertified(for:)
func (p_ ProcessInfo) IsDeviceCertifiedFor(performanceTier DeviceCertification /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDeviceCertifiedFor:"), performanceTier)
	return rv
}


// Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isOperatingSystemAtLeast(_:)
func (p_ ProcessInfo) IsOperatingSystemAtLeastVersion(version objc.IObject /* cross-framework OperatingSystemVersion */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOperatingSystemAtLeastVersion:"), version)
	return rv
}


// Synchronously perform an activity defined by a given block using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performActivity(options:reason:using:)
func (p_ ProcessInfo) PerformActivityWithOptionsReasonUsingBlock(options ActivityOptions, reason IString, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performActivityWithOptions:reason:usingBlock:"), options, reason, block)
}


// Performs the specified block asynchronously and notifies you if the process is about to be suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p_ ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason IString, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performExpiringActivityWithReason:usingBlock:"), reason, block)
}


// The number of active processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/activeProcessorCount
func (p_ ProcessInfo) ActiveProcessorCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](p_.ID, objc.Sel("activeProcessorCount"))
	return rv
}


// Array of strings with the command-line arguments for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/arguments
func (p_ ProcessInfo) Arguments() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](p_.ID, objc.Sel("arguments"))
	return rv
}


// A Boolean value indicating whether the app supports automatic termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/automaticTerminationSupportEnabled
func (p_ ProcessInfo) AutomaticTerminationSupportEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticTerminationSupportEnabled"))
	return rv
}


// A Boolean value indicating whether the app supports automatic termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/automaticTerminationSupportEnabled
func (p_ ProcessInfo) SetAutomaticTerminationSupportEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticTerminationSupportEnabled:"), value)
}


// The variable names (keys) and their values in the environment from which the process was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/environment
func (p_ ProcessInfo) Environment() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](p_.ID, objc.Sel("environment"))
	return rv
}


// Returns the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/fullUserName
func (p_ ProcessInfo) FullUserName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("fullUserName"))
	return rv
}


// Global unique identifier for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/globallyUniqueString
func (p_ ProcessInfo) GloballyUniqueString() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("globallyUniqueString"))
	return rv
}


// The name of the host computer on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/hostName
func (p_ ProcessInfo) HostName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("hostName"))
	return rv
}


// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isLowPowerModeEnabled
func (p_ ProcessInfo) LowPowerModeEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("lowPowerModeEnabled"))
	return rv
}


// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isMacCatalystApp
func (p_ ProcessInfo) MacCatalystApp() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("macCatalystApp"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnMac
func (p_ ProcessInfo) IOSAppOnMac() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("iOSAppOnMac"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnVision
func (p_ ProcessInfo) IOSAppOnVision() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("iOSAppOnVision"))
	return rv
}


// The version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemVersion
func (p_ ProcessInfo) OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */ {
	rv := objc.Send[OperatingSystemVersion](p_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}


// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemVersionString
func (p_ ProcessInfo) OperatingSystemVersionString() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("operatingSystemVersionString"))
	return rv
}


// The amount of physical memory on the computer in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/physicalMemory
func (p_ ProcessInfo) PhysicalMemory() uint64 /* primitive/slice/pointer. */ {
	rv := objc.Send[uint64](p_.ID, objc.Sel("physicalMemory"))
	return rv
}


// The identifier of the process (often called process ID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processIdentifier
func (p_ ProcessInfo) ProcessIdentifier() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("processIdentifier"))
	return rv
}


// Returns the process information agent for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processInfo
func (p_ ProcessInfo) ProcessInfo() IProcessInfo {
	rv := objc.Send[ProcessInfo](p_.ID, objc.Sel("processInfo"))
	return rv
}


// The name of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processName
func (p_ ProcessInfo) ProcessName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("processName"))
	return rv
}


// The name of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processName
func (p_ ProcessInfo) SetProcessName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessName:"), value)
}


// The number of processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processorCount
func (p_ ProcessInfo) ProcessorCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](p_.ID, objc.Sel("processorCount"))
	return rv
}


// The amount of time the system has been awake since the last time it was restarted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/systemUptime
func (p_ ProcessInfo) SystemUptime() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("systemUptime"))
	return rv
}


// The current thermal state of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/thermalState-swift.property
func (p_ ProcessInfo) ThermalState() ProcessInfoThermalState {
	rv := objc.Send[ProcessInfoThermalState](p_.ID, objc.Sel("thermalState"))
	return rv
}


// Returns the account name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/userName
func (p_ ProcessInfo) UserName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("userName"))
	return rv
}


// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/islowpowermodeenabled
func (p_ ProcessInfo) IsLowPowerModeEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isLowPowerModeEnabled"))
	return rv
}


// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/islowpowermodeenabled
func (p_ ProcessInfo) SetIsLowPowerModeEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsLowPowerModeEnabled:"), value)
}


// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/ismaccatalystapp
func (p_ ProcessInfo) IsMacCatalystApp() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMacCatalystApp"))
	return rv
}


// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/ismaccatalystapp
func (p_ ProcessInfo) SetIsMacCatalystApp(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMacCatalystApp:"), value)
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponmac
func (p_ ProcessInfo) IsiOSAppOnMac() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isiOSAppOnMac"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponmac
func (p_ ProcessInfo) SetIsiOSAppOnMac(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsiOSAppOnMac:"), value)
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponvision
func (p_ ProcessInfo) IsiOSAppOnVision() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isiOSAppOnVision"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponvision
func (p_ ProcessInfo) SetIsiOSAppOnVision(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsiOSAppOnVision:"), value)
}



