package dao

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB, _ = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:30306)/blog?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

func TestDao_CountTag(t *testing.T) {
	type fields struct {
		engine *gorm.DB
	}
	type args struct {
		name  string
		state uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "biny",
			args:    args{name: "标签1", state: 0},
			fields:  fields{testDB},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dao{
				engine: tt.fields.engine,
			}
			got, err := d.CountTag(tt.args.name, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dao.CountTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dao.CountTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
