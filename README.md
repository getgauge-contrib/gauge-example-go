# Gauge example in Go
[![Build Status](https://snap-ci.com/getgauge-contrib/gauge-example-go/branch/master/build_image)](https://snap-ci.com/getgauge-contrib/gauge-example-go/branch/master)

This is an example project for doing web automation testing with [Gauge](http://getgauge.io). This project tests some of the functionalities of the [active admin demo](https://github.com/getgauge/activeadmin-demo) app. This app is hosted as a Java WAR (with embedded Jetty).

## Running this example
The tests are run on Firefox by default.

### Prerequisites

This example requires the following softwares to run.
  * [Go](https://golang.org/)
  * [Gauge](http://getgauge.io/get-started/index.html)
  * Gauge Go plugin
    * can be installed using `gauge --install go`
  * Chrome driver
    * Download from [here](http://chromedriver.storage.googleapis.com/index.html)
    * On Mac: `brew install chromedriver`
  * Selenium server standalone
    * Available in `resources` directory of this repo. If you want a different version of it, it can be downloaded from [here](http://selenium-release.storage.googleapis.com/index.html)
  * Clone this repository in [GOPATH](https://apoorvam.github.io/golang/setup/2015/07/26/setting-up-golang-devbox.html).           
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

* [Specification](http://getgauge.io/documentation/user/current/specifications/index.html), [Scenario](http://getgauge.io/documentation/user/current/specifications/scenarios.html),  [Step](http://getgauge.io/documentation/user/current/specifications/steps.html), [Concepts](http://getgauge.io/documentation/user/current/specifications/concepts.html) and [Context Steps](http://getgauge.io/documentation/user/current/specifications/contexts.html)
* [Table parameters](http://getgauge.io/documentation/user/current/specifications/parameters.html#table-parameter)
* Using [External datasource (special param)](http://getgauge.io/documentation/user/current/specifications/parameters.html#special-parameters)
* Using [tags](http://getgauge.io/documentation/user/current/specifications/tags.html)
* Using Gauge with webdriver
