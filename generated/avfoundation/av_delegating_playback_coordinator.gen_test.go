// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewDelegatingPlaybackCoordinator

// ExampleDelegatingPlaybackCoordinator_ReapplyCurrentItemStateToPlaybackControlDelegate demonstrates using ReapplyCurrentItemStateToPlaybackControlDelegate on a DelegatingPlaybackCoordinator instance.
// Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
func ExampleDelegatingPlaybackCoordinator_ReapplyCurrentItemStateToPlaybackControlDelegate() {
	obj := avfoundation.NewDelegatingPlaybackCoordinator()
	obj.ReapplyCurrentItemStateToPlaybackControlDelegate()
	// Output:
	}

