// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AEAssessmentConfiguration] class.
type IAEAssessmentConfiguration interface {
	objectivec.IObject
	// properties:
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
	AllowsAccessibilityKeyboard() bool
	SetAllowsAccessibilityKeyboard(value bool)
	AllowsAccessibilityLiveCaptions() bool
	SetAllowsAccessibilityLiveCaptions(value bool)
	AllowsAccessibilityReader() bool
	SetAllowsAccessibilityReader(value bool)
	// methods:
	RemoveApplication(application IAEAssessmentApplication)
	SetConfigurationForApplication(configuration IAEAssessmentParticipantConfiguration, application IAEAssessmentApplication)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentConfigurationClass) Alloc() AEAssessmentConfiguration {
	rv := objc.Send[AEAssessmentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Removes the availability of a previously allowed app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/remove(_:)
func (a_ AEAssessmentConfiguration) RemoveApplication(application IAEAssessmentApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeApplication:"), application)
}


// Adds an app to the list of apps available during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/setConfiguration(_:for:)
func (a_ AEAssessmentConfiguration) SetConfigurationForApplication(configuration IAEAssessmentParticipantConfiguration, application IAEAssessmentApplication) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfiguration:forApplication:"), configuration, application)
}


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) AllowsKeyboardShortcuts() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsKeyboardShortcuts"))
	return rv
}


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) SetAllowsKeyboardShortcuts(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsKeyboardShortcuts:"), value)
}


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) AllowsPredictiveKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPredictiveKeyboard"))
	return rv
}


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsPredictiveKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPredictiveKeyboard:"), value)
}


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) AllowsScreenshots() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsScreenshots"))
	return rv
}


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) SetAllowsScreenshots(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsScreenshots:"), value)
}


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) AllowsSpellCheck() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsSpellCheck"))
	return rv
}


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) SetAllowsSpellCheck(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsSpellCheck:"), value)
}


// A Boolean value that indicates whether to allow Autocorrect during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) AutocorrectMode() AEAutocorrectMode {
	rv := objc.Send[AEAutocorrectMode](a_.ID, objc.Sel("autocorrectMode"))
	return rv
}


// A Boolean value that indicates whether to allow Autocorrect during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) SetAutocorrectMode(value AEAutocorrectMode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutocorrectMode:"), value)
}


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/configurationsByApplication
func (a_ AEAssessmentConfiguration) ConfigurationsByApplication() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/mainParticipantConfiguration
func (a_ AEAssessmentConfiguration) MainParticipantConfiguration() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}


// A Boolean value that indicates whether to allow alternative input methods in the Accessibility Keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitykeyboard
func (a_ AEAssessmentConfiguration) AllowsAccessibilityKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityKeyboard"))
	return rv
}


// A Boolean value that indicates whether to allow alternative input methods in the Accessibility Keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitykeyboard
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityKeyboard:"), value)
}


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitylivecaptions
func (a_ AEAssessmentConfiguration) AllowsAccessibilityLiveCaptions() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityLiveCaptions"))
	return rv
}


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitylivecaptions
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityLiveCaptions(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityLiveCaptions:"), value)
}


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilityreader
func (a_ AEAssessmentConfiguration) AllowsAccessibilityReader() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityReader"))
	return rv
}


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilityreader
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityReader(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityReader:"), value)
}


