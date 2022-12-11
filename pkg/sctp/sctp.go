// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package sctp

import (
	"net"

	"github.com/wangxn2015/onos-lib-go/pkg/sctp/addressing"

	"github.com/wangxn2015/onos-lib-go/pkg/sctp/connection"

	"github.com/wangxn2015/onos-lib-go/pkg/sctp/types"
)

// DialOptions is SCTP options
type DialOptions struct {
	addressFamily types.AddressFamily
	mode          types.SocketMode
	initMsg       types.InitMsg
	nonblocking   bool
}

// DialOption dial option function
type DialOption func(*DialOptions)

// WithAddressFamily sets address family
func WithAddressFamily(family types.AddressFamily) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.addressFamily = family

	}
}

// WithMode sets SCTP mode
func WithMode(mode types.SocketMode) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.mode = mode

	}
}

// WithInitMsg sets options
func WithInitMsg(initMsg types.InitMsg) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.initMsg = initMsg

	}
}

// WithNonBlocking sets nonblocking
func WithNonBlocking(nonblocking bool) func(options *DialOptions) {
	return func(options *DialOptions) {
		options.nonblocking = nonblocking
	}
}

// DialSCTP creates a new SCTP connection
func DialSCTP(addr net.Addr, opts ...DialOption) (*connection.SCTPConn, error) {
	dialOptions := &DialOptions{}
	for _, option := range opts {
		option(dialOptions)
	}
	cfg := connection.NewConfig(
		connection.WithAddressFamily(dialOptions.addressFamily),
		connection.WithOptions(dialOptions.initMsg),
		connection.WithMode(dialOptions.mode),
		connection.WithNonBlocking(dialOptions.nonblocking))
	conn, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		return nil, err
	}
	/*
		//wxn--> bind client addr here----------------------
		addrArray := make([]net.IPAddr, 0)
		laddr := net.IPAddr{
			//these two ips can be used for testing. when ransim is deployed on RIC node(and outside k8s),use the RIC address 113;
			//11楼RIC地址
			//IP: net.ParseIP("192.168.127.113"),
			//11楼 baicells RAN地址是
			//IP: net.ParseIP("192.168.126.182"),
		}
		addrArray = append(addrArray, laddr)
		localSCTPAddr := addressing.Address{
			IPAddrs: addrArray,
			//Port:    0,
			AddressFamily: types.Sctp4,
		}
		err = conn.Bind(&localSCTPAddr)
		if err != nil {
			//fmt.Printf("wxn----> SCTP client binding error: %v",  err)
			return nil, err
		}
		//---------------------------------------------------
	*/
	sctpAddress := addr.(*addressing.Address)
	err = conn.Connect(sctpAddress)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

// DialSCTP creates a new SCTP connection
func DialSCTPWithSctpClientBindAddress(addr net.Addr, sctpClientAddr string, opts ...DialOption) (*connection.SCTPConn, error) {
	dialOptions := &DialOptions{}
	for _, option := range opts {
		option(dialOptions)
	}
	cfg := connection.NewConfig(
		connection.WithAddressFamily(dialOptions.addressFamily),
		connection.WithOptions(dialOptions.initMsg),
		connection.WithMode(dialOptions.mode),
		connection.WithNonBlocking(dialOptions.nonblocking))
	conn, err := connection.NewSCTPConnection(cfg)
	if err != nil {
		return nil, err
	}

	//wxn--> bind client addr here
	addrArray := make([]net.IPAddr, 0)
	laddr := net.IPAddr{
		IP: net.ParseIP(sctpClientAddr),
		//11楼RIC地址
		//IP: net.ParseIP("192.168.127.113"),
		//11楼 baicells RAN地址是
		//IP: net.ParseIP("192.168.126.182"),
	}
	addrArray = append(addrArray, laddr)
	localSCTPAddr := addressing.Address{
		IPAddrs: addrArray,
		//Port:    0,
		AddressFamily: types.Sctp4,
	}
	err = conn.Bind(&localSCTPAddr)
	if err != nil {
		//fmt.Printf("wxn----> SCTP client binding error: %v",  err)
		return nil, err
	}
	//----------------------------
	sctpAddress := addr.(*addressing.Address)
	err = conn.Connect(sctpAddress)
	if err != nil {
		return nil, err
	}

	return conn, nil

}
