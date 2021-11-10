// Code generated by ./internal/cmd/genschema. DO NOT EDIT.

package buttplug

// User-set id for the message. 0 denotes system message and is reserved.
type ID int // [0, 4294967295]

// Signifies the server encountered an error while processing the message
// indicated by the id.
type Error struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID           ID      `json:"Id"`
	ErrorMessage string  `json:"ErrorMessage"`
	ErrorCode    float64 `json:"ErrorCode"`
}

// Used for connection/application testing. Causes server to echo back the
// string sent. Sending string of 'Error' will result in a server error.
type Test struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// String to be echo'd back from server. Setting this to 'Error' will cause
	// an error to be thrown.
	TestString string `json:"TestString"`
}

// Name of the device
type DeviceName string

// Index used for referencing the device in device messages.
type DeviceIndex int // [0, max]

// A list of the messages a device will accept on this server implementation.
type DeviceMessages []string

// Attributes for device message that have no attributes.
type NullMessageAttributes struct {
}

// Number of features on device.
type FeatureCount int // [1, max]

// Specifies granularity of each feature on the device.
type StepCount []int

// Attributes for device messages.
type GenericMessageAttributes struct {
	// Number of features on device.
	FeatureCount FeatureCount `json:"FeatureCount,omitempty"`
	// Specifies granularity of each feature on the device.
	StepCount []int `json:"StepCount,omitempty"`
}

// Attributes for raw device messages.
type RawMessageAttributes struct {
	Endpoints []string `json:"Endpoints,omitempty"`
}

// A list of the messages a device will accept on this server implementation.
type DeviceMessagesEx struct {
	// Attributes for device message that have no attributes.
	StopDeviceCmd NullMessageAttributes `json:"StopDeviceCmd,omitempty"`
	// Attributes for device messages.
	VibrateCmd GenericMessageAttributes `json:"VibrateCmd,omitempty"`
	// Attributes for device messages.
	LinearCmd GenericMessageAttributes `json:"LinearCmd,omitempty"`
	// Attributes for device messages.
	RotateCmd GenericMessageAttributes `json:"RotateCmd,omitempty"`
	// Attributes for device message that have no attributes.
	LovenseCmd NullMessageAttributes `json:"LovenseCmd,omitempty"`
	// Attributes for device message that have no attributes.
	VorzeA10CycloneCmd NullMessageAttributes `json:"VorzeA10CycloneCmd,omitempty"`
	// Attributes for device message that have no attributes.
	KiirooCmd NullMessageAttributes `json:"KiirooCmd,omitempty"`
	// Attributes for device message that have no attributes.
	SingleMotorVibrateCmd NullMessageAttributes `json:"SingleMotorVibrateCmd,omitempty"`
	// Attributes for device message that have no attributes.
	FleshlightLaunchFW12Cmd NullMessageAttributes `json:"FleshlightLaunchFW12Cmd,omitempty"`
	// Attributes for device message that have no attributes.
	BatteryLevelCmd NullMessageAttributes `json:"BatteryLevelCmd,omitempty"`
	// Attributes for device message that have no attributes.
	RSSILevelCmd NullMessageAttributes `json:"RSSILevelCmd,omitempty"`
	// Attributes for raw device messages.
	RawReadCmd RawMessageAttributes `json:"RawReadCmd,omitempty"`
	// Attributes for raw device messages.
	RawWriteCmd RawMessageAttributes `json:"RawWriteCmd,omitempty"`
	// Attributes for raw device messages.
	RawSubscribeCmd RawMessageAttributes `json:"RawSubscribeCmd,omitempty"`
	// Attributes for raw device messages.
	RawUnsubscribeCmd RawMessageAttributes `json:"RawUnsubscribeCmd,omitempty"`
}

// List of all available devices known to the system.
type DeviceList struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Array of device ids and names.
	Devices []struct {
		// Name of the device
		DeviceName DeviceName `json:"DeviceName"`
		// Index used for referencing the device in device messages.
		DeviceIndex    DeviceIndex `json:"DeviceIndex"`
		DeviceMessages interface{} `json:"DeviceMessages"`
	} `json:"Devices"`
}

// Used for non-direct-reply messages that can only be sent from server to
// client, using the reserved system message Id of 0.
type SystemID int // [0, 0]

// Notifies client that a device of a certain type has been added to the server.
type DeviceAdded struct {
	// Used for non-direct-reply messages that can only be sent from server to
	// client, using the reserved system message Id of 0.
	ID SystemID `json:"Id"`
	// Name of the device
	DeviceName DeviceName `json:"DeviceName"`
	// Index used for referencing the device in device messages.
	DeviceIndex    DeviceIndex `json:"DeviceIndex"`
	DeviceMessages interface{} `json:"DeviceMessages"`
}

// Stops the all actions currently being taken by a device.
type StopDeviceCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
}

// Request for server to stream log messages of a certain level to client.
type RequestLog struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
}

// Log message from the server.
type Log struct {
	// Used for non-direct-reply messages that can only be sent from server to
	// client, using the reserved system message Id of 0.
	ID SystemID `json:"Id"`
	// Log message from server.
	LogMessage string `json:"LogMessage"`
}

// Request server version, and relay client name.
type RequestServerInfo struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Name of the client software.
	ClientName string `json:"ClientName"`
	// Message template version of the client software.
	MessageVersion int `json:"MessageVersion,omitempty"`
}

// Server version information, in Major.Minor.Build format.
type ServerInfo struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Name of the server. Can be 0-length.
	ServerName string `json:"ServerName"`
	// Message template version of the server software.
	MessageVersion int `json:"MessageVersion"`
	// Major version of server.
	MajorVersion int `json:"MajorVersion,omitempty"`
	// Minor version of server.
	MinorVersion int `json:"MinorVersion,omitempty"`
	// Build version of server.
	BuildVersion int `json:"BuildVersion,omitempty"`
	// Maximum time (in milliseconds) the server will wait between ping messages
	// from client before shutting down.
	MaxPingTime int `json:"MaxPingTime"`
}

// Sends speed and position command to the Fleshlight Launch Device denoted by
// the device index.
type FleshlightLaunchFW12Cmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Speed at which to move to designated position.
	Speed int `json:"Speed"`
	// Position to which to move Fleshlight.
	Position int `json:"Position"`
}

// Sends a command string to a Lovense device. Command string will be verified
// by sender.
type LovenseCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Command to send to Lovense device.
	Command string `json:"Command"`
}

// Sends a vibrate command to a device that supports vibration.
type SingleMotorVibrateCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Device vibration speed (floating point, 0 < x < 1), stepping will be
	// device specific.
	Speed float64 `json:"Speed"`
}

// Sends a raw byte string to a Kiiroo Onyx/Pearl device.
type KiirooCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Command to send to Kiiroo device.
	Command string `json:"Command"`
}

// Request a raw byte array from a device. Should only be used for
// testing/development.
type RawReadCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Endpoint (from device config file) from which the data was retrieved.
	Endpoint string `json:"Endpoint"`
	// Amount of data to read from device, 0 to exhaust whatever is in immediate
	// buffer
	Length int `json:"Length"`
}

// Sends a raw byte array to a device. Should only be used for
// testing/development.
type RawWriteCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Endpoint (from device config file) to send command to.
	Endpoint string `json:"Endpoint"`
	// Raw byte string to send to device.
	Data []int `json:"Data"`
}

// Subscribe to an endpoint on a device to receive raw info back.
type RawSubscribeCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Endpoint (from device config file) from which the data was retrieved.
	Endpoint string `json:"Endpoint"`
}

// Unsubscribe to an endpoint on a device.
type RawUnsubscribeCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Endpoint (from device config file) from which the data was retrieved.
	Endpoint string `json:"Endpoint"`
}

// Raw byte array received from a device. Should only be used for
// testing/development.
type RawReading struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Endpoint (from device config file) from which the data was retrieved.
	Endpoint string `json:"Endpoint"`
	// Raw byte string received from device.
	Data []int `json:"Data"`
}

// Sends a raw byte string to a Kiiroo Onyx/Pearl device.
type VorzeA10CycloneCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Rotation speed command for the Cyclone.
	Speed int `json:"Speed"`
}

// Sends a vibrate command to a device that supports vibration.
type VibrateCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Device vibration speeds (floating point, 0 < x < 1) keyed on vibrator
	// number, stepping will be device specific.
	Speeds []struct {
		// Vibrator number.
		Index int `json:"Index"`
		// Vibration speed (floating point, 0 < x < 1), stepping will be device
		// specific.
		Speed float64 `json:"Speed"`
	} `json:"Speeds"`
}

// Sends a rotate command to a device that supports rotation.
type RotateCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Device rotation speeds (floating point, 0 < x < 1) keyed on rotator
	// number, stepping will be device specific.
	Rotations []struct {
		// Rotator number.
		Index int `json:"Index"`
		// Rotation speed (floating point, 0 < x < 1), stepping will be device
		// specific.
		Speed float64 `json:"Speed"`
	} `json:"Rotations"`
}

// Sends a linear movement command to a device that supports linear movements.
type LinearCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Device linear movement times (milliseconds) and positions (floating
	// point, 0 < x < 1) keyed on linear actuator number, stepping will be
	// device specific.
	Vectors []struct {
		// Linear actuator number.
		Index int `json:"Index"`
		// Linear movement time in milliseconds.
		Duration float64 `json:"Duration"`
		// Linear movement position (floating point, 0 < x < 1), stepping will be
		// device specific.
		Position float64 `json:"Position"`
	} `json:"Vectors"`
}

// Requests that a BatteryLevel be retreived.
type BatteryLevelCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
}

// Returns a BatteryLevel read from a device.
type BatteryLevelReading struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// Battery Level
	BatteryLevel float64 `json:"BatteryLevel"`
}

// Requests that a RSSI level be retreived.
type RSSILevelCmd struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
}

// Returns a BatteryLevel read from a device.
type RSSILevelReading struct {
	// User-set id for the message. 0 denotes system message and is reserved.
	ID ID `json:"Id"`
	// Index used for referencing the device in device messages.
	DeviceIndex DeviceIndex `json:"DeviceIndex"`
	// RSSI Level
	RSSILevel float64 `json:"RSSILevel"`
}