// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AppleScript] class.
var AppleScriptClass = _AppleScriptClass{objc.GetClass("NSAppleScript")}

type _AppleScriptClass struct {
	class objc.Class
}

type AppleScript struct {
	objc.ID
}

func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{
		ID: objc.ID(ptr),
	}
}




