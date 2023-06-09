package abi

type AbiGenOptions struct {
	AbiFile          string
	BinFile          string
	ToLanguage       string
	JsonFlag         string
	TypeFlag         string
	ExcFlag, PkgName string
	OutputName       string
	Alias            string
	SolFile          string
}
