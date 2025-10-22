// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASESamplerNodeDefinition

// ExampleNewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition demonstrates how to create a PHASESamplerNodeDefinition instance using NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition.
// Creates a sampler node with the given sound asset and mixer.
func ExampleNewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition() {
	_ = phase.NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition(
		"soundAssetIdentifier", // soundAssetIdentifier string
		phase.PHASEMixerDefinition{}, // mixerDefinition PHASEMixerDefinition
	)
	// Output:
}
// ExampleNewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier demonstrates how to create a PHASESamplerNodeDefinition instance using NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier.
// Creates a named sampler node with the given sound asset and mixer.
func ExampleNewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier() {
	_ = phase.NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier(
		"soundAssetIdentifier", // soundAssetIdentifier string
		phase.PHASEMixerDefinition{}, // mixerDefinition PHASEMixerDefinition
		"identifier", // identifier string
	)
	// Output:
}
