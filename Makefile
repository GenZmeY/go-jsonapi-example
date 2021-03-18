NAME     = go-jsonapi-example
GOCMD    = go
GOBUILD  = $(GOCMD) build
SRCMAIN  = ./cmd/$(NAME)
BINDIR   = bin
BIN      = $(BINDIR)/$(NAME)

.PHONY: all prep build check-build clean

all: build

prep: clean
	go mod init $(NAME); go mod tidy
	mkdir $(BINDIR)
	
build: prep
	$(GOBUILD) -o $(BIN) $(SRCMAIN)

check-build:
	test -e $(BIN)

clean:
	rm -rf $(BINDIR)
