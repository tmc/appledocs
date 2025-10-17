// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PresentationIntent] class.
var presentationIntentClass = _PresentationIntentClass{objc.GetClass("NSPresentationIntent")}

type _PresentationIntentClass struct {
	class objc.Class
}

// An interface definition for the [PresentationIntent] class.
type IPresentationIntent interface {
	objectivec.IObject
}

// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPresentationIntent

type PresentationIntent struct {
	objectivec.Object
}

// PresentationIntentFrom constructs a [PresentationIntent] from an unsafe.Pointer.
//
// A type that contains the Markdown formatting for blocks of text, like paragraphs, lists, code blocks, and parts of tables.
func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{objectivec.Object{objc.ID(ptr)}}
}



