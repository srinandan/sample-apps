grep -rli ./ -e '$2' *.yaml | xargs -i@ sed -i 's/$2/$2/g' @
