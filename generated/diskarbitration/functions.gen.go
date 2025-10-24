// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

/* debug [functions.gen.go]: Generating 39 functions for DiskArbitration */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DiskArbitration Functions (39 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_DAApprovalSessionCreate func(AllocatorRef) unsafe.Pointer
	_DAApprovalSessionGetTypeID func() TypeID
	_DAApprovalSessionScheduleWithRunLoop func(unsafe.Pointer, RunLoopRef, StringRef)
	_DAApprovalSessionUnscheduleFromRunLoop func(unsafe.Pointer, RunLoopRef, StringRef)
	_DADiskClaim func(DADiskRef, DADiskClaimOptions, DADiskClaimReleaseCallback, unsafe.Pointer, DADiskClaimCallback, unsafe.Pointer)
	_DADiskCopyDescription func(DADiskRef) DictionaryRef
	_DADiskCopyIOMedia func(DADiskRef) unsafe.Pointer
	_DADiskCopyWholeDisk func(DADiskRef) DADiskRef
	_DADiskCreateFromBSDName func(AllocatorRef, DASessionRef, unsafe.Pointer) DADiskRef
	_DADiskCreateFromIOMedia func(AllocatorRef, DASessionRef, unsafe.Pointer) DADiskRef
	_DADiskCreateFromVolumePath func(AllocatorRef, DASessionRef, URLRef) DADiskRef
	_DADiskEject func(DADiskRef, DADiskEjectOptions, DADiskEjectCallback, unsafe.Pointer)
	_DADiskGetBSDName func(DADiskRef) unsafe.Pointer
	_DADiskGetOptions func(DADiskRef) DADiskOptions
	_DADiskGetTypeID func() TypeID
	_DADiskIsClaimed func(DADiskRef) unsafe.Pointer
	_DADiskMount func(DADiskRef, URLRef, DADiskMountOptions, DADiskMountCallback, unsafe.Pointer)
	_DADiskMountWithArguments func(DADiskRef, URLRef, DADiskMountOptions, DADiskMountCallback, unsafe.Pointer, StringRef)
	_DADiskRename func(DADiskRef, StringRef, DADiskRenameOptions, DADiskRenameCallback, unsafe.Pointer)
	_DADiskSetOptions func(DADiskRef, DADiskOptions, unsafe.Pointer) DAReturn
	_DADiskUnclaim func(DADiskRef)
	_DADiskUnmount func(DADiskRef, DADiskUnmountOptions, DADiskUnmountCallback, unsafe.Pointer)
	_DADissenterCreate func(AllocatorRef, DAReturn, StringRef) DADissenterRef
	_DADissenterGetStatus func(DADissenterRef) DAReturn
	_DADissenterGetStatusString func(DADissenterRef) StringRef
	_DARegisterDiskAppearedCallback func(DASessionRef, DictionaryRef, DADiskAppearedCallback, unsafe.Pointer)
	_DARegisterDiskDescriptionChangedCallback func(DASessionRef, DictionaryRef, ArrayRef, DADiskDescriptionChangedCallback, unsafe.Pointer)
	_DARegisterDiskDisappearedCallback func(DASessionRef, DictionaryRef, DADiskDisappearedCallback, unsafe.Pointer)
	_DARegisterDiskEjectApprovalCallback func(DASessionRef, DictionaryRef, DADiskEjectApprovalCallback, unsafe.Pointer)
	_DARegisterDiskMountApprovalCallback func(DASessionRef, DictionaryRef, DADiskMountApprovalCallback, unsafe.Pointer)
	_DARegisterDiskPeekCallback func(DASessionRef, DictionaryRef, Index, DADiskPeekCallback, unsafe.Pointer)
	_DARegisterDiskUnmountApprovalCallback func(DASessionRef, DictionaryRef, DADiskUnmountApprovalCallback, unsafe.Pointer)
	_DASessionCreate func(AllocatorRef) DASessionRef
	_DASessionGetTypeID func() TypeID
	_DASessionScheduleWithRunLoop func(DASessionRef, RunLoopRef, StringRef)
	_DASessionSetDispatchQueue func(DASessionRef, unsafe.Pointer)
	_DASessionUnscheduleFromRunLoop func(DASessionRef, RunLoopRef, StringRef)
	_DAUnregisterApprovalCallback func(DASessionRef, unsafe.Pointer, unsafe.Pointer)
	_DAUnregisterCallback func(DASessionRef, unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_DAApprovalSessionCreate, lib, "DAApprovalSessionCreate")
	tryRegister(&_DAApprovalSessionGetTypeID, lib, "DAApprovalSessionGetTypeID")
	tryRegister(&_DAApprovalSessionScheduleWithRunLoop, lib, "DAApprovalSessionScheduleWithRunLoop")
	tryRegister(&_DAApprovalSessionUnscheduleFromRunLoop, lib, "DAApprovalSessionUnscheduleFromRunLoop")
	tryRegister(&_DADiskClaim, lib, "DADiskClaim")
	tryRegister(&_DADiskCopyDescription, lib, "DADiskCopyDescription")
	tryRegister(&_DADiskCopyIOMedia, lib, "DADiskCopyIOMedia")
	tryRegister(&_DADiskCopyWholeDisk, lib, "DADiskCopyWholeDisk")
	tryRegister(&_DADiskCreateFromBSDName, lib, "DADiskCreateFromBSDName")
	tryRegister(&_DADiskCreateFromIOMedia, lib, "DADiskCreateFromIOMedia")
	tryRegister(&_DADiskCreateFromVolumePath, lib, "DADiskCreateFromVolumePath")
	tryRegister(&_DADiskEject, lib, "DADiskEject")
	tryRegister(&_DADiskGetBSDName, lib, "DADiskGetBSDName")
	tryRegister(&_DADiskGetOptions, lib, "DADiskGetOptions")
	tryRegister(&_DADiskGetTypeID, lib, "DADiskGetTypeID")
	tryRegister(&_DADiskIsClaimed, lib, "DADiskIsClaimed")
	tryRegister(&_DADiskMount, lib, "DADiskMount")
	tryRegister(&_DADiskMountWithArguments, lib, "DADiskMountWithArguments")
	tryRegister(&_DADiskRename, lib, "DADiskRename")
	tryRegister(&_DADiskSetOptions, lib, "DADiskSetOptions")
	tryRegister(&_DADiskUnclaim, lib, "DADiskUnclaim")
	tryRegister(&_DADiskUnmount, lib, "DADiskUnmount")
	tryRegister(&_DADissenterCreate, lib, "DADissenterCreate")
	tryRegister(&_DADissenterGetStatus, lib, "DADissenterGetStatus")
	tryRegister(&_DADissenterGetStatusString, lib, "DADissenterGetStatusString")
	tryRegister(&_DARegisterDiskAppearedCallback, lib, "DARegisterDiskAppearedCallback")
	tryRegister(&_DARegisterDiskDescriptionChangedCallback, lib, "DARegisterDiskDescriptionChangedCallback")
	tryRegister(&_DARegisterDiskDisappearedCallback, lib, "DARegisterDiskDisappearedCallback")
	tryRegister(&_DARegisterDiskEjectApprovalCallback, lib, "DARegisterDiskEjectApprovalCallback")
	tryRegister(&_DARegisterDiskMountApprovalCallback, lib, "DARegisterDiskMountApprovalCallback")
	tryRegister(&_DARegisterDiskPeekCallback, lib, "DARegisterDiskPeekCallback")
	tryRegister(&_DARegisterDiskUnmountApprovalCallback, lib, "DARegisterDiskUnmountApprovalCallback")
	tryRegister(&_DASessionCreate, lib, "DASessionCreate")
	tryRegister(&_DASessionGetTypeID, lib, "DASessionGetTypeID")
	tryRegister(&_DASessionScheduleWithRunLoop, lib, "DASessionScheduleWithRunLoop")
	tryRegister(&_DASessionSetDispatchQueue, lib, "DASessionSetDispatchQueue")
	tryRegister(&_DASessionUnscheduleFromRunLoop, lib, "DASessionUnscheduleFromRunLoop")
	tryRegister(&_DAUnregisterApprovalCallback, lib, "DAUnregisterApprovalCallback")
	tryRegister(&_DAUnregisterCallback, lib, "DAUnregisterCallback")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// DAApprovalSessionCreate is a DiskArbitration function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionCreate
func DAApprovalSessionCreate(allocator AllocatorRef) unsafe.Pointer {
	return _DAApprovalSessionCreate(allocator)
}/* debug [functions.gen.go/function]: DAApprovalSessionCreate */

// DAApprovalSessionGetTypeID is a DiskArbitration function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionGetTypeID
func DAApprovalSessionGetTypeID() TypeID {
	return _DAApprovalSessionGetTypeID()
}/* debug [functions.gen.go/function]: DAApprovalSessionGetTypeID */

// DAApprovalSessionScheduleWithRunLoop is a DiskArbitration function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionScheduleWithRunLoop
func DAApprovalSessionScheduleWithRunLoop(session unsafe.Pointer, runLoop RunLoopRef, runLoopMode StringRef) {
	_DAApprovalSessionScheduleWithRunLoop(session, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: DAApprovalSessionScheduleWithRunLoop */

// DAApprovalSessionUnscheduleFromRunLoop is a DiskArbitration function.
//
// Added in macOS 10.4.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionUnscheduleFromRunLoop
func DAApprovalSessionUnscheduleFromRunLoop(session unsafe.Pointer, runLoop RunLoopRef, runLoopMode StringRef) {
	_DAApprovalSessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: DAApprovalSessionUnscheduleFromRunLoop */

// Claims the specified disk object for exclusive use.
//
// Added in macOS 10.4.
// Claims the specified disk object for exclusive use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaim(_:_:_:_:_:_:)
func DADiskClaim(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer) {
	_DADiskClaim(disk, options, release, releaseContext, callback, callbackContext)
}/* debug [functions.gen.go/function]: DADiskClaim */

// Obtains the Disk Arbitration description of the specified disk.
//
// Added in macOS 10.4.
// Obtains the Disk Arbitration description of the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyDescription(_:)
func DADiskCopyDescription(disk DADiskRef) DictionaryRef {
	return _DADiskCopyDescription(disk)
}/* debug [functions.gen.go/function]: DADiskCopyDescription */

// Obtains the I/O Kit media object for the specified disk.
//
// Added in macOS 10.4.
// Obtains the I/O Kit media object for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyIOMedia(_:)
func DADiskCopyIOMedia(disk DADiskRef) unsafe.Pointer {
	return _DADiskCopyIOMedia(disk)
}/* debug [functions.gen.go/function]: DADiskCopyIOMedia */

// Obtain the associated whole disk object for the specified disk.
//
// Added in macOS 10.4.
// Obtain the associated whole disk object for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyWholeDisk(_:)
func DADiskCopyWholeDisk(disk DADiskRef) DADiskRef {
	return _DADiskCopyWholeDisk(disk)
}/* debug [functions.gen.go/function]: DADiskCopyWholeDisk */

// Creates a new disk object.
//
// Added in macOS 10.4.
// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromBSDName(_:_:_:)
func DADiskCreateFromBSDName(allocator AllocatorRef, session DASessionRef, name unsafe.Pointer) DADiskRef {
	return _DADiskCreateFromBSDName(allocator, session, name)
}/* debug [functions.gen.go/function]: DADiskCreateFromBSDName */

// Creates a new disk object.
//
// Added in macOS 10.4.
// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromIOMedia(_:_:_:)
func DADiskCreateFromIOMedia(allocator AllocatorRef, session DASessionRef, media unsafe.Pointer) DADiskRef {
	return _DADiskCreateFromIOMedia(allocator, session, media)
}/* debug [functions.gen.go/function]: DADiskCreateFromIOMedia */

// Creates a new disk object.
//
// Added in macOS 10.7.
// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromVolumePath(_:_:_:)
func DADiskCreateFromVolumePath(allocator AllocatorRef, session DASessionRef, path URLRef) DADiskRef {
	return _DADiskCreateFromVolumePath(allocator, session, path)
}/* debug [functions.gen.go/function]: DADiskCreateFromVolumePath */

// Ejects the specified disk object.
//
// Added in macOS 10.4.
// Ejects the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEject(_:_:_:_:)
func DADiskEject(disk DADiskRef, options DADiskEjectOptions, callback DADiskEjectCallback, context unsafe.Pointer) {
	_DADiskEject(disk, options, callback, context)
}/* debug [functions.gen.go/function]: DADiskEject */

// Obtains the BSD device name for the specified disk.
//
// Added in macOS 10.4.
// Obtains the BSD device name for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetBSDName(_:)
func DADiskGetBSDName(disk DADiskRef) unsafe.Pointer {
	return _DADiskGetBSDName(disk)
}/* debug [functions.gen.go/function]: DADiskGetBSDName */

// Obtains the options for the specified disk.
//
// Added in macOS 10.4.
// Obtains the options for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetOptions(_:)
func DADiskGetOptions(disk DADiskRef) DADiskOptions {
	return _DADiskGetOptions(disk)
}/* debug [functions.gen.go/function]: DADiskGetOptions */

// Returns the type identifier of all DADisk instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all DADisk instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetTypeID()
func DADiskGetTypeID() TypeID {
	return _DADiskGetTypeID()
}/* debug [functions.gen.go/function]: DADiskGetTypeID */

// Reports whether or not the disk is claimed.
//
// Added in macOS 10.4.
// Reports whether or not the disk is claimed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskIsClaimed(_:)
func DADiskIsClaimed(disk DADiskRef) unsafe.Pointer {
	return _DADiskIsClaimed(disk)
}/* debug [functions.gen.go/function]: DADiskIsClaimed */

// Mounts the volume at the specified disk object.
//
// Added in macOS 10.4.
// Mounts the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMount(_:_:_:_:_:)
func DADiskMount(disk DADiskRef, path URLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer) {
	_DADiskMount(disk, path, options, callback, context)
}/* debug [functions.gen.go/function]: DADiskMount */

// Mounts the volume at the specified disk object, with the specified mount options.
//
// Added in macOS 10.4.
// Mounts the volume at the specified disk object, with the specified mount options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountWithArguments(_:_:_:_:_:_:)
func DADiskMountWithArguments(disk DADiskRef, path URLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments StringRef) {
	_DADiskMountWithArguments(disk, path, options, callback, context, arguments)
}/* debug [functions.gen.go/function]: DADiskMountWithArguments */

// Renames the volume at the specified disk object.
//
// Added in macOS 10.4.
// Renames the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRename(_:_:_:_:_:)
func DADiskRename(disk DADiskRef, name StringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer) {
	_DADiskRename(disk, name, options, callback, context)
}/* debug [functions.gen.go/function]: DADiskRename */

// Sets the options for the specified disk.
//
// Added in macOS 10.4.
// Sets the options for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskSetOptions(_:_:_:)
func DADiskSetOptions(disk DADiskRef, options DADiskOptions, value unsafe.Pointer) DAReturn {
	return _DADiskSetOptions(disk, options, value)
}/* debug [functions.gen.go/function]: DADiskSetOptions */

// Unclaims the specified disk object.
//
// Added in macOS 10.4.
// Unclaims the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnclaim(_:)
func DADiskUnclaim(disk DADiskRef) {
	_DADiskUnclaim(disk)
}/* debug [functions.gen.go/function]: DADiskUnclaim */

// Unmounts the volume at the specified disk object.
//
// Added in macOS 10.4.
// Unmounts the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmount(_:_:_:_:)
func DADiskUnmount(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer) {
	_DADiskUnmount(disk, options, callback, context)
}/* debug [functions.gen.go/function]: DADiskUnmount */

// Creates a new dissenter object.
//
// Added in macOS 10.4.
// Creates a new dissenter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterCreate(_:_:_:)
func DADissenterCreate(allocator AllocatorRef, status DAReturn, string_ StringRef) DADissenterRef {
	return _DADissenterCreate(allocator, status, string_)
}/* debug [functions.gen.go/function]: DADissenterCreate */

// Obtains the return code.
//
// Added in macOS 10.4.
// Obtains the return code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatus(_:)
func DADissenterGetStatus(dissenter DADissenterRef) DAReturn {
	return _DADissenterGetStatus(dissenter)
}/* debug [functions.gen.go/function]: DADissenterGetStatus */

// Obtains the return code string.
//
// Added in macOS 10.4.
// Obtains the return code string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatusString(_:)
func DADissenterGetStatusString(dissenter DADissenterRef) StringRef {
	return _DADissenterGetStatusString(dissenter)
}/* debug [functions.gen.go/function]: DADissenterGetStatusString */

// Registers a callback function to be called whenever a disk has appeared.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a disk has appeared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskAppearedCallback(_:_:_:_:)
func DARegisterDiskAppearedCallback(session DASessionRef, match DictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer) {
	_DARegisterDiskAppearedCallback(session, match, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskAppearedCallback */

// Registers a callback function to be called whenever a disk description has changed.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a disk description has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDescriptionChangedCallback(_:_:_:_:_:)
func DARegisterDiskDescriptionChangedCallback(session DASessionRef, match DictionaryRef, watch ArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer) {
	_DARegisterDiskDescriptionChangedCallback(session, match, watch, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskDescriptionChangedCallback */

// Registers a callback function to be called whenever a disk has disappeared.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a disk has disappeared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDisappearedCallback(_:_:_:_:)
func DARegisterDiskDisappearedCallback(session DASessionRef, match DictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer) {
	_DARegisterDiskDisappearedCallback(session, match, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskDisappearedCallback */

// Registers a callback function to be called whenever a volume is to be ejected.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a volume is to be ejected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskEjectApprovalCallback(_:_:_:_:)
func DARegisterDiskEjectApprovalCallback(session DASessionRef, match DictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer) {
	_DARegisterDiskEjectApprovalCallback(session, match, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskEjectApprovalCallback */

// Registers a callback function to be called whenever a volume is to be mounted.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a volume is to be mounted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskMountApprovalCallback(_:_:_:_:)
func DARegisterDiskMountApprovalCallback(session DASessionRef, match DictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer) {
	_DARegisterDiskMountApprovalCallback(session, match, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskMountApprovalCallback */

// Registers a callback function to be called whenever a disk has been probed.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a disk has been probed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskPeekCallback(_:_:_:_:_:)
func DARegisterDiskPeekCallback(session DASessionRef, match DictionaryRef, order Index, callback DADiskPeekCallback, context unsafe.Pointer) {
	_DARegisterDiskPeekCallback(session, match, order, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskPeekCallback */

// Registers a callback function to be called whenever a volume is to be unmounted.
//
// Added in macOS 10.4.
// Registers a callback function to be called whenever a volume is to be unmounted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskUnmountApprovalCallback(_:_:_:_:)
func DARegisterDiskUnmountApprovalCallback(session DASessionRef, match DictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer) {
	_DARegisterDiskUnmountApprovalCallback(session, match, callback, context)
}/* debug [functions.gen.go/function]: DARegisterDiskUnmountApprovalCallback */

// Creates a new session.
//
// Added in macOS 10.4.
// Creates a new session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionCreate(_:)
func DASessionCreate(allocator AllocatorRef) DASessionRef {
	return _DASessionCreate(allocator)
}/* debug [functions.gen.go/function]: DASessionCreate */

// Returns the type identifier of all DASession instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all DASession instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionGetTypeID()
func DASessionGetTypeID() TypeID {
	return _DASessionGetTypeID()
}/* debug [functions.gen.go/function]: DASessionGetTypeID */

// Schedules the session on a run loop.
//
// Added in macOS 10.4.
// Schedules the session on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionScheduleWithRunLoop(_:_:_:)
func DASessionScheduleWithRunLoop(session DASessionRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_DASessionScheduleWithRunLoop(session, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: DASessionScheduleWithRunLoop */

// Schedules the session on a dispatch queue.
//
// Added in macOS 10.7.
// Schedules the session on a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionSetDispatchQueue(_:_:)
func DASessionSetDispatchQueue(session DASessionRef, queue unsafe.Pointer) {
	_DASessionSetDispatchQueue(session, queue)
}/* debug [functions.gen.go/function]: DASessionSetDispatchQueue */

// Unschedules the session from a run loop.
//
// Added in macOS 10.4.
// Unschedules the session from a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionUnscheduleFromRunLoop(_:_:_:)
func DASessionUnscheduleFromRunLoop(session DASessionRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_DASessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: DASessionUnscheduleFromRunLoop */

// Unregisters a registered callback function.
//
// Added in macOS 10.4.
// Unregisters a registered callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterApprovalCallback
func DAUnregisterApprovalCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	_DAUnregisterApprovalCallback(session, callback, context)
}/* debug [functions.gen.go/function]: DAUnregisterApprovalCallback */

// Unregisters a registered callback function.
//
// Added in macOS 10.4.
// Unregisters a registered callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterCallback(_:_:_:)
func DAUnregisterCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	_DAUnregisterCallback(session, callback, context)
}/* debug [functions.gen.go/function]: DAUnregisterCallback */




