// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync

// Type aliases and typedefs
// CMMApplyTransformProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMApplyTransformProc
// CMMApplyTransformProc has base type: _Bool (*)(struct ColorSyncTransform *, unsigned long, unsigned long, unsigned long, void **, enum ColorSyncDataDepth, unsigned int, unsigned long, unsigned long, const void **, enum ColorSyncDataDepth, unsigned int, unsigned long, const struct __CFDictionary *)
type CMMApplyTransformProc uintptr
// CMMCreateTransformPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMCreateTransformPropertyProc
// CMMCreateTransformPropertyProc has base type: const void *(*)(struct ColorSyncTransform *, const void *, const struct __CFDictionary *)
type CMMCreateTransformPropertyProc uintptr
// CMMInitializeLinkProfileProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMInitializeLinkProfileProc
// CMMInitializeLinkProfileProc has base type: _Bool (*)(struct ColorSyncProfile *, const struct __CFArray *, const struct __CFDictionary *)
type CMMInitializeLinkProfileProc uintptr
// CMMInitializeTransformProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMInitializeTransformProc
// CMMInitializeTransformProc has base type: _Bool (*)(struct ColorSyncTransform *, const struct __CFArray *, const struct __CFDictionary *)
type CMMInitializeTransformProc uintptr
// ColorSyncCMMIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMIterateCallback
// ColorSyncCMMIterateCallback has base type: _Bool (*)(struct ColorSyncCMM *, void *)
type ColorSyncCMMIterateCallback uintptr
// ColorSyncDataLayout type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDataLayout
// ColorSyncDataLayout has base type: uint32_t
type ColorSyncDataLayout uintptr
// ColorSyncDeviceProfileIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceProfileIterateCallback
// ColorSyncDeviceProfileIterateCallback has base type: _Bool (*)(const struct __CFDictionary *, void *)
type ColorSyncDeviceProfileIterateCallback uintptr
// ColorSyncProfileRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfile
// ColorSyncProfileRef has base type: const struct ColorSyncProfile *
type ColorSyncProfileRef uintptr
// ColorSyncProfileIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIterateCallback
// ColorSyncProfileIterateCallback has base type: _Bool (*)(const struct __CFDictionary *, void *)
type ColorSyncProfileIterateCallback uintptr
// ColorSyncCMMRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMM
// ColorSyncCMMRef has base type: struct ColorSyncCMM *
type ColorSyncCMMRef uintptr
// ColorSyncMutableProfileRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncMutableProfile
// ColorSyncMutableProfileRef has base type: struct ColorSyncProfile *
type ColorSyncMutableProfileRef uintptr
// ColorSyncTransformRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransform
// ColorSyncTransformRef has base type: struct ColorSyncTransform *
type ColorSyncTransformRef uintptr

