# Build package binary for arm platforms: armV5
env GOOS=linux GOARCH=arm GOARM=5 go install <package>.go

# Build package binary for arm platforms: armV7
env GOOS=linux GOARCH=arm GOARM=7 go install <geometry>.go

# Steps to enable go code syntax highlighting in vim

sudo apt-get install vim-gocomplete gocode vim-syntax-go

vim-addon-manager install go-syntax

vim-addon-manager install gocode
