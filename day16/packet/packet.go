package packet

import (
	"github.com/BooleanCat/go-functional/v2/it"
	"slices"
	"strconv"
	"strings"
)

type Header struct {
	version byte
	pType   byte
}

func (h Header) IsLiteral() bool {
	return h.pType == 4
}

type Packet interface {
	Header() Header
	Value() int
	VersionSum() int
}

type Literal struct {
	header Header
	value  int
}

func (l Literal) VersionSum() int {
	return int(l.header.version)
}

func (l Literal) Header() Header {
	return l.header
}

func (l Literal) Value() int {
	return l.value
}

func parseLiteral(header Header, payload []byte) (Literal, []byte) {
	digit := payload[0:5]
	bin := strings.Builder{}
	for {
		bin.WriteString(string(digit[1:]))
		payload = payload[5:]
		if digit[0] == '0' {
			break
		} else {
			digit = payload[0:5]
		}
	}
	value, _ := strconv.ParseInt(bin.String(), 2, 0)
	return Literal{header, int(value)}, payload
}

type Operator struct {
	header Header
	value  []Packet
}

func (o Operator) VersionSum() int {
	sum := int(o.header.version)
	for _, child := range o.value {
		sum += child.VersionSum()
	}
	return sum
}

func (o Operator) Header() Header {
	return o.header
}

func (o Operator) Value() int {
	children := it.Map(slices.Values(o.value), func(child Packet) int {
		return child.Value()
	})

	switch o.header.pType {
	case 0: // +
		return it.Fold(children, func(l int, r int) int { return l + r }, 0)
	case 1: // *
		return it.Fold(children, func(l int, r int) int { return l * r }, 1)
	case 2: // min
		v, _ := it.Min(children)
		return v
	case 3: // max
		v, _ := it.Max(children)
		return v
	case 5:
		args := slices.Collect(children)
		if args[0] > args[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		args := slices.Collect(children)
		if args[0] < args[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		args := slices.Collect(children)
		if args[0] == args[1] {
			return 1
		} else {
			return 0
		}
	default:
		panic("Type not found")
	}
}

func parseOperator(header Header, payload []byte) (Operator, []byte) {
	lenType := payload[0]
	if lenType == '0' {
		len, _ := strconv.ParseInt(string(payload[1:16]), 2, 16)
		payload = payload[16:]
		value, _ := ParseN(payload[:len], -1)

		return Operator{header, value}, payload[len:]
	} else {
		cnt, _ := strconv.ParseInt(string(payload[1:12]), 2, 16)
		payload = payload[12:]
		value, rest := ParseN(payload, int(cnt))

		return Operator{header, value}, rest
	}

}

func parseNext(input []byte) (Packet, []byte) {
	ver, _ := strconv.ParseInt(string(input[0:3]), 2, 8)
	pType, _ := strconv.ParseInt(string(input[3:6]), 2, 8)
	header := Header{byte(ver), byte(pType)}

	if header.IsLiteral() {
		return parseLiteral(header, input[6:])
	} else {
		return parseOperator(header, input[6:])
	}
}

func ParseN(input []byte, n int) ([]Packet, []byte) {
	cnt := 1
	packet, rest := parseNext(input)
	results := []Packet{packet}
	for {
		if len(rest) <= 10 || cnt == n {
			break
		}

		cnt++
		packet, rest = parseNext(rest)
		results = append(results, packet)
	}

	return results, rest
}
