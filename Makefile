# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: test-sdl

libs:
	make -C sdl install
	make -C ttf install


test-sdl: test-sdl.go libs
	$(GC) test-sdl.go
	$(LD) -o $@ test-sdl.$(O)

clean:
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-sdl shoot.png
