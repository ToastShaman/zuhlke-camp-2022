---
marp: true
title: Contract Driven APIs
paginate: true
theme: default
backgroundColor: rgb(255 255 255)
color: rgb(77 77 77)
style: |
  @import url('https://fonts.googleapis.com/css2?family=Merriweather+Sans:ital,wght@0,400;0,600;1,400&display=swap');
  section { 
    font-family: 'Merriweather Sans', sans-serif;
    font-weight: 400;
    font-size: 1.4em; 
  }
  section::after {
    font-size: 60%;
  }
  a { 
    color: rgb(152 91 156); 
  }
  strong { 
    color: rgb(0 153 204);
    font-weight: 600;
  }
  blockquote {
    position: absolute;
    bottom: 1em;
    left: 3em;
    font-size: 60%;
  }
  section.title-slide h1 {
    color: #F5F5F5;
  }
  section.title-slide p {
    color: #DCDCDC;
  }
  section.interval-slide {
    text-align: center;
    font-size: 250%;
    font-weight: bold;
  }
  section.interval-slide p {
    color: white;
  }
  section.small-font {
    font-size: 95%;
  }
  section.smaller-font {
    font-size: 80%;
  }
  section.smallest-font {
    font-size: 60%;
  }
  table {
    margin: 0px auto;
  }
  img {
    display: block;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
  }
---
<!-- _backgroundColor: #222222 -->
<!-- _class: title-slide -->
<!-- _paginate: false -->
<!-- _footer: June 2022 -->

# Contract Driven APIs

An Introduction by Kevin Denver

---

## What is Contract Testing?

Contract testing is a technique for testing software application interfaces and integrations.

Contracts are typically either "provider driven" (where the team responsible for the provider application maintains and shares the contract with its consumers) or "consumer driven" (where the teams responsible for the consumers maintain and share the contracts).

> https://pactflow.io/what-is-contract-testing-page/
---
<!-- _class: small-font -->

## What is consumer driven contract testing?

Consumer driven contract testing is a type of contract testing that ensures that a provider is compatible with the expectations that the consumer has of it.

### Explicit Contracts

For synchronous protocols such as HTTP, a contract would typically contain a collection of request/response pairs. To verify the contract, each request would be sent to the provider, and the response compared to the expected one. If the actual responses match the expected responses, then the contract has been successfully verified. Any mismatches highlight an incompatibility between the consumer and provider.

Use of an explicit contract is best when the provider is able to incorporate the contract verification step into its own build, providing fast feedback to ensure that no breaking changes are made. This is ideal as it means that breaking changes should not be able to make it into a production release of the application or library.

### Implicit contracts

When explicit contracts cannot be used, a similar outcome can be achieved by using a provider test double when running consumer tests, and executing a test harness at regular intervals to ensure that the real provider and the doubled provider behave the same way. The test harness enforces an implicit contract with the provider.

Use of a test harness to enforce an implicit contract is best when the contract verification process cannot be done during the provider's build - for example, for a public API or an OSS code library. While it won't stop breaking changes being released, it will ensure that they are highlighted as soon as possible.

> https://pactflow.io/what-is-consumer-driven-contract-testing/

---

![bg 60%](https://docs.pact.io/img/how-pact-works/summary.png)

---

## Benefits

![bg right 90%](https://s3-ap-southeast-2.amazonaws.com/content-prod-529546285894/2019/07/screenshot-16.png)

* The ability to develop the consumer before the API

* The ability to drive out the requirements for your provider first, meaning you implement exactly and only what you need in the provider.

* The ability to immediately see which consumers will be broken if a change is made to the provider API.

> https://pactflow.io/what-are-the-benefits-of-contract-testing/

---

## Tooling

* https://github.com/stoplightio/prism

* https://github.com/apiaryio/dredd/

* https://pact.io/

---

## References

* [Consumer-Driven Contracts: A Service Evolution Pattern][1]

[1]: https://martinfowler.com/articles/consumerDrivenContracts.html