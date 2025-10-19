// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScriptClassDescription] class.
var (
	scriptClassDescriptionClass     _ScriptClassDescriptionClass
	scriptClassDescriptionClassOnce sync.Once
)

func getScriptClassDescriptionClass() _ScriptClassDescriptionClass {
	scriptClassDescriptionClassOnce.Do(func() {
		scriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}
	})
	return scriptClassDescriptionClass
}

type _ScriptClassDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ScriptClassDescription] class.
type IScriptClassDescription interface {
	IClassDescription
}

// A scriptable class that a macOS app supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription
type ScriptClassDescription struct {
	ClassDescription
}

// ScriptClassDescriptionFrom constructs a [ScriptClassDescription] from an unsafe.Pointer.
//
// A scriptable class that a macOS app supports.
func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: ClassDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptClassDescription) Autorelease() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptClassDescription creates a new ScriptClassDescription instance.
func NewScriptClassDescription() ScriptClassDescription {
	return getScriptClassDescriptionClass().New()
}




