// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountCommand] class.
var (
	CountCommandClass     _CountCommandClass
	CountCommandClassOnce sync.Once
)

func getCountCommandClass() _CountCommandClass {
	CountCommandClassOnce.Do(func() {
		CountCommandClass = _CountCommandClass{objc.GetClass("NSCountCommand")}
	})
	return CountCommandClass
}

type _CountCommandClass struct {
	class objc.Class
}

// An interface definition for the [CountCommand] class.
type ICountCommand interface {
	IScriptCommand
}

// A command that counts the number of objects of a specified class in the specified object container.
//
// An instance of counts the number of objects of a specified class in the specified object container (such as the number of words in a paragraph or document) and returns the result. is part of Cocoa’s built-in scripting support. It works automatically to support the command through key-value coding. Most applications don’t need to subclass or call its methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountCommand
type CountCommand struct {
	ScriptCommand
}

// CountCommandFrom constructs a [CountCommand] from an unsafe.Pointer.
//
// A command that counts the number of objects of a specified class in the specified object container.
func CountCommandFrom(ptr unsafe.Pointer) CountCommand {
	return CountCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CountCommandClass) Alloc() CountCommand {
	rv := objc.Send[CountCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CountCommandClass) New() CountCommand {
	rv := objc.Send[CountCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CountCommand) Init() CountCommand {
	rv := objc.Send[CountCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CountCommand) Autorelease() CountCommand {
	rv := objc.Send[CountCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCountCommand creates a new CountCommand instance.
func NewCountCommand() CountCommand {
	return getCountCommandClass().New()
}




