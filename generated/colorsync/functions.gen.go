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
	_ColorSyncAPIVersion func() uint32
	_ColorSyncCreateCodeFragment func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncIterateInstalledProfilesWithOptions func(ColorSyncProfileIterateCallback, []uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_ColorSyncProfileCopyData func(ColorSyncProfileRef, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateWithURLAndOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileGetTagCount func(ColorSyncProfileRef) uintptr
	_ColorSyncProfileIsHLGBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsMatrixBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsPQBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsWideGamut func(ColorSyncProfileRef) bool
	_ColorSyncTransformGetProfileSequence func(ColorSyncTransformRef) unsafe.Pointer
	_CGDisplayCreateUUIDFromDisplayID func(uint32) unsafe.Pointer
	_CGDisplayGetDisplayIDFromUUID func(unsafe.Pointer) uint32
	_ColorSyncCMMCopyCMMIdentifier func(ColorSyncCMMRef) unsafe.Pointer
	_ColorSyncCMMCopyLocalizedName func(ColorSyncCMMRef) unsafe.Pointer
	_ColorSyncCMMCreate func(unsafe.Pointer) ColorSyncCMMRef
	_ColorSyncCMMGetBundle func(ColorSyncCMMRef) unsafe.Pointer
	_ColorSyncCMMGetTypeID func() unsafe.Pointer
	_ColorSyncDeviceCopyDeviceInfo func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncDeviceSetCustomProfiles func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncIterateDeviceProfiles func(ColorSyncDeviceProfileIterateCallback, unsafe.Pointer)
	_ColorSyncIterateInstalledCMMs func(ColorSyncCMMIterateCallback, unsafe.Pointer)
	_ColorSyncIterateInstalledProfiles func(ColorSyncProfileIterateCallback, []uint32, unsafe.Pointer, unsafe.Pointer)
	_ColorSyncProfileContainsTag func(ColorSyncProfileRef, unsafe.Pointer) bool
	_ColorSyncProfileCopyDescriptionString func(ColorSyncProfileRef) unsafe.Pointer
	_ColorSyncProfileCopyHeader func(ColorSyncProfileRef) unsafe.Pointer
	_ColorSyncProfileCopyTag func(ColorSyncProfileRef, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCopyTagSignatures func(ColorSyncProfileRef) unsafe.Pointer
	_ColorSyncProfileCreate func(unsafe.Pointer, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileCreateDeviceProfile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileCreateDisplayTransferTablesFromVCGT func(ColorSyncProfileRef, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileCreateLink func(unsafe.Pointer, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileCreateMutable func() ColorSyncMutableProfileRef
	_ColorSyncProfileCreateMutableCopy func(ColorSyncProfileRef) ColorSyncMutableProfileRef
	_ColorSyncProfileCreateWithDisplayID func(uint32) ColorSyncProfileRef
	_ColorSyncProfileCreateWithName func(unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileCreateWithURL func(unsafe.Pointer, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileEstimateGamma func(ColorSyncProfileRef, unsafe.Pointer) float32
	_ColorSyncProfileEstimateGammaWithDisplayID func(unsafe.Pointer, unsafe.Pointer) float32
	_ColorSyncProfileGetDisplayTransferFormulaFromVCGT func(ColorSyncProfileRef, []float32, []float32, []float32, []float32, []float32, []float32, []float32, []float32, []float32) bool
	_ColorSyncProfileGetMD5 func(ColorSyncProfileRef) unsafe.Pointer
	_ColorSyncProfileGetTypeID func() unsafe.Pointer
	_ColorSyncProfileGetURL func(ColorSyncProfileRef, unsafe.Pointer) unsafe.Pointer
	_ColorSyncProfileInstall func(ColorSyncProfileRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncProfileRemoveTag func(ColorSyncMutableProfileRef, unsafe.Pointer)
	_ColorSyncProfileSetHeader func(ColorSyncMutableProfileRef, unsafe.Pointer)
	_ColorSyncProfileSetTag func(ColorSyncMutableProfileRef, unsafe.Pointer, unsafe.Pointer)
	_ColorSyncProfileUninstall func(ColorSyncProfileRef, unsafe.Pointer) bool
	_ColorSyncProfileVerify func(ColorSyncProfileRef, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncRegisterDevice func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncTransformConvert func(ColorSyncTransformRef, uintptr, uintptr, unsafe.Pointer, unsafe.Pointer, ColorSyncDataLayout, uintptr, unsafe.Pointer, unsafe.Pointer, ColorSyncDataLayout, uintptr, unsafe.Pointer) bool
	_ColorSyncTransformCopyProperty func(ColorSyncTransformRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ColorSyncTransformCreate func(unsafe.Pointer, unsafe.Pointer) ColorSyncTransformRef
	_ColorSyncTransformGetTypeID func() unsafe.Pointer
	_ColorSyncTransformSetProperty func(ColorSyncTransformRef, unsafe.Pointer, unsafe.Pointer)
	_ColorSyncUnregisterDevice func(unsafe.Pointer, unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_ColorSyncAPIVersion, lib, "ColorSyncAPIVersion")
	tryRegister(&_ColorSyncCreateCodeFragment, lib, "ColorSyncCreateCodeFragment")
	tryRegister(&_ColorSyncIterateInstalledProfilesWithOptions, lib, "ColorSyncIterateInstalledProfilesWithOptions")
	tryRegister(&_ColorSyncProfileCopyData, lib, "ColorSyncProfileCopyData")
	tryRegister(&_ColorSyncProfileCreateWithURLAndOptions, lib, "ColorSyncProfileCreateWithURLAndOptions")
	tryRegister(&_ColorSyncProfileGetTagCount, lib, "ColorSyncProfileGetTagCount")
	tryRegister(&_ColorSyncProfileIsHLGBased, lib, "ColorSyncProfileIsHLGBased")
	tryRegister(&_ColorSyncProfileIsMatrixBased, lib, "ColorSyncProfileIsMatrixBased")
	tryRegister(&_ColorSyncProfileIsPQBased, lib, "ColorSyncProfileIsPQBased")
	tryRegister(&_ColorSyncProfileIsWideGamut, lib, "ColorSyncProfileIsWideGamut")
	tryRegister(&_ColorSyncTransformGetProfileSequence, lib, "ColorSyncTransformGetProfileSequence")
	tryRegister(&_CGDisplayCreateUUIDFromDisplayID, lib, "CGDisplayCreateUUIDFromDisplayID")
	tryRegister(&_CGDisplayGetDisplayIDFromUUID, lib, "CGDisplayGetDisplayIDFromUUID")
	tryRegister(&_ColorSyncCMMCopyCMMIdentifier, lib, "ColorSyncCMMCopyCMMIdentifier")
	tryRegister(&_ColorSyncCMMCopyLocalizedName, lib, "ColorSyncCMMCopyLocalizedName")
	tryRegister(&_ColorSyncCMMCreate, lib, "ColorSyncCMMCreate")
	tryRegister(&_ColorSyncCMMGetBundle, lib, "ColorSyncCMMGetBundle")
	tryRegister(&_ColorSyncCMMGetTypeID, lib, "ColorSyncCMMGetTypeID")
	tryRegister(&_ColorSyncDeviceCopyDeviceInfo, lib, "ColorSyncDeviceCopyDeviceInfo")
	tryRegister(&_ColorSyncDeviceSetCustomProfiles, lib, "ColorSyncDeviceSetCustomProfiles")
	tryRegister(&_ColorSyncIterateDeviceProfiles, lib, "ColorSyncIterateDeviceProfiles")
	tryRegister(&_ColorSyncIterateInstalledCMMs, lib, "ColorSyncIterateInstalledCMMs")
	tryRegister(&_ColorSyncIterateInstalledProfiles, lib, "ColorSyncIterateInstalledProfiles")
	tryRegister(&_ColorSyncProfileContainsTag, lib, "ColorSyncProfileContainsTag")
	tryRegister(&_ColorSyncProfileCopyDescriptionString, lib, "ColorSyncProfileCopyDescriptionString")
	tryRegister(&_ColorSyncProfileCopyHeader, lib, "ColorSyncProfileCopyHeader")
	tryRegister(&_ColorSyncProfileCopyTag, lib, "ColorSyncProfileCopyTag")
	tryRegister(&_ColorSyncProfileCopyTagSignatures, lib, "ColorSyncProfileCopyTagSignatures")
	tryRegister(&_ColorSyncProfileCreate, lib, "ColorSyncProfileCreate")
	tryRegister(&_ColorSyncProfileCreateDeviceProfile, lib, "ColorSyncProfileCreateDeviceProfile")
	tryRegister(&_ColorSyncProfileCreateDisplayTransferTablesFromVCGT, lib, "ColorSyncProfileCreateDisplayTransferTablesFromVCGT")
	tryRegister(&_ColorSyncProfileCreateLink, lib, "ColorSyncProfileCreateLink")
	tryRegister(&_ColorSyncProfileCreateMutable, lib, "ColorSyncProfileCreateMutable")
	tryRegister(&_ColorSyncProfileCreateMutableCopy, lib, "ColorSyncProfileCreateMutableCopy")
	tryRegister(&_ColorSyncProfileCreateWithDisplayID, lib, "ColorSyncProfileCreateWithDisplayID")
	tryRegister(&_ColorSyncProfileCreateWithName, lib, "ColorSyncProfileCreateWithName")
	tryRegister(&_ColorSyncProfileCreateWithURL, lib, "ColorSyncProfileCreateWithURL")
	tryRegister(&_ColorSyncProfileEstimateGamma, lib, "ColorSyncProfileEstimateGamma")
	tryRegister(&_ColorSyncProfileEstimateGammaWithDisplayID, lib, "ColorSyncProfileEstimateGammaWithDisplayID")
	tryRegister(&_ColorSyncProfileGetDisplayTransferFormulaFromVCGT, lib, "ColorSyncProfileGetDisplayTransferFormulaFromVCGT")
	tryRegister(&_ColorSyncProfileGetMD5, lib, "ColorSyncProfileGetMD5")
	tryRegister(&_ColorSyncProfileGetTypeID, lib, "ColorSyncProfileGetTypeID")
	tryRegister(&_ColorSyncProfileGetURL, lib, "ColorSyncProfileGetURL")
	tryRegister(&_ColorSyncProfileInstall, lib, "ColorSyncProfileInstall")
	tryRegister(&_ColorSyncProfileRemoveTag, lib, "ColorSyncProfileRemoveTag")
	tryRegister(&_ColorSyncProfileSetHeader, lib, "ColorSyncProfileSetHeader")
	tryRegister(&_ColorSyncProfileSetTag, lib, "ColorSyncProfileSetTag")
	tryRegister(&_ColorSyncProfileUninstall, lib, "ColorSyncProfileUninstall")
	tryRegister(&_ColorSyncProfileVerify, lib, "ColorSyncProfileVerify")
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



// ColorSyncAPIVersion is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncAPIVersion()
func ColorSyncAPIVersion() uint32 {
	return _ColorSyncAPIVersion()
}

// ColorSyncCreateCodeFragment is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCreateCodeFragment(_:_:)
func ColorSyncCreateCodeFragment(profileSequence unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncCreateCodeFragment(profileSequence, options)
}

// ColorSyncIterateInstalledProfilesWithOptions is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfilesWithOptions(_:_:_:_:_:)
func ColorSyncIterateInstalledProfilesWithOptions(callBack ColorSyncProfileIterateCallback, seed []uint32, userInfo unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfilesWithOptions(callBack, seed, userInfo, options, error_)
}

// ColorSyncProfileCopyData is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyData(_:_:)
func ColorSyncProfileCopyData(prof ColorSyncProfileRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyData(prof, error_)
}

// ColorSyncProfileCreateWithURLAndOptions is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURLAndOptions(_:_:_:)
func ColorSyncProfileCreateWithURLAndOptions(url unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithURLAndOptions(url, options, error_)
}

// ColorSyncProfileGetTagCount is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTagCount(_:)
func ColorSyncProfileGetTagCount(p0 ColorSyncProfileRef) uintptr {
	return _ColorSyncProfileGetTagCount(p0)
}

// ColorSyncProfileIsHLGBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsHLGBased(_:)
func ColorSyncProfileIsHLGBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsHLGBased(p0)
}

// ColorSyncProfileIsMatrixBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsMatrixBased(_:)
func ColorSyncProfileIsMatrixBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsMatrixBased(p0)
}

// ColorSyncProfileIsPQBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsPQBased(_:)
func ColorSyncProfileIsPQBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsPQBased(p0)
}

// ColorSyncProfileIsWideGamut is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsWideGamut(_:)
func ColorSyncProfileIsWideGamut(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsWideGamut(p0)
}

// ColorSyncTransformGetProfileSequence is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetProfileSequence(_:)
func ColorSyncTransformGetProfileSequence(transform ColorSyncTransformRef) unsafe.Pointer {
	return _ColorSyncTransformGetProfileSequence(transform)
}

// CGDisplayCreateUUIDFromDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayCreateUUIDFromDisplayID(_:)
func CGDisplayCreateUUIDFromDisplayID(displayID uint32) unsafe.Pointer {
	return _CGDisplayCreateUUIDFromDisplayID(displayID)
}

// CGDisplayGetDisplayIDFromUUID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayGetDisplayIDFromUUID(_:)
func CGDisplayGetDisplayIDFromUUID(uuid unsafe.Pointer) uint32 {
	return _CGDisplayGetDisplayIDFromUUID(uuid)
}

// ColorSyncCMMCopyCMMIdentifier is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyCMMIdentifier(_:)
func ColorSyncCMMCopyCMMIdentifier(p0 ColorSyncCMMRef) unsafe.Pointer {
	return _ColorSyncCMMCopyCMMIdentifier(p0)
}

// ColorSyncCMMCopyLocalizedName is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyLocalizedName(_:)
func ColorSyncCMMCopyLocalizedName(p0 ColorSyncCMMRef) unsafe.Pointer {
	return _ColorSyncCMMCopyLocalizedName(p0)
}

// ColorSyncCMMCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCreate(_:)
func ColorSyncCMMCreate(cmmBundle unsafe.Pointer) ColorSyncCMMRef {
	return _ColorSyncCMMCreate(cmmBundle)
}

// ColorSyncCMMGetBundle is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetBundle(_:)
func ColorSyncCMMGetBundle(p0 ColorSyncCMMRef) unsafe.Pointer {
	return _ColorSyncCMMGetBundle(p0)
}

// ColorSyncCMMGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetTypeID()
func ColorSyncCMMGetTypeID() unsafe.Pointer {
	return _ColorSyncCMMGetTypeID()
}

// ColorSyncDeviceCopyDeviceInfo is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceCopyDeviceInfo(_:_:)
func ColorSyncDeviceCopyDeviceInfo(deviceClass unsafe.Pointer, devID unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncDeviceCopyDeviceInfo(deviceClass, devID)
}

// ColorSyncDeviceSetCustomProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceSetCustomProfiles(_:_:_:)
func ColorSyncDeviceSetCustomProfiles(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, profileInfo unsafe.Pointer) bool {
	return _ColorSyncDeviceSetCustomProfiles(deviceClass, deviceID, profileInfo)
}

// ColorSyncIterateDeviceProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateDeviceProfiles(_:_:)
func ColorSyncIterateDeviceProfiles(callBack ColorSyncDeviceProfileIterateCallback, userInfo unsafe.Pointer) {
	_ColorSyncIterateDeviceProfiles(callBack, userInfo)
}

// ColorSyncIterateInstalledCMMs is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledCMMs(_:_:)
func ColorSyncIterateInstalledCMMs(callBack ColorSyncCMMIterateCallback, userInfo unsafe.Pointer) {
	_ColorSyncIterateInstalledCMMs(callBack, userInfo)
}

// ColorSyncIterateInstalledProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfiles(_:_:_:_:)
func ColorSyncIterateInstalledProfiles(callBack ColorSyncProfileIterateCallback, seed []uint32, userInfo unsafe.Pointer, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfiles(callBack, seed, userInfo, error_)
}

// ColorSyncProfileContainsTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileContainsTag(_:_:)
func ColorSyncProfileContainsTag(prof ColorSyncProfileRef, signature unsafe.Pointer) bool {
	return _ColorSyncProfileContainsTag(prof, signature)
}

// ColorSyncProfileCopyDescriptionString is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyDescriptionString(_:)
func ColorSyncProfileCopyDescriptionString(prof ColorSyncProfileRef) unsafe.Pointer {
	return _ColorSyncProfileCopyDescriptionString(prof)
}

// ColorSyncProfileCopyHeader is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyHeader(_:)
func ColorSyncProfileCopyHeader(prof ColorSyncProfileRef) unsafe.Pointer {
	return _ColorSyncProfileCopyHeader(prof)
}

// ColorSyncProfileCopyTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTag(_:_:)
func ColorSyncProfileCopyTag(prof ColorSyncProfileRef, signature unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCopyTag(prof, signature)
}

// ColorSyncProfileCopyTagSignatures is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTagSignatures(_:)
func ColorSyncProfileCopyTagSignatures(prof ColorSyncProfileRef) unsafe.Pointer {
	return _ColorSyncProfileCopyTagSignatures(prof)
}

// ColorSyncProfileCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreate(_:_:)
func ColorSyncProfileCreate(data unsafe.Pointer, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreate(data, error_)
}

// ColorSyncProfileCreateDeviceProfile is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDeviceProfile(_:_:_:)
func ColorSyncProfileCreateDeviceProfile(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, profileID unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateDeviceProfile(deviceClass, deviceID, profileID)
}

// ColorSyncProfileCreateDisplayTransferTablesFromVCGT is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDisplayTransferTablesFromVCGT(_:_:)
func ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile ColorSyncProfileRef, nSamplesPerChannel unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile, nSamplesPerChannel)
}

// ColorSyncProfileCreateLink is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateLink(_:_:)
func ColorSyncProfileCreateLink(profileInfo unsafe.Pointer, options unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateLink(profileInfo, options)
}

// ColorSyncProfileCreateMutable is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutable()
func ColorSyncProfileCreateMutable() ColorSyncMutableProfileRef {
	return _ColorSyncProfileCreateMutable()
}

// ColorSyncProfileCreateMutableCopy is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutableCopy(_:)
func ColorSyncProfileCreateMutableCopy(prof ColorSyncProfileRef) ColorSyncMutableProfileRef {
	return _ColorSyncProfileCreateMutableCopy(prof)
}

// ColorSyncProfileCreateWithDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithDisplayID(_:)
func ColorSyncProfileCreateWithDisplayID(displayID uint32) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithDisplayID(displayID)
}

// ColorSyncProfileCreateWithName is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithName(_:)
func ColorSyncProfileCreateWithName(name unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithName(name)
}

// ColorSyncProfileCreateWithURL is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURL(_:_:)
func ColorSyncProfileCreateWithURL(url unsafe.Pointer, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithURL(url, error_)
}

// ColorSyncProfileEstimateGamma is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGamma(_:_:)
func ColorSyncProfileEstimateGamma(prof ColorSyncProfileRef, error_ unsafe.Pointer) float32 {
	return _ColorSyncProfileEstimateGamma(prof, error_)
}

// ColorSyncProfileEstimateGammaWithDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGammaWithDisplayID(_:_:)
func ColorSyncProfileEstimateGammaWithDisplayID(displayID unsafe.Pointer, error_ unsafe.Pointer) float32 {
	return _ColorSyncProfileEstimateGammaWithDisplayID(displayID, error_)
}

// ColorSyncProfileGetDisplayTransferFormulaFromVCGT is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetDisplayTransferFormulaFromVCGT(_:_:_:_:_:_:_:_:_:_:)
func ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile ColorSyncProfileRef, redMin []float32, redMax []float32, redGamma []float32, greenMin []float32, greenMax []float32, greenGamma []float32, blueMin []float32, blueMax []float32, blueGamma []float32) bool {
	return _ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}

// ColorSyncProfileGetMD5 is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetMD5(_:)
func ColorSyncProfileGetMD5(prof ColorSyncProfileRef) unsafe.Pointer {
	return _ColorSyncProfileGetMD5(prof)
}

// ColorSyncProfileGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTypeID()
func ColorSyncProfileGetTypeID() unsafe.Pointer {
	return _ColorSyncProfileGetTypeID()
}

// ColorSyncProfileGetURL is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetURL(_:_:)
func ColorSyncProfileGetURL(prof ColorSyncProfileRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncProfileGetURL(prof, error_)
}

// ColorSyncProfileInstall is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileInstall(_:_:_:_:)
func ColorSyncProfileInstall(profile ColorSyncProfileRef, domain unsafe.Pointer, subpath unsafe.Pointer, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileInstall(profile, domain, subpath, error_)
}

// ColorSyncProfileRemoveTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileRemoveTag(_:_:)
func ColorSyncProfileRemoveTag(prof ColorSyncMutableProfileRef, signature unsafe.Pointer) {
	_ColorSyncProfileRemoveTag(prof, signature)
}

// ColorSyncProfileSetHeader is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetHeader(_:_:)
func ColorSyncProfileSetHeader(prof ColorSyncMutableProfileRef, header unsafe.Pointer) {
	_ColorSyncProfileSetHeader(prof, header)
}

// ColorSyncProfileSetTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetTag(_:_:_:)
func ColorSyncProfileSetTag(prof ColorSyncMutableProfileRef, signature unsafe.Pointer, data unsafe.Pointer) {
	_ColorSyncProfileSetTag(prof, signature, data)
}

// ColorSyncProfileUninstall is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileUninstall(_:_:)
func ColorSyncProfileUninstall(profile ColorSyncProfileRef, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileUninstall(profile, error_)
}

// ColorSyncProfileVerify is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileVerify(_:_:_:)
func ColorSyncProfileVerify(prof ColorSyncProfileRef, errors unsafe.Pointer, warnings unsafe.Pointer) bool {
	return _ColorSyncProfileVerify(prof, errors, warnings)
}

// ColorSyncRegisterDevice is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncRegisterDevice(_:_:_:)
func ColorSyncRegisterDevice(deviceClass unsafe.Pointer, deviceID unsafe.Pointer, deviceInfo unsafe.Pointer) bool {
	return _ColorSyncRegisterDevice(deviceClass, deviceID, deviceInfo)
}

// ColorSyncTransformConvert is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformConvert(_:_:_:_:_:_:_:_:_:_:_:_:)
func ColorSyncTransformConvert(transform ColorSyncTransformRef, width uintptr, height uintptr, dst unsafe.Pointer, dstDepth unsafe.Pointer, dstLayout ColorSyncDataLayout, dstBytesPerRow uintptr, src unsafe.Pointer, srcDepth unsafe.Pointer, srcLayout ColorSyncDataLayout, srcBytesPerRow uintptr, options unsafe.Pointer) bool {
	return _ColorSyncTransformConvert(transform, width, height, dst, dstDepth, dstLayout, dstBytesPerRow, src, srcDepth, srcLayout, srcBytesPerRow, options)
}

// ColorSyncTransformCopyProperty is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCopyProperty(_:_:_:)
func ColorSyncTransformCopyProperty(transform ColorSyncTransformRef, key unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _ColorSyncTransformCopyProperty(transform, key, options)
}

// ColorSyncTransformCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCreate(_:_:)
func ColorSyncTransformCreate(profileSequence unsafe.Pointer, options unsafe.Pointer) ColorSyncTransformRef {
	return _ColorSyncTransformCreate(profileSequence, options)
}

// ColorSyncTransformGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetTypeID()
func ColorSyncTransformGetTypeID() unsafe.Pointer {
	return _ColorSyncTransformGetTypeID()
}

// ColorSyncTransformSetProperty is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformSetProperty(_:_:_:)
func ColorSyncTransformSetProperty(transform ColorSyncTransformRef, key unsafe.Pointer, property unsafe.Pointer) {
	_ColorSyncTransformSetProperty(transform, key, property)
}

// ColorSyncUnregisterDevice is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncUnregisterDevice(_:_:)
func ColorSyncUnregisterDevice(deviceClass unsafe.Pointer, deviceID unsafe.Pointer) bool {
	return _ColorSyncUnregisterDevice(deviceClass, deviceID)
}



