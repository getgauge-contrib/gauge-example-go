# Gauge example in Go

This is an example project for doing web automation testing with [Gauge](http://gauge.org). This project tests some of the functionalities of the [active admin demo](https://github.com/getgauge/activeadmin-demo) app. This app is hosted as a Java WAR (with embedded Jetty).

## Running this example
The tests are run on Firefox by default.

### Prerequisites

This example requires the following softwares to run.
  * [Go](https://golang.org/)
  * [Gauge](https://docs.gauge.org/getting_started/installing-gauge.html)
  * Gauge Go plugin
    * can be installed using `gauge --install go`
  * Chrome driver
    * Download from [here](http://chromedriver.storage.googleapis.com/index.html)
    * On Mac: `brew install chromedriver`
  * Selenium server standalone
    * Available in `resources` directory of this repo. If you want a different version of it, it can be downloaded from [here](http://selenium-release.storage.googleapis.com/index.html)
  * Clone this repository in [GOPATH](https://github.com/golang/go/wiki/GOPATH).           
  * `godep github.com/tebeka/selenium`       

### To set up the Selenium grid

```sh
java -jar <path_to_selenium_server_jar> -role hub
```
This uses port 4444 by default for its web interface.

To set up a node, run
```sh
java -jar <path_to_selenium_server_jar> -role node
```
Read more on setting selenium grid [here](https://github.com/apoorvam/gaugeGrid#setting-the-selenium-grid).

### Setting up the System Under Test (SUT)

* Download [activeadmin-demo.war](https://bintray.com/artifact/download/gauge/activeadmin-demo/activeadmin-demo.war)
* Bring up the SUT by executing the below command
```
java -jar activeadmin-demo.war
```
* The SUT should now be available at [http://localhost:8080/](http://localhost:8080)

### Execution

```
gauge specs
```
This uses Chrome as default browser for specs execution.

## Topics covered in the example

* [Specification](https://docs.gauge.org/latest/writing-specifications.html#specifications-spec), [Scenario](https://docs.gauge.org/latest/writing-specifications.html#scenario),  [Step](https://docs.gauge.org/latest/writing-specifications.html#step), [Concepts](https://docs.gauge.org/latest/writing-specifications.html#concepts) and [Context Steps](https://docs.gauge.org/latest/writing-specifications.html#longstart-context)
* [Table parameters](https://docs.gauge.org/latest/writing-specifications.html#table-parameters)
* Using [External datasource (special param)](https://docs.gauge.org/latest/writing-specifications.html#special-parameters)
* Using [tags](https://docs.gauge.org/latest/writing-specifications.html#tags)
* Using Gauge with webdriver
