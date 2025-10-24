// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coredata"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSObjectController */


/* debug [class_header]: Header for NSObjectController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ObjectController */
// An interface definition for the [ObjectController] class.
type IObjectController interface {
	IController
	
/* debug [class_interface_properties]: Properties for ObjectController */
	// properties:
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	CanAdd() bool
	CanRemove() bool
	Content() objc.ID
	SetContent(value objc.ID)
	EntityName() objc.IObject /* cross-framework: NSString */
	SetEntityName(value objc.IObject /* cross-framework: NSString */)
	FetchPredicate() foundation.Predicate
	SetFetchPredicate(value foundation.Predicate)
	Editable() bool
	SetEditable(value bool)
	ManagedObjectContext() coredata.ManagedObjectContext
	SetManagedObjectContext(value coredata.ManagedObjectContext)
	ObjectClass() objc.Class
	SetObjectClass(value objc.Class)
	SelectedObjects() objc.IObject /* cross-framework: NSArray */
	Selection() objc.ID
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ObjectController */
	// methods:
	Add(sender objc.IObject)
	AddObject(object objc.IObject)
	DefaultFetchRequest() coredata.FetchRequest
	Fetch(sender objc.IObject)
	FetchWithRequestMergeError(fetchRequest coredata.FetchRequest, merge bool, error_ objectivec.IObject) bool
	NewObject() objc.ID
	PrepareContent()
	Remove(sender objc.IObject)
	RemoveObject(object objc.IObject)
	ValidateUserInterfaceItem(item unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ObjectController */
// Alloc allocates a new instance without initialization.
func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ObjectController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ObjectController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/init(coder:)
func NewObjectControllerWithCoder(coder foundation.Coder) ObjectController {
	instance := getObjectControllerClass().Alloc()
	rv := objc.Send[ObjectController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewObjectControllerWithCoder */


// Initializes and returns an object with the given content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/init(content:)
func NewObjectControllerWithContent(content objc.IObject) ObjectController {
	instance := getObjectControllerClass().Alloc()
	rv := objc.Send[ObjectController](instance.ID, objc.Sel("initWithContent:"), content)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewObjectControllerWithContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ObjectController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ObjectController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ObjectController */

// Creates a new object and sets it as the receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/add(_:)
func (o_ ObjectController) Add(sender objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("add:"), sender)
}/* debug [instance_methods/method]: Add */


// Sets the receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/addObject(_:)
func (o_ ObjectController) AddObject(object objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObject:"), object)
}/* debug [instance_methods/method]: AddObject */


// Returns the default fetch request used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/defaultFetchRequest()
func (o_ ObjectController) DefaultFetchRequest() coredata.FetchRequest {
	rv := objc.Send[coredata.FetchRequest](o_.ID, objc.Sel("defaultFetchRequest"))
	return rv
}/* debug [instance_methods/method]: DefaultFetchRequest */


// Causes the receiver to fetch the data objects specified by the entity name and fetch predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/fetch(_:)
func (o_ ObjectController) Fetch(sender objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fetch:"), sender)
}/* debug [instance_methods/method]: Fetch */


// Subclasses should override this method to customize a fetch request, for example to specify fetch limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/fetch(with:merge:)
func (o_ ObjectController) FetchWithRequestMergeError(fetchRequest coredata.FetchRequest, merge bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("fetchWithRequest:merge:error:"), fetchRequest, merge, error_)
	return rv
}/* debug [instance_methods/method]: FetchWithRequestMergeError */


// Creates and returns a new object of the appropriate class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/newObject()
func (o_ ObjectController) NewObject() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("newObject"))
	return rv
}/* debug [instance_methods/method]: NewObject */


// Typically overridden by subclasses that require additional control over the creation of new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/prepareContent()
func (o_ ObjectController) PrepareContent() {
	objc.Send[objc.ID](o_.ID, objc.Sel("prepareContent"))
}/* debug [instance_methods/method]: PrepareContent */


// Removes the receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/remove(_:)
func (o_ ObjectController) Remove(sender objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("remove:"), sender)
}/* debug [instance_methods/method]: Remove */


// Removes a given object from the receiver’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/removeObject(_:)
func (o_ ObjectController) RemoveObject(object objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObject:"), object)
}/* debug [instance_methods/method]: RemoveObject */


// Returns whether the receiver can handle the action method for a user interface item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/validateUserInterfaceItem(_:)
func (o_ ObjectController) ValidateUserInterfaceItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}/* debug [instance_methods/method]: ValidateUserInterfaceItem */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ObjectController */

// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/automaticallyPreparesContent
func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("automaticallyPreparesContent"))
	return rv
}/* debug [instance_properties/getter]: automaticallyPreparesContent */


// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/automaticallyPreparesContent
func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutomaticallyPreparesContent:"), value)
}/* debug [instance_properties/setter]: automaticallyPreparesContent */


// A Boolean value that indicates whether an object can be added to the receiver using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/canAdd
func (o_ ObjectController) CanAdd() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canAdd"))
	return rv
}/* debug [instance_properties/getter]: canAdd */


// A Boolean value that indicates whether an object can be removed from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/canRemove
func (o_ ObjectController) CanRemove() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canRemove"))
	return rv
}/* debug [instance_properties/getter]: canRemove */


// The receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/content
func (o_ ObjectController) Content() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */


// The receiver’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/content
func (o_ ObjectController) SetContent(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setContent:"), value)
}/* debug [instance_properties/setter]: content */


// The entity name used by the receiver to create new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/entityName
func (o_ ObjectController) EntityName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("entityName"))
	return rv
}/* debug [instance_properties/getter]: entityName */


// The entity name used by the receiver to create new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/entityName
func (o_ ObjectController) SetEntityName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEntityName:"), value)
}/* debug [instance_properties/setter]: entityName */


// The receiver’s fetch predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/fetchPredicate
func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](o_.ID, objc.Sel("fetchPredicate"))
	return rv
}/* debug [instance_properties/getter]: fetchPredicate */


// The receiver’s fetch predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/fetchPredicate
func (o_ ObjectController) SetFetchPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFetchPredicate:"), value)
}/* debug [instance_properties/setter]: fetchPredicate */


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/isEditable
func (o_ ObjectController) Editable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/isEditable
func (o_ ObjectController) SetEditable(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// The receiver’s managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/managedObjectContext
func (o_ ObjectController) ManagedObjectContext() coredata.ManagedObjectContext {
	rv := objc.Send[coredata.ManagedObjectContext](o_.ID, objc.Sel("managedObjectContext"))
	return rv
}/* debug [instance_properties/getter]: managedObjectContext */


// The receiver’s managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/managedObjectContext
func (o_ ObjectController) SetManagedObjectContext(value coredata.ManagedObjectContext) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManagedObjectContext:"), value)
}/* debug [instance_properties/setter]: managedObjectContext */


// The object class to use when creating new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/objectClass
func (o_ ObjectController) ObjectClass() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("objectClass"))
	return rv
}/* debug [instance_properties/getter]: objectClass */


// The object class to use when creating new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/objectClass
func (o_ ObjectController) SetObjectClass(value objc.Class) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setObjectClass:"), value)
}/* debug [instance_properties/setter]: objectClass */


// An array of all objects to be affected by editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/selectedObjects
func (o_ ObjectController) SelectedObjects() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("selectedObjects"))
	return rv
}/* debug [instance_properties/getter]: selectedObjects */


// A proxy object representing the receiver’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/selection
func (o_ ObjectController) Selection() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("selection"))
	return rv
}/* debug [instance_properties/getter]: selection */


// A Boolean that indicates whether the receiver uses lazy fetching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/usesLazyFetching
func (o_ ObjectController) UsesLazyFetching() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("usesLazyFetching"))
	return rv
}/* debug [instance_properties/getter]: usesLazyFetching */


// A Boolean that indicates whether the receiver uses lazy fetching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController/usesLazyFetching
func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUsesLazyFetching:"), value)
}/* debug [instance_properties/setter]: usesLazyFetching */


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) IsEditable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean that indicates whether the receiver allows adding and removing objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/iseditable
func (o_ ObjectController) SetIsEditable(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSObjectController */


