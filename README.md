Every test follows the same flow: Set up the test environment (optional), feed the code under test input, capture the result, and compare it to the expected output. Note that inputs and results don't have to be arguments to a function. 

If the code under test is fetching data from a database then the input will be making sure the database contains appropriate test data (which may involve mocking at various levels). But, for our application, the common scenario of passing input arguments to a function and comparing the result to the function output is sufficient.

The way testing works is to include a _test.go file alongside the file under test, in the same folder and package.

Any function within this file that starts with Test will be tested. Such functions are expected to take a pointer to testing.T.
