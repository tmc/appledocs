// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Set] class.
var (
	SetClass     _SetClass
	SetClassOnce sync.Once
)

func getSetClass() _SetClass {
	SetClassOnce.Do(func() {
		SetClass = _SetClass{objc.GetClass("NSSet")}
	})
	return SetClass
}

type _SetClass struct {
	class objc.Class
}

// An interface definition for the [Set] class.
type ISet interface {
	objectivec.IObject
	AllObjects() unsafe.Pointer
	SetAllObjects(value unsafe.Pointer)
	Count() int
	SetCount(value int)
	Description() string
	SetDescription(value string)
}

// A static, unordered collection of unique objects.
//
// The , , and classes declare the programmatic interface to an unordered collection of objects. declares the programmatic interface for static sets of distinct objects. You establish a static set’s entries when it’s created, and can’t modify the entries after that. , on the other hand, declares a programmatic interface for dynamic sets of distinct objects. A dynamic — or mutable — set allows the addition and deletion of entries at any time, automatically allocating memory as needed. Use sets as an alternative to arrays when the order of elements isn’t important and you need to consider performance in testing whether the set contains an object. With an array, testing for membership is slower than with sets. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging. In Swift, use this class instead of a constant in cases where you require reference semantics.


// A static, unordered collection of unique objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSet
type Set struct {
	objectivec.Object
}

// SetFrom constructs a [Set] from an unsafe.Pointer.
//
// A static, unordered collection of unique objects.
func SetFrom(ptr unsafe.Pointer) Set {
	return Set{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SetClass) Alloc() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SetClass) New() Set {
	rv := objc.Send[Set](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Set) Init() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Set) Autorelease() Set {
	rv := objc.Send[Set](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSet creates a new Set instance.
func NewSet() Set {
	return getSetClass().New()
}



// An array containing the set’s members, or an empty array if the set has no members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/allobjects
func (s_ Set) AllObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("allObjects"))
	return rv
}


// An array containing the set’s members, or an empty array if the set has no members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/allobjects
func (s_ Set) SetAllObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllObjects:"), value)
}


// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (s_ Set) Count() int {
	rv := objc.Send[int](s_.ID, objc.Sel("count"))
	return rv
}


// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (s_ Set) SetCount(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCount:"), value)
}


// A string that represents the contents of the set, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/description
func (s_ Set) Description() string {
	rv := objc.Send[string](s_.ID, objc.Sel("description"))
	return rv
}


// A string that represents the contents of the set, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/description
func (s_ Set) SetDescription(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDescription:"), objc.String(value))
}



