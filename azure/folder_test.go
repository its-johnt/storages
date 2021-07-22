package azure

import (
	"github.com/stretchr/testify/assert"
	"github.com/its-johnt/storages/storage"
	"testing"
)

func TestAzureFolder(t *testing.T) {
	t.Skip("Credentials needed to run Azure Storage tests")

	storageFolder, err := ConfigureFolder("azure://test-container/test-folder/Sub0",
		make(map[string]string))

	assert.NoError(t, err)

	storage.RunFolderTest(storageFolder, t)
}
