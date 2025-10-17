// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextList] class.
var textListClass = _TextListClass{objc.GetClass("NSTextList")}

type _TextListClass struct {
	class objc.Class
}

// An interface definition for the [TextList] class.
type ITextList interface {
	objectivec.IObject
}

// A section of text that forms a single list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextList

type TextList struct {
	objectivec.Object
}

// TextListFrom constructs a [TextList] from an unsafe.Pointer.
//
// A section of text that forms a single list.
func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{objectivec.Object{objc.ID(ptr)}}
}



