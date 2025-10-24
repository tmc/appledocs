// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMutableIndexSet */


/* debug [class_header]: Header for NSMutableIndexSet */
// The class instance for the [MutableIndexSet] class.
var (
	MutableIndexSetClass     _MutableIndexSetClass
	MutableIndexSetClassOnce sync.Once
)

func getMutableIndexSetClass() _MutableIndexSetClass {
	MutableIndexSetClassOnce.Do(func() {
		MutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}
	})
	return MutableIndexSetClass
}

type _MutableIndexSetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableIndexSet */
// An interface definition for the [MutableIndexSet] class.
type IMutableIndexSet interface {
	IIndexSet
	
/* debug [class_interface_properties]: Properties for MutableIndexSet */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableIndexSet */
	// methods:
	AddIndex(value uint)
	AddIndexes(indexSet IIndexSet)
	AddIndexesInRange(range_ objc.IObject /* cross-framework: Range */)
	RemoveIndexes(indexSet IIndexSet)
	RemoveIndex(value uint)
	RemoveIndexesInRange(range_ objc.IObject /* cross-framework: Range */)
	RemoveAllIndexes()
	ShiftIndexesStartingAtIndexBy(index uint, delta int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableIndexSet */
// Alloc allocates a new instance without initialization.
func (mc _MutableIndexSetClass) Alloc() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableIndexSetClass) New() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableIndexSet) Init() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableIndexSet) Autorelease() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableIndexSet creates a new MutableIndexSet instance.
func NewMutableIndexSet() MutableIndexSet {
	return getMutableIndexSetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableIndexSet */
// A mutable collection of unique integer values that represent indexes in another collection.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class represents a mutable collection of unique unsigned integers, known as because of the way they are used. This collection is referred to as a . The inclusive range of valid indexes is ; trying to use indexes outside this range is invalid. The values in a mutable index set are always sorted, so the order in which values are added is irrelevant. Do not subclass the class.


// A mutable collection of unique integer values that represent indexes in another collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet
type MutableIndexSet struct {
	IndexSet
}

// MutableIndexSetFrom constructs a [MutableIndexSet] from an unsafe.Pointer.
//
// A mutable collection of unique integer values that represent indexes in another collection.
func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		IndexSet: IndexSetFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableIndexSet *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableIndexSet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableIndexSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableIndexSet */

// Adds an index to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(_:)-6dtkj
func (m_ MutableIndexSet) AddIndex(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addIndex:"), value)
}/* debug [instance_methods/method]: AddIndex */


// Adds the indexes in an index set to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(_:)-6zmti
func (m_ MutableIndexSet) AddIndexes(indexSet IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addIndexes:"), indexSet)
}/* debug [instance_methods/method]: AddIndexes */


// Adds the indexes in an index range to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(in:)
func (m_ MutableIndexSet) AddIndexesInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addIndexesInRange:"), range_)
}/* debug [instance_methods/method]: AddIndexesInRange */


// Removes the indexes in an index set from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-196u2
func (m_ MutableIndexSet) RemoveIndexes(indexSet IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndexes:"), indexSet)
}/* debug [instance_methods/method]: RemoveIndexes */


// Removes an index from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-5li0r
func (m_ MutableIndexSet) RemoveIndex(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndex:"), value)
}/* debug [instance_methods/method]: RemoveIndex */


// Removes the indexes in an index range from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(in:)
func (m_ MutableIndexSet) RemoveIndexesInRange(range_ objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndexesInRange:"), range_)
}/* debug [instance_methods/method]: RemoveIndexesInRange */


// Removes the receiver’s indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/removeAllIndexes()
func (m_ MutableIndexSet) RemoveAllIndexes() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllIndexes"))
}/* debug [instance_methods/method]: RemoveAllIndexes */


// Shifts a group of indexes to the left or the right within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/shiftIndexesStarting(at:by:)
func (m_ MutableIndexSet) ShiftIndexesStartingAtIndexBy(index uint, delta int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("shiftIndexesStartingAtIndex:by:"), index, delta)
}/* debug [instance_methods/method]: ShiftIndexesStartingAtIndexBy */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableIndexSet */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableIndexSet */



