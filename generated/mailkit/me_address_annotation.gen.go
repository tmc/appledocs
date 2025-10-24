// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEAddressAnnotation */


/* debug [class_header]: Header for MEAddressAnnotation */
// The class instance for the [MEAddressAnnotation] class.
var (
	MEAddressAnnotationClass     _MEAddressAnnotationClass
	MEAddressAnnotationClassOnce sync.Once
)

func getMEAddressAnnotationClass() _MEAddressAnnotationClass {
	MEAddressAnnotationClassOnce.Do(func() {
		MEAddressAnnotationClass = _MEAddressAnnotationClass{objc.GetClass("MEAddressAnnotation")}
	})
	return MEAddressAnnotationClass
}

type _MEAddressAnnotationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEAddressAnnotation */
// An interface definition for the [MEAddressAnnotation] class.
type IMEAddressAnnotation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEAddressAnnotation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEAddressAnnotation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEAddressAnnotation */
// Alloc allocates a new instance without initialization.
func (mc _MEAddressAnnotationClass) Alloc() MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEAddressAnnotationClass) New() MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEAddressAnnotation) Init() MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEAddressAnnotation) Autorelease() MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEAddressAnnotation creates a new MEAddressAnnotation instance.
func NewMEAddressAnnotation() MEAddressAnnotation {
	return getMEAddressAnnotationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEAddressAnnotation */
// An object that indicates the validity of an email address.
//
// Mail displays the status of an annotation as part of the address tokens in the To, Cc, and Bcc fields using a status icon and color.


// An object that indicates the validity of an email address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation
type MEAddressAnnotation struct {
	objectivec.Object
}

// MEAddressAnnotationFrom constructs a [MEAddressAnnotation] from an unsafe.Pointer.
//
// An object that indicates the validity of an email address.
func MEAddressAnnotationFrom(ptr unsafe.Pointer) MEAddressAnnotation {
	return MEAddressAnnotation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEAddressAnnotation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEAddressAnnotation */

// Indicates an address is invalid and may result in failure to deliver a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/error(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) ErrorWithLocalizedDescription(localizedDescription objc.IObject /* cross-framework: NSString */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("errorWithLocalizedDescription:"), localizedDescription)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ErrorWithLocalizedDescription) */


// Indicates an address is valid and correct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/success(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) SuccessWithLocalizedDescription(localizedDescription objc.IObject /* cross-framework: NSString */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("successWithLocalizedDescription:"), localizedDescription)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SuccessWithLocalizedDescription) */


// Indicates an address may be invalid or needs attention.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/warning(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) WarningWithLocalizedDescription(localizedDescription objc.IObject /* cross-framework: NSString */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("warningWithLocalizedDescription:"), localizedDescription)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WarningWithLocalizedDescription) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEAddressAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEAddressAnnotation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEAddressAnnotation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEAddressAnnotation */



