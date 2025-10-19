// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OrderedSet] class.
var orderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}

type _OrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [OrderedSet] class.
type IOrderedSet interface {
	objectivec.IObject
}

// A static, ordered collection of unique objects. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (oc _OrderedSetClass) Alloc() OrderedSet {
	rv := objc.Send[OrderedSet](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return orderedSetClass.New()
}




