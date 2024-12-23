package mysql

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var sampleAuth = []byte{
	0x49, 0x00, 0x00, 0x00, 0x0A, 0x38, 0x2E, 0x33, 0x2E, 0x30,
	0x00, 0xB5, 0x00, 0x00, 0x00, 0x6C, 0x04, 0x74, 0x45, 0x1E,
	0x3C, 0x79, 0x3B, 0x00, 0xFF, 0xFF, 0xFF, 0x02, 0x00, 0xFF,
	0xDF, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x59, 0x34, 0x1E, 0x44, 0x50, 0x2C, 0x07, 0x7D,
	0x0F, 0x0C, 0x14, 0x7E, 0x00, 0x63, 0x61, 0x63, 0x68, 0x69,
	0x6E, 0x67, 0x5F, 0x73, 0x68, 0x61, 0x32, 0x5F, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6F, 0x72, 0x64, 0x00,
}

func TestAuthPacket(t *testing.T) {
	reader := bytes.NewReader(sampleAuth)

	packet, err := NewAuthPacket(reader)
	require.NoError(t, err)

	require.Equal(t, 10, packet.ProtocolVersion)
	require.Equal(t, "8.3.0", packet.MySQLVersion)

	require.Equal(t, byte(0xff), packet.lowerCapabilities[0])
	require.Equal(t, byte(0xff), packet.lowerCapabilities[1])
}

func TestAuthPacket_RemoveSSLSupport(t *testing.T) {
	reader := bytes.NewReader(sampleAuth)

	packet, err := NewAuthPacket(reader)
	require.NoError(t, err)

	require.Equal(t, byte(0xff), packet.lowerCapabilities[0])
	require.Equal(t, byte(0xff), packet.lowerCapabilities[1])

	packet.RemoveSSLSupport()
	require.Equal(t, byte(0xf7), packet.lowerCapabilities[1])
}
