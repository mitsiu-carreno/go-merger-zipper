package declarations

type Declarations struct{
	_id 			string `bson:"_id,omitempty"`
	ANIO			string `bson:"ANIO"`
	INDICE 			string `bson:"INDICE"`
	ACUSE			string `bson:"ACUSE"`
	FECHA			string `bson:"FECHA"`
	DEPENDENCIA		string `bson:"DEPENDENCIA"`
	DIA				string `bson:"DIA"`
	DECLARACION		string `bson:"DECLARACION"`
	SOURCE			string `bson:"SOURCE"`
	MES				string `bson:"MES"`
	ARCHIVO 		string `bson:"ARCHIVO"`
	NOMBRE			string `bson:"NOMBRE"`
	FOLDER			string `bson:"FOLDER"`
}