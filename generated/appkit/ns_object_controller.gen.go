// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ObjectController] class.
var (
	ObjectControllerClass     _ObjectControllerClass
	ObjectControllerClassOnce sync.Once
)

func getObjectControllerClass() _ObjectControllerClass {
	ObjectControllerClassOnce.Do(func() {
		ObjectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}
	})
	return ObjectControllerClass
}

type _ObjectControllerClass struct {
	class objc.Class
}

// An interface definition for the [ObjectController] class.
type IObjectController interface {
	IController
	PrepareContent()
}

// A controller that can manage an object’s properties referenced by key-value paths.
//
// is a Cocoa bindings–compatible controller class. Properties of the content object of instances of this class can be bound to user interface elements to access and modify their values. By default, the content of an instance is an object. This allows a single instance to be used to manage many different properties referenced by key-value paths. The default content object class can be changed by calling , which subclasses must override. Your application should use a custom data class that is key-value compliant whenever possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController
type ObjectController struct {
	Controller
}

// ObjectControllerFrom constructs a [ObjectController] from an unsafe.Pointer.
//
// A controller that can manage an object’s properties referenced by key-value paths.
func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		Controller: ControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ObjectController) Init() ObjectController {
	rv := objc.Send[ObjectController](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ObjectController) Autorelease() ObjectController {
	rv := objc.Send[ObjectController](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObjectController creates a new ObjectController instance.
func NewObjectController() ObjectController {
	return getObjectControllerClass().New()
}


// Typically overridden by subclasses that require additional control over the creation of new objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/prepareContent()
func (o_ ObjectController) PrepareContent() {
	objc.Send[objc.ID](o_.ID, objc.Sel("prepareContent"))
}

// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/automaticallyPreparesContent
func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("automaticallyPreparesContent"))
	return rv
}


// SetAutomaticallyPreparesContent sets the value of the automaticallyPreparesContent property.
// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/automaticallyPreparesContent
func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutomaticallyPreparesContent:"), value)
}
// The receiver’s content object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/content
func (o_ ObjectController) Content() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("content"))
	return rv
}


// SetContent sets the value of the content property.
// The receiver’s content object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/content
func (o_ ObjectController) SetContent(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setContent:"), value)
}
// The object class to use when creating new objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/objectClass
func (o_ ObjectController) ObjectClass() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("objectClass"))
	return rv
}


// SetObjectClass sets the value of the objectClass property.
// The object class to use when creating new objects.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/objectClass
func (o_ ObjectController) SetObjectClass(value objc.Class) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setObjectClass:"), value)
}


