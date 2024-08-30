package barang

import (
	"time"
	"gorm.io/datatypes"
)

type Barang struct {
	Id     				 int       			`json:"id"`
	NamaBarang     string    			`json:"nama_barang"`
	JumlahStok     int       			`json:"jumlah_stok"`
	NomorSeri      string    			`json:"nomor_seri"`
	AdditionalInfo datatypes.JSON `json:"additional_info"`
	GambarBarang   string    			`json:"gambar_barang"`
	CreatedAt      time.Time 			`json:"created_at"`
	CreatedBy      string    			`json:"created_by"`
	UpdatedAt      *time.Time 		`json:"updated_at"`
	UpdatedBy      *string    		`json:"updated_by"`
}

func (Barang) TableName() string {
	return "barang"
}