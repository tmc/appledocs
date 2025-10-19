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

// An interface definition for the [ScriptCommandDescription] class.
type IScriptCommandDescription interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCommandDescription) Autorelease() ScriptCommandDescription {
	rv := objc.Send[ScriptCommandDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCommandDescription creates a new ScriptCommandDescription instance.
func NewScriptCommandDescription() ScriptCommandDescription {
	return scriptCommandDescriptionClass.New()
}




