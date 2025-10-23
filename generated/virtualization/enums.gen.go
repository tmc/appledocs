// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

// Enum types and constants
// VZDiskImageCachingMode - An integer that describes the disk image caching mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode
type VZDiskImageCachingMode uint

const (
	// VZDiskImageCachingModeAutomatic - Allows the virtualization framework to automatically determine whether to enable data caching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode/automatic
	VZDiskImageCachingModeAutomatic VZDiskImageCachingMode = 0
	// VZDiskImageCachingModeCached - Enables data caching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode/cached
	VZDiskImageCachingModeCached VZDiskImageCachingMode = 0
	// VZDiskImageCachingModeUncached - Disables data caching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageCachingMode/uncached
	VZDiskImageCachingModeUncached VZDiskImageCachingMode = 0
)

// VZDiskSynchronizationMode - Values that describe the synchronization modes available to the guest OS.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
type VZDiskSynchronizationMode uint

// VZMacAuxiliaryStorageInitializationOptions - Options you can set when creating new auxiliary storage.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/InitializationOptions
type VZMacAuxiliaryStorageInitializationOptions uint

const (
	// VZMacAuxiliaryStorageInitializationOptionAllowOverwrite - A Boolean value that indicates whether the VM can overwrite an existing auxiliary storage file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacAuxiliaryStorage/InitializationOptions/allowOverwrite
	VZMacAuxiliaryStorageInitializationOptionAllowOverwrite VZMacAuxiliaryStorageInitializationOptions = 0
)


