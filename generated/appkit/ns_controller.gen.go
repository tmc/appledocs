// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSController */


/* debug [class_header]: Header for NSController */
// The class instance for the [Controller] class.
var (
	ControllerClass     _ControllerClass
	ControllerClassOnce sync.Once
)

func getControllerClass() _ControllerClass {
	ControllerClassOnce.Do(func() {
		ControllerClass = _ControllerClass{objc.GetClass("NSController")}
	})
	return ControllerClass
}

type _ControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Controller */
// An interface definition for the [Controller] class.
type IController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Controller */
	// properties:
	Editing() bool
	IsEditing() bool
	SetIsEditing(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Controller */
	// methods:
	CommitEditing() bool
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject)
	DiscardEditing()
	ObjectDidBeginEditing(editor unsafe.Pointer)
	ObjectDidEndEditing(editor unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Controller */
// Alloc allocates a new instance without initialization.
func (cc _ControllerClass) Alloc() Controller {
	rv := objc.Send[Controller](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ControllerClass) New() Controller {
	rv := objc.Send[Controller](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Controller) Init() Controller {
	rv := objc.Send[Controller](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Controller) Autorelease() Controller {
	rv := objc.Send[Controller](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewController creates a new Controller instance.
func NewController() Controller {
	return getControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Controller */
// An abstract class that implements the and informal protocols required for controller classes.


// An abstract class that implements the and informal protocols required for controller classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController
type Controller struct {
	objectivec.Object
}

// ControllerFrom constructs a [Controller] from an unsafe.Pointer.
//
// An abstract class that implements the and informal protocols required for controller classes.
func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Controller */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/init(coder:)
func NewControllerWithCoder(coder foundation.Coder) Controller {
	instance := getControllerClass().Alloc()
	rv := objc.Send[Controller](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewControllerWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Controller */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Controller */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Controller */

// Attempts to commit any pending edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/commitEditing()
func (c_ Controller) CommitEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("commitEditing"))
	return rv
}/* debug [instance_methods/method]: CommitEditing */


// Attempts to commit any pending changes in known editors of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/commitEditing(withDelegate:didCommit:contextInfo:)
func (c_ Controller) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objc.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}/* debug [instance_methods/method]: CommitEditingWithDelegateDidCommitSelectorContextInfo */


// Discards any pending changes by registered editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/discardEditing()
func (c_ Controller) DiscardEditing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("discardEditing"))
}/* debug [instance_methods/method]: DiscardEditing */


// Invoked to inform the receiver that has uncommitted changes that can affect the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/objectDidBeginEditing(_:)
func (c_ Controller) ObjectDidBeginEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("objectDidBeginEditing:"), editor)
}/* debug [instance_methods/method]: ObjectDidBeginEditing */


// Invoked to inform the receiver that has committed or discarded its changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/objectDidEndEditing(_:)
func (c_ Controller) ObjectDidEndEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("objectDidEndEditing:"), editor)
}/* debug [instance_methods/method]: ObjectDidEndEditing */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Controller */

// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/isEditing
func (c_ Controller) Editing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editing"))
	return rv
}/* debug [instance_properties/getter]: editing */


// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting
func (c_ Controller) IsEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditing"))
	return rv
}/* debug [instance_properties/getter]: isEditing */


// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting
func (c_ Controller) SetIsEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditing:"), value)
}/* debug [instance_properties/setter]: isEditing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSController */


