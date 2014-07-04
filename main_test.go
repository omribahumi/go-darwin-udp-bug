package main

import (
	"net"
	"syscall"
	"testing"
)

func TestSyscalls(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		t.Fatalf("Unable to create UDP socket: %s", err)
	}

	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_RECVDSTADDR, 1)
	if err != nil {
		t.Fatalf("Unable to setsockopt() on UDP socket: %s", err)
	}
}

func TestGoUdpListener(t *testing.T) {
	serverAddr, err := net.ResolveUDPAddr("udp", ":1025")
	if err != nil {
		t.Fatalf("Unable to ResolveUDPAddr: %s", err)
	}

	sConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		t.Fatalf("Unable to ListenUDP: %s", err)
	}

	file, err := sConn.File()
	if err != nil {
		t.Fatalf("Unable to query for socket file: %s", err)
	}

	err = syscall.SetsockoptInt(int(file.Fd()), syscall.IPPROTO_IP, syscall.IP_RECVDSTADDR, 1)
	if err != nil {
		t.Fatalf("Unable to setsockopt() on UDP socket: %s", err)
	}
}
