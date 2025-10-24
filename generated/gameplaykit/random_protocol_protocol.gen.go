// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

// PRandom is the GKRandom protocol interface.
//
// The common interface for all randomization classes in (or usable with) GameplayKit.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gameplaykit/documentation/GameplayKit/GKRandom
type PRandom interface {
	// Required methods
	NextBool() bool/* debug [protocol_interface/required_method]: NextBool */
	NextInt() int/* debug [protocol_interface/required_method]: NextInt */
	NextIntWithUpperBound(upperBound uint) uint/* debug [protocol_interface/required_method]: NextIntWithUpperBound */
	NextUniform() float32/* debug [protocol_interface/required_method]: NextUniform */
}
