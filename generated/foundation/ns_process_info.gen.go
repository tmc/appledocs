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
	BeginActivityWithOptionsReason(options NSActivityOptions, reason string) objc.ID
	EndActivity(activity objectivec.IObject)
	OperatingSystem() uint
	PerformActivityWithOptionsReasonUsingBlock(options NSActivityOptions, reason string, block unsafe.Pointer)
	PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer)
	IOSAppOnMac() bool
	SystemUptime() TimeInterval
	ActiveProcessorCount() int
	SetActiveProcessorCount(value int)
	Arguments() string
	SetArguments(value string)
	AutomaticTerminationSupportEnabled() bool
	SetAutomaticTerminationSupportEnabled(value bool)
	Environment() string
	SetEnvironment(value string)
	FullUserName() string
	SetFullUserName(value string)
	GloballyUniqueString() string
	SetGloballyUniqueString(value string)
	HostName() string
	SetHostName(value string)
	IsLowPowerModeEnabled() bool
	SetIsLowPowerModeEnabled(value bool)
	IsMacCatalystApp() bool
	SetIsMacCatalystApp(value bool)
	IsiOSAppOnMac() bool
	SetIsiOSAppOnMac(value bool)
	IsiOSAppOnVision() bool
	SetIsiOSAppOnVision(value bool)
	OperatingSystemVersion() unsafe.Pointer
	SetOperatingSystemVersion(value unsafe.Pointer)
	OperatingSystemVersionString() string
	SetOperatingSystemVersionString(value string)
	PhysicalMemory() uint64
	SetPhysicalMemory(value uint64)
	ProcessIdentifier() unsafe.Pointer
	SetProcessIdentifier(value unsafe.Pointer)
	ProcessName() string
	SetProcessName(value string)
	ProcessorCount() int
	SetProcessorCount(value int)
	ThermalState() unsafe.Pointer
	SetThermalState(value unsafe.Pointer)
	UserName() string
	SetUserName(value string)
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



// Begin an activity using the given options and reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/beginActivity(options:reason:)
func (p_ ProcessInfo) BeginActivityWithOptionsReason(options NSActivityOptions, reason string) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("beginActivityWithOptions:reason:"), options, objc.String(reason))
	return rv
}


// Ends the given activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/endActivity(_:)
func (p_ ProcessInfo) EndActivity(activity objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("endActivity:"), activity)
}


// Returns a constant to indicate the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystem()
func (p_ ProcessInfo) OperatingSystem() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("operatingSystem"))
	return rv
}


// Synchronously perform an activity defined by a given block using the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performActivity(options:reason:using:)
func (p_ ProcessInfo) PerformActivityWithOptionsReasonUsingBlock(options NSActivityOptions, reason string, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performActivityWithOptions:reason:usingBlock:"), options, objc.String(reason), block)
}


// Performs the specified block asynchronously and notifies you if the process is about to be suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p_ ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performExpiringActivityWithReason:usingBlock:"), objc.String(reason), block)
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnMac
func (p_ ProcessInfo) IOSAppOnMac() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("iOSAppOnMac"))
	return rv
}


// The amount of time the system has been awake since the last time it was restarted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/systemUptime
func (p_ ProcessInfo) SystemUptime() TimeInterval {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("systemUptime"))
	return rv
}


// The number of active processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/activeprocessorcount
func (p_ ProcessInfo) ActiveProcessorCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("activeProcessorCount"))
	return rv
}


// The number of active processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/activeprocessorcount
func (p_ ProcessInfo) SetActiveProcessorCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActiveProcessorCount:"), value)
}


// Array of strings with the command-line arguments for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/arguments
func (p_ ProcessInfo) Arguments() string {
	rv := objc.Send[string](p_.ID, objc.Sel("arguments"))
	return rv
}


// Array of strings with the command-line arguments for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/arguments
func (p_ ProcessInfo) SetArguments(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArguments:"), objc.String(value))
}


// A Boolean value indicating whether the app supports automatic termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/automaticterminationsupportenabled
func (p_ ProcessInfo) AutomaticTerminationSupportEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("automaticTerminationSupportEnabled"))
	return rv
}


// A Boolean value indicating whether the app supports automatic termination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/automaticterminationsupportenabled
func (p_ ProcessInfo) SetAutomaticTerminationSupportEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutomaticTerminationSupportEnabled:"), value)
}


// The variable names (keys) and their values in the environment from which the process was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/environment
func (p_ ProcessInfo) Environment() string {
	rv := objc.Send[string](p_.ID, objc.Sel("environment"))
	return rv
}


// The variable names (keys) and their values in the environment from which the process was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/environment
func (p_ ProcessInfo) SetEnvironment(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnvironment:"), objc.String(value))
}


// Returns the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/fullusername
func (p_ ProcessInfo) FullUserName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("fullUserName"))
	return rv
}


// Returns the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/fullusername
func (p_ ProcessInfo) SetFullUserName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFullUserName:"), objc.String(value))
}


// Global unique identifier for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/globallyuniquestring
func (p_ ProcessInfo) GloballyUniqueString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("globallyUniqueString"))
	return rv
}


// Global unique identifier for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/globallyuniquestring
func (p_ ProcessInfo) SetGloballyUniqueString(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGloballyUniqueString:"), objc.String(value))
}


// The name of the host computer on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/hostname
func (p_ ProcessInfo) HostName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("hostName"))
	return rv
}


// The name of the host computer on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/hostname
func (p_ ProcessInfo) SetHostName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHostName:"), objc.String(value))
}


// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/islowpowermodeenabled
func (p_ ProcessInfo) IsLowPowerModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isLowPowerModeEnabled"))
	return rv
}


// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/islowpowermodeenabled
func (p_ ProcessInfo) SetIsLowPowerModeEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsLowPowerModeEnabled:"), value)
}


// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/ismaccatalystapp
func (p_ ProcessInfo) IsMacCatalystApp() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMacCatalystApp"))
	return rv
}


// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/ismaccatalystapp
func (p_ ProcessInfo) SetIsMacCatalystApp(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMacCatalystApp:"), value)
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponmac
func (p_ ProcessInfo) IsiOSAppOnMac() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isiOSAppOnMac"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponmac
func (p_ ProcessInfo) SetIsiOSAppOnMac(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsiOSAppOnMac:"), value)
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponvision
func (p_ ProcessInfo) IsiOSAppOnVision() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isiOSAppOnVision"))
	return rv
}


// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/isiosapponvision
func (p_ ProcessInfo) SetIsiOSAppOnVision(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsiOSAppOnVision:"), value)
}


// The version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversion
func (p_ ProcessInfo) OperatingSystemVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}


// The version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversion
func (p_ ProcessInfo) SetOperatingSystemVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOperatingSystemVersion:"), value)
}


// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversionstring
func (p_ ProcessInfo) OperatingSystemVersionString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("operatingSystemVersionString"))
	return rv
}


// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversionstring
func (p_ ProcessInfo) SetOperatingSystemVersionString(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOperatingSystemVersionString:"), objc.String(value))
}


// The amount of physical memory on the computer in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/physicalmemory
func (p_ ProcessInfo) PhysicalMemory() uint64 {
	rv := objc.Send[uint64](p_.ID, objc.Sel("physicalMemory"))
	return rv
}


// The amount of physical memory on the computer in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/physicalmemory
func (p_ ProcessInfo) SetPhysicalMemory(value uint64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPhysicalMemory:"), value)
}


// The identifier of the process (often called process ID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processidentifier
func (p_ ProcessInfo) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("processIdentifier"))
	return rv
}


// The identifier of the process (often called process ID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processidentifier
func (p_ ProcessInfo) SetProcessIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessIdentifier:"), value)
}


// The name of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processname
func (p_ ProcessInfo) ProcessName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("processName"))
	return rv
}


// The name of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processname
func (p_ ProcessInfo) SetProcessName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessName:"), objc.String(value))
}


// The number of processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processorcount
func (p_ ProcessInfo) ProcessorCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("processorCount"))
	return rv
}


// The number of processing cores available on the computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processorcount
func (p_ ProcessInfo) SetProcessorCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessorCount:"), value)
}


// The current thermal state of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/thermalstate-swift.property
func (p_ ProcessInfo) ThermalState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("thermalState"))
	return rv
}


// The current thermal state of the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/thermalstate-swift.property
func (p_ ProcessInfo) SetThermalState(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThermalState:"), value)
}


// Returns the account name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/username
func (p_ ProcessInfo) UserName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("userName"))
	return rv
}


// Returns the account name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/username
func (p_ ProcessInfo) SetUserName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), objc.String(value))
}



