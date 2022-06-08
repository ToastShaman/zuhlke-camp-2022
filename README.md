# Zuhlke Backend Camp 2022

You have been asked to implement the next generation of APIs for ZTravel, a new mobile experience to make booking, boarding and accessing everything about your flights easier than ever before.

## Requirements

The client has given you the following requirements:

1. All APIs need to be hosted on AWS.

1. The APIs need to be reachable from the internet.

    **Note**: You can use [AWS Lambdas][6] or an alternative technology such as [AWS Fargate][7], [Amazon EKS][8], etc.

1. The setup and configuration of the environment(s) needs to be automated and repeatable (Infrastructure as Code IaC).

    **Note**: You can use [Terraform][1] but maybe try one of these alternatives: [AWS CDK][2], [Pulumi][3] or [CDK for Terraform][4].

    **Note**: Alternatively, if you want to skip the IaC feel free to use something like [Serverless Framework][11]

1. For security reasons only clients with a valid API key are allowed to call the APIs.

    **Note**: Let [AWS KMS][9] do the heavy lifting for you

    **Note**: If you are feeling adventurous, try to explore the [SafetyNet Attestation API][10] and equivalent for other platforms

1. For security reasons the HTTP responses from the APIs need to be signed with a Public/Private key pair.

    ```text
    SIGN("{REQUEST_ID}:{HTTP_METHOD}:{URL_PATH}:{CURRENT_DATE}:{BODY}")
    ```

    **Note**: Let [AWS KMS][5] do the heavy lifting for you

1. All API/Application logs need to be structured (aka JSON).

    ```json
    {"level":"debug","Name":"Tom","time":1562212768}
    ```

     **Note**: This is just an example and it's up to you to define an appropriate structure.

1. To help with 24/7 support we want to measure the performance of our APIs and create dashboards.

## Flights API

```http
GET /v1/flights
Accept: application/json
Authorization: Bearer <MOBILE_API_KEY>

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Signature-Date: Mon, 11 Apr 2022 15:44:04 GMT
Signature: keyId=b0a20181-03bf-4e41-8c7b-35d67b583f9e,signature=MEUCIQCXBA6rjjRi2xSPBCGfxR34vjrjCrYIDXxSbijyQHcmjwIgOLE501mmaWnfYfT1OW4jtFrDPUq261BtBZOIuqt/XAc=

{
    "flights": [
        {
            "id": "f6aa0876-405d-48a8-87c1-f62dc37c6e69",
            "from": "LHR"
            "to": "ZRH",
            "departure": "2022-04-11T15:00:00Z"
        },
        {
            "id": "98d35664-155d-4cb1-a3e1-e6d759d74c5b",
            "from": "ZRH"
            "to": "LHR",
            "departure": "2022-04-12T15:00:00Z"
        },
    ]
}
```

## Booking API

```http
POST /v1/bookings
Accept: application/json
Authorization: Bearer <MOBILE_API_KEY>
Content-Type: application/json

{
    "passengers": [
        {
            "title": "Mr",
            "first_name": "Odlaw",
            "last_name": "Waldo",
            "email_address": "where-is-waldo@now.com"
        }
    ],
    "flights": {
        "outbound": "f6aa0876-405d-48a8-87c1-f62dc37c6e69",
        "inbound": "98d35664-155d-4cb1-a3e1-e6d759d74c5b"
    }
}

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Signature-Date: Mon, 11 Apr 2022 15:44:04 GMT
Signature: keyId=b0a20181-03bf-4e41-8c7b-35d67b583f9e,signature=MEUCIQCXBA6rjjRi2xSPBCGfxR34vjrjCrYIDXxSbijyQHcmjwIgOLE501mmaWnfYfT1OW4jtFrDPUq261BtBZOIuqt/XAc=

{
    "booking_id": "da5541fb-7473-4843-ad88-524968fb1344",
    "status": "PAYMENT_PENDING"
    "total_price": {
        "amount": "500.00",
        "currency": "GBP"
    }
}
```

## Third Party Integration - Payment Processor

```http
PUT /v1/bookings/{booking_id}
Accept: application/json
Authorization: Bearer <PAYMENT_PROCESSOR_API_KEY>
Content-Type: application/json

{
    "booking_id": "da5541fb-7473-4843-ad88-524968fb1344",
    "status": "PAYMENT_SUCCESSFUL"
}

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Signature-Date: Mon, 11 Apr 2022 15:44:04 GMT
Signature: keyId=b0a20181-03bf-4e41-8c7b-35d67b583f9e,signature=MEUCIQCXBA6rjjRi2xSPBCGfxR34vjrjCrYIDXxSbijyQHcmjwIgOLE501mmaWnfYfT1OW4jtFrDPUq261BtBZOIuqt/XAc=

{
    "booking_id": "da5541fb-7473-4843-ad88-524968fb1344",
    "status": "BOOKED"
    "total_price": {
        "amount": "500.00",
        "currency": "GBP"
    }
}
```

[1]: https://www.terraform.io/
[2]: https://aws.amazon.com/cdk/
[3]: https://www.pulumi.com/
[4]: https://www.terraform.io/cdktf
[5]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
[6]: https://aws.amazon.com/lambda/
[7]: https://aws.amazon.com/fargate/
[8]: https://aws.amazon.com/eks/
[9]: https://docs.aws.amazon.com/kms/latest/developerguide/overview.html
[10]: https://developer.android.com/training/safetynet/attestation
[11]: https://www.serverless.com/