// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coredata"
	"github.com/tmc/appledocs/generated/foundation"
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
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	Content() objc.ID
	SetContent(value objc.ID)
	ObjectClass() objc.Class
	SetObjectClass(value objc.Class)
	CanAdd() bool
	SetCanAdd(value bool)
	CanRemove() bool
	SetCanRemove(value bool)
	EntityName() string
	SetEntityName(value string)
	FetchPredicate() foundation.Predicate
	SetFetchPredicate(value foundation.IPredicate)
	IsEditable() bool
	SetIsEditable(value bool)
	ManagedObjectContext() coredata.ManagedObjectContext
	SetManagedObjectContext(value coredata.IManagedObjectContext)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	Selection() unsafe.Pointer
	SetSelection(value unsafe.Pointer)
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
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

// A Boolean value that indicates whether an object can be added to the receiver using
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canadd
func (o_ ObjectController) CanAdd() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canAdd"))
	return rv
}


// SetCanAdd sets the value of the canAdd property.
// A Boolean value that indicates whether an object can be added to the receiver using

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canadd
func (o_ ObjectController) SetCanAdd(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanAdd:"), value)
}

// A Boolean value that indicates whether an object can be removed from the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canremove
func (o_ ObjectController) CanRemove() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canRemove"))
	return rv
}


// SetCanRemove sets the value of the canRemove property.
// A Boolean value that indicates whether an object can be removed from the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canremove
func (o_ ObjectController) SetCanRemove(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanRemove:"), value)
}

// The entity name used by the receiver to create new objects.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/entityname
func (o_ ObjectController) EntityName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("entityName"))
	return rv
}


// SetEntityName sets the value of the entityName property.
// The entity name used by the receiver to create new objects.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/entityname
func (o_ ObjectController) SetEntityName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEntityName:"), objc.String(value))
}

// The receiver’s fetch predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/fetchpredicate
func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](o_.ID, objc.Sel("fetchPredicate"))
	return rv
}


// SetFetchPredicate sets the value of the fetchPredicate property.
// The receiver’s fetch predicate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/fetchpredicate
func (o_ ObjectController) SetFetchPredicate(value foundation.IPredicate) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFetchPredicate:"), value)
}

// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) IsEditable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEditable"))
	return rv
}


// SetIsEditable sets the value of the isEditable property.
// A Boolean that indicates whether the receiver allows adding and removing objects.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) SetIsEditable(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsEditable:"), value)
}

// The receiver’s managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/managedobjectcontext
func (o_ ObjectController) ManagedObjectContext() coredata.ManagedObjectContext {
	rv := objc.Send[coredata.ManagedObjectContext](o_.ID, objc.Sel("managedObjectContext"))
	return rv
}


// SetManagedObjectContext sets the value of the managedObjectContext property.
// The receiver’s managed object context.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/managedobjectcontext
func (o_ ObjectController) SetManagedObjectContext(value coredata.IManagedObjectContext) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManagedObjectContext:"), value)
}

// An array of all objects to be affected by editing.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selectedobjects
func (o_ ObjectController) SelectedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("selectedObjects"))
	return rv
}


// SetSelectedObjects sets the value of the selectedObjects property.
// An array of all objects to be affected by editing.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selectedobjects
func (o_ ObjectController) SetSelectedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSelectedObjects:"), value)
}

// A proxy object representing the receiver’s selection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selection
func (o_ ObjectController) Selection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("selection"))
	return rv
}


// SetSelection sets the value of the selection property.
// A proxy object representing the receiver’s selection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selection
func (o_ ObjectController) SetSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSelection:"), value)
}

// A Boolean that indicates whether the receiver uses lazy fetching.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/useslazyfetching
func (o_ ObjectController) UsesLazyFetching() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("usesLazyFetching"))
	return rv
}


// SetUsesLazyFetching sets the value of the usesLazyFetching property.
// A Boolean that indicates whether the receiver uses lazy fetching.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/useslazyfetching
func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUsesLazyFetching:"), value)
}



