### Chopped

Project Day/Night:

Gopher's take on Food Network's game show Chopped! If you haven't seen the show before, each contestant gets a basket of 8 ingredients and has to use them together in a recipe. For Golang cookery instead of 6 ingredients, every player/team gets 8 Go packages (plus a pantry of
standard library packages) to make an app or package with.

You can play either as a single player or on a team.

Examples:

By given bucket of "os", "net", "io", "crypto", “bufio”, “rand” Golang packages - you can create a weight based http load balancer - you
can create a load testing agent

By given bucket of "math", "context", "errors", "time", “ring”, “context” - you can create a primitive linear regression library

Winners:

Apps will be judged based on usefulness, creativity, completeness, and how well they showcase Go's strengths. Winners will be determined in the following categories.

1. The most creative use.
2. The most practical application.
3. The most innovative

Agenda:

1. Intro 15 minutes
2. Team registration 5 minutes
3. Ingredients distribution 5 minutes
4. Official start 90 minutes
5. Judging 15 minutes
6. Winner announcement 15 minutes
7. After party and networking ~

Ingredients instruction:

1. You can use one third party package and one extra ingredient without any penalty

2. You can use any additional number of third party/standard packages but you will receive a penalty in the final score

3. Ingredients will be chosen randomly from the list below. Group of ingredients will be unique by team.
4. Copy paste of code is not encouraged but allowed
5. You can user any of essential packages not listed in this list (such as bytes, strings, fmt, flag, errors):

https://golang.org/pkg/

Name Synopsis

bufio: Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that
also implements the interface but provides buffering and some help for textual I/O.

heap: Package heap provides heap operations for any type that implements heap.Interface.

list: Package list implements a doubly linked list.

ring: Package ring implements operations on circular lists.

context: Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API
boundaries and between processes.

crypto: Package crypto collects common cryptographic constants.

aes: Package aes implements AES encryption (formerly Rijndael), as defined in U.S. Federal Information Processing Standards Publication 197.

cipher: Package cipher implements standard block cipher modes that can be wrapped around low-level block cipher implementations.

hmac: Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards
Publication 198.

md5: Package md5 implements the MD5 hash algorithm as defined in RFC 1321.

rand: Package rand implements a cryptographically secure random number generator.

sha512: Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256 hash algorithms as defined in FIPS 180-4.

sql: Package sql provides a generic interface around SQL (or SQL-like) databases.

driver: Package driver defines interfaces to be implemented by database drivers as used by package sql.

elf: Package elf implements access to ELF object files.

csv: Package csv reads and writes comma-separated values (CSV) files.

json: Package json implements encoding and decoding of JSON as defined in RFC 7159.

xml: Package xml implements a simple XML 1.0 parser that understands XML name spaces.

importer: Package importer provides access to export data importers.

parser: Package parser implements a parser for Go source files.

scanner: Package scanner implements a scanner for Go source text.

token: Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens
(printing, predicates).
types: Package types declares the data types and implements the algorithms for type-checking of Go packages.

hash: Package hash provides interfaces for hash functions.

html: Package html provides functions for escaping and unescaping HTML text.

template: Package template (html/template) implements data-driven templates for generating HTML output safe against code injection.

image: Package image implements a basic 2-D image library.

color: Package color implements a basic color library.

palette: Package palette provides standard color palettes.

draw: Package draw provides image composition functions.

gif: Package gif implements a GIF image decoder and encoder.

jpeg: Package jpeg implements a JPEG image decoder and encoder.

png: Package png implements a PNG image decoder and encoder.

suffixarray: Package suffixarray implements substring search in logarithmic time using an in-memory suffix array.

io: Package io provides basic interfaces to I/O primitives.

ioutil: Package ioutil implements some I/O utility functions.

log: Package log implements a simple logging package.

math: Package math provides basic constants and mathematical functions.

big: Package big implements arbitrary-precision arithmetic (big numbers).

bits: Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.

cmplx: Package cmplx provides basic constants and mathematical functions for complex numbers.

rand: Package rand implements pseudo-random number generators.

mime: Package mime implements parts of the MIME spec.

multipart: Package multipart implements MIME multipart parsing, as defined in RFC 2046.

quotedprintable: Package quotedprintable implements quoted-printable encoding as specified by RFC 2045.

net: Package net provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.

http: Package http provides HTTP client and server implementations.

cookiejar: Package cookiejar implements an in-memory RFC 6265-compliant http.CookieJar.

fcgi: Package fcgi implements the FastCGI protocol.

httptest: Package httptest provides utilities for HTTP testing.

httptrace: Package httptrace provides mechanisms to trace the events within HTTP client requests.

httputil: Package httputil provides HTTP utility functions, complementing the more common ones in the net/http package.

pprof: Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

mail: Package mail implements parsing of mail messages.

rpc: Package rpc provides access to the exported methods of an object across a network or other I/O connection.

jsonrpc: Package jsonrpc implements a JSON-RPC 1.0 ClientCodec and ServerCodec for the rpc package.

smtp: Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321.

textproto: Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP.

url: Package url parses URLs and implements query escaping.

os: Package os provides a platform-independent interface to operating system functionality.

exec: Package exec runs external commands.

signal: Package signal implements access to incoming signals.

user: Package user allows user account lookups by name or id.

path: Package path implements utility routines for manipulating slash-separated paths.

filepath: Package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating
system-defined file paths.
plugin: Package plugin implements loading and symbol resolution of Go plugins.

reflect: Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.

regexp: Package regexp implements regular expression search.

syntax: Package syntax parses regular expressions into parse trees and compiles parse trees into programs.

runtime: Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines.

cgo: Package cgo contains runtime support for code generated by the cgo tool.

debug: Package debug contains facilities for programs to debug themselves while they are running.

pprof: Package pprof writes runtime profiling data in the format expected by the pprof visualization tool.

race: Package race implements data race detection logic.

trace: Package trace contains facilities for programs to generate traces for the Go execution tracer.

sort: Package sort provides primitives for sorting slices and user-defined collections.

strconv: Package strconv implements conversions to and from string representations of basic data types.

strings: Package strings implements simple functions to manipulate UTF-8 encoded strings.

sync: Package sync provides basic synchronization primitives such as mutual exclusion locks.

atomic: Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

syscall: Package syscall contains an interface to the low-level operating system primitives.

js: Package js gives access to the WebAssembly host environment when using the js/wasm architecture.

testing: Package testing provides support for automated testing of Go packages.

iotest: Package iotest implements Readers and Writers useful mainly for testing.

quick: Package quick implements utility functions to help with black box testing.

scanner: Package scanner provides a scanner and tokenizer for UTF-8-encoded text.

tabwriter: Package tabwriter implements a write filter (tabwriter.Writer) that translates tabbed columns in input into properly aligned
text.
template: Package template implements data-driven templates for generating textual output.

parse: Package parse builds parse trees for templates as defined by text/template and html/template.

time: Package time provides functionality for measuring and displaying time.

unicode: Package unicode provides data and functions to test some properties of Unicode code points.

unsafe: Package unsafe contains operations that step around the type safety of Go programs.
