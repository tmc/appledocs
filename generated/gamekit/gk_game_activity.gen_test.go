// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewGameActivity

// ExampleGameActivity_End demonstrates using End on a GameActivity instance.
// Ends the game activity if it’s not already ended.
func ExampleGameActivity_End() {
	obj := gamekit.NewGameActivity()
	obj.End()
	// Output:
	}

// ExampleGameActivity_MakeMatchRequest demonstrates using MakeMatchRequest on a GameActivity instance.
// Makes a match request object with information from the activity, which you can use to find matches for the local player.
func ExampleGameActivity_MakeMatchRequest() {
	obj := gamekit.NewGameActivity()
	_ = obj.MakeMatchRequest()
	// Output:
	}

// ExampleGameActivity_Pause demonstrates using Pause on a GameActivity instance.
// Pauses the game activity if it’s not already paused.
func ExampleGameActivity_Pause() {
	obj := gamekit.NewGameActivity()
	obj.Pause()
	// Output:
	}

// ExampleGameActivity_Resume demonstrates using Resume on a GameActivity instance.
// Resumes the game activity if it was paused.
func ExampleGameActivity_Resume() {
	obj := gamekit.NewGameActivity()
	obj.Resume()
	// Output:
	}

// ExampleGameActivity_Start demonstrates using Start on a GameActivity instance.
// Starts the game activity if it’s not already started.
func ExampleGameActivity_Start() {
	obj := gamekit.NewGameActivity()
	obj.Start()
	// Output:
	}

