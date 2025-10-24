// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook
import (
	"unsafe"
)


// C struct types
// QLGeneratorInterfaceStruct - An opaque reference that provides callbacks that the platform uses to interface with a Quick Look plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLGeneratorInterfaceStruct
type QLGeneratorInterfaceStruct struct {
	AddRef unsafe.Pointer
	CancelPreviewGeneration unsafe.Pointer
	CancelThumbnailGeneration unsafe.Pointer
	GeneratePreviewForURL unsafe.Pointer
	GenerateThumbnailForURL unsafe.Pointer
	QueryInterface unsafe.Pointer
	Release unsafe.Pointer
}/* debug [types.gen.go/struct]: QLGeneratorInterfaceStruct */





