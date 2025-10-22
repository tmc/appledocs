// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEGroupPresetSetting

// ExampleNewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType demonstrates how to create a PHASEGroupPresetSetting instance using NewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType.
// Creates a group preset setting.
func ExampleNewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType() {
	_ = phase.NewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType(
		0.0, // gain float64
		0.0, // rate float64
		phase.PHASECurveType{}, // gainCurveType PHASECurveType
		phase.PHASECurveType{}, // rateCurveType PHASECurveType
	)
	// Output:
}
