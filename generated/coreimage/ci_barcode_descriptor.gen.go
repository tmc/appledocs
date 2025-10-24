// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIBarcodeDescriptor */


/* debug [class_header]: Header for CIBarcodeDescriptor */
// The class instance for the [BarcodeDescriptor] class.
var (
	BarcodeDescriptorClass     _BarcodeDescriptorClass
	BarcodeDescriptorClassOnce sync.Once
)

func getBarcodeDescriptorClass() _BarcodeDescriptorClass {
	BarcodeDescriptorClassOnce.Do(func() {
		BarcodeDescriptorClass = _BarcodeDescriptorClass{objc.GetClass("CIBarcodeDescriptor")}
	})
	return BarcodeDescriptorClass
}

type _BarcodeDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BarcodeDescriptor */
// An interface definition for the [BarcodeDescriptor] class.
type IBarcodeDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BarcodeDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BarcodeDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BarcodeDescriptor */
// Alloc allocates a new instance without initialization.
func (bc _BarcodeDescriptorClass) Alloc() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BarcodeDescriptorClass) New() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BarcodeDescriptor) Init() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BarcodeDescriptor) Autorelease() BarcodeDescriptor {
	rv := objc.Send[BarcodeDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBarcodeDescriptor creates a new BarcodeDescriptor instance.
func NewBarcodeDescriptor() BarcodeDescriptor {
	return getBarcodeDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BarcodeDescriptor */
// An abstract base class that represents a machine-readable code’s attributes.
//
// Subclasses encapsulate the formal specification and fields specific to a code type. Each subclass is sufficient to recreate the unique symbol exactly as seen or used with a custom parser.


// An abstract base class that represents a machine-readable code’s attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIBarcodeDescriptor
type BarcodeDescriptor struct {
	objectivec.Object
}

// BarcodeDescriptorFrom constructs a [BarcodeDescriptor] from an unsafe.Pointer.
//
// An abstract base class that represents a machine-readable code’s attributes.
func BarcodeDescriptorFrom(ptr unsafe.Pointer) BarcodeDescriptor {
	return BarcodeDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BarcodeDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BarcodeDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BarcodeDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BarcodeDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BarcodeDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIBarcodeDescriptor */



