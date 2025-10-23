// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MEAddressAnnotation] class.
type IMEAddressAnnotation interface {
	objectivec.IObject
	// properties:
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MEAddressAnnotationClass) Alloc() MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Indicates an address is invalid and may result in failure to deliver a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/error(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) ErrorWithLocalizedDescription(localizedDescription string /* primitive/slice/pointer. */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("errorWithLocalizedDescription:"), objc.String(localizedDescription))
	return rv
}


// Indicates an address is valid and correct.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/success(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) SuccessWithLocalizedDescription(localizedDescription string /* primitive/slice/pointer. */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("successWithLocalizedDescription:"), objc.String(localizedDescription))
	return rv
}


// Indicates an address may be invalid or needs attention.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEAddressAnnotation/warning(withLocalizedDescription:)
func (mc _MEAddressAnnotationClass) WarningWithLocalizedDescription(localizedDescription string /* primitive/slice/pointer. */) MEAddressAnnotation {
	rv := objc.Send[MEAddressAnnotation](objc.ID(mc.class), objc.Sel("warningWithLocalizedDescription:"), objc.String(localizedDescription))
	return rv
}



