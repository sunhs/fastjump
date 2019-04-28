fj() {
    local output="$(fastjump ${@})"
    if [ -d ${output} ]
    then
        cd "${output}"
    else
        echo "${output}"
    fi
}