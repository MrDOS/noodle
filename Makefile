GOSRC=$(shell find . -type f -name '*.go')
VERSION=$(shell git describe --tags --dirty 2>/dev/null)

noodle: bin/noodle

bin/noodle: $(GOSRC) commands/description.txt
	go build \
		-o $@ \
		-ldflags "-X 'main.version=$(VERSION)'" \
		./...

commands/description.txt: README.md
	<$^ ./tools/para >$@
