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

func (handler *Handler) setup(services map[int]*types.Service) error {

	handler.ledGreen = rpio.Pin(2)
	handler.ledRed = rpio.Pin(3)
	handler.ledYellow = rpio.Pin(4)

	if err := rpio.Open(); err != nil {

		return err
	}

	// Cleanup (defer until end)
	// rpio.Close()

	// Ensure pins are in output mode
	handler.ledGreen.Output()
	handler.ledRed.Output()
	handler.ledYellow.Output()

	// Turn of both LEDs, set the pins to low.
	handler.ledGreen.Low()
	handler.ledRed.Low()
	handler.ledYellow.Low()

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
	fmt.Printf("%s -> %s for %d %s\n", svc.Name, price.Description, durationSeconds, price.UnitDescription)

	fmt.Println("POWER ON")
	time.Sleep(time.Duration(durationSeconds) * time.Second)
	fmt.Println("POWER OFF")
}

// EndServiceDelivery is called by Worldpay Within when a consumer wish to end delivery of a service
func (handler *Handler) EndServiceDelivery(serviceID int, serviceDeliveryToken types.ServiceDeliveryToken, unitsReceived int) {

	fmt.Printf("EndServiceDelivery. ServiceID = %d\n", serviceID)
	fmt.Printf("EndServiceDelivery. UnitsReceived = %d\n", unitsReceived)
	fmt.Printf("EndServiceDelivery. DeliveryToken = %+v\n", serviceDeliveryToken)

	if serviceID == 1 {

		handler.ledGreen.Low()

	} else if serviceID == 2 {
}

// GenericEvent handles general events
func (handler *Handler) GenericEvent(name string, message string, data interface{}) error {

	return nil
}
