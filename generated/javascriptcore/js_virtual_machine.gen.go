// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class JSVirtualMachine */


/* debug [class_header]: Header for JSVirtualMachine */
// The class instance for the [JSVirtualMachine] class.
var (
	JSVirtualMachineClass     _JSVirtualMachineClass
	JSVirtualMachineClassOnce sync.Once
)

func getJSVirtualMachineClass() _JSVirtualMachineClass {
	JSVirtualMachineClassOnce.Do(func() {
		JSVirtualMachineClass = _JSVirtualMachineClass{objc.GetClass("JSVirtualMachine")}
	})
	return JSVirtualMachineClass
}

type _JSVirtualMachineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for JSVirtualMachine */
// An interface definition for the [JSVirtualMachine] class.
type IJSVirtualMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for JSVirtualMachine */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for JSVirtualMachine */
	// methods:
	AddManagedReferenceWithOwner(object objc.IObject, owner objc.IObject)
	RemoveManagedReferenceWithOwner(object objc.IObject, owner objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for JSVirtualMachine */
// Alloc allocates a new instance without initialization.
func (jc _JSVirtualMachineClass) Alloc() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (jc _JSVirtualMachineClass) New() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](objc.ID(jc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (j_ JSVirtualMachine) Init() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j_ JSVirtualMachine) Autorelease() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](j_.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSVirtualMachine creates a new JSVirtualMachine instance.
func NewJSVirtualMachine() JSVirtualMachine {
	return getJSVirtualMachineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for JSVirtualMachine */
// A self-contained environment for JavaScript execution.
//
// You use this class for two main purposes: to support concurrent JavaScript execution, and to manage memory for objects that bridge between JavaScript and Objective-C or Swift.


// A self-contained environment for JavaScript execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine
type JSVirtualMachine struct {
	objectivec.Object
}

// JSVirtualMachineFrom constructs a [JSVirtualMachine] from an unsafe.Pointer.
//
// A self-contained environment for JavaScript execution.
func JSVirtualMachineFrom(ptr unsafe.Pointer) JSVirtualMachine {
	return JSVirtualMachine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for JSVirtualMachine */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for JSVirtualMachine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for JSVirtualMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for JSVirtualMachine */

// Notifies the JavaScriptCore virtual machine of an external object relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/addManagedReference(_:withOwner:)
func (j_ JSVirtualMachine) AddManagedReferenceWithOwner(object objc.IObject, owner objc.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("addManagedReference:withOwner:"), object, owner)
}/* debug [instance_methods/method]: AddManagedReferenceWithOwner */


// Notifies the JavaScriptCore virtual machine that a previously registered object relationship no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/removeManagedReference(_:withOwner:)
func (j_ JSVirtualMachine) RemoveManagedReferenceWithOwner(object objc.IObject, owner objc.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("removeManagedReference:withOwner:"), object, owner)
}/* debug [instance_methods/method]: RemoveManagedReferenceWithOwner */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for JSVirtualMachine */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class JSVirtualMachine */


