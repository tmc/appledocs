// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	AutomaticallyPreparesContent() bool /* primitive/slice/pointer. */
	SetAutomaticallyPreparesContent(value bool /* primitive/slice/pointer. */)
	CanAdd() bool /* primitive/slice/pointer. */
	SetCanAdd(value bool /* primitive/slice/pointer. */)
	CanRemove() bool /* primitive/slice/pointer. */
	SetCanRemove(value bool /* primitive/slice/pointer. */)
	Content() unsafe.Pointer
	SetContent(value unsafe.Pointer)
	EntityName() objc.IObject /* cross-framework: NSString */
	SetEntityName(value objc.IObject /* cross-framework: NSString */)
	FetchPredicate() objc.IObject /* cross-framework: Predicate */
	SetFetchPredicate(value objc.IObject /* cross-framework: Predicate */)
	IsEditable() bool /* primitive/slice/pointer. */
	SetIsEditable(value bool /* primitive/slice/pointer. */)
	ManagedObjectContext() objc.IObject /* cross-framework: ManagedObjectContext */
	SetManagedObjectContext(value objc.IObject /* cross-framework: ManagedObjectContext */)
	ObjectClass() unsafe.Pointer
	SetObjectClass(value unsafe.Pointer)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	Selection() unsafe.Pointer
	SetSelection(value unsafe.Pointer)
	UsesLazyFetching() bool /* primitive/slice/pointer. */
	SetUsesLazyFetching(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A controller that can manage an object’s properties referenced by key-value paths.
//
// is a Cocoa bindings–compatible controller class. Properties of the content object of instances of this class can be bound to user interface elements to access and modify their values. By default, the content of an instance is an object. This allows a single instance to be used to manage many different properties referenced by key-value paths. The default content object class can be changed by calling , which subclasses must override. Your application should use a custom data class that is key-value compliant whenever possible.


// A controller that can manage an object’s properties referenced by key-value paths.
//
// [Full Topic]
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



// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/automaticallypreparescontent
func (o_ ObjectController) AutomaticallyPreparesContent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("automaticallyPreparesContent"))
	return rv
}


// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/automaticallypreparescontent
func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutomaticallyPreparesContent:"), value)
}


// A Boolean value that indicates whether an object can be added to the receiver using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canadd
func (o_ ObjectController) CanAdd() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canAdd"))
	return rv
}


// A Boolean value that indicates whether an object can be added to the receiver using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canadd
func (o_ ObjectController) SetCanAdd(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanAdd:"), value)
}


// A Boolean value that indicates whether an object can be removed from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canremove
func (o_ ObjectController) CanRemove() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canRemove"))
	return rv
}


// A Boolean value that indicates whether an object can be removed from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/canremove
func (o_ ObjectController) SetCanRemove(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanRemove:"), value)
}


// The receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/content
func (o_ ObjectController) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("content"))
	return rv
}


// The receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/content
func (o_ ObjectController) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setContent:"), value)
}


// The entity name used by the receiver to create new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/entityname
func (o_ ObjectController) EntityName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("entityName"))
	return rv
}


// The entity name used by the receiver to create new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/entityname
func (o_ ObjectController) SetEntityName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEntityName:"), value)
}


// The receiver’s fetch predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/fetchpredicate
func (o_ ObjectController) FetchPredicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](o_.ID, objc.Sel("fetchPredicate"))
	return rv
}


// The receiver’s fetch predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/fetchpredicate
func (o_ ObjectController) SetFetchPredicate(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFetchPredicate:"), value)
}


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) IsEditable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) SetIsEditable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsEditable:"), value)
}


// The receiver’s managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/managedobjectcontext
func (o_ ObjectController) ManagedObjectContext() objc.IObject /* cross-framework: ManagedObjectContext */ {
	rv := objc.Send[ManagedObjectContext](o_.ID, objc.Sel("managedObjectContext"))
	return rv
}


// The receiver’s managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/managedobjectcontext
func (o_ ObjectController) SetManagedObjectContext(value objc.IObject /* cross-framework: ManagedObjectContext */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManagedObjectContext:"), value)
}


// The object class to use when creating new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/objectclass
func (o_ ObjectController) ObjectClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("objectClass"))
	return rv
}


// The object class to use when creating new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/objectclass
func (o_ ObjectController) SetObjectClass(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setObjectClass:"), value)
}


// An array of all objects to be affected by editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selectedobjects
func (o_ ObjectController) SelectedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("selectedObjects"))
	return rv
}


// An array of all objects to be affected by editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selectedobjects
func (o_ ObjectController) SetSelectedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSelectedObjects:"), value)
}


// A proxy object representing the receiver’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selection
func (o_ ObjectController) Selection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("selection"))
	return rv
}


// A proxy object representing the receiver’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/selection
func (o_ ObjectController) SetSelection(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSelection:"), value)
}


// A Boolean that indicates whether the receiver uses lazy fetching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/useslazyfetching
func (o_ ObjectController) UsesLazyFetching() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("usesLazyFetching"))
	return rv
}


// A Boolean that indicates whether the receiver uses lazy fetching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/useslazyfetching
func (o_ ObjectController) SetUsesLazyFetching(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUsesLazyFetching:"), value)
}



