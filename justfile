set quiet

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

setup day:
    ./setup.sh {{day}}

everyday:
    bash -c 'for i in {1..25}; do just setup "$i"; done'

solve day part:
    go run . {{day}} {{part}}

test day:
    go test go-aoc-template/solutions/day$(printf %02d {{day}})
