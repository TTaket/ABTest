#!/bin/bash
# kill all port
zsh ./scripts/sh_processkill/kill_all_process.sh

tmux kill-session -t all_server_session
