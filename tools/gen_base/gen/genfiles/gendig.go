package genfiles

import (
	"fmt"
	"genbase/gen/utils"
)

var (
	template_dig_dir = ""
	aim_dig_dir      = ""
)

func GenDig() {
	doinit()
	gen_dig()
}

func gen_dig() error {
	// read file
	filepath := template_dig_dir + "dig.go"
	content, err := utils.ReadFile(filepath)
	if err != nil {
		fmt.Printf("read failed: %v\n", err)
	}

	// transform file
	content, err = utils.TransformFile(content, []utils.Token{})
	if err != nil {
		fmt.Printf("TransformFile failed: %v\n", err)
	}

	// write file
	filepath = aim_dig_dir + "dig.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[5]Generated dig.go\n")
	return nil
}
