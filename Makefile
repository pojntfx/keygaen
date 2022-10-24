# Public variables
DESTDIR ?=

WWWROOT ?= /var/www/html
WWWPREFIX ?= /keygaen

PREFIX ?= /usr/local
OUTPUT_DIR ?= out
BUILD_DIR ?= build
STATIC_DIR ?= web
DST ?=

# Private variables
clis = keygaen-cli
pwas = keygaen-pwa
all: $(addprefix build-cli/,$(clis)) $(addprefix build-pwa/,$(pwas))

# Build
build: $(addprefix build-cli/,$(clis)) $(addprefix build-pwa/,$(pwas))

$(addprefix build-cli/,$(clis)):
ifdef DST
	go build -o $(DST) ./cmd/$(subst build-cli/,,$@)
else
	go build -o $(OUTPUT_DIR)/$(subst build-cli/,,$@) ./cmd/$(subst build-cli/,,$@)
endif

$(addprefix build-pwa/,$(pwas)): build-scss
	mkdir -p $(OUTPUT_DIR) $(BUILD_DIR)
	GOARCH=wasm GOOS=js go build -o $(STATIC_DIR)/app.wasm ./cmd/$(subst build-pwa/,,$@)
	go run ./cmd/$(subst build-pwa/,,$@) -dist $(BUILD_DIR) -prefix $(WWWPREFIX)
	cp -rf $(STATIC_DIR)/* $(BUILD_DIR)/web
	tar -cvzf $(OUTPUT_DIR)/$(subst build-pwa/,,$@).tar.gz -C $(BUILD_DIR) .

build-scss:
	npx sass -I . web/main.scss web/main.css

# Install
install: $(addprefix install-cli/,$(clis)) $(addprefix install-pwa/,$(pwas))

$(addprefix install-cli/,$(clis)):
	install -D -m 0755 $(OUTPUT_DIR)/$(subst install-cli/,,$@) $(DESTDIR)$(PREFIX)/bin/$(subst install-cli/,,$@)

$(addprefix install-pwa/,$(pwas)):
	mkdir -p $(DESTDIR)$(WWWROOT)$(WWWPREFIX)
	cp -rf $(BUILD_DIR)/* $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Uninstall
uninstall: $(addprefix uninstall-cli/,$(clis)) $(addprefix uninstall-pwa/,$(pwas))

$(addprefix uninstall-cli/,$(clis)):
	rm -f $(DESTDIR)$(PREFIX)/bin/$(subst uninstall-cli/,,$@)

$(addprefix uninstall-pwa/,$(pwas)):
	rm -rf $(DESTDIR)$(WWWROOT)$(WWWPREFIX)

# Run
run: $(addprefix run-cli/,$(clis)) $(addprefix run-pwa/,$(pwas))

$(addprefix run-cli/,$(clis)): build
	$(OUTPUT_DIR)/$(subst run-cli/,,$@) $(ARGS)

$(addprefix run-pwa/,$(pwas)): build-scss
	GOARCH=wasm GOOS=js go build -o $(STATIC_DIR)/app.wasm ./cmd/$(subst run-pwa/,,$@)
	go run ./cmd/$(subst run-pwa/,,$@) -serve

# Test
test:
	go test -timeout 3600s -parallel $(shell nproc) ./...

# Benchmark
benchmark:
	go test -timeout 3600s -bench=./... ./...

# Clean
clean:
	rm -rf $(OUTPUT_DIR) $(BUILD_DIR) $(STATIC_DIR)/app.wasm

# Dependencies
depend:
	npm i
	find node_modules/@patternfly/patternfly/ -name "*.css" -type f -delete
	rm -rf $(STATIC_DIR)/fonts
	mkdir -p $(STATIC_DIR)
	cp -r node_modules/@patternfly/patternfly/assets/fonts $(STATIC_DIR)

