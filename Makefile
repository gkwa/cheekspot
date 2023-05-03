SOURCES := $(shell find . -name '*.go')
TARGET := ./dist/cheekspot_darwin_amd64_v1/cheekspot

run: cheekspot
	./cheekspot

cheekspot: $(TARGET)
	cp $< $@

$(TARGET): $(SOURCES)
	gofumpt -w $<
	goreleaser build --single-target --snapshot --clean

.PHONY: clean
clean:
	rm -f cheekspot
	rm -f $(TARGET)
	rm -rf dist
