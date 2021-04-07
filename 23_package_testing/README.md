# **Golang Testing Package**
- In the software industry, there are clear differences between manual testing and automated testing. Where manual testing is used to ensure that the software code performs as expected and it requires time and effort. Most of the manual testing includes checking log files, external services, and the database for errors. In a dissimilar fashion, Automated testing is, well, automated where certain software/code perform the testing as the user would do. Because automated testing is done using an automation tool, exploration tests take less time and more test scripts, while increasing the overall scope of the tests.
- In Golang, package testing is responsible for different types of testing maybe it is performance testing, parallel testing, functional testing, or any possible combination of these all.

- Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form.

- The testing package provides support for automated testing of the Golang code. To run any test function use “go test” command, which automates the execution of any function of the form  TestXxx(*testing.T), where Xxx must not start with any lowercase letter.

- t.Error* will report test failures but continue executing the test. t.Fatal*  will report test failures and stop the test immediately.

- t.Run enables running “subtests”, one for each table entry. These are shown separately when executing go test -v.

- Before you can write any unit tests, you need some code for your tests to analyze. In this step, you will build a small program that sums two integers. In the subsequent steps, you will use go test to test the program.



- [x] **Golang testing package**
- https://golang.org/pkg/testing/
- https://golang.org/doc/tutorial/add-a-test
- https://gobyexample.com/testing
- https://www.digitalocean.com/community/tutorials/- - - how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
- https://www.golang-book.com/books/intro/12
- https://stackoverflow.com/questions/39591037/golang-testing-programs-that-involves-time
- https://www.geeksforgeeks.org/overview-of-testing-package-in-golang/
- https://tutorialedge.net/golang/intro-testing-in-go/
