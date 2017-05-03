package main

import (
	"fmt"
)

// Artifact is an artifact
type Artifact struct {
	builderID string
	files     []string
}

// BuilderId returns the builder id
func (a *Artifact) BuilderId() string {
	return a.builderID
}

// Files returns a list of files
func (a *Artifact) Files() []string {
	return a.files
}

// Id Returns the id of the artifact
func (*Artifact) Id() string {
	return "Dummy"
}

// String serializes the artifacts
func (a *Artifact) String() string {
	return fmt.Sprintf("Dummy files: %v", a.files)
}

// State returns state
func (a *Artifact) State(name string) interface{} {
	return nil
}

// Destroy removes the artifacts from disk
func (a *Artifact) Destroy() error {
	return nil
}
