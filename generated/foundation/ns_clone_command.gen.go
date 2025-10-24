// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CloneCommand] class.
var (
	CloneCommandClass     _CloneCommandClass
	CloneCommandClassOnce sync.Once
)

func getCloneCommandClass() _CloneCommandClass {
	CloneCommandClassOnce.Do(func() {
		CloneCommandClass = _CloneCommandClass{objc.GetClass("NSCloneCommand")}
	})
	return CloneCommandClass
}

type _CloneCommandClass struct {
	class objc.Class
}





// An interface definition for the [CloneCommand] class.
type ICloneCommand interface {
	IScriptCommand
	

	// properties:
	KeySpecifier() IScriptObjectSpecifier
	SetKeySpecifier(value IScriptObjectSpecifier)


	

	// methods:
	SetReceiversSpecifier(receiversRef IScriptObjectSpecifier)


}





// Alloc allocates a new instance without initialization.
func (cc _CloneCommandClass) Alloc() CloneCommand {
	rv := objc.Send[CloneCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CloneCommandClass) New() CloneCommand {
	rv := objc.Send[CloneCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloneCommand) Init() CloneCommand {
	rv := objc.Send[CloneCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloneCommand) Autorelease() CloneCommand {
	rv := objc.Send[CloneCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloneCommand creates a new CloneCommand instance.
func NewCloneCommand() CloneCommand {
	return getCloneCommandClass().New()
}





// A command that clones one or more scriptable objects.
//
// An instance of clones the specified scriptable object or objects (such as words, paragraphs, images, and so on) and inserts them in the specified location, or the default location if no location is specified. The cloned scriptable objects typically correspond to objects in the application, but aren’t required to. This command corresponds to AppleScript’s command. is part of Cocoa’s built-in scripting support. It works automatically to support the command through key-value coding. Most applications don’t need to subclass or invoke its methods. When an instance of is executed, it clones the specified objects by sending them messages.


// A command that clones one or more scriptable objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloneCommand
type CloneCommand struct {
	ScriptCommand
}

// CloneCommandFrom constructs a [CloneCommand] from an unsafe.Pointer.
//
// A command that clones one or more scriptable objects.
func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}




















// Sets the receiver’s object specifier;.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCloneCommand/setReceiversSpecifier(_:)
func (c_ CloneCommand) SetReceiversSpecifier(receiversRef IScriptObjectSpecifier) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReceiversSpecifier:"), receiversRef)
}







// Returns a specifier for the object or objects to be cloned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclonecommand/keyspecifier
func (c_ CloneCommand) KeySpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](c_.ID, objc.Sel("keySpecifier"))
	return rv
}


// Returns a specifier for the object or objects to be cloned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclonecommand/keyspecifier
func (c_ CloneCommand) SetKeySpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySpecifier:"), value)
}








