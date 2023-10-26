package bchutil

import "github.com/btcsuite/btcd/chaincfg"

var MainnetDNSSeeds = []chaincfg.DNSSeed{
	{Host: "seed.bitcoinabc.org", HasFiltering: true},
	{Host: "seed-abc.bitcoinforks.org", HasFiltering: true},
	{Host: "seed.bitcoinunlimited.info", HasFiltering: true},
	{Host: "seed.bitprim.org", HasFiltering: true},
	{Host: "seed.deadalnix.me", HasFiltering: true},
}

var TestnetDNSSeeds = []chaincfg.DNSSeed{
	{Host: "testnet-seed.bitcoinabc.org", HasFiltering: true},
	{Host: "testnet-seed-abc.bitcoinforks.org", HasFiltering: true},
	{Host: "testnet-seed.bitcoinunlimited.info", HasFiltering: true},
	{Host: "testnet-seed.bitprim.org", HasFiltering: true},
	{Host: "testnet-seed.deadalnix.me", HasFiltering: true},
}

func GetDNSSeed(params *chaincfg.Params) []chaincfg.DNSSeed {
	if params.Name == chaincfg.MainNetParams.Name {
		return MainnetDNSSeeds
	}
	return TestnetDNSSeeds
}
