// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	RemoveApplication(application unsafe.Pointer)
	SetConfigurationForApplication(configuration unsafe.Pointer, application unsafe.Pointer)
}

// Configuration information for an assessment session.
//
// Create a configuration instance and pass it to the initializer of an instance to create a new assessment session. Before using the configuration, indicate which exceptions you want to allow for the assessment session’s restrictions by setting values on the configuration instance. For example, you can set values to allow dictation and certain aspects of autocorrect: While you provide a configuration instance when creating a session on iOS, iPadOS, and macOS, specific exceptions apply only to certain platforms. In particular, on macOS, you can selectively make specific apps besides your own available during an assessment — for example, to allow users to access a calculator or a dictionary. All other exceptions apply only to iOS and iPadOS.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/remove(_:)
func (a_ AEAssessmentConfiguration) RemoveApplication(application unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeApplication:"), application)
}

// Adds an app to the list of apps available during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/setConfiguration(_:for:)
func (a_ AEAssessmentConfiguration) SetConfigurationForApplication(configuration unsafe.Pointer, application unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfiguration:forApplication:"), configuration, application)
}

// A Boolean value that indicates whether to allow the speech-related accessibility features during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilitySpeech
func (a_ AEAssessmentConfiguration) AllowsAccessibilitySpeech() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilitySpeech"))
	return rv
}


// SetAllowsAccessibilitySpeech sets the value of the allowsAccessibilitySpeech property.
// A Boolean value that indicates whether to allow the speech-related accessibility features during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilitySpeech
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilitySpeech(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilitySpeech:"), value)
}
// A Boolean value that indicates whether to allow accessibility typing feedback during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityTypingFeedback
func (a_ AEAssessmentConfiguration) AllowsAccessibilityTypingFeedback() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityTypingFeedback"))
	return rv
}


// SetAllowsAccessibilityTypingFeedback sets the value of the allowsAccessibilityTypingFeedback property.
// A Boolean value that indicates whether to allow accessibility typing feedback during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityTypingFeedback
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityTypingFeedback(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsAccessibilityTypingFeedback:"), value)
}
// A Boolean value that indicates whether to allow Handoff during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsActivityContinuation
func (a_ AEAssessmentConfiguration) AllowsActivityContinuation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsActivityContinuation"))
	return rv
}


// SetAllowsActivityContinuation sets the value of the allowsActivityContinuation property.
// A Boolean value that indicates whether to allow Handoff during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsActivityContinuation
func (a_ AEAssessmentConfiguration) SetAllowsActivityContinuation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsActivityContinuation:"), value)
}
// A Boolean value that indicates whether to allow Slide to Type to operate during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsContinuousPathKeyboard
func (a_ AEAssessmentConfiguration) AllowsContinuousPathKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsContinuousPathKeyboard"))
	return rv
}


// SetAllowsContinuousPathKeyboard sets the value of the allowsContinuousPathKeyboard property.
// A Boolean value that indicates whether to allow Slide to Type to operate during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsContinuousPathKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsContinuousPathKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsContinuousPathKeyboard:"), value)
}
// A Boolean value that indicates whether to allow the use of dictation during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsDictation
func (a_ AEAssessmentConfiguration) AllowsDictation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsDictation"))
	return rv
}


// SetAllowsDictation sets the value of the allowsDictation property.
// A Boolean value that indicates whether to allow the use of dictation during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsDictation
func (a_ AEAssessmentConfiguration) SetAllowsDictation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsDictation:"), value)
}
// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) AllowsKeyboardShortcuts() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsKeyboardShortcuts"))
	return rv
}


// SetAllowsKeyboardShortcuts sets the value of the allowsKeyboardShortcuts property.
// A Boolean value that indicates whether to allow keyboard shortcuts during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsKeyboardShortcuts
func (a_ AEAssessmentConfiguration) SetAllowsKeyboardShortcuts(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsKeyboardShortcuts:"), value)
}
// A Boolean value that indicates whether to allow password autofill during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPasswordAutoFill
func (a_ AEAssessmentConfiguration) AllowsPasswordAutoFill() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPasswordAutoFill"))
	return rv
}


// SetAllowsPasswordAutoFill sets the value of the allowsPasswordAutoFill property.
// A Boolean value that indicates whether to allow password autofill during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPasswordAutoFill
func (a_ AEAssessmentConfiguration) SetAllowsPasswordAutoFill(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPasswordAutoFill:"), value)
}
// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) AllowsPredictiveKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPredictiveKeyboard"))
	return rv
}


// SetAllowsPredictiveKeyboard sets the value of the allowsPredictiveKeyboard property.
// A Boolean value that indicates whether to enable the predictive keyboard during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPredictiveKeyboard
func (a_ AEAssessmentConfiguration) SetAllowsPredictiveKeyboard(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsPredictiveKeyboard:"), value)
}
// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) AllowsScreenshots() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsScreenshots"))
	return rv
}


// SetAllowsScreenshots sets the value of the allowsScreenshots property.
// A Boolean value that indicates whether to allow screenshots copied to the clipboard during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsScreenshots
func (a_ AEAssessmentConfiguration) SetAllowsScreenshots(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsScreenshots:"), value)
}
// A Boolean value that indicates whether to allow spell check during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) AllowsSpellCheck() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsSpellCheck"))
	return rv
}


// SetAllowsSpellCheck sets the value of the allowsSpellCheck property.
// A Boolean value that indicates whether to allow spell check during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsSpellCheck
func (a_ AEAssessmentConfiguration) SetAllowsSpellCheck(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsSpellCheck:"), value)
}
// A Boolean value that indicates whether to allow Autocorrect during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) AutocorrectMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("autocorrectMode"))
	return rv
}


// SetAutocorrectMode sets the value of the autocorrectMode property.
// A Boolean value that indicates whether to allow Autocorrect during an assessment.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/autocorrectMode-swift.property
func (a_ AEAssessmentConfiguration) SetAutocorrectMode(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutocorrectMode:"), value)
}
// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/configurationsByApplication
func (a_ AEAssessmentConfiguration) ConfigurationsByApplication() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}

// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/mainParticipantConfiguration
func (a_ AEAssessmentConfiguration) MainParticipantConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}



