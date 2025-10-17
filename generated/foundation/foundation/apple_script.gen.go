// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AppleScript] class.
var AppleScriptClass objc.Class

func init() {
	AppleScriptClass = objc.GetClass("NSAppleScript")
}

type AppleScript struct {
	objc.ID
}

func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{
		ID: objc.ID(ptr),
	}
}




