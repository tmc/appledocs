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

// An interface definition for the [ScriptSuiteRegistry] class.
type IScriptSuiteRegistry interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (sc _ScriptSuiteRegistryClass) Alloc() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScriptSuiteRegistryClass) New() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptSuiteRegistry) Init() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptSuiteRegistry) Autorelease() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptSuiteRegistry creates a new ScriptSuiteRegistry instance.
func NewScriptSuiteRegistry() ScriptSuiteRegistry {
	return scriptSuiteRegistryClass.New()
}




