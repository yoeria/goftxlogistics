package main

import (
	"reflect"
	"testing"
)

func Test_strategies_Update(t *testing.T) {
	type fields struct {
		EMA      bool
		SMA      bool
		STOCHRSI bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "TEST UpdateConfiguration: STRATEGIES",
			fields: fields{EMA: true, SMA: true, STOCHRSI: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &strategies{
				EMA:      tt.fields.EMA,
				SMA:      tt.fields.SMA,
				STOCHRSI: tt.fields.STOCHRSI,
			}
			s.Update()
		})
	}
}

func TestTrade_MarshalBinary(t *testing.T) {
	type fields struct {
		DatetimeEnter string
		DatetimeExit  string
		Enter         float64
		Exit          float64
		Net           float64
		NetPercent    float64
		Wallet        float64
		WalletNet     float64
		Fees          float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Trade{
				DatetimeEnter: tt.fields.DatetimeEnter,
				DatetimeExit:  tt.fields.DatetimeExit,
				Enter:         tt.fields.Enter,
				Exit:          tt.fields.Exit,
				Net:           tt.fields.Net,
				NetPercent:    tt.fields.NetPercent,
				Wallet:        tt.fields.Wallet,
				WalletNet:     tt.fields.WalletNet,
				Fees:          tt.fields.Fees,
			}
			got, err := tr.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("Trade.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trade.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrade_UnmarshalBinary(t *testing.T) {
	type fields struct {
		DatetimeEnter string
		DatetimeExit  string
		Enter         float64
		Exit          float64
		Net           float64
		NetPercent    float64
		Wallet        float64
		WalletNet     float64
		Fees          float64
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Trade{
				DatetimeEnter: tt.fields.DatetimeEnter,
				DatetimeExit:  tt.fields.DatetimeExit,
				Enter:         tt.fields.Enter,
				Exit:          tt.fields.Exit,
				Net:           tt.fields.Net,
				NetPercent:    tt.fields.NetPercent,
				Wallet:        tt.fields.Wallet,
				WalletNet:     tt.fields.WalletNet,
				Fees:          tt.fields.Fees,
			}
			if err := tr.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Trade.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
