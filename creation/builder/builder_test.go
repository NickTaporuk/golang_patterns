package builder

import (
	"reflect"
	"testing"
)

//func TestBuilderPattern(t *testing.T) {
//
//	manufacturingComplex := ManufacturingDirector{}
//
//	carBuilder := &CarBuilder{}
//
//	manufacturingComplex.SetBuilder(carBuilder)
//	manufacturingComplex.Construct()
//
//	car := carBuilder.Build()
//
//	if car.Wheels != 4 {
//		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
//	}
//	if car.Structure != "Car" {
//		t.Errorf("Structure on a car must be 'Car' and was %s\n",
//			car.Structure)
//	}
//
//	if car.Seats != 5 {
//		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
//	}
//
//	bikeBuilder := &BikeBuilder{}
//
//	manufacturingComplex.SetBuilder(bikeBuilder)
//	manufacturingComplex.Construct()
//
//	motorbike := bikeBuilder.GetVehicle()
//	motorbike.Seats = 1
//
//	if motorbike.Wheels != 2 {
//		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",
//			motorbike.Wheels)
//	}
//
//	if motorbike.Structure != "Motorbike" {
//		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",
//			motorbike.Structure)
//	}
//}

func TestManufacturingDirector_SetBuilder(t *testing.T) {
	type fields struct {
		builder BuildProcess
	}
	type args struct {
		b BuildProcess
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &ManufacturingDirector{
				builder: tt.fields.builder,
			}
			f.SetBuilder(tt.args.b)
		})
	}
}

func TestManufacturingDirector_Construct(t *testing.T) {
	type fields struct {
		builder BuildProcess
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &ManufacturingDirector{
				builder: tt.fields.builder,
			}
			f.Construct()
		})
	}
}

func TestCarBuilder_SetWheels(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarBuilder{
				v: tt.fields.v,
			}
			if got := c.SetWheels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetWheels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_SetSeats(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarBuilder{
				v: tt.fields.v,
			}
			if got := c.SetSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_SetStructure(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarBuilder{
				v: tt.fields.v,
			}
			if got := c.SetStructure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.SetStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_GetVehicle(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   VehicleProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarBuilder{
				v: tt.fields.v,
			}
			if got := c.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarBuilder_Build(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   VehicleProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarBuilder{
				v: tt.fields.v,
			}
			if got := c.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetWheels(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BikeBuilder{
				v: tt.fields.v,
			}
			if got := b.SetWheels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetWheels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetSeats(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BikeBuilder{
				v: tt.fields.v,
			}
			if got := b.SetSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_SetStructure(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   BuildProcess
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BikeBuilder{
				v: tt.fields.v,
			}
			if got := b.SetStructure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.SetStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBikeBuilder_GetVehicle(t *testing.T) {
	type fields struct {
		v VehicleProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   VehicleProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BikeBuilder{
				v: tt.fields.v,
			}
			if got := b.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BikeBuilder.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}
