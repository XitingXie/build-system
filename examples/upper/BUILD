{
  "targets": [
    {
      "label": "//examples/upper:upper",
      "srcs": ["../greeting/hello.txt", "../greeting/world.txt"],
      "deps": ["//examples/greeting:greeting"],
      "cmd": "cat $SRCS | tr '[:lower:]' '[:upper:]' > $OUT",
      "outs": ["upper.out"]
    }
  ]
}
