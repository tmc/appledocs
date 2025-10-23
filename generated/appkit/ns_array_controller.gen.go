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
	// properties:
	AlwaysUsesMultipleValuesMarker() bool /* primitive/slice/pointer. */
	SetAlwaysUsesMultipleValuesMarker(value bool /* primitive/slice/pointer. */)
	ArrangedObjects() unsafe.Pointer
	SetArrangedObjects(value unsafe.Pointer)
	AutomaticRearrangementKeyPaths() string /* primitive/slice/pointer. */
	SetAutomaticRearrangementKeyPaths(value string /* primitive/slice/pointer. */)
	AutomaticallyRearrangesObjects() bool /* primitive/slice/pointer. */
	SetAutomaticallyRearrangesObjects(value bool /* primitive/slice/pointer. */)
	AvoidsEmptySelection() bool /* primitive/slice/pointer. */
	SetAvoidsEmptySelection(value bool /* primitive/slice/pointer. */)
	CanInsert() bool /* primitive/slice/pointer. */
	SetCanInsert(value bool /* primitive/slice/pointer. */)
	CanSelectNext() bool /* primitive/slice/pointer. */
	SetCanSelectNext(value bool /* primitive/slice/pointer. */)
	CanSelectPrevious() bool /* primitive/slice/pointer. */
	SetCanSelectPrevious(value bool /* primitive/slice/pointer. */)
	ClearsFilterPredicateOnInsertion() bool /* primitive/slice/pointer. */
	SetClearsFilterPredicateOnInsertion(value bool /* primitive/slice/pointer. */)
	FilterPredicate() objc.IObject /* cross-framework: Predicate */
	SetFilterPredicate(value objc.IObject /* cross-framework: Predicate */)
	PreservesSelection() bool /* primitive/slice/pointer. */
	SetPreservesSelection(value bool /* primitive/slice/pointer. */)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	SelectionIndex() int /* primitive/slice/pointer. */
	SetSelectionIndex(value int /* primitive/slice/pointer. */)
	SelectionIndexes() foundation.objc.IObject /* cross-framework: IndexSet */
	SetSelectionIndexes(value foundation.objc.IObject /* cross-framework: IndexSet */)
	SelectsInsertedObjects() bool /* primitive/slice/pointer. */
	SetSelectsInsertedObjects(value bool /* primitive/slice/pointer. */)
	SortDescriptors() SortDescriptor /* not a class type */
	SetSortDescriptors(value SortDescriptor /* not a class type */)
	// methods:
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



// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/alwaysusesmultiplevaluesmarker
func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/alwaysusesmultiplevaluesmarker
func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}


// An array containing the receiver’s content objects arranged using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/arrangedobjects
func (a_ ArrayController) ArrangedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// An array containing the receiver’s content objects arranged using
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/arrangedobjects
func (a_ ArrayController) SetArrangedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setArrangedObjects:"), value)
}


// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticrearrangementkeypaths
func (a_ ArrayController) AutomaticRearrangementKeyPaths() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}


// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticrearrangementkeypaths
func (a_ ArrayController) SetAutomaticRearrangementKeyPaths(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticRearrangementKeyPaths:"), objc.String(value))
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) AutomaticallyRearrangesObjects() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) AvoidsEmptySelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) SetAvoidsEmptySelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) CanInsert() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canInsert"))
	return rv
}


// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) SetCanInsert(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanInsert:"), value)
}


// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) CanSelectNext() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectNext"))
	return rv
}


// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) SetCanSelectNext(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectNext:"), value)
}


// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) CanSelectPrevious() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectPrevious"))
	return rv
}


// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) SetCanSelectPrevious(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectPrevious:"), value)
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}


// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/filterpredicate
func (a_ ArrayController) FilterPredicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[Predicate](a_.ID, objc.Sel("filterPredicate"))
	return rv
}


// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/filterpredicate
func (a_ ArrayController) SetFilterPredicate(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterPredicate:"), value)
}


// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/preservesselection
func (a_ ArrayController) PreservesSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/preservesselection
func (a_ ArrayController) SetPreservesSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreservesSelection:"), value)
}


// An array containing the receiver’s selected objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectedobjects
func (a_ ArrayController) SelectedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("selectedObjects"))
	return rv
}


// An array containing the receiver’s selected objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectedobjects
func (a_ ArrayController) SetSelectedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectedObjects:"), value)
}


// The index of the first object in the receiver’s selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindex
func (a_ ArrayController) SelectionIndex() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](a_.ID, objc.Sel("selectionIndex"))
	return rv
}


// The index of the first object in the receiver’s selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindex
func (a_ ArrayController) SetSelectionIndex(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndex:"), value)
}


// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SelectionIndexes() foundation.objc.IObject /* cross-framework: IndexSet */ {
	rv := objc.Send[foundation.IndexSet](a_.ID, objc.Sel("selectionIndexes"))
	return rv
}


// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SetSelectionIndexes(value foundation.objc.IObject /* cross-framework: IndexSet */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndexes:"), value)
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SelectsInsertedObjects() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SetSelectsInsertedObjects(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SortDescriptors() SortDescriptor /* not a class type */ {
	rv := objc.Send[SortDescriptor](a_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SetSortDescriptors(value SortDescriptor /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSortDescriptors:"), value)
}



