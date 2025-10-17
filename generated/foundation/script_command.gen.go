// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCommand] class.
var scriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}

type _ScriptCommandClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCommand] class.
type IScriptCommand interface {
	objectivec.IObject
}

// A self-contained scripting statement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptCommand

type ScriptCommand struct {
	objectivec.Object
}

// ScriptCommandFrom constructs a [ScriptCommand] from an unsafe.Pointer.
//
// A self-contained scripting statement.
func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{objectivec.Object{objc.ID(ptr)}}
}



