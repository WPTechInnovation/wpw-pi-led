package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/psp/onlineworldpay"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// Application flags
var flagWPServiceKey string
var flagWPClientKey string

// Application Vars
var wpw wpwithin.WPWithin
var wpwHandler Handler

const (
	redDescr    string = "Turn on the red LED"
	greenDescr  string = "Turn on the green LED"
	yellowDescr string = "Turn on the yellow LED"
	oneSecond   string = "One second"
	oneMinute   string = "One minute"
)

func init() {

	flag.StringVar(&flagWPServiceKey, "wpservicekey", "", "Worldpay service key")
	flag.StringVar(&flagWPClientKey, "wpclientkey", "", "Worldpay client key")
}

func main() {

	flag.Parse()

	if strings.EqualFold(flagWPClientKey, "") {
		fmt.Println("Flag wpclientkey is required")
		os.Exit(1)
	} else if strings.EqualFold(flagWPServiceKey, "") {
		fmt.Println("Flag wpservicekey is required")
		os.Exit(1)
	}

	_wpw, err := wpwithin.Initialise("WPW Pi LED", "Worldpay Within Pi LED Demo")
	wpw = _wpw

	errCheck(err, "WorldpayWithin Initialise")

	doSetupServices()

	// wpwhandler accepts callbacks from worldpay within when service delivery begin/end is required.
	err = wpwHandler.setup()
	errCheck(err, "wpwHandler setup")
	wpw.SetEventHandler(&wpwHandler)

	pspConfig := map[string]string{
		onlineworldpay.CfgMerchantClientKey:  flagWPClientKey,
		onlineworldpay.CfgMerchantServiceKey: flagWPServiceKey,
	}
	err = wpw.InitProducer(pspConfig)

	errCheck(err, "Init producer")

	err = wpw.StartServiceBroadcast(0) // 0 = no timeout

	errCheck(err, "start service broadcast")

	// run the app until it is closed
	runForever()
}

func doSetupServices() {

	////////////////////////////////////////////
	// Green LED
	////////////////////////////////////////////

	svcGreenLed, err := types.NewService()
	errCheck(err, "Create new service - Green LED")
	svcGreenLed.ID = 1
	svcGreenLed.Name = "Big LED"
	svcGreenLed.Description = greenDescr

	priceGreenLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price - green led second")

	priceGreenLedSecond.Description = greenDescr
	priceGreenLedSecond.ID = 1
	priceGreenLedSecond.UnitDescription = oneSecond
	priceGreenLedSecond.UnitID = 1
	priceGreenLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       10,
		CurrencyCode: "GBP",
	}

	svcGreenLed.AddPrice(*priceGreenLedSecond)

	priceGreenLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price - green led minute")

	priceGreenLedMinute.Description = greenDescr
	priceGreenLedMinute.ID = 2
	priceGreenLedMinute.UnitDescription = oneMinute
	priceGreenLedMinute.UnitID = 2
	priceGreenLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       40, /* WOAH! This is minor units so means just 40p */
		CurrencyCode: "GBP",
	}

	svcGreenLed.AddPrice(*priceGreenLedMinute)

	err = wpw.AddService(svcGreenLed)
	errCheck(err, "Add service - green led")

	////////////////////////////////////////////
	// Red LED
	////////////////////////////////////////////

	svcRedLed, err := types.NewService()
	errCheck(err, "New service - red led")

	svcRedLed.ID = 2
	svcRedLed.Name = "Red LED"
	svcRedLed.Description = redDescr

	priceRedLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price - red led second")

	priceRedLedSecond.Description = redDescr
	priceRedLedSecond.ID = 3
	priceRedLedSecond.UnitDescription = oneSecond
	priceRedLedSecond.UnitID = 1
	priceRedLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       5,
		CurrencyCode: "GBP",
	}

	svcRedLed.AddPrice(*priceRedLedSecond)

	priceRedLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price - red led minute")

	priceRedLedMinute.Description = redDescr
	priceRedLedMinute.ID = 4
	priceRedLedMinute.UnitDescription = oneMinute
	priceRedLedMinute.UnitID = 2
	priceRedLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       20,
		CurrencyCode: "GBP",
	}

	svcRedLed.AddPrice(*priceRedLedMinute)

	err = wpw.AddService(svcRedLed)
	errCheck(err, "Add service - red led")

	////////////////////////////////////////////
	// Yellow LED
	////////////////////////////////////////////

	svcYellowLed, err := types.NewService()
	errCheck(err, "New service - yellow led")

	svcYellowLed.ID = 3
	svcYellowLed.Name = "Yellow LED"
	svcYellowLed.Description = yellowDescr

	priceYellowLedSecond, err := types.NewPrice()
	errCheck(err, "Create new price - yellow led second")

	priceYellowLedSecond.Description = yellowDescr
	priceYellowLedSecond.ID = 1
	priceYellowLedSecond.UnitDescription = oneSecond
	priceYellowLedSecond.UnitID = 1
	priceYellowLedSecond.PricePerUnit = &types.PricePerUnit{
		Amount:       5,
		CurrencyCode: "GBP",
	}

	svcYellowLed.AddPrice(*priceYellowLedSecond)

	priceYellowLedMinute, err := types.NewPrice()
	errCheck(err, "Create new price - yellow led minute")

	priceYellowLedMinute.Description = yellowDescr
	priceYellowLedMinute.ID = 2
	priceYellowLedMinute.UnitDescription = oneMinute
	priceYellowLedMinute.UnitID = 2
	priceYellowLedMinute.PricePerUnit = &types.PricePerUnit{
		Amount:       20,
		CurrencyCode: "GBP",
	}

	svcYellowLed.AddPrice(*priceYellowLedMinute)

	err = wpw.AddService(svcYellowLed)
	errCheck(err, "Add service - yellow led")
}

func errCheck(err error, hint string) {

	if err != nil {
		fmt.Printf("Did encounter error during: %s", hint)
		fmt.Println(err.Error())
		fmt.Println("Quitting...")
		os.Exit(1)
	}
}

func runForever() {

	done := make(chan bool)
	fnForever := func() {
		for {
			time.Sleep(time.Second * 10)
		}
	}

	go fnForever()

	<-done // Block forever
}
