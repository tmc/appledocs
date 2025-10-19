// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CreateCommand] class.
var (
	createCommandClass     _CreateCommandClass
	createCommandClassOnce sync.Once
)

func getCreateCommandClass() _CreateCommandClass {
	createCommandClassOnce.Do(func() {
		createCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}
	})
	return createCommandClass
}

type _CreateCommandClass struct {
	class objc.Class
}

// An interface definition for the [CreateCommand] class.
type ICreateCommand interface {
	IScriptCommand
}

// A command that creates a scriptable object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateCommand
type CreateCommand struct {
	ScriptCommand
}

// CreateCommandFrom constructs a [CreateCommand] from an unsafe.Pointer.
//
// A command that creates a scriptable object.
func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CreateCommandClass) Alloc() CreateCommand {
	rv := objc.Send[CreateCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CreateCommandClass) New() CreateCommand {
	rv := objc.Send[CreateCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CreateCommand) Init() CreateCommand {
	rv := objc.Send[CreateCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CreateCommand) Autorelease() CreateCommand {
	rv := objc.Send[CreateCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCreateCommand creates a new CreateCommand instance.
func NewCreateCommand() CreateCommand {
	return getCreateCommandClass().New()
}




