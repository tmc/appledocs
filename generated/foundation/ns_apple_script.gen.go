// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleScript] class.
var (
	AppleScriptClass     _AppleScriptClass
	AppleScriptClassOnce sync.Once
)

func getAppleScriptClass() _AppleScriptClass {
	AppleScriptClassOnce.Do(func() {
		AppleScriptClass = _AppleScriptClass{objc.GetClass("NSAppleScript")}
	})
	return AppleScriptClass
}

type _AppleScriptClass struct {
	class objc.Class
}

// An interface definition for the [AppleScript] class.
type IAppleScript interface {
	objectivec.IObject
	ExecuteAppleEventError(event unsafe.Pointer, errorInfo unsafe.Pointer) unsafe.Pointer
}

// An object that provides the ability to load, compile, and execute scripts.
//
// This class provides applications with the ability to load a script from a URL or from a text string compile or execute a script or an individual Apple event obtain an containing the reply from an executed script or event obtain an attributed string for a compiled script, suitable for display in a script editor obtain various kinds of information about any errors that may occur When you create an instance of object, you can use a URL to specify a script that can be in either text or compiled form, or you can supply the script as a string. Should an error occur when compiling or executing the script, several of the methods return a dictionary containing error information. The keys for obtaining error information, such as , are described in the Constants section. See also NSAppleScript Additions Reference in the Application Kit framework, which defines a method that returns the syntax-highlighted source code for a script.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript
type AppleScript struct {
	objectivec.Object
}

// AppleScriptFrom constructs a [AppleScript] from an unsafe.Pointer.
//
// An object that provides the ability to load, compile, and execute scripts.
func AppleScriptFrom(ptr unsafe.Pointer) AppleScript {
	return AppleScript{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleScriptClass) Alloc() AppleScript {
	rv := objc.Send[AppleScript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleScriptClass) New() AppleScript {
	rv := objc.Send[AppleScript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleScript) Init() AppleScript {
	rv := objc.Send[AppleScript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleScript) Autorelease() AppleScript {
	rv := objc.Send[AppleScript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleScript creates a new AppleScript instance.
func NewAppleScript() AppleScript {
	return getAppleScriptClass().New()
}


// Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/executeAppleEvent(_:error:)
func (a_ AppleScript) ExecuteAppleEventError(event unsafe.Pointer, errorInfo unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("executeAppleEvent:error:"), event, errorInfo)
	return rv
}

// The script source for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/source
func (a_ AppleScript) Source() string {
	rv := objc.Send[string](a_.ID, objc.Sel("source"))
	return rv
}


// SetSource sets the value of the source property.
// The script source for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/source
func (a_ AppleScript) SetSource(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSource:"), objc.String(value))
}

// Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/richtextsource
func (a_ AppleScript) RichTextSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("richTextSource"))
	return rv
}


// SetRichTextSource sets the value of the richTextSource property.
// Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/richtextsource
func (a_ AppleScript) SetRichTextSource(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRichTextSource:"), value)
}

// A Boolean value that indicates whether the receiver’s script has been compiled.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/iscompiled
func (a_ AppleScript) IsCompiled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompiled"))
	return rv
}


// SetIsCompiled sets the value of the isCompiled property.
// A Boolean value that indicates whether the receiver’s script has been compiled.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/iscompiled
func (a_ AppleScript) SetIsCompiled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompiled:"), value)
}



