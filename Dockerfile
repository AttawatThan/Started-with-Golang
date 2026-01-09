FROM golang:1.25-bookworm

RUN apt-get update && apt-get install -y git zsh build-essential

WORKDIR /workspace

# Setting Proxy
ENV GOPROXY=https://proxy.golang.org,direct

# Install essential Go tools for VS Code:
# 1. gopls: The "Brain" (IntelliSense, Auto-complete, Jump to definition)
RUN go install golang.org/x/tools/gopls@v0.21.0
# 2. dlv: The "Debugger" (Required to use F5 / Run & Debug)
RUN go install github.com/go-delve/delve/cmd/dlv@v1.26.0
# 3. staticcheck: The "Linter" (Checks for bugs and bad practices)
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.6.1