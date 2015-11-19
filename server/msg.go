// Copyright (c) 2014 The SkyDNS Authors. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

package server

import "github.com/miekg/dns"

// Fit will make m fit the size. If a message is larger than size then entire
// additional section is dropped. If it is still to large and the transport
// is udp we return a truncated message.
// If the transport is tcp we return the message and set the returned bool
// to true, the server must then return SERVFAIL, because the generated answer
// is just too big.
func Fit(m *dns.Msg, size int, tcp bool) (*dns.Msg, bool) {
	if m.Len() > size {
		// TODO(miek): Check for OPT Records at the end and keep those.
		m.Extra = nil
	}
	if m.Len() < size {
		return m, false
	}

	// With TCP setting TC does not mean anything.
	if !tcp {
		m.Truncated = true
		return m, false
	}

	// Additional section is gone, still too big and can't set TC.
	return m, true
}
