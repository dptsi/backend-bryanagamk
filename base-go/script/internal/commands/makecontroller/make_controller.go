package makecontroller

import (
	"fmt"
	"os"
	"strings"

	"github.com/mikestefanello/hooks"
	"github.com/stoewer/go-strcase"
	"its.ac.id/base-go/script/internal/app"
)

func init() {
	app.HookBoot.Listen(func(event hooks.Event[*app.Script]) {
		event.Msg.AddCommand(app.Command{
			Name:        "make:controller",
			Description: "Create new controller",
			Usage:       "make:controller <module_name> <controller_name>",
			Handler:     makeController,
		})
	})
}

func makeController(args []string) {
	if len(args) == 0 {
		fmt.Println("No module name provided")
		return
	}
	if len(args) == 1 {
		fmt.Println("No controller name provided")
		return
	}
	module := args[0]
	name := args[1]
	path := fmt.Sprintf("modules/%s", module)

	if _, err := os.Stat(path); err != nil {
		fmt.Printf("Module %s doesn't exist\n", module)
		return
	}
	os.MkdirAll(fmt.Sprintf("%s/internal/presentation/controllers", path), os.ModePerm)
	if err := createSkeleton(module, name, path); err != nil {
		fmt.Println(err)
		return
	}
}

func createSkeleton(module string, name string, path string) error {
	snakeCased := strings.ToLower(strcase.SnakeCase(name))
	controllerPath := fmt.Sprintf("%s/internal/presentation/controllers/%s_controller.go", path, snakeCased)
	file, err := os.Create(controllerPath)
	if err != nil {
		return err
	}
	fmt.Fprintf(
		file,
		`package controllers

import (
	"github.com/samber/do"
)

type %sController struct {
	i *do.Injector
}

func New%sController() *%sController {
	i := do.DefaultInjector
	return &%sController{i: i}
}
		`,
		name,
		name,
		name,
		name,
	)
	file.Close()

	return nil
}
