%1 - modulename
%2 - modulename(大写)
%3 - moudlename(小写)
%4 - rpc 服务名称（最底下的service）
%5 - rpc Service Client

//func
%001 - Function_01（小写）
%002 - Function_01（大写）
%003 - Function_02（大写）
%004 - Function_02（大写）
。。。。

// 特殊指定
%0001 - 当前函数小写
%0002 - 当前函数大写
%00001 - 用于填充不定数多的函数 在 client.go 
%00002 - 用于填充不定数多的函数 在 service.go
%00003 - 用于填充不定数多的函数 在 srv_service.go


base:
  _%1: experiment #modulename # example:experiment
  _%2: Experiment #modulename(大写) # example:Experiment
  _%3: experiment #modulename(小写) # example:experiment
  _%4: ExperimentService #rpc 服务名称（最底下的service） # example:ExperimentService
  _%5: NewExperimentServiceClient #New rpc Service Client  # example:NewExperimentServiceClient

func:
  - key: _%001
    value: createexperiment #Function_01（小写） # example:createexperiment
  - key: _%002
    value: CreateExperiment #Function_01（大写） # example:CreateExperiment
  - key: _%003
    value: getexperiment #Function_02（小写）# example:getexperiment
  - key: _%004
    value: GetExperiment #Function_02（大写） # example:GetExperiment