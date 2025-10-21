// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptSuiteRegistry] class.
var (
	ScriptSuiteRegistryClass     _ScriptSuiteRegistryClass
	ScriptSuiteRegistryClassOnce sync.Once
)

func getScriptSuiteRegistryClass() _ScriptSuiteRegistryClass {
	ScriptSuiteRegistryClassOnce.Do(func() {
		ScriptSuiteRegistryClass = _ScriptSuiteRegistryClass{objc.GetClass("NSScriptSuiteRegistry")}
	})
	return ScriptSuiteRegistryClass
}

type _ScriptSuiteRegistryClass struct {
	class objc.Class
}

// An interface definition for the [ScriptSuiteRegistry] class.
type IScriptSuiteRegistry interface {
	objectivec.IObject
	CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode unsafe.Pointer, appleEventIDCode unsafe.Pointer) unsafe.Pointer
	LoadSuiteWithDictionaryFromBundle(suiteDeclaration objc.ID, bundle unsafe.Pointer)
	LoadSuitesFromBundle(bundle unsafe.Pointer)
	RegisterCommandDescription(commandDescription unsafe.Pointer)
	RegisterClassDescription(classDescription unsafe.Pointer)
	SuiteForAppleEventCode(appleEventCode unsafe.Pointer) string
}

// The top-level repository of scriptability information for an app at runtime.
//
// Scriptability information specifies the terminology available for use in scripts that target an application. It also provides information, used by AppleScript and by Cocoa, about how support for that terminology is implemented in the application. This information includes descriptions of the scriptable object classes in an application and of the commands the application supports. There are two standard formats for supplying scriptability information: the older script suite format, consisting of a script suite file and one or more script terminology files, and the newer scripting definition (or sdef) format, consisting of a single sdef file. There is one instance of per scriptable application. This registry object collects scriptability information when the application first needs to respond to an Apple event for which Cocoa hasn’t installed a default event handler. It then creates one instance of for each object class and one instance of for each command class, and installs a command handler for each command. When a user executes an AppleScript script, Apple events are sent to the targeted application. Using the information stored in the registry object, Cocoa automatically converts incoming Apple events into script commands (based on or a subclass) that manipulate objects in the application. The public methods of are used primarily by Cocoa’s built-in scripting support. You should not need to create a subclass of . For information on scriptability information formats, loading of scriptability information, and related topics, see “Scriptability Information” in in .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry
type ScriptSuiteRegistry struct {
	objectivec.Object
}

// ScriptSuiteRegistryFrom constructs a [ScriptSuiteRegistry] from an unsafe.Pointer.
//
// The top-level repository of scriptability information for an app at runtime.
func ScriptSuiteRegistryFrom(ptr unsafe.Pointer) ScriptSuiteRegistry {
	return ScriptSuiteRegistry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptSuiteRegistryClass) Alloc() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptSuiteRegistryClass) New() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptSuiteRegistry) Init() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptSuiteRegistry) Autorelease() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptSuiteRegistry creates a new ScriptSuiteRegistry instance.
func NewScriptSuiteRegistry() ScriptSuiteRegistry {
	return getScriptSuiteRegistryClass().New()
}


// Returns the command description identified by a suite’s four-character Apple event code of the class ( ) and the four-character Apple event code of the command ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/commandDescription(withAppleEventClass:andAppleEventCode:)
func (s_ ScriptSuiteRegistry) CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode unsafe.Pointer, appleEventIDCode unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("commandDescriptionWithAppleEventClass:andAppleEventCode:"), appleEventClassCode, appleEventIDCode)
	return rv
}

// Loads the suite definition encapsulated in ; previously, this suite definition was parsed from a property list contained in a framework or in .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/loadSuite(with:from:)
func (s_ ScriptSuiteRegistry) LoadSuiteWithDictionaryFromBundle(suiteDeclaration objc.ID, bundle unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadSuiteWithDictionary:fromBundle:"), suiteDeclaration, bundle)
}

// Loads the suite definitions in bundle , invoking for each suite found.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/loadSuites(from:)
func (s_ ScriptSuiteRegistry) LoadSuitesFromBundle(bundle unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadSuitesFromBundle:"), bundle)
}

// Registers command description for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the command name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/register(_:)-5mq91
func (s_ ScriptSuiteRegistry) RegisterCommandDescription(commandDescription unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerCommandDescription:"), commandDescription)
}

// Registers class description for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the class name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/register(_:)-9aplw
func (s_ ScriptSuiteRegistry) RegisterClassDescription(classDescription unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerClassDescription:"), classDescription)
}

// Returns the name of the suite definition associated with the given four-character Apple event code, .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/suite(forAppleEventCode:)
func (s_ ScriptSuiteRegistry) SuiteForAppleEventCode(appleEventCode unsafe.Pointer) string {
	rv := objc.Send[string](s_.ID, objc.Sel("suiteForAppleEventCode:"), appleEventCode)
	return rv
}



