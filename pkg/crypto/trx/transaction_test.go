package trx

import (
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

func TestCreateTransaction(t *testing.T) {
	type args struct {
		sender   string
		receiver string
		amount   float64
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTransaction(tt.args.sender, tt.args.receiver, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignTransaction(t *testing.T) {
	type args struct {
		transaction []byte
		privateKey  *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignTransaction(tt.args.transaction, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroadcastTransaction(t *testing.T) {
	type args struct {
		signedTx []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BroadcastTransaction(tt.args.signedTx); (err != nil) != tt.wantErr {
				t.Errorf("BroadcastTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendTrx(t *testing.T) {
	type args struct {
		sender   *models.Wallet
		receiver string
		amount   float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendTrx(tt.args.sender, tt.args.receiver, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SendTrx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
