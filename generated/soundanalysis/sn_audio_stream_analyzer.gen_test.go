// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis_test

import (
	"github.com/tmc/appledocs/generated/soundanalysis"
)

// Suppress unused import errors
var _ = soundanalysis.NewSNAudioStreamAnalyzer

// ExampleSNAudioStreamAnalyzer_CompleteAnalysis demonstrates using CompleteAnalysis on a SNAudioStreamAnalyzer instance.
// Notifies the analyzer when it receives the final audio buffer.
func ExampleSNAudioStreamAnalyzer_CompleteAnalysis() {
	obj := soundanalysis.NewSNAudioStreamAnalyzer()
	obj.CompleteAnalysis()
	// Output:
	}

// ExampleSNAudioStreamAnalyzer_RemoveAllRequests demonstrates using RemoveAllRequests on a SNAudioStreamAnalyzer instance.
// Removes all the sound analysis requests from the audio stream analyzer.
func ExampleSNAudioStreamAnalyzer_RemoveAllRequests() {
	obj := soundanalysis.NewSNAudioStreamAnalyzer()
	obj.RemoveAllRequests()
	// Output:
	}

