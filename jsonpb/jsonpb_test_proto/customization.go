package jsonpb

import "fmt"

// MarshalJSON prefix flattens object to a string with the prefix Prefix.
func (s *SelfMarshaller) MarshalJSON() (b []byte, err error) {
	b = []byte(fmt.Sprintf("\"Prefix:%s\"", s.ValueToMod))
	return
}
