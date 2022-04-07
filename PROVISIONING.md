# Provisioning

This project uses a feature from Visual Studio Code called [Visual Studio Code Remote - Containers][1].

You will need to run `npm install` once you've started the development container.

You will need to install [Node.js][2] version 16 if you decide not to use the development container.

## First time setup of CDK for your AWS account

See [Bootstrapping CDK][3] and [How to use the CDK CLI command][4].

```bash
export CDK_DEFAULT_ACCOUNT="<ACCOUNT_NUMBER>"
export CDK_DEFAULT_REGION="eu-west-2"

cdk bootstrap aws://<ACCOUNT_NUMBER>/eu-west-2
```

## Deploying all environments

```bash
cdk deploy --all --parameters owner=<YOUR_USERNAME>
```

See [infrastructure/README.md](./infrastructure/README.md) on how to use CDK.

[1]: https://code.visualstudio.com/docs/remote/containers
[2]: https://nodejs.dev/
[3]: https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html
[4]: https://docs.aws.amazon.com/cdk/v2/guide/cli.html
