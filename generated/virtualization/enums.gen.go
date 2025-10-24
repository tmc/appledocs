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

// VZDiskImageSynchronizationMode - An integer that describes the disk image synchronization mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode
type VZDiskImageSynchronizationMode uint

const (
	// VZDiskImageSynchronizationModeFsync - Synchronizes data to the drive using the system’s best-effort synchronization mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode/fsync
	VZDiskImageSynchronizationModeFsync VZDiskImageSynchronizationMode = 0
	// VZDiskImageSynchronizationModeFull - Synchronizes data to the permanent storage holding the disk image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode/full
	VZDiskImageSynchronizationModeFull VZDiskImageSynchronizationMode = 0
	// VZDiskImageSynchronizationModeNone - Disables data synchronization with the permanent storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskImageSynchronizationMode/none
	VZDiskImageSynchronizationModeNone VZDiskImageSynchronizationMode = 0
)

// VZDiskSynchronizationMode - Values that describe the synchronization modes available to the guest OS.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode
type VZDiskSynchronizationMode uint

const (
	// VZDiskSynchronizationModeFull - Perform all synchronization operations as requested by the guest OS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode/full
	VZDiskSynchronizationModeFull VZDiskSynchronizationMode = 0
	// VZDiskSynchronizationModeNone - Don’t synchronize the data with the permanent storage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZDiskSynchronizationMode/none
	VZDiskSynchronizationModeNone VZDiskSynchronizationMode = 0
)

// VZErrorCode - Errors you might encounter when configuring or using a virtual machine.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code
type VZErrorCode uint

const (
	// VZErrorDeviceAlreadyAttached - The device already has an attachment to the VM.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/deviceAlreadyAttached
	VZErrorDeviceAlreadyAttached VZErrorCode = 0
	// VZErrorDeviceInitializationFailure - A device initialization failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/deviceInitializationFailure
	VZErrorDeviceInitializationFailure VZErrorCode = 0
	// VZErrorDeviceNotFound - The framework can’t find the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/deviceNotFound
	VZErrorDeviceNotFound VZErrorCode = 0
	// VZErrorInstallationFailed - An error occurred during installation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/installationFailed
	VZErrorInstallationFailed VZErrorCode = 0
	// VZErrorInstallationRequiresUpdate - The VM requires a software update in order to complete the installation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/installationRequiresUpdate
	VZErrorInstallationRequiresUpdate VZErrorCode = 0
	// VZErrorInternal - An internal error, such as the VM unexpectedly stopping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/internalError
	VZErrorInternal VZErrorCode = 0
	// VZErrorInvalidDiskImage - An invalid disk-image error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidDiskImage
	VZErrorInvalidDiskImage VZErrorCode = 0
	// VZErrorInvalidRestoreImage - The restore image is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidRestoreImage
	VZErrorInvalidRestoreImage VZErrorCode = 0
	// VZErrorInvalidRestoreImageCatalog - The restore image catalog is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidRestoreImageCatalog
	VZErrorInvalidRestoreImageCatalog VZErrorCode = 0
	// VZErrorInvalidVirtualMachineConfiguration - An invalid configuration error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidVirtualMachineConfiguration
	VZErrorInvalidVirtualMachineConfiguration VZErrorCode = 0
	// VZErrorInvalidVirtualMachineState - An invalid state error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidVirtualMachineState
	VZErrorInvalidVirtualMachineState VZErrorCode = 0
	// VZErrorInvalidVirtualMachineStateTransition - An invalid state transition error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/invalidVirtualMachineStateTransition
	VZErrorInvalidVirtualMachineStateTransition VZErrorCode = 0
	// VZErrorNetworkBlockDeviceDisconnected - The network block device client disconnected from the server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/networkBlockDeviceDisconnected
	VZErrorNetworkBlockDeviceDisconnected VZErrorCode = 0
	// VZErrorNetworkBlockDeviceNegotiationFailed - The connection or the negotiation with the network block device server failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/networkBlockDeviceNegotiationFailed
	VZErrorNetworkBlockDeviceNegotiationFailed VZErrorCode = 0
	// VZErrorNetworkError - A network error, such as a failed connection error, occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/networkError
	VZErrorNetworkError VZErrorCode = 0
	// VZErrorNoSupportedRestoreImagesInCatalog - The restore image catalog has no supported restore images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/noSupportedRestoreImagesInCatalog
	VZErrorNoSupportedRestoreImagesInCatalog VZErrorCode = 0
	// VZErrorNotSupported - The host computer or operating system isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/notSupported
	VZErrorNotSupported VZErrorCode = 0
	// VZErrorOperationCancelled - The code that indicates user canceled the installation of Rosetta or the app canceled the installation of a guest OS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/operationCancelled
	VZErrorOperationCancelled VZErrorCode = 0
	// VZErrorOutOfDiskSpace - The host is out of disk space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/outOfDiskSpace
	VZErrorOutOfDiskSpace VZErrorCode = 0
	// VZErrorRestore - The VM failed to restore from save file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/restore
	VZErrorRestore VZErrorCode = 0
	// VZErrorRestoreImageCatalogLoadFailed - The restore image catalog failed to load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/restoreImageCatalogLoadFailed
	VZErrorRestoreImageCatalogLoadFailed VZErrorCode = 0
	// VZErrorRestoreImageLoadFailed - The restore image failed to load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/restoreImageLoadFailed
	VZErrorRestoreImageLoadFailed VZErrorCode = 0
	// VZErrorSave - The VM failed to save to the save file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/save
	VZErrorSave VZErrorCode = 0
	// VZErrorUSBControllerNotFound - The framework can’t find the controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/usbControllerNotFound
	VZErrorUSBControllerNotFound VZErrorCode = 0
	// VZErrorVirtualMachineLimitExceeded - Unable to create an additional VM.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZError/Code/virtualMachineLimitExceeded
	VZErrorVirtualMachineLimitExceeded VZErrorCode = 0
)

// VZLinuxRosettaAvailability - Constants that describe the availability and installation status of Rosetta.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability
type VZLinuxRosettaAvailability uint

const (
	// VZLinuxRosettaAvailabilityInstalled - Rosetta is available on the host system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability/installed
	VZLinuxRosettaAvailabilityInstalled VZLinuxRosettaAvailability = 0
	// VZLinuxRosettaAvailabilityNotInstalled - Rosetta isn’t installed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability/notInstalled
	VZLinuxRosettaAvailabilityNotInstalled VZLinuxRosettaAvailability = 0
	// VZLinuxRosettaAvailabilityNotSupported - The current hardware or software configuration doesn’t support Rosetta.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability/notSupported
	VZLinuxRosettaAvailabilityNotSupported VZLinuxRosettaAvailability = 0
)

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

// VZVirtualMachineState - The execution states of the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum
type VZVirtualMachineState uint

const (
	// VZVirtualMachineStateError - The VM encountered an unrecoverable error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/error
	VZVirtualMachineStateError VZVirtualMachineState = 0
	// VZVirtualMachineStatePaused - The framework has paused a started VM.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/paused
	VZVirtualMachineStatePaused VZVirtualMachineState = 0
	// VZVirtualMachineStatePausing - The VM is pausing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/pausing
	VZVirtualMachineStatePausing VZVirtualMachineState = 0
	// VZVirtualMachineStateRestoring - The VM is restoring from a saved state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/restoring
	VZVirtualMachineStateRestoring VZVirtualMachineState = 0
	// VZVirtualMachineStateResuming - The VM is resuming from a paused state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/resuming
	VZVirtualMachineStateResuming VZVirtualMachineState = 0
	// VZVirtualMachineStateRunning - The VM is running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/running
	VZVirtualMachineStateRunning VZVirtualMachineState = 0
	// VZVirtualMachineStateSaving - The VM is saving its state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/saving
	VZVirtualMachineStateSaving VZVirtualMachineState = 0
	// VZVirtualMachineStateStarting - The VM is configuring the hardware preparing to run.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/starting
	VZVirtualMachineStateStarting VZVirtualMachineState = 0
	// VZVirtualMachineStateStopped - The VM isn’t running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/stopped
	VZVirtualMachineStateStopped VZVirtualMachineState = 0
	// VZVirtualMachineStateStopping - The VM is stopping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtualMachine/State-swift.enum/stopping
	VZVirtualMachineStateStopping VZVirtualMachineState = 0
)


