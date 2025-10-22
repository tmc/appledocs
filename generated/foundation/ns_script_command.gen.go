// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptCommand] class.
var (
	ScriptCommandClass     _ScriptCommandClass
	ScriptCommandClassOnce sync.Once
)

func getScriptCommandClass() _ScriptCommandClass {
	ScriptCommandClassOnce.Do(func() {
		ScriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}
	})
	return ScriptCommandClass
}

type _ScriptCommandClass struct {
	class objc.Class
}

// An interface definition for the [ScriptCommand] class.
type IScriptCommand interface {
	objectivec.IObject
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type ScriptCommand struct {
	objectivec.Object
}

// ScriptCommandFrom constructs a [ScriptCommand] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptCommandClass) Alloc() ScriptCommand {
	rv := objc.Send[ScriptCommand](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptCommandClass) New() ScriptCommand {
	rv := objc.Send[ScriptCommand](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptCommand) Init() ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptCommand) Autorelease() ScriptCommand {
	rv := objc.Send[ScriptCommand](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptCommand creates a new ScriptCommand instance.
func NewScriptCommand() ScriptCommand {
	return getScriptCommandClass().New()
}




