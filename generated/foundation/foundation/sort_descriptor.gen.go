// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SortDescriptor] class.
var SortDescriptorClass objc.Class

func init() {
	SortDescriptorClass = objc.GetClass("NSSortDescriptor")
}

type SortDescriptor struct {
	objc.ID
}

func SortDescriptorFrom(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc SortDescriptor) Alloc() SortDescriptor {
	ret := objc.ID(SortDescriptorClass).Send(objc.RegisterName("alloc"))
	return SortDescriptor{ret}
}

// Init initializes the instance.
func (s_ SortDescriptor) Init() SortDescriptor {
	ret := s_.ID.Send(objc.RegisterName("init"))
	return SortDescriptor{ret}
}
// Creates a sort descriptor with a specified string key path and sort order. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSSortDescriptor/init(key:ascending:)
func NewSortDescriptorWithKeyAscending(key string, ascending bool) SortDescriptor {
	instance := SortDescriptor{}.Alloc()
	sel := objc.RegisterName("initWithKey:ascending:")
	ret := instance.ID.Send(sel, key, ascending)
	instance = SortDescriptor{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates and returns a sort descriptor initialized with the specified key path and ordering, and a comparator block. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSSortDescriptor/sortDescriptorWithKey:ascending:comparator:
func (sc SortDescriptor) SortDescriptorWithKeyAscendingComparator(key string, ascending bool, cmptr unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sortDescriptorWithKey:ascending:comparator:")
	ret := objc.ID(SortDescriptorClass).Send(sel, key, ascending, cmptr)
	return unsafe.Pointer(ret)
}
// Returns a comparison result value that indicates the sort order of two objects. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSSortDescriptor/compare(_:to:)
func (s_ SortDescriptor) CompareObjectToObject(object1 objc.ID, object2 objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("compareObject:toObject:")
	ret := s_.ID.Send(sel, object1, object2)
	return unsafe.Pointer(ret)
}

