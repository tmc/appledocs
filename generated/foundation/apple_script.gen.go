// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleScript] class.
var appleScriptClass = _AppleScriptClass{objc.GetClass("NSAppleScript")}

type _AppleScriptClass struct {
	class objc.Class
}

// An interface definition for the [AppleScript] class.
type IAppleScript interface {
	objectivec.IObject
}

// An object that provides the ability to load, compile, and execute scripts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript

type AppleScript struct {
	objectivec.Object
}

// AppleScriptFrom constructs a [AppleScript] from an unsafe.Pointer.
//
// An object that provides the ability to load, compile, and execute scripts.
func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{objectivec.Object{objc.ID(ptr)}}
}



