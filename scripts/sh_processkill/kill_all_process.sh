#!/bin/bash

ports=("50051" "50052")

for port in "${ports[@]}"; do
    lsof -i :$port | awk 'NR!=1 {print $2}' | xargs kill -9
done