// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [JSVirtualMachine] class.
type IJSVirtualMachine interface {
	objectivec.IObject
	// properties:
	// methods:
	AddManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject)
	RemoveManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject)
}

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

// Alloc allocates a new instance without initialization.
func (jc _JSVirtualMachineClass) Alloc() JSVirtualMachine {
	rv := objc.Send[JSVirtualMachine](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Notifies the JavaScriptCore virtual machine of an external object relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/addManagedReference(_:withOwner:)
func (j_ JSVirtualMachine) AddManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("addManagedReference:withOwner:"), object, owner)
}


// Notifies the JavaScriptCore virtual machine that a previously registered object relationship no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSVirtualMachine/removeManagedReference(_:withOwner:)
func (j_ JSVirtualMachine) RemoveManagedReferenceWithOwner(object objectivec.IObject, owner objectivec.IObject) {
	objc.Send[objc.ID](j_.ID, objc.Sel("removeManagedReference:withOwner:"), object, owner)
}


