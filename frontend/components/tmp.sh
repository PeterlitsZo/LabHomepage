#!/usr/bin/env bash

nameList=(Paper Person Resource)
declare -A lowName=([Paper]=paper [Person]=person [Resource]=resource [xxxxxx]=XXXXXX)

for NAME in ${nameList[*]}; do
    for FOLDER in Create${NAME} ${NAME} ${NAME}List; do
        cd $FOLDER
        for FILE in $(ls); do
            NEWFILE=$(echo $FILE | sed 's/News/'$NAME'/g' | sed 's/news/'${lowName[$NAME]}'/g')
            if [[ "$FILE" != "$NEWFILE" ]]; then
                mv $FILE $NEWFILE
            fi
            sed -i'' -e "s/News/$NAME/g" $NEWFILE
            sed -i'' -e "s/news/${lowName[$NAME]}/g" $NEWFILE
        done
        cd ..
    done
done
