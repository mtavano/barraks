package barraks

// Item represents a basic informations for every barrak ite
type Item struct {
  ID        string `json:"id" db:"id"`
  Name      string `json:"name" db:"name"`
  ImgURL    string `json:"imag_url,omitempty" db:"img_url"`
  Unit      string `json:"unit" db:"unit"`
  Stock     int `json:"stock" db:"stock"`
  MinStock  int `json:"min_stock" db:"min_stock"`
}
