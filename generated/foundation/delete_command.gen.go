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

// A command that deletes a scriptable object.
//
// An instance of deletes the specified scriptable object or objects (such as words, paragraphs, and so on). Suppose, for example, a user executes a script that sends the command to the Sketch sample application (located in ). Cocoa creates an object to perform the operation. When the command is executed, it uses the key-value coding mechanism (by invoking ) to remove the specified object or objects from their container. See the description for for related information. is part of Cocoa’s built-in scripting support. Most applications don’t need to subclass or call its methods.
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




