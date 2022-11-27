package validator

import (
	"chains-gotest-backend/utils"
)

var (
	ChainMetaValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType": {utils.NotEmpty()}}
	CoinMetaDataValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}}
	CrossChainAssetMappingValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}}
	CrossChainLockProxyValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}}
	CrossChainProxyHashMapValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType": {utils.NotEmpty()}, "Name": {utils.NotEmpty()}}
)
