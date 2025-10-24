// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEDecodedMessageBanner */


/* debug [class_header]: Header for MEDecodedMessageBanner */
// The class instance for the [MEDecodedMessageBanner] class.
var (
	MEDecodedMessageBannerClass     _MEDecodedMessageBannerClass
	MEDecodedMessageBannerClassOnce sync.Once
)

func getMEDecodedMessageBannerClass() _MEDecodedMessageBannerClass {
	MEDecodedMessageBannerClassOnce.Do(func() {
		MEDecodedMessageBannerClass = _MEDecodedMessageBannerClass{objc.GetClass("MEDecodedMessageBanner")}
	})
	return MEDecodedMessageBannerClass
}

type _MEDecodedMessageBannerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEDecodedMessageBanner */
// An interface definition for the [MEDecodedMessageBanner] class.
type IMEDecodedMessageBanner interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEDecodedMessageBanner */
	// properties:
	Dismissable() bool
	PrimaryActionTitle() objc.IObject /* cross-framework: NSString */
	Title() objc.IObject /* cross-framework: NSString */
	IsDismissable() bool
	SetIsDismissable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEDecodedMessageBanner */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEDecodedMessageBanner */
// Alloc allocates a new instance without initialization.
func (mc _MEDecodedMessageBannerClass) Alloc() MEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEDecodedMessageBannerClass) New() MEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEDecodedMessageBanner) Init() MEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEDecodedMessageBanner) Autorelease() MEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodedMessageBanner creates a new MEDecodedMessageBanner instance.
func NewMEDecodedMessageBanner() MEDecodedMessageBanner {
	return getMEDecodedMessageBannerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEDecodedMessageBanner */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner
type MEDecodedMessageBanner struct {
	objectivec.Object
}

// MEDecodedMessageBannerFrom constructs a [MEDecodedMessageBanner] from an unsafe.Pointer.
func MEDecodedMessageBannerFrom(ptr unsafe.Pointer) MEDecodedMessageBanner {
	return MEDecodedMessageBanner{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEDecodedMessageBanner */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/init(title:primaryActionTitle:dismissable:)
func NewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable(title objc.IObject /* cross-framework: NSString */, primaryActionTitle objc.IObject /* cross-framework: NSString */, dismissable bool) MEDecodedMessageBanner {
	instance := getMEDecodedMessageBannerClass().Alloc()
	rv := objc.Send[MEDecodedMessageBanner](instance.ID, objc.Sel("initWithTitle:primaryActionTitle:dismissable:"), title, primaryActionTitle, dismissable)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEDecodedMessageBanner */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEDecodedMessageBanner */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEDecodedMessageBanner */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEDecodedMessageBanner */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/isDismissable
func (m_ MEDecodedMessageBanner) Dismissable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("dismissable"))
	return rv
}/* debug [instance_properties/getter]: dismissable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/primaryActionTitle
func (m_ MEDecodedMessageBanner) PrimaryActionTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("primaryActionTitle"))
	return rv
}/* debug [instance_properties/getter]: primaryActionTitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/title
func (m_ MEDecodedMessageBanner) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/isdismissable
func (m_ MEDecodedMessageBanner) IsDismissable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDismissable"))
	return rv
}/* debug [instance_properties/getter]: isDismissable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/isdismissable
func (m_ MEDecodedMessageBanner) SetIsDismissable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDismissable:"), value)
}/* debug [instance_properties/setter]: isDismissable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEDecodedMessageBanner */


