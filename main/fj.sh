fj() {
    local output
    output=$(fastjump $@)
    code=$?
    if [[ $# -eq 1 && $1 != "-l" && $code -eq 0 ]]
    then cd $output
    elif [[ $1 = "-l" || $code -ne 0 ]]
    then echo $output
    else echo "succeeded"
    fi
}