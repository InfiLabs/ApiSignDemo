package pack

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"sort"
)

func PackUint16(w io.Writer, n uint16) error {
	return binary.Write(w, binary.LittleEndian, n)
}

func PackUint32(w io.Writer, n uint32) error {
	return binary.Write(w, binary.LittleEndian, n)
}

func PackUint64(w io.Writer, n uint64) error {
	return binary.Write(w, binary.LittleEndian, n)
}

func PackString(w io.Writer, s string) error {
	err := PackUint16(w, uint16(len(s)))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(s))
	return err
}

func PackHexString(w io.Writer, s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return PackString(w, string(b))
}

func PackExtra(w io.Writer, extra map[uint16]string) error {
	var keys []int
	if err := PackUint16(w, uint16(len(extra))); err != nil {
		return err
	}
	for k := range extra {
		keys = append(keys, int(k))
	}
	//should sorted keys
	sort.Ints(keys)

	for _, k := range keys {
		v := extra[uint16(k)]
		if err := PackUint16(w, uint16(k)); err != nil {
			return err
		}
		if err := PackString(w, v); err != nil {
			return err
		}
	}
	return nil
}

func PackMapUint32(w io.Writer, extra map[uint16]uint32) error {
	var keys []int
	if err := PackUint16(w, uint16(len(extra))); err != nil {
		return err
	}
	for k := range extra {
		keys = append(keys, int(k))
	}
	//should sorted keys
	sort.Ints(keys)

	for _, k := range keys {
		v := extra[uint16(k)]
		if err := PackUint16(w, uint16(k)); err != nil {
			return err
		}
		if err := PackUint32(w, v); err != nil {
			return err
		}
	}
	return nil
}
