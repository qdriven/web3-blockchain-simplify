package address

var whiteListAddresses = make([]WhiteListAddress, 0)

type WhiteListAddress struct {
	Address   string
	Enable    bool
	ChainName string
}

func IsNotInWhiteList(out, chainName string) bool {
	addresses := whiteListAddresses
	for _, addr := range addresses {
		if addr.Address == out && addr.Enable && addr.ChainName == chainName {
			return false
		}
	}
	return true
}

func AddAddressToWhiteList(address WhiteListAddress) {
	whiteListAddresses = append(whiteListAddresses, address)
}

func LoadWhiteList() {
	//TODO: load white list into memory
}
