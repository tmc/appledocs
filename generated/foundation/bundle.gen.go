// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Bundle] class.
var (
	bundleClass     _BundleClass
	bundleClassOnce sync.Once
)

func getBundleClass() _BundleClass {
	bundleClassOnce.Do(func() {
		bundleClass = _BundleClass{objc.GetClass("NSBundle")}
	})
	return bundleClass
}

type _BundleClass struct {
	class objc.Class
}

// An interface definition for the [Bundle] class.
type IBundle interface {
	objectivec.IObject
	ObjectForInfoDictionaryKey(key string) objc.ID
}

// A representation of the code and resources stored in a bundle directory on disk.
//
// Apple uses bundles to represent apps, frameworks, plug-ins, and many other specific types of content. Bundles organize their contained resources into well-defined subdirectories, and bundle structures vary depending on the platform and the type of the bundle. By using a bundle object, you can access a bundle’s resources without knowing the structure of the bundle. The bundle object provides a single interface for locating items, taking into account the bundle structure, user preferences, available localizations, and other relevant factors. Any executable can use a bundle object to locate resources, either inside an app’s bundle or in a known bundle located elsewhere. You don’t use a bundle object to locate files in a container directory or in other parts of the file system. The general pattern for using a bundle object is as follows: Create a bundle object for the intended bundle directory. Use the methods of the bundle object to locate or load the needed resource. Use other system APIs to interact with the resource. Some types of frequently used resources can be located and opened without a bundle. For example, when loading images, you store images in asset catalogs and load them using the methods of or . Similarly, for string resources, you use to load individual strings instead of loading the entire file yourself.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getBundleClass().New()
}


// Returns the object with which the specified class is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleForClass:"), aClass)
	return rv
}


// Returns the object with which the specified class is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}

// Returns the value associated with the specified key in the receiver’s information property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("objectForInfoDictionaryKey:"), objc.String(key))
	return rv
}

// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (b_ Bundle) AppStoreReceiptURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}


// The receiver’s bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundleIdentifier
func (b_ Bundle) BundleIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// The bundle’s principal class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/principalClass
func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("principalClass"))
	return rv
}



