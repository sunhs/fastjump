fj() {
    local args
    local argc
    local output
    local code
    if [ $# -eq 0 ]
    then
        args="~"
        argc=1
    else
        args=$@
        argc=$#
    fi
    output=$(fastjump $args)
    code=$?
    if [[ $argc -eq 1 && $args != "-l" && $code -eq 0 ]]
    then cd "$output"
    elif [[ $1 = "-l" || $code -ne 0 ]]
    then echo "$output"
    else echo "succeeded"
    fi
}
