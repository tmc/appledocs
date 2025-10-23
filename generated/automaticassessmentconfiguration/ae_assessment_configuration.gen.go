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
	AllowsAccessibilitySpeech() bool /* primitive/slice/pointer. */
	SetAllowsAccessibilitySpeech(value bool /* primitive/slice/pointer. */)
	AllowsAccessibilityTypingFeedback() bool /* primitive/slice/pointer. */
	SetAllowsAccessibilityTypingFeedback(value bool /* primitive/slice/pointer. */)
	AllowsActivityContinuation() bool /* primitive/slice/pointer. */
	SetAllowsActivityContinuation(value bool /* primitive/slice/pointer. */)
	AllowsContinuousPathKeyboard() bool /* primitive/slice/pointer. */
	SetAllowsContinuousPathKeyboard(value bool /* primitive/slice/pointer. */)
	AllowsDictation() bool /* primitive/slice/pointer. */
	SetAllowsDictation(value bool /* primitive/slice/pointer. */)
	AllowsKeyboardShortcuts() bool /* primitive/slice/pointer. */
	SetAllowsKeyboardShortcuts(value bool /* primitive/slice/pointer. */)
	AllowsPasswordAutoFill() bool /* primitive/slice/pointer. */
	SetAllowsPasswordAutoFill(value bool /* primitive/slice/pointer. */)
	AllowsPredictiveKeyboard() bool /* primitive/slice/pointer. */
	SetAllowsPredictiveKeyboard(value bool /* primitive/slice/pointer. */)
	AllowsScreenshots() bool /* primitive/slice/pointer. */
	SetAllowsScreenshots(value bool /* primitive/slice/pointer. */)
	AllowsSpellCheck() bool /* primitive/slice/pointer. */
	SetAllowsSpellCheck(value bool /* primitive/slice/pointer. */)
	AutocorrectMode() AEAutocorrectMode
	SetAutocorrectMode(value AEAutocorrectMode)
	ConfigurationsByApplication() foundation.IDictionary /* already interface */
	MainParticipantConfiguration() IAEAssessmentParticipantConfiguration
	AllowsAccessibilityKeyboard() bool /* primitive/slice/pointer. */
	SetAllowsAccessibilityKeyboard(value bool /* primitive/slice/pointer. */)
	AllowsAccessibilityLiveCaptions() bool /* primitive/slice/pointer. */
	SetAllowsAccessibilityLiveCaptions(value bool /* primitive/slice/pointer. */)
	AllowsAccessibilityReader() bool /* primitive/slice/pointer. */
	SetAllowsAccessibilityReader(value bool /* primitive/slice/pointer. */)
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


// A Boolean value that indicates whether to allow the speech-related accessibility features during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilitySpeech
func (a_ AEAssessmentConfiguration) AllowsAccessibilitySpeech() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilitySpeech"))
	return rv
}


// A Boolean value that indicates whether to allow the speech-related accessibility features during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilitySpeech
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilitySpeech(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilitySpeech:"), value)
}


// A Boolean value that indicates whether to allow accessibility typing feedback during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityTypingFeedback
func (a_ AEAssessmentConfiguration) AllowsAccessibilityTypingFeedback() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityTypingFeedback"))
	return rv
}


// A Boolean value that indicates whether to allow accessibility typing feedback during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityTypingFeedback
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityTypingFeedback(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityTypingFeedback:"), value)
}


// A Boolean value that indicates whether to allow Handoff during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsActivityContinuation
func (a_ AEAssessmentConfiguration) AllowsActivityContinuation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsActivityContinuation"))
	return rv
}


// A Boolean value that indicates whether to allow Handoff during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsActivityContinuation
func (a_ AEAssessmentConfiguration) SetAllowsActivityContinuation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsActivityContinuation:"), value)
}


// A Boolean value that indicates whether to allow Slide to Type to operate during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsContinuousPathKeyboard
func (a_ AEAssessmentConfiguration) AllowsContinuousPathKeyboard() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsContinuousPathKeyboard"))
	return rv
}


// A Boolean value that indicates whether to allow Slide to Type to operate during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsContinuousPathKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsContinuousPathKeyboard(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsContinuousPathKeyboard:"), value)
}


// A Boolean value that indicates whether to allow the use of dictation during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsDictation
func (a_ AEAssessmentConfiguration) AllowsDictation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsDictation"))
	return rv
}


// A Boolean value that indicates whether to allow the use of dictation during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsDictation
func (a_ AEAssessmentConfiguration) SetAllowsDictation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsDictation:"), value)
}


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) AllowsKeyboardShortcuts() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsKeyboardShortcuts"))
	return rv
}


// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) SetAllowsKeyboardShortcuts(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsKeyboardShortcuts:"), value)
}


// A Boolean value that indicates whether to allow password autofill during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPasswordAutoFill
func (a_ AEAssessmentConfiguration) AllowsPasswordAutoFill() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPasswordAutoFill"))
	return rv
}


// A Boolean value that indicates whether to allow password autofill during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPasswordAutoFill
func (a_ AEAssessmentConfiguration) SetAllowsPasswordAutoFill(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPasswordAutoFill:"), value)
}


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) AllowsPredictiveKeyboard() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPredictiveKeyboard"))
	return rv
}


// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsPredictiveKeyboard(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPredictiveKeyboard:"), value)
}


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) AllowsScreenshots() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsScreenshots"))
	return rv
}


// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) SetAllowsScreenshots(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsScreenshots:"), value)
}


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) AllowsSpellCheck() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsSpellCheck"))
	return rv
}


// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) SetAllowsSpellCheck(value bool /* primitive/slice/pointer. */) {
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
func (a_ AEAssessmentConfiguration) ConfigurationsByApplication() foundation.IDictionary /* already interface */ {
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
func (a_ AEAssessmentConfiguration) AllowsAccessibilityKeyboard() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityKeyboard"))
	return rv
}


// A Boolean value that indicates whether to allow alternative input methods in the Accessibility Keyboard during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitykeyboard
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityKeyboard(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityKeyboard:"), value)
}


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitylivecaptions
func (a_ AEAssessmentConfiguration) AllowsAccessibilityLiveCaptions() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityLiveCaptions"))
	return rv
}


// A Boolean value that indicates whether to allow Live Captions during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilitylivecaptions
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityLiveCaptions(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityLiveCaptions:"), value)
}


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilityreader
func (a_ AEAssessmentConfiguration) AllowsAccessibilityReader() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityReader"))
	return rv
}


// A Boolean value that indicates whether to allow the Accessibility Reader during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/allowsaccessibilityreader
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityReader(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityReader:"), value)
}



