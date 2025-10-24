// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAEAssessmentSessionDelegate is the AEAssessmentSessionDelegate protocol interface.
//
// An interface that the session uses to provide information about session state changes to a delegate.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 13.4+
//   - iPadOS 13.4+
//   - macOS 10.15.4+
//
// See: doc://com.apple.automaticassessmentconfiguration/documentation/AutomaticAssessmentConfiguration/AEAssessmentSessionDelegate
type PAEAssessmentSessionDelegate interface {
	// Optional methods
	AssessmentSessionFailedToBeginWithError(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)
	HasAssessmentSessionFailedToBeginWithError() bool
	AssessmentSessionFailedToUpdateToConfigurationError(session IAEAssessmentSession, configuration IAEAssessmentConfiguration, error_ objc.IObject /* cross-framework: Error */)
	HasAssessmentSessionFailedToUpdateToConfigurationError() bool
	AssessmentSessionWasInterruptedWithError(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)
	HasAssessmentSessionWasInterruptedWithError() bool
	AssessmentSessionDidBegin(session IAEAssessmentSession)
	HasAssessmentSessionDidBegin() bool
	AssessmentSessionDidEnd(session IAEAssessmentSession)
	HasAssessmentSessionDidEnd() bool
	AssessmentSessionDidUpdate(session IAEAssessmentSession)
	HasAssessmentSessionDidUpdate() bool
}

// AEAssessmentSessionDelegate is a delegate implementation builder for the PAEAssessmentSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AEAssessmentSessionDelegate struct {
	_AssessmentSessionFailedToBeginWithError func(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)
	_AssessmentSessionFailedToUpdateToConfigurationError func(session IAEAssessmentSession, configuration IAEAssessmentConfiguration, error_ objc.IObject /* cross-framework: Error */)
	_AssessmentSessionWasInterruptedWithError func(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)
	_AssessmentSessionDidBegin func(session IAEAssessmentSession)
	_AssessmentSessionDidEnd func(session IAEAssessmentSession)
	_AssessmentSessionDidUpdate func(session IAEAssessmentSession)
}

// SetAssessmentSessionFailedToBeginWithError sets the handler for the AssessmentSessionFailedToBeginWithError delegate method.
//
// Tells the delegate that the session failed to start.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionFailedToBeginWithError(f func(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)) {
	d._AssessmentSessionFailedToBeginWithError = f
}

// SetAssessmentSessionFailedToUpdateToConfigurationError sets the handler for the AssessmentSessionFailedToUpdateToConfigurationError delegate method.
//
// Tells the delegate that a configuration update failed.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionFailedToUpdateToConfigurationError(f func(session IAEAssessmentSession, configuration IAEAssessmentConfiguration, error_ objc.IObject /* cross-framework: Error */)) {
	d._AssessmentSessionFailedToUpdateToConfigurationError = f
}

// SetAssessmentSessionWasInterruptedWithError sets the handler for the AssessmentSessionWasInterruptedWithError delegate method.
//
// Tells the delegate that a system failure interrupted the session.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionWasInterruptedWithError(f func(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */)) {
	d._AssessmentSessionWasInterruptedWithError = f
}

// SetAssessmentSessionDidBegin sets the handler for the AssessmentSessionDidBegin delegate method.
//
// Tells the delegate that the session started.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionDidBegin(f func(session IAEAssessmentSession)) {
	d._AssessmentSessionDidBegin = f
}

// SetAssessmentSessionDidEnd sets the handler for the AssessmentSessionDidEnd delegate method.
//
// Tells the delegate that the session ended.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionDidEnd(f func(session IAEAssessmentSession)) {
	d._AssessmentSessionDidEnd = f
}

// SetAssessmentSessionDidUpdate sets the handler for the AssessmentSessionDidUpdate delegate method.
//
// Tells the delegate that a configuration update succeeded.
func (d *AEAssessmentSessionDelegate) SetAssessmentSessionDidUpdate(f func(session IAEAssessmentSession)) {
	d._AssessmentSessionDidUpdate = f
}

// AssessmentSessionFailedToBeginWithError implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionFailedToBeginWithError(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */) {
	if d._AssessmentSessionFailedToBeginWithError != nil {
		d._AssessmentSessionFailedToBeginWithError(session, error_)
	}
}

// HasAssessmentSessionFailedToBeginWithError returns true if a handler for AssessmentSessionFailedToBeginWithError has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionFailedToBeginWithError() bool {
	return d._AssessmentSessionFailedToBeginWithError != nil
}

// AssessmentSessionFailedToUpdateToConfigurationError implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionFailedToUpdateToConfigurationError(session IAEAssessmentSession, configuration IAEAssessmentConfiguration, error_ objc.IObject /* cross-framework: Error */) {
	if d._AssessmentSessionFailedToUpdateToConfigurationError != nil {
		d._AssessmentSessionFailedToUpdateToConfigurationError(session, configuration, error_)
	}
}

// HasAssessmentSessionFailedToUpdateToConfigurationError returns true if a handler for AssessmentSessionFailedToUpdateToConfigurationError has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionFailedToUpdateToConfigurationError() bool {
	return d._AssessmentSessionFailedToUpdateToConfigurationError != nil
}

// AssessmentSessionWasInterruptedWithError implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionWasInterruptedWithError(session IAEAssessmentSession, error_ objc.IObject /* cross-framework: Error */) {
	if d._AssessmentSessionWasInterruptedWithError != nil {
		d._AssessmentSessionWasInterruptedWithError(session, error_)
	}
}

// HasAssessmentSessionWasInterruptedWithError returns true if a handler for AssessmentSessionWasInterruptedWithError has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionWasInterruptedWithError() bool {
	return d._AssessmentSessionWasInterruptedWithError != nil
}

// AssessmentSessionDidBegin implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionDidBegin(session IAEAssessmentSession) {
	if d._AssessmentSessionDidBegin != nil {
		d._AssessmentSessionDidBegin(session)
	}
}

// HasAssessmentSessionDidBegin returns true if a handler for AssessmentSessionDidBegin has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionDidBegin() bool {
	return d._AssessmentSessionDidBegin != nil
}

// AssessmentSessionDidEnd implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionDidEnd(session IAEAssessmentSession) {
	if d._AssessmentSessionDidEnd != nil {
		d._AssessmentSessionDidEnd(session)
	}
}

// HasAssessmentSessionDidEnd returns true if a handler for AssessmentSessionDidEnd has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionDidEnd() bool {
	return d._AssessmentSessionDidEnd != nil
}

// AssessmentSessionDidUpdate implements the PAEAssessmentSessionDelegate interface.
func (d *AEAssessmentSessionDelegate) AssessmentSessionDidUpdate(session IAEAssessmentSession) {
	if d._AssessmentSessionDidUpdate != nil {
		d._AssessmentSessionDidUpdate(session)
	}
}

// HasAssessmentSessionDidUpdate returns true if a handler for AssessmentSessionDidUpdate has been set.
func (d *AEAssessmentSessionDelegate) HasAssessmentSessionDidUpdate() bool {
	return d._AssessmentSessionDidUpdate != nil
}
