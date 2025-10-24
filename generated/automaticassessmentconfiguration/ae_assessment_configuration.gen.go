// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AEAssessmentConfiguration */


/* debug [class_header]: Header for AEAssessmentConfiguration */
// The class instance for the [AEAssessmentConfiguration] class.
var (
	AEAssessmentConfigurationClass     _AEAssessmentConfigurationClass
	AEAssessmentConfigurationClassOnce sync.Once
)

func getAEAssessmentConfigurationClass() _AEAssessmentConfigurationClass {
	AEAssessmentConfigurationClassOnce.Do(func() {
		AEAssessmentConfigurationClass = _AEAssessmentConfigurationClass{objc.GetClass("AEAssessmentConfiguration")}
	})
	return AEAssessmentConfigurationClass
}

type _AEAssessmentConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AEAssessmentConfiguration */
// An interface definition for the [AEAssessmentConfiguration] class.
type IAEAssessmentConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AEAssessmentConfiguration */
	// properties:
	AllowsAccessibilityKeyboard() bool
	SetAllowsAccessibilityKeyboard(value bool)
	AllowsAccessibilityLiveCaptions() bool
	SetAllowsAccessibilityLiveCaptions(value bool)
	AllowsAccessibilityReader() bool
	SetAllowsAccessibilityReader(value bool)
	AllowsKeyboardShortcuts() bool
	SetAllowsKeyboardShortcuts(value bool)
	AllowsPredictiveKeyboard() bool
	SetAllowsPredictiveKeyboard(value bool)
	AllowsScreenshots() bool
	SetAllowsScreenshots(value bool)
	AllowsSpellCheck() bool
	SetAllowsSpellCheck(value bool)
	AutocorrectMode() AEAutocorrectMode
	SetAutocorrectMode(value AEAutocorrectMode)
	ConfigurationsByApplication() foundation.IDictionary
	MainParticipantConfiguration() IAEAssessmentParticipantConfiguration
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AEAssessmentConfiguration */
	// methods:
	RemoveApplication(application IAEAssessmentApplication)
	SetConfigurationForApplication(configuration IAEAssessmentParticipantConfiguration, application IAEAssessmentApplication)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AEAssessmentConfiguration */
// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentConfigurationClass) Alloc() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AEAssessmentConfigurationClass) New() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AEAssessmentConfiguration) Init() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AEAssessmentConfiguration) Autorelease() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEAssessmentConfiguration creates a new AEAssessmentConfiguration instance.
func NewAEAssessmentConfiguration() AEAssessmentConfiguration {
	return getAEAssessmentConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AEAssessmentConfiguration */
// Configuration information for an assessment session.
//
// Create a configuration instance and pass it to the initializer of an instance to create a new assessment session. Before using the configuration, indicate which exceptions you want to allow for the assessment session’s restrictions by setting values on the configuration instance. For example, you can set values to allow dictation and certain aspects of autocorrect: While you provide a configuration instance when creating a session on iOS, iPadOS, and macOS, specific exceptions apply only to certain platforms. In particular, on macOS, you can selectively make specific apps besides your own available during an assessment — for example, to allow users to access a calculator or a dictionary. All other exceptions apply only to iOS and iPadOS.


// Configuration information for an assessment session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration
type AEAssessmentConfiguration struct {
	objectivec.Object
}

// AEAssessmentConfigurationFrom constructs a [AEAssessmentConfiguration] from an unsafe.Pointer.
//
// Configuration information for an assessment session.
func AEAssessmentConfigurationFrom(ptr unsafe.Pointer) AEAssessmentConfiguration {
	return AEAssessmentConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AEAssessmentConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AEAssessmentConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AEAssessmentConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AEAssessmentConfiguration */

// Removes the availability of a previously allowed app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/remove(_:)
func (a_ AEAssessmentConfiguration) RemoveApplication(application IAEAssessmentApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeApplication:"), application)
}/* debug [instance_methods/method]: RemoveApplication */


// Adds an app to the list of apps available during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/setConfiguration(_:for:)
func (a_ AEAssessmentConfiguration) SetConfigurationForApplication(configuration IAEAssessmentParticipantConfiguration, application IAEAssessmentApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfiguration:forApplication:"), configuration, application)
}/* debug [instance_methods/method]: SetConfigurationForApplication */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AEAssessmentConfiguration */

// A Boolean value that indicates whether to allow alternative input methods in the Accessibility Keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityKeyboard
func (a_ AEAssessmentConfiguration) AllowsAccessibilityKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityKeyboard"))
	return rv
}/* debug [instance_properties/getter]: allowsAccessibilityKeyboard */


// A Boolean value that indicates whether to allow alternative input methods in the Accessibility Keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityKeyboard:"), value)
}/* debug [instance_properties/setter]: allowsAccessibilityKeyboard */


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityLiveCaptions
func (a_ AEAssessmentConfiguration) AllowsAccessibilityLiveCaptions() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityLiveCaptions"))
	return rv
}/* debug [instance_properties/getter]: allowsAccessibilityLiveCaptions */


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityLiveCaptions
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityLiveCaptions(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityLiveCaptions:"), value)
}/* debug [instance_properties/setter]: allowsAccessibilityLiveCaptions */


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityReader
func (a_ AEAssessmentConfiguration) AllowsAccessibilityReader() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityReader"))
	return rv
}/* debug [instance_properties/getter]: allowsAccessibilityReader */


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityReader
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityReader(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityReader:"), value)
}/* debug [instance_properties/setter]: allowsAccessibilityReader */


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) AllowsKeyboardShortcuts() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsKeyboardShortcuts"))
	return rv
}/* debug [instance_properties/getter]: allowsKeyboardShortcuts */


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) SetAllowsKeyboardShortcuts(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsKeyboardShortcuts:"), value)
}/* debug [instance_properties/setter]: allowsKeyboardShortcuts */


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) AllowsPredictiveKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPredictiveKeyboard"))
	return rv
}/* debug [instance_properties/getter]: allowsPredictiveKeyboard */


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsPredictiveKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPredictiveKeyboard:"), value)
}/* debug [instance_properties/setter]: allowsPredictiveKeyboard */


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) AllowsScreenshots() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsScreenshots"))
	return rv
}/* debug [instance_properties/getter]: allowsScreenshots */


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) SetAllowsScreenshots(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsScreenshots:"), value)
}/* debug [instance_properties/setter]: allowsScreenshots */


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) AllowsSpellCheck() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsSpellCheck"))
	return rv
}/* debug [instance_properties/getter]: allowsSpellCheck */


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) SetAllowsSpellCheck(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsSpellCheck:"), value)
}/* debug [instance_properties/setter]: allowsSpellCheck */


// A Boolean value that indicates whether to allow Autocorrect during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) AutocorrectMode() AEAutocorrectMode {
	rv := objc.Send[AEAutocorrectMode](a_.ID, objc.Sel("autocorrectMode"))
	return rv
}/* debug [instance_properties/getter]: autocorrectMode */


// A Boolean value that indicates whether to allow Autocorrect during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) SetAutocorrectMode(value AEAutocorrectMode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutocorrectMode:"), value)
}/* debug [instance_properties/setter]: autocorrectMode */


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/configurationsByApplication
func (a_ AEAssessmentConfiguration) ConfigurationsByApplication() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}/* debug [instance_properties/getter]: configurationsByApplication */


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/mainParticipantConfiguration
func (a_ AEAssessmentConfiguration) MainParticipantConfiguration() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}/* debug [instance_properties/getter]: mainParticipantConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AEAssessmentConfiguration */


