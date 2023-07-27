package hk

//go:generate go run golang.org/x/tools/cmd/stringer -type=Category -linecomment -output=category_string.go

// Category represents an Apple HomeKit® accessory category.
type Category uint8

// All available accessory categories.
const (
	CategoryUnknown            Category = iota // Unknown
	CategoryOther                              // Other
	CategoryBridge                             // Bridge
	CategoryFan                                // Fan
	CategoryGarageDoorOpener                   // Garage Door Opener
	CategoryLightbulb                          // Lightbulb
	CategoryDoorLock                           // Door Lock
	CategoryOutlet                             // Outlet
	CategorySwitch                             // Switch
	CategoryThermostat                         // Thermostat
	CategorySensor                             // Sensor
	CategorySecuritySystem                     // Security System
	CategoryDoor                               // Door
	CategoryWindow                             // Window
	CategoryWindowCovering                     // Window Covering
	CategoryProgrammableSwitch                 // Programmable Switch
	_                                          // Range Extender (?)
	CategoryIPCamera                           // IP Camera
	CategoryVideoDoorbell                      // Video Doorbell
	CategoryAirPurifier                        // Air Purifier
	CategoryHeater                             // Heater
	CategoryAirConditioner                     // Air Conditioner
	CategoryHumidifier                         // Humidifier
	CategoryDehumidifier                       // Dehumidifier
	_                                          // Apple TV (?)
	_                                          // <reserved>
	_                                          // Speaker (?)
	_                                          // Airport (?)
	CategorySprinklers                         // Sprinklers
	CategoryFaucets                            // Faucets
	CategoryShowerSystems                      // Shower Systems
	_                                          // Television (?)
	_                                          // Remote Control (?)
)
