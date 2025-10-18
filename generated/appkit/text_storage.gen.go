// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextStorage] class.
var textStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}

type _TextStorageClass struct {
	class objc.Class
}

// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	foundation.IMutableAttributedString
}

// The fundamental storage mechanism of TextKit that contains the text managed by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorage

type TextStorage struct {
	foundation.MutableAttributedString
}

// TextStorageFrom constructs a [TextStorage] from an unsafe.Pointer.
//
// The fundamental storage mechanism of TextKit that contains the text managed by the system.
func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: foundation.MutableAttributedStringFrom(ptr),
	}
}



