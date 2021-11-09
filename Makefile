# Variables
DESTDIR ?=
WWWROOT ?= /var/www/html
WWWPREFIX ?= /keygaen

all: build

# Build
build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm main.go
	go run main.go -prefix $(WWWPREFIX)
	cp -rf web/* out/web/web
	tar -cvzf out/keygaen.tar.gz -C out/web .

# Install
install:
	mkdir -p $(DESTDIR)$(WWWROOT)$(WWWPREFIX)
	cp -rf out/web/* $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Uninstall
uninstall:
	rm -rf $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Run
run:
	GOARCH=wasm GOOS=js go build -o web/app.wasm main.go
	go run main.go -serve

# Clean
clean:
	rm -rf out web/app.wasm
