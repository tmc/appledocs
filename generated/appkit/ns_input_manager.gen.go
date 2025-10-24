// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInputManager */


/* debug [class_header]: Header for NSInputManager */
// The class instance for the [InputManager] class.
var (
	InputManagerClass     _InputManagerClass
	InputManagerClassOnce sync.Once
)

func getInputManagerClass() _InputManagerClass {
	InputManagerClassOnce.Do(func() {
		InputManagerClass = _InputManagerClass{objc.GetClass("NSInputManager")}
	})
	return InputManagerClass
}

type _InputManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InputManager */
// An interface definition for the [InputManager] class.
type IInputManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for InputManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InputManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InputManager */
// Alloc allocates a new instance without initialization.
func (ic _InputManagerClass) Alloc() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InputManagerClass) New() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputManager) Init() InputManager {
	rv := objc.Send[InputManager](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputManager) Autorelease() InputManager {
	rv := objc.Send[InputManager](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputManager creates a new InputManager instance.
func NewInputManager() InputManager {
	return getInputManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InputManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager
type InputManager struct {
	objectivec.Object
}

// InputManagerFrom constructs a [InputManager] from an unsafe.Pointer.
func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InputManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager/initWithName:host:
func NewInputManagerWithNameHost(inputServerName objc.IObject /* cross-framework: NSString */, hostName objc.IObject /* cross-framework: NSString */) InputManager {
	instance := getInputManagerClass().Alloc()
	rv := objc.Send[InputManager](instance.ID, objc.Sel("initWithName:host:"), inputServerName, hostName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInputManagerWithNameHost */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InputManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager/currentInputManager
func (ic _InputManagerClass) CurrentInputManager() IInputManager {
	rv := objc.Send[InputManager](objc.ID(ic.class), objc.Sel("currentInputManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentInputManager) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager/cycleToNextInputLanguage:
func (ic _InputManagerClass) CycleToNextInputLanguage(sender objc.IObject) {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("cycleToNextInputLanguage:"), sender)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CycleToNextInputLanguage) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager/cycleToNextInputServerInLanguage:
func (ic _InputManagerClass) CycleToNextInputServerInLanguage(sender objc.IObject) {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("cycleToNextInputServerInLanguage:"), sender)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CycleToNextInputServerInLanguage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InputManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InputManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InputManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInputManager */


