package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type rubyArgs struct {
	rbEnvEnabled   bool
	rbEnvVersion   string
	hasRB          bool
	hasRakeFile    bool
	hasGemFile     bool
	displayVersion bool
}

func bootStrapRubyTest(args *rubyArgs) *ruby {
	env := new(MockedEnvironment)
	env.On("hasCommand", "rbenv").Return(args.rbEnvEnabled)
	env.On("runCommand", "rbenv", []string{"rvm-prompt"}).Return(args.rbEnvVersion, nil)
	env.On("hasFiles", "*.rb").Return(args.hasRB)
	env.On("hasFiles", "Rakefile").Return(args.hasRakeFile)
	env.On("hasFiles", "Gemfile").Return(args.hasGemFile)
	props := &properties{
		values: map[Property]interface{}{
			DisplayVersion: args.displayVersion,
		},
	}
	r := &ruby{
		env:   env,
		props: props,
	}
	return r
}

func TestRubyWriterDisabled(t *testing.T) {
	args := &rubyArgs{}
	ruby := bootStrapRubyTest(args)
	assert.False(t, ruby.enabled(), "ruby is not enabled")
}
