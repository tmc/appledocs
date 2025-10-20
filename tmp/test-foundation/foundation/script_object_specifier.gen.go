// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ScriptObjectSpecifierClass _ScriptObjectSpecifierClass

func init() {
	ScriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}
}

type _ScriptObjectSpecifierClass struct {
	class objc.Class
}

type ScriptObjectSpecifier struct {
	objc.ID
}

func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{
		ID: objc.ID(ptr),
	}
}




