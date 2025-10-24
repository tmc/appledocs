// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSExtentPacker */


/* debug [class_header]: Header for FSExtentPacker */
// The class instance for the [FSExtentPacker] class.
var (
	FSExtentPackerClass     _FSExtentPackerClass
	FSExtentPackerClassOnce sync.Once
)

func getFSExtentPackerClass() _FSExtentPackerClass {
	FSExtentPackerClassOnce.Do(func() {
		FSExtentPackerClass = _FSExtentPackerClass{objc.GetClass("FSExtentPacker")}
	})
	return FSExtentPackerClass
}

type _FSExtentPackerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSExtentPacker */
// An interface definition for the [FSExtentPacker] class.
type IFSExtentPacker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSExtentPacker */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSExtentPacker */
	// methods:
	PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset unsafe.Pointer, physicalOffset unsafe.Pointer, length uintptr /* not a class type */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSExtentPacker */
// Alloc allocates a new instance without initialization.
func (fc _FSExtentPackerClass) Alloc() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSExtentPackerClass) New() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSExtentPacker) Init() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSExtentPacker) Autorelease() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSExtentPacker creates a new FSExtentPacker instance.
func NewFSExtentPacker() FSExtentPacker {
	return getFSExtentPackerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSExtentPacker */
// A type that directs the kernel to map space on disk to a specific file managed by this file system.
//
// provide the kernel the logical-to-physical mapping of a given file. An extent describes a physical offset on disk, and a length and a logical offset within the file. Rather than working with extents directly, you use this type’s methods to provide or “pack” extent information, which FSKit then passes to the kernel.


// A type that directs the kernel to map space on disk to a specific file managed by this file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentPacker
type FSExtentPacker struct {
	objectivec.Object
}

// FSExtentPackerFrom constructs a [FSExtentPacker] from an unsafe.Pointer.
//
// A type that directs the kernel to map space on disk to a specific file managed by this file system.
func FSExtentPackerFrom(ptr unsafe.Pointer) FSExtentPacker {
	return FSExtentPacker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSExtentPacker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSExtentPacker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSExtentPacker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSExtentPacker */

// Packs a single extent to send to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentPacker/packExtent(resource:type:logicalOffset:physicalOffset:length:)
func (f_ FSExtentPacker) PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset unsafe.Pointer, physicalOffset unsafe.Pointer, length uintptr /* not a class type */) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("packExtentWithResource:type:logicalOffset:physicalOffset:length:"), resource, type_, logicalOffset, physicalOffset, length)
	return rv
}/* debug [instance_methods/method]: PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSExtentPacker */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSExtentPacker */



