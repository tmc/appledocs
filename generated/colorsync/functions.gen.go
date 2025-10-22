// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ColorSync Functions (56 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CGDisplayCreateUUIDFromDisplayID func(unsafe.Pointer) unsafe.Pointer
	_CGDisplayGetDisplayIDFromUUID func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncAPIVersion func() unsafe.Pointer
	_ColorSyncCMMCopyCMMIdentifier func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncCMMCopyLocalizedName func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncCMMGetBundle func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncCreateCodeFragment func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncDeviceCopyDeviceInfo func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncDeviceSetCustomProfiles func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncIterateInstalledCMMs func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncIterateInstalledProfiles func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncIterateInstalledProfilesWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileContainsTag func(unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncProfileCopyData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCopyDescriptionString func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCopyHeader func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateWithDisplayID func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateWithURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateWithURLAndOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileEstimateGammaWithDisplayID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileGetMD5 func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileGetTagCount func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileGetTypeID func() unsafe.Pointer
	_ColorSyncProfileInstall func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncProfileIsHLGBased func(unsafe.Pointer) bool
	_ColorSyncProfileIsMatrixBased func(unsafe.Pointer) bool
	_ColorSyncProfileIsPQBased func(unsafe.Pointer) bool
	_ColorSyncProfileIsWideGamut func(unsafe.Pointer) bool
	_ColorSyncProfileVerify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncTransformGetProfileSequence func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncCMMCreate func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncCMMGetTypeID func() unsafe.Pointer
	_ColorSyncIterateDeviceProfiles func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCopyTag func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCopyTagSignatures func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateDeviceProfile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateDisplayTransferTablesFromVCGT func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateLink func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateMutable func() unsafe.Pointer
	_ColorSyncProfileCreateMutableCopy func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateWithName func(unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileEstimateGamma func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileGetDisplayTransferFormulaFromVCGT func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncProfileGetURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileRemoveTag func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileSetHeader func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileSetTag func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileUninstall func(unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncRegisterDevice func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncTransformConvert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncTransformCopyProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncTransformGetTypeID func() unsafe.Pointer
	_ColorSyncTransformSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncUnregisterDevice func(unsafe.Pointer, unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CGDisplayCreateUUIDFromDisplayID, lib, "CGDisplayCreateUUIDFromDisplayID")
	tryRegister(&_CGDisplayGetDisplayIDFromUUID, lib, "CGDisplayGetDisplayIDFromUUID")
	tryRegister(&_ColorSyncAPIVersion, lib, "ColorSyncAPIVersion")
	tryRegister(&_ColorSyncCMMCopyCMMIdentifier, lib, "ColorSyncCMMCopyCMMIdentifier")
	tryRegister(&_ColorSyncCMMCopyLocalizedName, lib, "ColorSyncCMMCopyLocalizedName")
	tryRegister(&_ColorSyncCMMGetBundle, lib, "ColorSyncCMMGetBundle")
	tryRegister(&_ColorSyncCreateCodeFragment, lib, "ColorSyncCreateCodeFragment")
	tryRegister(&_ColorSyncDeviceCopyDeviceInfo, lib, "ColorSyncDeviceCopyDeviceInfo")
	tryRegister(&_ColorSyncDeviceSetCustomProfiles, lib, "ColorSyncDeviceSetCustomProfiles")
	tryRegister(&_ColorSyncIterateInstalledCMMs, lib, "ColorSyncIterateInstalledCMMs")
	tryRegister(&_ColorSyncIterateInstalledProfiles, lib, "ColorSyncIterateInstalledProfiles")
	tryRegister(&_ColorSyncIterateInstalledProfilesWithOptions, lib, "ColorSyncIterateInstalledProfilesWithOptions")
	tryRegister(&_ColorSyncProfileContainsTag, lib, "ColorSyncProfileContainsTag")
	tryRegister(&_ColorSyncProfileCopyData, lib, "ColorSyncProfileCopyData")
	tryRegister(&_ColorSyncProfileCopyDescriptionString, lib, "ColorSyncProfileCopyDescriptionString")
	tryRegister(&_ColorSyncProfileCopyHeader, lib, "ColorSyncProfileCopyHeader")
	tryRegister(&_ColorSyncProfileCreateWithDisplayID, lib, "ColorSyncProfileCreateWithDisplayID")
	tryRegister(&_ColorSyncProfileCreateWithURL, lib, "ColorSyncProfileCreateWithURL")
	tryRegister(&_ColorSyncProfileCreateWithURLAndOptions, lib, "ColorSyncProfileCreateWithURLAndOptions")
	tryRegister(&_ColorSyncProfileEstimateGammaWithDisplayID, lib, "ColorSyncProfileEstimateGammaWithDisplayID")
	tryRegister(&_ColorSyncProfileGetMD5, lib, "ColorSyncProfileGetMD5")
	tryRegister(&_ColorSyncProfileGetTagCount, lib, "ColorSyncProfileGetTagCount")
	tryRegister(&_ColorSyncProfileGetTypeID, lib, "ColorSyncProfileGetTypeID")
	tryRegister(&_ColorSyncProfileInstall, lib, "ColorSyncProfileInstall")
	tryRegister(&_ColorSyncProfileIsHLGBased, lib, "ColorSyncProfileIsHLGBased")
	tryRegister(&_ColorSyncProfileIsMatrixBased, lib, "ColorSyncProfileIsMatrixBased")
	tryRegister(&_ColorSyncProfileIsPQBased, lib, "ColorSyncProfileIsPQBased")
	tryRegister(&_ColorSyncProfileIsWideGamut, lib, "ColorSyncProfileIsWideGamut")
	tryRegister(&_ColorSyncProfileVerify, lib, "ColorSyncProfileVerify")
	tryRegister(&_ColorSyncTransformGetProfileSequence, lib, "ColorSyncTransformGetProfileSequence")
	tryRegister(&_ColorSyncCMMCreate, lib, "ColorSyncCMMCreate")
	tryRegister(&_ColorSyncCMMGetTypeID, lib, "ColorSyncCMMGetTypeID")
	tryRegister(&_ColorSyncIterateDeviceProfiles, lib, "ColorSyncIterateDeviceProfiles")
	tryRegister(&_ColorSyncProfileCopyTag, lib, "ColorSyncProfileCopyTag")
	tryRegister(&_ColorSyncProfileCopyTagSignatures, lib, "ColorSyncProfileCopyTagSignatures")
	tryRegister(&_ColorSyncProfileCreate, lib, "ColorSyncProfileCreate")
	tryRegister(&_ColorSyncProfileCreateDeviceProfile, lib, "ColorSyncProfileCreateDeviceProfile")
	tryRegister(&_ColorSyncProfileCreateDisplayTransferTablesFromVCGT, lib, "ColorSyncProfileCreateDisplayTransferTablesFromVCGT")
	tryRegister(&_ColorSyncProfileCreateLink, lib, "ColorSyncProfileCreateLink")
	tryRegister(&_ColorSyncProfileCreateMutable, lib, "ColorSyncProfileCreateMutable")
	tryRegister(&_ColorSyncProfileCreateMutableCopy, lib, "ColorSyncProfileCreateMutableCopy")
	tryRegister(&_ColorSyncProfileCreateWithName, lib, "ColorSyncProfileCreateWithName")
	tryRegister(&_ColorSyncProfileEstimateGamma, lib, "ColorSyncProfileEstimateGamma")
	tryRegister(&_ColorSyncProfileGetDisplayTransferFormulaFromVCGT, lib, "ColorSyncProfileGetDisplayTransferFormulaFromVCGT")
	tryRegister(&_ColorSyncProfileGetURL, lib, "ColorSyncProfileGetURL")
	tryRegister(&_ColorSyncProfileRemoveTag, lib, "ColorSyncProfileRemoveTag")
	tryRegister(&_ColorSyncProfileSetHeader, lib, "ColorSyncProfileSetHeader")
	tryRegister(&_ColorSyncProfileSetTag, lib, "ColorSyncProfileSetTag")
	tryRegister(&_ColorSyncProfileUninstall, lib, "ColorSyncProfileUninstall")
	tryRegister(&_ColorSyncRegisterDevice, lib, "ColorSyncRegisterDevice")
	tryRegister(&_ColorSyncTransformConvert, lib, "ColorSyncTransformConvert")
	tryRegister(&_ColorSyncTransformCopyProperty, lib, "ColorSyncTransformCopyProperty")
	tryRegister(&_ColorSyncTransformCreate, lib, "ColorSyncTransformCreate")
	tryRegister(&_ColorSyncTransformGetTypeID, lib, "ColorSyncTransformGetTypeID")
	tryRegister(&_ColorSyncTransformSetProperty, lib, "ColorSyncTransformSetProperty")
	tryRegister(&_ColorSyncUnregisterDevice, lib, "ColorSyncUnregisterDevice")
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



// CGDisplayCreateUUIDFromDisplayID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayCreateUUIDFromDisplayID(_:)
func CGDisplayCreateUUIDFromDisplayID(displayID unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayCreateUUIDFromDisplayID(displayID)
	}


// CGDisplayGetDisplayIDFromUUID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayGetDisplayIDFromUUID(_:)
func CGDisplayGetDisplayIDFromUUID(uuid unsafe.Pointer) unsafe.Pointer {
	return _CGDisplayGetDisplayIDFromUUID(uuid)
	}


// ColorSyncAPIVersion is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncAPIVersion()
func ColorSyncAPIVersion() unsafe.Pointer {
	return _ColorSyncAPIVersion()
	}


// ColorSyncCMMCopyCMMIdentifier is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyCMMIdentifier(_:)
func ColorSyncCMMCopyCMMIdentifier(p0 unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCMMCopyCMMIdentifier(p0)
	}


// ColorSyncCMMCopyLocalizedName is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyLocalizedName(_:)
func ColorSyncCMMCopyLocalizedName(p0 unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCMMCopyLocalizedName(p0)
	}


// ColorSyncCMMGetBundle is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetBundle(_:)
func ColorSyncCMMGetBundle(p0 unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCMMGetBundle(p0)
	}


// ColorSyncCreateCodeFragment is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCreateCodeFragment(_:_:)
func ColorSyncCreateCodeFragment(profileSequence unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCreateCodeFragment(profileSequence, options)
	}


// ColorSyncDeviceCopyDeviceInfo is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceCopyDeviceInfo(_:_:)
func ColorSyncDeviceCopyDeviceInfo(deviceClass unsafe.Pointer, devID unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncDeviceCopyDeviceInfo(deviceClass, devID)
	}


// ColorSyncDeviceSetCustomProfiles is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceSetCustomProfiles(_:_:_:)
func ColorSyncDeviceSetCustomProfiles(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, profileInfo unsafe.Pointer) bool {
	return _ColorSyncDeviceSetCustomProfiles(deviceClass, deviceID, profileInfo)
	}


// ColorSyncIterateInstalledCMMs is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledCMMs(_:_:)
func ColorSyncIterateInstalledCMMs(callBack unsafe.Pointer, userInfo unsafe.Pointer) {
	_ColorSyncIterateInstalledCMMs(callBack, userInfo)
	}


// ColorSyncIterateInstalledProfiles is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfiles(_:_:_:_:)
func ColorSyncIterateInstalledProfiles(callBack unsafe.Pointer, seed unsafe.Pointer, userInfo unsafe.Pointer, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfiles(callBack, seed, userInfo, error_)
	}


// ColorSyncIterateInstalledProfilesWithOptions is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfilesWithOptions(_:_:_:_:_:)
func ColorSyncIterateInstalledProfilesWithOptions(callBack unsafe.Pointer, seed unsafe.Pointer, userInfo unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfilesWithOptions(callBack, seed, userInfo, options, error_)
	}


// ColorSyncProfileContainsTag is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileContainsTag(_:_:)
func ColorSyncProfileContainsTag(prof unsafe.Pointer, signature unsafe.Pointer) bool {
	return _ColorSyncProfileContainsTag(prof, signature)
	}


// ColorSyncProfileCopyData is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyData(_:_:)
func ColorSyncProfileCopyData(prof unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyData(prof, error_)
	}


// ColorSyncProfileCopyDescriptionString is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyDescriptionString(_:)
func ColorSyncProfileCopyDescriptionString(prof unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyDescriptionString(prof)
	}


// ColorSyncProfileCopyHeader is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyHeader(_:)
func ColorSyncProfileCopyHeader(prof unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyHeader(prof)
	}


// ColorSyncProfileCreateWithDisplayID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithDisplayID(_:)
func ColorSyncProfileCreateWithDisplayID(displayID unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateWithDisplayID(displayID)
	}


// ColorSyncProfileCreateWithURL is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURL(_:_:)
func ColorSyncProfileCreateWithURL(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateWithURL(url, error_)
	}


// ColorSyncProfileCreateWithURLAndOptions is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURLAndOptions(_:_:_:)
func ColorSyncProfileCreateWithURLAndOptions(url unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateWithURLAndOptions(url, options, error_)
	}


// ColorSyncProfileEstimateGammaWithDisplayID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGammaWithDisplayID(_:_:)
func ColorSyncProfileEstimateGammaWithDisplayID(displayID unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileEstimateGammaWithDisplayID(displayID, error_)
	}


// ColorSyncProfileGetMD5 is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetMD5(_:)
func ColorSyncProfileGetMD5(prof unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileGetMD5(prof)
	}


// ColorSyncProfileGetTagCount is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTagCount(_:)
func ColorSyncProfileGetTagCount(p0 unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileGetTagCount(p0)
	}


// ColorSyncProfileGetTypeID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTypeID()
func ColorSyncProfileGetTypeID() unsafe.Pointer {
	return _ColorSyncProfileGetTypeID()
	}


// ColorSyncProfileInstall is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileInstall(_:_:_:_:)
func ColorSyncProfileInstall(profile unsafe.Pointer, domain unsafe.Pointer, subpath unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileInstall(profile, domain, subpath, error_)
	}


// ColorSyncProfileIsHLGBased is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsHLGBased(_:)
func ColorSyncProfileIsHLGBased(p0 unsafe.Pointer) bool {
	return _ColorSyncProfileIsHLGBased(p0)
	}


// ColorSyncProfileIsMatrixBased is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsMatrixBased(_:)
func ColorSyncProfileIsMatrixBased(p0 unsafe.Pointer) bool {
	return _ColorSyncProfileIsMatrixBased(p0)
	}


// ColorSyncProfileIsPQBased is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsPQBased(_:)
func ColorSyncProfileIsPQBased(p0 unsafe.Pointer) bool {
	return _ColorSyncProfileIsPQBased(p0)
	}


// ColorSyncProfileIsWideGamut is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsWideGamut(_:)
func ColorSyncProfileIsWideGamut(p0 unsafe.Pointer) bool {
	return _ColorSyncProfileIsWideGamut(p0)
	}


// ColorSyncProfileVerify is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileVerify(_:_:_:)
func ColorSyncProfileVerify(prof unsafe.Pointer, errors unsafe.Pointer, warnings unsafe.Pointer) bool {
	return _ColorSyncProfileVerify(prof, errors, warnings)
	}


// ColorSyncTransformGetProfileSequence is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetProfileSequence(_:)
func ColorSyncTransformGetProfileSequence(transform unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncTransformGetProfileSequence(transform)
	}


// ColorSyncCMMCreate is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCreate(_:)
func ColorSyncCMMCreate(cmmBundle unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCMMCreate(cmmBundle)
	}


// ColorSyncCMMGetTypeID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetTypeID()
func ColorSyncCMMGetTypeID() unsafe.Pointer {
	return _ColorSyncCMMGetTypeID()
	}


// ColorSyncIterateDeviceProfiles is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateDeviceProfiles(_:_:)
func ColorSyncIterateDeviceProfiles(callBack unsafe.Pointer, userInfo unsafe.Pointer) {
	_ColorSyncIterateDeviceProfiles(callBack, userInfo)
	}


// ColorSyncProfileCopyTag is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTag(_:_:)
func ColorSyncProfileCopyTag(prof unsafe.Pointer, signature unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyTag(prof, signature)
	}


// ColorSyncProfileCopyTagSignatures is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTagSignatures(_:)
func ColorSyncProfileCopyTagSignatures(prof unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyTagSignatures(prof)
	}


// ColorSyncProfileCreate is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreate(_:_:)
func ColorSyncProfileCreate(data unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreate(data, error_)
	}


// ColorSyncProfileCreateDeviceProfile is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDeviceProfile(_:_:_:)
func ColorSyncProfileCreateDeviceProfile(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, profileID unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateDeviceProfile(deviceClass, deviceID, profileID)
	}


// ColorSyncProfileCreateDisplayTransferTablesFromVCGT is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDisplayTransferTablesFromVCGT(_:_:)
func ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile unsafe.Pointer, nSamplesPerChannel unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile, nSamplesPerChannel)
	}


// ColorSyncProfileCreateLink is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateLink(_:_:)
func ColorSyncProfileCreateLink(profileInfo unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateLink(profileInfo, options)
	}


// ColorSyncProfileCreateMutable is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutable()
func ColorSyncProfileCreateMutable() unsafe.Pointer {
	return _ColorSyncProfileCreateMutable()
	}


// ColorSyncProfileCreateMutableCopy is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutableCopy(_:)
func ColorSyncProfileCreateMutableCopy(prof unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateMutableCopy(prof)
	}


// ColorSyncProfileCreateWithName is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithName(_:)
func ColorSyncProfileCreateWithName(name unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateWithName(name)
	}


// ColorSyncProfileEstimateGamma is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGamma(_:_:)
func ColorSyncProfileEstimateGamma(prof unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileEstimateGamma(prof, error_)
	}


// ColorSyncProfileGetDisplayTransferFormulaFromVCGT is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetDisplayTransferFormulaFromVCGT(_:_:_:_:_:_:_:_:_:_:)
func ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile unsafe.Pointer, redMin unsafe.Pointer, redMax unsafe.Pointer, redGamma unsafe.Pointer, greenMin unsafe.Pointer, greenMax unsafe.Pointer, greenGamma unsafe.Pointer, blueMin unsafe.Pointer, blueMax unsafe.Pointer, blueGamma unsafe.Pointer) bool {
	return _ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
	}


// ColorSyncProfileGetURL is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetURL(_:_:)
func ColorSyncProfileGetURL(prof unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileGetURL(prof, error_)
	}


// ColorSyncProfileRemoveTag is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileRemoveTag(_:_:)
func ColorSyncProfileRemoveTag(prof unsafe.Pointer, signature unsafe.Pointer) {
	_ColorSyncProfileRemoveTag(prof, signature)
	}


// ColorSyncProfileSetHeader is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetHeader(_:_:)
func ColorSyncProfileSetHeader(prof unsafe.Pointer, header unsafe.Pointer) {
	_ColorSyncProfileSetHeader(prof, header)
	}


// ColorSyncProfileSetTag is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetTag(_:_:_:)
func ColorSyncProfileSetTag(prof unsafe.Pointer, signature unsafe.Pointer, data unsafe.Pointer) {
	_ColorSyncProfileSetTag(prof, signature, data)
	}


// ColorSyncProfileUninstall is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileUninstall(_:_:)
func ColorSyncProfileUninstall(profile unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileUninstall(profile, error_)
	}


// ColorSyncRegisterDevice is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncRegisterDevice(_:_:_:)
func ColorSyncRegisterDevice(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, deviceInfo unsafe.Pointer) bool {
	return _ColorSyncRegisterDevice(deviceClass, deviceID, deviceInfo)
	}


// ColorSyncTransformConvert is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformConvert(_:_:_:_:_:_:_:_:_:_:_:_:)
func ColorSyncTransformConvert(transform unsafe.Pointer, width unsafe.Pointer, height unsafe.Pointer, dst unsafe.Pointer, dstDepth unsafe.Pointer, dstLayout unsafe.Pointer, dstBytesPerRow unsafe.Pointer, src unsafe.Pointer, srcDepth unsafe.Pointer, srcLayout unsafe.Pointer, srcBytesPerRow unsafe.Pointer, options unsafe.Pointer) bool {
	return _ColorSyncTransformConvert(transform, width, height, dst, dstDepth, dstLayout, dstBytesPerRow, src, srcDepth, srcLayout, srcBytesPerRow, options)
	}


// ColorSyncTransformCopyProperty is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCopyProperty(_:_:_:)
func ColorSyncTransformCopyProperty(transform unsafe.Pointer, key unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncTransformCopyProperty(transform, key, options)
	}


// ColorSyncTransformCreate is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCreate(_:_:)
func ColorSyncTransformCreate(profileSequence unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncTransformCreate(profileSequence, options)
	}


// ColorSyncTransformGetTypeID is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetTypeID()
func ColorSyncTransformGetTypeID() unsafe.Pointer {
	return _ColorSyncTransformGetTypeID()
	}


// ColorSyncTransformSetProperty is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformSetProperty(_:_:_:)
func ColorSyncTransformSetProperty(transform unsafe.Pointer, key unsafe.Pointer, property unsafe.Pointer) {
	_ColorSyncTransformSetProperty(transform, key, property)
	}


// ColorSyncUnregisterDevice is a ColorSync function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncUnregisterDevice(_:_:)
func ColorSyncUnregisterDevice(deviceClass unsafe.Pointer, deviceID unsafe.Pointer) bool {
	return _ColorSyncUnregisterDevice(deviceClass, deviceID)
	}




