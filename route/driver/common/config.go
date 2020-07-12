package common

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	MagicCookieKey: "DriverKey",
	MagicCookieValue: "DriverValue",
}