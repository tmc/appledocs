// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/avfaudio"
)

// PCXProviderDelegate is the CXProviderDelegate protocol interface.
//
// A collection of methods that a telephony provider object calls.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.callkit/documentation/CallKit/CXProviderDelegate
type PCXProviderDelegate interface {
	// Required methods
	ProviderDidReset(provider ICXProvider)/* debug [protocol_interface/required_method]: ProviderDidReset */
	// Optional methods
	ProviderDidActivateAudioSession(provider ICXProvider, audioSession avfaudio.AudioSession)
	HasProviderDidActivateAudioSession() bool
	ProviderDidDeactivateAudioSession(provider ICXProvider, audioSession avfaudio.AudioSession)
	HasProviderDidDeactivateAudioSession() bool
	ProviderExecuteTransaction(provider ICXProvider, transaction ICXTransaction) bool
	HasProviderExecuteTransaction() bool
	ProviderPerformStartCallAction(provider ICXProvider, action ICXStartCallAction)
	HasProviderPerformStartCallAction() bool
	ProviderPerformSetTranslatingCallAction(provider ICXProvider, action ICXSetTranslatingCallAction)
	HasProviderPerformSetTranslatingCallAction() bool
	ProviderPerformPlayDTMFCallAction(provider ICXProvider, action ICXPlayDTMFCallAction)
	HasProviderPerformPlayDTMFCallAction() bool
	ProviderPerformSetMutedCallAction(provider ICXProvider, action ICXSetMutedCallAction)
	HasProviderPerformSetMutedCallAction() bool
	ProviderPerformSetHeldCallAction(provider ICXProvider, action ICXSetHeldCallAction)
	HasProviderPerformSetHeldCallAction() bool
	ProviderPerformEndCallAction(provider ICXProvider, action ICXEndCallAction)
	HasProviderPerformEndCallAction() bool
	ProviderPerformSetGroupCallAction(provider ICXProvider, action ICXSetGroupCallAction)
	HasProviderPerformSetGroupCallAction() bool
	ProviderPerformAnswerCallAction(provider ICXProvider, action ICXAnswerCallAction)
	HasProviderPerformAnswerCallAction() bool
	ProviderTimedOutPerformingAction(provider ICXProvider, action ICXAction)
	HasProviderTimedOutPerformingAction() bool
	ProviderDidBegin(provider ICXProvider)
	HasProviderDidBegin() bool
}

// CXProviderDelegate is a delegate implementation builder for the PCXProviderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CXProviderDelegate struct {
	_ProviderDidActivateAudioSession func(provider ICXProvider, audioSession avfaudio.AudioSession)
	_ProviderDidDeactivateAudioSession func(provider ICXProvider, audioSession avfaudio.AudioSession)
	_ProviderExecuteTransaction func(provider ICXProvider, transaction ICXTransaction) bool
	_ProviderPerformStartCallAction func(provider ICXProvider, action ICXStartCallAction)
	_ProviderPerformSetTranslatingCallAction func(provider ICXProvider, action ICXSetTranslatingCallAction)
	_ProviderPerformPlayDTMFCallAction func(provider ICXProvider, action ICXPlayDTMFCallAction)
	_ProviderPerformSetMutedCallAction func(provider ICXProvider, action ICXSetMutedCallAction)
	_ProviderPerformSetHeldCallAction func(provider ICXProvider, action ICXSetHeldCallAction)
	_ProviderPerformEndCallAction func(provider ICXProvider, action ICXEndCallAction)
	_ProviderPerformSetGroupCallAction func(provider ICXProvider, action ICXSetGroupCallAction)
	_ProviderPerformAnswerCallAction func(provider ICXProvider, action ICXAnswerCallAction)
	_ProviderTimedOutPerformingAction func(provider ICXProvider, action ICXAction)
	_ProviderDidBegin func(provider ICXProvider)
	_ProviderDidReset func(provider ICXProvider)
}

// SetProviderDidActivateAudioSession sets the handler for the ProviderDidActivateAudioSession delegate method.
//
// Called when the provider’s audio session is activated.
func (d *CXProviderDelegate) SetProviderDidActivateAudioSession(f func(provider ICXProvider, audioSession avfaudio.AudioSession)) {
	d._ProviderDidActivateAudioSession = f
}

// SetProviderDidDeactivateAudioSession sets the handler for the ProviderDidDeactivateAudioSession delegate method.
//
// Called when the provider’s audio session is deactivated.
func (d *CXProviderDelegate) SetProviderDidDeactivateAudioSession(f func(provider ICXProvider, audioSession avfaudio.AudioSession)) {
	d._ProviderDidDeactivateAudioSession = f
}

// SetProviderExecuteTransaction sets the handler for the ProviderExecuteTransaction delegate method.
//
// Called when a transaction is executed by a call controller.
func (d *CXProviderDelegate) SetProviderExecuteTransaction(f func(provider ICXProvider, transaction ICXTransaction) bool) {
	d._ProviderExecuteTransaction = f
}

// SetProviderPerformStartCallAction sets the handler for the ProviderPerformStartCallAction delegate method.
//
// Called when the provider performs the specified start call action.
func (d *CXProviderDelegate) SetProviderPerformStartCallAction(f func(provider ICXProvider, action ICXStartCallAction)) {
	d._ProviderPerformStartCallAction = f
}

// SetProviderPerformSetTranslatingCallAction sets the handler for the ProviderPerformSetTranslatingCallAction delegate method.
//
// Called when the provider performs the specified set translation action.
func (d *CXProviderDelegate) SetProviderPerformSetTranslatingCallAction(f func(provider ICXProvider, action ICXSetTranslatingCallAction)) {
	d._ProviderPerformSetTranslatingCallAction = f
}

// SetProviderPerformPlayDTMFCallAction sets the handler for the ProviderPerformPlayDTMFCallAction delegate method.
//
// Called when the provider performs the specified play DTMF (dual tone multifrequency) call action.
func (d *CXProviderDelegate) SetProviderPerformPlayDTMFCallAction(f func(provider ICXProvider, action ICXPlayDTMFCallAction)) {
	d._ProviderPerformPlayDTMFCallAction = f
}

// SetProviderPerformSetMutedCallAction sets the handler for the ProviderPerformSetMutedCallAction delegate method.
//
// Called when the provider performs the specified set muted call action.
func (d *CXProviderDelegate) SetProviderPerformSetMutedCallAction(f func(provider ICXProvider, action ICXSetMutedCallAction)) {
	d._ProviderPerformSetMutedCallAction = f
}

// SetProviderPerformSetHeldCallAction sets the handler for the ProviderPerformSetHeldCallAction delegate method.
//
// Called when the provider performs the specified set held call action.
func (d *CXProviderDelegate) SetProviderPerformSetHeldCallAction(f func(provider ICXProvider, action ICXSetHeldCallAction)) {
	d._ProviderPerformSetHeldCallAction = f
}

// SetProviderPerformEndCallAction sets the handler for the ProviderPerformEndCallAction delegate method.
//
// Called when the provider performs the specified end call action.
func (d *CXProviderDelegate) SetProviderPerformEndCallAction(f func(provider ICXProvider, action ICXEndCallAction)) {
	d._ProviderPerformEndCallAction = f
}

// SetProviderPerformSetGroupCallAction sets the handler for the ProviderPerformSetGroupCallAction delegate method.
//
// Called when the provider performs the specified set group call action.
func (d *CXProviderDelegate) SetProviderPerformSetGroupCallAction(f func(provider ICXProvider, action ICXSetGroupCallAction)) {
	d._ProviderPerformSetGroupCallAction = f
}

// SetProviderPerformAnswerCallAction sets the handler for the ProviderPerformAnswerCallAction delegate method.
//
// Called when the provider performs the specified answer call action.
func (d *CXProviderDelegate) SetProviderPerformAnswerCallAction(f func(provider ICXProvider, action ICXAnswerCallAction)) {
	d._ProviderPerformAnswerCallAction = f
}

// SetProviderTimedOutPerformingAction sets the handler for the ProviderTimedOutPerformingAction delegate method.
//
// Called when the provider performs the specified action times out.
func (d *CXProviderDelegate) SetProviderTimedOutPerformingAction(f func(provider ICXProvider, action ICXAction)) {
	d._ProviderTimedOutPerformingAction = f
}

// SetProviderDidBegin sets the handler for the ProviderDidBegin delegate method.
//
// Called when the provider begins.
func (d *CXProviderDelegate) SetProviderDidBegin(f func(provider ICXProvider)) {
	d._ProviderDidBegin = f
}

// SetProviderDidReset sets the handler for the ProviderDidReset delegate method.
//
// Called when the provider is reset.
func (d *CXProviderDelegate) SetProviderDidReset(f func(provider ICXProvider)) {
	d._ProviderDidReset = f
}

// ProviderDidActivateAudioSession implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderDidActivateAudioSession(provider ICXProvider, audioSession avfaudio.AudioSession) {
	if d._ProviderDidActivateAudioSession != nil {
		d._ProviderDidActivateAudioSession(provider, audioSession)
	}
}

// HasProviderDidActivateAudioSession returns true if a handler for ProviderDidActivateAudioSession has been set.
func (d *CXProviderDelegate) HasProviderDidActivateAudioSession() bool {
	return d._ProviderDidActivateAudioSession != nil
}

// ProviderDidDeactivateAudioSession implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderDidDeactivateAudioSession(provider ICXProvider, audioSession avfaudio.AudioSession) {
	if d._ProviderDidDeactivateAudioSession != nil {
		d._ProviderDidDeactivateAudioSession(provider, audioSession)
	}
}

// HasProviderDidDeactivateAudioSession returns true if a handler for ProviderDidDeactivateAudioSession has been set.
func (d *CXProviderDelegate) HasProviderDidDeactivateAudioSession() bool {
	return d._ProviderDidDeactivateAudioSession != nil
}

// ProviderExecuteTransaction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderExecuteTransaction(provider ICXProvider, transaction ICXTransaction) bool {
	if d._ProviderExecuteTransaction != nil {
		return d._ProviderExecuteTransaction(provider, transaction)
	}
	var zero bool
	return zero
}

// HasProviderExecuteTransaction returns true if a handler for ProviderExecuteTransaction has been set.
func (d *CXProviderDelegate) HasProviderExecuteTransaction() bool {
	return d._ProviderExecuteTransaction != nil
}

// ProviderPerformStartCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformStartCallAction(provider ICXProvider, action ICXStartCallAction) {
	if d._ProviderPerformStartCallAction != nil {
		d._ProviderPerformStartCallAction(provider, action)
	}
}

// HasProviderPerformStartCallAction returns true if a handler for ProviderPerformStartCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformStartCallAction() bool {
	return d._ProviderPerformStartCallAction != nil
}

// ProviderPerformSetTranslatingCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformSetTranslatingCallAction(provider ICXProvider, action ICXSetTranslatingCallAction) {
	if d._ProviderPerformSetTranslatingCallAction != nil {
		d._ProviderPerformSetTranslatingCallAction(provider, action)
	}
}

// HasProviderPerformSetTranslatingCallAction returns true if a handler for ProviderPerformSetTranslatingCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformSetTranslatingCallAction() bool {
	return d._ProviderPerformSetTranslatingCallAction != nil
}

// ProviderPerformPlayDTMFCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformPlayDTMFCallAction(provider ICXProvider, action ICXPlayDTMFCallAction) {
	if d._ProviderPerformPlayDTMFCallAction != nil {
		d._ProviderPerformPlayDTMFCallAction(provider, action)
	}
}

// HasProviderPerformPlayDTMFCallAction returns true if a handler for ProviderPerformPlayDTMFCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformPlayDTMFCallAction() bool {
	return d._ProviderPerformPlayDTMFCallAction != nil
}

// ProviderPerformSetMutedCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformSetMutedCallAction(provider ICXProvider, action ICXSetMutedCallAction) {
	if d._ProviderPerformSetMutedCallAction != nil {
		d._ProviderPerformSetMutedCallAction(provider, action)
	}
}

// HasProviderPerformSetMutedCallAction returns true if a handler for ProviderPerformSetMutedCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformSetMutedCallAction() bool {
	return d._ProviderPerformSetMutedCallAction != nil
}

// ProviderPerformSetHeldCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformSetHeldCallAction(provider ICXProvider, action ICXSetHeldCallAction) {
	if d._ProviderPerformSetHeldCallAction != nil {
		d._ProviderPerformSetHeldCallAction(provider, action)
	}
}

// HasProviderPerformSetHeldCallAction returns true if a handler for ProviderPerformSetHeldCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformSetHeldCallAction() bool {
	return d._ProviderPerformSetHeldCallAction != nil
}

// ProviderPerformEndCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformEndCallAction(provider ICXProvider, action ICXEndCallAction) {
	if d._ProviderPerformEndCallAction != nil {
		d._ProviderPerformEndCallAction(provider, action)
	}
}

// HasProviderPerformEndCallAction returns true if a handler for ProviderPerformEndCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformEndCallAction() bool {
	return d._ProviderPerformEndCallAction != nil
}

// ProviderPerformSetGroupCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformSetGroupCallAction(provider ICXProvider, action ICXSetGroupCallAction) {
	if d._ProviderPerformSetGroupCallAction != nil {
		d._ProviderPerformSetGroupCallAction(provider, action)
	}
}

// HasProviderPerformSetGroupCallAction returns true if a handler for ProviderPerformSetGroupCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformSetGroupCallAction() bool {
	return d._ProviderPerformSetGroupCallAction != nil
}

// ProviderPerformAnswerCallAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderPerformAnswerCallAction(provider ICXProvider, action ICXAnswerCallAction) {
	if d._ProviderPerformAnswerCallAction != nil {
		d._ProviderPerformAnswerCallAction(provider, action)
	}
}

// HasProviderPerformAnswerCallAction returns true if a handler for ProviderPerformAnswerCallAction has been set.
func (d *CXProviderDelegate) HasProviderPerformAnswerCallAction() bool {
	return d._ProviderPerformAnswerCallAction != nil
}

// ProviderTimedOutPerformingAction implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderTimedOutPerformingAction(provider ICXProvider, action ICXAction) {
	if d._ProviderTimedOutPerformingAction != nil {
		d._ProviderTimedOutPerformingAction(provider, action)
	}
}

// HasProviderTimedOutPerformingAction returns true if a handler for ProviderTimedOutPerformingAction has been set.
func (d *CXProviderDelegate) HasProviderTimedOutPerformingAction() bool {
	return d._ProviderTimedOutPerformingAction != nil
}

// ProviderDidBegin implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderDidBegin(provider ICXProvider) {
	if d._ProviderDidBegin != nil {
		d._ProviderDidBegin(provider)
	}
}

// HasProviderDidBegin returns true if a handler for ProviderDidBegin has been set.
func (d *CXProviderDelegate) HasProviderDidBegin() bool {
	return d._ProviderDidBegin != nil
}

// ProviderDidReset implements the PCXProviderDelegate interface.
func (d *CXProviderDelegate) ProviderDidReset(provider ICXProvider) {
	if d._ProviderDidReset != nil {
		d._ProviderDidReset(provider)
	}
}

// HasProviderDidReset returns true if a handler for ProviderDidReset has been set.
func (d *CXProviderDelegate) HasProviderDidReset() bool {
	return d._ProviderDidReset != nil
}
