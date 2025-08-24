package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"plugin"

	"neuralops/pkg/plugin"
)

const (
	pluginCodePath = "/app/plugin.go"
	pluginSOPath   = "/app/plugin.so"
	inputData      = "some input data"
)

func main() {
	// In a real implementation, the generated code would be passed as an
	// input artifact from the orchestrator. For now, we'll just check
	// if the file exists.
	if _, err := os.Stat(pluginCodePath); os.IsNotExist(err) {
		// For local testing, create a dummy plugin file
		dummyCode := `
package main

import (
	"context"
	"fmt"
)

type Runner struct{}

func (r *Runner) Apply(ctx context.Context, input []byte) ([]byte, error) {
	fmt.Println("Hello from the dummy plugin!")
	return nil, nil
}

func New() interface{} {
	return &Runner{}
}
`
		if err := ioutil.WriteFile(pluginCodePath, []byte(dummyCode), 0644); err != nil {
			log.Fatalf("failed to create dummy plugin file: %v", err)
		}
	}


	// Compile the plugin
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginSOPath, pluginCodePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to compile plugin: %v", err)
	}

	// Load the plugin
	p, err := plugin.Open(pluginSOPath)
	if err != nil {
		log.Fatalf("failed to open plugin: %v", err)
	}

	// Look up the New symbol
	newFuncSymbol, err := p.Lookup("New")
	if err != nil {
		log.Fatalf("failed to lookup New symbol: %v", err)
	}

	newFunc, ok := newFuncSymbol.(func() interface{})
	if !ok {
		log.Fatalf("unexpected type from module symbol")
	}

	runner, ok := newFunc().(plugin.Runner)
	if !ok {
		log.Fatalf("plugin does not implement plugin.Runner interface")
	}


	// Execute the plugin
	output, err := runner.Apply(context.Background(), []byte(inputData))
	if err != nil {
		log.Fatalf("failed to run plugin: %v", err)
	}

	fmt.Printf("Plugin output: %s\n", string(output))
}
