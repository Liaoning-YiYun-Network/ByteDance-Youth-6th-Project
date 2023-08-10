package util

import (
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

// UUID 生成UUID
func UUID() (string, error) {
	rand.Seed(uint64(time.Now().Unix()))
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// UUIDWithoutHyphen 生成不带横杠的UUID
func UUIDWithoutHyphen() (string, error) {
	rand.Seed(uint64(time.Now().Unix()))
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	str := strings.ReplaceAll(u.String(), "-", "")
	return str, nil
}
