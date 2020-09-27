.PHONY:
run: build
	docker run -it --rm  setitimer bash
# -v "$(PWD)":/src -w /src

.PHONY:
build:
	docker build -t setitimer .

