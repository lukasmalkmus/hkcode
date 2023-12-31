// Code generated by "stringer -type=Category -linecomment -output=category_string.go"; DO NOT EDIT.

package hk

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CategoryUnknown-0]
	_ = x[CategoryOther-1]
	_ = x[CategoryBridge-2]
	_ = x[CategoryFan-3]
	_ = x[CategoryGarageDoorOpener-4]
	_ = x[CategoryLightbulb-5]
	_ = x[CategoryDoorLock-6]
	_ = x[CategoryOutlet-7]
	_ = x[CategorySwitch-8]
	_ = x[CategoryThermostat-9]
	_ = x[CategorySensor-10]
	_ = x[CategorySecuritySystem-11]
	_ = x[CategoryDoor-12]
	_ = x[CategoryWindow-13]
	_ = x[CategoryWindowCovering-14]
	_ = x[CategoryProgrammableSwitch-15]
	_ = x[CategoryIPCamera-17]
	_ = x[CategoryVideoDoorbell-18]
	_ = x[CategoryAirPurifier-19]
	_ = x[CategoryHeater-20]
	_ = x[CategoryAirConditioner-21]
	_ = x[CategoryHumidifier-22]
	_ = x[CategoryDehumidifier-23]
	_ = x[CategorySprinklers-28]
	_ = x[CategoryFaucets-29]
	_ = x[CategoryShowerSystems-30]
}

const (
	_Category_name_0 = "UnknownOtherBridgeFanGarage Door OpenerLightbulbDoor LockOutletSwitchThermostatSensorSecurity SystemDoorWindowWindow CoveringProgrammable Switch"
	_Category_name_1 = "IP CameraVideo DoorbellAir PurifierHeaterAir ConditionerHumidifierDehumidifier"
	_Category_name_2 = "SprinklersFaucetsShower Systems"
)

var (
	_Category_index_0 = [...]uint8{0, 7, 12, 18, 21, 39, 48, 57, 63, 69, 79, 85, 100, 104, 110, 125, 144}
	_Category_index_1 = [...]uint8{0, 9, 23, 35, 41, 56, 66, 78}
	_Category_index_2 = [...]uint8{0, 10, 17, 31}
)

func (i Category) String() string {
	switch {
	case i <= 15:
		return _Category_name_0[_Category_index_0[i]:_Category_index_0[i+1]]
	case 17 <= i && i <= 23:
		i -= 17
		return _Category_name_1[_Category_index_1[i]:_Category_index_1[i+1]]
	case 28 <= i && i <= 30:
		i -= 28
		return _Category_name_2[_Category_index_2[i]:_Category_index_2[i+1]]
	default:
		return "Category(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
