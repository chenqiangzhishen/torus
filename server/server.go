package server

import "github.com/barakmich/agro"
import "github.com/barakmich/agro/storage"

type server struct {
	cold     agro.KeyStore
	metadata agro.Metadata
	inodes   agro.INodeStore
}

func NewMemoryServer() agro.Server {
	md := agro.CreateMetadata("temp", "")
	return &server{
		cold:     storage.OpenTempKeyStore(),
		metadata: md,
		inodes:   storage.OpenTempINodeStore(),
	}
}