// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptSuiteRegistry] class.
var scriptSuiteRegistryClass = _ScriptSuiteRegistryClass{objc.GetClass("NSScriptSuiteRegistry")}

type _ScriptSuiteRegistryClass struct {
	class objc.Class
}

// The top-level repository of scriptability information for an app at runtime. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry

type ScriptSuiteRegistry struct {
	objectivec.Object
}

// ScriptSuiteRegistryFrom constructs a [ScriptSuiteRegistry] from an unsafe.Pointer.
//
// The top-level repository of scriptability information for an app at runtime.
func ScriptSuiteRegistryFrom(ptr unsafe.Pointer) ScriptSuiteRegistry {
	return ScriptSuiteRegistry{objectivec.Object{objc.ID(ptr)}}
}



