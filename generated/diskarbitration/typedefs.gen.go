// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration
import (
"unsafe"
)

// Type aliases and typedefs
// DADiskRef - Type of a reference to DADisk instances.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADisk
// DADiskRef has base type: struct __DADisk *
type DADiskRef uintptr
// DADiskAppearedCallback - Type of the callback function used by DARegisterDiskAppearedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskAppearedCallback
// DADiskAppearedCallback is a callback function
// C type: void (*)(struct __DADisk *, void *)
type DADiskAppearedCallback = func(unsafe.Pointer, unsafe.Pointer)
// DADiskClaimCallback - Type of the callback function used by DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimCallback
// DADiskClaimCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskClaimCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// DADiskClaimOptions - Options for DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimOptions
// DADiskClaimOptions has base type: UInt32
type DADiskClaimOptions uintptr
// DADiskClaimReleaseCallback - Type of the callback function used by DADiskClaim().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskClaimReleaseCallback
// DADiskClaimReleaseCallback is a callback function
// C type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskClaimReleaseCallback = func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// DADiskDescriptionChangedCallback - Type of the callback function used by DARegisterDiskDescriptionChangedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskDescriptionChangedCallback
// DADiskDescriptionChangedCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __CFArray *, void *)
type DADiskDescriptionChangedCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// DADiskDisappearedCallback - Type of the callback function used by DARegisterDiskDisappearedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskDisappearedCallback
// DADiskDisappearedCallback is a callback function
// C type: void (*)(struct __DADisk *, void *)
type DADiskDisappearedCallback = func(unsafe.Pointer, unsafe.Pointer)
// DADiskEjectApprovalCallback - Type of the callback function used by DARegisterDiskEjectApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectApprovalCallback
// DADiskEjectApprovalCallback is a callback function
// C type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskEjectApprovalCallback = func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// DADiskEjectCallback - Type of the callback function used by DADiskEject().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectCallback
// DADiskEjectCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskEjectCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// DADiskEjectOptions - Options for DADiskEject().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskEjectOptions
// DADiskEjectOptions has base type: UInt32
type DADiskEjectOptions uintptr
// DADiskMountApprovalCallback - Type of the callback function used by DARegisterDiskMountApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountApprovalCallback
// DADiskMountApprovalCallback is a callback function
// C type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskMountApprovalCallback = func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// DADiskMountCallback - Type of the callback function used by DADiskMount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskMountCallback
// DADiskMountCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskMountCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
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
// DADiskPeekCallback is a callback function
// C type: void (*)(struct __DADisk *, void *)
type DADiskPeekCallback = func(unsafe.Pointer, unsafe.Pointer)
// DADiskRenameCallback - Type of the callback function used by DADiskRename().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameCallback
// DADiskRenameCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskRenameCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// DADiskRenameOptions - Options for DADiskRename().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskRenameOptions
// DADiskRenameOptions has base type: UInt32
type DADiskRenameOptions uintptr
// DADiskUnmountApprovalCallback - Type of the callback function used by DARegisterDiskUnmountApprovalCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountApprovalCallback
// DADiskUnmountApprovalCallback is a callback function
// C type: const struct __DADissenter *(*)(struct __DADisk *, void *)
type DADiskUnmountApprovalCallback = func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// DADiskUnmountCallback - Type of the callback function used by DADiskUnmount().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmountCallback
// DADiskUnmountCallback is a callback function
// C type: void (*)(struct __DADisk *, const struct __DADissenter *, void *)
type DADiskUnmountCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
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

