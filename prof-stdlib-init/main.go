/*
Command prof-stdlib-init is the profiling Go runtime stdib init functions CPU and memory.

This code written by @bradfitz, references is here:
 https://github.com/golang/go/issues/26775
 https://play.golang.org/p/9ervXCWzV_z

Usage

In your terminal,

    $ go build -o prof-stdlib-init .

    $ GODEBUG=memprofilerate=1 ./prof-stdlib-init
    408912

    $ go tool pprof /tmp/all.mem.prof
    Type: inuse_space
    Time: Oct 21, 2019 at 3:58am (JST)
    Entering interactive mode (type "help" for commands, "o" for options)

    (pprof) top 50
    Showing nodes accounting for 319.28kB, 87.98% of 362.90kB total
    Dropped 236 nodes (cum <= 1.81kB)
    Showing top 50 nodes out of 106
          flat  flat%   sum%        cum   cum%
      148.12kB 40.82% 40.82%   148.17kB 40.83%  runtime.procresize
       18.95kB  5.22% 46.04%    18.97kB  5.23%  encoding/xml.init
       14.28kB  3.94% 49.98%    14.28kB  3.94%  syscall.copyenv
       14.25kB  3.93% 53.90%    14.25kB  3.93%  runtime.malg
       12.17kB  3.35% 57.26%    12.17kB  3.35%  unicode.init
       11.31kB  3.12% 60.37%    13.62kB  3.75%  sync.(*Map).LoadOrStore
       11.02kB  3.04% 63.41%    11.02kB  3.04%  vendor/golang.org/x/net/http2/hpack.(*headerFieldTable).addEntry
        9.17kB  2.53% 65.94%     9.25kB  2.55%  html/template.init
        8.62kB  2.38% 68.31%    10.11kB  2.79%  net/http.init
           8kB  2.20% 70.52%       16kB  4.41%  runtime.allocm
        7.16kB  1.97% 72.49%     7.17kB  1.98%  debug/dwarf.init
           6kB  1.65% 74.14%        6kB  1.65%  syscall.runtime_envs
        5.25kB  1.45% 75.59%     5.25kB  1.45%  math/rand.NewSource
        4.03kB  1.11% 76.70%     4.03kB  1.11%  errors.New
        3.92kB  1.08% 77.78%     3.92kB  1.08%  go/types.init
        3.62kB     1% 78.78%    13.80kB  3.80%  encoding/gob.validUserType
        3.62kB     1% 79.78%     3.62kB     1%  regexp/syntax.init
        2.91kB   0.8% 80.58%     2.91kB   0.8%  go/types.(*Scope).Insert
        2.70kB  0.74% 81.32%     9.80kB  2.70%  encoding/gob.init
        2.25kB  0.62% 81.94%     2.25kB  0.62%  compress/flate.(*huffmanEncoder).generate
        2.09kB  0.58% 82.52%     2.09kB  0.58%  image/jpeg.(*huffmanLUT).init
        2.01kB  0.55% 83.07%     2.01kB  0.55%  sync.(*Pool).pinSlow
           2kB  0.55% 83.63%        5kB  1.38%  runtime.mcommoninit
        1.89kB  0.52% 84.15%     1.89kB  0.52%  compress/flate.newHuffmanEncoder
        1.75kB  0.48% 84.63%     2.80kB  0.77%  net.init
        1.64kB  0.45% 85.08%     2.16kB  0.59%  go/doc.init
        1.59kB  0.44% 85.52%     3.57kB  0.98%  encoding/gob.bootstrapType
        1.39kB  0.38% 85.90%     2.05kB  0.56%  go/build.init
        1.17kB  0.32% 86.23%     2.64kB  0.73%  text/template.init
        1.06kB  0.29% 86.52%    16.22kB  4.47%  encoding/gob.RegisterName
        1.02kB  0.28% 86.80%     4.29kB  1.18%  encoding/gob.buildTypeInfo
        0.88kB  0.24% 87.04%     3.27kB   0.9%  encoding/gob.newTypeObject
        0.88kB  0.24% 87.28%     3.80kB  1.05%  reflect.(*rtype).ptrTo
        0.75kB  0.21% 87.49%     1.84kB  0.51%  html/template.New
        0.73kB   0.2% 87.69%    18.38kB  5.06%  encoding/gob.init.1
        0.61kB  0.17% 87.86%     5.37kB  1.48%  net/http/pprof.init
        0.14kB 0.039% 87.90%     2.51kB  0.69%  go/types.defPredeclaredTypes
        0.09kB 0.026% 87.92%   191.63kB 52.81%  runtime.main
        0.08kB 0.022% 87.95%     2.40kB  0.66%  go/types.NewPackage
        0.05kB 0.013% 87.96%     6.05kB  1.67%  syscall.init
        0.05kB 0.013% 87.97%    11.16kB  3.07%  vendor/golang.org/x/net/http2/hpack.newStaticTable
        0.02kB 0.0043% 87.98%     2.24kB  0.62%  fmt.Sprintf
        0.02kB 0.0043% 87.98%     5.31kB  1.46%  math/rand.init
             0     0% 87.98%     2.55kB   0.7%  compress/flate.init.0
             0     0% 87.98%    16.22kB  4.47%  encoding/gob.Register
             0     0% 87.98%     3.27kB   0.9%  encoding/gob.getBaseType
             0     0% 87.98%     3.27kB   0.9%  encoding/gob.getType
             0     0% 87.98%     4.29kB  1.18%  encoding/gob.getTypeInfo
             0     0% 87.98%     3.80kB  1.05%  encoding/gob.implementsInterface
             0     0% 87.98%     4.80kB  1.32%  encoding/gob.mustGetTypeInfo

    (pprof) top --cum 50
    Showing nodes accounting for 277.38kB, 76.43% of 362.90kB total
    Dropped 236 nodes (cum <= 1.81kB)
    Showing top 50 nodes out of 106
          flat  flat%   sum%        cum   cum%
        0.09kB 0.026% 0.026%   191.63kB 52.81%  runtime.main
             0     0% 0.026%   189.45kB 52.20%  runtime.doInit
             0     0% 0.026%   148.55kB 40.93%  runtime.rt0_go
      148.12kB 40.82% 40.84%   148.17kB 40.83%  runtime.procresize
             0     0% 40.84%   148.17kB 40.83%  runtime.schedinit
       18.95kB  5.22% 46.07%    18.97kB  5.23%  encoding/xml.init
        0.73kB   0.2% 46.27%    18.38kB  5.06%  encoding/gob.init.1
             0     0% 46.27%    18.12kB  4.99%  runtime.mstart
             0     0% 46.27%    16.22kB  4.47%  encoding/gob.Register
        1.06kB  0.29% 46.56%    16.22kB  4.47%  encoding/gob.RegisterName
             0     0% 46.56%    16.22kB  4.47%  encoding/gob.registerBasics
           8kB  2.20% 48.77%       16kB  4.41%  runtime.allocm
             0     0% 48.77%    14.63kB  4.03%  os.init
             0     0% 48.77%    14.30kB  3.94%  os.Getwd
             0     0% 48.77%    14.28kB  3.94%  os.Getenv
             0     0% 48.77%    14.28kB  3.94%  sync.(*Once).Do
             0     0% 48.77%    14.28kB  3.94%  sync.(*Once).doSlow
             0     0% 48.77%    14.28kB  3.94%  syscall.Getenv
       14.28kB  3.94% 52.70%    14.28kB  3.94%  syscall.copyenv
       14.25kB  3.93% 56.63%    14.25kB  3.93%  runtime.malg
             0     0% 56.63%       14kB  3.86%  runtime.newm
             0     0% 56.63%    13.80kB  3.80%  encoding/gob.userType
        3.62kB     1% 57.63%    13.80kB  3.80%  encoding/gob.validUserType
       11.31kB  3.12% 60.74%    13.62kB  3.75%  sync.(*Map).LoadOrStore
       12.17kB  3.35% 64.10%    12.17kB  3.35%  unicode.init
             0     0% 64.10%    11.75kB  3.24%  runtime.systemstack
             0     0% 64.10%    11.22kB  3.09%  vendor/golang.org/x/net/http2/hpack.init
        0.05kB 0.013% 64.11%    11.16kB  3.07%  vendor/golang.org/x/net/http2/hpack.newStaticTable
       11.02kB  3.04% 67.15%    11.02kB  3.04%  vendor/golang.org/x/net/http2/hpack.(*headerFieldTable).addEntry
             0     0% 67.15%    10.12kB  2.79%  runtime.newproc.func1
             0     0% 67.15%    10.12kB  2.79%  runtime.newproc1
        8.62kB  2.38% 69.52%    10.11kB  2.79%  net/http.init
             0     0% 69.52%       10kB  2.76%  runtime.startm
        2.70kB  0.74% 70.27%     9.80kB  2.70%  encoding/gob.init
             0     0% 70.27%     9.28kB  2.56%  go/types.init.0
        9.17kB  2.53% 72.79%     9.25kB  2.55%  html/template.init
             0     0% 72.79%        8kB  2.20%  runtime.schedule
             0     0% 72.79%        8kB  2.20%  runtime.wakep
             0     0% 72.79%     7.70kB  2.12%  html/template.(*Template).Parse
             0     0% 72.79%     7.70kB  2.12%  text/template.(*Template).Parse
             0     0% 72.79%     7.30kB  2.01%  text/template/parse.Parse
        7.16kB  1.97% 74.77%     7.17kB  1.98%  debug/dwarf.init
             0     0% 74.77%     6.80kB  1.87%  text/template/parse.(*Tree).Parse
             0     0% 74.77%     6.78kB  1.87%  text/template/parse.(*Tree).parse
             0     0% 74.77%     6.50kB  1.79%  text/template/parse.(*Tree).textOrAction
             0     0% 74.77%     6.38kB  1.76%  runtime.mstart1
        0.05kB 0.013% 74.78%     6.05kB  1.67%  syscall.init
             0     0% 74.78%        6kB  1.65%  runtime.resetspinning
           6kB  1.65% 76.43%        6kB  1.65%  syscall.runtime_envs
             0     0% 76.43%     5.81kB  1.60%  text/template/parse.(*Tree).action
*/
package main

import (
	_ "archive/tar"
	_ "archive/zip"
	_ "bufio"
	_ "bytes"
	_ "compress/bzip2"
	_ "compress/flate"
	_ "compress/gzip"
	_ "compress/lzw"
	_ "compress/zlib"
	_ "container/heap"
	_ "container/list"
	_ "container/ring"
	_ "context"
	_ "crypto"
	_ "crypto/aes"
	_ "crypto/cipher"
	_ "crypto/des"
	_ "crypto/dsa"
	_ "crypto/ecdsa"
	_ "crypto/elliptic"
	_ "crypto/hmac"
	_ "crypto/md5"
	_ "crypto/rand"
	_ "crypto/rc4"
	_ "crypto/rsa"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "crypto/subtle"
	_ "crypto/tls"
	_ "crypto/x509"
	_ "crypto/x509/pkix"
	_ "database/sql"
	_ "database/sql/driver"
	_ "debug/dwarf"
	_ "debug/elf"
	_ "debug/gosym"
	_ "debug/macho"
	_ "debug/pe"
	_ "debug/plan9obj"
	_ "encoding"
	_ "encoding/ascii85"
	_ "encoding/asn1"
	_ "encoding/base32"
	_ "encoding/base64"
	_ "encoding/binary"
	_ "encoding/csv"
	_ "encoding/gob"
	_ "encoding/hex"
	_ "encoding/json"
	_ "encoding/pem"
	_ "encoding/xml"
	_ "errors"
	_ "expvar"
	_ "flag"
	_ "fmt"
	_ "go/ast"
	_ "go/build"
	_ "go/constant"
	_ "go/doc"
	_ "go/format"
	_ "go/importer"
	_ "go/parser"
	_ "go/printer"
	_ "go/scanner"
	_ "go/token"
	_ "go/types"
	_ "hash"
	_ "hash/adler32"
	_ "hash/crc32"
	_ "hash/crc64"
	_ "hash/fnv"
	_ "html"
	_ "html/template"
	_ "image"
	_ "image/color"
	_ "image/color/palette"
	_ "image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "index/suffixarray"
	_ "io"
	_ "io/ioutil"
	"log"
	_ "log/syslog"
	_ "math"
	_ "math/big"
	_ "math/bits"
	_ "math/cmplx"
	_ "math/rand"
	_ "mime"
	_ "mime/multipart"
	_ "mime/quotedprintable"
	_ "net"
	_ "net/http"
	_ "net/http/cgi"
	_ "net/http/cookiejar"
	_ "net/http/fcgi"
	_ "net/http/httptest"
	_ "net/http/httptrace"
	_ "net/http/httputil"
	_ "net/http/pprof"
	_ "net/mail"
	_ "net/rpc"
	_ "net/rpc/jsonrpc"
	_ "net/smtp"
	_ "net/textproto"
	_ "net/url"
	"os"
	_ "os/exec"
	_ "os/signal"
	_ "os/user"
	_ "path"
	_ "path/filepath"
	_ "plugin"
	_ "reflect"
	_ "regexp"
	_ "regexp/syntax"
	"runtime"
	_ "runtime/cgo"
	_ "runtime/debug"
	"runtime/pprof"
	_ "runtime/race"
	_ "runtime/trace"
	_ "sort"
	_ "strconv"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "testing"
	_ "testing/iotest"
	_ "testing/quick"
	_ "text/scanner"
	_ "text/tabwriter"
	_ "text/template"
	_ "text/template/parse"
	_ "time"
	_ "unicode"
	_ "unicode/utf16"
	_ "unicode/utf8"
	_ "unsafe"
)

func main() {
	runtime.GC()
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	println(ms.HeapAlloc)
	f, err := os.Create("/tmp/all.mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatalf("WriteHeapProfile: %v", err)
	}
}
