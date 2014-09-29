.PHONY: all
# TODO: some of the hard-coded paths may be determined at runtime.

# Where libopentxs-core.so etc. live.
LIBPATH=/usr/local/lib/x86_64-linux-gnu

# Where "runtime.h" lives: adjust according to local Go runtime install.
GOINCLUDE=/usr/local/share/go/pkg/linux_amd64

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

	# This file confuses `go install', remove it.
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
