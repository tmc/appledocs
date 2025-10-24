// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox_test

import (
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// Suppress unused import errors
var _ = audiotoolbox.NewAudioUnit

// ExampleAudioUnit_DeallocateRenderResources demonstrates using DeallocateRenderResources on a AudioUnit instance.
// Deallocates resources required to render audio.
func ExampleAudioUnit_DeallocateRenderResources() {
	obj := audiotoolbox.NewAudioUnit()
	obj.DeallocateRenderResources()
	// Output:
	}

// ExampleAudioUnit_Reset demonstrates using Reset on a AudioUnit instance.
// Resets transitory rendering state to its initial state.
func ExampleAudioUnit_Reset() {
	obj := audiotoolbox.NewAudioUnit()
	obj.Reset()
	// Output:
	}

// ExampleAudioUnit_StopHardware demonstrates using StopHardware on a AudioUnit instance.
// Stops the audio hardware.
func ExampleAudioUnit_StopHardware() {
	obj := audiotoolbox.NewAudioUnit()
	obj.StopHardware()
	// Output:
	}

