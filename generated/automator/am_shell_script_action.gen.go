// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AMShellScriptAction] class.
type IAMShellScriptAction interface {
	IAMBundleAction
}

// An object that represents Automator actions whose runtime behavior is driven by a shell script or by a Perl or Python script.
//
// When you create a Shell Script Automator Action project in Xcode, the project template supplies an instance as the Principal Class of the action bundle. This ready-made instance provides a default implementation of the method that uses the logic defined in the script. You can substitute your own subclass of for Principal Class if you need to.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AMShellScriptActionClass) Alloc() AMShellScriptAction {
	rv := objc.Send[AMShellScriptAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A string to use as the delimiter between items in the string passed to the action through standard input.
//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/inputFieldSeparator
func (a_ AMShellScriptAction) InputFieldSeparator() string {
	rv := objc.Send[string](a_.ID, objc.Sel("inputFieldSeparator"))
	return rv
}

// A string to use as a delimiter in the string output by the action.
//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/outputFieldSeparator
func (a_ AMShellScriptAction) OutputFieldSeparator() string {
	rv := objc.Send[string](a_.ID, objc.Sel("outputFieldSeparator"))
	return rv
}

// A Boolean value that indicates whether you want automatic remapping of carriage return ( ) to newline ( ) characters in the input string.
//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMShellScriptAction/remapLineEndings
func (a_ AMShellScriptAction) RemapLineEndings() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("remapLineEndings"))
	return rv
}



