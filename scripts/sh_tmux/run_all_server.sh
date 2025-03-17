#!/bin/bash

all_server=("run_experiment_server" "run_userb_server")

#//执行每一个all server
tmux new-session -d -s all_server_session

tmux send-keys "source ./scripts/alias.sh" C-m
tmux send-keys "begin_docker" C-m
tmux split-window
tmux split-window
tmux select-pane -U

# server
for server in "${all_server[@]}"; do
    tmux split-window -h
    tmux send-keys "source ./scripts/alias.sh" C-m
    tmux send-keys "$server" C-m
done
tmux select-pane -t 1
tmux kill-pane
tmux select-pane -D

# control
tmux send-keys "source ./scripts/alias.sh" C-m
tmux send-keys "docker_etcdctl_all" C-m

tmux split-window -h
tmux send-keys "source ./scripts/alias.sh" C-m
tmux send-keys "docker_mysql" C-m

tmux split-window -h
tmux send-keys "source ./scripts/alias.sh" C-m
tmux send-keys "cd_log" C-m



tmux select-pane -t 0
tmux attach-session -t all_server_session