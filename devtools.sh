#!/bin/bash

echo GOBIN :${GOBIN} 경로에 개발 도구를 설치합니다

go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v golang.org/x/tools/cmd/goimports@latest
go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install -v github.com/rogpeppe/godef@latest
go install -v github.com/swaggo/swag/cmd/swag@latest
go install -v github.com/stamblerre/gocode@latest
go install -v golang.org/x/tools/gopls@latest
go install -v honnef.co/go/tools/cmd/staticcheck@latest