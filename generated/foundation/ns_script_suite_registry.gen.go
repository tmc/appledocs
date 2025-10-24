// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScriptSuiteRegistry */


/* debug [class_header]: Header for NSScriptSuiteRegistry */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptSuiteRegistry */
// An interface definition for the [ScriptSuiteRegistry] class.
type IScriptSuiteRegistry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptSuiteRegistry */
	// properties:
	SuiteNames() IString
	SetSuiteNames(value IString)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptSuiteRegistry */
	// methods:
	CommandDescriptionsInSuite(suiteName IString) IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptSuiteRegistry */
// Alloc allocates a new instance without initialization.
func (sc _ScriptSuiteRegistryClass) Alloc() ScriptSuiteRegistry {
	rv := objc.Send[ScriptSuiteRegistry](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptSuiteRegistry */
// The top-level repository of scriptability information for an app at runtime.
//
// Scriptability information specifies the terminology available for use in scripts that target an application. It also provides information, used by AppleScript and by Cocoa, about how support for that terminology is implemented in the application. This information includes descriptions of the scriptable object classes in an application and of the commands the application supports. There are two standard formats for supplying scriptability information: the older script suite format, consisting of a script suite file and one or more script terminology files, and the newer scripting definition (or sdef) format, consisting of a single sdef file. There is one instance of per scriptable application. This registry object collects scriptability information when the application first needs to respond to an Apple event for which Cocoa hasn’t installed a default event handler. It then creates one instance of for each object class and one instance of for each command class, and installs a command handler for each command. When a user executes an AppleScript script, Apple events are sent to the targeted application. Using the information stored in the registry object, Cocoa automatically converts incoming Apple events into script commands (based on or a subclass) that manipulate objects in the application. The public methods of are used primarily by Cocoa’s built-in scripting support. You should not need to create a subclass of . For information on scriptability information formats, loading of scriptability information, and related topics, see “Scriptability Information” in in .


// The top-level repository of scriptability information for an app at runtime.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptSuiteRegistry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptSuiteRegistry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptSuiteRegistry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptSuiteRegistry */

// Returns the command descriptions contained in the suite identified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/commandDescriptions(inSuite:)
func (s_ ScriptSuiteRegistry) CommandDescriptionsInSuite(suiteName IString) IDictionary {
	rv := objc.Send[Dictionary](s_.ID, objc.Sel("commandDescriptionsInSuite:"), suiteName)
	return rv
}/* debug [instance_methods/method]: CommandDescriptionsInSuite */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptSuiteRegistry */

// Returns the names of the suite definitions currently loaded by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/suitenames
func (s_ ScriptSuiteRegistry) SuiteNames() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("suiteNames"))
	return rv
}/* debug [instance_properties/getter]: suiteNames */


// Returns the names of the suite definitions currently loaded by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptsuiteregistry/suitenames
func (s_ ScriptSuiteRegistry) SetSuiteNames(value IString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuiteNames:"), value)
}/* debug [instance_properties/setter]: suiteNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScriptSuiteRegistry */



