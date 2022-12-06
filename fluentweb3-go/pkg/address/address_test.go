package address

import (
	"fmt"
	"github.com/smartystreets/assertions"
	"testing"
)

func TestNewWrapperAddress(t *testing.T) {

	wrappedAddress := NewWrapperAddress()
	assertions.ShouldNotBeNil(wrappedAddress.Address)
	assertions.ShouldNotBeNil(wrappedAddress.AddressStr)
	assertions.ShouldNotBeNil(wrappedAddress.PrivKeyStr)
	assertions.ShouldNotBeNil(wrappedAddress.PubKeyStr)
	fmt.Println("address is ", wrappedAddress.AddressStr)
	fmt.Println("priKey is ", wrappedAddress.PrivKeyStr)
}
