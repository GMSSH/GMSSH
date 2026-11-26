#!/bin/bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

Install() {
    # do-something
    echo 'Successify'
}

Uninstall() {
    # do-something
}

CleanData() {
    # do-something
}

action=$GM_APP_OPT_ACTION
clean_data=$GM_UNINSTALL_CLEAN
if [ "$action" == 'install' ]; then
    Install
else
    Uninstall
    if [ "$clean_data" == 'true' ]; then
        CleanData
    fi
fi
