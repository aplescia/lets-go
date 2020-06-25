package secretsmanager_test

import (
	sm "github.com/aplescia-chwy/lets-go/aws/secretsmanager"
	"testing"
)

func TestGetSecret(t *testing.T) {
	secretValue := sm.GetSecret("someDummySecret")
	t.Log(secretValue)
}
