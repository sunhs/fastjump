fj() {
    local output="$(fastjump ${@})"
    if [[ "$#" -eq 1 && "$1" != "-l" ]]
    then cd "${output}"
    else echo "${output}"
    fi
}
