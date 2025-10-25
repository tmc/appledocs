// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAppleScript */


/* debug [class_header]: Header for NSAppleScript */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AppleScript */
// An interface definition for the [AppleScript] class.
type IAppleScript interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AppleScript */
	// properties:
	Compiled() bool
	RichTextSource() IAttributedString
	Source() IString
	IsCompiled() bool
	SetIsCompiled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AppleScript */
	// methods:
	CompileAndReturnError(errorInfo IDictionary) bool
	ExecuteAndReturnError(errorInfo IDictionary) IAppleEventDescriptor
	ExecuteAppleEventError(event IAppleEventDescriptor, errorInfo IDictionary) IAppleEventDescriptor
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AppleScript */
// Alloc allocates a new instance without initialization.
func (ac _AppleScriptClass) Alloc() AppleScript {
	rv := objc.Send[AppleScript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AppleScript */
// An object that provides the ability to load, compile, and execute scripts.
//
// This class provides applications with the ability to load a script from a URL or from a text string compile or execute a script or an individual Apple event obtain an containing the reply from an executed script or event obtain an attributed string for a compiled script, suitable for display in a script editor obtain various kinds of information about any errors that may occur When you create an instance of object, you can use a URL to specify a script that can be in either text or compiled form, or you can supply the script as a string. Should an error occur when compiling or executing the script, several of the methods return a dictionary containing error information. The keys for obtaining error information, such as , are described in the Constants section. See also NSAppleScript Additions Reference in the Application Kit framework, which defines a method that returns the syntax-highlighted source code for a script.


// An object that provides the ability to load, compile, and execute scripts.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AppleScript */

// Initializes a newly allocated script instance from the source identified by the passed URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(contentsOf:error:)
func NewAppleScriptWithContentsOfURLError(url IURL, errorInfo IDictionary) AppleScript {
	instance := getAppleScriptClass().Alloc()
	rv := objc.Send[AppleScript](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, errorInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAppleScriptWithContentsOfURLError */


// Initializes a newly allocated script instance from the passed source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(source:)
func NewAppleScriptWithSource(source IString) AppleScript {
	instance := getAppleScriptClass().Alloc()
	rv := objc.Send[AppleScript](instance.ID, objc.Sel("initWithSource:"), source)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAppleScriptWithSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AppleScript */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AppleScript */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AppleScript */

// Compiles the receiver, if it is not already compiled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/compileAndReturnError(_:)
func (a_ AppleScript) CompileAndReturnError(errorInfo IDictionary) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compileAndReturnError:"), errorInfo)
	return rv
}/* debug [instance_methods/method]: CompileAndReturnError */


// Executes the receiver, compiling it first if it is not already compiled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/executeAndReturnError(_:)
func (a_ AppleScript) ExecuteAndReturnError(errorInfo IDictionary) IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("executeAndReturnError:"), errorInfo)
	return rv
}/* debug [instance_methods/method]: ExecuteAndReturnError */


// Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/executeAppleEvent(_:error:)
func (a_ AppleScript) ExecuteAppleEventError(event IAppleEventDescriptor, errorInfo IDictionary) IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("executeAppleEvent:error:"), event, errorInfo)
	return rv
}/* debug [instance_methods/method]: ExecuteAppleEventError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AppleScript */

// A Boolean value that indicates whether the receiver’s script has been compiled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/isCompiled
func (a_ AppleScript) Compiled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("compiled"))
	return rv
}/* debug [instance_properties/getter]: compiled */


// Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/richTextSource
func (a_ AppleScript) RichTextSource() IAttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("richTextSource"))
	return rv
}/* debug [instance_properties/getter]: richTextSource */


// The script source for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleScript/source
func (a_ AppleScript) Source() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// A Boolean value that indicates whether the receiver’s script has been compiled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/iscompiled
func (a_ AppleScript) IsCompiled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isCompiled"))
	return rv
}/* debug [instance_properties/getter]: isCompiled */


// A Boolean value that indicates whether the receiver’s script has been compiled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsapplescript/iscompiled
func (a_ AppleScript) SetIsCompiled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsCompiled:"), value)
}/* debug [instance_properties/setter]: isCompiled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAppleScript */


