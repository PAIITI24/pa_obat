package model

import (
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
	DosisObat     string         `json:"dosis_obat" gorm:"type:varchar(50)"`
	BentukSediaan string         `json:"bentuk_sediaan" gorm:"type:varchar(50)"`
	Harga         float32        `json:"harga" gorm:"type:float"`
	Gambar        string         `json:"gambar" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	KategoriObat  []KategoriObat `json:"kategori,omitempty" gorm:"many2many:kategorisasi"`
}
