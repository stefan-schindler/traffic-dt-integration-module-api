package api

type IDatagram interface {
	GetIndex() int
	SetIndex(index int)

	GetType() string
	SetType(newType string)

	GetTimestamp() string
	SetTimestamp(timestamp string)
}

// BaseDatagram Datagram
type BaseDatagram struct {
	Index     int    `json:"index"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	// Checksum TODO
}

func (datagram *BaseDatagram) GetIndex() int {
	return datagram.Index
}

func (datagram *BaseDatagram) SetIndex(index int) {
	datagram.Index = index
}

func (datagram *BaseDatagram) GetType() string {
	return datagram.Type
}

func (datagram *BaseDatagram) SetType(newType string) {
	datagram.Type = newType
}

func (datagram *BaseDatagram) GetTimestamp() string {
	return datagram.Timestamp
}

func (datagram *BaseDatagram) SetTimestamp(timestamp string) {
	datagram.Timestamp = timestamp
}

type ConnectDatagram struct {
	BaseDatagram
}

type SubscribeDatagram struct {
	BaseDatagram
	Content  string  `json:"content"`
	Interval float32 `json:"interval"`
}

type UnsubscribeDatagram struct {
	BaseDatagram
	Content string `json:"content"`
}

type AcknowledgeDatagram struct {
	BaseDatagram
	AcknowledgingIndex int `json:"acknowledgingIndex"`
}

type KeepAliveDatagram struct {
	BaseDatagram
}

type RequestAreaDatagram struct {
	BaseDatagram
}

type AreaDatagram struct {
	BaseDatagram
	TopLeft     PositionJSON `json:"topLeft"`
	BottomRight PositionJSON `json:"bottomRight"`
}

type WarningDatagram struct {
	WarnVehicleDatagram
	VehicleId int `json:"vehicleId"`
}

type WarnVehicleDatagram struct {
	BaseDatagram
	TimeToCollision   float32 `json:"timeToCollision"`
	CollisionType     string  `json:"collisionType"`
	CollisionSeverity string  `json:"collisionSeverity"`
}

type UpdateVehiclesDatagram struct {
	BaseDatagram
	Vehicles []UpdateVehiclesVehicle `json:"vehicles"`
}

type UpdateVehicleDatagram struct {
	BaseDatagram
	Vehicle UpdateVehicleVehicle `json:"vehicle"`
}

type UpdateVehiclesVehicle struct {
	Id           int          `json:"id"`
	Type         string       `json:"type"`
	Speed        float32      `json:"speed"`
	Acceleration float32      `json:"acceleration"`
	Heading      float32      `json:"heading"`
	Position     PositionJSON `json:"position"`
}

type UpdateVehicleVehicle struct {
	Vin          string       `json:"vin"`
	Speed        float32      `json:"speed"`
	Acceleration float32      `json:"acceleration"`
	Heading      float32      `json:"heading"`
	Position     PositionJSON `json:"position"`
}

type PositionJSON struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
