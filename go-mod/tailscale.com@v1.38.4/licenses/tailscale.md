# Tailscale CLI and daemon dependencies

The following open source dependencies are used to build the [tailscale][] and
[tailscaled][] commands. These are primarily used on Linux and BSD variants as
well as an [option for macOS][].

[tailscale]: https://pkg.go.dev/tailscale.com/cmd/tailscale
[tailscaled]: https://pkg.go.dev/tailscale.com/cmd/tailscaled
[option for macOS]: https://tailscale.com/kb/1065/macos-variants/

## Go Packages

Some packages may only be included on certain architectures or operating systems.


 - [filippo.io/edwards25519](https://pkg.go.dev/filippo.io/edwards25519) ([BSD-3-Clause](https://github.com/FiloSottile/edwards25519/blob/v1.0.0-rc.1/LICENSE))
 - [github.com/Microsoft/go-winio](https://pkg.go.dev/github.com/Microsoft/go-winio) ([MIT](https://github.com/Microsoft/go-winio/blob/v0.6.0/LICENSE))
 - [github.com/akutz/memconn](https://pkg.go.dev/github.com/akutz/memconn) ([Apache-2.0](https://github.com/akutz/memconn/blob/v0.1.0/LICENSE))
 - [github.com/alexbrainman/sspi](https://pkg.go.dev/github.com/alexbrainman/sspi) ([BSD-3-Clause](https://github.com/alexbrainman/sspi/blob/909beea2cc74/LICENSE))
 - [github.com/anmitsu/go-shlex](https://pkg.go.dev/github.com/anmitsu/go-shlex) ([MIT](https://github.com/anmitsu/go-shlex/blob/38f4b401e2be/LICENSE))
 - [github.com/aws/aws-sdk-go-v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/v1.17.3/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/config](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/config) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/config/v1.11.0/config/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/credentials](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/credentials) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/credentials/v1.6.4/credentials/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/feature/ec2/imds](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/ec2/imds) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/feature/ec2/imds/v1.8.2/feature/ec2/imds/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/internal/configsources](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/internal/configsources) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/internal/configsources/v1.1.27/internal/configsources/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/internal/endpoints/v2](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/internal/endpoints/v2.4.21/internal/endpoints/v2/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/internal/ini](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/internal/ini) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/internal/ini/v1.3.2/internal/ini/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/internal/sync/singleflight](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/internal/sync/singleflight) ([BSD-3-Clause](https://github.com/aws/aws-sdk-go-v2/blob/v1.17.3/internal/sync/singleflight/LICENSE))
 - [github.com/aws/aws-sdk-go-v2/service/internal/presigned-url](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/internal/presigned-url) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/service/internal/presigned-url/v1.5.2/service/internal/presigned-url/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/service/ssm](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/ssm) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/service/ssm/v1.35.0/service/ssm/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/service/sso](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sso) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/service/sso/v1.6.2/service/sso/LICENSE.txt))
 - [github.com/aws/aws-sdk-go-v2/service/sts](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts) ([Apache-2.0](https://github.com/aws/aws-sdk-go-v2/blob/service/sts/v1.11.1/service/sts/LICENSE.txt))
 - [github.com/aws/smithy-go](https://pkg.go.dev/github.com/aws/smithy-go) ([Apache-2.0](https://github.com/aws/smithy-go/blob/v1.13.5/LICENSE))
 - [github.com/aws/smithy-go/internal/sync/singleflight](https://pkg.go.dev/github.com/aws/smithy-go/internal/sync/singleflight) ([BSD-3-Clause](https://github.com/aws/smithy-go/blob/v1.13.5/internal/sync/singleflight/LICENSE))
 - [github.com/coreos/go-iptables/iptables](https://pkg.go.dev/github.com/coreos/go-iptables/iptables) ([Apache-2.0](https://github.com/coreos/go-iptables/blob/v0.6.0/LICENSE))
 - [github.com/creack/pty](https://pkg.go.dev/github.com/creack/pty) ([MIT](https://github.com/creack/pty/blob/v1.1.17/LICENSE))
 - [github.com/dblohm7/wingoes](https://pkg.go.dev/github.com/dblohm7/wingoes) ([BSD-3-Clause](https://github.com/dblohm7/wingoes/blob/6ac47ab19aa5/LICENSE))
 - [github.com/fxamacker/cbor/v2](https://pkg.go.dev/github.com/fxamacker/cbor/v2) ([MIT](https://github.com/fxamacker/cbor/blob/v2.4.0/LICENSE))
 - [github.com/go-ole/go-ole](https://pkg.go.dev/github.com/go-ole/go-ole) ([MIT](https://github.com/go-ole/go-ole/blob/v1.2.6/LICENSE))
 - [github.com/godbus/dbus/v5](https://pkg.go.dev/github.com/godbus/dbus/v5) ([BSD-2-Clause](https://github.com/godbus/dbus/blob/v5.0.6/LICENSE))
 - [github.com/golang/groupcache/lru](https://pkg.go.dev/github.com/golang/groupcache/lru) ([Apache-2.0](https://github.com/golang/groupcache/blob/41bb18bfe9da/LICENSE))
 - [github.com/google/btree](https://pkg.go.dev/github.com/google/btree) ([Apache-2.0](https://github.com/google/btree/blob/v1.0.1/LICENSE))
 - [github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid) ([BSD-3-Clause](https://github.com/google/uuid/blob/v1.3.0/LICENSE))
 - [github.com/hdevalence/ed25519consensus](https://pkg.go.dev/github.com/hdevalence/ed25519consensus) ([BSD-3-Clause](https://github.com/hdevalence/ed25519consensus/blob/c00d1f31bab3/LICENSE))
 - [github.com/illarion/gonotify](https://pkg.go.dev/github.com/illarion/gonotify) ([MIT](https://github.com/illarion/gonotify/blob/v1.0.1/LICENSE))
 - [github.com/insomniacslk/dhcp](https://pkg.go.dev/github.com/insomniacslk/dhcp) ([BSD-3-Clause](https://github.com/insomniacslk/dhcp/blob/de60144f33f8/LICENSE))
 - [github.com/jmespath/go-jmespath](https://pkg.go.dev/github.com/jmespath/go-jmespath) ([Apache-2.0](https://github.com/jmespath/go-jmespath/blob/v0.4.0/LICENSE))
 - [github.com/josharian/native](https://pkg.go.dev/github.com/josharian/native) ([MIT](https://github.com/josharian/native/blob/5c7d0dd6ab86/license))
 - [github.com/jsimonetti/rtnetlink](https://pkg.go.dev/github.com/jsimonetti/rtnetlink) ([MIT](https://github.com/jsimonetti/rtnetlink/blob/d380b505068b/LICENSE.md))
 - [github.com/kballard/go-shellquote](https://pkg.go.dev/github.com/kballard/go-shellquote) ([MIT](https://github.com/kballard/go-shellquote/blob/95032a82bc51/LICENSE))
 - [github.com/klauspost/compress](https://pkg.go.dev/github.com/klauspost/compress) ([Apache-2.0](https://github.com/klauspost/compress/blob/v1.15.4/LICENSE))
 - [github.com/klauspost/compress/internal/snapref](https://pkg.go.dev/github.com/klauspost/compress/internal/snapref) ([BSD-3-Clause](https://github.com/klauspost/compress/blob/v1.15.4/internal/snapref/LICENSE))
 - [github.com/klauspost/compress/zstd/internal/xxhash](https://pkg.go.dev/github.com/klauspost/compress/zstd/internal/xxhash) ([MIT](https://github.com/klauspost/compress/blob/v1.15.4/zstd/internal/xxhash/LICENSE.txt))
 - [github.com/kortschak/wol](https://pkg.go.dev/github.com/kortschak/wol) ([BSD-3-Clause](https://github.com/kortschak/wol/blob/da482cc4850a/LICENSE))
 - [github.com/kr/fs](https://pkg.go.dev/github.com/kr/fs) ([BSD-3-Clause](https://github.com/kr/fs/blob/v0.1.0/LICENSE))
 - [github.com/mattn/go-colorable](https://pkg.go.dev/github.com/mattn/go-colorable) ([MIT](https://github.com/mattn/go-colorable/blob/v0.1.12/LICENSE))
 - [github.com/mattn/go-isatty](https://pkg.go.dev/github.com/mattn/go-isatty) ([MIT](https://github.com/mattn/go-isatty/blob/v0.0.14/LICENSE))
 - [github.com/mdlayher/genetlink](https://pkg.go.dev/github.com/mdlayher/genetlink) ([MIT](https://github.com/mdlayher/genetlink/blob/v1.2.0/LICENSE.md))
 - [github.com/mdlayher/netlink](https://pkg.go.dev/github.com/mdlayher/netlink) ([MIT](https://github.com/mdlayher/netlink/blob/v1.7.1/LICENSE.md))
 - [github.com/mdlayher/sdnotify](https://pkg.go.dev/github.com/mdlayher/sdnotify) ([MIT](https://github.com/mdlayher/sdnotify/blob/v1.0.0/LICENSE.md))
 - [github.com/mdlayher/socket](https://pkg.go.dev/github.com/mdlayher/socket) ([MIT](https://github.com/mdlayher/socket/blob/v0.4.0/LICENSE.md))
 - [github.com/mitchellh/go-ps](https://pkg.go.dev/github.com/mitchellh/go-ps) ([MIT](https://github.com/mitchellh/go-ps/blob/v1.0.0/LICENSE.md))
 - [github.com/peterbourgon/ff/v3](https://pkg.go.dev/github.com/peterbourgon/ff/v3) ([Apache-2.0](https://github.com/peterbourgon/ff/blob/v3.1.2/LICENSE))
 - [github.com/pkg/errors](https://pkg.go.dev/github.com/pkg/errors) ([BSD-2-Clause](https://github.com/pkg/errors/blob/v0.9.1/LICENSE))
 - [github.com/pkg/sftp](https://pkg.go.dev/github.com/pkg/sftp) ([BSD-2-Clause](https://github.com/pkg/sftp/blob/v1.13.4/LICENSE))
 - [github.com/skip2/go-qrcode](https://pkg.go.dev/github.com/skip2/go-qrcode) ([MIT](https://github.com/skip2/go-qrcode/blob/da1b6568686e/LICENSE))
 - [github.com/tailscale/certstore](https://pkg.go.dev/github.com/tailscale/certstore) ([MIT](https://github.com/tailscale/certstore/blob/78d6e1c49d8d/LICENSE.md))
 - [github.com/tailscale/golang-x-crypto](https://pkg.go.dev/github.com/tailscale/golang-x-crypto) ([BSD-3-Clause](https://github.com/tailscale/golang-x-crypto/blob/bc99ab8c2d17/LICENSE))
 - [github.com/tailscale/netlink](https://pkg.go.dev/github.com/tailscale/netlink) ([Apache-2.0](https://github.com/tailscale/netlink/blob/cabfb018fe85/LICENSE))
 - [github.com/tailscale/wireguard-go](https://pkg.go.dev/github.com/tailscale/wireguard-go) ([MIT](https://github.com/tailscale/wireguard-go/blob/4fa124729667/LICENSE))
 - [github.com/tcnksm/go-httpstat](https://pkg.go.dev/github.com/tcnksm/go-httpstat) ([MIT](https://github.com/tcnksm/go-httpstat/blob/v0.2.0/LICENSE))
 - [github.com/toqueteos/webbrowser](https://pkg.go.dev/github.com/toqueteos/webbrowser) ([MIT](https://github.com/toqueteos/webbrowser/blob/v1.2.0/LICENSE.md))
 - [github.com/u-root/u-root/pkg/termios](https://pkg.go.dev/github.com/u-root/u-root/pkg/termios) ([BSD-3-Clause](https://github.com/u-root/u-root/blob/948a78c969ad/LICENSE))
 - [github.com/u-root/uio](https://pkg.go.dev/github.com/u-root/uio) ([BSD-3-Clause](https://github.com/u-root/uio/blob/c3537552635f/LICENSE))
 - [github.com/vishvananda/netlink/nl](https://pkg.go.dev/github.com/vishvananda/netlink/nl) ([Apache-2.0](https://github.com/vishvananda/netlink/blob/650dca95af54/LICENSE))
 - [github.com/vishvananda/netns](https://pkg.go.dev/github.com/vishvananda/netns) ([Apache-2.0](https://github.com/vishvananda/netns/blob/50045581ed74/LICENSE))
 - [github.com/x448/float16](https://pkg.go.dev/github.com/x448/float16) ([MIT](https://github.com/x448/float16/blob/v0.8.4/LICENSE))
 - [go4.org/mem](https://pkg.go.dev/go4.org/mem) ([Apache-2.0](https://github.com/go4org/mem/blob/927187094b94/LICENSE))
 - [go4.org/netipx](https://pkg.go.dev/go4.org/netipx) ([BSD-3-Clause](https://github.com/go4org/netipx/blob/7e7bdc8411bf/LICENSE))
 - [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) ([BSD-3-Clause](https://cs.opensource.google/go/x/crypto/+/v0.6.0:LICENSE))
 - [golang.org/x/exp](https://pkg.go.dev/golang.org/x/exp) ([BSD-3-Clause](https://cs.opensource.google/go/x/exp/+/47842c84:LICENSE))
 - [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) ([BSD-3-Clause](https://cs.opensource.google/go/x/net/+/v0.7.0:LICENSE))
 - [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) ([BSD-3-Clause](https://cs.opensource.google/go/x/sync/+/v0.1.0:LICENSE))
 - [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) ([BSD-3-Clause](https://cs.opensource.google/go/x/sys/+/v0.5.0:LICENSE))
 - [golang.org/x/term](https://pkg.go.dev/golang.org/x/term) ([BSD-3-Clause](https://cs.opensource.google/go/x/term/+/v0.5.0:LICENSE))
 - [golang.org/x/text](https://pkg.go.dev/golang.org/x/text) ([BSD-3-Clause](https://cs.opensource.google/go/x/text/+/v0.7.0:LICENSE))
 - [golang.org/x/time/rate](https://pkg.go.dev/golang.org/x/time/rate) ([BSD-3-Clause](https://cs.opensource.google/go/x/time/+/579cf78f:LICENSE))
 - [golang.zx2c4.com/wintun](https://pkg.go.dev/golang.zx2c4.com/wintun) ([MIT](https://git.zx2c4.com/wintun-go/tree/LICENSE?id=0fa3db229ce2))
 - [golang.zx2c4.com/wireguard/windows/tunnel/winipcfg](https://pkg.go.dev/golang.zx2c4.com/wireguard/windows/tunnel/winipcfg) ([MIT](https://git.zx2c4.com/wireguard-windows/tree/COPYING?h=v0.5.3))
 - [gopkg.in/yaml.v2](https://pkg.go.dev/gopkg.in/yaml.v2) ([Apache-2.0](https://github.com/go-yaml/yaml/blob/v2.4.0/LICENSE))
 - [gvisor.dev/gvisor/pkg](https://pkg.go.dev/gvisor.dev/gvisor/pkg) ([Apache-2.0](https://github.com/google/gvisor/blob/703fd9b7fbc0/LICENSE))
 - [inet.af/peercred](https://pkg.go.dev/inet.af/peercred) ([BSD-3-Clause](https://github.com/inetaf/peercred/blob/0893ea02156a/LICENSE))
 - [inet.af/wf](https://pkg.go.dev/inet.af/wf) ([BSD-3-Clause](https://github.com/inetaf/wf/blob/50d96caab2f6/LICENSE))
 - [k8s.io/client-go/util/homedir](https://pkg.go.dev/k8s.io/client-go/util/homedir) ([Apache-2.0](https://github.com/kubernetes/client-go/blob/v0.25.0/LICENSE))
 - [nhooyr.io/websocket](https://pkg.go.dev/nhooyr.io/websocket) ([MIT](https://github.com/nhooyr/websocket/blob/v1.8.7/LICENSE.txt))
 - [sigs.k8s.io/yaml](https://pkg.go.dev/sigs.k8s.io/yaml) ([MIT](https://github.com/kubernetes-sigs/yaml/blob/v1.3.0/LICENSE))
 - [software.sslmate.com/src/go-pkcs12](https://pkg.go.dev/software.sslmate.com/src/go-pkcs12) ([BSD-3-Clause](https://github.com/SSLMate/go-pkcs12/blob/v0.2.0/LICENSE))
 - [tailscale.com](https://pkg.go.dev/tailscale.com) ([BSD-3-Clause](https://github.com/tailscale/tailscale/blob/HEAD/LICENSE))
 - [tailscale.com/tempfork/gliderlabs/ssh](https://pkg.go.dev/tailscale.com/tempfork/gliderlabs/ssh) ([BSD-3-Clause](https://github.com/tailscale/tailscale/blob/HEAD/tempfork/gliderlabs/ssh/LICENSE))