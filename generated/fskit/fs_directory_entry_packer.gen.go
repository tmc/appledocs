// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSDirectoryEntryPacker */


/* debug [class_header]: Header for FSDirectoryEntryPacker */
// The class instance for the [FSDirectoryEntryPacker] class.
var (
	FSDirectoryEntryPackerClass     _FSDirectoryEntryPackerClass
	FSDirectoryEntryPackerClassOnce sync.Once
)

func getFSDirectoryEntryPackerClass() _FSDirectoryEntryPackerClass {
	FSDirectoryEntryPackerClassOnce.Do(func() {
		FSDirectoryEntryPackerClass = _FSDirectoryEntryPackerClass{objc.GetClass("FSDirectoryEntryPacker")}
	})
	return FSDirectoryEntryPackerClass
}

type _FSDirectoryEntryPackerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSDirectoryEntryPacker */
// An interface definition for the [FSDirectoryEntryPacker] class.
type IFSDirectoryEntryPacker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSDirectoryEntryPacker */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSDirectoryEntryPacker */
	// methods:
	PackEntryWithNameItemTypeItemIDNextCookieAttributes(name IFSFileName, itemType FSItemType, itemID FSItemID, nextCookie FSDirectoryCookie /* typedef */, attributes IFSItemAttributes) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSDirectoryEntryPacker */
// Alloc allocates a new instance without initialization.
func (fc _FSDirectoryEntryPackerClass) Alloc() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSDirectoryEntryPackerClass) New() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSDirectoryEntryPacker) Init() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSDirectoryEntryPacker) Autorelease() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSDirectoryEntryPacker creates a new FSDirectoryEntryPacker instance.
func NewFSDirectoryEntryPacker() FSDirectoryEntryPacker {
	return getFSDirectoryEntryPackerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSDirectoryEntryPacker */
// An object used to provide items during a directory enumeration.
//
// You use this type in your implementation of . Packing allows your implementation to provide information FSKit needs, including each item’s name, type, and identifier (such as an inode number). Some directory enumerations require other attributes, as indicated by the sent to the enumerate method.


// An object used to provide items during a directory enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker
type FSDirectoryEntryPacker struct {
	objectivec.Object
}

// FSDirectoryEntryPackerFrom constructs a [FSDirectoryEntryPacker] from an unsafe.Pointer.
//
// An object used to provide items during a directory enumeration.
func FSDirectoryEntryPackerFrom(ptr unsafe.Pointer) FSDirectoryEntryPacker {
	return FSDirectoryEntryPacker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSDirectoryEntryPacker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSDirectoryEntryPacker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSDirectoryEntryPacker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSDirectoryEntryPacker */

// Provides a directory entry during enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker/packEntry(name:itemType:itemID:nextCookie:attributes:)
func (f_ FSDirectoryEntryPacker) PackEntryWithNameItemTypeItemIDNextCookieAttributes(name IFSFileName, itemType FSItemType, itemID FSItemID, nextCookie FSDirectoryCookie /* typedef */, attributes IFSItemAttributes) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("packEntryWithName:itemType:itemID:nextCookie:attributes:"), name, itemType, itemID, nextCookie, attributes)
	return rv
}/* debug [instance_methods/method]: PackEntryWithNameItemTypeItemIDNextCookieAttributes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSDirectoryEntryPacker */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSDirectoryEntryPacker */



