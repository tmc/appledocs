// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextCheckingResult] class.
var textCheckingResultClass = _TextCheckingResultClass{objc.GetClass("NSTextCheckingResult")}

type _TextCheckingResultClass struct {
	class objc.Class
}

// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult

type TextCheckingResult struct {
	objectivec.Object
}

// TextCheckingResultFrom constructs a [TextCheckingResult] from an unsafe.Pointer.
//
// An occurrence of textual content found during the analysis of a block of text, such as when matching a regular expression.
func TextCheckingResultFrom(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{objectivec.Object{objc.ID(ptr)}}
}



