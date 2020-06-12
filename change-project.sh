#!/bin/bash

if [ "$#" -ne 3 ]; then
   echo "requires three parameter to run: change-project mac old-value new-value"
   exit 1
fi
 
if [ ! "$1" = "mac" -o "$1" = "lin" ]; then
 echo "invalid type: type can be mac or lin"
fi 

if [[ "$1" == "mac" ]]
  then     
    find . -type f -name '*.yaml' -exec sed -i '' s/$2/$3/ {} +
else
  grep -rli ./ -e '$2' *.yaml | xargs -i@ sed -i 's/$2/$3/g' @
fi
