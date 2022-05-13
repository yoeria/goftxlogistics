package strategies

import (
	"reflect"
	"testing"

	talib "github.com/markcheno/go-talib"
	"github.com/sdcoffey/techan"
)

func TestStratData_calcStochRSI(t *testing.T) {
	type fields struct {
		techanSeries *techan.TimeSeries
		inReal       []float64
		timeFrame    int
		kPeriod      int
		dPeriod      int
		maType       talib.MaType
	}
	tests := []struct {
		name   string
		fields fields
		want   []float64
		want1  []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StratData{
				techanSeries: tt.fields.techanSeries,
				inReal:       tt.fields.inReal,
				timeFrame:    tt.fields.timeFrame,
				kPeriod:      tt.fields.kPeriod,
				dPeriod:      tt.fields.dPeriod,
				maType:       tt.fields.maType,
			}
			got, got1 := s.calcStochRSI()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StratData.calcStochRSI() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StratData.calcStochRSI() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStratData_calcFastStochastic(t *testing.T) {
	type fields struct {
		techanSeries *techan.TimeSeries
		inReal       []float64
		timeFrame    int
		kPeriod      int
		dPeriod      int
		maType       talib.MaType
	}
	tests := []struct {
		name   string
		fields fields
		want   techan.Indicator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StratData{
				techanSeries: tt.fields.techanSeries,
				inReal:       tt.fields.inReal,
				timeFrame:    tt.fields.timeFrame,
				kPeriod:      tt.fields.kPeriod,
				dPeriod:      tt.fields.dPeriod,
				maType:       tt.fields.maType,
			}
			if got := s.calcFastStochastic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StratData.calcFastStochastic() = %v, want %v", got, tt.want)
			}
		})
	}
}
