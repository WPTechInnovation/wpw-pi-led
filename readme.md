# wpw-pi-led

Worldpay Within Pi LED demo

# Setup

* Register an account with Worldpay for online payments [here](http://online.worldpay.com)
* Download Worldpay Within SDK: `go get github.com/WPTechInnovation/worldpay-within-sdk`

# Usage
The following commands need to be run out of both the `producer` and `consumer` directories

* Build the demo apps (producer & consumer): `go build`
* Run producer: `producer -wpservicekey <svc_key> -wpclientkey <client_key>`
  * Take note of some info..
* Run consumer: `consumer`
* Once the consumer has notified of completed purchase and the LED has powered on, the transaction should be available at the Worldpay online payments dashboard

Note: The producer will run indefinitely until it is forced to quit. Once the consumer starts, it will scan for 15 seconds before selecting the service.  For this reason, it is advisable to launch the producer first.
