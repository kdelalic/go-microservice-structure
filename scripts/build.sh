#!/bin/bash

# This is the list of all makefiles that we've already built. We don't include the 
# root makefile by default.
BUILT=`readlink -e ${PWD}/Makefile`
echo "${BUILT}" > builtlist

# Main build function. Takes a directory as input.
build () {
  echo "Build input = $1"
  DIRNAME=`echo $1`
  MKFILE=`echo "${DIRNAME}/Makefile"`
  SLASHES=${PWD//[^\/]/}

  # Try walking up the path until we find a makefile.
  for (( n=${#SLASHES}; n>0; --n )); do
    if [ -f $MKFILE ]; then
      echo "Found Makefile in ${DIRNAME}"
      break
    else
      DIRNAME="${DIRNAME}/.."
      MKFILE=`echo "${DIRNAME}/Makefile"`
    fi
  done
  
  # Get the full path of the makefile.
  MKFILE_FULL=`readlink -e ${MKFILE}`
  
  # Build only if it's not on our list of built makefiles.
  BUILT=$(<builtlist)
  if [[ $BUILT != *"${MKFILE_FULL}"* ]]; then
    echo "Build ${DIRNAME} (${MKFILE_FULL})"
    
    # Main build command.
    INCLUDE_MAKEFILE=$MKFILE make release
   
    # Add item to our list of built makefiles.
    BUILT=`echo "${BUILT};${MKFILE_FULL}"`
    echo "${BUILT}" > builtlist
  else
    echo "Skip ${MKFILE_FULL} (already built, or root)"
  fi
}

# Prebuild function. Takes a file as input.
processline () {
  line=$1
  echo "Process ${line}"

  if [[ $line == vendor* ]] || [[ $line == pkg* ]]; then
    # The changed line is common. We will iterate through all dirs except hidden ones, 'vendor',
    # and 'pkg' to see if build is necessary.
    find . -type d -not -path "*/\.*" | grep -v 'vendor' | grep -v 'pkg' | while read item; do
      # Get the current package's full list of golang dependencies (recursive).
      PKG_GODEPS=`go list -f '{{ .Deps }}' $item`
      
      if [ $? -eq 0 ]; then
        LINE_DIR=`dirname $line`

        # See if this package has a dependency with the changed file. If so, proceed with build.
        if [[ $PKG_GODEPS = *"${LINE_DIR}"* ]]; then
          echo "'${item}' has a dependency with '${LINE_DIR}'"
          # Remove the './' prefix (output from 'find' command).
          TO_BUILD=`echo "${item}" | cut -c 3-`
          build $TO_BUILD
        fi
      fi
    done
  else
    # The changed line belongs to a service.
    TO_BUILD=`dirname $line`
    build $TO_BUILD
  fi
}

git remote set-branches --add origin main
git fetch

COMMIT_RANGE="origin/main"

if [ "$TRAVIS_BRANCH" == "main" ]; then
  COMMIT_RANGE=$TRAVIS_COMMIT_RANGE
fi

echo "Commit range ${COMMIT_RANGE}"

git diff --name-only $COMMIT_RANGE | while read line; do
  processline $line
  echo "-"
done