set quiet

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

solve day part:
    go run . {{day}} {{part}}

test day:
    go test go-aoc-template/solutions/day{{day}}
