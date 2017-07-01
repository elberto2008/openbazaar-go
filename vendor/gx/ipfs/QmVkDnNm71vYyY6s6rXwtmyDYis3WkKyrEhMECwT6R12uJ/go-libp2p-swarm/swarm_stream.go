package swarm

import (
	"time"

	inet "gx/ipfs/QmRscs8KxrSmSv4iuevHv8JfuUzHBMoqiaHzxfDRiksd6e/go-libp2p-net"
	ps "gx/ipfs/QmYBtUjY2Q3jY8eYAva99G44zefLZ5Ksgk66VG81ueKQp1/go-peerstream"
	protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
)

// Stream is a wrapper around a ps.Stream that exposes a way to get
// our Conn and Swarm (instead of just the ps.Conn and ps.Swarm)
type Stream ps.Stream

// Stream returns the underlying peerstream.Stream
func (s *Stream) Stream() *ps.Stream {
	return (*ps.Stream)(s)
}

// Conn returns the Conn associated with this Stream, as an inet.Conn
func (s *Stream) Conn() inet.Conn {
	return s.SwarmConn()
}

// SwarmConn returns the Conn associated with this Stream, as a *Conn
func (s *Stream) SwarmConn() *Conn {
	return (*Conn)(s.Stream().Conn())
}

// Read reads bytes from a stream.
func (s *Stream) Read(p []byte) (n int, err error) {
	return s.Stream().Read(p)
}

// Write writes bytes to a stream, flushing for each call.
func (s *Stream) Write(p []byte) (n int, err error) {
	return s.Stream().Write(p)
}

// Close closes the stream, indicating this side is finished
// with the stream.
func (s *Stream) Close() error {
	return s.Stream().Close()
}

func (s *Stream) Protocol() protocol.ID {
	return (*ps.Stream)(s).Protocol()
}

func (s *Stream) SetProtocol(p protocol.ID) {
	(*ps.Stream)(s).SetProtocol(p)
}

func (s *Stream) SetDeadline(t time.Time) error {
	return s.Stream().SetDeadline(t)
}

func (s *Stream) SetReadDeadline(t time.Time) error {
	return s.Stream().SetReadDeadline(t)
}

func (s *Stream) SetWriteDeadline(t time.Time) error {
	return s.Stream().SetWriteDeadline(t)
}
