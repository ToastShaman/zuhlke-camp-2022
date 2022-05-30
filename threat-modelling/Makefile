SLIDES_SRC := $(shell find src -name "*.md")
SLIDES_PDF := $(addsuffix .pdf, $(basename $(SLIDES_SRC)))
PWD := $(shell pwd)

all: clean example pdf
pdf: clean $(SLIDES_PDF)

%.pdf: %.md
	docker run --rm --init -v $(PWD):/home/marp/app/ -e LANG=EN marpteam/marp-cli $< --allow-local-files -o ./build/$(notdir $@)

example:
	docker run --rm -it -v $(PWD):/data minlag/mermaid-cli -i /data/example/example.md -o /data/build/example.md -t dark -b transparent -s 3
.PHONY: example

clean:
	rm -rf ./build && mkdir ./build
.PHONY: clean