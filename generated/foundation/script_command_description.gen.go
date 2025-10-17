// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCommandDescription] class.
var scriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}

type _ScriptCommandDescriptionClass struct {
	class objc.Class
}

// A script command that a macOS app supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommandDescription

type ScriptCommandDescription struct {
	objectivec.Object
}

// ScriptCommandDescriptionFrom constructs a [ScriptCommandDescription] from an unsafe.Pointer.
//
// A script command that a macOS app supports.
func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{objectivec.Object{objc.ID(ptr)}}
}



