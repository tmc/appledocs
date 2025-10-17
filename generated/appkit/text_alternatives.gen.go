// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAlternatives] class.
var textAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}

type _TextAlternativesClass struct {
	class objc.Class
}

// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	objectivec.IObject
}

// A list of alternative strings for a piece of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlternatives

type TextAlternatives struct {
	objectivec.Object
}

// TextAlternativesFrom constructs a [TextAlternatives] from an unsafe.Pointer.
//
// A list of alternative strings for a piece of text.
func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{objectivec.Object{objc.ID(ptr)}}
}



