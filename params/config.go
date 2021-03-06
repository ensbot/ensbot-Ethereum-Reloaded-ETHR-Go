// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/Ethereum-Reloaded/ETHR-Go/common"
)

// Genesis hashes to enforce below configs on.
var (
	
	MainnetGenesisHash = common.HexToHash("0x3550649845766f08bd2115aa6e8a37d5d8d5e9995b3489974234ff1f5a16d8f2")
	TestnetGenesisHash = common.HexToHash("0xc0aa9949ba05d4e30bff37bcdc4e5d9523a55a0fb34f7ad6abab1a4981ab5971")
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(22082018),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: nil,
		ForkMasternode:      big.NewInt(0),
		ForkSmartContract:   big.NewInt(0),
		EraV1Block:          big.NewInt(1000000),
		EraV2Block:          big.NewInt(3500000),
		EraV3Block:          big.NewInt(8000000),
		EraV4Block:          big.NewInt(10000000),
		EraV5Block:          big.NewInt(15000000),
		EraV6Block:          big.NewInt(18000000),
		EraV7Block:          big.NewInt(20000000),
		EraV8Block:          big.NewInt(20500000),
		EraV9Block:          big.NewInt(30000000),
		Ethash:              new(EthashConfig), 
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(29021982),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(1),
		EIP158Block:         big.NewInt(1),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: nil,
		ForkMasternode:      big.NewInt(0),
		ForkSmartContract:   big.NewInt(0),
		EraV1Block:          big.NewInt(1000000),
		EraV2Block:          big.NewInt(3500000),
		EraV3Block:          big.NewInt(8000000),
		EraV4Block:          big.NewInt(10000000),
		EraV5Block:          big.NewInt(15000000),
		EraV6Block:          big.NewInt(18000000),
		EraV7Block:          big.NewInt(20000000),
		EraV8Block:          big.NewInt(20500000),
		EraV9Block:          big.NewInt(30000000),
		Ethash:              new(EthashConfig),
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: nil,
		ForkMasternode:      nil,
		ForkSmartContract:   nil,
		EraV1Block:          big.NewInt(1000000),
		EraV2Block:          big.NewInt(3500000),
		EraV3Block:          big.NewInt(8000000),
		EraV4Block:          big.NewInt(10000000),
		EraV5Block:          big.NewInt(15000000),
		EraV6Block:          big.NewInt(18000000),
		EraV7Block:          big.NewInt(20000000),
		EraV8Block:          big.NewInt(20500000),
		EraV9Block:          big.NewInt(30000000),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), new(EthashConfig), nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, &CliqueConfig{Period: 0, Epoch: 30000}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), new(EthashConfig), nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already on byzantium)
	ForkMasternode         *big.Int `json:"ForkMasternode,omitempty"`         // Roller switch block (nil = no fork, 0 = already on Roller)
	ForkSmartContract     *big.Int `json:"forkSmartContract,omitempty"`     //second fork Roller release
	EraV1Block     *big.Int `json:"EraV1Block,omitempty"`
	EraV2Block     *big.Int `json:"EraV2Block,omitempty"`
	EraV3Block     *big.Int `json:"EraV31Block,omitempty"`
	EraV4Block     *big.Int `json:"EraV4Block,omitempty"`
	EraV5Block     *big.Int `json:"EraV5Block,omitempty"`
	EraV6Block     *big.Int `json:"EraV6Block,omitempty"`
	EraV7Block     *big.Int `json:"EraV7Block,omitempty"`
	EraV8Block     *big.Int `json:"EraV8Block,omitempty"`
	EraV9Block     *big.Int `json:"EraV9Block,omitempty"`
	
	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {

	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Fork Masternode: %v ForkSmartContract: %v EraV1Block: %v EraV2Block: %v EraV3Block: %v EraV4Block: %v EraV5Block: %v EraV6Block: %v EraV7Block: %v EraV8Block: %v EraV9Block: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.ForkMasternode,
		c.ForkSmartContract,
		c.EraV1Block,
		c.EraV2Block,
		c.EraV3Block,
		c.EraV4Block,
		c.EraV5Block,
		c.EraV6Block,
		c.EraV7Block,
		c.EraV8Block,
		c.EraV9Block,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Roller fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsRoller returns whether num is either equal to the Roller fork block or greater.
func (c *ChainConfig) IsForkMasternode(num *big.Int) bool {
	return isForked(c.ForkMasternode, num)
}

// IsForkSmartContract returns whether num is either equal to the IsForkSmartContract fork block or greater.
func (c *ChainConfig) IsForkSmartContract(num *big.Int) bool {
	return isForked(c.ForkSmartContract, num)
}

func (c *ChainConfig) IsEraV1(num *big.Int) bool {
	return isForked(c.EraV1Block, num)
}

func (c *ChainConfig) IsEraV2(num *big.Int) bool {
	return isForked(c.EraV2Block, num)
}

func (c *ChainConfig) IsEraV3(num *big.Int) bool {
	return isForked(c.EraV3Block, num)
}

func (c *ChainConfig) IsEraV4(num *big.Int) bool {
	return isForked(c.EraV4Block, num)
}

func (c *ChainConfig) IsEraV5(num *big.Int) bool {
	return isForked(c.EraV5Block, num)
}

func (c *ChainConfig) IsEraV6(num *big.Int) bool {
	return isForked(c.EraV6Block, num)
}

func (c *ChainConfig) IsEraV7(num *big.Int) bool {
	return isForked(c.EraV7Block, num)
}

func (c *ChainConfig) IsEraV8(num *big.Int) bool {
	return isForked(c.EraV8Block, num)
}

func (c *ChainConfig) IsEraV9(num *big.Int) bool {
	return isForked(c.EraV9Block, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.ForkMasternode, newcfg.ForkMasternode, head) {
		return newCompatError("Roller fork block Masternode", c.ForkMasternode, newcfg.ForkMasternode)
	}
	if isForkIncompatible(c.ForkSmartContract, newcfg.ForkSmartContract, head) {
		return newCompatError("Baneslayer fork block smartcontract", c.ForkSmartContract, newcfg.ForkSmartContract)
	}
	if isForkIncompatible(c.EraV1Block, newcfg.EraV1Block, head) {
		return newCompatError("EraV1 start block", c.EraV1Block, newcfg.EraV1Block)
	}
	if isForkIncompatible(c.EraV2Block, newcfg.EraV2Block, head) {
		return newCompatError("EraV2 start block", c.EraV2Block, newcfg.EraV2Block)
	}
	if isForkIncompatible(c.EraV3Block, newcfg.EraV3Block, head) {
		return newCompatError("EraV3 start block", c.EraV3Block, newcfg.EraV3Block)
	}
	if isForkIncompatible(c.EraV4Block, newcfg.EraV4Block, head) {
		return newCompatError("EraV4 start block", c.EraV4Block, newcfg.EraV4Block)
	}
	if isForkIncompatible(c.EraV5Block, newcfg.EraV5Block, head) {
		return newCompatError("EraV5 start block", c.EraV5Block, newcfg.EraV5Block)
	}
	if isForkIncompatible(c.EraV6Block, newcfg.EraV6Block, head) {
		return newCompatError("EraV6 start block", c.EraV6Block, newcfg.EraV6Block)
	}
	if isForkIncompatible(c.EraV7Block, newcfg.EraV7Block, head) {
		return newCompatError("EraV7 start block", c.EraV7Block, newcfg.EraV7Block)
	}
	if isForkIncompatible(c.EraV8Block, newcfg.EraV8Block, head) {
		return newCompatError("EraV8 start block", c.EraV8Block, newcfg.EraV8Block)
	}
	if isForkIncompatible(c.EraV9Block, newcfg.EraV9Block, head) {
		return newCompatError("EraV9 start block", c.EraV9Block, newcfg.EraV9Block)
	}
	
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                   *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158 bool
	IsByzantium                               bool
	IsConstantinople                          bool
	IsForkMasternode                          bool
	IsForkSmartContract                       bool
	IsEraV1                                   bool
	IsEraV2                                   bool
	IsEraV3                                   bool
	IsEraV4                                   bool
	IsEraV5                                   bool
	IsEraV6                                   bool
	IsEraV7                                   bool
	IsEraV8                                   bool
	IsEraV9                                   bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{ChainID: new(big.Int).Set(chainID), IsHomestead: c.IsHomestead(num), IsEIP150: c.IsEIP150(num), IsEIP155: c.IsEIP155(num), IsEIP158: c.IsEIP158(num), IsByzantium: c.IsByzantium(num), IsConstantinople: c.IsConstantinople(num), IsForkMasternode: c.IsForkMasternode(num), IsForkSmartContract: c.IsForkSmartContract(num) IsEraV1: c.IsEraV1(num) IsEraV2: c.IsEraV2(num) IsEraV3: c.IsEraV3(num) IsEraV4: c.IsEraV4(num) IsEraV5: c.IsEraV5(num) IsEraV6: c.IsEraV6(num) IsEraV7: c.IsEraV7(num) IsEraV8: c.IsEraV8(num) IsEraV9: c.IsEraV9(num)}
}
