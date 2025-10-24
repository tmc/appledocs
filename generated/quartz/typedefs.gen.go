// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz
import (
"unsafe"
)

// Type aliases and typedefs
// QCPlugInBufferReleaseCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInBufferReleaseCallback
// QCPlugInBufferReleaseCallback is a callback function
// C type: void (*)(const void *, void *)
type QCPlugInBufferReleaseCallback = func(unsafe.Pointer, unsafe.Pointer)
// QCPlugInTextureReleaseCallback type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInTextureReleaseCallback
// QCPlugInTextureReleaseCallback is a callback function
// C type: void (*)(struct _CGLContextObject *, unsigned int, void *)
type QCPlugInTextureReleaseCallback = func(unsafe.Pointer, uint32, unsafe.Pointer)

