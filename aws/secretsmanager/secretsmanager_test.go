package secretsmanager_test

import (
	sm "github.com/Chewy-Inc/lets-go/aws/secretsmanager"
	"testing"
)

func TestGetSecret(t *testing.T) {
	secretValue := sm.GetSecret("someDummySecret")
	t.Log(secretValue)
}
