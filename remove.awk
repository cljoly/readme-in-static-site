#!/usr/bin/awk -f

BEGIN { removing = 0; inserting = 0 }

/^<!--+ remove -+->$/ {
    removing = 1
    next
}
/^<!--+ end_remove -+->$/ {
    removing = 0
    next
}

/^<!--+ insert$/ {
    inserting = 1
    next
}
/^end_insert -+->$/ {
    inserting = 0
    next
}

removing == 0 { print $0 }
