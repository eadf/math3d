# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all install clean

all:
	gomake -C math3df
	gomake -C math3dd

install: all
	gomake -C math3df install
	gomake -C math3dd install

clean:
	gomake -C math3df clean
	gomake -C math3dd clean
