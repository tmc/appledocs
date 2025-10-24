// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHObject] class.
var (
	PHObjectClass     _PHObjectClass
	PHObjectClassOnce sync.Once
)

func getPHObjectClass() _PHObjectClass {
	PHObjectClassOnce.Do(func() {
		PHObjectClass = _PHObjectClass{objc.GetClass("PHObject")}
	})
	return PHObjectClass
}

type _PHObjectClass struct {
	class objc.Class
}

// An interface definition for the [PHObject] class.
type IPHObject interface {
	objectivec.IObject
	// properties:
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	Hash() int
	SetHash(value int)
	// methods:
}

// The abstract superclass for Photos model objects (assets and collections).
//
// You do not create or use instances of this class directly. Instead, work with instances of its concrete subclasses— , , , and . Because the class implements the and methods in terms of its property, you can use techniques that depend on these methods to keep track of asset and collection objects.

// The abstract superclass for Photos model objects (assets and collections).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHObject
type PHObject struct {
	objectivec.Object
}

// PHObjectFrom constructs a [PHObject] from an unsafe.Pointer.
//
// The abstract superclass for Photos model objects (assets and collections).
func PHObjectFrom(ptr unsafe.Pointer) PHObject {
	return PHObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHObjectClass) Alloc() PHObject {
	rv := objc.Send[PHObject](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHObjectClass) New() PHObject {
	rv := objc.Send[PHObject](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHObject) Init() PHObject {
	rv := objc.Send[PHObject](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHObject) Autorelease() PHObject {
	rv := objc.Send[PHObject](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHObject creates a new PHObject instance.
func NewPHObject() PHObject {
	return getPHObjectClass().New()
}

// A unique string that persistently identifies the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHObject/localIdentifier
func (p_ PHObject) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localIdentifier"))
	return rv
}

// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (p_ PHObject) Hash() int {
	rv := objc.Send[int](p_.ID, objc.Sel("hash"))
	return rv
}

// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (p_ PHObject) SetHash(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHash:"), value)
}
