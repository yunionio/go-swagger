package generate

import (
	"github.com/go-swagger/go-swagger/generator"
	flags "github.com/jessevdk/go-flags"
)

// Markdown generates a markdown representation of the spec
type Markdown struct {
	Spec           flags.Filename `long:"spec" short:"f" description:"the spec file to use (default swagger.{json,yml,yaml})"`
	TemplateDir    flags.Filename `long:"template-dir" short:"T" description:"alternative template override directory"`
	ConfigFile     flags.Filename `long:"config-file" short:"C" description:"configuration file to use for overriding template options"`
	Output         flags.Filename `long:"output" short:"o" description:"the file to write to"`
	DumpData       bool           `long:"dump-data" description:"when present dumps the json for the template generator instead of generating files"`
	SkipModels     bool           `long:"skip-models" description:"no models will be generated when this flag is specified"`
	SkipOperations bool           `long:"skip-operations" description:"no operations will be generated when this flag is specified"`
	SkipSupport    bool           `long:"skip-support" description:"no supporting files will be generated when this flag is specified"`
	Operations     []string       `long:"operation" short:"O" description:"specify an operation to include, repeat for multiple"`
	Tags           []string       `long:"tags" description:"the tags to include, if not specified defaults to all"`
	Models         []string       `long:"model" short:"M" description:"specify a model to include, repeat for multiple"`
}

// Execute runs this command
func (s *Markdown) Execute(args []string) error {
	opts := &generator.GenOpts{
		Spec:              string(s.Spec),
		Target:            "./",
		APIPackage:        "operations",
		ModelPackage:      "models",
		ServerPackage:     "restapi",
		ClientPackage:     "client",
		Principal:         "interface{}",
		DefaultScheme:     "http",
		IncludeModel:      !s.SkipModels,
		IncludeValidator:  !s.SkipModels,
		IncludeHandler:    !s.SkipOperations,
		IncludeParameters: !s.SkipOperations,
		IncludeResponses:  !s.SkipOperations,
		IncludeURLBuilder: !s.SkipOperations,
		IncludeSupport:    !s.SkipSupport,
		ExcludeSpec:       true,
		TemplateDir:       string(s.TemplateDir),
		DumpData:          s.DumpData,
		Models:            s.Models,
		Operations:        s.Operations,
		Tags:              s.Tags,
		Name:              "",
	}

	return generator.GenerateMarkdown(string(s.Output), s.Models, s.Operations, opts)
}
