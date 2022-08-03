package codec

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func MarshalProto(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func UnmarshalProto[T proto.Message](bytea []byte, msg T) (T, error) {
	err := proto.Unmarshal(bytea, msg)
	return msg, err
}

func Proto2Hex(msg proto.Message) (string, error) {
	bytea, err := proto.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal proto: %w", err)
	}

	return hex.EncodeToString(bytea), nil
}

func MustProto2Hex(msg proto.Message) string {
	got, err := Proto2Hex(msg)
	if err != nil {
		panic(err)
	}

	return got
}

func Hex2Proto[T proto.Message](input string, msg T) (T, error) {
	bytea, err := hex.DecodeString(input)
	if err != nil {
		return msg, fmt.Errorf("failed to decode hex: %w", err)
	}

	err = proto.Unmarshal(bytea, msg)
	if err != nil {
		return msg, fmt.Errorf("failed to unmarshall proto: %w", err)
	}

	return msg, nil
}

func MustHex2Proto[T proto.Message](input string, msg T) T {
	out, err := Hex2Proto[T](input, msg)
	if err != nil {
		panic(err)
	}

	return out
}
