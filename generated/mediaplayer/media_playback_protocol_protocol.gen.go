// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

// PMediaPlayback is the MPMediaPlayback protocol interface.
//
// A protocol that defines the interface for controlling audio media playback.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPMediaPlayback
type PMediaPlayback interface {
	// Required methods
	BeginSeekingBackward()/* debug [protocol_interface/required_method]: BeginSeekingBackward */
	BeginSeekingForward()/* debug [protocol_interface/required_method]: BeginSeekingForward */
	EndSeeking()/* debug [protocol_interface/required_method]: EndSeeking */
	Pause()/* debug [protocol_interface/required_method]: Pause */
	Play()/* debug [protocol_interface/required_method]: Play */
	PrepareToPlay()/* debug [protocol_interface/required_method]: PrepareToPlay */
	Stop()/* debug [protocol_interface/required_method]: Stop */
}
