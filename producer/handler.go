package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/types"
)

// Handler handles the events coming from Worldpay Within
type Handler struct {
	ledGreen rpio.Pin
	ledRed   rpio.Pin
	ledBlue  rpio.Pin
	services map[int]*types.Service
}

func (handler *Handler) setup(services map[int]*types.Service, ignoreGPIO bool) error {

	if services == nil {

		return errors.New("Services must be set.")
	}

	handler.services = services
	handler.ledRed = rpio.Pin(2)
	handler.ledGreen = rpio.Pin(3)
	handler.ledBlue = rpio.Pin(4)

	gpioErr := rpio.Open()

	if gpioErr != nil {
		fmt.Println("Failed to open Raspberry Pi GPIO")

		if !ignoreGPIO {

			return gpioErr
		}

		return nil
	}
	fmt.Println("Did open Raspberry Pi GPIO")

	// Cleanup (defer until end)
	// rpio.Close()

	// Ensure pins are in output mode
	handler.ledGreen.Output()
	handler.ledRed.Output()
	handler.ledBlue.Output()
	fmt.Println("Did set GPIO pins to output type")

	// Turn of both LEDs, set the pins to low.
	handler.ledGreen.Low()
	handler.ledRed.Low()
	handler.ledBlue.Low()
	fmt.Println("Did set GPIO pins to low")

	return nil
}

// BeginServiceDelivery is called by Worldpay Within when a consumer wish to begin delivery of a service
func (handler *Handler) BeginServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsToSupply int) {

	fmt.Printf("BeginServiceDelivery. ServiceID = %d\n", serviceID)
	fmt.Printf("BeginServiceDelivery. UnitsToSupply = %d\n", unitsToSupply)
	fmt.Printf("BeginServiceDelivery. DeliveryToken = %+v\n", serviceDeliveryToken.Key)
	fmt.Println()
	svc := handler.services[serviceID]

	if &svc == nil {

		fmt.Printf("Service %d not found", serviceID)
		return
	}

	price := svc.Prices[1]

	durationSeconds := unitsToSupply * (unitsInTime[price.ID])
	fmt.Println("Warning, hardcoded price selection due to WPW design flaw. i.e. This event doesn't know what price was selected..")
	fmt.Printf("(%d) %s -> %s for %d %s\n", svc.ID, svc.Name, price.Description, durationSeconds, price.UnitDescription)

	fmt.Print("POWER ON ")
	switch svc.ID {

	case 1:
		fmt.Println("RED LED")
		handler.ledRed.High()
	case 2:
		fmt.Println("GREEN LED")
		handler.ledGreen.High()
	case 3:
		fmt.Println("BLUE LED")
		handler.ledBlue.High()
	default:
		fmt.Println("Unknown service id")
	}

	time.Sleep(time.Duration(durationSeconds) * time.Second)

	handler.EndServiceDelivery(serviceID, serviceDeliveryToken, unitsToSupply)
}

// EndServiceDelivery is called by Worldpay Within when a consumer wish to end delivery of a service
func (handler *Handler) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	fmt.Printf("EndServiceDelivery. ServiceID = %d\n", serviceID)
	fmt.Printf("EndServiceDelivery. UnitsReceived = %d\n", unitsReceived)
	fmt.Printf("EndServiceDelivery. DeliveryToken = %+v\n", serviceDeliveryToken.Key)
	fmt.Println()
	svc := handler.services[serviceID]

	if &svc == nil {

		fmt.Printf("Service %d not found", serviceID)
		return
	}

	fmt.Printf("%d - %s\n", svc.ID, svc.Name)

	fmt.Print("POWER OFF ")
	switch svc.ID {

	case 1:
		fmt.Println("RED LED")
		handler.ledRed.Low()
	case 2:
		fmt.Println("GREEN LED")
		handler.ledGreen.Low()
	case 3:
		fmt.Println("BLUE LED")
		handler.ledBlue.Low()
	default:
		fmt.Println("Unknown service id")
	}
}

// GenericEvent handles general events
func (handler *Handler) GenericEvent(name string, message string, data interface{}) error {

	return nil
}
