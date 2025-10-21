// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MEDecodedMessageBanner] class.
type IMEDecodedMessageBanner interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner
type MEDecodedMessageBanner struct {
	objectivec.Object
}

// MEDecodedMessageBannerFrom constructs a [MEDecodedMessageBanner] from an unsafe.Pointer.
func MEDecodedMessageBannerFrom(ptr unsafe.Pointer) MEDecodedMessageBanner {
	return MEDecodedMessageBanner{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEDecodedMessageBannerClass) Alloc() MEDecodedMessageBanner {
	rv := objc.Send[MEDecodedMessageBanner](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/init(title:primaryActionTitle:dismissable:)
func NewMEDecodedMessageBannerWithTitlePrimaryActionTitleDismissable(title appkit.string, primaryActionTitle appkit.string, dismissable bool) MEDecodedMessageBanner {
	instance := getMEDecodedMessageBannerClass().Alloc()
	rv := objc.Send[MEDecodedMessageBanner](instance.ID, objc.Sel("initWithTitle:primaryActionTitle:dismissable:"), title, primaryActionTitle, dismissable)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/isDismissable
func (m_ MEDecodedMessageBanner) Dismissable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("dismissable"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEDecodedMessageBanner/primaryActionTitle
func (m_ MEDecodedMessageBanner) PrimaryActionTitle() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("primaryActionTitle"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/isdismissable
func (m_ MEDecodedMessageBanner) IsDismissable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isDismissable"))
	return rv
}


// SetIsDismissable sets the value of the isDismissable property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/isdismissable
func (m_ MEDecodedMessageBanner) SetIsDismissable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsDismissable:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/title
func (m_ MEDecodedMessageBanner) Title() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
//
// [Full Topic]: https://developer.apple.com/documentation/mailkit/medecodedmessagebanner/title
func (m_ MEDecodedMessageBanner) SetTitle(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitle:"), value)
}


