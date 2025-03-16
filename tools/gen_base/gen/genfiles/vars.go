package genfiles

import (
	"fmt"
	"genbase/gen/utils"
)

var (
	Config utils.Config
)

var (
	aim_dir      = "./aim/"
	gen_dir      = "./gen/"
	template_dir = gen_dir + "template/"
	config_dir   = gen_dir + "config/"
)
var (
	yaml_file_path = config_dir + "config.yml"
	isInit         = false
)

func doinit() {
	if isInit {
		return
	}
	isInit = true
	err := utils.ReadYamlFile(yaml_file_path, &Config)
	if err != nil {
		fmt.Printf("ReadYamlFile failed: %v\n", err)
		return
	}

	// 填充client_func_names 和 client_func_file
	for i := 0; i < len(Config.FuncInfo); i += 2 {
		client_func_names = append(client_func_names, Config.FuncInfo[i].Value)
		client_func_file = append(client_func_file, "cli_"+Config.FuncInfo[i].Value+".go")
	}

	// 填充service_func_names 和 service_func_file
	for i := 0; i < len(Config.FuncInfo); i += 2 {
		service_func_names = append(service_func_names, Config.FuncInfo[i].Value)
		service_func_file = append(service_func_file, "svc_"+Config.FuncInfo[i].Value+".go")
	}

	// var
	//client
	template_client_dir = template_dir + "client/"
	aim_client_dir = aim_dir + Config.BaseInfo.ModuleName + "/client/"

	//cmd
	template_cmd_client_dir = template_dir + "cmd/client_main/"
	aim_cmd_client_dir = aim_dir + Config.BaseInfo.ModuleName + "/cmd/client_main/"
	template_cmd_server_dir = template_dir + "cmd/server_main/"
	aim_cmd_server_dir = aim_dir + Config.BaseInfo.ModuleName + "/cmd/server_main/"

	//dig
	template_dig_dir = template_dir + "dig/"
	aim_dig_dir = aim_dir + Config.BaseInfo.ModuleName + "/dig/"

	//internal
	template_internal_dir = template_dir + "internal/"
	template_internal_config_dir = template_internal_dir + "config/"
	template_internal_handler_dir = template_internal_dir + "handler/"
	template_internal_server_dir = template_internal_dir + "server/"
	template_internal_service_dir = template_internal_dir + "service/"
	aim_internal_dir = aim_dir + Config.BaseInfo.ModuleName + "/internal/"
	aim_internal_config_dir = aim_internal_dir + "config/"
	aim_internal_handler_dir = aim_internal_dir + "handler/"
	aim_internal_server_dir = aim_internal_dir + "server/"
	aim_internal_service_dir = aim_internal_dir + "service/"
}
