DIAGRAMS_SRC := $(shell find src -name "*.md")
DIAGRAMS_PNG := $(addsuffix .pdf, $(basename $(DIAGRAMS_SRC)))
PWD := $(shell pwd)

pdf: clean example $(DIAGRAMS_PNG)

%.pdf: %.md
	docker run --rm --init -v $(PWD):/home/marp/app/ -e LANG=EN marpteam/marp-cli $< -o ./build/$(notdir $@)

example: clean
	docker run --rm -it -v $(PWD):/data minlag/mermaid-cli -i /data/example/example.md -o /data/build/example.md -t dark -b white -s 3
.PHONY: example

clean:
	rm -rf ./build && mkdir ./build
.PHONY: clean