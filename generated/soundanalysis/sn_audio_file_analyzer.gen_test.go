// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis_test

import (
	"github.com/tmc/appledocs/generated/soundanalysis"
)

// Suppress unused import errors
var _ = soundanalysis.NewSNAudioFileAnalyzer

// ExampleSNAudioFileAnalyzer_Analyze demonstrates using Analyze on a SNAudioFileAnalyzer instance.
// Analyzes the audio file synchronously.
func ExampleSNAudioFileAnalyzer_Analyze() {
	obj := soundanalysis.NewSNAudioFileAnalyzer()
	obj.Analyze()
	// Output:
}

// ExampleSNAudioFileAnalyzer_CancelAnalysis demonstrates using CancelAnalysis on a SNAudioFileAnalyzer instance.
// Cancels all the asynchronous sound analysis requests the analyzer is currently processing.
func ExampleSNAudioFileAnalyzer_CancelAnalysis() {
	obj := soundanalysis.NewSNAudioFileAnalyzer()
	obj.CancelAnalysis()
	// Output:
}

// ExampleSNAudioFileAnalyzer_RemoveAllRequests demonstrates using RemoveAllRequests on a SNAudioFileAnalyzer instance.
// Removes all the sound analysis requests from the audio file analyzer.
func ExampleSNAudioFileAnalyzer_RemoveAllRequests() {
	obj := soundanalysis.NewSNAudioFileAnalyzer()
	obj.RemoveAllRequests()
	// Output:
}
