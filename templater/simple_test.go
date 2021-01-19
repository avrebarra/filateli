package templater_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/avrebarra/filateli/postman"
	"github.com/avrebarra/filateli/templater"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSimple(t *testing.T) {
	type args struct {
		cfg templater.ConfigSimple
	}
	tests := []struct {
		name    string
		args    args
		wantT   templater.Templater
		wantErr bool
	}{
		{
			name:    "ok",
			args:    args{},
			wantT:   &templater.Simple{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, err := templater.NewSimple(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSimple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("NewSimple() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestSimple_Apply(t *testing.T) {
	// shared
	ctx := context.Background()
	e, err := templater.NewSimple(templater.ConfigSimple{})
	require.Nil(t, err)

	co, err := loadSampleCollection()
	require.Nil(t, err)

	codoc, err := loadSampleCollectionDoc()
	require.Nil(t, err)

	// test
	type args struct {
		ctx  context.Context
		coll postman.Collection
	}
	tests := []struct {
		name    string
		args    args
		wantDoc []byte
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx:  ctx,
				coll: co,
			},
			wantDoc: codoc,
			wantErr: false,
		},
		{
			name:    "empty input",
			args:    args{},
			wantDoc: []byte(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDoc, err := e.Apply(tt.args.ctx, tt.args.coll)
			if (err != nil) != tt.wantErr {
				t.Errorf("Simple.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("COK")
			fmt.Println(string(gotDoc))
			fmt.Println("COK")
			assert.Equal(t, string(gotDoc), string(tt.wantDoc))
		})
	}
}

func loadSampleCollection() (c postman.Collection, err error) {
	content, err := ioutil.ReadFile("../fixtures/mimpiyangtetapsemu.postman_collection.json")
	if err != nil {
		return
	}

	err = c.ParseFrom(bytes.NewBuffer(content))
	if err != nil {
		return
	}

	return
}

func loadSampleCollectionDoc() (d []byte, err error) {
	d, err = ioutil.ReadFile("../fixtures/mimpiyangtetapsemu.md")
	if err != nil {
		return
	}
	return
}
