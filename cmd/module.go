package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// moduleCmd generates a new dawn module
var moduleCmd = &cobra.Command{
	Use:     "module NAME",
	Aliases: []string{"m"},
	Short:   "Generate a new dawn module",
	Example: moduleExamples,
	Args:    cobra.MinimumNArgs(1),
	RunE:    moduleRunE,
}

func moduleRunE(cmd *cobra.Command, args []string) (err error) {
	now := time.Now()

	name := args[0]

	dir, _ := os.Getwd()

	modulePath := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, name)
	if err := createModule(modulePath, name); err != nil {
		return err
	}

	cmd.Printf(moduleCreatedTemplate,
		modulePath, name, formatLatency(time.Since(now)))

	return
}

func createModule(modulePath, name string) (err error) {
	if err = os.Mkdir(modulePath, 0750); err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = os.RemoveAll(modulePath)
		}
	}()

	// create module.go
	if err = createFile(fmt.Sprintf("%s%c%s.go", modulePath, os.PathSeparator, name),
		moduleContent(name)); err != nil {
		return
	}

	// create module_test.go
	return createFile(fmt.Sprintf("%s%c%s_test.go", modulePath, os.PathSeparator, name),
		moduleTestContent(name))
}

func moduleContent(name string) string {
	temp := `package {{module}}

import (
	"github.com/go-dawn/dawn"
	"github.com/gofiber/fiber/v2"
)

type {{module}}Module struct {
	dawn.Module
}

// New returns the module
func New() dawn.Moduler {
	return &{{module}}Module{
	}
}

func (m *{{module}}Module) String() string {
	return "{{module}}"
}

func (m *{{module}}Module) Init() dawn.Cleanup {
	// you can implement me or remove me

	// Read config and init module

	return func() {
		// Put cleanup stuff here if any
	}
}

func (m *{{module}}Module) Boot() {
	// you can implement me or remove me
}

func (m *{{module}}Module) RegisterRoutes(router fiber.Router) {
	// implement me or remove me
}`
	return strings.ReplaceAll(temp, "{{module}}", name)
}

func moduleTestContent(name string) string {
	temp := `package {{module}}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Module_Name(t *testing.T) {
	assert.Equal(t, "{{module}}", New().String())
}

func Test_Module_Init(t *testing.T) {
	m := &{{module}}Module{}

	m.Init()()

	// more assertions
}

func Test_Module_Boot(t *testing.T) {
	m := &{{module}}Module{}

	m.Boot()

	// more assertions
}`
	return strings.ReplaceAll(temp, "{{module}}", name)
}

var (
	moduleCreatedTemplate = `
Scaffolding module in %s

  Done. Now run:

  cd %s
  go test . -cover

🎊  Done in %s.
`

	moduleExamples = `  dawn module hello
    Generates a new module dir named hello includes boilerplate code`
)
