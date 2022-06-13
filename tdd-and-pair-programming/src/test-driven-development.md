---
marp: true
title: Test Driven Development (TDD)
paginate: true
theme: default
backgroundColor: rgb(247 249 251)
style: |
  section { font-size: 1.4em; }
  h2 { 
    display: inline-block;
    padding-left: 0.6em;
    padding-top: 0.1em;
    background: black;
    color: #ffffff;
  }
  a { color: #5386e4; }
  strong { color: #E3056C; }
  blockquote {
    position: absolute;
    bottom: 1em;
    left: 3em;
    font-size: 60%;
    /* margin-top: auto; */
  }
  section.interval-slide {
    text-align: center;
    color: white;
    font-size: 250%;
    font-weight: bold;
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
<!-- _paginate: false -->
<!-- _footer: June 2022 -->

<style scoped>
  h1 { color: #F5F5F5; }
  p { color: #DCDCDC; }
</style>

![bg 50% opacity:.4](https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.1.1/svgs/solid/laptop-code.svg)

# Test Driven Development (TDD)

An Introduction by Kevin Denver

---

## Red-Green-Refactor: The Building Blocks of TDD

Test-driven development follows a three-phase process:

* **Red**. We write a failing test (including possible compilation failures). We run the test suite to verify the failing tests.

* **Green**. We write just enough production code to make the test green. We run the test suite to verify this.

* **Refactor**. We remove any code smells. These may be due to duplication, hardcoded values, or improper use of language idioms. If we break any tests during refactoring, we prioritize getting them back to green before exiting this phase.

![bg right 75%](./assets/TDD.png)

---

## Benefits of TDD

* Writing the tests first requires you to really consider what do you want from the code
* Fast feedback
* Creates a detailed specification
* Tells you whether your last change (or refactoring) has broken previously working code
* Allows the design to evolve and adapt to your changing understanding of the problem
* Forces radical simplification of the code, you will only write codes in response to the requirements of the tests
* Shortens the development Time to Market
* Reduces the time it takes to refactor the code or fix bugs
* Cuts development costs
* Improves quality
* Test Driven Development gives programmers the confidence to change the larger architecture of an application when adding new functionality. Without the flexibility of TDD, developers frequently add new functionality by virtually bolting it to the existing application without true integration – clearly, this can cause problems down the road

---

## Pitfalls of TDD

* Tests need to be treated with the same respect and care as production code
* Excessive use of mocking can hinder your ability to refactor
  * See [Mockists Are Dead. Long Live Classicists][7]
* Testing the wrong thing or at the wrong level/abstraction

![bg right 80%](https://martinfowler.com/articles/practical-test-pyramid/testPyramid.png)

---

## References

* [Learning Test-Driven Development][1]
* [Mocks Aren't Stubs][2]
* [Scientific Research Into Pair Programming][3]
* [The Art of Agile Development, 2nd Edition][4]
* [Detroit and London Schools of Test-Driven Development][5]
* [Test Driven Development: By Example][6]
* [Mockists Are Dead. Long Live Classicists][7]
* [The Practical Test Pyramid][8]

[1]: https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/
[2]: https://martinfowler.com/articles/mocksArentStubs.html
[3]: https://tuple.app/pair-programming-guide/scientific-research-into-pair-programming
[4]: https://learning.oreilly.com/library/view/the-art-of/9781492080688/
[5]: https://blog.devgenius.io/detroit-and-london-schools-of-test-driven-development-3d2f8dca71e5
[6]: https://learning.oreilly.com/library/view/test-driven-development/0321146530/
[7]: https://www.thoughtworks.com/insights/blog/mockists-are-dead-long-live-classicists
[8]: https://martinfowler.com/articles/practical-test-pyramid.html