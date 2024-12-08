set quiet
releasepath := './bin/aocgosolutions'

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

solve day:
    mkdir -p bin && go build -o {{releasepath}} go-aoc-template && {{releasepath}} {{day}}

test day:
    go test -v -count=1 go-aoc-template/solutions/day$(printf %02d {{day}}) | grep "^Part"

setup day:
    ./setup.sh {{day}}

everyday:
    bash -c 'for i in {1..25}; do just setup "$i"; done'
