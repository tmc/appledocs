// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync
import (
"unsafe"
)

// Type aliases and typedefs
// MApplyTransformProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMApplyTransformProc
// CMMApplyTransformProc is a callback function
// C type: _Bool (*)(struct ColorSyncTransform *, unsigned long, unsigned long, unsigned long, void **, enum ColorSyncDataDepth, unsigned int, unsigned long, unsigned long, const void **, enum ColorSyncDataDepth, unsigned int, unsigned long, const struct __CFDictionary *)
type MApplyTransformProc = func(unsafe.Pointer, uint, uint, uint, unsafe.Pointer, ColorSyncDataDepth, uint32, uint, uint, unsafe.Pointer, ColorSyncDataDepth, uint32, uint, unsafe.Pointer) bool
// MCreateTransformPropertyProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMCreateTransformPropertyProc
// CMMCreateTransformPropertyProc is a callback function
// C type: const void *(*)(struct ColorSyncTransform *, const void *, const struct __CFDictionary *)
type MCreateTransformPropertyProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
// MInitializeLinkProfileProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMInitializeLinkProfileProc
// CMMInitializeLinkProfileProc is a callback function
// C type: _Bool (*)(struct ColorSyncProfile *, const struct __CFArray *, const struct __CFDictionary *)
type MInitializeLinkProfileProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// MInitializeTransformProc type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/CMMInitializeTransformProc
// CMMInitializeTransformProc is a callback function
// C type: _Bool (*)(struct ColorSyncTransform *, const struct __CFArray *, const struct __CFDictionary *)
type MInitializeTransformProc = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
// ColorSyncCMMRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMM
// ColorSyncCMMRef has base type: struct ColorSyncCMM *
type ColorSyncCMMRef uintptr
// ColorSyncCMMIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncCMMIterateCallback
// ColorSyncCMMIterateCallback is a callback function
// C type: _Bool (*)(struct ColorSyncCMM *, void *)
type ColorSyncCMMIterateCallback = func(unsafe.Pointer, unsafe.Pointer) bool
// ColorSyncDataLayout type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDataLayout
// ColorSyncDataLayout has base type: uint32_t
type ColorSyncDataLayout uintptr
// ColorSyncDeviceProfileIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncDeviceProfileIterateCallback
// ColorSyncDeviceProfileIterateCallback is a callback function
// C type: _Bool (*)(const struct __CFDictionary *, void *)
type ColorSyncDeviceProfileIterateCallback = func(unsafe.Pointer, unsafe.Pointer) bool
// ColorSyncMutableProfileRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncMutableProfile
// ColorSyncMutableProfileRef has base type: struct ColorSyncProfile *
type ColorSyncMutableProfileRef uintptr
// ColorSyncProfileRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfile
// ColorSyncProfileRef has base type: const struct ColorSyncProfile *
type ColorSyncProfileRef uintptr
// ColorSyncProfileIterateCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncProfileIterateCallback
// ColorSyncProfileIterateCallback is a callback function
// C type: _Bool (*)(const struct __CFDictionary *, void *)
type ColorSyncProfileIterateCallback = func(unsafe.Pointer, unsafe.Pointer) bool
// ColorSyncTransformRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ColorSync/ColorSyncTransform
// ColorSyncTransformRef has base type: struct ColorSyncTransform *
type ColorSyncTransformRef uintptr

