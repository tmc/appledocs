// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox_test

import (
	"github.com/tmc/appledocs/generated/audiotoolbox"
)

// Suppress unused import errors
var _ = audiotoolbox.NewAudioUnitBusArray

// ExampleNewAudioUnitBusArrayWithAudioUnitBusType demonstrates how to create a AudioUnitBusArray instance using NewAudioUnitBusArrayWithAudioUnitBusType.
// Initializes an empty bus array.
func ExampleNewAudioUnitBusArrayWithAudioUnitBusType() {
	_ = audiotoolbox.NewAudioUnitBusArrayWithAudioUnitBusType(
		audiotoolbox.AUAudioUnit{},      // owner AUAudioUnit
		audiotoolbox.AudioUnitBusType{}, // busType AudioUnitBusType
	)
	// Output:
}

// ExampleNewAudioUnitBusArrayWithAudioUnitBusTypeBusses demonstrates how to create a AudioUnitBusArray instance using NewAudioUnitBusArrayWithAudioUnitBusTypeBusses.
// Initializes a bus array by making a copy of the supplied busses.
func ExampleNewAudioUnitBusArrayWithAudioUnitBusTypeBusses() {
	_ = audiotoolbox.NewAudioUnitBusArrayWithAudioUnitBusTypeBusses(
		audiotoolbox.AUAudioUnit{},      // owner AUAudioUnit
		audiotoolbox.AudioUnitBusType{}, // busType AudioUnitBusType
		[]audiotoolbox.AudioUnitBus{},   // busArray []AudioUnitBus
	)
	// Output:
}
