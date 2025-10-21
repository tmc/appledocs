// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

// Type aliases and typedefs
// DADiskRef - Type of a reference to DADisk instances.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADisk
// DADiskRef has base type: struct __DADisk *
type DADiskRef uintptr
// DADiskAppearedCallback - Type of the callback function used by DARegisterDiskAppearedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskAppearedCallback
// DADiskAppearedCallback has base type: void (*)(struct __DADisk *, void *)
type DADiskAppearedCallback uintptr
// DADiskClaimCallback - Type of the callback function used by DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimCallback
// DADiskClaimCallback has base type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskClaimCallback uintptr
// DADiskClaimOptions - Options for DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimOptions
// DADiskClaimOptions has base type: UInt32
type DADiskClaimOptions uintptr
// DADiskClaimReleaseCallback - Type of the callback function used by DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimReleaseCallback
// DADiskClaimReleaseCallback has base type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskClaimReleaseCallback uintptr
// DADiskDescriptionChangedCallback - Type of the callback function used by DARegisterDiskDescriptionChangedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskDescriptionChangedCallback
// DADiskDescriptionChangedCallback has base type: void (*)(struct __DADisk *, const struct __CFArray *, void *)
type DADiskDescriptionChangedCallback uintptr
// DADiskDisappearedCallback - Type of the callback function used by DARegisterDiskDisappearedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskDisappearedCallback
// DADiskDisappearedCallback has base type: void (*)(struct __DADisk *, void *)
type DADiskDisappearedCallback uintptr
// DADiskEjectApprovalCallback - Type of the callback function used by DARegisterDiskEjectApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectApprovalCallback
// DADiskEjectApprovalCallback has base type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskEjectApprovalCallback uintptr
// DADiskEjectCallback - Type of the callback function used by DADiskEject().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectCallback
// DADiskEjectCallback has base type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskEjectCallback uintptr
// DADiskEjectOptions - Options for DADiskEject().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectOptions
// DADiskEjectOptions has base type: UInt32
type DADiskEjectOptions uintptr
// DADiskMountApprovalCallback - Type of the callback function used by DARegisterDiskMountApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountApprovalCallback
// DADiskMountApprovalCallback has base type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskMountApprovalCallback uintptr
// DADiskMountCallback - Type of the callback function used by DADiskMount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountCallback
// DADiskMountCallback has base type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskMountCallback uintptr
// DADiskMountOptions - Options for DADiskMount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountOptions
// DADiskMountOptions has base type: UInt32
type DADiskMountOptions uintptr
// DADiskOptions - Options for DADiskGetOptions() and DADiskSetOptions().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskOptions
// DADiskOptions has base type: UInt32
type DADiskOptions uintptr
// DADiskPeekCallback - Type of the callback function used by DARegisterDiskPeekCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskPeekCallback
// DADiskPeekCallback has base type: void (*)(struct __DADisk *, void *)
type DADiskPeekCallback uintptr
// DADiskRenameCallback - Type of the callback function used by DADiskRename().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameCallback
// DADiskRenameCallback has base type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskRenameCallback uintptr
// DADiskRenameOptions - Options for DADiskRename().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameOptions
// DADiskRenameOptions has base type: UInt32
type DADiskRenameOptions uintptr
// DADiskUnmountApprovalCallback - Type of the callback function used by DARegisterDiskUnmountApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountApprovalCallback
// DADiskUnmountApprovalCallback has base type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskUnmountApprovalCallback uintptr
// DADiskUnmountCallback - Type of the callback function used by DADiskUnmount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountCallback
// DADiskUnmountCallback has base type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskUnmountCallback uintptr
// DADiskUnmountOptions - Options for DADiskUnmount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountOptions
// DADiskUnmountOptions has base type: UInt32
type DADiskUnmountOptions uintptr
// DADissenterRef - Type of a reference to DADissenter instances.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADissenter
// DADissenterRef has base type: const struct __DADissenter *
type DADissenterRef uintptr
// DAReturn - A return code.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DAReturn
// DAReturn has base type: mach_error_t
type DAReturn uintptr
// DASessionRef - Type of a reference to DASession instances.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASession
// DASessionRef has base type: struct __DASession *
type DASessionRef uintptr

