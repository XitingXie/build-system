{
  "targets": [
    {
      "label": "//examples/greeting:hello",
      "srcs": ["hello.txt"],
      "deps": [],
      "cmd": "cp $SRCS $OUT",
      "outs": ["hello.out"]
    },
    {
      "label": "//examples/greeting:world",
      "srcs": ["world.txt"],
      "deps": [],
      "cmd": "cp $SRCS $OUT",
      "outs": ["world.out"]
    },
    {
      "label": "//examples/greeting:greeting",
      "srcs": ["hello.txt", "world.txt"],
      "deps": [
        "//examples/greeting:hello",
        "//examples/greeting:world"
      ],
      "cmd": "cat $SRCS > $OUT",
      "outs": ["greeting.out"]
    }
  ]
}
