// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AMShellScriptAction */


/* debug [class_header]: Header for AMShellScriptAction */
// The class instance for the [AMShellScriptAction] class.
var (
	AMShellScriptActionClass     _AMShellScriptActionClass
	AMShellScriptActionClassOnce sync.Once
)

func getAMShellScriptActionClass() _AMShellScriptActionClass {
	AMShellScriptActionClassOnce.Do(func() {
		AMShellScriptActionClass = _AMShellScriptActionClass{objc.GetClass("AMShellScriptAction")}
	})
	return AMShellScriptActionClass
}

type _AMShellScriptActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AMShellScriptAction */
// An interface definition for the [AMShellScriptAction] class.
type IAMShellScriptAction interface {
	IAMBundleAction
	
/* debug [class_interface_properties]: Properties for AMShellScriptAction */
	// properties:
	InputFieldSeparator() objc.IObject /* cross-framework: NSString */
	OutputFieldSeparator() objc.IObject /* cross-framework: NSString */
	RemapLineEndings() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AMShellScriptAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AMShellScriptAction */
// Alloc allocates a new instance without initialization.
func (ac _AMShellScriptActionClass) Alloc() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AMShellScriptActionClass) New() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMShellScriptAction) Init() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMShellScriptAction) Autorelease() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMShellScriptAction creates a new AMShellScriptAction instance.
func NewAMShellScriptAction() AMShellScriptAction {
	return getAMShellScriptActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AMShellScriptAction */
// An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
//
// When you create a Shell Script Automator Action project in Xcode, the project template supplies an instance as the Principal Class of the action bundle. This ready-made instance provides a default implementation of the method that uses the logic defined in the script. You can substitute your own subclass of for Principal Class if you need to.


// An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction
type AMShellScriptAction struct {
	AMBundleAction
}

// AMShellScriptActionFrom constructs a [AMShellScriptAction] from an unsafe.Pointer.
//
// An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
func AMShellScriptActionFrom(ptr unsafe.Pointer) AMShellScriptAction {
	return AMShellScriptAction{
		AMBundleAction: AMBundleActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AMShellScriptAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AMShellScriptAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AMShellScriptAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AMShellScriptAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AMShellScriptAction */

// A string to use as the delimiter between items in the string passed to the action through standard input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/inputFieldSeparator
func (a_ AMShellScriptAction) InputFieldSeparator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("inputFieldSeparator"))
	return rv
}/* debug [instance_properties/getter]: inputFieldSeparator */


// A string to use as a delimiter in the string output by the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/outputFieldSeparator
func (a_ AMShellScriptAction) OutputFieldSeparator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("outputFieldSeparator"))
	return rv
}/* debug [instance_properties/getter]: outputFieldSeparator */


// A Boolean value that indicates whether you want automatic remapping of carriage return ( ) to newline ( ) characters in the input string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/remapLineEndings
func (a_ AMShellScriptAction) RemapLineEndings() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("remapLineEndings"))
	return rv
}/* debug [instance_properties/getter]: remapLineEndings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AMShellScriptAction */



