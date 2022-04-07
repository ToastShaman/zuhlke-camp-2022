import { GetSecretValueCommand, SecretsManagerClient } from "@aws-sdk/client-secrets-manager";

export async function GetSecretValue(id: string): Promise<string> {
    const client = new SecretsManagerClient({ region: 'eu-west-2' });
    const result = await client.send(new GetSecretValueCommand({ SecretId: id }));
    return result.SecretString!;
}