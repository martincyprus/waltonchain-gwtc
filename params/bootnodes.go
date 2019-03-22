// Copyright 2015 The go-wtc Authors
// This file is part of the go-wtc library.
//
// The go-wtc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-wtc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-wtc library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Wtc network.
var MainnetBootnodes = []string{
	"enode://7fe123c5eff54599b9c7bfd1b82200f0b8a0ae483f77b218466e502aca182a1cddf289c9dab2c43b27df70baea2ef0be5d899c81ba5ddc3ee78f839f014160f3@47.244.105.248:10101",
	"enode://103545ef8eca922dd475df4fd83202f7702d57fe28cdc19f3cf4999e4b3a9fcc3d5c32a030c66199e9d38b8edef9a8a74187b5b940c808e3e7aab6851279a40b@47.52.96.205:10101",
	"enode://4a9c6157762a49d4974a6a74959fc13ceb9b55f0895a363cf5535e7a155d3f51c9ae55f6803734d370ce9dc364e4107e3a09a279c02fbf31a52625d04e8d6fef@27.102.128.188:10101",
	"enode://5d2b1c62569daec741dab8d63bb49ba0582bf7e4f3834c7be7464754dd34b3ec4f37d459090315ab88ac7f40d894212c334b6344f9ebca5296819c68dc3173be@192.168.89.8:10101",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
