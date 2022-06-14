all: pair-programming threat-modelling
.PHONY: all

pair-programming:
	(cd ./pair-programming && make)
.PHONY: pair-programming

threat-modelling:
	(cd ./threat-modelling && make)
.PHONY: threat-modelling

tdd:
	(cd ./tdd-and-pair-programming && make)
.PHONY: tdd
