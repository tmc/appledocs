// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AutomaticRearrangementKeyPaths() []string
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	ArrangedObjects() unsafe.Pointer
	SetArrangedObjects(value unsafe.Pointer)
	AutomaticallyRearrangesObjects() bool
	SetAutomaticallyRearrangesObjects(value bool)
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	CanInsert() bool
	SetCanInsert(value bool)
	CanSelectNext() bool
	SetCanSelectNext(value bool)
	CanSelectPrevious() bool
	SetCanSelectPrevious(value bool)
	ClearsFilterPredicateOnInsertion() bool
	SetClearsFilterPredicateOnInsertion(value bool)
	FilterPredicate() foundation.Predicate
	SetFilterPredicate(value foundation.IPredicate)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	SelectionIndex() int
	SetSelectionIndex(value int)
	SelectionIndexes() foundation.IndexSet
	SetSelectionIndexes(value foundation.IIndexSet)
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() foundation.SortDescriptor
	SetSortDescriptors(value foundation.ISortDescriptor)
}

// A bindings-compatible controller that manages a collection of objects.
//
// Typically the collection that an manages is an array, however, if the controller manages a relationship of a managed object (see ) the collection may be a set. provides selection management and sorting capabilities.
//
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

// Alloc allocates a new instance without initialization.
func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticRearrangementKeyPaths
func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}

// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/alwaysusesmultiplevaluesmarker
func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// SetAlwaysUsesMultipleValuesMarker sets the value of the alwaysUsesMultipleValuesMarker property.
// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/alwaysusesmultiplevaluesmarker
func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}

// An array containing the receiver’s content objects arranged using
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/arrangedobjects
func (a_ ArrayController) ArrangedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// SetArrangedObjects sets the value of the arrangedObjects property.
// An array containing the receiver’s content objects arranged using

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/arrangedobjects
func (a_ ArrayController) SetArrangedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArrangedObjects:"), value)
}

// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}


// SetAutomaticallyRearrangesObjects sets the value of the automaticallyRearrangesObjects property.
// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}

// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// SetAvoidsEmptySelection sets the value of the avoidsEmptySelection property.
// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}

// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) CanInsert() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canInsert"))
	return rv
}


// SetCanInsert sets the value of the canInsert property.
// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) SetCanInsert(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanInsert:"), value)
}

// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectNext"))
	return rv
}


// SetCanSelectNext sets the value of the canSelectNext property.
// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) SetCanSelectNext(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectNext:"), value)
}

// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectPrevious"))
	return rv
}


// SetCanSelectPrevious sets the value of the canSelectPrevious property.
// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) SetCanSelectPrevious(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectPrevious:"), value)
}

// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}


// SetClearsFilterPredicateOnInsertion sets the value of the clearsFilterPredicateOnInsertion property.
// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}

// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/filterpredicate
func (a_ ArrayController) FilterPredicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](a_.ID, objc.Sel("filterPredicate"))
	return rv
}


// SetFilterPredicate sets the value of the filterPredicate property.
// A predicate used by the receiver to filter the array controller contents

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/filterpredicate
func (a_ ArrayController) SetFilterPredicate(value foundation.IPredicate) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterPredicate:"), value)
}

// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/preservesselection
func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesSelection"))
	return rv
}


// SetPreservesSelection sets the value of the preservesSelection property.
// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/preservesselection
func (a_ ArrayController) SetPreservesSelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreservesSelection:"), value)
}

// An array containing the receiver’s selected objects
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectedobjects
func (a_ ArrayController) SelectedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("selectedObjects"))
	return rv
}


// SetSelectedObjects sets the value of the selectedObjects property.
// An array containing the receiver’s selected objects

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectedobjects
func (a_ ArrayController) SetSelectedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedObjects:"), value)
}

// The index of the first object in the receiver’s selection
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindex
func (a_ ArrayController) SelectionIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("selectionIndex"))
	return rv
}


// SetSelectionIndex sets the value of the selectionIndex property.
// The index of the first object in the receiver’s selection

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindex
func (a_ ArrayController) SetSelectionIndex(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndex:"), value)
}

// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SelectionIndexes() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](a_.ID, objc.Sel("selectionIndexes"))
	return rv
}


// SetSelectionIndexes sets the value of the selectionIndexes property.
// An index set containing the indexes of the receiver’s currently selected objects in the content array

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SetSelectionIndexes(value foundation.IIndexSet) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndexes:"), value)
}

// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// SetSelectsInsertedObjects sets the value of the selectsInsertedObjects property.
// A Boolean value that indicates whether the receiver automatically selects inserted objects

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}

// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SortDescriptors() foundation.SortDescriptor {
	rv := objc.Send[foundation.SortDescriptor](a_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// An array of sort descriptor objects, used by the receiver to arrange its content.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SetSortDescriptors(value foundation.ISortDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSortDescriptors:"), value)
}



