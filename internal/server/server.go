// Package server contains server-side code
package server

import (
	"fmt"
	"log"

	"github.com/SOMAS2020/SOMAS2020/internal/common"
)

// Server represents the primary server interface exposed to the simulation.
type Server interface {
	// EntryPoint function that returns a list of historic common.GameStates until the
	// game ends.
	EntryPoint() []common.GameState
}

// SOMASServer implements Server.
type SOMASServer struct {
	gameState common.GameState
}

// SOMASServerFactory returns an instance of the main server we use.
func SOMASServerFactory() Server {
	return &SOMASServer{
		gameState: common.GameState{
			Day:         1,
			ClientInfos: getClientInfoFromRegisteredClients(common.RegisteredClients),
		},
	}
}

// EntryPoint function that returns a list of historic common.GameStates until the
// game ends.
func (s *SOMASServer) EntryPoint() []common.GameState {
	states := []common.GameState{s.gameState}

	for anyClientsAlive(s.gameState.ClientInfos) {
		s.gameState.Day++
		s.runRound()
		states = append(states, s.gameState)
	}

	return states
}

// runRound runs a round (day) of the game.
func (s *SOMASServer) runRound() {
	// TODO: Implement round logic
	s.getEcho("HELLO WORLD!")
	s.killAllClients()
}

// getEcho retrieves an echo from all the clients and make sure they are the same.
func (s *SOMASServer) getEcho(str string) error {
	cis := s.gameState.ClientInfos
	for _, id := range common.TeamIDs {
		ci := cis[id]
		c := ci.Client
		got := c.Echo(str)
		if str != got {
			return fmt.Errorf("Echo error: want '%v' got '%v' from %v",
				str, got, c.GetID())
		}
		s.logf("Received echo `%v` from %v", str, c.GetID())
	}
	return nil
}

// killAllClients sets all the Alive states of the clients to false to end the game.
// Only used for testing to preemptively end the game.
func (s *SOMASServer) killAllClients() {
	for _, id := range common.TeamIDs {
		ci := s.gameState.ClientInfos[id]
		ci.Alive = false
		s.gameState.ClientInfos[id] = ci
	}
}

// logf is the server's default logger.
func (s *SOMASServer) logf(format string, a ...interface{}) {
	log.Printf("[SERVER]: %v", fmt.Sprintf(format, a...))
}
