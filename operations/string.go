package operations

import (
	"fmt"
	"strconv"
	"strings"
)

func (m Matrix) String() string {
	x, y := m.GetIJ()
	s := make([]string, 0, (x+2)*(y+2))
	for _, row := range m {
		s = append(s, "\n|\t")
		for _, entry := range row {
			s = append(s, fmt.Sprintf("%s\t", strconv.Itoa(entry)))
		}
		s = append(s, "|\n")
	}
	return strings.Join(s, "")
}
