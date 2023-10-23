#!/usr/bin/bash

set -eux

sudo apt-get update
sudo apt-get install -y bash-completion

go install github.com/go-task/task/v3/cmd/task@latest
wget https://raw.githubusercontent.com/go-task/task/main/completion/bash/task.bash -O ~/task.bash
echo -e "\n. ~/task.bash" >> ~/.bashrc
