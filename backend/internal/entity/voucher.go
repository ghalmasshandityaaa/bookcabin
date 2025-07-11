package entity

type Voucher struct {
	ID           int          `json:"id" gorm:"column:id;type:int;autoIncrement;primaryKey"`
	CrewID       string       `json:"crew_id" gorm:"column:crew_id;type:text;not null"`
	CrewName     string       `json:"crew_name" gorm:"column:crew_name;type:text;not null"`
	FlightNumber string       `json:"flight_number" gorm:"column:flight_number;type:text;not null"`
	FlightDate   string       `json:"flight_date" gorm:"column:flight_date;type:date;not null"`
	AircraftType AircraftType `json:"aircraft_type" gorm:"column:aircraft_type;type:text;not null"`
	Seat1        string       `json:"seat1" gorm:"column:seat1;type:text;not null"`
	Seat2        string       `json:"seat2" gorm:"column:seat2;type:text"`
	Seat3        string       `json:"seat3" gorm:"column:seat3;type:text"`
	CreatedAt    string       `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

type CreateVoucherProps struct {
	CrewID       string       `json:"crew_id" validate:"required,numeric"`
	CrewName     string       `json:"crew_name" validate:"required,min=2,max=100"`
	FlightNumber string       `json:"flight_number" validate:"required,min=3,max=100"`
	FlightDate   string       `json:"flight_date" validate:"required,is-valid-date"`
	AircraftType AircraftType `json:"aircraft_type" validate:"required,oneof=ATR 'Airbus 320' 'Boeing 737 Max'"`
	Seat1        string       `json:"seat1" validate:"required,min=2,max=10"`
	Seat2        string       `json:"seat2" validate:"omitempty,min=2,max=10"`
	Seat3        string       `json:"seat3" validate:"omitempty,min=2,max=10"`
}

func NewVoucher(props *CreateVoucherProps) *Voucher {
	return &Voucher{
		CrewID:       props.CrewID,
		CrewName:     props.CrewName,
		FlightNumber: props.FlightNumber,
		FlightDate:   props.FlightDate,
		AircraftType: props.AircraftType,
		Seat1:        props.Seat1,
		Seat2:        props.Seat2,
		Seat3:        props.Seat3,
	}
}

func (a *Voucher) TableName() string {
	return "vouchers"
}
