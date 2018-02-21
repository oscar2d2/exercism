package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Record struct {
	Team  string
	Win   int
	Draw  int
	Loss  int
	Point int
}

type RecordList []Record

func (l RecordList) Len() int { return len(l) }
func (l RecordList) Less(i, j int) bool {
	return l[i].Point < l[j].Point || (l[i].Point == l[j].Point && l[i].Team > l[j].Team)
}
func (l RecordList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func Tally(in io.Reader, out io.Writer) error {
	var stat = make(map[string]*Record)
	input := make([]byte, 3000)
	n, err := in.Read(input)

	if err != nil {
		return err
	}

	input = input[:n]
	x := strings.TrimSpace(string(input))
	strs := strings.Split(x, "\n")

	for _, match := range strs {
		if strings.HasPrefix(match, "#") || match == "" {
			continue
		}
		details := strings.Split(match, ";")

		if len(details) != 3 {
			return errors.New("Invalid format")
		}

		team1 := details[0]
		team2 := details[1]
		result := details[2]

		if stat[team1] == nil {
			stat[team1] = &Record{Team: team1}
		}
		if stat[team2] == nil {
			stat[team2] = &Record{Team: team2}
		}

		switch result {
		case "draw":
			stat[team1].Draw++
			stat[team2].Draw++
		case "win":
			stat[team1].Win++
			stat[team2].Loss++
		case "loss":
			stat[team1].Loss++
			stat[team2].Win++
		default:
			return errors.New("Invalid result")
		}

	}

	board := make([]Record, len(stat))
	board = board[:0]

	for _, v := range stat {
		v.Point = v.Win*3 + v.Draw
		board = append(board, *v)
	}
	sort.Sort(sort.Reverse(RecordList(board)))
	_, err = out.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, v := range board {
		s := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", v.Team, v.Win+v.Draw+v.Loss, v.Win, v.Draw, v.Loss, v.Point)
		_, err = out.Write([]byte(s))
	}

	return nil
}
