// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ArrangedObjects() unsafe.Pointer
	SetArrangedObjects(value unsafe.Pointer)
	AutomaticRearrangementKeyPaths() objc.IObject /* cross-framework: NSString */
	SetAutomaticRearrangementKeyPaths(value objc.IObject /* cross-framework: NSString */)
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
	FilterPredicate() objc.IObject /* cross-framework: Predicate */
	SetFilterPredicate(value objc.IObject /* cross-framework: Predicate */)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	SelectionIndex() int
	SetSelectionIndex(value int)
	SelectionIndexes() objc.IObject /* cross-framework: IndexSet */
	SetSelectionIndexes(value objc.IObject /* cross-framework: IndexSet */)
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() objectivec.IObject
	SetSortDescriptors(value objectivec.IObject)
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
func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the receiver always returns the multiple values marker when multiple objects are selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/alwaysusesmultiplevaluesmarker
func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
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
func (a_ ArrayController) AutomaticRearrangementKeyPaths() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}


// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticrearrangementkeypaths
func (a_ ArrayController) SetAutomaticRearrangementKeyPaths(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticRearrangementKeyPaths:"), value)
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("automaticallyRearrangesObjects"))
	return rv
}


// A Boolean that indicates if the receiver automatically rearranges its content to correspond to the current sort descriptors and filter predicates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/automaticallyrearrangesobjects
func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutomaticallyRearrangesObjects:"), value)
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the receiver requires that the content array attempt to maintain a selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/avoidsemptyselection
func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) CanInsert() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canInsert"))
	return rv
}


// Returns a Boolean value that indicates whether an object can be inserted into the receiver’s content collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/caninsert
func (a_ ArrayController) SetCanInsert(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanInsert:"), value)
}


// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectNext"))
	return rv
}


// A Boolean value indicating whether the next object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectnext
func (a_ ArrayController) SetCanSelectNext(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectNext:"), value)
}


// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSelectPrevious"))
	return rv
}


// A Boolean value indicating whether the previous object, relative to the current selection, in the receiver’s content array can be selected
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/canselectprevious
func (a_ ArrayController) SetCanSelectPrevious(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanSelectPrevious:"), value)
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("clearsFilterPredicateOnInsertion"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically clears an existing filter predicate when new items are inserted or added to the content
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/clearsfilterpredicateoninsertion
func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setClearsFilterPredicateOnInsertion:"), value)
}


// A predicate used by the receiver to filter the array controller contents
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/filterpredicate
func (a_ ArrayController) FilterPredicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[foundation.Predicate](a_.ID, objc.Sel("filterPredicate"))
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
func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the receiver will attempt to preserve the current selection when the content changes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/preservesselection
func (a_ ArrayController) SetPreservesSelection(value bool) {
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
func (a_ ArrayController) SelectionIndex() int {
	rv := objc.Send[int](a_.ID, objc.Sel("selectionIndex"))
	return rv
}


// The index of the first object in the receiver’s selection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindex
func (a_ ArrayController) SetSelectionIndex(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndex:"), value)
}


// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SelectionIndexes() objc.IObject /* cross-framework: IndexSet */ {
	rv := objc.Send[foundation.IndexSet](a_.ID, objc.Sel("selectionIndexes"))
	return rv
}


// An index set containing the indexes of the receiver’s currently selected objects in the content array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectionindexes
func (a_ ArrayController) SetSelectionIndexes(value objc.IObject /* cross-framework: IndexSet */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectionIndexes:"), value)
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the receiver automatically selects inserted objects
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/selectsinsertedobjects
func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SortDescriptors() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array of sort descriptor objects, used by the receiver to arrange its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsarraycontroller/sortdescriptors
func (a_ ArrayController) SetSortDescriptors(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSortDescriptors:"), value)
}



