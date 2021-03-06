package pl_validator

import (
	"testing"
)

func TestValidatePesel_correctData(t *testing.T) {
	peselValidator := PeselValidator{}
	for _, data := range correctData {
		err := peselValidator.validate(data)
		if err != nil {
			t.Errorf("Err should be nil")
		}
	}
}

func TestValidatePesel_dataTooLong(t *testing.T) {
	expectedError := "pesel number is too long"
	peselValidator := PeselValidator{}
	for _, data := range tooLongData {
		err := peselValidator.validate(data)
		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	}
}
func TestValidatePesel_dataTooShort(t *testing.T) {
	expectedError := "pesel number is too short"
	peselValidator := PeselValidator{}
	for _, data := range tooShortData {
		err := peselValidator.validate(data)
		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	}
}

func TestValidatePesel_incorrectMonth(t *testing.T) {
	expectedError := "incorrect month. (range 1-12)"
	peselValidator := PeselValidator{}
	for _, data := range incorrectMonthData {
		err := peselValidator.validate(data)
		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	}
}
func TestValidatePesel_incorrectDay(t *testing.T) {
	expectedError := "incorrect day. (range 1-31)"
	peselValidator := PeselValidator{}
	for _, data := range incorrectDayData {
		err := peselValidator.validate(data)
		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	}
}

func TestValidatePesel_correct(t *testing.T) {
	peselValidator := PeselValidator{}
	for _, data := range correctData {
		result := peselValidator.validate(data)
		if result != nil {
			t.Errorf(result.Error())
		}
	}
}

func TestValidatePesel_incorrect(t *testing.T) {
	peselValidator := PeselValidator{}
	for _, data := range incorrectData {
		result := peselValidator.validate(data)
		if result != nil && result.Error() != "wrong pesel controlsum" {
			t.Errorf(result.Error())
		}
	}
}

var incorrectMonthData = []string{
	"00007713909",
	"00140098765",
}
var incorrectDayData = []string{
	"00017713909",
	"00120098765",
}

var tooLongData = []string{
	"999999999994567",
	"123456789012",
}
var tooShortData = []string{
	"0",
	"9999991234",
}
var correctData = []string{
	"78060258156",
	"85041245377",
	"81011884898",
	"48062598914",
	"47040668557",
	"86090524275",
	"62013029384",
	"77022577258",
	"78040618488",
	"96080641594",
	"99041164297",
	"57112266455",
	"91021964561",
	"54101461552",
	"86082942342",
	"99041827235",
	"46100189223",
	"80090189629",
	"61082819739",
	"70060674891",
	"93040279222",
	"88111298765",
	"86010634499",
	"86073199872",
	"54071284254",
	"57021549759",
	"83122377319",
	"45120356268",
	"60102479768",
	"50011035733",
	"79010161997",
	"82061418567",
	"51020388793",
	"85030938147",
	"64011539241",
	"49061532994",
	"63102894447",
	"67032892481",
	"94070258584",
	"64062848965",
	"53022742825",
	"71092685763",
	"54121199741",
	"89031315426",
	"60092268988",
	"82091941598",
	"65070965729",
	"58080435261",
	"48040571386",
	"51020592114",
	"91042336426",
	"77090334555",
	"80121214489",
	"85082911383",
	"55092836787",
	"72081952387",
	"52071676585",
	"72041682749",
	"74051223763",
	"94052666268",
	"60043064328",
	"65050484741",
	"89050882286",
	"91042548616",
	"51121384296",
	"60100594881",
	"88121583561",
	"86092282391",
	"92042369227",
	"51110577337",
	"52091956995",
	"88080868244",
	"65102914284",
	"87053015384",
	"82091794868",
	"64072482193",
	"46032475751",
	"77021745573",
	"49092012544",
	"71021224142",
	"60090218664",
	"58120565967",
	"81113057316",
	"65013121346",
	"71063023477",
	"69090282457",
	"51062563859",
	"74092237886",
	"61020714382",
	"74041551869",
	"55101224792",
	"98102963552",
	"73080553272",
	"89851202672",
	"67822806603",
	"13841202928",
	"25853008762",
	"61872106072",
	"45421500818",
	"58503009446",
	"21470801056",
	"53491800514",
	"75411008664",
}

var incorrectData = []string{
	"16080641594",
	"46100489223",
	"02321164427",
	"61082859739",
}
