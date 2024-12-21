// Copyright 2024 slw <150657601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package storage defines redis storage.
package storage

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"

	"github.com/buger/jsonparser"
	uuid "github.com/satori/go.uuid"

	"github.com/shiqinfeng1/gomono-layout/third_party/forked/murmur3"
)

const defaultHashAlgorithm = "murmur64"

// GenerateToken generate token, if hashing algorithm is empty, use legacy key generation.
func GenerateToken(orgID, keyID, hashAlgorithm string) (string, error) {
	if keyID == "" {
		keyID = strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "")
	}

	if hashAlgorithm != "" {
		_, err := hashFunction(hashAlgorithm)
		if err != nil {
			hashAlgorithm = defaultHashAlgorithm
		}

		jsonToken := fmt.Sprintf(`{"org":"%s","id":"%s","h":"%s"}`, orgID, keyID, hashAlgorithm)

		return base64.StdEncoding.EncodeToString([]byte(jsonToken)), err
	}

	// Legacy keys
	return orgID + keyID, nil
}

// B64JSONPrefix stand for `{"` in base64.
const B64JSONPrefix = "ey"

// TokenHashAlgo ...
func TokenHashAlgo(token string) string {
	// Legacy tokens not b64 and not JSON records
	if strings.HasPrefix(token, B64JSONPrefix) {
		if jsonToken, err := base64.StdEncoding.DecodeString(token); err == nil {
			hashAlgo, _ := jsonparser.GetString(jsonToken, "h")

			return hashAlgo
		}
	}

	return ""
}

// TokenOrg ...
func TokenOrg(token string) string {
	if strings.HasPrefix(token, B64JSONPrefix) {
		if jsonToken, err := base64.StdEncoding.DecodeString(token); err == nil {
			// Checking error in case if it is a legacy tooken which just by accided has the same b64JSON prefix
			if org, err := jsonparser.GetString(jsonToken, "org"); err == nil {
				return org
			}
		}
	}

	// 24 is mongo bson id length
	if len(token) > 24 {
		return token[:24]
	}

	return ""
}

// Defines algorithm constant.
var (
	HashSha256    = "sha256"
	HashMurmur32  = "murmur32"
	HashMurmur64  = "murmur64"
	HashMurmur128 = "murmur128"
)

func hashFunction(algorithm string) (hash.Hash, error) {
	switch algorithm {
	case HashSha256:
		return sha256.New(), nil
	case HashMurmur64:
		return murmur3.New64(), nil
	case HashMurmur128:
		return murmur3.New128(), nil
	case "", HashMurmur32:
		return murmur3.New32(), nil
	default:
		return murmur3.New32(), fmt.Errorf("unknown key hash function: %s. Falling back to murmur32", algorithm)
	}
}

// HashStr return hash the give string and return.
func HashStr(in string) string {
	h, _ := hashFunction(TokenHashAlgo(in))
	_, _ = h.Write([]byte(in))

	return hex.EncodeToString(h.Sum(nil))
}

// HashKey return hash the give string and return.
func HashKey(in string) string {
	return HashStr(in)
}
