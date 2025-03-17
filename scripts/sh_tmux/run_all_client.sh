#!/bin/bash

all_client=("run_experiment_client" "run_userb_client")

#//执行每一个all client
tmux new-session -d -s all_client_session

tmux send-keys "source ./scripts/alias.sh" C-m
tmux send-keys "begin_docker" C-m

for client in "${all_client[@]}"; do
    tmux split-window 
    tmux send-keys "source ./scripts/alias.sh" C-m
    tmux send-keys "$client" C-m
    tmux select-layout tiled
done

tmux select-pane -t 0
tmux attach-session -t all_client_session