# bdf2go.awk
# based on the book "Linuxから目覚めるぼくらのゲームボーイ" ISBN 978-4797325645 

BEGIN {
   print "// cp437 8x8"
   print "package font"
   print ""
   print "var Char8x8 = [][]byte{"
}

/^ENCODING/ {
    idx = $2
    found = 0
    while (getline > 0) { # Search BITMAP tag
        if ($0 != "BITMAP" )
            continue
        else {
            found = 1
            break
        }
    }
    if ( ! found ) die(idx, "Illegal BDF source")

    if ( idx >= 0 && idx <= 255) {
        printf( "    { ")
        for ( i = 0; i < 7; i++ ) {
            getline
            printf("0x%s, ", $0)
        }
        getline
        if ( idx != 255)
            printf("0x%s }, \n", $0)
        else
            printf("0x%s } }\n", $0)
    }
}

function die(char, msg) {
    print "Error: " msg " in ASCII code # " char
    exit
}