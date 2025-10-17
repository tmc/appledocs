// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Bundle] class.
var bundleClass = _BundleClass{objc.GetClass("NSBundle")}

type _BundleClass struct {
	class objc.Class
}

// A representation of the code and resources stored in a bundle directory on disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle

type Bundle struct {
	objectivec.Object
}

// BundleFrom constructs a [Bundle] from an unsafe.Pointer.
//
// A representation of the code and resources stored in a bundle directory on disk.
func BundleFrom(ptr unsafe.Pointer) Bundle {
	return Bundle{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (bc _BundleClass) Alloc() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (bc _BundleClass) New() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Bundle) Init() Bundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Bundle) Autorelease() Bundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBundle creates a new Bundle instance.
func NewBundle() Bundle {
	return bundleClass.New()
}
// Returns the object with which the specified class is associated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[Bundle](objc.ID(bundleClass.class), objc.Sel("bundleForClass:"), aClass)
	rv.Autorelease()
	return rv
}


// Returns the object with which the specified class is associated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}
// Returns the value associated with the specified key in the receiver’s information property list. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("objectForInfoDictionaryKey:"), key)
	return rv
}

