// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAgentDelegate is the GKAgentDelegate protocol interface.
//
// Implement this protocol to synchronize the state of an agent with its visual representation in your game.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gameplaykit/documentation/GameplayKit/GKAgentDelegate
type PAgentDelegate interface {
	// Optional methods
	AgentDidUpdate(agent IGKAgent)
	HasAgentDidUpdate() bool
	AgentWillUpdate(agent IGKAgent)
	HasAgentWillUpdate() bool
}

// AgentDelegate is a delegate implementation builder for the PAgentDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AgentDelegate struct {
	_AgentDidUpdate func(agent IGKAgent)
	_AgentWillUpdate func(agent IGKAgent)
}

// SetAgentDidUpdate sets the handler for the AgentDidUpdate delegate method.
//
// Tells the delegate that an agent has just performed a simulation step.
func (d *AgentDelegate) SetAgentDidUpdate(f func(agent IGKAgent)) {
	d._AgentDidUpdate = f
}

// SetAgentWillUpdate sets the handler for the AgentWillUpdate delegate method.
//
// Tells the delegate that an agent is about to perform its next simulation step.
func (d *AgentDelegate) SetAgentWillUpdate(f func(agent IGKAgent)) {
	d._AgentWillUpdate = f
}

// AgentDidUpdate implements the PAgentDelegate interface.
func (d *AgentDelegate) AgentDidUpdate(agent IGKAgent) {
	if d._AgentDidUpdate != nil {
		d._AgentDidUpdate(agent)
	}
}

// HasAgentDidUpdate returns true if a handler for AgentDidUpdate has been set.
func (d *AgentDelegate) HasAgentDidUpdate() bool {
	return d._AgentDidUpdate != nil
}

// AgentWillUpdate implements the PAgentDelegate interface.
func (d *AgentDelegate) AgentWillUpdate(agent IGKAgent) {
	if d._AgentWillUpdate != nil {
		d._AgentWillUpdate(agent)
	}
}

// HasAgentWillUpdate returns true if a handler for AgentWillUpdate has been set.
func (d *AgentDelegate) HasAgentWillUpdate() bool {
	return d._AgentWillUpdate != nil
}
