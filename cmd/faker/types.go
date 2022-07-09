package faker

var fakerTypes = []string{
	"default",
	"uuid",
	"characters",
	"city",
	"color",
	"company",
	"country",
	"email",
	"username",
	"firstname",
	"lastname",
	"fullname",
	"gender",
	"age",
	"jobtitle",
	"day",
	"month",
	"year",
	"sentence",
	"sentences",
	"paragraph",
	"paragraphs",
	"street",
	"streetaddress",
	"title",
	"industry",
	"brand",
	"bool",
}

func IsFakerType(fakeType string) bool {
	for _, fakerType := range fakerTypes {
		if fakerType == fakeType {
			return true
		}
	} 

	return false
}

func NotBrokenID(key, value string) bool {	
	if key == "id" && (value == "default" || value == "uuid") {
		return true
	} 

	if key != "id" && (value == "default" || value == "uuid") {
		return false
	} 

	return true
}