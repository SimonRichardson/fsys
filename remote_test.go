package fsys

import (
	"testing"
	"testing/quick"

	"github.com/pkg/errors"
)

func TestConfigBuild(t *testing.T) {
	t.Parallel()

	t.Run("build", func(t *testing.T) {
		fn := func(id, secret, token, region, bucket, key, encryption string) bool {
			config, err := BuildConfig(
				WithID(id),
				WithSecret(secret),
				WithToken(token),
				WithRegion(region),
				WithBucket(bucket),
				WithEncryption(true),
				WithKMSKey(key),
				WithServerSideEncryption(encryption),
			)
			if err != nil {
				t.Fatal(err)
			}
			return config.ID == id &&
				config.Secret == secret &&
				config.Token == token &&
				config.Region == region &&
				config.Bucket == bucket &&
				config.KMSKey == key &&
				config.ServerSideEncryption == encryption &&
				config.Type == KMSRemoteAccessType
		}

		if err := quick.Check(fn, nil); err != nil {
			t.Error(err)
		}
	})

	t.Run("invalid build", func(t *testing.T) {
		_, err := BuildConfig(
			func(config *RemoteConfig) error {
				return errors.Errorf("bad")
			},
		)

		if expected, actual := false, err == nil; expected != actual {
			t.Errorf("expected: %t, actual: %t", expected, actual)
		}
	})
}
