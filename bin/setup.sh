#!/bin/zsh

check_requrired_command() {
  for cmd in "$@"
  do
    if ! command -v $cmd &> /dev/null
    then
      echo -e "\033[0;31mplease install $cmd!\033[0m"
      exit 1
    fi
  done
}

install_npm_dependencies() {
  for dir in "$@"
  do
    (
      cd $dir
      pnpm install --frozen-lockfile
    ) &
  done
  wait
}

main() {
  check_requrired_command lefthook pnpm task docker
  lefthook install
  install_npm_dependencies frontend swagger
}

main
