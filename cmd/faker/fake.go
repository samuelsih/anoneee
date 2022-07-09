package faker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/samuelsih/anoneee/utils"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

var (
	defaultID int
)

func init() {
	defaultID = 0
	rand.Seed(time.Now().Unix())
}

func Generate(data map[string]any) (map[string]any, error) {
	resultMap := make(map[string]any)

	for key, value := range data {
		fakeData, err := summonFakeData(value.(string))
		if err != nil {
			return nil, err
		}

		resultMap[key] = fakeData
	}

	return resultMap, nil
}

func summonFakeData(fakeType string) (any, error) {
	switch fakeType {
	case "default":
		defaultID++
		return defaultID, nil

	case "uuid":
		return uuid.New().String(), nil

	case "characters":
		return fake.Character(), nil

	case "city":
		return fake.City(), nil

	case "color":
		return fake.Color(), nil

	case "company":
		return fake.Company(), nil

	case "country":
		return fake.Country(), nil

	case "email":
		return fake.EmailAddress(), nil

	case "username":
		return fake.UserName(), nil

	case "firstname":
		return fake.FirstName(), nil

	case "lastname":
		return fake.LastName(), nil

	case "fullname":
		return fake.FullName(), nil

	case "gender":
		return fake.Gender(), nil

	case "age":
		return randInt(10, 60), nil

	case "jobtitle":
		return fake.JobTitle(), nil

	case "day":
		return fake.WeekDay(), nil

	case "month":
		return fake.Month(), nil

	case "year":
		return fake.Year(1970, 2050), nil

	case "sentence":
		return fake.Sentence(), nil

	case "sentences":
		return fake.Sentences(), nil

	case "paragraph":
		return fake.Paragraph(), nil

	case "paragraphs":
		return fake.Paragraphs(), nil

	case "street":
		return fake.Street(), nil

	case "streetaddress":
		return fake.StreetAddress(), nil

	case "title":
		return fake.Title(), nil

	case "industry":
		return fake.Industry(), nil

	case "brand":
		return fake.Brand(), nil

	case "bool":
		return randBool(), nil
	}

	return "", utils.CustomErrReturn(fmt.Sprintf("unknown faker type %v", fakeType))
}
