package ethkey

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Ethkey holds private key used in Ethereum wallets
type Ethkey struct {
	key *ecdsa.PrivateKey
}

// NewEthkey generates a new Ethereum private key object
func NewEthkey() Ethkey {
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}
	return Ethkey{key}
}

// Private returns private key string in hex format
func (kp Ethkey) Private() string {
	privateBytes := crypto.FromECDSA(kp.key)
	return hexutil.Encode(privateBytes)[2:]
}

// publicECDSA returns the public ECDSA from private key
func (kp Ethkey) publicECDSA() *ecdsa.PublicKey {
	publicKey := kp.key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot extract public key")
	}
	return publicKeyECDSA
}

// Public returns public key string in hex format
func (kp Ethkey) Public() string {
	publicKeyBytes := crypto.FromECDSAPub(kp.publicECDSA())
	return hexutil.Encode(publicKeyBytes)[4:]
}

// Address returns ethereum address in hex format
func (kp Ethkey) Address() string {
	publicECDSA := kp.publicECDSA()
	return crypto.PubkeyToAddress(*publicECDSA).Hex()
}

// PrintNewWallet creates a new wallet and prints out keys and Ethereum address
func PrintNewWallet() {
	kp := NewEthkey()
	fmt.Printf("────\nPriv: %s\nPubl: %s\nAddr: %s\n────\n", kp.Private(), kp.Public(), kp.Address())
}
