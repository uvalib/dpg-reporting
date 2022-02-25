#!/usr/bin/env bash
#
# run application
#

# run the server
cd bin; ./dpg-reporting-svc -dbhost $DBHOST -dbport $DBPORT -dbname $DBNAME -dbuser $DBUSER -dbpass $DBPASS

# return the status
exit $?

#
# end of file
#
