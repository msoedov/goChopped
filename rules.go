package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Smart contract for project day

func main() {
	teams := 10
	var shared = []string{
		"bufio", // Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
		"fmt",
		"bytes",
		"strings",
		"strconv", // Package strconv implements conversions to and from string representations of basic data types.
		"fmt",
		"flag",
		"errors",
		"exec",   // Package exec runs external commands.
		"signal", // Package signal implements access to incoming signals.
		"time",   // Package time provides functionality for measuring and displaying time.
		"sync",   // Package sync provides basic synchronization primitives such as mutual exclusion locks.
		"os",     // Package os provides a platform-independent interface to operating system functionality.
		"log",    // Package log implements a simple logging package.
		"net",    // Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
	}

	var packages = []string{
		"heap",            // Package heap provides heap operations for any type that implements heap.Interface.
		"list",            // Package list implements a doubly linked list.
		"ring",            // Package ring implements operations on circular lists.
		"context",         // Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.
		"crypto",          // Package crypto collects common cryptographic constants.
		"aes",             // Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.
		"cipher",          // Package cipher implements standard block cipher modes that can be wrapped around low-level block cipher implementations.
		"hmac",            // Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards Publication 198.
		"md5",             // Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
		"rand",            // Package rand implements a cryptographically secure random number generator.
		"sha512",          // Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256 hash algorithms as defined in FIPS 180-4.
		"sql",             // Package sql provides a generic interface around SQL (or SQL-like) databases.
		"driver",          // Package driver defines interfaces to be implemented by database drivers as used by package sql.
		"elf",             // Package elf implements access to ELF object files.
		"csv",             // Package csv reads and writes comma-separated values (CSV) files.
		"json",            // Package json implements encoding and decoding of JSON as defined in RFC 7159.
		"xml",             // Package xml implements a simple XML 1.0 parser that understands XML name spaces.
		"importer",        // Package importer provides access to export data importers.
		"parser",          // Package parser implements a parser for Go source files.
		"scanner",         // Package scanner implements a scanner for Go source text.
		"token",           // Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates).
		"types",           // Package types declares the data types and implements the algorithms for type-checking of Go packages.
		"hash",            // Package hash provides interfaces for hash functions.
		"html",            // Package html provides functions for escaping and unescaping HTML text.
		"template",        // Package template (html/template) implements data-driven templates for generating HTML output safe against code injection.
		"image",           // Package image implements a basic 2-D image library.
		"color",           // Package color implements a basic color library.
		"palette",         // Package palette provides standard color palettes.
		"draw",            // Package draw provides image composition functions.
		"gif",             // Package gif implements a GIF image decoder and encoder.
		"jpeg",            // Package jpeg implements a JPEG image decoder and encoder.
		"png",             // Package png implements a PNG image decoder and encoder.
		"suffixarray",     // Package suffixarray implements substring search in logarithmic time using an in-memory suffix array.
		"io",              // Package io provides basic interfaces to I/O primitives.
		"ioutil",          // Package ioutil implements some I/O utility functions.
		"math",            // Package math provides basic constants and mathematical functions.
		"big",             // Package big implements arbitrary-precision arithmetic (big numbers).
		"bits",            // Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.
		"cmplx",           // Package cmplx provides basic constants and mathematical functions for complex numbers.
		"rand",            // Package rand implements pseudo-random number generators.
		"mime",            // Package mime implements parts of the MIME spec.
		"multipart",       // Package multipart implements MIME multipart parsing, as defined in RFC 2046.
		"quotedprintable", // Package quotedprintable implements quoted-printable encoding as specified by RFC 2045.
		"http",            // Package http provides HTTP client and server implementations.
		"cookiejar",       // Package cookiejar implements an in-memory RFC 6265-compliant http.CookieJar.
		"fcgi",            // Package fcgi implements the FastCGI protocol.
		"httptest",        // Package httptest provides utilities for HTTP testing.
		"httptrace",       // Package httptrace provides mechanisms to trace the events within HTTP client requests.
		"httputil",        // Package httputil provides HTTP utility functions, complementing the more common ones in the net/http package.
		"pprof",           // Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.
		"mail",            // Package mail implements parsing of mail messages.
		"rpc",             // Package rpc provides access to the exported methods of an object across a network or other I/O connection.
		"jsonrpc",         // Package jsonrpc implements a JSON-RPC 1.0 ClientCodec and ServerCodec for the rpc package.
		"smtp",            // Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321.
		"textproto",       // Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP.
		"url",             // Package url parses URLs and implements query escaping.
		"user",            // Package user allows user account lookups by name or id.
		"path",            // Package path implements utility routines for manipulating slash-separated paths.
		"filepath",        // Package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.
		"plugin",          // Package plugin implements loading and symbol resolution of Go plugins.
		"reflect",         // Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.
		"regexp",          // Package regexp implements regular expression search.
		"syntax",          // Package syntax parses regular expressions into parse trees and compiles parse trees into programs.
		"runtime",         // Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.
		"cgo",             // Package cgo contains runtime support for code generated by the cgo tool.
		"debug",           // Package debug contains facilities for programs to debug themselves while they are running.
		"pprof",           // Package pprof writes runtime profiling data in the format expected by the pprof visualization tool.
		"race",            // Package race implements data race detection logic.
		"trace",           // Package trace contains facilities for programs to generate traces for the Go execution tracer.
		"sort",            // Package sort provides primitives for sorting slices and user-defined collections.
		"atomic",          // Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.
		"syscall",         // Package syscall contains an interface to the low-level operating system primitives.
		"js",              // Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
		"testing",         // Package testing provides support for automated testing of Go packages.
		"iotest",          // Package iotest implements Readers and Writers useful mainly for testing.
		"quick",           // Package quick implements utility functions to help with black box testing.
		"scanner",         // Package scanner provides a scanner and tokenizer for UTF-8-encoded text.
		"tabwriter",       // Package tabwriter implements a write filter (tabwriter.Writer) that translates tabbed columns in input into properly aligned text.
		"template",        // Package template implements data-driven templates for generating textual output.
		"parse",           // Package parse builds parse trees for templates as defined by text/template and html/template.
		"unicode",         // Package unicode provides data and functions to test some properties of Unicode code points.
		"unsafe",          // Package unsafe contains operations that step around the type safety of Go programs.

	}
	rand.Seed(time.Now().Unix())

	groupSize := len(packages) / teams
	if groupSize >= 9 {
		groupSize = 9
	}
	randomize := rand.Perm(len(packages))

	fmt.Printf("Shared libraries: %#v\n", shared)

	for i := 0; i < teams; i++ {
		fmt.Printf("Team %d:  \n", i+1)
		for _, v := range randomize[i*groupSize : (i+1)*groupSize] {
			fmt.Printf("%s, ", packages[v])
		}
		fmt.Printf("\n\n---------------------------\n")
	}

	fmt.Printf("Penalties to total score\n")
	fmt.Printf("Extra package from standart lib %.2f\n", PenaltyWithCoef(0.88))
	fmt.Printf("Extra vendor package %.2f\n", PenaltyWithCoef(0.65))

	fmt.Printf("Ranking\n")

	fmt.Printf("Common:\n")
	fmt.Printf("Code style and structure 10 points max\n")
	fmt.Printf("\n---------------------------\n")
	fmt.Printf("By category:\n")
	fmt.Printf("No obvious and creative usage 15 points max * penalty \n")
	fmt.Printf("Practical application 15 points max * penalty\n")
	fmt.Printf("Most innovative  15 points max * penalty \n")
}

// PenaltyWithCoef sequence
func PenaltyWithCoef(coef float32) []float32 {
	penalty := []float32{coef}
	last := coef
	for i := 0; i < 5; i++ {
		last = coef * last
		penalty = append(penalty, last)

	}
	return penalty
}