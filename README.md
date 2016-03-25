Zeroless
========

[![Build Status](https://travis-ci.org/zmqless/go-zeroless.svg?branch=master)](https://travis-ci.org/zmqless/go-zeroless)
[![Coverage Status](https://coveralls.io/repos/zmqless/go-zeroless/badge.svg?branch=master&service=github)](https://coveralls.io/github/zmqless/go-zeroless?branch=master)
[![GoDoc](https://godoc.org/github.com/zmqless/go-zeroless?status.svg)](https://godoc.org/github.com/zmqless/go-zeroless)
[![License](https://img.shields.io/badge/license-LGPLv2+-blue.svg)](https://www.gnu.org/licenses/lgpl-2.1.html)

Yet another [ØMQ] wrapper for Go. However, differing from [zmq4], which
tries to stay very close to the C++ implementation, this project aims to
make distributed systems employing [ØMQ] as gopher as possible.

Being simpler to use, Zeroless doesn't supports all of the fine aspects
and features of [ØMQ]. However, you can expect to find all the message
passing patterns you were accustomed to (i.e. pair, request/reply,
publisher/subscriber, push/pull). Despite that, the only transport
available is TCP.

Installation
------------

Use `go get`:

    $ go get github.com/zmqless/go-zeroless

Go API
------

In the `zeroless` package, two structs can be used to define how distributed
entities are related (i.e. ``Server`` and ``Client``). To put it bluntly, with
the exception of the pair pattern, a client may be connected to multiple
servers, while a server may accept incoming connections from multiple clients.

Both servers and clients are able to create a *channel*. So that you can iterate
over incoming messages and/or transmit a message.

Testing
-------

To run all the tests:

    $ go test

License
-------

Copyright 2015 Lucas Lira Gomes x8lucas8x@gmail.com

This library is free software; you can redistribute it and/or modify it
under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 2.1 of the License, or (at
your option) any later version.

This library is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this library. If not, see http://www.gnu.org/licenses/.

[ØMQ]: http://www.zeromq.org
[zmq4]: https://github.com/pebbe/zmq4
