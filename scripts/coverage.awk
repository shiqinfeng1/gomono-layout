#!/usr/bin/env awk

{
  print $0
  if (match($0, /^total:/)) {
    sub(/%/, "", $NF);
    print("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
    printf("test coverage is %s%%(quality gate is %s%%)\n", $NF, target)
    if (int($NF) < target) {
      printf("!!! test coverage does not meet expectations: %d%%, please add test cases !!!\n", target)
    }
    print("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
  }
}
