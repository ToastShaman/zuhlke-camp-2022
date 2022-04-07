package aws

import (
	"context"
	"crypto/sha256"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

type Signer interface {
	Sign(ctx context.Context, content []byte) (string, error)
	GetKeyId() string
}

type KmsSigner struct {
	KeyId  string
	Client *kms.Client
}

func (k *KmsSigner) Sign(ctx context.Context, content []byte) ([]byte, error) {
	hasher := sha256.New()
	hasher.Write(content)
	hash := hasher.Sum(nil)

	output, err := k.Client.Sign(ctx, &kms.SignInput{
		KeyId:            aws.String(k.KeyId),
		Message:          hash,
		SigningAlgorithm: types.SigningAlgorithmSpecEcdsaSha256,
		MessageType:      types.MessageTypeDigest,
	})

	if err != nil {
		return nil, err
	}

	return output.Signature, nil
}

func (k *KmsSigner) GetKeyId() string {
	return k.KeyId
}

func NewKMSClient(cfg aws.Config) *kms.Client {
	return kms.NewFromConfig(cfg)
}

func NewSigner(keyId string, client *kms.Client) *KmsSigner {
	return &KmsSigner{
		KeyId:  keyId,
		Client: client,
	}
}
