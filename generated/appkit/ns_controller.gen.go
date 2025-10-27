// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Controller] class.
type IController interface {
	objectivec.IObject
	

	// properties:
	Editing() bool
	IsEditing() bool
	SetIsEditing(value bool)


	

	// methods:
	CommitEditing() bool
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject)
	DiscardEditing()
	ObjectDidBeginEditing(editor unsafe.Pointer)
	ObjectDidEndEditing(editor unsafe.Pointer)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/init(coder:)
func NewControllerWithCoder(coder foundation.foundation.INSCoder) Controller {
	instance := getControllerClass().Alloc()
	rv := objc.Send[Controller](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

















// Attempts to commit any pending edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/commitEditing()
func (c_ Controller) CommitEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("commitEditing"))
	return rv
}


// Attempts to commit any pending changes in known editors of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/commitEditing(withDelegate:didCommit:contextInfo:)
func (c_ Controller) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}


// Discards any pending changes by registered editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/discardEditing()
func (c_ Controller) DiscardEditing() {
	objc.Send[objc.ID](c_.ID, objc.Sel("discardEditing"))
}


// Invoked to inform the receiver that has uncommitted changes that can affect the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/objectDidBeginEditing(_:)
func (c_ Controller) ObjectDidBeginEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("objectDidBeginEditing:"), editor)
}


// Invoked to inform the receiver that has committed or discarded its changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/objectDidEndEditing(_:)
func (c_ Controller) ObjectDidEndEditing(editor unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("objectDidEndEditing:"), editor)
}







// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController/isEditing
func (c_ Controller) Editing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("editing"))
	return rv
}


// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting
func (c_ Controller) IsEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditing"))
	return rv
}


// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting
func (c_ Controller) SetIsEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditing:"), value)
}







