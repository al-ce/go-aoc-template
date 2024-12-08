set quiet
releasepath := './bin/aocgosolutions'
_dayf := "'$(printf %02d)'"

get year day:
    aocgofetch {{year}} {{day}} > inputs/{{day}}

solve day:
    mkdir -p bin && go build -o {{releasepath}} go-aoc-template && {{releasepath}} {{day}}

test day:
    go run $(printf "./solutions/day%02d/test%02d/test%02d.go" {{day}} {{day}} {{day}})

setup day:  # Create template for a specific day (!overwrites current file!)
    ./setup.sh {{day}}

everyday:  # Create template for every day (!overwrites all files!)
    bash -c 'for i in {1..25}; do just setup "$i"; done'
