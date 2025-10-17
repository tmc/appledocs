// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextContentManager] class.
var textContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}

type _TextContentManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objectivec.IObject
}

// An abstract class that defines the interface and a default implementation for managing the text document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager

type TextContentManager struct {
	objectivec.Object
}

// TextContentManagerFrom constructs a [TextContentManager] from an unsafe.Pointer.
//
// An abstract class that defines the interface and a default implementation for managing the text document contents.
func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{objectivec.Object{objc.ID(ptr)}}
}



