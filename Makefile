# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all install clean

all:
	gomake -C math3d32
	gomake -C math3d64

install: all
	gomake -C math3d32 install
	gomake -C math3d64 install

clean:
	gomake -C math3d32 clean
	gomake -C math3d64 clean
