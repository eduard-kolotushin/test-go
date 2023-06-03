package secman_test

import (
	"log"
	"strings"
	"testing"

	"github.com/eduard-kolotushin/test-go/secman"
)

func TestGetSecret(t *testing.T) {
	data, _ := secman.GetListSecrets("secret")
	if data == nil {
		t.Error("Lol, not true")
	}
	log.Print(data)
}

func getAllPath(paths []string) ([]string, []string) {
	if len(paths) == 0 {
		return nil, nil
	}
	var secrets, dirs []string
	for _, path := range paths {
		data, _ := secman.GetListSecrets(path)
		if data == nil {
			secrets = append(secrets, path)
			continue
		}
		for _, rpath := range data["keys"].([]interface{}) {
			dirs = append(dirs, strings.Trim(path, "/")+"/"+strings.Trim(rpath.(string), "/"))
		}
	}
	return secrets, dirs
}

func getAllPaths(paths []string) []string {
	if len(paths) == 0 {
		return nil
	}
	dirs := paths
	var secrets, result []string
	for dirs != nil {
		secrets, dirs = getAllPath(dirs)
		if len(secrets) != 0 {
			result = append(result, secrets...)
		}
	}
	return result
}

func TestGetAllPaths(t *testing.T) {
	data := getAllPaths([]string{"secret"})
	if data == nil {
		t.Error("Lol, not true")
	}
	log.Print(data)
}
