package libada

//go:generate stringer -type=Network
type Network byte

const (
	Testnet Network = 0
	Mainnet Network = 1
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
	default:
		return 0
	}
}
