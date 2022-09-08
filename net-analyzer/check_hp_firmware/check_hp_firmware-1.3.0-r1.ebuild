EAPI=8
SRC_URI="https://github.com/NETWAYS/${PN}/archive/refs/tags/v${PV}.tar.gz"
SLOT="0"
RESTRICT="fetch"

pkg_nofetch() {
	echo "SRC_URI: ${SRC_URI}"
}

