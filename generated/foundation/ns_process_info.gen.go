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
	BeginActivityWithOptionsReason(options unsafe.Pointer, reason string) objc.ID
	DisableAutomaticTermination(reason string)
	EnableAutomaticTermination(reason string)
	EnableSuddenTermination()
	EndActivity(activity objc.ID)
	IsDeviceCertifiedFor(performanceTier unsafe.Pointer) bool
	IsOperatingSystemAtLeastVersion(version unsafe.Pointer) bool
	OperatingSystemName() string
	PerformActivityWithOptionsReasonUsingBlock(options unsafe.Pointer, reason string, block unsafe.Pointer)
	PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer)
}

// A collection of information about the current process.
//
// Each process has a single, shared object known as a that can return information such as arguments, environment variables, host name, and process name. The class method returns the shared agent for the current process. For example, the following line returns the object, which then provides the name of the current process: The class also includes the property, which returns an structure identifying the operating system version on which the process is executing. objects attempt to interpret environment variables and command-line arguments in the user’s default C string encoding if they can’t convert to Unicode as UTF-8 strings. If neither the Unicode nor C string conversion works, the object ignores these values.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processInfo
func (pc _ProcessInfoClass) ProcessInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("processInfo"))
	return rv
}
// Begin an activity using the given options and reason.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/beginActivity(options:reason:)
func (p_ ProcessInfo) BeginActivityWithOptionsReason(options unsafe.Pointer, reason string) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("beginActivityWithOptions:reason:"), options, objc.String(reason))
	return rv
}

// Disables automatic termination for the application.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/disableAutomaticTermination(_:)
func (p_ ProcessInfo) DisableAutomaticTermination(reason string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("disableAutomaticTermination:"), objc.String(reason))
}

// Enables automatic termination for the application.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableAutomaticTermination(_:)
func (p_ ProcessInfo) EnableAutomaticTermination(reason string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("enableAutomaticTermination:"), objc.String(reason))
}

// Enables the application for quick killing using sudden termination.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableSuddenTermination()
func (p_ ProcessInfo) EnableSuddenTermination() {
	objc.Send[objc.ID](p_.ID, objc.Sel("enableSuddenTermination"))
}

// Ends the given activity.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/endActivity(_:)
func (p_ ProcessInfo) EndActivity(activity objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("endActivity:"), activity)
}

// Indicates whether the device supports the requested performance tier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isDeviceCertified(for:)
func (p_ ProcessInfo) IsDeviceCertifiedFor(performanceTier unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDeviceCertifiedFor:"), performanceTier)
	return rv
}

// Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isOperatingSystemAtLeast(_:)
func (p_ ProcessInfo) IsOperatingSystemAtLeastVersion(version unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOperatingSystemAtLeastVersion:"), version)
	return rv
}

// Returns a string containing the name of the operating system on which the process is executing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemName()
func (p_ ProcessInfo) OperatingSystemName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("operatingSystemName"))
	return rv
}

// Synchronously perform an activity defined by a given block using the given options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performActivity(options:reason:using:)
func (p_ ProcessInfo) PerformActivityWithOptionsReasonUsingBlock(options unsafe.Pointer, reason string, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performActivityWithOptions:reason:usingBlock:"), options, objc.String(reason), block)
}

// Performs the specified block asynchronously and notifies you if the process is about to be suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/performExpiringActivity(withReason:using:)
func (p_ ProcessInfo) PerformExpiringActivityWithReasonUsingBlock(reason string, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performExpiringActivityWithReason:usingBlock:"), objc.String(reason), block)
}

// The number of active processing cores available on the computer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/activeProcessorCount
func (p_ ProcessInfo) ActiveProcessorCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("activeProcessorCount"))
	return rv
}

// Array of strings with the command-line arguments for the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/arguments
func (p_ ProcessInfo) Arguments() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("arguments"))
	return rv
}

// The variable names (keys) and their values in the environment from which the process was launched.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/environment
func (p_ ProcessInfo) Environment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("environment"))
	return rv
}

// Returns the full name of the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/fullUserName
func (p_ ProcessInfo) FullUserName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("fullUserName"))
	return rv
}

// Global unique identifier for the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/globallyUniqueString
func (p_ ProcessInfo) GloballyUniqueString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("globallyUniqueString"))
	return rv
}

// A Boolean value that indicates the current state of Low Power Mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isLowPowerModeEnabled
func (p_ ProcessInfo) LowPowerModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("lowPowerModeEnabled"))
	return rv
}

// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isMacCatalystApp
func (p_ ProcessInfo) MacCatalystApp() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("macCatalystApp"))
	return rv
}

// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnMac
func (p_ ProcessInfo) IOSAppOnMac() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("iOSAppOnMac"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnVision
func (p_ ProcessInfo) IOSAppOnVision() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("iOSAppOnVision"))
	return rv
}

// A string containing the version of the operating system on which the process is executing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemVersionString
func (p_ ProcessInfo) OperatingSystemVersionString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("operatingSystemVersionString"))
	return rv
}

// The amount of physical memory on the computer in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/physicalMemory
func (p_ ProcessInfo) PhysicalMemory() uint64 {
	rv := objc.Send[uint64](p_.ID, objc.Sel("physicalMemory"))
	return rv
}

// The identifier of the process (often called process ID).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processIdentifier
func (p_ ProcessInfo) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("processIdentifier"))
	return rv
}

// Returns the process information agent for the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processInfo
func (p_ ProcessInfo) ProcessInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("processInfo"))
	return rv
}

// The name of the process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processName
func (p_ ProcessInfo) ProcessName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("processName"))
	return rv
}


// SetProcessName sets the value of the processName property.
// The name of the process.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/processName
func (p_ ProcessInfo) SetProcessName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProcessName:"), objc.String(value))
}
// The amount of time the system has been awake since the last time it was restarted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ProcessInfo/systemUptime
func (p_ ProcessInfo) SystemUptime() TimeInterval {
	rv := objc.Send[TimeInterval](p_.ID, objc.Sel("systemUptime"))
	return rv
}



