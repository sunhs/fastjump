fj() {
    if [ $# -eq 1 ] && [[ "$1" =~ "^-.*" ]]; then
        if [ "$1" = "-" ] || [ "$1" -lt 0 ] 2>/dev/null; then
            cd "$1"
        else
            fj_cli $@
        fi
        return 0
    fi

    local dir
    local code

    dir="$(fj_cli jump $@)"
    code=$?
    if [ $code -eq 0 ]; then
        cd "$dir"
    else
        echo "no match"
    fi
}
