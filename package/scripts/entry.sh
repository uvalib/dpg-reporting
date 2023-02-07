#!/usr/bin/env bash
#
# run application
#

# run the server
cd bin; ./dpg-reporting-svc   \
   -tsadmin $TS_ADMIN_URL     \
   -tsadmin $TS_API_URL       \
   -dbhost  $DBHOST           \
   -dbport  $DBPORT           \
   -dbname  $DBNAME           \
   -dbuser  $DBUSER           \
   -dbpass  $DBPASS

# return the status
exit $?

#
# end of file
#
