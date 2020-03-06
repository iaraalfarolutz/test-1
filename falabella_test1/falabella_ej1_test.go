package tlvToMap
import(
  "testing"
  "reflect"
)


func TestToMap(t *testing.T) {
  emptyMap := map[string]string{}
  correctMap := map[string]string{
    "A05": "AB398765UJ1",
    "N23": "00",
  }
  onlyFirstCorrect := map[string]string{
    "A05": "AB398765UJ1",
  }
  tt := []struct{
    tlvInput []byte
    resultMap map[string]string
    err error
    }{
      { []byte("A0511AB398765UJ1N230200"), correctMap, nil},
      { []byte("A0511AB398765UJ1N23020"), onlyFirstCorrect, InputError{"Index out of range"}},
      { []byte(""), emptyMap, InputError{"The input cant be empty"}},
      { []byte("A052"), emptyMap, InputError{"The input must provide at least the first 2 fields even if the value is empty"}},
      { []byte("F0511AB398765UJ1N230200"), emptyMap, InputError{"The first character of the type field should be A or N"}},
      { []byte("A0511AB398765UJ1F230200"), onlyFirstCorrect, InputError{"The first character of the type field should be A or N"}},
      { []byte("A05-2AB398765UJ1N230200"), emptyMap, InputError{"The lenght of the value must be positive"}},
      { []byte("A0511AB398765UJ1N23-200"), onlyFirstCorrect, InputError{"The lenght of the value must be positive"}},
      { []byte("N0511AB398765UJ1N230200"), emptyMap, InputError{"The value must match the value type defined on the first field"}},
      { []byte("A0511AB398765UJ1N23020A"), onlyFirstCorrect, InputError{"The value must match the value type defined on the first field"}},
    }

    for _, tc := range tt {
      myMap, err := ToMap(tc.tlvInput)
      if !reflect.DeepEqual(myMap, tc.resultMap){
        t.Errorf("ToMap(%v): expected %+q, actual %+q", string(tc.tlvInput), tc.resultMap, myMap)
      }
      if err != tc.err {
        t.Errorf("ToMap(%d): expected %d, actual %d", tc.tlvInput, tc.err, err)
      }

    }
  }
