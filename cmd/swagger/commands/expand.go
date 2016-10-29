package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
)

// ExpandSpec is a command that resolves all the references in a swagger document
type ExpandSpec struct {
	// SchemaURL string `long:"schema" description:"The schema url to use" default:"http://swagger.io/v2/schema.json"`
}

// Execute validates the spec
func (c *ExpandSpec) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("The expand command requires the swagger document url to be specified")
	}

	swaggerDoc := args[0]
	specDoc, err := loads.Spec(swaggerDoc)
	if err != nil {
		return err
	}

	sp := specDoc.Spec()
	if e := spec.ExpandSpec(sp); e != nil {
		return e
	}
	b, err := json.MarshalIndent(sp, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
