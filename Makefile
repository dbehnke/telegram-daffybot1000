DAFFYBOTBIN=daffybot1000

all: bin/$(DAFFYBOTBIN)

.PHONY: gb
gb:
	go get github.com/constabulary/gb/... && which gb

bin/$(DAFFYBOTBIN): gb
	gb build

clean:
	rm -r -f bin
	rm -r -f pkg
