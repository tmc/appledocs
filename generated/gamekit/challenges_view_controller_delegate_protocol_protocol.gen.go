// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PChallengesViewControllerDelegate is the GKChallengesViewControllerDelegate protocol interface.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 15.4)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKChallengesViewControllerDelegate
type PChallengesViewControllerDelegate interface {
	// Required methods
	ChallengesViewControllerDidFinish(viewController IGKChallengesViewController)/* debug [protocol_interface/required_method]: ChallengesViewControllerDidFinish */
}

// ChallengesViewControllerDelegate is a delegate implementation builder for the PChallengesViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ChallengesViewControllerDelegate struct {
	_ChallengesViewControllerDidFinish func(viewController IGKChallengesViewController)
}

// SetChallengesViewControllerDidFinish sets the handler for the ChallengesViewControllerDidFinish delegate method.
func (d *ChallengesViewControllerDelegate) SetChallengesViewControllerDidFinish(f func(viewController IGKChallengesViewController)) {
	d._ChallengesViewControllerDidFinish = f
}

// ChallengesViewControllerDidFinish implements the PChallengesViewControllerDelegate interface.
func (d *ChallengesViewControllerDelegate) ChallengesViewControllerDidFinish(viewController IGKChallengesViewController) {
	if d._ChallengesViewControllerDidFinish != nil {
		d._ChallengesViewControllerDidFinish(viewController)
	}
}

// HasChallengesViewControllerDidFinish returns true if a handler for ChallengesViewControllerDidFinish has been set.
func (d *ChallengesViewControllerDelegate) HasChallengesViewControllerDidFinish() bool {
	return d._ChallengesViewControllerDidFinish != nil
}
