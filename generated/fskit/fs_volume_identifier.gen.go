// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class FSVolumeIdentifier */


/* debug [class_header]: Header for FSVolumeIdentifier */
// The class instance for the [FSVolumeIdentifier] class.
var (
	FSVolumeIdentifierClass     _FSVolumeIdentifierClass
	FSVolumeIdentifierClassOnce sync.Once
)

func getFSVolumeIdentifierClass() _FSVolumeIdentifierClass {
	FSVolumeIdentifierClassOnce.Do(func() {
		FSVolumeIdentifierClass = _FSVolumeIdentifierClass{objc.GetClass("FSVolumeIdentifier")}
	})
	return FSVolumeIdentifierClass
}

type _FSVolumeIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSVolumeIdentifier */
// An interface definition for the [FSVolumeIdentifier] class.
type IFSVolumeIdentifier interface {
	IFSEntityIdentifier
	
/* debug [class_interface_properties]: Properties for FSVolumeIdentifier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSVolumeIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSVolumeIdentifier */
// Alloc allocates a new instance without initialization.
func (fc _FSVolumeIdentifierClass) Alloc() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSVolumeIdentifierClass) New() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSVolumeIdentifier) Init() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSVolumeIdentifier) Autorelease() FSVolumeIdentifier {
	rv := objc.Send[FSVolumeIdentifier](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeIdentifier creates a new FSVolumeIdentifier instance.
func NewFSVolumeIdentifier() FSVolumeIdentifier {
	return getFSVolumeIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSVolumeIdentifier */
// A type that identifies a volume.
//
// For most volumes, the volume identifier is the UUID identifying the volume. Network file systems may access the same underlying volume using different authentication credentials. To handle this situation, add qualifying data to identify the specific container, as discussed in the superclass, .


// A type that identifies a volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/Identifier
type FSVolumeIdentifier struct {
	FSEntityIdentifier
}

// FSVolumeIdentifierFrom constructs a [FSVolumeIdentifier] from an unsafe.Pointer.
//
// A type that identifies a volume.
func FSVolumeIdentifierFrom(ptr unsafe.Pointer) FSVolumeIdentifier {
	return FSVolumeIdentifier{
		FSEntityIdentifier: FSEntityIdentifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSVolumeIdentifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSVolumeIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSVolumeIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSVolumeIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSVolumeIdentifier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSVolumeIdentifier */



