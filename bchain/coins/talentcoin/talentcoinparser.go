package talentcoin

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xe5cfcde1
	TestnetMagic wire.BitcoinNet = 0xe5cccce4
	RegtestMagic wire.BitcoinNet = 0xe5cfcce3
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{127}
	MainNetParams.ScriptHashAddrID = []byte{15}
	MainNetParams.Bech32HRPSegwit = "tlt"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{37}
	TestNetParams.ScriptHashAddrID = []byte{42}
	TestNetParams.Bech32HRPSegwit = "ttlt"
}

// TalentcoinParser handle
type TalentcoinParser struct {
	*btc.BitcoinParser
}

// NewTalentcoinParser returns new TalentcoinParser instance
func NewTalentcoinParser(params *chaincfg.Params, c *btc.Configuration) *TalentcoinParser {
	return &TalentcoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Vertcoin network,
// and the test Vertcoin network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
