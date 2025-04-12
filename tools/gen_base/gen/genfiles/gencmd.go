package genfiles

import (
	"fmt"
	"genbase/gen/utils"
)

var (
	template_cmd_client_dir = ""
	aim_cmd_client_dir      = ""
	template_cmd_server_dir = ""
	aim_cmd_server_dir      = ""
)

func GenCmd() {
	doinit()
	gen_client_cmd()
	gen_server_cmd()
}

func gen_client_cmd() error {
	// read file
	filepath := template_cmd_client_dir + "client.go"
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
	filepath = aim_cmd_client_dir + "client.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[3]Generated cmd/client.go\n")
	return nil
}

func gen_server_cmd() error {
	// read file
	filepath := template_cmd_server_dir + "server.go"
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
	filepath = aim_cmd_server_dir + "server.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[4]Generated cmd/server.go\n")
	return nil
}
