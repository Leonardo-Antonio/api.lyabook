package entity

type Digital struct {
	Format string `bson:"format,omitempty" json:"format,omitempty" xml:"format,omitempty"`
	Src    string `bson:"src,omitempty" json:"src,omitempty" xml:"src,omitempty"`
	//Details []string `bson:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
}

type Fisico struct {
	Format string `bson:"format,omitempty" json:"format,omitempty" xml:"format,omitempty"`
	Log    string `bson:"log,omitempty" json:"log,omitempty" xml:"log,omitempty"`
	Lat    string `bson:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	Stock  int    `bson:"stock,omitempty" json:"stock,omitempty" xml:"stock,omitempty"`
	//Details []string `bson:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	Active bool `bson:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
}

type Format struct {
	Digital Digital `bson:"digital,omitempty" json:"digital,omitempty" xml:"digital,omitempty"`
	Fisico  Fisico  `bson:"fisico,omitempty" json:"fisico,omitempty" xml:"fisico,omitempty"`
}
