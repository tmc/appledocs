// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// DiskArbitration Functions (39 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_DAApprovalSessionCreate func(unsafe.Pointer) unsafe.Pointer
	_DAApprovalSessionGetTypeID func() unsafe.Pointer
	_DAApprovalSessionScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DAApprovalSessionUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskClaim func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_DADiskCopyIOMedia func(unsafe.Pointer) unsafe.Pointer
	_DADiskCopyWholeDisk func(unsafe.Pointer) unsafe.Pointer
	_DADiskCreateFromBSDName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskCreateFromIOMedia func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskCreateFromVolumePath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskEject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskGetBSDName func(unsafe.Pointer) unsafe.Pointer
	_DADiskGetOptions func(unsafe.Pointer) unsafe.Pointer
	_DADiskGetTypeID func() unsafe.Pointer
	_DADiskIsClaimed func(unsafe.Pointer) unsafe.Pointer
	_DADiskMount func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskMountWithArguments func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskRename func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskSetOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADiskUnclaim func(unsafe.Pointer) unsafe.Pointer
	_DADiskUnmount func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADissenterCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DADissenterGetStatus func(unsafe.Pointer) unsafe.Pointer
	_DADissenterGetStatusString func(unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskAppearedCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskDescriptionChangedCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskDisappearedCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskEjectApprovalCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskMountApprovalCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskPeekCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DARegisterDiskUnmountApprovalCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DASessionCreate func(unsafe.Pointer) unsafe.Pointer
	_DASessionGetTypeID func() unsafe.Pointer
	_DASessionScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DASessionSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DASessionUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DAUnregisterApprovalCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_DAUnregisterCallback func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
func DAApprovalSessionCreate(allocator unsafe.Pointer) unsafe.Pointer {
	return _DAApprovalSessionCreate(allocator)
	}


// DAApprovalSessionGetTypeID is a DiskArbitration function.
//
// Added in macOS 10.4.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionGetTypeID
func DAApprovalSessionGetTypeID() unsafe.Pointer {
	return _DAApprovalSessionGetTypeID()
	}


// DAApprovalSessionScheduleWithRunLoop is a DiskArbitration function.
//
// Added in macOS 10.4.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionScheduleWithRunLoop
func DAApprovalSessionScheduleWithRunLoop(session unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_DAApprovalSessionScheduleWithRunLoop(session, runLoop, runLoopMode)
	}


// DAApprovalSessionUnscheduleFromRunLoop is a DiskArbitration function.
//
// Added in macOS 10.4.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionUnscheduleFromRunLoop
func DAApprovalSessionUnscheduleFromRunLoop(session unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_DAApprovalSessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
	}


// Claims the specified disk object for exclusive use.
//
// Added in macOS 10.4.

// Claims the specified disk object for exclusive use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaim(_:_:_:_:_:_:)
func DADiskClaim(disk unsafe.Pointer, options unsafe.Pointer, release unsafe.Pointer, releaseContext unsafe.Pointer, callback unsafe.Pointer, callbackContext unsafe.Pointer) {
	_DADiskClaim(disk, options, release, releaseContext, callback, callbackContext)
	}


// Obtains the Disk Arbitration description of the specified disk.
//
// Added in macOS 10.4.

// Obtains the Disk Arbitration description of the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyDescription(_:)
func DADiskCopyDescription(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskCopyDescription(disk)
	}


// Obtains the I/O Kit media object for the specified disk.
//
// Added in macOS 10.4.

// Obtains the I/O Kit media object for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyIOMedia(_:)
func DADiskCopyIOMedia(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskCopyIOMedia(disk)
	}


// Obtain the associated whole disk object for the specified disk.
//
// Added in macOS 10.4.

// Obtain the associated whole disk object for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyWholeDisk(_:)
func DADiskCopyWholeDisk(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskCopyWholeDisk(disk)
	}


// Creates a new disk object.
//
// Added in macOS 10.4.

// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromBSDName(_:_:_:)
func DADiskCreateFromBSDName(allocator unsafe.Pointer, session unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _DADiskCreateFromBSDName(allocator, session, name)
	}


// Creates a new disk object.
//
// Added in macOS 10.4.

// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromIOMedia(_:_:_:)
func DADiskCreateFromIOMedia(allocator unsafe.Pointer, session unsafe.Pointer, media unsafe.Pointer) unsafe.Pointer {
	return _DADiskCreateFromIOMedia(allocator, session, media)
	}


// Creates a new disk object.
//
// Added in macOS 10.7.

// Creates a new disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromVolumePath(_:_:_:)
func DADiskCreateFromVolumePath(allocator unsafe.Pointer, session unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _DADiskCreateFromVolumePath(allocator, session, path)
	}


// Ejects the specified disk object.
//
// Added in macOS 10.4.

// Ejects the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEject(_:_:_:_:)
func DADiskEject(disk unsafe.Pointer, options unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DADiskEject(disk, options, callback, context)
	}


// Obtains the BSD device name for the specified disk.
//
// Added in macOS 10.4.

// Obtains the BSD device name for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetBSDName(_:)
func DADiskGetBSDName(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskGetBSDName(disk)
	}


// Obtains the options for the specified disk.
//
// Added in macOS 10.4.

// Obtains the options for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetOptions(_:)
func DADiskGetOptions(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskGetOptions(disk)
	}


// Returns the type identifier of all DADisk instances.
//
// Added in macOS 10.4.

// Returns the type identifier of all DADisk instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskGetTypeID()
func DADiskGetTypeID() unsafe.Pointer {
	return _DADiskGetTypeID()
	}


// Reports whether or not the disk is claimed.
//
// Added in macOS 10.4.

// Reports whether or not the disk is claimed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskIsClaimed(_:)
func DADiskIsClaimed(disk unsafe.Pointer) unsafe.Pointer {
	return _DADiskIsClaimed(disk)
	}


// Mounts the volume at the specified disk object.
//
// Added in macOS 10.4.

// Mounts the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMount(_:_:_:_:_:)
func DADiskMount(disk unsafe.Pointer, path unsafe.Pointer, options unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DADiskMount(disk, path, options, callback, context)
	}


// Mounts the volume at the specified disk object, with the specified mount options.
//
// Added in macOS 10.4.

// Mounts the volume at the specified disk object, with the specified mount options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountWithArguments(_:_:_:_:_:_:)
func DADiskMountWithArguments(disk unsafe.Pointer, path unsafe.Pointer, options unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer, arguments unsafe.Pointer) {
	_DADiskMountWithArguments(disk, path, options, callback, context, arguments)
	}


// Renames the volume at the specified disk object.
//
// Added in macOS 10.4.

// Renames the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRename(_:_:_:_:_:)
func DADiskRename(disk unsafe.Pointer, name unsafe.Pointer, options unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DADiskRename(disk, name, options, callback, context)
	}


// Sets the options for the specified disk.
//
// Added in macOS 10.4.

// Sets the options for the specified disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskSetOptions(_:_:_:)
func DADiskSetOptions(disk unsafe.Pointer, options unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _DADiskSetOptions(disk, options, value)
	}


// Unclaims the specified disk object.
//
// Added in macOS 10.4.

// Unclaims the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnclaim(_:)
func DADiskUnclaim(disk unsafe.Pointer) {
	_DADiskUnclaim(disk)
	}


// Unmounts the volume at the specified disk object.
//
// Added in macOS 10.4.

// Unmounts the volume at the specified disk object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmount(_:_:_:_:)
func DADiskUnmount(disk unsafe.Pointer, options unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DADiskUnmount(disk, options, callback, context)
	}


// Creates a new dissenter object.
//
// Added in macOS 10.4.

// Creates a new dissenter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterCreate(_:_:_:)
func DADissenterCreate(allocator unsafe.Pointer, status unsafe.Pointer, string_ unsafe.Pointer) unsafe.Pointer {
	return _DADissenterCreate(allocator, status, string_)
	}


// Obtains the return code.
//
// Added in macOS 10.4.

// Obtains the return code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatus(_:)
func DADissenterGetStatus(dissenter unsafe.Pointer) unsafe.Pointer {
	return _DADissenterGetStatus(dissenter)
	}


// Obtains the return code string.
//
// Added in macOS 10.4.

// Obtains the return code string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatusString(_:)
func DADissenterGetStatusString(dissenter unsafe.Pointer) unsafe.Pointer {
	return _DADissenterGetStatusString(dissenter)
	}


// Registers a callback function to be called whenever a disk has appeared.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a disk has appeared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskAppearedCallback(_:_:_:_:)
func DARegisterDiskAppearedCallback(session unsafe.Pointer, match unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskAppearedCallback(session, match, callback, context)
	}


// Registers a callback function to be called whenever a disk description has changed.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a disk description has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDescriptionChangedCallback(_:_:_:_:_:)
func DARegisterDiskDescriptionChangedCallback(session unsafe.Pointer, match unsafe.Pointer, watch unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskDescriptionChangedCallback(session, match, watch, callback, context)
	}


// Registers a callback function to be called whenever a disk has disappeared.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a disk has disappeared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDisappearedCallback(_:_:_:_:)
func DARegisterDiskDisappearedCallback(session unsafe.Pointer, match unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskDisappearedCallback(session, match, callback, context)
	}


// Registers a callback function to be called whenever a volume is to be ejected.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a volume is to be ejected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskEjectApprovalCallback(_:_:_:_:)
func DARegisterDiskEjectApprovalCallback(session unsafe.Pointer, match unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskEjectApprovalCallback(session, match, callback, context)
	}


// Registers a callback function to be called whenever a volume is to be mounted.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a volume is to be mounted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskMountApprovalCallback(_:_:_:_:)
func DARegisterDiskMountApprovalCallback(session unsafe.Pointer, match unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskMountApprovalCallback(session, match, callback, context)
	}


// Registers a callback function to be called whenever a disk has been probed.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a disk has been probed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskPeekCallback(_:_:_:_:_:)
func DARegisterDiskPeekCallback(session unsafe.Pointer, match unsafe.Pointer, order unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskPeekCallback(session, match, order, callback, context)
	}


// Registers a callback function to be called whenever a volume is to be unmounted.
//
// Added in macOS 10.4.

// Registers a callback function to be called whenever a volume is to be unmounted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskUnmountApprovalCallback(_:_:_:_:)
func DARegisterDiskUnmountApprovalCallback(session unsafe.Pointer, match unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DARegisterDiskUnmountApprovalCallback(session, match, callback, context)
	}


// Creates a new session.
//
// Added in macOS 10.4.

// Creates a new session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionCreate(_:)
func DASessionCreate(allocator unsafe.Pointer) unsafe.Pointer {
	return _DASessionCreate(allocator)
	}


// Returns the type identifier of all DASession instances.
//
// Added in macOS 10.4.

// Returns the type identifier of all DASession instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionGetTypeID()
func DASessionGetTypeID() unsafe.Pointer {
	return _DASessionGetTypeID()
	}


// Schedules the session on a run loop.
//
// Added in macOS 10.4.

// Schedules the session on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionScheduleWithRunLoop(_:_:_:)
func DASessionScheduleWithRunLoop(session unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_DASessionScheduleWithRunLoop(session, runLoop, runLoopMode)
	}


// Schedules the session on a dispatch queue.
//
// Added in macOS 10.7.

// Schedules the session on a dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionSetDispatchQueue(_:_:)
func DASessionSetDispatchQueue(session unsafe.Pointer, queue unsafe.Pointer) {
	_DASessionSetDispatchQueue(session, queue)
	}


// Unschedules the session from a run loop.
//
// Added in macOS 10.4.

// Unschedules the session from a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASessionUnscheduleFromRunLoop(_:_:_:)
func DASessionUnscheduleFromRunLoop(session unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_DASessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
	}


// Unregisters a registered callback function.
//
// Added in macOS 10.4.

// Unregisters a registered callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterApprovalCallback
func DAUnregisterApprovalCallback(session unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DAUnregisterApprovalCallback(session, callback, context)
	}


// Unregisters a registered callback function.
//
// Added in macOS 10.4.

// Unregisters a registered callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterCallback(_:_:_:)
func DAUnregisterCallback(session unsafe.Pointer, callback unsafe.Pointer, context unsafe.Pointer) {
	_DAUnregisterCallback(session, callback, context)
	}




