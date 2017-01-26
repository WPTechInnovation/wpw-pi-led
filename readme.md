# wpw-pi-led

Worldpay Within Pi LED demo

This demo demonstrates enabling a Raspberry Pi computer as a service producer containing 3 coloured LEDs that can be powered on once a payment is made. There is also a consumer that can be configured to seek the services of the producer and select 1 of 3 LEDs to be powered on.

# Software Setup

* Register an account with Worldpay for online payments [here](http://online.worldpay.com)
* Download/install dependencies, please run the following from the project root:
  * `go get ./...`

# Building the physical demo

* Parts list
 * 2 x Raspberry Pi 3 Model B. Other models should work but GPIO Pinout may differ.
 * 1 x Breadboard.
 * 3 x LED {Red, Green, Blue}.
 * 3 x Resistors. I used 1K Ohms (Brown, Black, Red).
 * 4 x Hookup wires. Preferably female on one end and male on the other.

* [Raspberry Pi 3 GPIO Pinout](https://www.myelectronicslab.com/wp-content/uploads/2016/06/raspbery-pi-3-gpio-pinout-40-pin-header-block-connector-.png)
* I used Pins { #06=GND, #03=Red, #05=Green, #07=Blue }
* // TODO Photo close-ups of build

# Usage

## Producer

* From the producer directory use `go build` to build the application
* Command line help can be found by using `producer -h`
* Run producer: `producer -wpservicekey <svc_key> -wpclientkey <client_key>`
* Note that `-ignoregpio` can be specified if you are not running a Raspberry Pi. Program will ignore errors setting up GPIO ports.

Once the producer is run it will setup the services, prices, PSP configuration etc. There should be enough information on screen to explain what has occurred. Some of the information may be relevant when starting the consumer.

## Consumer
* From the consumer directory use `go build` to build the application
* Command line help can be found by using `consumer -h`
* Run consumer `consumer -produceruuid <producer uuid> -serviceid <svc_id> -priceid <price_id> -unitquantity <quantity>`
* Note: the above parameters can be found by running the producer and looking at the producer overview on screen.
* Note: `-interactive` can be useful to step through the application as it runs. Press return when the program pauses to proceed to next section.
