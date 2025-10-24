// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMutableArray */


/* debug [class_header]: Header for NSMutableArray */
// The class instance for the [MutableArray] class.
var (
	MutableArrayClass     _MutableArrayClass
	MutableArrayClassOnce sync.Once
)

func getMutableArrayClass() _MutableArrayClass {
	MutableArrayClassOnce.Do(func() {
		MutableArrayClass = _MutableArrayClass{objc.GetClass("NSMutableArray")}
	})
	return MutableArrayClass
}

type _MutableArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableArray */
// An interface definition for the [MutableArray] class.
type IMutableArray interface {
	IArray
	
/* debug [class_interface_properties]: Properties for MutableArray */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableArray */
	// methods:
	AddObject(anObject objectivec.IObject)
	AddObjectsFromArray(otherArray []objc.ID)
	ApplyDifference(difference unsafe.Pointer)
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	FilterUsingPredicate(predicate IPredicate)
	InsertObjectAtIndex(anObject objectivec.IObject, index uint)
	InsertObjectsAtIndexes(objects []objc.ID, indexes IIndexSet)
	RemoveObject(anObject objectivec.IObject)
	RemoveObjectInRange(anObject objectivec.IObject, range_ objc.IObject /* cross-framework: Range */)
	RemoveAllObjects()
	RemoveLastObject()
	RemoveObjectAtIndex(index uint)
	RemoveObjectIdenticalTo(anObject objectivec.IObject)
	RemoveObjectIdenticalToInRange(anObject objectivec.IObject, range_ objc.IObject /* cross-framework: Range */)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	RemoveObjectsInRange(range_ objc.IObject /* cross-framework: Range */)
	RemoveObjectsInArray(otherArray []objc.ID)
	ReplaceObjectAtIndexWithObject(index uint, anObject objectivec.IObject)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID)
	ReplaceObjectsInRangeWithObjectsFromArray(range_ objc.IObject /* cross-framework: Range */, otherArray []objc.ID)
	ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ objc.IObject /* cross-framework: Range */, otherArray []objc.ID, otherRange objc.IObject /* cross-framework: Range */)
	SetArray(otherArray []objc.ID)
	SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint)
	SortUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject)
	SortUsingComparator(cmptr Comparator /* not a class type */)
	SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */)
	SortUsingDescriptors(sortDescriptors []SortDescriptor)
	SortUsingSelector(comparator objc.SEL)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableArray */
// Alloc allocates a new instance without initialization.
func (mc _MutableArrayClass) Alloc() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableArrayClass) New() MutableArray {
	rv := objc.Send[MutableArray](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableArray) Init() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableArray) Autorelease() MutableArray {
	rv := objc.Send[MutableArray](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableArray creates a new MutableArray instance.
func NewMutableArray() MutableArray {
	return getMutableArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableArray */
// A dynamic ordered collection of objects.
//
// You can use this type in Swift instead of an variable in cases that require reference semantics. The class declares the programmatic interface to objects that manage a modifiable array of objects. This class adds insertion and deletion operations to the basic array-handling behavior inherited from . NSMutableArray is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A dynamic ordered collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray
type MutableArray struct {
	Array
}

// MutableArrayFrom constructs a [MutableArray] from an unsafe.Pointer.
//
// A dynamic ordered collection of objects.
func MutableArrayFrom(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: ArrayFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableArray */

// Returns an array, initialized with enough memory to initially hold a given number of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(capacity:)
func NewMutableArrayWithCapacity(numItems uint) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableArrayWithCapacity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(coder:)
func NewMutableArrayWithCoder(coder ICoder) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableArrayWithCoder */


// Initializes a newly allocated mutable array with the contents of the file specified by a given path
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/initWithContentsOfFile:
func NewMutableArrayWithContentsOfFile(path IString) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithContentsOfFile:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableArrayWithContentsOfFile */


// Initialized a newly allocated mutable array with the contents of the location specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/initWithContentsOfURL:
func NewMutableArrayWithContentsOfURL(url IURL) MutableArray {
	instance := getMutableArrayClass().Alloc()
	rv := objc.Send[MutableArray](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableArrayWithContentsOfURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableArray */

// Creates and returns an object with enough allocated memory to initially hold a given number of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/arrayWithCapacity:
func (mc _MutableArrayClass) ArrayWithCapacity(numItems uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("arrayWithCapacity:"), numItems)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArrayWithCapacity) */


// Creates and returns a mutable array containing the contents of the file specified by the given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/arrayWithContentsOfFile:
func (mc _MutableArrayClass) ArrayWithContentsOfFile(path IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("arrayWithContentsOfFile:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArrayWithContentsOfFile) */


// Creates and returns a mutable array containing the contents specified by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/init(contentsOfURL:)
func (mc _MutableArrayClass) ArrayWithContentsOfURL(url IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("arrayWithContentsOfURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArrayWithContentsOfURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableArray */

// Inserts a given object at the end of the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/add(_:)
func (m_ MutableArray) AddObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObject:"), anObject)
}/* debug [instance_methods/method]: AddObject */


// Adds the objects contained in another given array to the end of the receiving array’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/addObjects(from:)
func (m_ MutableArray) AddObjectsFromArray(otherArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addObjectsFromArray:"), otherArray)
}/* debug [instance_methods/method]: AddObjectsFromArray */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/applyDifference:
func (m_ MutableArray) ApplyDifference(difference unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyDifference:"), difference)
}/* debug [instance_methods/method]: ApplyDifference */


// Exchanges the objects in the array at given indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/exchangeObject(at:withObjectAt:)
func (m_ MutableArray) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}/* debug [instance_methods/method]: ExchangeObjectAtIndexWithObjectAtIndex */


// Evaluates a given predicate against the array’s content and leaves only objects that match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/filter(using:)
func (m_ MutableArray) FilterUsingPredicate(predicate IPredicate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("filterUsingPredicate:"), predicate)
}/* debug [instance_methods/method]: FilterUsingPredicate */


// Inserts a given object into the array’s contents at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-5dbx5
func (m_ MutableArray) InsertObjectAtIndex(anObject objectivec.IObject, index uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObject:atIndex:"), anObject, index)
}/* debug [instance_methods/method]: InsertObjectAtIndex */


// Inserts the objects in the provided array into the receiving array at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-73pln
func (m_ MutableArray) InsertObjectsAtIndexes(objects []objc.ID, indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertObjects:atIndexes:"), objects, indexes)
}/* debug [instance_methods/method]: InsertObjectsAtIndexes */


// Removes all occurrences in the array of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:)
func (m_ MutableArray) RemoveObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:"), anObject)
}/* debug [instance_methods/method]: RemoveObject */


// Removes all occurrences within a specified range in the array of a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/remove(_:in:)
func (m_ MutableArray) RemoveObjectInRange(anObject objectivec.IObject, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObject:inRange:"), anObject, range_)
}/* debug [instance_methods/method]: RemoveObjectInRange */


// Empties the array of all its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeAllObjects()
func (m_ MutableArray) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}/* debug [instance_methods/method]: RemoveAllObjects */


// Removes the object with the highest-valued index in the array
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeLastObject()
func (m_ MutableArray) RemoveLastObject() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeLastObject"))
}/* debug [instance_methods/method]: RemoveLastObject */


// Removes the object at .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(at:)
func (m_ MutableArray) RemoveObjectAtIndex(index uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectAtIndex:"), index)
}/* debug [instance_methods/method]: RemoveObjectAtIndex */


// Removes all occurrences of a given object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:)
func (m_ MutableArray) RemoveObjectIdenticalTo(anObject objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectIdenticalTo:"), anObject)
}/* debug [instance_methods/method]: RemoveObjectIdenticalTo */


// Removes all occurrences of within the specified range in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObject(identicalTo:in:)
func (m_ MutableArray) RemoveObjectIdenticalToInRange(anObject objectivec.IObject, range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectIdenticalTo:inRange:"), anObject, range_)
}/* debug [instance_methods/method]: RemoveObjectIdenticalToInRange */


// Removes the objects at the specified indexes from the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(at:)
func (m_ MutableArray) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsAtIndexes:"), indexes)
}/* debug [instance_methods/method]: RemoveObjectsAtIndexes */


// Removes from the array each of the objects within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-1udmn
func (m_ MutableArray) RemoveObjectsInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInRange:"), range_)
}/* debug [instance_methods/method]: RemoveObjectsInRange */


// Removes from the receiving array the objects in another given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(in:)-4yb26
func (m_ MutableArray) RemoveObjectsInArray(otherArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInArray:"), otherArray)
}/* debug [instance_methods/method]: RemoveObjectsInArray */


// Replaces the object at with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObject(at:with:)
func (m_ MutableArray) ReplaceObjectAtIndexWithObject(index uint, anObject objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectAtIndex:withObject:"), index, anObject)
}/* debug [instance_methods/method]: ReplaceObjectAtIndexWithObject */


// Replaces the objects in the receiving array at locations specified with the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(at:with:)
func (m_ MutableArray) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objects)
}/* debug [instance_methods/method]: ReplaceObjectsAtIndexesWithObjects */


// Replaces the objects in the receiving array specified by a given range with all of the objects from a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:)
func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArray(range_ objc.IObject /* cross-framework: Range */, otherArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:"), range_, otherArray)
}/* debug [instance_methods/method]: ReplaceObjectsInRangeWithObjectsFromArray */


// Replaces the objects in the receiving array specified by one given range with the objects in another array specified by another range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/replaceObjects(in:withObjectsFrom:range:)
func (m_ MutableArray) ReplaceObjectsInRangeWithObjectsFromArrayRange(range_ objc.IObject /* cross-framework: Range */, otherArray []objc.ID, otherRange objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsInRange:withObjectsFromArray:range:"), range_, otherArray, otherRange)
}/* debug [instance_methods/method]: ReplaceObjectsInRangeWithObjectsFromArrayRange */


// Sets the receiving array’s elements to those in another given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/setArray(_:)
func (m_ MutableArray) SetArray(otherArray []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArray:"), otherArray)
}/* debug [instance_methods/method]: SetArray */


// Replaces the object at the index with the new object, possibly adding the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/setObject:atIndexedSubscript:
func (m_ MutableArray) SetObjectAtIndexedSubscript(obj objectivec.IObject, idx uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:atIndexedSubscript:"), obj, idx)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */


// Sorts the receiver in ascending order as defined by the comparison function .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(_:context:)
func (m_ MutableArray) SortUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingFunction:context:"), compare, context)
}/* debug [instance_methods/method]: SortUsingFunctionContext */


// Sorts the receiver in ascending order using the comparison method specified by a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(comparator:)
func (m_ MutableArray) SortUsingComparator(cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingComparator:"), cmptr)
}/* debug [instance_methods/method]: SortUsingComparator */


// Sorts the receiver in ascending order using the specified options and the comparison method specified by a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(options:usingComparator:)
func (m_ MutableArray) SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}/* debug [instance_methods/method]: SortWithOptionsUsingComparator */


// Sorts the receiver using a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-4eh07
func (m_ MutableArray) SortUsingDescriptors(sortDescriptors []SortDescriptor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}/* debug [instance_methods/method]: SortUsingDescriptors */


// Sorts the receiver in ascending order, as determined by the comparison method specified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableArray/sort(using:)-537vs
func (m_ MutableArray) SortUsingSelector(comparator objc.SEL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingSelector:"), comparator)
}/* debug [instance_methods/method]: SortUsingSelector */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableArray */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableArray */


