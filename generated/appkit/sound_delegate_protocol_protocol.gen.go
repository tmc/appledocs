// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSoundDelegate is the NSSoundDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSoundDelegate
type PSoundDelegate interface {
	// Optional methods
	SoundDidFinishPlaying(sound ISound, flag bool)
	HasSoundDidFinishPlaying() bool
}

// SoundDelegate is a delegate implementation builder for the PSoundDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SoundDelegate struct {
	_SoundDidFinishPlaying func(sound ISound, flag bool)
}

// SetSoundDidFinishPlaying sets the handler for the SoundDidFinishPlaying delegate method.
//
// This delegate method is called when an   instance has completed playback of its sound data.
func (d *SoundDelegate) SetSoundDidFinishPlaying(f func(sound ISound, flag bool)) {
	d._SoundDidFinishPlaying = f
}

// SoundDidFinishPlaying implements the PSoundDelegate interface.
func (d *SoundDelegate) SoundDidFinishPlaying(sound ISound, flag bool) {
	if d._SoundDidFinishPlaying != nil {
		d._SoundDidFinishPlaying(sound, flag)
	}
}

// HasSoundDidFinishPlaying returns true if a handler for SoundDidFinishPlaying has been set.
func (d *SoundDelegate) HasSoundDidFinishPlaying() bool {
	return d._SoundDidFinishPlaying != nil
}
