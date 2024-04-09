all: present run

present:
	GOBIN=$(PWD) go install golang.org/x/tools/cmd/present@latest

.PHONY: run
run:
	$(PWD)/present

.PHONY: clean
clean:
	rm -f ./present
