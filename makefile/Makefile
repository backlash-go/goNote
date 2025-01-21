
include ./make-rules/golang.mk


GO := go
SHELL := /bin/bash


x.txt: m.txt c.txt
	cat m.txt c.txt > x.txt

m.txt: a.txt b.txt
	cat a.txt b.txt > m.txt

.PHONY: clean
clean:
	rm -f m.txt
	rm -f x.txt

cd:
	@pwd
	cd ..
	pwd
cd_ok:
	pwd && cd .. ;

has_error:
	-rm zzz.txt
	echo 'ok'

test: g1
	echo #test

td:
	ls
	@echo #test

var-kept:
	@export foo=bar; \
	echo "foo=[$$foo]"

tes:
	# this is test
	@echo todo


.PHONY: clean-o
clean-o:
	rm -rf *.o

file1.o: a.txt
	ls

.PHONY: go.build.%
go.build.%:
ifneq ($(GO),go)
	echo "ifneq"
endif
	echo "endif"


tf1 tf2:
	touch $@

cpf: a.txt b.txt
	cp $< cpf.txt

print:
	$(eval GO_VERSION := $(shell $(GO) version))
	echo $(GO_VERSION)




