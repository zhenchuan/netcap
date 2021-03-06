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

package encoder

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

var (
	// LiveMode switch for all encoders
	LiveMode bool

	// BlockSize is the file system block size
	BlockSize int

	allEncoderNames = make(map[string]struct{})
	errorMap        *AtomicCounterMap
)

func SetErrorMap(m *AtomicCounterMap) {
	errorMap = m
}

func invalidProto(name string) {
	fmt.Println("invalid protocol: " + ansi.Red + name + ansi.Reset)
	ShowEncoders()
	os.Exit(1)
}
