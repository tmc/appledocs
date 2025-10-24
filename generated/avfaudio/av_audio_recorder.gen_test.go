// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioRecorder

// ExampleAudioRecorder_DeleteRecording demonstrates using DeleteRecording on a AudioRecorder instance.
// Deletes a recorded audio file.
func ExampleAudioRecorder_DeleteRecording() {
	obj := avfaudio.NewAudioRecorder()
	_ = obj.DeleteRecording()
	// Output:
	}

// ExampleAudioRecorder_Pause demonstrates using Pause on a AudioRecorder instance.
// Pauses an audio recording.
func ExampleAudioRecorder_Pause() {
	obj := avfaudio.NewAudioRecorder()
	obj.Pause()
	// Output:
	}

// ExampleAudioRecorder_PrepareToRecord demonstrates using PrepareToRecord on a AudioRecorder instance.
// Creates an audio file and prepares the system for recording.
//
// Note: This example is not executed because PrepareToRecord crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleAudioRecorder_PrepareToRecord() {
	obj := avfaudio.NewAudioRecorder()
	_ = obj.PrepareToRecord()
	}

// ExampleAudioRecorder_Record demonstrates using Record on a AudioRecorder instance.
// Starts or resumes audio recording.
func ExampleAudioRecorder_Record() {
	obj := avfaudio.NewAudioRecorder()
	_ = obj.Record()
	// Output:
	}

// ExampleAudioRecorder_Stop demonstrates using Stop on a AudioRecorder instance.
// Stops recording and closes the audio file.
func ExampleAudioRecorder_Stop() {
	obj := avfaudio.NewAudioRecorder()
	obj.Stop()
	// Output:
	}

// ExampleAudioRecorder_UpdateMeters demonstrates using UpdateMeters on a AudioRecorder instance.
// Refreshes the average and peak power values for all channels of an audio recorder.
func ExampleAudioRecorder_UpdateMeters() {
	obj := avfaudio.NewAudioRecorder()
	obj.UpdateMeters()
	// Output:
	}

