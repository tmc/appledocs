// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [sortedAttributes] class.
var (
	SortedAttributesClass     _sortedAttributesClass
	SortedAttributesClassOnce sync.Once
)

func getsortedAttributesClass() _sortedAttributesClass {
	SortedAttributesClassOnce.Do(func() {
		SortedAttributesClass = _sortedAttributesClass{objc.GetClass("sortedAttributes")}
	})
	return SortedAttributesClass
}

type _sortedAttributesClass struct {
	class objc.Class
}

// An interface definition for the [sortedAttributes] class.
type IsortedAttributes interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/sortedAttributes-c.ivar
type sortedAttributes struct {
	objectivec.Object
}

// sortedAttributesFrom constructs a [sortedAttributes] from an unsafe.Pointer.
func sortedAttributesFrom(ptr unsafe.Pointer) sortedAttributes {
	return sortedAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _sortedAttributesClass) Alloc() sortedAttributes {
	rv := objc.Send[sortedAttributes](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _sortedAttributesClass) New() sortedAttributes {
	rv := objc.Send[sortedAttributes](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ sortedAttributes) Init() sortedAttributes {
	rv := objc.Send[sortedAttributes](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ sortedAttributes) Autorelease() sortedAttributes {
	rv := objc.Send[sortedAttributes](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsortedAttributes creates a new sortedAttributes instance.
func NewsortedAttributes() sortedAttributes {
	return getsortedAttributesClass().New()
}




