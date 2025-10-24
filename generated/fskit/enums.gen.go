// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

/* debug [enums.gen.go]: Generating 17 enums for FSKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum FSBlockmapFlags (2 cases) */
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

/* debug [enums.gen.go]: Processing enum FSCompleteIOFlags (3 cases) */
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

/* debug [enums.gen.go]: Processing enum FSErrorCode (7 cases) */
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

/* debug [enums.gen.go]: Processing enum FSExtentType (2 cases) */
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

/* debug [enums.gen.go]: Processing enum FSItemAttribute (18 cases) */
// FSItemAttribute - A value that indicates a set of item attributes to get or set.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute
type FSItemAttribute uint

const (
	// FSItemAttributeAccessTime - The last-accessed time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/accessTime
	FSItemAttributeAccessTime FSItemAttribute = 0
	// FSItemAttributeAddedTime - The time added attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/addedTime
	FSItemAttributeAddedTime FSItemAttribute = 0
	// FSItemAttributeAllocSize - The allocated size attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/allocSize
	FSItemAttributeAllocSize FSItemAttribute = 0
	// FSItemAttributeBackupTime - The backup time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/backupTime
	FSItemAttributeBackupTime FSItemAttribute = 0
	// FSItemAttributeBirthTime - The creation time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/birthTime
	FSItemAttributeBirthTime FSItemAttribute = 0
	// FSItemAttributeChangeTime - The last-changed time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/changeTime
	FSItemAttributeChangeTime FSItemAttribute = 0
	// FSItemAttributeFileID - The file ID attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/fileID
	FSItemAttributeFileID FSItemAttribute = 0
	// FSItemAttributeFlags - The flags attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/flags
	FSItemAttributeFlags FSItemAttribute = 0
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
	// FSItemAttributeModifyTime - The last-modified time attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/modifyTime
	FSItemAttributeModifyTime FSItemAttribute = 0
	// FSItemAttributeParentID - The parent ID attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/parentID
	FSItemAttributeParentID FSItemAttribute = 0
	// FSItemAttributeSize - The size attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/size
	FSItemAttributeSize FSItemAttribute = 0
	// FSItemAttributeSupportsLimitedXAttrs - The supports limited extended attributes attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/supportsLimitedXAttrs
	FSItemAttributeSupportsLimitedXAttrs FSItemAttribute = 0
	// FSItemAttributeType - The type attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/type
	FSItemAttributeType FSItemAttribute = 0
	// FSItemAttributeUID - The user ID (uid) attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Attribute/uid
	FSItemAttributeUID FSItemAttribute = 0
)

/* debug [enums.gen.go]: Processing enum FSItemID (3 cases) */
// FSItemID - The unique identifier for an item.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier
type FSItemID uint

const (
	// FSItemIDInvalid - The identifier for an invalid item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier/invalid
	FSItemIDInvalid FSItemID = 0
	// FSItemIDParentOfRoot - The identifier for an item that serves as the parent of the root directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier/parentOfRoot
	FSItemIDParentOfRoot FSItemID = 0
	// FSItemIDRootDirectory - The item identifier for the root directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier/rootDirectory
	FSItemIDRootDirectory FSItemID = 0
)

/* debug [enums.gen.go]: Processing enum FSItemType (8 cases) */
// FSItemType - An enumeration of item types, such as file, directory, or symbolic link.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType
type FSItemType uint

const (
	// FSItemTypeBlockDevice - The item type of a block device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/blockDevice
	FSItemTypeBlockDevice FSItemType = 0
	// FSItemTypeCharDevice - The item type of a character device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/charDevice
	FSItemTypeCharDevice FSItemType = 0
	// FSItemTypeDirectory - The item type of a directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/directory
	FSItemTypeDirectory FSItemType = 0
	// FSItemTypeFIFO - The item type of a first-in/first-out named pipe.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/fifo
	FSItemTypeFIFO FSItemType = 0
	// FSItemTypeFile - The item type of a regular file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/file
	FSItemTypeFile FSItemType = 0
	// FSItemTypeSocket - The item type of a socket.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/socket
	FSItemTypeSocket FSItemType = 0
	// FSItemTypeSymlink - The item type of a symbolic link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/symlink
	FSItemTypeSymlink FSItemType = 0
	// FSItemTypeUnknown - The item type of an unknown item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/unknown
	FSItemTypeUnknown FSItemType = 0
)

/* debug [enums.gen.go]: Processing enum FSMatchResult (4 cases) */
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

/* debug [enums.gen.go]: Processing enum FSAccessMask (17 cases) */
// FSAccessMask - A bitmask of access rights.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask
type FSAccessMask uint

const (
	// FSAccessAddFile - The file system allows adding files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/addFile
	FSAccessAddFile FSAccessMask = 0
	// FSAccessAddSubdirectory - The file system allows adding subdirectories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/addSubdirectory
	FSAccessAddSubdirectory FSAccessMask = 0
	// FSAccessAppendData - The file system allows appending data to a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/appendData
	FSAccessAppendData FSAccessMask = 0
	// FSAccessDelete - The file system allows deleting a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/delete
	FSAccessDelete FSAccessMask = 0
	// FSAccessDeleteChild - The file system allows deleting subdirectories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/deleteChild
	FSAccessDeleteChild FSAccessMask = 0
	// FSAccessExecute - The file system allows file executuion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/execute
	FSAccessExecute FSAccessMask = 0
	// FSAccessListDirectory - The file system allows listing directory contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/listDirectory
	FSAccessListDirectory FSAccessMask = 0
	// FSAccessReadAttributes - The file system allows reading file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/readAttributes
	FSAccessReadAttributes FSAccessMask = 0
	// FSAccessReadData - The file system allows reading data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/readData
	FSAccessReadData FSAccessMask = 0
	// FSAccessReadSecurity - The file system allows reading a file’s security descriptors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/readSecurity
	FSAccessReadSecurity FSAccessMask = 0
	// FSAccessReadXattr - The file system allows reading extended file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/readXattr
	FSAccessReadXattr FSAccessMask = 0
	// FSAccessSearch - The file system allows searching files.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/search
	FSAccessSearch FSAccessMask = 0
	// FSAccessTakeOwnership - The file system allows taking ownership of a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/takeOwnership
	FSAccessTakeOwnership FSAccessMask = 0
	// FSAccessWriteAttributes - The file system allows writing file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/writeAttributes
	FSAccessWriteAttributes FSAccessMask = 0
	// FSAccessWriteData - The file system allows writing data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/writeData
	FSAccessWriteData FSAccessMask = 0
	// FSAccessWriteSecurity - The file system allows writing a file’s security descriptors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/writeSecurity
	FSAccessWriteSecurity FSAccessMask = 0
	// FSAccessWriteXattr - The file system allows writing extended file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/AccessMask/writeXattr
	FSAccessWriteXattr FSAccessMask = 0
)

/* debug [enums.gen.go]: Processing enum FSVolumeCaseFormat (3 cases) */
// FSVolumeCaseFormat - An enumeration of case-sensitivity support types.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/CaseFormat
type FSVolumeCaseFormat uint

const (
	// FSVolumeCaseFormatInsensitive - The volume isn’t case sensitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/CaseFormat/insensitive
	FSVolumeCaseFormatInsensitive FSVolumeCaseFormat = 0
	// FSVolumeCaseFormatInsensitiveCasePreserving - The volume isn’t case sensitive, but supports preserving the case of file and directory names.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/CaseFormat/insensitiveCasePreserving
	FSVolumeCaseFormatInsensitiveCasePreserving FSVolumeCaseFormat = 0
	// FSVolumeCaseFormatSensitive - The volume is case sensitive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/CaseFormat/sensitive
	FSVolumeCaseFormatSensitive FSVolumeCaseFormat = 0
)

/* debug [enums.gen.go]: Processing enum FSItemDeactivationOptions (4 cases) */
// FSItemDeactivationOptions - Options to specify the item deactivation policy.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions
type FSItemDeactivationOptions uint

const (
	// FSItemDeactivationAlways - An option to always perform deactivation calls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions/always
	FSItemDeactivationAlways FSItemDeactivationOptions = 0
	// FSItemDeactivationForPreallocatedItems - An option to process deactivation for for files with preallocated space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions/forPreallocatedItems
	FSItemDeactivationForPreallocatedItems FSItemDeactivationOptions = 0
	// FSItemDeactivationForRemovedItems - An option to process deactivation for open-unlinked items at the moment of last close.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivationOptions/forRemovedItems
	FSItemDeactivationForRemovedItems FSItemDeactivationOptions = 0
	// FSItemDeactivationNever - An option to never perform deactivation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItemDeactivationOptions/FSItemDeactivationNever
	FSItemDeactivationNever FSItemDeactivationOptions = 0
)

/* debug [enums.gen.go]: Processing enum FSVolumeOpenModes (2 cases) */
// FSVolumeOpenModes - Defined modes for opening a file.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/OpenModes
type FSVolumeOpenModes uint

const (
	// FSVolumeOpenModesRead - The read mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/OpenModes/read
	FSVolumeOpenModesRead FSVolumeOpenModes = 0
	// FSVolumeOpenModesWrite - The write mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/OpenModes/write
	FSVolumeOpenModesWrite FSVolumeOpenModes = 0
)

/* debug [enums.gen.go]: Processing enum FSPreallocateFlags (4 cases) */
// FSPreallocateFlags - Behavior flags for preallocation operations.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags
type FSPreallocateFlags uint

const (
	// FSPreallocateFlagsAll - Allocates all requested space or no space at all.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags/all
	FSPreallocateFlagsAll FSPreallocateFlags = 0
	// FSPreallocateFlagsContiguous - Allocates contiguous space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags/contiguous
	FSPreallocateFlagsContiguous FSPreallocateFlags = 0
	// FSPreallocateFlagsFromEOF - Allocates space from the physical end of file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags/fromEOF
	FSPreallocateFlagsFromEOF FSPreallocateFlags = 0
	// FSPreallocateFlagsPersist - Allocates space that isn’t freed when deleting the descriptor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateFlags/persist
	FSPreallocateFlagsPersist FSPreallocateFlags = 0
)

/* debug [enums.gen.go]: Processing enum FSSetXattrPolicy (4 cases) */
// FSSetXattrPolicy - Flags to specify the policy when setting extended file attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy
type FSSetXattrPolicy uint

const (
	// FSSetXattrPolicyAlwaysSet - Set the value, regardless of previous state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy/alwaysSet
	FSSetXattrPolicyAlwaysSet FSSetXattrPolicy = 0
	// FSSetXattrPolicyDelete - Delete the value, failing if the extended attribute doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy/delete
	FSSetXattrPolicyDelete FSSetXattrPolicy = 0
	// FSSetXattrPolicyMustCreate - Set the value, but fail if the extended attribute already exists.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy/mustCreate
	FSSetXattrPolicyMustCreate FSSetXattrPolicy = 0
	// FSSetXattrPolicyMustReplace - Set the value, but fail if the extended attribute doesn’t already exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy/mustReplace
	FSSetXattrPolicyMustReplace FSSetXattrPolicy = 0
)

/* debug [enums.gen.go]: Processing enum FSContainerState (4 cases) */
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

/* debug [enums.gen.go]: Processing enum FSDeactivateOptions (1 cases) */
// FSDeactivateOptions - Options that affect the behavior of deactivate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDeactivateOptions
type FSDeactivateOptions uint

const (
	// FSDeactivateOptionsForce - An option to force deactivation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDeactivateOptions/force
	FSDeactivateOptionsForce FSDeactivateOptions = 0
)

/* debug [enums.gen.go]: Processing enum FSSyncFlags (3 cases) */
// FSSyncFlags - Behavior flags for use with synchronization calls.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSSyncFlags
type FSSyncFlags uint

const (
	// FSSyncFlagsDWait - A flag for synchronized I/O with data-integrity completion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSSyncFlags/dWait
	FSSyncFlagsDWait FSSyncFlags = 0
	// FSSyncFlagsNoWait - A flag for synchronized I/O that starts I/O but doesn’t wait for it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSSyncFlags/noWait
	FSSyncFlagsNoWait FSSyncFlags = 0
	// FSSyncFlagsWait - A flag for synchronized I/O with file-integrity completion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSSyncFlags/wait
	FSSyncFlagsWait FSSyncFlags = 0
)


