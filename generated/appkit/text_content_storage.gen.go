// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextContentStorage] class.
var textContentStorageClass = _TextContentStorageClass{objc.GetClass("NSTextContentStorage")}

type _TextContentStorageClass struct {
	class objc.Class
}

// A concrete object for managing your view’s text content and generating the text elements necessary for layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentStorage

type TextContentStorage struct {
	TextContentManager
}

// TextContentStorageFrom constructs a [TextContentStorage] from an unsafe.Pointer.
//
// A concrete object for managing your view’s text content and generating the text elements necessary for layout.
func TextContentStorageFrom(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		TextContentManager: TextContentManagerFrom(ptr),
	}
}



