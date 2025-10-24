// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLBinaryArchiveDescriptor */


/* debug [class_header]: Header for MTLBinaryArchiveDescriptor */
// The class instance for the [BinaryArchiveDescriptor] class.
var (
	BinaryArchiveDescriptorClass     _BinaryArchiveDescriptorClass
	BinaryArchiveDescriptorClassOnce sync.Once
)

func getBinaryArchiveDescriptorClass() _BinaryArchiveDescriptorClass {
	BinaryArchiveDescriptorClassOnce.Do(func() {
		BinaryArchiveDescriptorClass = _BinaryArchiveDescriptorClass{objc.GetClass("MTLBinaryArchiveDescriptor")}
	})
	return BinaryArchiveDescriptorClass
}

type _BinaryArchiveDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BinaryArchiveDescriptor */
// An interface definition for the [BinaryArchiveDescriptor] class.
type IBinaryArchiveDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BinaryArchiveDescriptor */
	// properties:
	Url() objc.IObject /* cross-framework: NSURL */
	SetUrl(value objc.IObject /* cross-framework: NSURL */)
	MTLBinaryArchiveDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BinaryArchiveDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BinaryArchiveDescriptor */
// Alloc allocates a new instance without initialization.
func (bc _BinaryArchiveDescriptorClass) Alloc() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BinaryArchiveDescriptorClass) New() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryArchiveDescriptor) Init() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryArchiveDescriptor) Autorelease() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryArchiveDescriptor creates a new BinaryArchiveDescriptor instance.
func NewBinaryArchiveDescriptor() BinaryArchiveDescriptor {
	return getBinaryArchiveDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BinaryArchiveDescriptor */
// A description of a binary shader archive that you want to create.


// A description of a binary shader archive that you want to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor
type BinaryArchiveDescriptor struct {
	objectivec.Object
}

// BinaryArchiveDescriptorFrom constructs a [BinaryArchiveDescriptor] from an unsafe.Pointer.
//
// A description of a binary shader archive that you want to create.
func BinaryArchiveDescriptorFrom(ptr unsafe.Pointer) BinaryArchiveDescriptor {
	return BinaryArchiveDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BinaryArchiveDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BinaryArchiveDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BinaryArchiveDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BinaryArchiveDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BinaryArchiveDescriptor */

// A URL to a Metal binary archive file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor/url
func (b_ BinaryArchiveDescriptor) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](b_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A URL to a Metal binary archive file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor/url
func (b_ BinaryArchiveDescriptor) SetUrl(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */


// The domain for Metal binary archive errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchivedomain
func (b_ BinaryArchiveDescriptor) MTLBinaryArchiveDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("MTLBinaryArchiveDomain"))
	return rv
}/* debug [instance_properties/getter]: MTLBinaryArchiveDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLBinaryArchiveDescriptor */



