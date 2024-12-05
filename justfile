set quiet
releasepath := './bin/aocgosolutions'

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

solve day part:
    mkdir -p bin && go build -o {{releasepath}} go-aoc-template && {{releasepath}} {{day}} {{part}}

test day:
    go test go-aoc-template/solutions/day$(printf %02d {{day}})

setup day:
    ./setup.sh {{day}}

everyday:
    bash -c 'for i in {1..25}; do just setup "$i"; done'
