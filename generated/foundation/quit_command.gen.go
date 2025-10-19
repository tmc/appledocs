// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QuitCommand] class.
var (
	quitCommandClass     _QuitCommandClass
	quitCommandClassOnce sync.Once
)

func getQuitCommandClass() _QuitCommandClass {
	quitCommandClassOnce.Do(func() {
		quitCommandClass = _QuitCommandClass{objc.GetClass("NSQuitCommand")}
	})
	return quitCommandClass
}

type _QuitCommandClass struct {
	class objc.Class
}

// An interface definition for the [QuitCommand] class.
type IQuitCommand interface {
	IScriptCommand
}

// A command that quits the specified app.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSQuitCommand
type QuitCommand struct {
	ScriptCommand
}

// QuitCommandFrom constructs a [QuitCommand] from an unsafe.Pointer.
//
// A command that quits the specified app.
func QuitCommandFrom(ptr unsafe.Pointer) QuitCommand {
	return QuitCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QuitCommandClass) Alloc() QuitCommand {
	rv := objc.Send[QuitCommand](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuitCommandClass) New() QuitCommand {
	rv := objc.Send[QuitCommand](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuitCommand) Init() QuitCommand {
	rv := objc.Send[QuitCommand](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuitCommand) Autorelease() QuitCommand {
	rv := objc.Send[QuitCommand](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuitCommand creates a new QuitCommand instance.
func NewQuitCommand() QuitCommand {
	return getQuitCommandClass().New()
}




