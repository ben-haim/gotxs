.PHONY: all
# Some of the hard-coded paths could be determined at runtime

# this is where libopentxs-core.so etc live
LIBPATH=/usr/local/lib64

# this is where "runtime.h"
GOINCLUDE=/usr/lib/go/pkg/linux_amd64

LIBNAME=libopentxs-golang.so

all: $(LIBNAME) opentxs.a

opentxs_wrap.cxx opentxs_gc.c opentxs.go: opentxs.i
	swig -soname $(LIBNAME) -c++ -go -intgosize 64 opentxs.i

$(LIBNAME): opentxs_wrap.cxx
	g++ -g -fpic -shared -std=c++11 -DEXPORT= \
		-Iinclude/ \
		-L$(LIBPATH) \
		$< \
		-lopentxs-client \
		-lopentxs-ext \
		-lopentxs-cash \
		-lopentxs-basket \
		-lopentxs-core \
		-o $@

	# we need to remove this file because
	# it confuses `go install'
	rm $<

opentxs.a: opentxs_gc.c opentxs.go
	rm -f $@
	go tool 6c -I $(GOINCLUDE) -D _64BIT opentxs_gc.c
	go tool 6g opentxs.go
	go tool pack grc opentxs.a opentxs.6 opentxs_gc.6

clean:
	rm -f *.so
	rm -f *.dylib
	rm -f *.a
	rm -f *.6
	rm -f *.go
	rm -f *.c
	rm -f *.cxx

opentxs.txt:
	godoc . > $@

install: opentxs.a $(LIBNAME)
	sudo cp -v $(LIBNAME) $(LIBPATH)
	sudo ldconfig
