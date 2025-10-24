// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioEngine

// ExampleNewAudioEngine demonstrates how to create a AudioEngine instance.
// Creates an audio engine instance for rendering in real time.
func ExampleNewAudioEngine() {
	_ = avfaudio.NewAudioEngine()
	// Output:
}
// ExampleAudioEngine_DisableManualRenderingMode demonstrates using DisableManualRenderingMode on a AudioEngine instance.
// Sets the engine to render to or from an audio device.
func ExampleAudioEngine_DisableManualRenderingMode() {
	obj := avfaudio.NewAudioEngine()
	obj.DisableManualRenderingMode()
	// Output:
	}

// ExampleAudioEngine_Pause demonstrates using Pause on a AudioEngine instance.
// Pauses the audio engine.
func ExampleAudioEngine_Pause() {
	obj := avfaudio.NewAudioEngine()
	obj.Pause()
	// Output:
	}

// ExampleAudioEngine_Prepare demonstrates using Prepare on a AudioEngine instance.
// Prepares the audio engine for starting.
//
// Note: This example is not executed because Prepare crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleAudioEngine_Prepare() {
	obj := avfaudio.NewAudioEngine()
	obj.Prepare()
	}

// ExampleAudioEngine_Reset demonstrates using Reset on a AudioEngine instance.
// Resets all audio nodes in the audio engine.
func ExampleAudioEngine_Reset() {
	obj := avfaudio.NewAudioEngine()
	obj.Reset()
	// Output:
	}

// ExampleAudioEngine_Stop demonstrates using Stop on a AudioEngine instance.
// Stops the audio engine and releases any previously prepared resources.
func ExampleAudioEngine_Stop() {
	obj := avfaudio.NewAudioEngine()
	obj.Stop()
	// Output:
	}

