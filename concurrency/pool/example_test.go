// This sample program demonstrates how to use the pool package
// to share a simulated set of database connections.
package pool

import (
	"io"
	"reflect"
	"testing"

	"github.com/goinaction/code/chapter7/patterns/pool"
)

func Test_dbConnection_Close(t *testing.T) {
	tests := []struct {
		name    string
		dbConn  *dbConnection
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dbConn.Close(); (err != nil) != tt.wantErr {
				t.Errorf("dbConnection.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    io.Closer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("createConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example()
		})
	}
}

func Test_performQueries(t *testing.T) {
	type args struct {
		query int
		p     *pool.Pool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			performQueries(tt.args.query, tt.args.p)
		})
	}
}
