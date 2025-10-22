// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AMBundleAction] class.
var (
	AMBundleActionClass     _AMBundleActionClass
	AMBundleActionClassOnce sync.Once
)

func getAMBundleActionClass() _AMBundleActionClass {
	AMBundleActionClassOnce.Do(func() {
		AMBundleActionClass = _AMBundleActionClass{objc.GetClass("AMBundleAction")}
	})
	return AMBundleActionClass
}

type _AMBundleActionClass struct {
	class objc.Class
}

// An interface definition for the [AMBundleAction] class.
type IAMBundleAction interface {
	IAMAction
	AwakeFromBundle()
	Bundle() foundation.Bundle
	HasView() bool
	Parameters() unsafe.Pointer
	SetParameters(value unsafe.Pointer)
	View() appkit.View
}

// An object that represents an Automator action that’s a loadable bundle.
//
// Automator loads action bundles from standard locations in the file system: , , and . objects have several important properties: The object associated with the action’s physical bundle The action’s view, which holds its user interface A parameters dictionary that reflects the settings in the user interface When you create a Cocoa Automator Action project in Xcode, the project template includes a custom subclass of . This custom class uses the name of the project. You must provide an implementation of , which is declared by the superclass . If you add any instance variables, you must override the method and the method of to work with them.


// An object that represents an Automator action that’s a loadable bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction

type AMBundleAction struct {
	AMAction
}

// AMBundleActionFrom constructs a [AMBundleAction] from an unsafe.Pointer.
//
// An object that represents an Automator action that’s a loadable bundle.
func AMBundleActionFrom(ptr unsafe.Pointer) AMBundleAction {
	return AMBundleAction{
		AMAction: AMActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AMBundleActionClass) Alloc() AMBundleAction {
	rv := objc.Send[AMBundleAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AMBundleActionClass) New() AMBundleAction {
	rv := objc.Send[AMBundleAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMBundleAction) Init() AMBundleAction {
	rv := objc.Send[AMBundleAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMBundleAction) Autorelease() AMBundleAction {
	rv := objc.Send[AMBundleAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMBundleAction creates a new AMBundleAction instance.
func NewAMBundleAction() AMBundleAction {
	return getAMBundleActionClass().New()
}




// Allows the action object to perform setup tasks requiring the presence of all bundle objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/awakeFromBundle()

func (a_ AMBundleAction) AwakeFromBundle() {
	objc.Send[objc.ID](a_.ID, objc.Sel("awakeFromBundle"))
}


// The action’s bundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/bundle

func (a_ AMBundleAction) Bundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](a_.ID, objc.Sel("bundle"))
	return rv
}


// A Boolean value that indicates whether the action has a view associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/hasView

func (a_ AMBundleAction) HasView() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hasView"))
	return rv
}


// The action’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/parameters

func (a_ AMBundleAction) Parameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("parameters"))
	return rv
}


// The action’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/parameters

func (a_ AMBundleAction) SetParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParameters:"), value)
}


// The action’s view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMBundleAction/view

func (a_ AMBundleAction) View() appkit.View {
	rv := objc.Send[appkit.View](a_.ID, objc.Sel("view"))
	return rv
}



