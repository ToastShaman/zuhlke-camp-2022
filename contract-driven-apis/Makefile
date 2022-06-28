SLIDES_SRC := $(shell find src -name "*.md")
SLIDES_PDF := $(addsuffix .pdf, $(basename $(SLIDES_SRC)))
PWD := $(shell pwd)

all: clean pdf
pdf: clean $(SLIDES_PDF)

%.pdf: %.md
	docker run --rm --init -v $(PWD):/home/marp/app/ -e LANG=EN marpteam/marp-cli $< --allow-local-files -o ./build/$(notdir $@)

clean:
	rm -rf ./build && mkdir ./build
.PHONY: clean