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
	ActiveProcessorCount() int
	SetActiveProcessorCount(value int)
	Arguments() IString
	SetArguments(value IString)
	AutomaticTerminationSupportEnabled() bool
	SetAutomaticTerminationSupportEnabled(value bool)
	Environment() IString
	SetEnvironment(value IString)
	FullUserName() IString
	SetFullUserName(value IString)
	GloballyUniqueString() IString
	SetGloballyUniqueString(value IString)
	HostName() IString
	SetHostName(value IString)
	IsLowPowerModeEnabled() bool
	SetIsLowPowerModeEnabled(value bool)
	IsMacCatalystApp() bool
	SetIsMacCatalystApp(value bool)
	IsiOSAppOnMac() bool
	SetIsiOSAppOnMac(value bool)
	IsiOSAppOnVision() bool
	SetIsiOSAppOnVision(value bool)
	OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */
	SetOperatingSystemVersion(value objc.IObject /* cross-framework: OperatingSystemVersion */)
	OperatingSystemVersionString() IString
	SetOperatingSystemVersionString(value IString)
	PhysicalMemory() uint64
	SetPhysicalMemory(value uint64)
	ProcessIdentifier() unsafe.Pointer
	SetProcessIdentifier(value unsafe.Pointer)
	ProcessName() IString
	SetProcessName(value IString)
	ProcessorCount() int
	SetProcessorCount(value int)
	SystemUptime() float64
	SetSystemUptime(value float64)
	ThermalState() unsafe.Pointer
	SetThermalState(value unsafe.Pointer)
	UserName() IString
	SetUserName(value IString)
	// methods:
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
func (p_ ProcessInfo) Arguments() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("arguments"))
	return rv
}


// Array of strings with the command-line arguments for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/arguments
func (p_ ProcessInfo) SetArguments(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setArguments:"), value)
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
func (p_ ProcessInfo) Environment() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("environment"))
	return rv
}


// The variable names (keys) and their values in the environment from which the process was launched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/environment
func (p_ ProcessInfo) SetEnvironment(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnvironment:"), value)
}


// Returns the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/fullusername
func (p_ ProcessInfo) FullUserName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("fullUserName"))
	return rv
}


// Returns the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/fullusername
func (p_ ProcessInfo) SetFullUserName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFullUserName:"), value)
}


// Global unique identifier for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/globallyuniquestring
func (p_ ProcessInfo) GloballyUniqueString() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("globallyUniqueString"))
	return rv
}


// Global unique identifier for the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/globallyuniquestring
func (p_ ProcessInfo) SetGloballyUniqueString(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGloballyUniqueString:"), value)
}


// The name of the host computer on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/hostname
func (p_ ProcessInfo) HostName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("hostName"))
	return rv
}


// The name of the host computer on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/hostname
func (p_ ProcessInfo) SetHostName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHostName:"), value)
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
func (p_ ProcessInfo) OperatingSystemVersion() objc.IObject /* cross-framework: OperatingSystemVersion */ {
	rv := objc.Send[OperatingSystemVersion](p_.ID, objc.Sel("operatingSystemVersion"))
	return rv
}


// The version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversion
func (p_ ProcessInfo) SetOperatingSystemVersion(value objc.IObject /* cross-framework: OperatingSystemVersion */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOperatingSystemVersion:"), value)
}


// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversionstring
func (p_ ProcessInfo) OperatingSystemVersionString() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("operatingSystemVersionString"))
	return rv
}


// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/operatingsystemversionstring
func (p_ ProcessInfo) SetOperatingSystemVersionString(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOperatingSystemVersionString:"), value)
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
func (p_ ProcessInfo) ProcessName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("processName"))
	return rv
}


// The name of the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/processname
func (p_ ProcessInfo) SetProcessName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessName:"), value)
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


// The amount of time the system has been awake since the last time it was restarted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/systemuptime
func (p_ ProcessInfo) SystemUptime() float64 {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("systemUptime"))
	return rv
}


// The amount of time the system has been awake since the last time it was restarted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/systemuptime
func (p_ ProcessInfo) SetSystemUptime(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSystemUptime:"), value)
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
func (p_ ProcessInfo) UserName() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("userName"))
	return rv
}


// Returns the account name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/processinfo/username
func (p_ ProcessInfo) SetUserName(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}



