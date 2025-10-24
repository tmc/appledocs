// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SCContentSharingPicker */


/* debug [class_header]: Header for SCContentSharingPicker */
// The class instance for the [ContentSharingPicker] class.
var (
	ContentSharingPickerClass     _ContentSharingPickerClass
	ContentSharingPickerClassOnce sync.Once
)

func getContentSharingPickerClass() _ContentSharingPickerClass {
	ContentSharingPickerClassOnce.Do(func() {
		ContentSharingPickerClass = _ContentSharingPickerClass{objc.GetClass("SCContentSharingPicker")}
	})
	return ContentSharingPickerClass
}

type _ContentSharingPickerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentSharingPicker */
// An interface definition for the [ContentSharingPicker] class.
type IContentSharingPicker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentSharingPicker */
	// properties:
	DefaultConfiguration() ISCContentSharingPickerConfiguration
	SetDefaultConfiguration(value ISCContentSharingPickerConfiguration)
	Active() bool
	SetActive(value bool)
	MaximumStreamCount() objc.IObject /* cross-framework: NSNumber */
	SetMaximumStreamCount(value objc.IObject /* cross-framework: NSNumber */)
	Configuration() IContentSharingPickerConfiguration
	SetConfiguration(value IContentSharingPickerConfiguration)
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentSharingPicker */
	// methods:
	AddObserver(observer unsafe.Pointer)
	Present()
	PresentPickerForStream(stream ISCStream)
	PresentPickerForStreamUsingContentStyle(stream ISCStream, contentStyle ShareableContentStyle)
	PresentPickerUsingContentStyle(contentStyle ShareableContentStyle)
	RemoveObserver(observer unsafe.Pointer)
	SetConfigurationForStream(pickerConfig ISCContentSharingPickerConfiguration, stream ISCStream)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentSharingPicker */
// Alloc allocates a new instance without initialization.
func (cc _ContentSharingPickerClass) Alloc() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentSharingPickerClass) New() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentSharingPicker) Init() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentSharingPicker) Autorelease() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentSharingPicker creates a new ContentSharingPicker instance.
func NewContentSharingPicker() ContentSharingPicker {
	return getContentSharingPickerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentSharingPicker */
// An instance of a picker presented by the operating system for managing frame-capture streams.


// An instance of a picker presented by the operating system for managing frame-capture streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker
type ContentSharingPicker struct {
	objectivec.Object
}

// ContentSharingPickerFrom constructs a [ContentSharingPicker] from an unsafe.Pointer.
//
// An instance of a picker presented by the operating system for managing frame-capture streams.
func ContentSharingPickerFrom(ptr unsafe.Pointer) ContentSharingPicker {
	return ContentSharingPicker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentSharingPicker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentSharingPicker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentSharingPicker */

// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (cc _ContentSharingPickerClass) SharedPicker() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("sharedPicker"))
	return rv
}/* debug [class_properties_class/property]: sharedPicker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentSharingPicker */

// Adds an observer instance to notify of changes in the content-sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/add(_:)
func (c_ ContentSharingPicker) AddObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addObserver:"), observer)
}/* debug [instance_methods/method]: AddObserver */


// Displays the picker with no active selection for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present()
func (c_ ContentSharingPicker) Present() {
	objc.Send[objc.ID](c_.ID, objc.Sel("present"))
}/* debug [instance_methods/method]: Present */


// Displays the picker with an already running capture stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(for:)
func (c_ ContentSharingPicker) PresentPickerForStream(stream ISCStream) {
	objc.Send[objc.ID](c_.ID, objc.Sel("presentPickerForStream:"), stream)
}/* debug [instance_methods/method]: PresentPickerForStream */


// Displays the picker with an existing capture stream, allowing for a single type of capture selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(for:using:)
func (c_ ContentSharingPicker) PresentPickerForStreamUsingContentStyle(stream ISCStream, contentStyle ShareableContentStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("presentPickerForStream:usingContentStyle:"), stream, contentStyle)
}/* debug [instance_methods/method]: PresentPickerForStreamUsingContentStyle */


// Displays the picker for a single type of capture selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(using:)
func (c_ ContentSharingPicker) PresentPickerUsingContentStyle(contentStyle ShareableContentStyle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("presentPickerUsingContentStyle:"), contentStyle)
}/* debug [instance_methods/method]: PresentPickerUsingContentStyle */


// Removes an observer instance from the content-sharing picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/remove(_:)
func (c_ ContentSharingPicker) RemoveObserver(observer unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeObserver:"), observer)
}/* debug [instance_methods/method]: RemoveObserver */


// Sets the configuration for the content capture picker for a capture stream, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/setConfiguration:forStream:
func (c_ ContentSharingPicker) SetConfigurationForStream(pickerConfig ISCContentSharingPickerConfiguration, stream ISCStream) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:forStream:"), pickerConfig, stream)
}/* debug [instance_methods/method]: SetConfigurationForStream */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentSharingPicker */

// The default configuration to use for the content capture picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/defaultConfiguration-9v5sa
func (c_ ContentSharingPicker) DefaultConfiguration() ISCContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}/* debug [instance_properties/getter]: defaultConfiguration */


// The default configuration to use for the content capture picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/defaultConfiguration-9v5sa
func (c_ ContentSharingPicker) SetDefaultConfiguration(value ISCContentSharingPickerConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}/* debug [instance_properties/setter]: defaultConfiguration */


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/isActive
func (c_ ContentSharingPicker) Active() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/isActive
func (c_ ContentSharingPicker) SetActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActive:"), value)
}/* debug [instance_properties/setter]: active */


// The maximum number of streams the content capture picker allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/maximumStreamCount-66khx
func (c_ ContentSharingPicker) MaximumStreamCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("maximumStreamCount"))
	return rv
}/* debug [instance_properties/getter]: maximumStreamCount */


// The maximum number of streams the content capture picker allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/maximumStreamCount-66khx
func (c_ ContentSharingPicker) SetMaximumStreamCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumStreamCount:"), value)
}/* debug [instance_properties/setter]: maximumStreamCount */


// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (c_ ContentSharingPicker) SharedPicker() ISCContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("sharedPicker"))
	return rv
}/* debug [instance_properties/getter]: sharedPicker */


// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) Configuration() IContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) SetConfiguration(value IContentSharingPickerConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SCContentSharingPicker */



