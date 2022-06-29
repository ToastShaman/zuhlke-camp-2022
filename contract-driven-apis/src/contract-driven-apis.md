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
    font-size: 1.2em;
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

![bg 90%](https://martinfowler.com/articles/practical-test-pyramid/testPyramid.png)

---

![bg right 60%](https://martinfowler.com/articles/practical-test-pyramid/httpIntegrationTest.png)

## Disadvantages of End-To-End Test

* Slow
* Brittle / Flaky
* Introduces coupling between teams
* Increases blockers and dependencies
* You can't start testing until the dependency has been fulfilled
* Re-creating a test environment for a real-world scenario might not be possible

---
<!-- _class: small-font -->

## Test Doubles

![bg right 70%](https://martinfowler.com/articles/practical-test-pyramid/unitTest.png)

* **Mocks** and **Stubs** are two different kinds of Test Doubles.

* You can use test doubles to replace objects you'd use in production with an implementation that helps you with testing.

* **Stubs** provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.

* **Mocks** are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.

> https://martinfowler.com/bliki/TestDouble.html

---

![bg right 70%](https://martinfowler.com/articles/practical-test-pyramid/unitTest.png)

## Disadvantages of Test Doubles

* No way of knowing whether it'll work in production
* No way of knowing when the dependency changes

---

![bg right 70%](https://martinfowler.com/articles/practical-test-pyramid/contract_tests.png)

As you often spread the consuming and providing services across different teams you find yourself in the situation where you have to clearly specify the interface between these services (the so called **contract**).

---
![bg right 70%](https://martinfowler.com/articles/practical-test-pyramid/contract_tests.png)

* Write a long and detailed interface specification (**the contract**)
* Implement the providing service according to the defined contract
* Throw the interface specification over the fence to the consuming team
* Wait until they implement their part of consuming the interface
* Run some large-scale manual system test to see if everything works
* Hope that both teams stick to the interface definition forever and don't screw up
* [OpenAPI][3] or [API Blueprint][4]

> https://martinfowler.com/articles/practical-test-pyramid.html

---

![bg right 70%](https://martinfowler.com/articles/practical-test-pyramid/cdc_tests.png)

Consumer-Driven Contract tests (CDC tests) let the **consumers drive the implementation of a contract**. Using CDC, consumers of an interface write tests that check the interface for all data they need from that interface. The consuming team then **publishes** these tests so that the publishing team can fetch and execute these tests easily. The providing team can now develop their API by running the CDC tests. Once all tests pass they know they have implemented everything the consuming team needs.

> https://martinfowler.com/articles/practical-test-pyramid.html
---

## Benefits

![bg right 90%](https://s3-ap-southeast-2.amazonaws.com/content-prod-529546285894/2019/07/screenshot-16.png)

* The ability to develop the consumer before the API

* The ability to drive out the requirements for your provider first, meaning you implement exactly and only what you need in the provider.

* The ability to immediately see which consumers will be broken if a change is made to the provider API.

> https://pactflow.io/what-are-the-benefits-of-contract-testing/

---

## Tooling

* https://pact.io/
* https://github.com/stoplightio/prism
* https://github.com/apiaryio/dredd/

---

## References

* [Consumer-Driven Contracts: A Service Evolution Pattern][1]
* [The Practical Test Pyramid][2]
* [OpenAPI Specification][3]
* [API Blueprint][4]
* [A Comprehensive Guide to Contract Testing APIs in a Service Oriented Architecture][5]

[1]: https://martinfowler.com/articles/consumerDrivenContracts.html
[2]: https://martinfowler.com/articles/practical-test-pyramid.html
[3]: https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.1.0.md
[4]: https://apiblueprint.org/
[5]: https://lirantal.medium.com/a-comprehensive-guide-to-contract-testing-apis-in-a-service-oriented-architecture-5695ccf9ac5a