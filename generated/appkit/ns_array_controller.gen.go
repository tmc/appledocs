// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ArrayController] class.
var (
	ArrayControllerClass     _ArrayControllerClass
	ArrayControllerClassOnce sync.Once
)

func getArrayControllerClass() _ArrayControllerClass {
	ArrayControllerClassOnce.Do(func() {
		ArrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}
	})
	return ArrayControllerClass
}

type _ArrayControllerClass struct {
	class objc.Class
}





// An interface definition for the [ArrayController] class.
type IArrayController interface {
	IObjectController
	

	// properties:
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	ArrangedObjects() objc.ID
	AutomaticallyRearrangesObjects() bool
	SetAutomaticallyRearrangesObjects(value bool)
	AutomaticRearrangementKeyPaths() []string
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	CanInsert() bool
	CanSelectNext() bool
	CanSelectPrevious() bool
	ClearsFilterPredicateOnInsertion() bool
	SetClearsFilterPredicateOnInsertion(value bool)
	FilterPredicate() foundation.Predicate
	SetFilterPredicate(value foundation.Predicate)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedObjects() foundation.foundation.INSArray
	SelectionIndex() uint
	SelectionIndexes() foundation.IndexSet
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.SortDescriptor)


	

	// methods:
	Add(sender objectivec.IObject)
	AddObjects(objects foundation.foundation.INSArray)
	AddObject(object objectivec.IObject)
	AddSelectedObjects(objects foundation.foundation.INSArray) bool
	AddSelectionIndexes(indexes foundation.IndexSet) bool
	ArrangeObjects(objects foundation.foundation.INSArray) foundation.Array
	DidChangeArrangementCriteria()
	Insert(sender objectivec.IObject)
	InsertObjectAtArrangedObjectIndex(object objectivec.IObject, index uint)
	InsertObjectsAtArrangedObjectIndexes(objects foundation.foundation.INSArray, indexes foundation.IndexSet)
	RearrangeObjects()
	Remove(sender objectivec.IObject)
	RemoveObjectAtArrangedObjectIndex(index uint)
	RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IndexSet)
	RemoveObjects(objects foundation.foundation.INSArray)
	RemoveObject(object objectivec.IObject)
	RemoveSelectedObjects(objects foundation.foundation.INSArray) bool
	RemoveSelectionIndexes(indexes foundation.IndexSet) bool
	SelectNext(sender objectivec.IObject)
	SelectPrevious(sender objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArrayController) Init() ArrayController {
	rv := objc.Send[ArrayController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArrayController) Autorelease() ArrayController {
	rv := objc.Send[ArrayController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArrayController creates a new ArrayController instance.
func NewArrayController() ArrayController {
	return getArrayControllerClass().New()
}





// A bindings-compatible controller that manages a collection of objects.
//
// Typically the collection that an manages is an array, however, if the controller manages a relationship of a managed object (see ) the collection may be a set. provides selection management and sorting capabilities.


// A bindings-compatible controller that manages a collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController
type ArrayController struct {
	ObjectController
}

// ArrayControllerFrom constructs a [ArrayController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a collection of objects.
func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}




















// Creates and adds a new object to the receiver’s content and arranged objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/add(_:)
func (a_ ArrayController) Add(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("add:"), sender)
}


// Adds to the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/add(contentsOf:)
func (a_ ArrayController) AddObjects(objects foundation.foundation.INSArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObjects:"), objects)
}


// Adds to the receiver’s content collection and the arranged objects array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/addObject(_:)
func (a_ ArrayController) AddObject(object objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addObject:"), object)
}


// Adds the specified objects from the receiver’s content array to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/addSelectedObjects(_:)
func (a_ ArrayController) AddSelectedObjects(objects foundation.foundation.INSArray) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addSelectedObjects:"), objects)
	return rv
}


// Adds the objects at the specified indexes in the receiver’s content array to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/addSelectionIndexes(_:)
func (a_ ArrayController) AddSelectionIndexes(indexes foundation.IndexSet) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("addSelectionIndexes:"), indexes)
	return rv
}


// Returns a given array, appropriately sorted and filtered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/arrange(_:)
func (a_ ArrayController) ArrangeObjects(objects foundation.foundation.INSArray) foundation.Array {
	rv := objc.Send[foundation.Array](a_.ID, objc.Sel("arrangeObjects:"), objects)
	return rv
}


// Invoked when any criteria for arranging objects change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/didChangeArrangementCriteria()
func (a_ ArrayController) DidChangeArrangementCriteria() {
	objc.Send[objc.ID](a_.ID, objc.Sel("didChangeArrangementCriteria"))
}


// Creates a new object and inserts it into the receiver’s content array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(_:)
func (a_ ArrayController) Insert(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insert:"), sender)
}


// Inserts into the receiver’s arranged objects array at the location specified by , and adds it to the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(_:atArrangedObjectIndex:)
func (a_ ArrayController) InsertObjectAtArrangedObjectIndex(object objectivec.IObject, index uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertObject:atArrangedObjectIndex:"), object, index)
}


// Inserts s into the receiver’s arranged objects array at the locations specified in , and adds it to the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/insert(contentsOf:atArrangedObjectIndexes:)
func (a_ ArrayController) InsertObjectsAtArrangedObjectIndexes(objects foundation.foundation.INSArray, indexes foundation.IndexSet) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertObjects:atArrangedObjectIndexes:"), objects, indexes)
}


// Triggers filtering of the receiver’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/rearrangeObjects()
func (a_ ArrayController) RearrangeObjects() {
	objc.Send[objc.ID](a_.ID, objc.Sel("rearrangeObjects"))
}


// Removes the receiver’s selected objects from the content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(_:)
func (a_ ArrayController) Remove(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("remove:"), sender)
}


// Removes the object at the specified in the receiver’s arranged objects from the receiver’s content array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(atArrangedObjectIndex:)
func (a_ ArrayController) RemoveObjectAtArrangedObjectIndex(index uint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObjectAtArrangedObjectIndex:"), index)
}


// Removes the objects at the specified in the receiver’s arranged objects from the content array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(atArrangedObjectIndexes:)
func (a_ ArrayController) RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IndexSet) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObjectsAtArrangedObjectIndexes:"), indexes)
}


// Removes from the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/remove(contentsOf:)
func (a_ ArrayController) RemoveObjects(objects foundation.foundation.INSArray) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObjects:"), objects)
}


// Removes from the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/removeObject(_:)
func (a_ ArrayController) RemoveObject(object objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeObject:"), object)
}


// Removes the specified objects from the receiver’s current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/removeSelectedObjects(_:)
func (a_ ArrayController) RemoveSelectedObjects(objects foundation.foundation.INSArray) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeSelectedObjects:"), objects)
	return rv
}


// Removes the object as the specified indexes from the receiver’s current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/removeSelectionIndexes(_:)
func (a_ ArrayController) RemoveSelectionIndexes(indexes foundation.IndexSet) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removeSelectionIndexes:"), indexes)
	return rv
}


// Selects the next object, relative to the current selection, in the receiver’s arranged content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectNext(_:)
func (a_ ArrayController) SelectNext(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectNext:"), sender)
}


// Selects the previous object, relative to the current selection, in the receiver’s arranged content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectPrevious(_:)
func (a_ ArrayController) SelectPrevious(sender objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("selectPrevious:"), sender)
}







// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/alwaysUsesMultipleValuesMarker
func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/alwaysUsesMultipleValuesMarker
func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}


// An array containing the receiver’s content objects arranged using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/arrangedObjects
func (a_ ArrayController) ArrangedObjects() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticallyRearrangesObjects
func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticallyRearrangesObjects
func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}


// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticRearrangementKeyPaths
func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/avoidsEmptySelection
func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/avoidsEmptySelection
func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/canInsert
func (a_ ArrayController) CanInsert() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canInsert"))
	return rv
}


// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/canSelectNext
func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectNext"))
	return rv
}


// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/canSelectPrevious
func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectPrevious"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/clearsFilterPredicateOnInsertion
func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/clearsFilterPredicateOnInsertion
func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}


// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/filterPredicate
func (a_ ArrayController) FilterPredicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](a_.ID, objc.Sel("filterPredicate"))
	return rv
}


// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/filterPredicate
func (a_ ArrayController) SetFilterPredicate(value foundation.Predicate) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterPredicate:"), value)
}


// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/preservesSelection
func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/preservesSelection
func (a_ ArrayController) SetPreservesSelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreservesSelection:"), value)
}


// An array containing the receiver’s selected objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectedObjects
func (a_ ArrayController) SelectedObjects() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("selectedObjects"))
	return rv
}


// The index of the first object in the receiver’s selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectionIndex
func (a_ ArrayController) SelectionIndex() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("selectionIndex"))
	return rv
}


// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectionIndexes
func (a_ ArrayController) SelectionIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](a_.ID, objc.Sel("selectionIndexes"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectsInsertedObjects
func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/selectsInsertedObjects
func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/sortDescriptors
func (a_ ArrayController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Send[[]foundation.SortDescriptor](a_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/sortDescriptors
func (a_ ArrayController) SetSortDescriptors(value []foundation.SortDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}








