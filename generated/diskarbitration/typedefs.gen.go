// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration
import (
"unsafe"
)

// Type aliases and typedefs
// DADiskDisappearedCallback - Type of the callback function used by DARegisterDiskDisappearedCallback().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskDisappearedCallback
// DADiskDisappearedCallback is a callback function
// C type: void (*)(struct __DADisk *, void *)
type DADiskDisappearedCallback = func(unsafe.Pointer, unsafe.Pointer)
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
// DADiskOptions - Options for DADiskGetOptions() and DADiskSetOptions().
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DADiskOptions
// DADiskOptions has base type: UInt32
type DADiskOptions uintptr
// DASessionRef - Type of a reference to DASession instances.
//
// [Full Topic]: https://developer.apple.com/documentation/DiskArbitration/DASession
// DASessionRef has base type: struct __DASession *
type DASessionRef uintptr

