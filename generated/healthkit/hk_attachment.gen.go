// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class HKAttachment */


/* debug [class_header]: Header for HKAttachment */
// The class instance for the [HKAttachment] class.
var (
	HKAttachmentClass     _HKAttachmentClass
	HKAttachmentClassOnce sync.Once
)

func getHKAttachmentClass() _HKAttachmentClass {
	HKAttachmentClassOnce.Do(func() {
		HKAttachmentClass = _HKAttachmentClass{objc.GetClass("HKAttachment")}
	})
	return HKAttachmentClass
}

type _HKAttachmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAttachment */
// An interface definition for the [HKAttachment] class.
type IHKAttachment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKAttachment */
	// properties:
	ContentType() uniformtypeidentifiers.UTType
	CreationDate() objc.IObject /* cross-framework: NSDate */
	Identifier() foundation.UUID
	Metadata() foundation.IDictionary
	Name() objc.IObject /* cross-framework: NSString */
	Size() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAttachment */
// Alloc allocates a new instance without initialization.
func (hc _HKAttachmentClass) Alloc() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAttachmentClass) New() HKAttachment {
	rv := objc.Send[HKAttachment](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAttachment) Init() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAttachment) Autorelease() HKAttachment {
	rv := objc.Send[HKAttachment](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAttachment creates a new HKAttachment instance.
func NewHKAttachment() HKAttachment {
	return getHKAttachmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAttachment */
// A file that is attached to a sample in the HealthKit store.
//
// To access the attachment’s data, get a data reader from the attachment store for each attachment. You can then asynchronously access the whole data object. Alternatively, you can access the file’s contents as an asynchronous sequence of bytes.


// A file that is attached to a sample in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment
type HKAttachment struct {
	objectivec.Object
}

// HKAttachmentFrom constructs a [HKAttachment] from an unsafe.Pointer.
//
// A file that is attached to a sample in the HealthKit store.
func HKAttachmentFrom(ptr unsafe.Pointer) HKAttachment {
	return HKAttachment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAttachment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAttachment */

// The type of data stored in the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/contentType
func (h_ HKAttachment) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](h_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The attachment’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/creationDate
func (h_ HKAttachment) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The universally unique identifier for the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/identifier
func (h_ HKAttachment) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](h_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Additional data associated with the attachment in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/metadata
func (h_ HKAttachment) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](h_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// The name of the attached file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/name
func (h_ HKAttachment) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The attachment’s size (in bytes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachment/size
func (h_ HKAttachment) Size() int {
	rv := objc.Send[int](h_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAttachment */



