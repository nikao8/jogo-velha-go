package domain_contracts_services

type IHash interface {
	Hash(value string) string
}
