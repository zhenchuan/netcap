/*
 * NETCAP - Traffic Analysis Framework
 * Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package types

import (
	"encoding/hex"
	"strconv"
)

func (n NTP) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"LeapIndicator",
		"Version",
		"Mode",
		"Stratum",
		"Poll",
		"Precision",
		"RootDelay",
		"RootDispersion",
		"ReferenceID",
		"ReferenceTimestamp",
		"OriginTimestamp",
		"ReceiveTimestamp",
		"TransmitTimestamp",
		"ExtensionBytes",
	})
}

func (n NTP) CSVRecord() []string {
	return filter([]string{
		formatTimestamp(n.Timestamp),
		formatInt32(n.LeapIndicator),                         // int32
		formatInt32(n.Version),                               // int32
		formatInt32(n.Mode),                                  // int32
		formatInt32(n.Stratum),                               // int32
		formatInt32(n.Poll),                                  // int32
		formatInt32(n.Precision),                             // int32
		strconv.FormatUint(uint64(n.RootDelay), 10),          // uint32
		strconv.FormatUint(uint64(n.RootDispersion), 10),     // uint32
		strconv.FormatUint(uint64(n.ReferenceID), 10),        // uint32
		strconv.FormatUint(uint64(n.ReferenceTimestamp), 10), // uint64
		strconv.FormatUint(uint64(n.OriginTimestamp), 10),    // uint64
		strconv.FormatUint(uint64(n.ReceiveTimestamp), 10),   // uint64
		strconv.FormatUint(uint64(n.TransmitTimestamp), 10),  // uint64
		hex.EncodeToString(n.ExtensionBytes),                 // []byte
	})
}

func (n NTP) NetcapTimestamp() string {
	return n.Timestamp
}
