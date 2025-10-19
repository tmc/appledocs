// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DeleteCommand] class.
var (
	deleteCommandClass     _DeleteCommandClass
	deleteCommandClassOnce sync.Once
)

func getDeleteCommandClass() _DeleteCommandClass {
	deleteCommandClassOnce.Do(func() {
		deleteCommandClass = _DeleteCommandClass{objc.GetClass("NSDeleteCommand")}
	})
	return deleteCommandClass
}

type _DeleteCommandClass struct {
	class objc.Class
}

// An interface definition for the [DeleteCommand] class.
type IDeleteCommand interface {
	IScriptCommand
}

// A command that deletes a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDeleteCommand
type DeleteCommand struct {
	ScriptCommand
}

// DeleteCommandFrom constructs a [DeleteCommand] from an unsafe.Pointer.
//
// A command that deletes a scriptable object.
func DeleteCommandFrom(ptr unsafe.Pointer) DeleteCommand {
	return DeleteCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DeleteCommandClass) Alloc() DeleteCommand {
	rv := objc.Send[DeleteCommand](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DeleteCommandClass) New() DeleteCommand {
	rv := objc.Send[DeleteCommand](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DeleteCommand) Init() DeleteCommand {
	rv := objc.Send[DeleteCommand](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DeleteCommand) Autorelease() DeleteCommand {
	rv := objc.Send[DeleteCommand](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDeleteCommand creates a new DeleteCommand instance.
func NewDeleteCommand() DeleteCommand {
	return getDeleteCommandClass().New()
}




