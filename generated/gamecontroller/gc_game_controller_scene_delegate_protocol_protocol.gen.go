// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/gameplaykit"
)

// PGCGameControllerSceneDelegate is the GCGameControllerSceneDelegate protocol interface.
//
// Availability:
//   - iOS 18.0+
//   - iPadOS 18.0+
//
// See: doc://com.apple.gamecontroller/documentation/GameController/GCGameControllerSceneDelegate
type PGCGameControllerSceneDelegate interface {
	// Required methods
	SceneDidActivateGameControllerWithContext(scene gameplaykit.Scene, context IGCGameControllerActivationContext)/* debug [protocol_interface/required_method]: SceneDidActivateGameControllerWithContext */
}

// GCGameControllerSceneDelegate is a delegate implementation builder for the PGCGameControllerSceneDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GCGameControllerSceneDelegate struct {
	_SceneDidActivateGameControllerWithContext func(scene gameplaykit.Scene, context IGCGameControllerActivationContext)
}

// SetSceneDidActivateGameControllerWithContext sets the handler for the SceneDidActivateGameControllerWithContext delegate method.
func (d *GCGameControllerSceneDelegate) SetSceneDidActivateGameControllerWithContext(f func(scene gameplaykit.Scene, context IGCGameControllerActivationContext)) {
	d._SceneDidActivateGameControllerWithContext = f
}

// SceneDidActivateGameControllerWithContext implements the PGCGameControllerSceneDelegate interface.
func (d *GCGameControllerSceneDelegate) SceneDidActivateGameControllerWithContext(scene gameplaykit.Scene, context IGCGameControllerActivationContext) {
	if d._SceneDidActivateGameControllerWithContext != nil {
		d._SceneDidActivateGameControllerWithContext(scene, context)
	}
}

// HasSceneDidActivateGameControllerWithContext returns true if a handler for SceneDidActivateGameControllerWithContext has been set.
func (d *GCGameControllerSceneDelegate) HasSceneDidActivateGameControllerWithContext() bool {
	return d._SceneDidActivateGameControllerWithContext != nil
}
