// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [OrderedSet] class.
var (
	OrderedSetClass     _OrderedSetClass
	OrderedSetClassOnce sync.Once
)

func getOrderedSetClass() _OrderedSetClass {
	OrderedSetClassOnce.Do(func() {
		OrderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}
	})
	return OrderedSetClass
}

type _OrderedSetClass struct {
	class objc.Class
}





// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objectivec.IObject
	

	// properties:
	Array() []objc.ID
	Description() IString
	FirstObject() objectivec.IObject
	LastObject() objectivec.IObject
	ReversedOrderedSet() unsafe.Pointer
	Set() unsafe.Pointer
	Reversed() IOrderedSet
	SetReversed(value IOrderedSet)


	

	// methods:
	ContainsObject(object objectivec.IObject) bool
	DescriptionWithLocale(locale objectivec.IObject) IString
	DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) IString
	DifferenceFromOrderedSet(other unsafe.Pointer) unsafe.Pointer
	DifferenceFromOrderedSetWithOptions(other unsafe.Pointer, options OrderedCollectionDifferenceCalculationOptions) unsafe.Pointer
	DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest(other unsafe.Pointer, options OrderedCollectionDifferenceCalculationOptions, block bool) unsafe.Pointer
	EnumerateObjectsUsingBlock(block unsafe.Pointer)
	EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block unsafe.Pointer)
	EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer)
	FilteredOrderedSetUsingPredicate(p IPredicate) unsafe.Pointer
	GetObjectsRange(objects []objc.ID, range_ Range)
	IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) uint
	IndexOfObject(object objectivec.IObject) uint
	IndexOfObjectInSortedRangeOptionsUsingComparator(object objectivec.IObject, range_ Range, opts BinarySearchingOptions, cmp Comparator /* not a class type */) uint
	IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate unsafe.Pointer) uint
	IndexOfObjectPassingTest(predicate unsafe.Pointer) uint
	IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet
	IndexesOfObjectsPassingTest(predicate unsafe.Pointer) IIndexSet
	IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet
	IntersectsOrderedSet(other unsafe.Pointer) bool
	IntersectsSet(set unsafe.Pointer) bool
	IsEqualToOrderedSet(other unsafe.Pointer) bool
	IsSubsetOfOrderedSet(other unsafe.Pointer) bool
	IsSubsetOfSet(set unsafe.Pointer) bool
	ObjectAtIndex(idx uint) objectivec.IObject
	ObjectEnumerator() unsafe.Pointer
	ObjectsAtIndexes(indexes IIndexSet) []objc.ID
	OrderedSetByApplyingDifference(difference unsafe.Pointer) unsafe.Pointer
	ReverseObjectEnumerator() unsafe.Pointer
	SortedArrayUsingComparator(cmptr Comparator /* not a class type */) []objc.ID
	SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) []objc.ID
	SortedArrayUsingDescriptors(sortDescriptors []SortDescriptor) []objc.ID
	ObjectAtIndexedSubscript(idx uint) objectivec.IObject


}





// Alloc allocates a new instance without initialization.
func (oc _OrderedSetClass) Alloc() OrderedSet {
	rv := objc.Send[OrderedSet](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OrderedSetClass) New() OrderedSet {
	rv := objc.Send[OrderedSet](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OrderedSet) Init() OrderedSet {
	rv := objc.Send[OrderedSet](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OrderedSet) Autorelease() OrderedSet {
	rv := objc.Send[OrderedSet](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrderedSet creates a new OrderedSet instance.
func NewOrderedSet() OrderedSet {
	return getOrderedSetClass().New()
}





// A static, ordered collection of unique objects.
//
// declares the programmatic interface for static sets of distinct objects. You establish a static set’s entries when it’s created, and thereafter the entries can’t be modified. , on the other hand, declares a programmatic interface for dynamic sets of distinct objects. A dynamic—or mutable—set allows the addition and deletion of entries at any time, automatically allocating memory as needed. You can use ordered sets as an alternative to arrays when the order of elements is important and performance in testing whether an object is contained in the set is a consideration—testing for membership of an array is slower than testing for membership of a set.


// A static, ordered collection of unique objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet
type OrderedSet struct {
	objectivec.Object
}

// OrderedSetFrom constructs a [OrderedSet] from an unsafe.Pointer.
//
// A static, ordered collection of unique objects.
func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{objectivec.Object{objc.ID(ptr)}}
}






// Initializes a newly allocated set with the objects that are contained in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:)
func NewOrderedSetWithArray(array []objc.ID) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithArray:"), array)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:copyItems:)
func NewOrderedSetWithArrayCopyItems(set []objc.ID, flag bool) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithArray:copyItems:"), set, flag)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated set with the objects that are contained in the specified range of an array, optionally copying the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(array:range:copyItems:)
func NewOrderedSetWithArrayRangeCopyItems(set []objc.ID, range_ Range, flag bool) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithArray:range:copyItems:"), set, range_, flag)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(coder:)
func NewOrderedSetWithCoder(coder ICoder) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(object:)
func NewOrderedSetWithObject(object objectivec.IObject) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithObject:"), object)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated set with members taken from the specified list of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/initWithObjects:
func NewOrderedSetWithObjects(firstObj objectivec.IObject) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithObjects:"), firstObj)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated set with a specified number of objects from a given C array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(objects:count:)-2ai32
func NewOrderedSetWithObjectsCount(objects []objc.ID, cnt uint) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithObjects:count:"), objects, cnt)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the contents of a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:)
func NewOrderedSetWithOrderedSet(set unsafe.Pointer) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithOrderedSet:"), set)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the contents of a set, optionally copying the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:copyItems:)
func NewOrderedSetWithOrderedSetCopyItems(set unsafe.Pointer, flag bool) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithOrderedSet:copyItems:"), set, flag)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the contents of an ordered set, optionally copying the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(orderedSet:range:copyItems:)
func NewOrderedSetWithOrderedSetRangeCopyItems(set unsafe.Pointer, range_ Range, flag bool) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithOrderedSet:range:copyItems:"), set, range_, flag)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the contents of a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:)
func NewOrderedSetWithSet(set unsafe.Pointer) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithSet:"), set)
	rv.Autorelease()
	return rv
}


// Initializes a new ordered set with the contents of a set, optionally copying the objects in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(set:copyItems:)
func NewOrderedSetWithSetCopyItems(set unsafe.Pointer, flag bool) OrderedSet {
	instance := getOrderedSetClass().Alloc()
	rv := objc.Send[OrderedSet](instance.ID, objc.Sel("initWithSet:copyItems:"), set, flag)
	rv.Autorelease()
	return rv
}







// Creates and returns a set containing a specified number of objects from a given C array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/init(objects:count:)-3ny0m
func (oc _OrderedSetClass) OrderedSetWithObjectsCount(objects []objc.ID, cnt uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithObjects:count:"), objects, cnt)
	return rv
}


// Creates and returns an empty ordered set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSet
func (oc _OrderedSetClass) OrderedSet() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSet"))
	return rv
}


// Creates and returns a set containing a uniqued collection of the objects contained in a given array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithArray:
func (oc _OrderedSetClass) OrderedSetWithArray(array []objc.ID) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithArray:"), array)
	return rv
}


// Creates and returns a new ordered set for a specified range of objects in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithArray:range:copyItems:
func (oc _OrderedSetClass) OrderedSetWithArrayRangeCopyItems(array []objc.ID, range_ Range, flag bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithArray:range:copyItems:"), array, range_, flag)
	return rv
}


// Creates and returns a ordered set that contains a single given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithObject:
func (oc _OrderedSetClass) OrderedSetWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithObject:"), object)
	return rv
}


// Creates and returns a ordered set containing the objects in a given argument list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithObjects:
func (oc _OrderedSetClass) OrderedSetWithObjects(firstObj objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithObjects:"), firstObj)
	return rv
}


// Creates and returns an ordered set containing the objects from another ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithOrderedSet:
func (oc _OrderedSetClass) OrderedSetWithOrderedSet(set unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithOrderedSet:"), set)
	return rv
}


// Creates and returns a new ordered set for a specified range of objects in an ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithOrderedSet:range:copyItems:
func (oc _OrderedSetClass) OrderedSetWithOrderedSetRangeCopyItems(set unsafe.Pointer, range_ Range, flag bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithOrderedSet:range:copyItems:"), set, range_, flag)
	return rv
}


// Creates and returns an ordered set with the contents of a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithSet:
func (oc _OrderedSetClass) OrderedSetWithSet(set unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithSet:"), set)
	return rv
}


// Creates and returns an ordered set with the contents of a set, optionally copying the items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetWithSet:copyItems:
func (oc _OrderedSetClass) OrderedSetWithSetCopyItems(set unsafe.Pointer, flag bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("orderedSetWithSet:copyItems:"), set, flag)
	return rv
}












// Raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/addObserver(_:forKeyPath:options:context:)
func (o_ OrderedSet) AddObserverForKeyPathOptionsContext(observer objectivec.IObject, keyPath IString, options uint, context objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObserver:forKeyPath:options:context:"), observer, keyPath, options, context)
}


// Returns a Boolean value that indicates whether a given object is present in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/contains(_:)
func (o_ OrderedSet) ContainsObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("containsObject:"), object)
	return rv
}


// Returns a string that represents the contents of the ordered set, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description(withLocale:)
func (o_ OrderedSet) DescriptionWithLocale(locale objectivec.IObject) IString {
	rv := objc.Send[String](o_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}


// Returns a string that represents the contents of the ordered set, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description(withLocale:indent:)
func (o_ OrderedSet) DescriptionWithLocaleIndent(locale objectivec.IObject, level uint) IString {
	rv := objc.Send[String](o_.ID, objc.Sel("descriptionWithLocale:indent:"), locale, level)
	return rv
}


// Compares two ordered sets to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:
func (o_ OrderedSet) DifferenceFromOrderedSet(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("differenceFromOrderedSet:"), other)
	return rv
}


// Compares two ordered sets, with options, to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:withOptions:
func (o_ OrderedSet) DifferenceFromOrderedSetWithOptions(other unsafe.Pointer, options OrderedCollectionDifferenceCalculationOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("differenceFromOrderedSet:withOptions:"), other, options)
	return rv
}


// Compares two ordered sets, using the provided block and with options, to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/differenceFromOrderedSet:withOptions:usingEquivalenceTest:
func (o_ OrderedSet) DifferenceFromOrderedSetWithOptionsUsingEquivalenceTest(other unsafe.Pointer, options OrderedCollectionDifferenceCalculationOptions, block bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("differenceFromOrderedSet:withOptions:usingEquivalenceTest:"), other, options, block)
	return rv
}


// Executes a given block using each object in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(_:)
func (o_ OrderedSet) EnumerateObjectsUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("enumerateObjectsUsingBlock:"), block)
}


// Executes a given block using the objects in the ordered set at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(at:options:using:)
func (o_ OrderedSet) EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), s, opts, block)
}


// Executes a given block using each object in the set, using the specified enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/enumerateObjects(options:using:)
func (o_ OrderedSet) EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}


// Evaluates a given predicate against each object in the receiving ordered set and returns a new ordered set containing the objects for which the predicate returns true.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/filtered(using:)
func (o_ OrderedSet) FilteredOrderedSetUsingPredicate(p IPredicate) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("filteredOrderedSetUsingPredicate:"), p)
	return rv
}


// Copies the objects contained in the ordered set that fall within the specified range to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/getObjects:range:
func (o_ OrderedSet) GetObjectsRange(objects []objc.ID, range_ Range) {
	objc.Send[objc.ID](o_.ID, objc.Sel("getObjects:range:"), objects, range_)
}


// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(_:ofObjectPassingTest:)
func (o_ OrderedSet) IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, predicate)
	return rv
}


// Returns the index of the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(of:)
func (o_ OrderedSet) IndexOfObject(object objectivec.IObject) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObject:"), object)
	return rv
}


// Returns the index, within a specified range, of an object compared with elements in the ordered set using a given NSComparator block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(of:inSortedRange:options:usingComparator:)
func (o_ OrderedSet) IndexOfObjectInSortedRangeOptionsUsingComparator(object objectivec.IObject, range_ Range, opts BinarySearchingOptions, cmp Comparator /* not a class type */) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObject:inSortedRange:options:usingComparator:"), object, range_, opts, cmp)
	return rv
}


// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(ofObjectAt:options:passingTest:)
func (o_ OrderedSet) IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), s, opts, predicate)
	return rv
}


// Returns the index of the object in the ordered set that passes a test in a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/index(ofObjectPassingTest:)
func (o_ OrderedSet) IndexOfObjectPassingTest(predicate unsafe.Pointer) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}


// Returns the index, from a given set of indexes, of the object in the ordered set that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(ofObjectsAt:options:passingTest:)
func (o_ OrderedSet) IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](o_.ID, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), s, opts, predicate)
	return rv
}


// Returns the index of the object in the ordered set that passes a test in a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(ofObjectsPassingTest:)
func (o_ OrderedSet) IndexesOfObjectsPassingTest(predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](o_.ID, objc.Sel("indexesOfObjectsPassingTest:"), predicate)
	return rv
}


// Returns the index of an object in the ordered set that passes a test in a given block for a given set of enumeration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/indexes(options:ofObjectsPassingTest:)
func (o_ OrderedSet) IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate unsafe.Pointer) IIndexSet {
	rv := objc.Send[IndexSet](o_.ID, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, predicate)
	return rv
}


// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/intersects(_:)
func (o_ OrderedSet) IntersectsOrderedSet(other unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("intersectsOrderedSet:"), other)
	return rv
}


// Returns a Boolean value that indicates whether at least one object in the receiving ordered set is also present in another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/intersectsSet(_:)
func (o_ OrderedSet) IntersectsSet(set unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("intersectsSet:"), set)
	return rv
}


// Compares the receiving ordered set to another ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isEqual(to:)
func (o_ OrderedSet) IsEqualToOrderedSet(other unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEqualToOrderedSet:"), other)
	return rv
}


// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isSubset(of:)-7brc
func (o_ OrderedSet) IsSubsetOfOrderedSet(other unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSubsetOfOrderedSet:"), other)
	return rv
}


// Returns a Boolean value that indicates whether every object in the receiving ordered set is also present in another given set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/isSubset(of:)-8zx9x
func (o_ OrderedSet) IsSubsetOfSet(set unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSubsetOfSet:"), set)
	return rv
}


// Returns the object at the specified index of the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/object(at:)
func (o_ OrderedSet) ObjectAtIndex(idx uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("objectAtIndex:"), idx)
	return rv
}


// Returns an enumerator object that lets you access each object in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/objectEnumerator()
func (o_ OrderedSet) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("objectEnumerator"))
	return rv
}


// Returns the objects in the ordered set at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/objects(at:)
func (o_ OrderedSet) ObjectsAtIndexes(indexes IIndexSet) []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("objectsAtIndexes:"), indexes)
	return rv
}


// Creates a new ordered set by applying a difference object to an existing ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/orderedSetByApplyingDifference:
func (o_ OrderedSet) OrderedSetByApplyingDifference(difference unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("orderedSetByApplyingDifference:"), difference)
	return rv
}


// Raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/removeObserver(_:forKeyPath:)
func (o_ OrderedSet) RemoveObserverForKeyPath(observer objectivec.IObject, keyPath IString) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:forKeyPath:"), observer, keyPath)
}


// Raises an exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/removeObserver(_:forKeyPath:context:)
func (o_ OrderedSet) RemoveObserverForKeyPathContext(observer objectivec.IObject, keyPath IString, context objectivec.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:forKeyPath:context:"), observer, keyPath, context)
}


// Returns an enumerator object that lets you access each object in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/reverseObjectEnumerator()
func (o_ OrderedSet) ReverseObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("reverseObjectEnumerator"))
	return rv
}


// Invokes on each of the receiver’s members using the specified value and key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/setValue(_:forKey:)
func (o_ OrderedSet) SetValueForKey(value objectivec.IObject, key IString) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forKey:"), value, key)
}


// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given block
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(comparator:)
func (o_ OrderedSet) SortedArrayUsingComparator(cmptr Comparator /* not a class type */) []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return rv
}


// Returns an array that lists the receiving ordered set’s elements in ascending order, as determined by the comparison method specified by a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(options:usingComparator:)
func (o_ OrderedSet) SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator /* not a class type */) []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return rv
}


// Returns an array of the ordered set’s elements sorted as specified by a given array of sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/sortedArray(using:)
func (o_ OrderedSet) SortedArrayUsingDescriptors(sortDescriptors []SortDescriptor) []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}


// Returns the object at the specified index of the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/subscript(_:)
func (o_ OrderedSet) ObjectAtIndexedSubscript(idx uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}


// Returns an ordered set containing the results of invoking using key on each of the ordered set’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/value(forKey:)
func (o_ OrderedSet) ValueForKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKey:"), key)
	return rv
}







// A representation of the ordered set as an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/array
func (o_ OrderedSet) Array() []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("array"))
	return rv
}


// A string that represents the contents of the ordered set, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/description
func (o_ OrderedSet) Description() IString {
	rv := objc.Send[String](o_.ID, objc.Sel("description"))
	return rv
}


// The first object in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/firstObject
func (o_ OrderedSet) FirstObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("firstObject"))
	return rv
}


// The last object in the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/lastObject
func (o_ OrderedSet) LastObject() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](o_.ID, objc.Sel("lastObject"))
	return rv
}


// An ordered set in the reverse order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/reversed
func (o_ OrderedSet) ReversedOrderedSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("reversedOrderedSet"))
	return rv
}


// A representation of the set containing the contents of the ordered set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrderedSet/set
func (o_ OrderedSet) Set() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("set"))
	return rv
}


// An ordered set in the reverse order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/reversed
func (o_ OrderedSet) Reversed() IOrderedSet {
	rv := objc.Send[OrderedSet](o_.ID, objc.Sel("reversed"))
	return rv
}


// An ordered set in the reverse order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/reversed
func (o_ OrderedSet) SetReversed(value IOrderedSet) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setReversed:"), value)
}







