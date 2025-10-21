// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHObjectPlaceholder] class.
var (
	PHObjectPlaceholderClass     _PHObjectPlaceholderClass
	PHObjectPlaceholderClassOnce sync.Once
)

func getPHObjectPlaceholderClass() _PHObjectPlaceholderClass {
	PHObjectPlaceholderClassOnce.Do(func() {
		PHObjectPlaceholderClass = _PHObjectPlaceholderClass{objc.GetClass("PHObjectPlaceholder")}
	})
	return PHObjectPlaceholderClass
}

type _PHObjectPlaceholderClass struct {
	class objc.Class
}

// An interface definition for the [PHObjectPlaceholder] class.
type IPHObjectPlaceholder interface {
	IPHObject
}

// A read-only proxy object that represents a Photos asset or collection to create.
//
// You obtain object placeholders when you use change requests to create assets, collections, or collection lists. After the change request completes, you can use the object placeholder to fetch the newly created object. You can also use an object placeholder to make additional change requests involving the object to create. For example, the following code uses a placeholder to add a newly created asset to an album. A placeholder always has the same local identifier as the asset, collection, or collection list that it represents. To find the object that corresponds to a placeholder, read the placeholder’s property and use it to fetch the actual object. Alternatively, because the class implements the and methods in terms of its property, you can also find the object for a placeholder using techniques that depend on these methods.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHObjectPlaceholder
type PHObjectPlaceholder struct {
	PHObject
}

// PHObjectPlaceholderFrom constructs a [PHObjectPlaceholder] from an unsafe.Pointer.
//
// A read-only proxy object that represents a Photos asset or collection to create.
func PHObjectPlaceholderFrom(ptr unsafe.Pointer) PHObjectPlaceholder {
	return PHObjectPlaceholder{
		PHObject: PHObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHObjectPlaceholderClass) Alloc() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHObjectPlaceholderClass) New() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHObjectPlaceholder) Init() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHObjectPlaceholder) Autorelease() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHObjectPlaceholder creates a new PHObjectPlaceholder instance.
func NewPHObjectPlaceholder() PHObjectPlaceholder {
	return getPHObjectPlaceholderClass().New()
}


// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (p_ PHObjectPlaceholder) Hash() int {
	rv := objc.Send[int](p_.ID, objc.Sel("hash"))
	return rv
}


// SetHash sets the value of the hash property.
// Returns an integer that can be used as a table address in a hash table structure.

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (p_ PHObjectPlaceholder) SetHash(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHash:"), value)
}

// A unique string that persistently identifies the object.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHObjectPlaceholder) LocalIdentifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("localIdentifier"))
	return rv
}


// SetLocalIdentifier sets the value of the localIdentifier property.
// A unique string that persistently identifies the object.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHObjectPlaceholder) SetLocalIdentifier(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalIdentifier:"), value)
}



