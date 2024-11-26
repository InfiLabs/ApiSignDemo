package utils

import (
	"bytes"
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

func UnPackUint16(r io.Reader) (uint16, error) {
	var n uint16
	err := binary.Read(r, binary.LittleEndian, &n)
	return n, err
}

func UnPackUint32(r io.Reader) (uint32, error) {
	var n uint32
	err := binary.Read(r, binary.LittleEndian, &n)
	return n, err
}

func UnPackUint64(r io.Reader) (uint64, error) {
	var n uint64
	err := binary.Read(r, binary.LittleEndian, &n)
	return n, err
}

func UnPackString(r io.Reader) (string, error) {
	n, err := UnPackUint16(r)
	if err != nil {
		return "", err
	}

	buf := make([]byte, n)
	if _, err = r.Read(buf); err != nil {
		return "", err
	}
	s := string(buf[:])
	return s, err
}

func UnPackContent(buff []byte) (string, uint64, string, uint64, string, error) {
	in := bytes.NewReader(buff)
	recordId, err := UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	issueAt, err := UnPackUint64(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	loginName, err := UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	validTime, err := UnPackUint64(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	m, err := UnPackString(in)
	if err != nil {
		return "", 0, "", 0, "", err
	}

	return recordId, issueAt, loginName, validTime, m, nil
}

func UnPackMessages(msgStr string) (string, uint64, map[uint16]uint32, error) {
	msgMap := make(map[uint16]uint32)

	msgByte := []byte(msgStr)
	in := bytes.NewReader(msgByte)

	signature, err := UnPackString(in)
	if err != nil {
		return "", 0, msgMap, err
	}
	salt, err := UnPackUint64(in)
	if err != nil {
		return "", 0, msgMap, err
	}

	length, err := UnPackUint16(in)
	if err != nil {
		return "", 0, msgMap, err
	}
	for i := uint16(0); i < length; i++ {
		key, err := UnPackUint16(in)
		if err != nil {
			return "", 0, msgMap, err
		}
		value, err := UnPackUint32(in)
		if err != nil {
			return "", 0, msgMap, err
		}
		msgMap[key] = value
	}

	return signature, salt, msgMap, nil
}
