package main

/*
The Dummy Artifact builder creates dummy artifacts based on configuration
*/

import (
	"fmt"

	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type dummyConfig struct {
	BuilderID string   `mapstructure:"builder_id"`
	Files     []string `mapstructure:"files"`
}

// Builder is this builder
type Builder struct {
	config dummyConfig
}

// Prepare prepares the build
func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	err := config.Decode(&b.config, &config.DecodeOpts{
		Interpolate: true,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)

	if err != nil {
		return nil, err
	}

	// Accumulate any errors
	errs := new(packer.MultiError)

	if b.config.BuilderID == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("builder_id must be set"))
	}

	if b.config.Files == nil || len(b.config.Files) == 0 {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("files must be set"))
	}

	if len(errs.Errors) > 0 {
		return nil, errs
	}

	return nil, nil
}

// Run runs the build
func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	artifact := new(Artifact)

	artifact.files = b.config.Files
	artifact.builderID = b.config.BuilderID

	return artifact, nil
}

// Cancel cancels a possibly running Builder
func (b *Builder) Cancel() {
}
