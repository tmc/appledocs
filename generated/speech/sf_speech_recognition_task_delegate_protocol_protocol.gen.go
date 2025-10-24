// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

// PSFSpeechRecognitionTaskDelegate is the SFSpeechRecognitionTaskDelegate protocol interface.
//
// A protocol with methods for managing multi-utterance speech recognition requests.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// See: doc://com.apple.speech/documentation/Speech/SFSpeechRecognitionTaskDelegate
type PSFSpeechRecognitionTaskDelegate interface {
	// Optional methods
	SpeechRecognitionDidDetectSpeech(task ISFSpeechRecognitionTask)
	HasSpeechRecognitionDidDetectSpeech() bool
	SpeechRecognitionTaskDidFinishRecognition(task ISFSpeechRecognitionTask, recognitionResult ISFSpeechRecognitionResult)
	HasSpeechRecognitionTaskDidFinishRecognition() bool
	SpeechRecognitionTaskDidFinishSuccessfully(task ISFSpeechRecognitionTask, successfully bool)
	HasSpeechRecognitionTaskDidFinishSuccessfully() bool
	SpeechRecognitionTaskDidHypothesizeTranscription(task ISFSpeechRecognitionTask, transcription ISFTranscription)
	HasSpeechRecognitionTaskDidHypothesizeTranscription() bool
	SpeechRecognitionTaskDidProcessAudioDuration(task ISFSpeechRecognitionTask, duration float64)
	HasSpeechRecognitionTaskDidProcessAudioDuration() bool
	SpeechRecognitionTaskFinishedReadingAudio(task ISFSpeechRecognitionTask)
	HasSpeechRecognitionTaskFinishedReadingAudio() bool
	SpeechRecognitionTaskWasCancelled(task ISFSpeechRecognitionTask)
	HasSpeechRecognitionTaskWasCancelled() bool
}

// SFSpeechRecognitionTaskDelegate is a delegate implementation builder for the PSFSpeechRecognitionTaskDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SFSpeechRecognitionTaskDelegate struct {
	_SpeechRecognitionDidDetectSpeech                 func(task ISFSpeechRecognitionTask)
	_SpeechRecognitionTaskDidFinishRecognition        func(task ISFSpeechRecognitionTask, recognitionResult ISFSpeechRecognitionResult)
	_SpeechRecognitionTaskDidFinishSuccessfully       func(task ISFSpeechRecognitionTask, successfully bool)
	_SpeechRecognitionTaskDidHypothesizeTranscription func(task ISFSpeechRecognitionTask, transcription ISFTranscription)
	_SpeechRecognitionTaskDidProcessAudioDuration     func(task ISFSpeechRecognitionTask, duration float64)
	_SpeechRecognitionTaskFinishedReadingAudio        func(task ISFSpeechRecognitionTask)
	_SpeechRecognitionTaskWasCancelled                func(task ISFSpeechRecognitionTask)
}

// SetSpeechRecognitionDidDetectSpeech sets the handler for the SpeechRecognitionDidDetectSpeech delegate method.
//
// Tells the delegate when the task first detects speech in the source audio.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionDidDetectSpeech(f func(task ISFSpeechRecognitionTask)) {
	d._SpeechRecognitionDidDetectSpeech = f
}

// SetSpeechRecognitionTaskDidFinishRecognition sets the handler for the SpeechRecognitionTaskDidFinishRecognition delegate method.
//
// Tells the delegate when the final utterance is recognized.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskDidFinishRecognition(f func(task ISFSpeechRecognitionTask, recognitionResult ISFSpeechRecognitionResult)) {
	d._SpeechRecognitionTaskDidFinishRecognition = f
}

// SetSpeechRecognitionTaskDidFinishSuccessfully sets the handler for the SpeechRecognitionTaskDidFinishSuccessfully delegate method.
//
// Tells the delegate when the recognition of all requested utterances is finished.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskDidFinishSuccessfully(f func(task ISFSpeechRecognitionTask, successfully bool)) {
	d._SpeechRecognitionTaskDidFinishSuccessfully = f
}

// SetSpeechRecognitionTaskDidHypothesizeTranscription sets the handler for the SpeechRecognitionTaskDidHypothesizeTranscription delegate method.
//
// Tells the delegate that a hypothesized transcription is available.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskDidHypothesizeTranscription(f func(task ISFSpeechRecognitionTask, transcription ISFTranscription)) {
	d._SpeechRecognitionTaskDidHypothesizeTranscription = f
}

// SetSpeechRecognitionTaskDidProcessAudioDuration sets the handler for the SpeechRecognitionTaskDidProcessAudioDuration delegate method.
//
// Tells the delegate how much audio has been processed by the task.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskDidProcessAudioDuration(f func(task ISFSpeechRecognitionTask, duration float64)) {
	d._SpeechRecognitionTaskDidProcessAudioDuration = f
}

// SetSpeechRecognitionTaskFinishedReadingAudio sets the handler for the SpeechRecognitionTaskFinishedReadingAudio delegate method.
//
// Tells the delegate when the task is no longer accepting new audio input, even if final processing is in progress.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskFinishedReadingAudio(f func(task ISFSpeechRecognitionTask)) {
	d._SpeechRecognitionTaskFinishedReadingAudio = f
}

// SetSpeechRecognitionTaskWasCancelled sets the handler for the SpeechRecognitionTaskWasCancelled delegate method.
//
// Tells the delegate that the task has been canceled.
func (d *SFSpeechRecognitionTaskDelegate) SetSpeechRecognitionTaskWasCancelled(f func(task ISFSpeechRecognitionTask)) {
	d._SpeechRecognitionTaskWasCancelled = f
}

// SpeechRecognitionDidDetectSpeech implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionDidDetectSpeech(task ISFSpeechRecognitionTask) {
	if d._SpeechRecognitionDidDetectSpeech != nil {
		d._SpeechRecognitionDidDetectSpeech(task)
	}
}

// HasSpeechRecognitionDidDetectSpeech returns true if a handler for SpeechRecognitionDidDetectSpeech has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionDidDetectSpeech() bool {
	return d._SpeechRecognitionDidDetectSpeech != nil
}

// SpeechRecognitionTaskDidFinishRecognition implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskDidFinishRecognition(task ISFSpeechRecognitionTask, recognitionResult ISFSpeechRecognitionResult) {
	if d._SpeechRecognitionTaskDidFinishRecognition != nil {
		d._SpeechRecognitionTaskDidFinishRecognition(task, recognitionResult)
	}
}

// HasSpeechRecognitionTaskDidFinishRecognition returns true if a handler for SpeechRecognitionTaskDidFinishRecognition has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskDidFinishRecognition() bool {
	return d._SpeechRecognitionTaskDidFinishRecognition != nil
}

// SpeechRecognitionTaskDidFinishSuccessfully implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskDidFinishSuccessfully(task ISFSpeechRecognitionTask, successfully bool) {
	if d._SpeechRecognitionTaskDidFinishSuccessfully != nil {
		d._SpeechRecognitionTaskDidFinishSuccessfully(task, successfully)
	}
}

// HasSpeechRecognitionTaskDidFinishSuccessfully returns true if a handler for SpeechRecognitionTaskDidFinishSuccessfully has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskDidFinishSuccessfully() bool {
	return d._SpeechRecognitionTaskDidFinishSuccessfully != nil
}

// SpeechRecognitionTaskDidHypothesizeTranscription implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskDidHypothesizeTranscription(task ISFSpeechRecognitionTask, transcription ISFTranscription) {
	if d._SpeechRecognitionTaskDidHypothesizeTranscription != nil {
		d._SpeechRecognitionTaskDidHypothesizeTranscription(task, transcription)
	}
}

// HasSpeechRecognitionTaskDidHypothesizeTranscription returns true if a handler for SpeechRecognitionTaskDidHypothesizeTranscription has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskDidHypothesizeTranscription() bool {
	return d._SpeechRecognitionTaskDidHypothesizeTranscription != nil
}

// SpeechRecognitionTaskDidProcessAudioDuration implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskDidProcessAudioDuration(task ISFSpeechRecognitionTask, duration float64) {
	if d._SpeechRecognitionTaskDidProcessAudioDuration != nil {
		d._SpeechRecognitionTaskDidProcessAudioDuration(task, duration)
	}
}

// HasSpeechRecognitionTaskDidProcessAudioDuration returns true if a handler for SpeechRecognitionTaskDidProcessAudioDuration has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskDidProcessAudioDuration() bool {
	return d._SpeechRecognitionTaskDidProcessAudioDuration != nil
}

// SpeechRecognitionTaskFinishedReadingAudio implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskFinishedReadingAudio(task ISFSpeechRecognitionTask) {
	if d._SpeechRecognitionTaskFinishedReadingAudio != nil {
		d._SpeechRecognitionTaskFinishedReadingAudio(task)
	}
}

// HasSpeechRecognitionTaskFinishedReadingAudio returns true if a handler for SpeechRecognitionTaskFinishedReadingAudio has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskFinishedReadingAudio() bool {
	return d._SpeechRecognitionTaskFinishedReadingAudio != nil
}

// SpeechRecognitionTaskWasCancelled implements the PSFSpeechRecognitionTaskDelegate interface.
func (d *SFSpeechRecognitionTaskDelegate) SpeechRecognitionTaskWasCancelled(task ISFSpeechRecognitionTask) {
	if d._SpeechRecognitionTaskWasCancelled != nil {
		d._SpeechRecognitionTaskWasCancelled(task)
	}
}

// HasSpeechRecognitionTaskWasCancelled returns true if a handler for SpeechRecognitionTaskWasCancelled has been set.
func (d *SFSpeechRecognitionTaskDelegate) HasSpeechRecognitionTaskWasCancelled() bool {
	return d._SpeechRecognitionTaskWasCancelled != nil
}
