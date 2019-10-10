package main

import (
	_ "net"
)

// DNS resolve workaround for android in pure go

// this only work before any Lookup call and net.dnsReadConfig() failed
// go:linkname defaultNS _.defaultNS
var defaultNS = []string{"8.8.4.4:53", "1.1.1.1:53", "8.8.4.4:53"}
