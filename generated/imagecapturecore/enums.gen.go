// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

/* debug [enums.gen.go]: Generating 23 enums for ImageCaptureCore */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ICLegacyReturnCode (22 cases) */
// ICLegacyReturnCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code
type ICLegacyReturnCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/cannotYieldDevice
	ICLegacyReturnCodeCannotYieldDevice ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/communicationErr
	ICLegacyReturnCodeCommunicationErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/dataTypeNotFoundErr
	ICLegacyReturnCodeDataTypeNotFoundErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceAlreadyOpenErr
	ICLegacyReturnCodeDeviceAlreadyOpenErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceGUIDNotFoundErr
	ICLegacyReturnCodeDeviceGUIDNotFoundErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceInternalErr
	ICLegacyReturnCodeDeviceInternalErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceInvalidParamErr
	ICLegacyReturnCodeDeviceInvalidParamErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceIOServicePathNotFoundErr
	ICLegacyReturnCodeDeviceIOServicePathNotFoundErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceLocationIDNotFoundErr
	ICLegacyReturnCodeDeviceLocationIDNotFoundErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceMemoryAllocationErr
	ICLegacyReturnCodeDeviceMemoryAllocationErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceNotFoundErr
	ICLegacyReturnCodeDeviceNotFoundErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceNotOpenErr
	ICLegacyReturnCodeDeviceNotOpenErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceUnsupportedErr
	ICLegacyReturnCodeDeviceUnsupportedErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/extensionInternalErr
	ICLegacyReturnCodeExtensionInternalErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/fileCorruptedErr
	ICLegacyReturnCodeFileCorruptedErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/frameworkInternalErr
	ICLegacyReturnCodeFrameworkInternalErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/indexOutOfRangeErr
	ICLegacyReturnCodeIndexOutOfRangeErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/invalidObjectErr
	ICLegacyReturnCodeInvalidObjectErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/invalidPropertyErr
	ICLegacyReturnCodeInvalidPropertyErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/invalidSessionErr
	ICLegacyReturnCodeInvalidSessionErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/ioPendingErr
	ICLegacyReturnCodeIOPendingErr ICLegacyReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/propertyTypeNotFoundErr
	ICLegacyReturnCodePropertyTypeNotFoundErr ICLegacyReturnCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnCode (40 cases) */
// ICReturnCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code
type ICReturnCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/communicationTimedOut
	ICReturnCommunicationTimedOut ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deleteFilesCanceled
	ICReturnDeleteFilesCanceled ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deleteFilesFailed
	ICReturnDeleteFilesFailed ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceCommandGeneralFailure
	ICReturnDeviceCommandGeneralFailure ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceCouldNotPair
	ICReturnDeviceCouldNotPair ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceCouldNotUnpair
	ICReturnDeviceCouldNotUnpair ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceFailedToCloseSession
	ICReturnDeviceFailedToCloseSession ICReturnCode = 0
	// ICReturnDeviceFailedToCompleteTransfer - Failed to complete a data transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceFailedToCompleteTransfer
	ICReturnDeviceFailedToCompleteTransfer ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceFailedToOpenSession
	ICReturnDeviceFailedToOpenSession ICReturnCode = 0
	// ICReturnDeviceFailedToSendData - Failed to send data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceFailedToSendData
	ICReturnDeviceFailedToSendData ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceFailedToTakePicture
	ICReturnDeviceFailedToTakePicture ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceIsBusyEnumerating
	ICReturnDeviceIsBusyEnumerating ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceIsPasscodeLocked
	ICReturnDeviceIsPasscodeLocked ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceNeedsCredentials
	ICReturnDeviceNeedsCredentials ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareInstallationCanceled
	ICReturnDeviceSoftwareInstallationCanceled ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareInstallationCompleted
	ICReturnDeviceSoftwareInstallationCompleted ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareInstallationFailed
	ICReturnDeviceSoftwareInstallationFailed ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareIsBeingInstalled
	ICReturnDeviceSoftwareIsBeingInstalled ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareNotAvailable
	ICReturnDeviceSoftwareNotAvailable ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/deviceSoftwareNotInstalled
	ICReturnDeviceSoftwareNotInstalled ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/downloadCanceled
	ICReturnDownloadCanceled ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/downloadFailed
	ICReturnDownloadFailed ICReturnCode = 0
	// ICReturnExFATVolumeInvalid - EXFAT volume is invalid, and cannot be enumerated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/exFATVolumeInvalid
	ICReturnExFATVolumeInvalid ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/failedToCompletePassThroughCommand
	ICReturnFailedToCompletePassThroughCommand ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/failedToCompleteSendMessageRequest
	ICReturnFailedToCompleteSendMessageRequest ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/failedToDisabeTethering
	ICReturnFailedToDisabeTethering ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/failedToEnabeTethering
	ICReturnFailedToEnabeTethering ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/invalidParam
	ICReturnInvalidParam ICReturnCode = 0
	// ICReturnMultiErrorDictionary - Multierror
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/multiErrorDictionary
	ICReturnMultiErrorDictionary ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/receivedUnsolicitedScannerErrorInfo
	ICReturnReceivedUnsolicitedScannerErrorInfo ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/receivedUnsolicitedScannerStatusInfo
	ICReturnReceivedUnsolicitedScannerStatusInfo ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scannerFailedToCompleteOverviewScan
	ICReturnScannerFailedToCompleteOverviewScan ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scannerFailedToCompleteScan
	ICReturnScannerFailedToCompleteScan ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scannerFailedToSelectFunctionalUnit
	ICReturnScannerFailedToSelectFunctionalUnit ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scannerInUseByLocalUser
	ICReturnScannerInUseByLocalUser ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scannerInUseByRemoteUser
	ICReturnScannerInUseByRemoteUser ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/scanOperationCanceled
	ICReturnScanOperationCanceled ICReturnCode = 0
	// ICReturnSessionNotOpened - Session is not open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/sessionNotOpened
	ICReturnSessionNotOpened ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/success
	ICReturnSuccess ICReturnCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturn/Code/uploadFailed
	ICReturnUploadFailed ICReturnCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnConnectionErrorCode (8 cases) */
// ICReturnConnectionErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code
type ICReturnConnectionErrorCode uint

const (
	// ICReturnConnectionClosedSessionSuddenly - Device closed session without request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/closedSessionSuddenly
	ICReturnConnectionClosedSessionSuddenly ICReturnConnectionErrorCode = 0
	// ICReturnConnectionDriverExited - Device driver exited without request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/driverExited
	ICReturnConnectionDriverExited ICReturnConnectionErrorCode = 0
	// ICReturnConnectionEjectedSuddenly - Device ejected without request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/ejectedSuddenly
	ICReturnConnectionEjectedSuddenly ICReturnConnectionErrorCode = 0
	// ICReturnConnectionEjectFailed - Device reports eject has failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/ejectFailed
	ICReturnConnectionEjectFailed ICReturnConnectionErrorCode = 0
	// ICReturnConnectionFailedToOpen - Failed to open a connection to the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/failedToOpen
	ICReturnConnectionFailedToOpen ICReturnConnectionErrorCode = 0
	// ICReturnConnectionFailedToOpenDevice - Failed to open the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/failedToOpenDevice
	ICReturnConnectionFailedToOpenDevice ICReturnConnectionErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/notAuthorizedToOpenDevice
	ICReturnConnectionNotAuthorizedToOpenDevice ICReturnConnectionErrorCode = 0
	// ICReturnConnectionSessionAlreadyOpen - Device reports session is already open.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code/sessionAlreadyOpen
	ICReturnConnectionSessionAlreadyOpen ICReturnConnectionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnDownloadErrorCode (2 cases) */
// ICReturnDownloadErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code
type ICReturnDownloadErrorCode uint

const (
	// ICReturnDownloadFileWritable - The destination file is not writable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code/fileWritable
	ICReturnDownloadFileWritable ICReturnDownloadErrorCode = 0
	// ICReturnDownloadPathInvalid - The destination path is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code/pathInvalid
	ICReturnDownloadPathInvalid ICReturnDownloadErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnMetadataErrorCode (4 cases) */
// ICReturnMetadataErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code
type ICReturnMetadataErrorCode uint

const (
	// ICReturnMetadataAlreadyFetching - Item metadata request is being serviced.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code/alreadyFetching
	ICReturnMetadataAlreadyFetching ICReturnMetadataErrorCode = 0
	// ICReturnMetadataCanceled - Item metadata request has been canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code/canceled
	ICReturnMetadataCanceled ICReturnMetadataErrorCode = 0
	// ICReturnMetadataInvalid - Item metadata request completed with invalid result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code/invalid
	ICReturnMetadataInvalid ICReturnMetadataErrorCode = 0
	// ICReturnMetadataNotAvailable - Item does not have metadata available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnMetadataError/Code/notAvailable
	ICReturnMetadataNotAvailable ICReturnMetadataErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnObjectErrorCode (5 cases) */
// ICReturnObjectErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code
type ICReturnObjectErrorCode uint

const (
	// ICReturnCodeObjectCouldNotBeRead - The object could not be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code/codeObjectCouldNotBeRead
	ICReturnCodeObjectCouldNotBeRead ICReturnObjectErrorCode = 0
	// ICReturnCodeObjectDataEmpty - The object data is empty.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code/codeObjectDataEmpty
	ICReturnCodeObjectDataEmpty ICReturnObjectErrorCode = 0
	// ICReturnCodeObjectDataOffsetInvalid - The object data offset is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code/codeObjectDataOffsetInvalid
	ICReturnCodeObjectDataOffsetInvalid ICReturnObjectErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code/codeObjectDataRequestTooLarge
	ICReturnCodeObjectDataRequestTooLarge ICReturnObjectErrorCode = 0
	// ICReturnCodeObjectDoesNotExist - The object does not exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code/codeObjectDoesNotExist
	ICReturnCodeObjectDoesNotExist ICReturnObjectErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnPTPDeviceErrorCode (2 cases) */
// ICReturnPTPDeviceErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnPTPDeviceError/Code
type ICReturnPTPDeviceErrorCode uint

const (
	// ICReturnPTPFailedToSendCommand - Sending a PTP command failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnPTPDeviceError/Code/failedToSendCommand
	ICReturnPTPFailedToSendCommand ICReturnPTPDeviceErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnPTPDeviceError/Code/notAuthorizedToSendCommand
	ICReturnPTPNotAuthorizedToSendCommand ICReturnPTPDeviceErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICReturnThumbnailErrorCode (4 cases) */
// ICReturnThumbnailErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code
type ICReturnThumbnailErrorCode uint

const (
	// ICReturnThumbnailAlreadyFetching - Item thumbnail request is being serviced.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/alreadyFetching
	ICReturnThumbnailAlreadyFetching ICReturnThumbnailErrorCode = 0
	// ICReturnThumbnailCanceled - Item thumbnail request has been canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/canceled
	ICReturnThumbnailCanceled ICReturnThumbnailErrorCode = 0
	// ICReturnThumbnailInvalid - Item thumbnail request completed with invalid result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/invalid
	ICReturnThumbnailInvalid ICReturnThumbnailErrorCode = 0
	// ICReturnThumbnailNotAvailable - Item does not have thumbnail available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/notAvailable
	ICReturnThumbnailNotAvailable ICReturnThumbnailErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ICDeviceLocationType (4 cases) */
// ICDeviceLocationType - The location of the image capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType
type ICDeviceLocationType uint

const (
	// ICDeviceLocationTypeBluetooth - A paired Bluetooth device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType/bluetooth
	ICDeviceLocationTypeBluetooth ICDeviceLocationType = 0
	// ICDeviceLocationTypeBonjour - A supported Bonjour services device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationType/bonjour
	ICDeviceLocationTypeBonjour ICDeviceLocationType = 0
	// ICDeviceLocationTypeLocal - A device that's directly attached to the Mac through its USB or FireWire port.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicelocationtype/icdevicelocationtypelocal
	ICDeviceLocationTypeLocal ICDeviceLocationType = 0
	// ICDeviceLocationTypeShared - A device that’s shared by other Mac hosts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicelocationtype/icdevicelocationtypeshared
	ICDeviceLocationTypeShared ICDeviceLocationType = 0
)

/* debug [enums.gen.go]: Processing enum ICDeviceLocationTypeMask (5 cases) */
// ICDeviceLocationTypeMask - Masks for detecting different device locations.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask
type ICDeviceLocationTypeMask uint

const (
	// ICDeviceLocationTypeMaskBluetooth - A mask for detecting a paired Bluetooth device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask/bluetooth
	ICDeviceLocationTypeMaskBluetooth ICDeviceLocationTypeMask = 0
	// ICDeviceLocationTypeMaskBonjour - A mask for detecting a network device that publishes a Bonjour service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask/bonjour
	ICDeviceLocationTypeMaskBonjour ICDeviceLocationTypeMask = 0
	// ICDeviceLocationTypeMaskLocal - A mask for detecting a local device, such as USB or FireWire.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicelocationtypemask/icdevicelocationtypemasklocal
	ICDeviceLocationTypeMaskLocal ICDeviceLocationTypeMask = 0
	// ICDeviceLocationTypeMaskRemote - A mask for detecting a remote device, such as a shared, Bonjour, or Bluetooth device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicelocationtypemask/icdevicelocationtypemaskremote
	ICDeviceLocationTypeMaskRemote ICDeviceLocationTypeMask = 0
	// ICDeviceLocationTypeMaskShared - A mask for detecting a device shared by another Mac host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicelocationtypemask/icdevicelocationtypemaskshared
	ICDeviceLocationTypeMaskShared ICDeviceLocationTypeMask = 0
)

/* debug [enums.gen.go]: Processing enum ICDeviceType (2 cases) */
// ICDeviceType - The type of image capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceType
type ICDeviceType uint

const (
	// ICDeviceTypeCamera - The device is a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceType/camera
	ICDeviceTypeCamera ICDeviceType = 0
	// ICDeviceTypeScanner - The device is a scanner.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicetype/icdevicetypescanner
	ICDeviceTypeScanner ICDeviceType = 0
)

/* debug [enums.gen.go]: Processing enum ICDeviceTypeMask (2 cases) */
// ICDeviceTypeMask - Masks for detecting different device types.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTypeMask
type ICDeviceTypeMask uint

const (
	// ICDeviceTypeMaskCamera - A mask for detecting a camera.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTypeMask/camera
	ICDeviceTypeMaskCamera ICDeviceTypeMask = 0
	// ICDeviceTypeMaskScanner - A mask for detecting a scanner.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icdevicetypemask/icdevicetypemaskscanner
	ICDeviceTypeMaskScanner ICDeviceTypeMask = 0
)

/* debug [enums.gen.go]: Processing enum ICEXIFOrientationType (8 cases) */
// ICEXIFOrientationType - The file’s orientation type.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICEXIFOrientationType
type ICEXIFOrientationType uint

const (
	// ICEXIFOrientation1 - Normal
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation1
	ICEXIFOrientation1 ICEXIFOrientationType = 0
	// ICEXIFOrientation2 - Flipped horizontally
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation2
	ICEXIFOrientation2 ICEXIFOrientationType = 0
	// ICEXIFOrientation3 - Rotated 180°
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation3
	ICEXIFOrientation3 ICEXIFOrientationType = 0
	// ICEXIFOrientation4 - Flipped vertically
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation4
	ICEXIFOrientation4 ICEXIFOrientationType = 0
	// ICEXIFOrientation5 - Rotated 90° CCW and flipped vertically
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation5
	ICEXIFOrientation5 ICEXIFOrientationType = 0
	// ICEXIFOrientation6 - Rotated 90° CCW
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation6
	ICEXIFOrientation6 ICEXIFOrientationType = 0
	// ICEXIFOrientation7 - Rotated 90° CW and flipped vertically
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation7
	ICEXIFOrientation7 ICEXIFOrientationType = 0
	// ICEXIFOrientation8 - Rotated 90° CW
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icexiforientationtype/icexiforientation8
	ICEXIFOrientation8 ICEXIFOrientationType = 0
)

/* debug [enums.gen.go]: Processing enum ICMediaPresentation (2 cases) */
// ICMediaPresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation
type ICMediaPresentation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation/convertedAssets
	ICMediaPresentationConvertedAssets ICMediaPresentation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icmediapresentation/icmediapresentationoriginalassets
	ICMediaPresentationOriginalAssets ICMediaPresentation = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerBitDepth (3 cases) */
// ICScannerBitDepth - The number of bits per channel in the scanned image.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth
type ICScannerBitDepth uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth16Bits
	ICScannerBitDepth16Bits ICScannerBitDepth = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth1Bit
	ICScannerBitDepth1Bit ICScannerBitDepth = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth8Bits
	ICScannerBitDepth8Bits ICScannerBitDepth = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerColorDataFormatType (2 cases) */
// ICScannerColorDataFormatType - The color data formats relevant to multichannel data.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerColorDataFormatType
type ICScannerColorDataFormatType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerColorDataFormatType/chunky
	ICScannerColorDataFormatTypeChunky ICScannerColorDataFormatType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannercolordataformattype/icscannercolordataformattypeplanar
	ICScannerColorDataFormatTypePlanar ICScannerColorDataFormatType = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerDocumentType (72 cases) */
// ICScannerDocumentType - The supported document size types.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDocumentType
type ICScannerDocumentType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype10
	ICScannerDocumentType10 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype10r
	ICScannerDocumentType10R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype110
	ICScannerDocumentType110 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype11r
	ICScannerDocumentType11R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype12r
	ICScannerDocumentType12R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype135
	ICScannerDocumentType135 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype2a0
	ICScannerDocumentType2A0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype3r
	ICScannerDocumentType3R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype4a0
	ICScannerDocumentType4A0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype4r
	ICScannerDocumentType4R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype5r
	ICScannerDocumentType5R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype6r
	ICScannerDocumentType6R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttype8r
	ICScannerDocumentType8R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea0
	ICScannerDocumentTypeA0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea1
	ICScannerDocumentTypeA1 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea2
	ICScannerDocumentTypeA2 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea3
	ICScannerDocumentTypeA3 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea4
	ICScannerDocumentTypeA4 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea5
	ICScannerDocumentTypeA5 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea6
	ICScannerDocumentTypeA6 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea7
	ICScannerDocumentTypeA7 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea8
	ICScannerDocumentTypeA8 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypea9
	ICScannerDocumentTypeA9 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeapsc
	ICScannerDocumentTypeAPSC ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeapsh
	ICScannerDocumentTypeAPSH ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeapsp
	ICScannerDocumentTypeAPSP ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeb5
	ICScannerDocumentTypeB5 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypebusinesscard
	ICScannerDocumentTypeBusinessCard ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec0
	ICScannerDocumentTypeC0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec1
	ICScannerDocumentTypeC1 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec10
	ICScannerDocumentTypeC10 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec2
	ICScannerDocumentTypeC2 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec3
	ICScannerDocumentTypeC3 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec4
	ICScannerDocumentTypeC4 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec5
	ICScannerDocumentTypeC5 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec6
	ICScannerDocumentTypeC6 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec7
	ICScannerDocumentTypeC7 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec8
	ICScannerDocumentTypeC8 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypec9
	ICScannerDocumentTypeC9 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypedefault
	ICScannerDocumentTypeDefault ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypee
	ICScannerDocumentTypeE ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob0
	ICScannerDocumentTypeISOB0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob1
	ICScannerDocumentTypeISOB1 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob10
	ICScannerDocumentTypeISOB10 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob2
	ICScannerDocumentTypeISOB2 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob3
	ICScannerDocumentTypeISOB3 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob4
	ICScannerDocumentTypeISOB4 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob5
	ICScannerDocumentTypeISOB5 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob6
	ICScannerDocumentTypeISOB6 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob7
	ICScannerDocumentTypeISOB7 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob8
	ICScannerDocumentTypeISOB8 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeisob9
	ICScannerDocumentTypeISOB9 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb0
	ICScannerDocumentTypeJISB0 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb1
	ICScannerDocumentTypeJISB1 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb10
	ICScannerDocumentTypeJISB10 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb2
	ICScannerDocumentTypeJISB2 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb3
	ICScannerDocumentTypeJISB3 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb4
	ICScannerDocumentTypeJISB4 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb6
	ICScannerDocumentTypeJISB6 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb7
	ICScannerDocumentTypeJISB7 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb8
	ICScannerDocumentTypeJISB8 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypejisb9
	ICScannerDocumentTypeJISB9 ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypelf
	ICScannerDocumentTypeLF ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypemf
	ICScannerDocumentTypeMF ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypes10r
	ICScannerDocumentTypeS10R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypes12r
	ICScannerDocumentTypeS12R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypes8r
	ICScannerDocumentTypeS8R ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeusexecutive
	ICScannerDocumentTypeUSExecutive ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeusledger
	ICScannerDocumentTypeUSLedger ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeuslegal
	ICScannerDocumentTypeUSLegal ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeusletter
	ICScannerDocumentTypeUSLetter ICScannerDocumentType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdocumenttype/icscannerdocumenttypeusstatement
	ICScannerDocumentTypeUSStatement ICScannerDocumentType = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerFeatureType (4 cases) */
// ICScannerFeatureType - The types of scanner features.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureType
type ICScannerFeatureType uint

const (
	// ICScannerFeatureTypeBoolean - A feature with a value of YES or NO.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureType/boolean
	ICScannerFeatureTypeBoolean ICScannerFeatureType = 0
	// ICScannerFeatureTypeEnumeration - A feature that can have one of several discrete values, strings, or numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFeatureType/enumeration
	ICScannerFeatureTypeEnumeration ICScannerFeatureType = 0
	// ICScannerFeatureTypeRange - A feature with a value that lies within a range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturetype/icscannerfeaturetyperange
	ICScannerFeatureTypeRange ICScannerFeatureType = 0
	// ICScannerFeatureTypeTemplate - A group of one or more rectangular scan areas that can be used with a scanner functional unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfeaturetype/icscannerfeaturetypetemplate
	ICScannerFeatureTypeTemplate ICScannerFeatureType = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerFunctionalUnitState (3 cases) */
// ICScannerFunctionalUnitState - Flags to indicate the state of the scanner functional unit.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitState
type ICScannerFunctionalUnitState uint

const (
	// ICScannerFunctionalUnitStateOverviewScanInProgress - A flag indicating that the functional unit is performing an overview scan.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitstate/icscannerfunctionalunitstateoverviewscaninprogress
	ICScannerFunctionalUnitStateOverviewScanInProgress ICScannerFunctionalUnitState = 0
	// ICScannerFunctionalUnitStateReady - A flag indicating that the functional unit is ready for operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitstate/icscannerfunctionalunitstateready
	ICScannerFunctionalUnitStateReady ICScannerFunctionalUnitState = 0
	// ICScannerFunctionalUnitStateScanInProgress - A flag indicating that the functional unit is performing a scan.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitstate/icscannerfunctionalunitstatescaninprogress
	ICScannerFunctionalUnitStateScanInProgress ICScannerFunctionalUnitState = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerFunctionalUnitType (4 cases) */
// ICScannerFunctionalUnitType - The types of scanner functional units.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitType
type ICScannerFunctionalUnitType uint

const (
	// ICScannerFunctionalUnitTypeDocumentFeeder - A document feeder functional unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitType/documentFeeder
	ICScannerFunctionalUnitTypeDocumentFeeder ICScannerFunctionalUnitType = 0
	// ICScannerFunctionalUnitTypeFlatbed - A flatbed functional unit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitType/flatbed
	ICScannerFunctionalUnitTypeFlatbed ICScannerFunctionalUnitType = 0
	// ICScannerFunctionalUnitTypeNegativeTransparency - A transparency functional unit for scanning negatives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunittype/icscannerfunctionalunittypenegativetransparency
	ICScannerFunctionalUnitTypeNegativeTransparency ICScannerFunctionalUnitType = 0
	// ICScannerFunctionalUnitTypePositiveTransparency - A transparency functional unit for scanning positives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunittype/icscannerfunctionalunittypepositivetransparency
	ICScannerFunctionalUnitTypePositiveTransparency ICScannerFunctionalUnitType = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerMeasurementUnit (6 cases) */
// ICScannerMeasurementUnit - The unit of measurement used by the scanner.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerMeasurementUnit
type ICScannerMeasurementUnit uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerMeasurementUnit/centimeters
	ICScannerMeasurementUnitCentimeters ICScannerMeasurementUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannermeasurementunit/icscannermeasurementunitinches
	ICScannerMeasurementUnitInches ICScannerMeasurementUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannermeasurementunit/icscannermeasurementunitpicas
	ICScannerMeasurementUnitPicas ICScannerMeasurementUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannermeasurementunit/icscannermeasurementunitpixels
	ICScannerMeasurementUnitPixels ICScannerMeasurementUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannermeasurementunit/icscannermeasurementunitpoints
	ICScannerMeasurementUnitPoints ICScannerMeasurementUnit = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannermeasurementunit/icscannermeasurementunittwips
	ICScannerMeasurementUnitTwips ICScannerMeasurementUnit = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerPixelDataType (9 cases) */
// ICScannerPixelDataType - The pixel data types.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType
type ICScannerPixelDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType/BW
	ICScannerPixelDataTypeBW ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType/CIEXYZ
	ICScannerPixelDataTypeCIEXYZ ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType/CMY
	ICScannerPixelDataTypeCMY ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType/CMYK
	ICScannerPixelDataTypeCMYK ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType/gray
	ICScannerPixelDataTypeGray ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerpixeldatatype/icscannerpixeldatatypepalette
	ICScannerPixelDataTypePalette ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerpixeldatatype/icscannerpixeldatatypergb
	ICScannerPixelDataTypeRGB ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerpixeldatatype/icscannerpixeldatatypeyuv
	ICScannerPixelDataTypeYUV ICScannerPixelDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerpixeldatatype/icscannerpixeldatatypeyuvk
	ICScannerPixelDataTypeYUVK ICScannerPixelDataType = 0
)

/* debug [enums.gen.go]: Processing enum ICScannerTransferMode (2 cases) */
// ICScannerTransferMode - The modes for transferring scan data from the scanner functional unit.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerTransferMode
type ICScannerTransferMode uint

const (
	// ICScannerTransferModeFileBased - The mode for transferring the scan as a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerTransferMode/fileBased
	ICScannerTransferModeFileBased ICScannerTransferMode = 0
	// ICScannerTransferModeMemoryBased - The mode for transferring the scan as data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannertransfermode/icscannertransfermodememorybased
	ICScannerTransferModeMemoryBased ICScannerTransferMode = 0
)


