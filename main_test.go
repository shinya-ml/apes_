package main

import "testing"

func TestCalcPoint(t *testing.T) {
    cases := []struct{
        current_rank string
        kill_num int64
        rank_in_match int64
        got_point int64
    }{
        {
            current_rank: "bronze",
            kill_num: 5,
            rank_in_match: 1,
            got_point: 250,
        },
    }
}

