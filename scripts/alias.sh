#!/bin/bash

# docker
alias begin_docker="docker-compose -f ./docker-compose.yml up -d"
alias end_docker="docker-compose -f ./docker-compose.yml down"
alias restart_docker="docker-compose -f ./docker-compose.yml restart"

# all
alias run_all_server="zsh ./scripts/sh_tmux/run_all_server.sh"
alias run_all_client="zsh ./scripts/sh_tmux/run_all_client.sh"
alias kill_all_server="zsh ./scripts/sh_tmux/kill_all_server.sh"
alias kill_all_client="zsh ./scripts/sh_tmux/kill_all_client.sh"
alias kill_all_process="zsh ./scripts/sh_processkill/kill_all_process.sh"

# server and client
alias run_experiment_server="go run ./apps/Micro/experiment/cmd/server_main/server.go"
alias run_experiment_client="go run ./apps/Micro/experiment/cmd/client_main/client.go"
alias run_userb_server="go run ./apps/Micro/userb/cmd/server_main/server.go"
alias run_userb_client="go run ./apps/Micro/userb/cmd/client_main/client.go"


# etcd
ETCD_HOST=127.0.0.1
ETCD_ENDPOINTS=${ETCD_HOST}:12379
alias docker_etcdctl_all="etcdctl --endpoints=${ETCD_ENDPOINTS} get --prefix '' "
alias etcdctl_all="etcdctl  get --prefix '' "

# mysql
alias docker_mysql_userb="mysql -h 127.0.0.1 -P 13306 -u root -p'abtest2025'"
alias docker_mysql_experiment="mysql -h 127.0.0.1 -P 13307 -u root -p'abtest2025'"

# proto
alias protoc_gen="sh ./scripts/genproto.sh"

# log
alias cd_log="cd /var/log/ABTest"