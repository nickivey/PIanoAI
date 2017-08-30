package main

import "errors"

// MarkovAI is an implementation of an AI that aims to
// improvise in realtime. In this implementation, the current
// history of real playing is used to generate transition
// probabilities which are used to reconstruct new licks.
type MarkovAI struct {
	// BeatsBetweenLicks specifies the amount of space
	// between each lick before adding an "end"
	BeatsBetweenLicks int
	// keep track of whether it is learning,
	// so learning can be done asynchronously
	IsLearning bool
	HasLearned bool
	// transition matrix for probabilities
	velocityMap map[int]map[int]float64
	// transition matrix for duration of the note
	durationMap map[int]map[int]float64
	// transition matrix for pitches
	pitchMap map[int]map[int]float64
	// transition matrix for the amount of time to wait before
	// playing another note (even if the current note is still playing)
	transitionTimeMap map[int]map[int]float64
}

func (m *MarkovAI) toggleLearning(l bool) {
	m.IsLearning = l
}

// Learn is for calculating the matricies for the transition
// probabilities
func (m *MarkovAI) Learn() (err error) {
	m.toggleLearning(true)
	defer m.toggleLearning(false)

	// TODO: Generate the transition matricies

	m.HasLearned = true
	return
}

// Lick generates a sequence of chords using the Markov
// probabilities. Must run Learn() beforehand.
func (m *MarkovAI) Lick() (lick map[string]Chord, err error) {
	if !m.HasLearned || m.IsLearning {
		return errors.New("Learning must be finished")
	}

	lick = make(map[string]Chord)
	// TODO: Generate lick from the transition probabilities

	return
}
