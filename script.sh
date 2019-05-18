#!/bin/bash


sudo apt-get update
sudo apt-get -y upgrade
cd /tmp
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
sudo tar -xvf go1.11.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
source ~/.profile
#
cd /tmp
mkdir Servidor
cd Servidor
git clone https://github.com/Caceres-Joseph/-SO1-Proyecto1_Servidor.git
cd /tmp/Servidor/-SO1-Proyecto1_Servidor/
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/cpu
go run memoria.go &
