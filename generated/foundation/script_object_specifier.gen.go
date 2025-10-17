// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScriptObjectSpecifier] class.
var ScriptObjectSpecifierClass objc.Class

func init() {
	ScriptObjectSpecifierClass = objc.GetClass("NSScriptObjectSpecifier")
}

type ScriptObjectSpecifier struct {
	objc.ID
}

func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{
		ID: objc.ID(ptr),
	}
}



