package libada

//go:generate stringer -type=Network
type Network byte

const (
	Testnet Network = 0
	Mainnet Network = 1
	Preprod Network = 2
	Preview Network = 3
)

func (n Network) Id() uint8 {
	return byte(n)
}

func (n Network) ProtocolMagic() uint32 {
	switch n {
	case Testnet:
		return 1097911063
	case Mainnet:
		return 764824073
	case Preprod:
		return 1
	case Preview:
		return 2
	default:
		return 0
	}
}
