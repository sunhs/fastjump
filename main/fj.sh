fj() {
    local args
    local argc
    local output
    local code
    if [ $# -eq 0 ]; then
        args="~"
        argc=1
    else
        args=$@
        argc=$#
    fi
    output=$(fastjump $args)
    code=$?
    if [[ $argc -eq 1 && $args != "-l" && $code -eq 0 ]]; then
        cd "$output"
    elif [[ $1 = "-l" || $code -ne 0 ]]; then
        echo "$output"
    else echo "Succeeded."
    fi
}

fjrc() {
    if [ ! $# -eq 0 ]; then
        echo "Wrong arguments. Use fjrc directly to remove the current dir from the DB."
        return 1
    fi
    cnt=$(fj -l | grep $(pwd) | wc -l)
    while [ ! $cnt -eq 0 ]; do
        i=$(fj -l | grep $(pwd) | head -n 1 | awk '{print $1}')
        fj -r $i
        cnt=$(( $cnt - 1 ))
    done
}