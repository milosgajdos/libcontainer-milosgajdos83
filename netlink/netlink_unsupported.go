// +build !linux

package netlink

import (
	"errors"
	"net"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

func NetworkGetRoutes() ([]Route, error) {
	return nil, ErrNotImplemented
}

func NetworkLinkAdd(name string, linkType string) error {
	return ErrNotImplemented
}

func NetworkLinkDel(name string) error {
	return ErrNotImplemented
}

func NetworkLinkUp(iface *net.Interface) error {
	return ErrNotImplemented
}

func NetworkLinkAddIp(iface *net.Interface, ip net.IP, ipNet *net.IPNet) error {
	return ErrNotImplemented
}

func AddRoute(destination, source, gateway, device string) error {
	return ErrNotImplemented
}

func AddDefaultGw(ip, device string) error {
	return ErrNotImplemented
}

func NetworkSetMTU(iface *net.Interface, mtu int) error {
	return ErrNotImplemented
}

func NetworkSetMacAddress(iface *net.Interface, macaddr string) error {
	return ErrNotImplemented
}

func NetworkCreateVethPair(name1, name2 string) error {
	return ErrNotImplemented
}

func NetworkChangeName(iface *net.Interface, newName string) error {
	return ErrNotImplemented
}

func NetworkSetNsFd(iface *net.Interface, fd int) error {
	return ErrNotImplemented
}

func NetworkSetNsPid(iface *net.Interface, nspid int) error {
	return ErrNotImplemented
}

func NetworkSetMaster(iface, master *net.Interface) error {
	return ErrNotImplemented
}

func NetworkSetNoMaster(iface *net.Interface) error {
	return ErrNotImplemented
}

func NetworkLinkDown(iface *net.Interface) error {
	return ErrNotImplemented
}

func NetworkLinkAddMacVlan(masterDev, macVlanDev string, mode string) error {
	return ErrNotImplemented
}

func NetworkLinkAddVlan(masterDev, vlanDev string, vlanId uint16) error {
	return ErrNotImplemented
}
