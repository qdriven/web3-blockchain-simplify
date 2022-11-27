package service

import (
	"errors"
	"go-chains/chains/models"
	"go-chains/global"
	"gorm.io/gorm"
)

func CreateChainMetaData(meta *models.ChainMetaData) (err error) {
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ? ",
		meta.ChainId, meta.ChainName).First(&models.ChainMetaData{})

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		return errors.New("duplicated chain Id existing ")
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllChainMetaData() []models.ChainMetaData {
	result := make([]models.ChainMetaData, 0)
	global.GVA_DB.Find(&result)
	return result
}

func UpdateChainMetaData(meta *models.ChainMetaData) error {
	result := &models.ChainMetaData{}
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ?",
		meta.ChainId, meta.ChainName).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}

func CreateCoinMetaData(meta *models.CoinMetaData) (err error) {
	return UpdateCoinMetaData(meta)
}

func UpdateCoinMetaData(meta *models.CoinMetaData) error {
	result := &models.CoinMetaData{}
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ?",
		meta.ChainId, meta.ChainName).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllCoinMetaData() []models.CoinMetaData {
	result := make([]models.CoinMetaData, 0)
	global.GVA_DB.Find(&result)
	return result
}

func CreateCrossChainLockProxy(meta *models.CrossChainLockProxy) (err error) {
	return UpdateCrossChainLockProxy(meta)
}

func UpdateCrossChainLockProxy(meta *models.CrossChainLockProxy) error {
	result := &models.CrossChainLockProxy{}
	queryResult := global.GVA_DB.Where("chain_id = ? AND chain_name = ?",
		meta.ChainId, meta.ChainName).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllCrossChainLockProxy() []models.CrossChainLockProxy {
	result := make([]models.CrossChainLockProxy, 0)
	global.GVA_DB.Find(&result)
	return result
}

func CreateCrossChainAssetMapping(meta *models.CrossChainAssetMapping) (err error) {
	return UpdateCrossChainAssetMapping(meta)
}

func UpdateCrossChainAssetMapping(meta *models.CrossChainAssetMapping) error {
	result := &models.CrossChainAssetMapping{}
	queryResult := global.GVA_DB.Where("from_asset_hash = ? AND lock_proxy_address = ?",
		meta.FromAssetHash, meta.LockProxyAddress).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllCrossChainAssetMapping() []models.CrossChainAssetMapping {
	result := make([]models.CrossChainAssetMapping, 0)
	global.GVA_DB.Find(&result)
	return result
}

func CreateCrossChainProxyHashMap(meta *models.CrossChainProxyHashMap) (err error) {
	return UpdateCrossChainProxyHashMap(meta)
}

func UpdateCrossChainProxyHashMap(meta *models.CrossChainProxyHashMap) error {
	result := &models.CrossChainProxyHashMap{}
	queryResult := global.GVA_DB.Where("to_chain_id = ? AND lock_proxy_address = ?",
		meta.ToChainId, meta.LockProxyAddress).First(result)

	if !errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		meta.ID = result.ID
		global.GVA_DB.Updates(meta)
	}
	return global.GVA_DB.Create(&meta).Error
}

func QueryAllCrossChainProxyHashMap() []models.CrossChainProxyHashMap {
	result := make([]models.CrossChainProxyHashMap, 0)
	global.GVA_DB.Find(&result)
	return result
}
