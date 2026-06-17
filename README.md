# dnsdist
A Go client for managing dnsdist over the encrypted control socket.

This library provides a typed API for dnsdist operations and a TCP transport that handles:

* NaCl secretbox encryption/decryption
* dnsdist nonce handshake
* command retries on connection/command failures
* connection lifecycle (connect, reconnect, close)


## Status
This project is currently partially implemented.

*Implemented features*
* Top-level client:
    * `NewClient(...)`
    * `Ping()`
    * `Close()`

* Rules API:
    * `Add(rule, action, opts...)`
    * `Remove(id)`
    * `List(opts...)`
    * `Exist(id)`
    * `Clear()`

* Rule builders:
    * `AllRule()`
    * `QNameRule(name)`
    * `QNameSuffixRule(names...)`
    * `QTypeRule(qtype)`
    * `NetmaskGroupRule(netmasks...)`
    * `AndRule(rules...)`
    * `OrRule(rules...)`
    * `NotRule(rule)`

* Action builders:
    * `SpoofAction(ips, opts...)`

* TCP transport
    * `NewTCPTransport(key, opts...)`
    * `WithHost(ip)`
    * `WithHostName(hostname)`
    * `WithPort(port)`
    * `WithTimeout(duration)`
    * `WithNumRetriesOnCommandFailure(n)`

*Unimplemented features*
These APIs exist but currently return an "un-implemented" error:

* Pools API:
    * `Pools().List()`

* Servers API:
    * `Servers().List()`
    * `Servers().Add(addr)`
    * `Servers().Remove(id)`
    * `Servers().SetUp(id)`
    * `Servers().SetDown(id)`

* Stats API:
    * `Stats().Version()`
    * `Stats().Dump()`
    * `Stats().Counters()`

## Installation
```bash
go get github.com/NorskHelsenett/dnsdist
```

## Requirements
* Go module support
* A running dnsdist instance with control socket enabled
* A valid base64-encoded 32-byte control socket key

## Quick Start!
```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/NorskHelsenett/dnsdist"
    "github.com/NorskHelsenett/dnsdist/transport/tcp"
)

func main() {
    // Example key placeholder. Replace with your real dnsdist control key.
    key := "REPLACE_WITH_BASE64_32_BYTE_KEY"

    t, err := tcp.NewTCPTransport(
        key,
        tcp.WithHost("127.0.0.1"),
        tcp.WithPort(5199),
        tcp.WithTimeout(5*time.Second),
        tcp.WithNumRetriesOnCommandFailure(2),
    )
    if err != nil {
        log.Fatalf("transport init failed: %v", err)
    }

    c := dnsdist.NewClient(t)
    defer c.Close()

    if err := c.Ping(); err != nil {
        log.Fatalf("dnsdist ping failed: %v", err)
    }

    fmt.Println("connected to dnsdist control socket")
}
```

## Usage example
### Spoof an FQDN
```go
    name := "awesome-name-for-rule"

    err := c.Rules().Add(
        rules.QNameRule("example.com"),
        rules.SpoofAction([]string{"127.0.0.1""::1"}),
        rules.GlobalRuleOptions{Name: &name},
    )
    if err != nil {
        log.Fatalf("add rule failed: %v", err)
    }
```

### Listing, check, and remove rules
```go
        // Basic list
    out, err := c.Rules().List()
    if err != nil {
        log.Fatalf("list rules failed: %v", err)
    }
    fmt.Println(out)

    // List with options
    showUUIDs := true
    width := 120
    out, err = c.Rules().List(&rules.ListOptions{
        ShowUUIDs:         &showUUIDs,
        TruncateRuleWidth: &width,
    })
    if err != nil {
        log.Fatalf("list rules with options failed: %v", err)
    }
    fmt.Println(out)

    // Check existence
    exists, err := c.Rules().Exist("awesome-name-for-rule")
    if err != nil {
        log.Fatalf("exist check failed: %v", err)
    }
    fmt.Printf("exists: %v\n", exists)

    // Remove one rule (id/name/uuid)
    if err := c.Rules().Remove("<name>/<uuid>/<id>"); err != nil {
        log.Fatalf("remove rule failed: %v", err)
    }

    // Remove all rules
    if err := c.Rules().Clear(); err != nil {
        log.Fatalf("clear rules failed: %v", err)
    }
```

*Rules Types*
Supported query types:
* QTypeA
* QTypeNS
* QTypeCNAME
* QTypeSOA
* QTypePTR
* QTypeTXT
* QTypeMX
* QTypeAAAA
* QTypeSRV
* QTypeANY

*Rule options*
* GlobalRuleOptions:
    * Name
    * UUID
* ListOptions:
    * ShowUUIDs
    * TruncateRuleWidth
* SpoofAction options:
    * AA
    * AD
    * RA
    * TTL

for more information about rule options, please refer to the official dnsdist controllsocket [docoumentation](https://www.dnsdist.org/reference/rules-management.html)

## Contributing
Contributions are welcome, especially for unimplemented Pools, Servers, and Stats features.
