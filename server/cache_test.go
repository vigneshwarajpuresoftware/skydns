// Copyright (c) 2014 The SkyDNS Authors. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

package server

import (
	"testing"

	"github.com/miekg/dns"
	"github.com/skynetservices/skydns/cache"
)

func TestCacheTruncated(t *testing.T) {
	s := newTestServer(t, true)
	m := &dns.Msg{}
	m.SetQuestion("skydns.test.", dns.TypeSRV)
	m.Truncated = true
	s.rcache.InsertMessage(cache.Key(m.Question[0], false, false), m)

	// Now asking for this should result in a non-truncated answer.
	resp, _ := dns.Exchange(m, "127.0.0.1:"+StrPort)
	if resp.Truncated {
		t.Fatal("truncated bit should be false")
	}
}

// Store a large message in the cache, then query with a smaller bufsize and check
// we get back a smaller message.
// TODO(miek).
/*
func testCacheStoreLarge(t *testing.T) {
	s := newTestServer(t, true)
	defer s.Stop()

	c := new(dns.Client)
	m := new(dns.Msg)

	for i := 0; i < 2000; i++ {
		is := strconv.Itoa(i)
		m := &msg.Service{
			Host: "2001::" + is, Key: "machine" + is + ".machines.skydns.test.",
		}
		addService(t, s, m.Key, 0, m)
		defer delService(t, s, m.Key)
	}
	m.SetQuestion("machines.skydns.test.", dns.TypeSRV)
	resp, _, err := c.Exchange(m, "127.0.0.1:"+StrPort)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", resp)

	if resp.Rcode != dns.RcodeSuccess {
		t.Fatalf("expecting server failure, got %d", resp.Rcode)
	}
}
*/
