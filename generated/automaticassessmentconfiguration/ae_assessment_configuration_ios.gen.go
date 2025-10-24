//go:build darwin && ios

// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AEAssessmentConfiguration


// iOS-only properties

// A Boolean value that indicates whether to allow the speech-related accessibility features during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilitySpeech
func (a_ AEAssessmentConfiguration) AllowsAccessibilitySpeech() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilitySpeech"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilitySpeech(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsAccessibilitySpeech:"), value)
}

// A Boolean value that indicates whether to allow accessibility typing feedback during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsAccessibilityTypingFeedback
func (a_ AEAssessmentConfiguration) AllowsAccessibilityTypingFeedback() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsAccessibilityTypingFeedback"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsAccessibilityTypingFeedback(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsAccessibilityTypingFeedback:"), value)
}

// A Boolean value that indicates whether to allow Handoff during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsActivityContinuation
func (a_ AEAssessmentConfiguration) AllowsActivityContinuation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsActivityContinuation"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsActivityContinuation(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsActivityContinuation:"), value)
}

// A Boolean value that indicates whether to allow Slide to Type to operate during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsContinuousPathKeyboard
func (a_ AEAssessmentConfiguration) AllowsContinuousPathKeyboard() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsContinuousPathKeyboard"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsContinuousPathKeyboard(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsContinuousPathKeyboard:"), value)
}

// A Boolean value that indicates whether to allow the use of dictation during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsDictation
func (a_ AEAssessmentConfiguration) AllowsDictation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsDictation"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsDictation(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsDictation:"), value)
}

// A Boolean value that indicates whether to allow password autofill during an assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/allowsPasswordAutoFill
func (a_ AEAssessmentConfiguration) AllowsPasswordAutoFill() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsPasswordAutoFill"))
	return rv
}
func (a_ AEAssessmentConfiguration) SetAllowsPasswordAutoFill(value bool) {
	a_.ID.Send(objc.RegisterName("setAllowsPasswordAutoFill:"), value)
}





