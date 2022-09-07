package sdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"reflect"
)

type TxEvent struct {
	BlockNumber uint64
	Contract    common.Address
	EventId     common.Hash
	Topic1      common.Hash
	Topic2      common.Hash
	Data        []byte
}

func genIntType(rule int64, size uint) []byte {
	var topic [common.HashLength]byte
	if rule < 0 {
		// if a rule is negative, we need to put it into two's complement.
		// extended to common.Hashlength bytes.
		topic = [common.HashLength]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
	}
	for i := uint(0); i < size; i++ {
		topic[common.HashLength-i-1] = byte(rule >> (i * 8))
	}
	return topic[:]
}

func MakeTopics(query ...[]interface{}) ([][]common.Hash, error) {
	topics := make([][]common.Hash, len(query))
	for i, filter := range query {
		for _, rule := range filter {
			var topic common.Hash

			// Try to generate the topic based on simple types
			switch rule := rule.(type) {
			case common.Hash:
				copy(topic[:], rule[:])
			case common.Address:
				copy(topic[common.HashLength-common.AddressLength:], rule[:])
			case *big.Int:
				blob := rule.Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case bool:
				if rule {
					topic[common.HashLength-1] = 1
				}
			case int8:
				copy(topic[:], genIntType(int64(rule), 1))
			case int16:
				copy(topic[:], genIntType(int64(rule), 2))
			case int32:
				copy(topic[:], genIntType(int64(rule), 4))
			case int64:
				copy(topic[:], genIntType(rule, 8))
			case uint8:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint16:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint32:
				blob := new(big.Int).SetUint64(uint64(rule)).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case uint64:
				blob := new(big.Int).SetUint64(rule).Bytes()
				copy(topic[common.HashLength-len(blob):], blob)
			case string:
				hash := crypto.Keccak256Hash([]byte(rule))
				copy(topic[:], hash[:])
			case []byte:
				hash := crypto.Keccak256Hash(rule)
				copy(topic[:], hash[:])

			default:
				// todo(rjl493456442) according solidity documentation, indexed event
				// parameters that are not value types i.e. arrays and structs are not
				// stored directly but instead a keccak256-hash of an encoding is stored.
				//
				// We only convert stringS and bytes to hash, still need to deal with
				// array(both fixed-size and dynamic-size) and struct.

				// Attempt to generate the topic from funky types
				val := reflect.ValueOf(rule)
				switch {
				// static byte array
				case val.Kind() == reflect.Array && reflect.TypeOf(rule).Elem().Kind() == reflect.Uint8:
					reflect.Copy(reflect.ValueOf(topic[:val.Len()]), val)
				default:
					return nil, fmt.Errorf("unsupported indexed type: %T", rule)
				}
			}
			topics[i] = append(topics[i], topic)
		}
	}
	return topics, nil
}
