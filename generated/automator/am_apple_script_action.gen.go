// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AMAppleScriptAction] class.
var (
	AMAppleScriptActionClass     _AMAppleScriptActionClass
	AMAppleScriptActionClassOnce sync.Once
)

func getAMAppleScriptActionClass() _AMAppleScriptActionClass {
	AMAppleScriptActionClassOnce.Do(func() {
		AMAppleScriptActionClass = _AMAppleScriptActionClass{objc.GetClass("AMAppleScriptAction")}
	})
	return AMAppleScriptActionClass
}

type _AMAppleScriptActionClass struct {
	class objc.Class
}

// An interface definition for the [AMAppleScriptAction] class.
type IAMAppleScriptAction interface {
	IAMBundleAction
	// properties:
	Script() unsafe.Pointer
	SetScript(value unsafe.Pointer)
	// methods:
}

// An object that represents Automator actions whose runtime behavior is driven by an AppleScript script.
//
// An object holds the compiled script as an instance of the class. By default, the object is instantiated from the script in the Xcode project file . When you create a Automator Applescript Action project in Xcode, the project template supplies an instance as File’s Owner of the action bundle. This ready-made instance provides a default implementation of the method that uses the logic defined in the script. You can substitute your own subclass of for File’s Owner if you need to.


// An object that represents Automator actions whose runtime behavior is driven by an AppleScript script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMAppleScriptAction
type AMAppleScriptAction struct {
	AMBundleAction
}

// AMAppleScriptActionFrom constructs a [AMAppleScriptAction] from an unsafe.Pointer.
//
// An object that represents Automator actions whose runtime behavior is driven by an AppleScript script.
func AMAppleScriptActionFrom(ptr unsafe.Pointer) AMAppleScriptAction {
	return AMAppleScriptAction{
		AMBundleAction: AMBundleActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AMAppleScriptActionClass) Alloc() AMAppleScriptAction {
	rv := objc.Send[AMAppleScriptAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AMAppleScriptActionClass) New() AMAppleScriptAction {
	rv := objc.Send[AMAppleScriptAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMAppleScriptAction) Init() AMAppleScriptAction {
	rv := objc.Send[AMAppleScriptAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMAppleScriptAction) Autorelease() AMAppleScriptAction {
	rv := objc.Send[AMAppleScriptAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMAppleScriptAction creates a new AMAppleScriptAction instance.
func NewAMAppleScriptAction() AMAppleScriptAction {
	return getAMAppleScriptActionClass().New()
}



// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amapplescriptaction/script
func (a_ AMAppleScriptAction) Script() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("script"))
	return rv
}


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automator/amapplescriptaction/script
func (a_ AMAppleScriptAction) SetScript(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setScript:"), value)
}



