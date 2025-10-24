// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCContentSharingPickerConfiguration */


/* debug [class_header]: Header for SCContentSharingPickerConfiguration */
// The class instance for the [ContentSharingPickerConfiguration] class.
var (
	ContentSharingPickerConfigurationClass     _ContentSharingPickerConfigurationClass
	ContentSharingPickerConfigurationClassOnce sync.Once
)

func getContentSharingPickerConfigurationClass() _ContentSharingPickerConfigurationClass {
	ContentSharingPickerConfigurationClassOnce.Do(func() {
		ContentSharingPickerConfigurationClass = _ContentSharingPickerConfigurationClass{objc.GetClass("SCContentSharingPickerConfiguration")}
	})
	return ContentSharingPickerConfigurationClass
}

type _ContentSharingPickerConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentSharingPickerConfiguration */
// An interface definition for the [ContentSharingPickerConfiguration] class.
type IContentSharingPickerConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentSharingPickerConfiguration */
	// properties:
	AllowedPickerModes() ContentSharingPickerMode
	SetAllowedPickerModes(value ContentSharingPickerMode)
	AllowsChangingSelectedContent() bool
	SetAllowsChangingSelectedContent(value bool)
	ExcludedBundleIDs() []string
	SetExcludedBundleIDs(value []string)
	ExcludedWindowIDs() []foundation.Number
	SetExcludedWindowIDs(value []foundation.Number)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentSharingPickerConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentSharingPickerConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _ContentSharingPickerConfigurationClass) Alloc() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentSharingPickerConfigurationClass) New() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentSharingPickerConfiguration) Init() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentSharingPickerConfiguration) Autorelease() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentSharingPickerConfiguration creates a new ContentSharingPickerConfiguration instance.
func NewContentSharingPickerConfiguration() ContentSharingPickerConfiguration {
	return getContentSharingPickerConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentSharingPickerConfiguration */
// An instance for configuring the system content-sharing picker.


// An instance for configuring the system content-sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class
type ContentSharingPickerConfiguration struct {
	objectivec.Object
}

// ContentSharingPickerConfigurationFrom constructs a [ContentSharingPickerConfiguration] from an unsafe.Pointer.
//
// An instance for configuring the system content-sharing picker.
func ContentSharingPickerConfigurationFrom(ptr unsafe.Pointer) ContentSharingPickerConfiguration {
	return ContentSharingPickerConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentSharingPickerConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentSharingPickerConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentSharingPickerConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentSharingPickerConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentSharingPickerConfiguration */

// The content-selection modes supported by the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowedPickerModes
func (c_ ContentSharingPickerConfiguration) AllowedPickerModes() ContentSharingPickerMode {
	rv := objc.Send[ContentSharingPickerMode](c_.ID, objc.Sel("allowedPickerModes"))
	return rv
}/* debug [instance_properties/getter]: allowedPickerModes */


// The content-selection modes supported by the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowedPickerModes
func (c_ ContentSharingPickerConfiguration) SetAllowedPickerModes(value ContentSharingPickerMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedPickerModes:"), value)
}/* debug [instance_properties/setter]: allowedPickerModes */


// A Boolean value that indicates if the present stream can change to a different source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowsChangingSelectedContent
func (c_ ContentSharingPickerConfiguration) AllowsChangingSelectedContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsChangingSelectedContent"))
	return rv
}/* debug [instance_properties/getter]: allowsChangingSelectedContent */


// A Boolean value that indicates if the present stream can change to a different source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowsChangingSelectedContent
func (c_ ContentSharingPickerConfiguration) SetAllowsChangingSelectedContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsChangingSelectedContent:"), value)
}/* debug [instance_properties/setter]: allowsChangingSelectedContent */


// A list of bundle IDs to exclude from the sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedBundleIDs
func (c_ ContentSharingPickerConfiguration) ExcludedBundleIDs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("excludedBundleIDs"))
	return rv
}/* debug [instance_properties/getter]: excludedBundleIDs */


// A list of bundle IDs to exclude from the sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedBundleIDs
func (c_ ContentSharingPickerConfiguration) SetExcludedBundleIDs(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setExcludedBundleIDs:"), nsArray)
}/* debug [instance_properties/setter]: excludedBundleIDs */


// A list of window IDs to exclude from the sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedWindowIDs
func (c_ ContentSharingPickerConfiguration) ExcludedWindowIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("excludedWindowIDs"))
	return rv
}/* debug [instance_properties/getter]: excludedWindowIDs */


// A list of window IDs to exclude from the sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedWindowIDs
func (c_ ContentSharingPickerConfiguration) SetExcludedWindowIDs(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setExcludedWindowIDs:"), nsArray)
}/* debug [instance_properties/setter]: excludedWindowIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCContentSharingPickerConfiguration */



