// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class QLPreviewReplyAttachment */


/* debug [class_header]: Header for QLPreviewReplyAttachment */
// The class instance for the [PreviewReplyAttachment] class.
var (
	PreviewReplyAttachmentClass     _PreviewReplyAttachmentClass
	PreviewReplyAttachmentClassOnce sync.Once
)

func getPreviewReplyAttachmentClass() _PreviewReplyAttachmentClass {
	PreviewReplyAttachmentClassOnce.Do(func() {
		PreviewReplyAttachmentClass = _PreviewReplyAttachmentClass{objc.GetClass("QLPreviewReplyAttachment")}
	})
	return PreviewReplyAttachmentClass
}

type _PreviewReplyAttachmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewReplyAttachment */
// An interface definition for the [PreviewReplyAttachment] class.
type IPreviewReplyAttachment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewReplyAttachment */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewReplyAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewReplyAttachment */
// Alloc allocates a new instance without initialization.
func (pc _PreviewReplyAttachmentClass) Alloc() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewReplyAttachmentClass) New() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewReplyAttachment) Init() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewReplyAttachment) Autorelease() PreviewReplyAttachment {
	rv := objc.Send[PreviewReplyAttachment](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewReplyAttachment creates a new PreviewReplyAttachment instance.
func NewPreviewReplyAttachment() PreviewReplyAttachment {
	return getPreviewReplyAttachmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewReplyAttachment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReplyAttachment
type PreviewReplyAttachment struct {
	objectivec.Object
}

// PreviewReplyAttachmentFrom constructs a [PreviewReplyAttachment] from an unsafe.Pointer.
func PreviewReplyAttachmentFrom(ptr unsafe.Pointer) PreviewReplyAttachment {
	return PreviewReplyAttachment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewReplyAttachment */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReplyAttachment/init(data:contentType:)
func NewPreviewReplyAttachmentWithDataContentType(data objc.IObject /* cross-framework: NSData */, contentType uniformtypeidentifiers.UTType) PreviewReplyAttachment {
	instance := getPreviewReplyAttachmentClass().Alloc()
	rv := objc.Send[PreviewReplyAttachment](instance.ID, objc.Sel("initWithData:contentType:"), data, contentType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewReplyAttachmentWithDataContentType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewReplyAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewReplyAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewReplyAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewReplyAttachment */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewReplyAttachment */


