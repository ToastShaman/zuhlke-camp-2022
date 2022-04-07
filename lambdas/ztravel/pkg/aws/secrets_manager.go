package aws

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretSupplier = func() string

type Secrets interface {
	GetSecretValue(ctx context.Context, key string) (string, error)
	GetSecretValueFn(ctx context.Context, key string) SecretSupplier
}

type SecretsManager struct {
	Client *secretsmanager.Client
}

func (s *SecretsManager) GetSecretValue(ctx context.Context, key string) (string, error) {
	value, err := s.Client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	})

	if err != nil {
		return "", err
	}

	return aws.ToString(value.SecretString), err
}

func (s *SecretsManager) GetSecretValueFn(ctx context.Context, key string) SecretSupplier {
	cache := make(map[string]string)

	return func() string {
		if val, found := cache[key]; found {
			return val
		}

		result, err := s.GetSecretValue(ctx, key)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to retrieve secret value")
		}

		cache[key] = result

		return result
	}
}

func NewSecretsManagerClient(cfg aws.Config) *secretsmanager.Client {
	return secretsmanager.NewFromConfig(cfg)
}

func NewSecretsManager(client *secretsmanager.Client) *SecretsManager {
	return &SecretsManager{
		Client: client,
	}
}
