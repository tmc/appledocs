// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingCollaborationModeRestriction] class.
var (
	SharingCollaborationModeRestrictionClass     _SharingCollaborationModeRestrictionClass
	SharingCollaborationModeRestrictionClassOnce sync.Once
)

func getSharingCollaborationModeRestrictionClass() _SharingCollaborationModeRestrictionClass {
	SharingCollaborationModeRestrictionClassOnce.Do(func() {
		SharingCollaborationModeRestrictionClass = _SharingCollaborationModeRestrictionClass{objc.GetClass("NSSharingCollaborationModeRestriction")}
	})
	return SharingCollaborationModeRestrictionClass
}

type _SharingCollaborationModeRestrictionClass struct {
	class objc.Class
}

// An interface definition for the [SharingCollaborationModeRestriction] class.
type ISharingCollaborationModeRestriction interface {
	objectivec.IObject
	// properties:
	AlertDismissButtonTitle() objc.IObject /* cross-framework: NSString */
	AlertMessage() objc.IObject /* cross-framework: NSString */
	AlertRecoverySuggestionButtonLaunchURL() objc.IObject /* cross-framework: NSURL */
	AlertRecoverySuggestionButtonTitle() objc.IObject /* cross-framework: NSString */
	AlertTitle() objc.IObject /* cross-framework: NSString */
	DisabledMode() SharingCollaborationMode
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction
type SharingCollaborationModeRestriction struct {
	objectivec.Object
}

// SharingCollaborationModeRestrictionFrom constructs a [SharingCollaborationModeRestriction] from an unsafe.Pointer.
func SharingCollaborationModeRestrictionFrom(ptr unsafe.Pointer) SharingCollaborationModeRestriction {
	return SharingCollaborationModeRestriction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SharingCollaborationModeRestrictionClass) Alloc() SharingCollaborationModeRestriction {
	rv := objc.Send[SharingCollaborationModeRestriction](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SharingCollaborationModeRestrictionClass) New() SharingCollaborationModeRestriction {
	rv := objc.Send[SharingCollaborationModeRestriction](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingCollaborationModeRestriction) Init() SharingCollaborationModeRestriction {
	rv := objc.Send[SharingCollaborationModeRestriction](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingCollaborationModeRestriction) Autorelease() SharingCollaborationModeRestriction {
	rv := objc.Send[SharingCollaborationModeRestriction](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingCollaborationModeRestriction creates a new SharingCollaborationModeRestriction instance.
func NewSharingCollaborationModeRestriction() SharingCollaborationModeRestriction {
	return getSharingCollaborationModeRestrictionClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:)
func NewSharingCollaborationModeRestrictionWithDisabledMode(disabledMode SharingCollaborationMode) SharingCollaborationModeRestriction {
	instance := getSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[SharingCollaborationModeRestriction](instance.ID, objc.Sel("initWithDisabledMode:"), disabledMode)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessage(disabledMode SharingCollaborationMode, alertTitle objc.IObject /* cross-framework: NSString */, alertMessage objc.IObject /* cross-framework: NSString */) SharingCollaborationModeRestriction {
	instance := getSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[SharingCollaborationModeRestriction](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:"), disabledMode, alertTitle, alertMessage)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitle(disabledMode SharingCollaborationMode, alertTitle objc.IObject /* cross-framework: NSString */, alertMessage objc.IObject /* cross-framework: NSString */, alertDismissButtonTitle objc.IObject /* cross-framework: NSString */) SharingCollaborationModeRestriction {
	instance := getSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[SharingCollaborationModeRestriction](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:"), disabledMode, alertTitle, alertMessage, alertDismissButtonTitle)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/init(disabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunch:)
func NewSharingCollaborationModeRestrictionWithDisabledModeAlertTitleAlertMessageAlertDismissButtonTitleAlertRecoverySuggestionButtonTitleAlertRecoverySuggestionButtonLaunchURL(disabledMode SharingCollaborationMode, alertTitle objc.IObject /* cross-framework: NSString */, alertMessage objc.IObject /* cross-framework: NSString */, alertDismissButtonTitle objc.IObject /* cross-framework: NSString */, alertRecoverySuggestionButtonTitle objc.IObject /* cross-framework: NSString */, alertRecoverySuggestionButtonLaunchURL objc.IObject /* cross-framework: NSURL */) SharingCollaborationModeRestriction {
	instance := getSharingCollaborationModeRestrictionClass().Alloc()
	rv := objc.Send[SharingCollaborationModeRestriction](instance.ID, objc.Sel("initWithDisabledMode:alertTitle:alertMessage:alertDismissButtonTitle:alertRecoverySuggestionButtonTitle:alertRecoverySuggestionButtonLaunchURL:"), disabledMode, alertTitle, alertMessage, alertDismissButtonTitle, alertRecoverySuggestionButtonTitle, alertRecoverySuggestionButtonLaunchURL)
	rv.Autorelease()
	return rv
}



// The label on the alert button which will simply confirm that the alert was viewed and dismiss it Defaults to “OK”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertDismissButtonTitle
func (s_ SharingCollaborationModeRestriction) AlertDismissButtonTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("alertDismissButtonTitle"))
	return rv
}


// The message of the alert if a reason for disabling is provided
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertMessage
func (s_ SharingCollaborationModeRestriction) AlertMessage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("alertMessage"))
	return rv
}


// The URL that is opened when the user selects the recovery suggestion, if any
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertRecoverySuggestionButtonLaunchURL
func (s_ SharingCollaborationModeRestriction) AlertRecoverySuggestionButtonLaunchURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("alertRecoverySuggestionButtonLaunchURL"))
	return rv
}


// The label on the recovery suggestion button if it is provided
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertRecoverySuggestionButtonTitle
func (s_ SharingCollaborationModeRestriction) AlertRecoverySuggestionButtonTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("alertRecoverySuggestionButtonTitle"))
	return rv
}


// The title of the alert if a reason for disabling is provided
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/alertTitle
func (s_ SharingCollaborationModeRestriction) AlertTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("alertTitle"))
	return rv
}


// The type of sharing which should be disabled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/CollaborationModeRestriction/disabledMode
func (s_ SharingCollaborationModeRestriction) DisabledMode() SharingCollaborationMode {
	rv := objc.Send[SharingCollaborationMode](s_.ID, objc.Sel("disabledMode"))
	return rv
}


