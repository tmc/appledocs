// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync

/* debug [functions.gen.go]: Generating 56 functions for ColorSync */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// ColorSync Functions (56 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_ColorSyncProfileCreateWithURLAndOptions func(URLRef, DictionaryRef, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileGetTagCount func(ColorSyncProfileRef) uintptr
	_CGDisplayCreateUUIDFromDisplayID func(uint32) UUIDRef
	_CGDisplayGetDisplayIDFromUUID func(UUIDRef) uint32
	_ColorSyncAPIVersion func() uint32
	_ColorSyncCMMCopyCMMIdentifier func(ColorSyncCMMRef) StringRef
	_ColorSyncCMMCopyLocalizedName func(ColorSyncCMMRef) StringRef
	_ColorSyncCMMCreate func(BundleRef) ColorSyncCMMRef
	_ColorSyncCMMGetBundle func(ColorSyncCMMRef) BundleRef
	_ColorSyncCMMGetTypeID func() TypeID
	_ColorSyncCreateCodeFragment func(ArrayRef, DictionaryRef) TypeRef
	_ColorSyncDeviceCopyDeviceInfo func(StringRef, UUIDRef) DictionaryRef
	_ColorSyncDeviceSetCustomProfiles func(StringRef, UUIDRef, DictionaryRef) bool
	_ColorSyncIterateDeviceProfiles func(ColorSyncDeviceProfileIterateCallback, unsafe.Pointer)
	_ColorSyncIterateInstalledCMMs func(ColorSyncCMMIterateCallback, unsafe.Pointer)
	_ColorSyncIterateInstalledProfiles func(ColorSyncProfileIterateCallback, []uint32, unsafe.Pointer, unsafe.Pointer)
	_ColorSyncIterateInstalledProfilesWithOptions func(ColorSyncProfileIterateCallback, []uint32, unsafe.Pointer, DictionaryRef, unsafe.Pointer)
	_ColorSyncProfileContainsTag func(ColorSyncProfileRef, StringRef) bool
	_ColorSyncProfileCopyData func(ColorSyncProfileRef, unsafe.Pointer) DataRef
	_ColorSyncProfileCopyDescriptionString func(ColorSyncProfileRef) StringRef
	_ColorSyncProfileCopyHeader func(ColorSyncProfileRef) DataRef
	_ColorSyncProfileCopyTag func(ColorSyncProfileRef, StringRef) DataRef
	_ColorSyncProfileCopyTagSignatures func(ColorSyncProfileRef) ArrayRef
	_ColorSyncProfileCreate func(DataRef, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileCreateDeviceProfile func(StringRef, UUIDRef, TypeRef) ColorSyncProfileRef
	_ColorSyncProfileCreateDisplayTransferTablesFromVCGT func(ColorSyncProfileRef, unsafe.Pointer) DataRef
	_ColorSyncProfileCreateLink func(ArrayRef, DictionaryRef) ColorSyncProfileRef
	_ColorSyncProfileCreateMutable func() ColorSyncMutableProfileRef
	_ColorSyncProfileCreateMutableCopy func(ColorSyncProfileRef) ColorSyncMutableProfileRef
	_ColorSyncProfileCreateWithDisplayID func(uint32) ColorSyncProfileRef
	_ColorSyncProfileCreateWithName func(StringRef) ColorSyncProfileRef
	_ColorSyncProfileCreateWithURL func(URLRef, unsafe.Pointer) ColorSyncProfileRef
	_ColorSyncProfileEstimateGamma func(ColorSyncProfileRef, unsafe.Pointer) float32
	_ColorSyncProfileEstimateGammaWithDisplayID func(unsafe.Pointer, unsafe.Pointer) float32
	_ColorSyncProfileGetDisplayTransferFormulaFromVCGT func(ColorSyncProfileRef, []float32, []float32, []float32, []float32, []float32, []float32, []float32, []float32, []float32) bool
	_ColorSyncProfileGetMD5 func(ColorSyncProfileRef) ColorSyncMD5
	_ColorSyncProfileGetTypeID func() TypeID
	_ColorSyncProfileGetURL func(ColorSyncProfileRef, unsafe.Pointer) URLRef
	_ColorSyncProfileInstall func(ColorSyncProfileRef, StringRef, StringRef, unsafe.Pointer) bool
	_ColorSyncProfileIsHLGBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsMatrixBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsPQBased func(ColorSyncProfileRef) bool
	_ColorSyncProfileIsWideGamut func(ColorSyncProfileRef) bool
	_ColorSyncProfileRemoveTag func(ColorSyncMutableProfileRef, StringRef)
	_ColorSyncProfileSetHeader func(ColorSyncMutableProfileRef, DataRef)
	_ColorSyncProfileSetTag func(ColorSyncMutableProfileRef, StringRef, DataRef)
	_ColorSyncProfileUninstall func(ColorSyncProfileRef, unsafe.Pointer) bool
	_ColorSyncProfileVerify func(ColorSyncProfileRef, unsafe.Pointer, unsafe.Pointer) bool
	_ColorSyncRegisterDevice func(StringRef, UUIDRef, DictionaryRef) bool
	_ColorSyncTransformConvert func(ColorSyncTransformRef, uintptr, uintptr, unsafe.Pointer, ColorSyncDataDepth, ColorSyncDataLayout, uintptr, unsafe.Pointer, ColorSyncDataDepth, ColorSyncDataLayout, uintptr, DictionaryRef) bool
	_ColorSyncTransformCopyProperty func(ColorSyncTransformRef, TypeRef, DictionaryRef) TypeRef
	_ColorSyncTransformCreate func(ArrayRef, DictionaryRef) ColorSyncTransformRef
	_ColorSyncTransformGetProfileSequence func(ColorSyncTransformRef) ArrayRef
	_ColorSyncTransformGetTypeID func() TypeID
	_ColorSyncTransformSetProperty func(ColorSyncTransformRef, TypeRef, TypeRef)
	_ColorSyncUnregisterDevice func(StringRef, UUIDRef) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_ColorSyncProfileCreateWithURLAndOptions, lib, "ColorSyncProfileCreateWithURLAndOptions")
	tryRegister(&_ColorSyncProfileGetTagCount, lib, "ColorSyncProfileGetTagCount")
	tryRegister(&_CGDisplayCreateUUIDFromDisplayID, lib, "CGDisplayCreateUUIDFromDisplayID")
	tryRegister(&_CGDisplayGetDisplayIDFromUUID, lib, "CGDisplayGetDisplayIDFromUUID")
	tryRegister(&_ColorSyncAPIVersion, lib, "ColorSyncAPIVersion")
	tryRegister(&_ColorSyncCMMCopyCMMIdentifier, lib, "ColorSyncCMMCopyCMMIdentifier")
	tryRegister(&_ColorSyncCMMCopyLocalizedName, lib, "ColorSyncCMMCopyLocalizedName")
	tryRegister(&_ColorSyncCMMCreate, lib, "ColorSyncCMMCreate")
	tryRegister(&_ColorSyncCMMGetBundle, lib, "ColorSyncCMMGetBundle")
	tryRegister(&_ColorSyncCMMGetTypeID, lib, "ColorSyncCMMGetTypeID")
	tryRegister(&_ColorSyncCreateCodeFragment, lib, "ColorSyncCreateCodeFragment")
	tryRegister(&_ColorSyncDeviceCopyDeviceInfo, lib, "ColorSyncDeviceCopyDeviceInfo")
	tryRegister(&_ColorSyncDeviceSetCustomProfiles, lib, "ColorSyncDeviceSetCustomProfiles")
	tryRegister(&_ColorSyncIterateDeviceProfiles, lib, "ColorSyncIterateDeviceProfiles")
	tryRegister(&_ColorSyncIterateInstalledCMMs, lib, "ColorSyncIterateInstalledCMMs")
	tryRegister(&_ColorSyncIterateInstalledProfiles, lib, "ColorSyncIterateInstalledProfiles")
	tryRegister(&_ColorSyncIterateInstalledProfilesWithOptions, lib, "ColorSyncIterateInstalledProfilesWithOptions")
	tryRegister(&_ColorSyncProfileContainsTag, lib, "ColorSyncProfileContainsTag")
	tryRegister(&_ColorSyncProfileCopyData, lib, "ColorSyncProfileCopyData")
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
	tryRegister(&_ColorSyncProfileIsHLGBased, lib, "ColorSyncProfileIsHLGBased")
	tryRegister(&_ColorSyncProfileIsMatrixBased, lib, "ColorSyncProfileIsMatrixBased")
	tryRegister(&_ColorSyncProfileIsPQBased, lib, "ColorSyncProfileIsPQBased")
	tryRegister(&_ColorSyncProfileIsWideGamut, lib, "ColorSyncProfileIsWideGamut")
	tryRegister(&_ColorSyncProfileRemoveTag, lib, "ColorSyncProfileRemoveTag")
	tryRegister(&_ColorSyncProfileSetHeader, lib, "ColorSyncProfileSetHeader")
	tryRegister(&_ColorSyncProfileSetTag, lib, "ColorSyncProfileSetTag")
	tryRegister(&_ColorSyncProfileUninstall, lib, "ColorSyncProfileUninstall")
	tryRegister(&_ColorSyncProfileVerify, lib, "ColorSyncProfileVerify")
	tryRegister(&_ColorSyncRegisterDevice, lib, "ColorSyncRegisterDevice")
	tryRegister(&_ColorSyncTransformConvert, lib, "ColorSyncTransformConvert")
	tryRegister(&_ColorSyncTransformCopyProperty, lib, "ColorSyncTransformCopyProperty")
	tryRegister(&_ColorSyncTransformCreate, lib, "ColorSyncTransformCreate")
	tryRegister(&_ColorSyncTransformGetProfileSequence, lib, "ColorSyncTransformGetProfileSequence")
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



// ColorSyncProfileCreateWithURLAndOptions is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURLAndOptions(_:_:_:)
func ColorSyncProfileCreateWithURLAndOptions(url URLRef, options DictionaryRef, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithURLAndOptions(url, options, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateWithURLAndOptions */

// ColorSyncProfileGetTagCount is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTagCount(_:)
func ColorSyncProfileGetTagCount(p0 ColorSyncProfileRef) uintptr {
	return _ColorSyncProfileGetTagCount(p0)
}/* debug [functions.gen.go/function]: ColorSyncProfileGetTagCount */

// CGDisplayCreateUUIDFromDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayCreateUUIDFromDisplayID(_:)
func CGDisplayCreateUUIDFromDisplayID(displayID uint32) UUIDRef {
	return _CGDisplayCreateUUIDFromDisplayID(displayID)
}/* debug [functions.gen.go/function]: CGDisplayCreateUUIDFromDisplayID */

// CGDisplayGetDisplayIDFromUUID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CGDisplayGetDisplayIDFromUUID(_:)
func CGDisplayGetDisplayIDFromUUID(uuid UUIDRef) uint32 {
	return _CGDisplayGetDisplayIDFromUUID(uuid)
}/* debug [functions.gen.go/function]: CGDisplayGetDisplayIDFromUUID */

// ColorSyncAPIVersion is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncAPIVersion()
func ColorSyncAPIVersion() uint32 {
	return _ColorSyncAPIVersion()
}/* debug [functions.gen.go/function]: ColorSyncAPIVersion */

// ColorSyncCMMCopyCMMIdentifier is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyCMMIdentifier(_:)
func ColorSyncCMMCopyCMMIdentifier(p0 ColorSyncCMMRef) StringRef {
	return _ColorSyncCMMCopyCMMIdentifier(p0)
}/* debug [functions.gen.go/function]: ColorSyncCMMCopyCMMIdentifier */

// ColorSyncCMMCopyLocalizedName is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCopyLocalizedName(_:)
func ColorSyncCMMCopyLocalizedName(p0 ColorSyncCMMRef) StringRef {
	return _ColorSyncCMMCopyLocalizedName(p0)
}/* debug [functions.gen.go/function]: ColorSyncCMMCopyLocalizedName */

// ColorSyncCMMCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMCreate(_:)
func ColorSyncCMMCreate(cmmBundle BundleRef) ColorSyncCMMRef {
	return _ColorSyncCMMCreate(cmmBundle)
}/* debug [functions.gen.go/function]: ColorSyncCMMCreate */

// ColorSyncCMMGetBundle is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetBundle(_:)
func ColorSyncCMMGetBundle(p0 ColorSyncCMMRef) BundleRef {
	return _ColorSyncCMMGetBundle(p0)
}/* debug [functions.gen.go/function]: ColorSyncCMMGetBundle */

// ColorSyncCMMGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMGetTypeID()
func ColorSyncCMMGetTypeID() TypeID {
	return _ColorSyncCMMGetTypeID()
}/* debug [functions.gen.go/function]: ColorSyncCMMGetTypeID */

// ColorSyncCreateCodeFragment is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCreateCodeFragment(_:_:)
func ColorSyncCreateCodeFragment(profileSequence ArrayRef, options DictionaryRef) TypeRef {
	return _ColorSyncCreateCodeFragment(profileSequence, options)
}/* debug [functions.gen.go/function]: ColorSyncCreateCodeFragment */

// ColorSyncDeviceCopyDeviceInfo is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceCopyDeviceInfo(_:_:)
func ColorSyncDeviceCopyDeviceInfo(deviceClass StringRef, devID UUIDRef) DictionaryRef {
	return _ColorSyncDeviceCopyDeviceInfo(deviceClass, devID)
}/* debug [functions.gen.go/function]: ColorSyncDeviceCopyDeviceInfo */

// ColorSyncDeviceSetCustomProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceSetCustomProfiles(_:_:_:)
func ColorSyncDeviceSetCustomProfiles(deviceClass StringRef, deviceID UUIDRef, profileInfo DictionaryRef) bool {
	return _ColorSyncDeviceSetCustomProfiles(deviceClass, deviceID, profileInfo)
}/* debug [functions.gen.go/function]: ColorSyncDeviceSetCustomProfiles */

// ColorSyncIterateDeviceProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateDeviceProfiles(_:_:)
func ColorSyncIterateDeviceProfiles(callBack ColorSyncDeviceProfileIterateCallback, userInfo unsafe.Pointer) {
	_ColorSyncIterateDeviceProfiles(callBack, userInfo)
}/* debug [functions.gen.go/function]: ColorSyncIterateDeviceProfiles */

// ColorSyncIterateInstalledCMMs is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledCMMs(_:_:)
func ColorSyncIterateInstalledCMMs(callBack ColorSyncCMMIterateCallback, userInfo unsafe.Pointer) {
	_ColorSyncIterateInstalledCMMs(callBack, userInfo)
}/* debug [functions.gen.go/function]: ColorSyncIterateInstalledCMMs */

// ColorSyncIterateInstalledProfiles is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfiles(_:_:_:_:)
func ColorSyncIterateInstalledProfiles(callBack ColorSyncProfileIterateCallback, seed []uint32, userInfo unsafe.Pointer, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfiles(callBack, seed, userInfo, error_)
}/* debug [functions.gen.go/function]: ColorSyncIterateInstalledProfiles */

// ColorSyncIterateInstalledProfilesWithOptions is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncIterateInstalledProfilesWithOptions(_:_:_:_:_:)
func ColorSyncIterateInstalledProfilesWithOptions(callBack ColorSyncProfileIterateCallback, seed []uint32, userInfo unsafe.Pointer, options DictionaryRef, error_ unsafe.Pointer) {
	_ColorSyncIterateInstalledProfilesWithOptions(callBack, seed, userInfo, options, error_)
}/* debug [functions.gen.go/function]: ColorSyncIterateInstalledProfilesWithOptions */

// ColorSyncProfileContainsTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileContainsTag(_:_:)
func ColorSyncProfileContainsTag(prof ColorSyncProfileRef, signature StringRef) bool {
	return _ColorSyncProfileContainsTag(prof, signature)
}/* debug [functions.gen.go/function]: ColorSyncProfileContainsTag */

// ColorSyncProfileCopyData is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyData(_:_:)
func ColorSyncProfileCopyData(prof ColorSyncProfileRef, error_ unsafe.Pointer) DataRef {
	return _ColorSyncProfileCopyData(prof, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileCopyData */

// ColorSyncProfileCopyDescriptionString is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyDescriptionString(_:)
func ColorSyncProfileCopyDescriptionString(prof ColorSyncProfileRef) StringRef {
	return _ColorSyncProfileCopyDescriptionString(prof)
}/* debug [functions.gen.go/function]: ColorSyncProfileCopyDescriptionString */

// ColorSyncProfileCopyHeader is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyHeader(_:)
func ColorSyncProfileCopyHeader(prof ColorSyncProfileRef) DataRef {
	return _ColorSyncProfileCopyHeader(prof)
}/* debug [functions.gen.go/function]: ColorSyncProfileCopyHeader */

// ColorSyncProfileCopyTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTag(_:_:)
func ColorSyncProfileCopyTag(prof ColorSyncProfileRef, signature StringRef) DataRef {
	return _ColorSyncProfileCopyTag(prof, signature)
}/* debug [functions.gen.go/function]: ColorSyncProfileCopyTag */

// ColorSyncProfileCopyTagSignatures is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCopyTagSignatures(_:)
func ColorSyncProfileCopyTagSignatures(prof ColorSyncProfileRef) ArrayRef {
	return _ColorSyncProfileCopyTagSignatures(prof)
}/* debug [functions.gen.go/function]: ColorSyncProfileCopyTagSignatures */

// ColorSyncProfileCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreate(_:_:)
func ColorSyncProfileCreate(data DataRef, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreate(data, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreate */

// ColorSyncProfileCreateDeviceProfile is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDeviceProfile(_:_:_:)
func ColorSyncProfileCreateDeviceProfile(deviceClass StringRef, deviceID UUIDRef, profileID TypeRef) ColorSyncProfileRef {
	return _ColorSyncProfileCreateDeviceProfile(deviceClass, deviceID, profileID)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateDeviceProfile */

// ColorSyncProfileCreateDisplayTransferTablesFromVCGT is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateDisplayTransferTablesFromVCGT(_:_:)
func ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile ColorSyncProfileRef, nSamplesPerChannel unsafe.Pointer) DataRef {
	return _ColorSyncProfileCreateDisplayTransferTablesFromVCGT(profile, nSamplesPerChannel)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateDisplayTransferTablesFromVCGT */

// ColorSyncProfileCreateLink is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateLink(_:_:)
func ColorSyncProfileCreateLink(profileInfo ArrayRef, options DictionaryRef) ColorSyncProfileRef {
	return _ColorSyncProfileCreateLink(profileInfo, options)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateLink */

// ColorSyncProfileCreateMutable is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutable()
func ColorSyncProfileCreateMutable() ColorSyncMutableProfileRef {
	return _ColorSyncProfileCreateMutable()
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateMutable */

// ColorSyncProfileCreateMutableCopy is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateMutableCopy(_:)
func ColorSyncProfileCreateMutableCopy(prof ColorSyncProfileRef) ColorSyncMutableProfileRef {
	return _ColorSyncProfileCreateMutableCopy(prof)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateMutableCopy */

// ColorSyncProfileCreateWithDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithDisplayID(_:)
func ColorSyncProfileCreateWithDisplayID(displayID uint32) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithDisplayID(displayID)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateWithDisplayID */

// ColorSyncProfileCreateWithName is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithName(_:)
func ColorSyncProfileCreateWithName(name StringRef) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithName(name)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateWithName */

// ColorSyncProfileCreateWithURL is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileCreateWithURL(_:_:)
func ColorSyncProfileCreateWithURL(url URLRef, error_ unsafe.Pointer) ColorSyncProfileRef {
	return _ColorSyncProfileCreateWithURL(url, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileCreateWithURL */

// ColorSyncProfileEstimateGamma is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGamma(_:_:)
func ColorSyncProfileEstimateGamma(prof ColorSyncProfileRef, error_ unsafe.Pointer) float32 {
	return _ColorSyncProfileEstimateGamma(prof, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileEstimateGamma */

// ColorSyncProfileEstimateGammaWithDisplayID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileEstimateGammaWithDisplayID(_:_:)
func ColorSyncProfileEstimateGammaWithDisplayID(displayID unsafe.Pointer, error_ unsafe.Pointer) float32 {
	return _ColorSyncProfileEstimateGammaWithDisplayID(displayID, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileEstimateGammaWithDisplayID */

// ColorSyncProfileGetDisplayTransferFormulaFromVCGT is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetDisplayTransferFormulaFromVCGT(_:_:_:_:_:_:_:_:_:_:)
func ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile ColorSyncProfileRef, redMin []float32, redMax []float32, redGamma []float32, greenMin []float32, greenMax []float32, greenGamma []float32, blueMin []float32, blueMax []float32, blueGamma []float32) bool {
	return _ColorSyncProfileGetDisplayTransferFormulaFromVCGT(profile, redMin, redMax, redGamma, greenMin, greenMax, greenGamma, blueMin, blueMax, blueGamma)
}/* debug [functions.gen.go/function]: ColorSyncProfileGetDisplayTransferFormulaFromVCGT */

// ColorSyncProfileGetMD5 is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetMD5(_:)
func ColorSyncProfileGetMD5(prof ColorSyncProfileRef) ColorSyncMD5 {
	return _ColorSyncProfileGetMD5(prof)
}/* debug [functions.gen.go/function]: ColorSyncProfileGetMD5 */

// ColorSyncProfileGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetTypeID()
func ColorSyncProfileGetTypeID() TypeID {
	return _ColorSyncProfileGetTypeID()
}/* debug [functions.gen.go/function]: ColorSyncProfileGetTypeID */

// ColorSyncProfileGetURL is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileGetURL(_:_:)
func ColorSyncProfileGetURL(prof ColorSyncProfileRef, error_ unsafe.Pointer) URLRef {
	return _ColorSyncProfileGetURL(prof, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileGetURL */

// ColorSyncProfileInstall is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileInstall(_:_:_:_:)
func ColorSyncProfileInstall(profile ColorSyncProfileRef, domain StringRef, subpath StringRef, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileInstall(profile, domain, subpath, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileInstall */

// ColorSyncProfileIsHLGBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsHLGBased(_:)
func ColorSyncProfileIsHLGBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsHLGBased(p0)
}/* debug [functions.gen.go/function]: ColorSyncProfileIsHLGBased */

// ColorSyncProfileIsMatrixBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsMatrixBased(_:)
func ColorSyncProfileIsMatrixBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsMatrixBased(p0)
}/* debug [functions.gen.go/function]: ColorSyncProfileIsMatrixBased */

// ColorSyncProfileIsPQBased is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsPQBased(_:)
func ColorSyncProfileIsPQBased(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsPQBased(p0)
}/* debug [functions.gen.go/function]: ColorSyncProfileIsPQBased */

// ColorSyncProfileIsWideGamut is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIsWideGamut(_:)
func ColorSyncProfileIsWideGamut(p0 ColorSyncProfileRef) bool {
	return _ColorSyncProfileIsWideGamut(p0)
}/* debug [functions.gen.go/function]: ColorSyncProfileIsWideGamut */

// ColorSyncProfileRemoveTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileRemoveTag(_:_:)
func ColorSyncProfileRemoveTag(prof ColorSyncMutableProfileRef, signature StringRef) {
	_ColorSyncProfileRemoveTag(prof, signature)
}/* debug [functions.gen.go/function]: ColorSyncProfileRemoveTag */

// ColorSyncProfileSetHeader is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetHeader(_:_:)
func ColorSyncProfileSetHeader(prof ColorSyncMutableProfileRef, header DataRef) {
	_ColorSyncProfileSetHeader(prof, header)
}/* debug [functions.gen.go/function]: ColorSyncProfileSetHeader */

// ColorSyncProfileSetTag is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileSetTag(_:_:_:)
func ColorSyncProfileSetTag(prof ColorSyncMutableProfileRef, signature StringRef, data DataRef) {
	_ColorSyncProfileSetTag(prof, signature, data)
}/* debug [functions.gen.go/function]: ColorSyncProfileSetTag */

// ColorSyncProfileUninstall is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileUninstall(_:_:)
func ColorSyncProfileUninstall(profile ColorSyncProfileRef, error_ unsafe.Pointer) bool {
	return _ColorSyncProfileUninstall(profile, error_)
}/* debug [functions.gen.go/function]: ColorSyncProfileUninstall */

// ColorSyncProfileVerify is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileVerify(_:_:_:)
func ColorSyncProfileVerify(prof ColorSyncProfileRef, errors unsafe.Pointer, warnings unsafe.Pointer) bool {
	return _ColorSyncProfileVerify(prof, errors, warnings)
}/* debug [functions.gen.go/function]: ColorSyncProfileVerify */

// ColorSyncRegisterDevice is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncRegisterDevice(_:_:_:)
func ColorSyncRegisterDevice(deviceClass StringRef, deviceID UUIDRef, deviceInfo DictionaryRef) bool {
	return _ColorSyncRegisterDevice(deviceClass, deviceID, deviceInfo)
}/* debug [functions.gen.go/function]: ColorSyncRegisterDevice */

// ColorSyncTransformConvert is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformConvert(_:_:_:_:_:_:_:_:_:_:_:_:)
func ColorSyncTransformConvert(transform ColorSyncTransformRef, width uintptr, height uintptr, dst unsafe.Pointer, dstDepth ColorSyncDataDepth, dstLayout ColorSyncDataLayout, dstBytesPerRow uintptr, src unsafe.Pointer, srcDepth ColorSyncDataDepth, srcLayout ColorSyncDataLayout, srcBytesPerRow uintptr, options DictionaryRef) bool {
	return _ColorSyncTransformConvert(transform, width, height, dst, dstDepth, dstLayout, dstBytesPerRow, src, srcDepth, srcLayout, srcBytesPerRow, options)
}/* debug [functions.gen.go/function]: ColorSyncTransformConvert */

// ColorSyncTransformCopyProperty is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCopyProperty(_:_:_:)
func ColorSyncTransformCopyProperty(transform ColorSyncTransformRef, key TypeRef, options DictionaryRef) TypeRef {
	return _ColorSyncTransformCopyProperty(transform, key, options)
}/* debug [functions.gen.go/function]: ColorSyncTransformCopyProperty */

// ColorSyncTransformCreate is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformCreate(_:_:)
func ColorSyncTransformCreate(profileSequence ArrayRef, options DictionaryRef) ColorSyncTransformRef {
	return _ColorSyncTransformCreate(profileSequence, options)
}/* debug [functions.gen.go/function]: ColorSyncTransformCreate */

// ColorSyncTransformGetProfileSequence is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetProfileSequence(_:)
func ColorSyncTransformGetProfileSequence(transform ColorSyncTransformRef) ArrayRef {
	return _ColorSyncTransformGetProfileSequence(transform)
}/* debug [functions.gen.go/function]: ColorSyncTransformGetProfileSequence */

// ColorSyncTransformGetTypeID is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformGetTypeID()
func ColorSyncTransformGetTypeID() TypeID {
	return _ColorSyncTransformGetTypeID()
}/* debug [functions.gen.go/function]: ColorSyncTransformGetTypeID */

// ColorSyncTransformSetProperty is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransformSetProperty(_:_:_:)
func ColorSyncTransformSetProperty(transform ColorSyncTransformRef, key TypeRef, property TypeRef) {
	_ColorSyncTransformSetProperty(transform, key, property)
}/* debug [functions.gen.go/function]: ColorSyncTransformSetProperty */

// ColorSyncUnregisterDevice is a ColorSync function.
//
// Added in macOS 10.13.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncUnregisterDevice(_:_:)
func ColorSyncUnregisterDevice(deviceClass StringRef, deviceID UUIDRef) bool {
	return _ColorSyncUnregisterDevice(deviceClass, deviceID)
}/* debug [functions.gen.go/function]: ColorSyncUnregisterDevice */




