ALL := admincoretest

ifeq ($(P), LINUX)
	GOBUILD := GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install
else
	GOBUILD := go install
endif

export CURRENT_DIR = $(shell pwd)
# 更新 GOPATH 变量, 这样可以共享 GOPATH/src 下的文件包, 避免重复下载依赖包
export GOPATH := $(CURRENT_DIR):$(GOPATH)

all:
	$(GOBUILD) $(ALL)

run:
	$(GOBUILD) $(ALL)
	bin/admincoretest conf/admincoretest.yaml

clean:
	rm -fr pkg bin/admincoretest
