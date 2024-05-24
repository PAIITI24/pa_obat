package model

import (
	"encoding/json"
	"time"
)

type KategoriObat struct {
	ID               int       `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	NamaKategoriObat string    `json:"nama_kategori_obat" gorm:"type:varchar(50)"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	Obat             []Obat    `json:"kategori,omitempty" gorm:"many2many:kategorisasi"`
}

type Obat struct {
	ID            int            `json:"id" gorm:"primaryKey,autoIncrement"`
	NamaObat      string         `json:"nama_obat" gorm:"type:varchar(50)"`
	JumlahStok    uint           `json:"jumlah_stok" gorm:"type:int"`
	DosisObat     string         `json:"dosis_obat" gorm:"type:varchar(50)"`
	BentukSediaan string         `json:"bentuk_sediaan" gorm:"type:varchar(50)"`
	Harga         float32        `json:"harga" gorm:"type:float"`
	Gambar        string         `json:"gambar" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	KategoriObat  []KategoriObat `json:"kategori,omitempty" gorm:"many2many:kategorisasi"`
}

func (K *KategoriObat) MarshalJSON() ([]byte, error) {
	type Alias KategoriObat

	return json.Marshal(&struct {
		CreatedAt string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		*Alias
	}{
		CreatedAt: K.CreatedAt.Format("02/01/2006"),
		UpdatedAt: K.UpdatedAt.Format("02/01/2006"),
		Alias:     (*Alias)(K),
	})
}

func (O *Obat) MarshalJSON() ([]byte, error) {
	type Alias Obat
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	}{
		Alias:     (*Alias)(O),
		CreatedAt: O.CreatedAt.Format("02/01/2006"),
		UpdatedAt: O.UpdatedAt.Format("02/01/2006"),
	})
}
