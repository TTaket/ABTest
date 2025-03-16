package genfiles

import (
	"fmt"
	"genbase/gen/utils"
	"strings"
)

var (
	template_client_dir = ""
	aim_client_dir      = ""
	client_func_names   = []string{}
	client_func_file    = []string{}
)

func GenClient() {
	doinit()
	gen_client()
	gen_client_func()
}

func gen_client() error {
	// read file
	content, err := utils.ReadFile(template_client_dir + "client.go")
	if err != nil {
		fmt.Printf("read failed: %v\n", err)
	}
	// transform file
	// 特别填充 00001
	var specialToken = utils.Token{
		Before: "%00001",
	}
	{
		for i := 1; i < len(Config.FuncInfo); i += 2 {
			BaseInfo := `%002(ctx context.Context, in *pb.%002Request, opts ...grpc.CallOption) (*pb.%002Response, error)`
			ChangeInfo := strings.ReplaceAll(BaseInfo, "%002", Config.FuncInfo[i].Value)
			specialToken.After += ChangeInfo + "\n\t"
		}
	}
	content, err = utils.TransformFile(content, []utils.Token{specialToken})
	if err != nil {
		fmt.Printf("TransformFile failed: %v\n", err)
	}
	// write file
	file_name := aim_client_dir + "client.go"
	err = utils.WriteFile(file_name, content)
	if err != nil {
		fmt.Printf("%s failed: %v\n", file_name, err)
	}

	fmt.Printf("[1]Generated client.go\n")
	return nil
}

func gen_client_func() {
	ok := true
	for i := 0; i < len(client_func_names); i++ {
		// read file
		content, err := utils.ReadFile(template_client_dir + "cli_func.go")
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
		file_name := aim_client_dir + client_func_file[i]
		err = utils.WriteFile(file_name, content)
		if err != nil {
			fmt.Printf("%s failed: %v\n", file_name, err)
			ok = false
		}
	}
	if ok {
		fmt.Println("[2]Generated client funcs")
	} else {
		fmt.Println("Failed to generate client funcs")
	}

}
