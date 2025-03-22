package genfiles

import (
	"fmt"
	"genbase/gen/utils"
	"strings"
)

var (
	template_internal_dir         = ""
	template_internal_config_dir  = ""
	template_internal_handler_dir = ""
	template_internal_server_dir  = ""
	template_internal_service_dir = ""

	aim_internal_dir         = ""
	aim_internal_config_dir  = ""
	aim_internal_handler_dir = ""
	aim_internal_server_dir  = ""
	aim_internal_service_dir = ""

	service_func_names = []string{}
	service_func_file  = []string{}
)

func GenInternal() {
	doinit()
	gen_config()
	gen_handler()
	gen_server()
	gen_service()
	gen_service_func()
}

func gen_config() error {
	// read file
	filepath := template_internal_config_dir + "config.go"
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
	filepath = aim_internal_config_dir + "config.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[6]Generated config.go\n")
	return nil
}

func gen_handler() error {
	// read file
	filepath := template_internal_handler_dir + "handler.go"
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
	filepath = aim_internal_handler_dir + "handler.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[7]Generated handler.go\n")
	return nil
}

func gen_server() error {
	// read file
	filepath := template_internal_server_dir + "server.go"
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
	filepath = aim_internal_server_dir + "server.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[8]Generated server.go\n")

	// srv_service.go
	// read file
	filepath = template_internal_server_dir + "srv_service.go"
	content, err = utils.ReadFile(filepath)
	if err != nil {
		fmt.Printf("read failed: %v\n", err)
	}

	// transform file
	// 特别填充 00003
	var specialToken = utils.Token{
		Before: "%00003",
	}
	{
		for i := 1; i < len(Config.FuncInfo); i += 2 {
			BaseInfo := `
func (s *%3Server) %002(ctx context.Context, req *pb.%002Request) (resp *pb.%002Response, err error) {
	return s.%3Service.%002(ctx, req)
}`
			ChangeInfo := strings.ReplaceAll(BaseInfo, "%002", Config.FuncInfo[i].Value)
			specialToken.After += ChangeInfo + "\n\n\t"
		}
	}
	content, err = utils.TransformFile(content, []utils.Token{specialToken})
	if err != nil {
		fmt.Printf("TransformFile failed: %v\n", err)
	}

	// write file
	filepath = aim_internal_server_dir + "srv_service.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[9]Generated src_service.go\n")
	return nil
}

func gen_service() error {
	// read file
	filepath := template_internal_service_dir + "service.go"
	content, err := utils.ReadFile(filepath)
	if err != nil {
		fmt.Printf("read failed: %v\n", err)
	}

	// transform file// 特别填充 00002
	var specialToken = utils.Token{
		Before: "%00002",
	}
	{
		for i := 1; i < len(Config.FuncInfo); i += 2 {
			BaseInfo := `%002(ctx context.Context, in *pb.%002Request) (*pb.%002Response, error)`
			ChangeInfo := strings.ReplaceAll(BaseInfo, "%002", Config.FuncInfo[i].Value)
			specialToken.After += ChangeInfo + "\n\t"
		}
	}
	content, err = utils.TransformFile(content, []utils.Token{specialToken})
	if err != nil {
		fmt.Printf("TransformFile failed: %v\n", err)
	}

	// write file
	filepath = aim_internal_service_dir + "service.go"
	err = utils.WriteFile(filepath, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", filepath, err)
	}

	fmt.Printf("[10]Generated service.go\n")
	return nil
}

func gen_service_func() {
	ok := true
	for i := 0; i < len(service_func_names); i++ {
		// read file
		filepath := template_internal_service_dir + "svc_func.go"
		content, err := utils.ReadFile(filepath)
		if err != nil {
			fmt.Printf("ReadFile failed: %v\n", err)
			ok = false
		}
		// transform file
		content, err = utils.TransformFile(content, []utils.Token{
			{Before: "%0001", After: Config.FuncInfo[i*2].Value},
			{Before: "%0002", After: Config.FuncInfo[i*2+1].Value},
		})
		if err != nil {
			fmt.Printf("TransformFile failed: %v\n", err)
			ok = false
		}
		// write file
		file_name := aim_internal_service_dir + service_func_file[i]
		err = utils.WriteFile(file_name, content)
		if err != nil {
			fmt.Printf("%s failed: %v\n", file_name, err)
			ok = false
		}
	}
	if ok {
		fmt.Println("[11]Generated server funcs")
	} else {
		fmt.Println("Failed to generate server funcs")
	}

}
