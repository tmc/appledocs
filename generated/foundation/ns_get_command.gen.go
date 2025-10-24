// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GetCommand] class.
var (
	GetCommandClass     _GetCommandClass
	GetCommandClassOnce sync.Once
)

func getGetCommandClass() _GetCommandClass {
	GetCommandClassOnce.Do(func() {
		GetCommandClass = _GetCommandClass{objc.GetClass("NSGetCommand")}
	})
	return GetCommandClass
}

type _GetCommandClass struct {
	class objc.Class
}





// An interface definition for the [GetCommand] class.
type IGetCommand interface {
	IScriptCommand
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GetCommandClass) Alloc() GetCommand {
	rv := objc.Send[GetCommand](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GetCommandClass) New() GetCommand {
	rv := objc.Send[GetCommand](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GetCommand) Init() GetCommand {
	rv := objc.Send[GetCommand](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GetCommand) Autorelease() GetCommand {
	rv := objc.Send[GetCommand](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGetCommand creates a new GetCommand instance.
func NewGetCommand() GetCommand {
	return getGetCommandClass().New()
}





// A command that retrieves a value or object from a scriptable object.
//
// An instance of gets the specified value or object from the specified scriptable object: for example, the words from a paragraph or the name of a document. When an instance of is executed, it evaluates the specified receivers, gathers the specified data, if any, and packages it in a return Apple event. is part of Cocoa’s built-in scripting support. It works automatically to support the command through key-value coding. Most applications don’t need to subclass or call its methods. For information on working with commands, see in .


// A command that retrieves a value or object from a scriptable object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetCommand
type GetCommand struct {
	ScriptCommand
}

// GetCommandFrom constructs a [GetCommand] from an unsafe.Pointer.
//
// A command that retrieves a value or object from a scriptable object.
func GetCommandFrom(ptr unsafe.Pointer) GetCommand {
	return GetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}































