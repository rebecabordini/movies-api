package movie

type Movie struct {
	Id string `bson:"_id"`
	Title string `bson:"_title"`
	Director Person `bson:_director`
	Year string `bson:_year`
	Stars []Person `bson:,inline`
}

type Person struct {
	Id string `bson:_id`
	Name string `bson:name`
	Age int `bson:age`
}
