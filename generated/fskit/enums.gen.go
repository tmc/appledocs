// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

// Enum types and constants
// FSBlockmapFlags - Flags that describe the behavior of a blockmap operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags
type FSBlockmapFlags uint

const (
	// FSBlockmapFlagsRead - A flag that describes a read operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags/read
	FSBlockmapFlagsRead FSBlockmapFlags = 0
	// FSBlockmapFlagsWrite - A flag that describes a write operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags/write
	FSBlockmapFlagsWrite FSBlockmapFlags = 0
)

// FSCompleteIOFlags - Flags that describe the behavior of an I/O completion operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags
type FSCompleteIOFlags uint

const (
	// FSCompleteIOFlagsAsync - A flag that requests that the file system module flush metadata I/O asynchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags/async
	FSCompleteIOFlagsAsync FSCompleteIOFlags = 0
	// FSCompleteIOFlagsRead - A flag that describes a read operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags/read
	FSCompleteIOFlagsRead FSCompleteIOFlags = 0
	// FSCompleteIOFlagsWrite - A flag that describes a write operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags/write
	FSCompleteIOFlagsWrite FSCompleteIOFlags = 0
)

// FSContainerState - An enumeration of container state values.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerState
type FSContainerState uint

const (
	// FSContainerStateActive - The container is active, and one or more volumes are active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerState/active
	FSContainerStateActive FSContainerState = 0
	// FSContainerStateBlocked - The container is blocked from transitioning from the not-ready state to the ready state by a potentially-recoverable error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerState/blocked
	FSContainerStateBlocked FSContainerState = 0
	// FSContainerStateNotReady - The container isn’t ready.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerState/notReady
	FSContainerStateNotReady FSContainerState = 0
	// FSContainerStateReady - The container is ready, but inactive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerState/ready
	FSContainerStateReady FSContainerState = 0
)

// FSDeactivateOptions - Options that affect the behavior of deactivate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDeactivateOptions
type FSDeactivateOptions uint

// FSErrorCode - A code that indicates a specific FSKit error.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code
type FSErrorCode uint

const (
	// FSErrorInvalidDirectoryCookie - While enumerating a directory, the given cookie didn’t resolve to a valid directory entry.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/invalidDirectoryCookie
	FSErrorInvalidDirectoryCookie FSErrorCode = 0
	// FSErrorModuleLoadFailed - The module failed to load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/moduleLoadFailed
	FSErrorModuleLoadFailed FSErrorCode = 0
	// FSErrorResourceDamaged - The resource is damaged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/resourceDamaged
	FSErrorResourceDamaged FSErrorCode = 0
	// FSErrorResourceUnrecognized - FSKit didn’t recognize the resource, and probing failed to find a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/resourceUnrecognized
	FSErrorResourceUnrecognized FSErrorCode = 0
	// FSErrorResourceUnusable - FSKit recognizes the resource, but the resource isn’t usable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/resourceUnusable
	FSErrorResourceUnusable FSErrorCode = 0
	// FSErrorStatusOperationInProgress - An operation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/statusOperationInProgress
	FSErrorStatusOperationInProgress FSErrorCode = 0
	// FSErrorStatusOperationPaused - An operation is paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/statusOperationPaused
	FSErrorStatusOperationPaused FSErrorCode = 0
)

// FSExtentType - An enumeration of types of extents.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentType
type FSExtentType uint

const (
	// FSExtentTypeData - An extent type to indicate valid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentType/data
	FSExtentTypeData FSExtentType = 0
	// FSExtentTypeZeroFill - An extent type to indicate uninitialized data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentType/zeroFill
	FSExtentTypeZeroFill FSExtentType = 0
)

// FSItemAttribute - A value that indicates a set of item attributes to get or set.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute
type FSItemAttribute uint

const (
	// FSItemAttributeAccessTime - The last-accessed time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/accessTime
	FSItemAttributeAccessTime FSItemAttribute = 0
	// FSItemAttributeGID - The group ID (gid) attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/gid
	FSItemAttributeGID FSItemAttribute = 0
	// FSItemAttributeInhibitKernelOffloadedIO - The inhibit kernel offloaded I/O attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/inhibitKernelOffloadedIO
	FSItemAttributeInhibitKernelOffloadedIO FSItemAttribute = 0
	// FSItemAttributeLinkCount - The link count attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/linkCount
	FSItemAttributeLinkCount FSItemAttribute = 0
	// FSItemAttributeMode - The mode attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/mode
	FSItemAttributeMode FSItemAttribute = 0
	// FSItemAttributeSupportsLimitedXAttrs - The supports limited extended attributes attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/supportsLimitedXAttrs
	FSItemAttributeSupportsLimitedXAttrs FSItemAttribute = 0
)

// FSItemID enum type
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier
type FSItemID uint

// FSItemType - An enumeration of item types, such as file, directory, or symbolic link.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType
type FSItemType uint

// FSMatchResult - A type that represents the recognition and usability of a probed resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult
type FSMatchResult uint

const (
	// FSMatchResultNotRecognized - The probe doesn’t recognize the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/notRecognized
	FSMatchResultNotRecognized FSMatchResult = 0
	// FSMatchResultRecognized - The probe recognizes the resource but can’t use it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/recognized
	FSMatchResultRecognized FSMatchResult = 0
	// FSMatchResultUsable - The probe recognizes the resource and is ready to use it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/usable
	FSMatchResultUsable FSMatchResult = 0
	// FSMatchResultUsableButLimited - The probe recognizes the resource and is ready to use it, but only in a limited capacity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/usableButLimited
	FSMatchResultUsableButLimited FSMatchResult = 0
)

// FSSyncFlags - Behavior flags for use with synchronization calls.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSSyncFlags
type FSSyncFlags uint

// FSAccessMask - A bitmask of access rights.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask
type FSAccessMask uint

// FSItemDeactivationOptions - Options to specify the item deactivation policy.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions
type FSItemDeactivationOptions uint

// FSVolumeOpenModes - Defined modes for opening a file.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/OpenModes
type FSVolumeOpenModes uint

// FSPreallocateFlags - Behavior flags for preallocation operations.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags
type FSPreallocateFlags uint

// FSSetXattrPolicy - Flags to specify the policy when setting extended file attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy
type FSSetXattrPolicy uint


